{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-23.11";
  };

  outputs = {nixpkgs, ...}: let
    system = "x86_64-linux";
    pkgs = nixpkgs.legacyPackages.${system};
  in {
    devShells.${system}.default = pkgs.mkShell {
      packages = [
          pkgs.go
          pkgs.gom
          pkgs.gopls
          pkgs.gotools
          pkgs.go-tools
      ];
    };
  };
}