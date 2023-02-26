pkgs.mkShell {
  buildInputs = with pkgs; [
    cargo
    rustc
    git-ignore
    cargo-watch
    rust-analyzer
    clippy
    darwin.apple_sdk.frameworks.Security
  ];

  RUST_SRC_PATH = "${pkgs.rust.packages.stable.rustPlatform.rustLibSrc}";
}
