package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net"
	"net/netip"
	"os"
	"path/filepath"
	"sync"

	"golang.zx2c4.com/wireguard/conn"
	"golang.zx2c4.com/wireguard/device"
	"golang.zx2c4.com/wireguard/tun/netstack"
)

func main() {
	tun, tnet, err := netstack.CreateNetTUN(
		[]netip.Addr{netip.MustParseAddr("192.168.4.29")},
		[]netip.Addr{},
		device.DefaultMTU,
	)
	if err != nil {
		log.Panic(err)
	}

	dev := device.NewDevice(tun, conn.NewDefaultBind(), device.NewLogger(device.LogLevelVerbose, ""))

	cwd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	localPrivateKey, err := os.ReadFile(filepath.Join(cwd, ".wg/sin"))
	if err != nil {
		log.Panic(err)
	}

	localPublicKey, err := os.ReadFile(filepath.Join(cwd, ".wg/sin.pub"))
	if err != nil {
		log.Panic(err)
	}

	wgConf := bytes.NewBuffer(nil)
	fmt.Fprintf(wgConf, "private_key=%s\n", hex.EncodeToString(localPrivateKey))
	fmt.Fprintf(wgConf, "public_key=%s\n", hex.EncodeToString(localPublicKey))
	fmt.Fprintf(wgConf, "listen_port=58120\n")
	fmt.Fprintf(wgConf, "allowed_ip=%s\n", "192.168.4.28/32")
	fmt.Fprintf(wgConf, "persistent_keepalive_interval=25\n")

	if err := dev.IpcSetOperation(bufio.NewReader(wgConf)); err != nil {
		log.Panic(err)
	}

	dev.Up()

	listener, err := tnet.ListenTCP(&net.TCPAddr{Port: 80})
	if err != nil {
		log.Panic(err)
	}

	source, err := listener.Accept()
	if err != nil {
		log.Panic(err)
	}
	defer source.Close()

	go func() {
		target, err := net.Dial("tcp", ":80")
		if err != nil {
			log.Print(err)
			return
		}
		defer target.Close()

		wg := &sync.WaitGroup{}

		wg.Add(2)

		copyFunc := func(dst net.Conn, src net.Conn) {
			defer wg.Done()
			io.Copy(dst, src)
		}

		go copyFunc(target, source)
		go copyFunc(source, target)

		wg.Wait()
	}()
}
