{
  description = "A bot to post gitea events to a matrix channel";

  # Nixpkgs / NixOS version to use.
  inputs.nixpkgs.url = "nixpkgs/nixos-unstable";

  outputs = { self, nixpkgs }:
    let

      # System types to support.
      supportedSystems =
        [ "x86_64-linux" "x86_64-darwin" "aarch64-linux" "aarch64-darwin" ];

      # Helper function to generate an attrset '{ x86_64-linux = f "x86_64-linux"; ... }'.
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;

      # Nixpkgs instantiated for supported system types.
      nixpkgsFor = forAllSystems (system:
        import nixpkgs {
          inherit system;
          overlays = [ self.overlays.default ];
        });
    in
    {

      # A Nixpkgs overlay.
      overlays.default = final: prev: {
        gitea-matrix-bot = with final;
          buildGoModule {

            pname = "gitea-matrix-bot";
            version = "master";
            src = ./.;
            vendorHash =
              "sha256-4g/pUJ0gGYj6Y3rDINZBbcrEwloA5QCLdviECB1QrIQ=";

            # TODO fix tests
            doCheck = false;
            meta = with lib; {
              description = "A bot to post gitea events to a matrix channel";
              homepage = "https://github.com/pinpox/gitea-matrix-bot";
              license = licenses.gpl3;
              maintainers = with maintainers; [ pinpox ];
            };
          };
      };

      # Package
      packages = forAllSystems (system: {
        inherit (nixpkgsFor.${system}) gitea-matrix-bot;
        default = self.packages.${system}.gitea-matrix-bot;
      });

      # TODO: Add module
    };
}
