.POSIX
.PHONY: *

wireguard:
	mkdir -p .wg
	rm -rf .wg/sin .wg/sin.pub
	rm -rf .wg/iad .wg/iad.pub
	rm -rf .wg/cdg .wg/cdg.pub
	nix-shell --run 'wg genkey | tee .wg/sin | wg pubkey > .wg/sin.pub'
	nix-shell --run 'wg genkey | tee .wg/iad | wg pubkey > .wg/iad.pub'
	nix-shell --run 'wg genkey | tee .wg/cdg | wg pubkey > .wg/cdg.pub'
