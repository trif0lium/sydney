.POSIX:
.PHONY: *

wireguard:
	mkdir -p .wg
	rm -rf .wg/sin .wg/sin.pub .wg/sin_iad.pub .wg/sin_cdg.pub
	rm -rf .wg/iad .wg/iad.pub .wg/iad_sin.pub .wg/iad_cdg.pub
	rm -rf .wg/cdg .wg/cdg.pub .wg/cdg_sin.pub .wg/cdg_iad.pub
	nix-shell --run 'wg genkey | tee .wg/sin | wg pubkey > .wg/sin.pub'
	nix-shell --run 'wg genkey | tee .wg/iad | wg pubkey > .wg/iad.pub'
	nix-shell --run 'wg genkey | tee .wg/cdg | wg pubkey > .wg/cdg.pub'

build: wireguard
	rm -rf bin
	mkdir -p bin
	go build -o bin/egressproxy -ldflags '-w -s' egressproxy/main.go
	go build -o bin/ingressproxy -ldflags '-w -s' ingressproxy/main.go
	go build -o bin/testserver -ldflags '-w -s' testserver/main.go
