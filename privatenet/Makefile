.POSIX:
.PHONY: *

wireguard:
	mkdir -p .wg
	rm -rf .wg/sin .wg/sin.pub .wg/sin_iad.pub .wg/sin_cdg.pub
	rm -rf .wg/iad .wg/iad.pub .wg/iad_sin.pub .wg/iad_cdg.pub
	rm -rf .wg/cdg .wg/cdg.pub .wg/cdg_sin.pub .wg/cdg_iad.pub
	nix-shell --run 'wg genkey | tee .wg/sin | wg pubkey > .wg/sin.pub'
	nix-shell --run 'wg pubkey < .wg/sin > .wg/sin_iad.pub'
	nix-shell --run 'wg pubkey < .wg/sin > .wg/sin_cdg.pub'
	nix-shell --run 'wg genkey | tee .wg/iad | wg pubkey > .wg/iad.pub'
	nix-shell --run 'wg pubkey < .wg/iad > .wg/iad_sin.pub'
	nix-shell --run 'wg pubkey < .wg/iad > .wg/iad_cdg.pub'
	nix-shell --run 'wg genkey | tee .wg/cdg | wg pubkey > .wg/cdg.pub'
	nix-shell --run 'wg pubkey < .wg/cdg > .wg/cdg_sin.pub'
	nix-shell --run 'wg pubkey < .wg/cdg > .wg/cdg_iad.pub'
