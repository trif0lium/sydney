package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/netip"
	"os"
	"path/filepath"

	"golang.zx2c4.com/wireguard/conn"
	"golang.zx2c4.com/wireguard/device"
	"golang.zx2c4.com/wireguard/tun/netstack"
)

func main() {
	hostIPs, err := net.LookupIP("host.docker.internal")
	if err != nil {
		log.Panic(err)
	}

	hostIP := hostIPs[0].String()
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

	localPrivateKeyBase64, err := os.ReadFile(filepath.Join(cwd, ".wg/sin"))
	if err != nil {
		log.Panic(err)
	}

	localPrivateKey, err := base64.StdEncoding.DecodeString(string(localPrivateKeyBase64))
	if err != nil {
		log.Panic(err)
	}

	remotePublicKeyBase64, err := os.ReadFile(filepath.Join(cwd, ".wg/iad.pub"))
	if err != nil {
		log.Panic(err)
	}

	remotePublicKey, err := base64.StdEncoding.DecodeString(string(remotePublicKeyBase64))
	if err != nil {
		log.Panic(err)
	}

	wgConf := bytes.NewBuffer(nil)
	fmt.Fprintf(wgConf, "private_key=%s\n", hex.EncodeToString(localPrivateKey))
	fmt.Fprintf(wgConf, "listen_port=58120\n")

	fmt.Fprintf(wgConf, "public_key=%s\n", hex.EncodeToString(remotePublicKey))
	fmt.Fprintf(wgConf, "allowed_ip=%s\n", "192.168.4.28/32")
	fmt.Fprintf(wgConf, "endpoint=%s\n", hostIP+":58121")
	fmt.Fprintf(wgConf, "persistent_keepalive_interval=25\n")

	if err := dev.IpcSetOperation(bufio.NewReader(wgConf)); err != nil {
		log.Panic(err)
	}

	dev.Up()

	listener, err := tnet.ListenTCP(&net.TCPAddr{Port: 80})
	if err != nil {
		log.Panic(err)
	}
	defer listener.Close()

	for {
		local, err := listener.Accept()
		if err != nil {
			log.Panic(err)
		}

		remote, err := net.Dial("tcp", "127.0.0.1:80")
		if err != nil {
			log.Panic(err)
		}

		runTunnel(local, remote)
	}
}

func runTunnel(local, remote net.Conn) {
	defer local.Close()
	defer remote.Close()
	done := make(chan struct{}, 2)

	go func() {
		io.Copy(local, remote)
		done <- struct{}{}
	}()

	go func() {
		io.Copy(remote, local)
		done <- struct{}{}
	}()

	<-done
}
