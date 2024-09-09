# CF Sample Plugin

CF Sample Plugin is a Cloud Foundry CLI plugin written in Go. It provides various commands, such as listing deployed applications and running example commands. This project uses [GoReleaser](https://goreleaser.com) to automate the release and multi-platform builds.

## Installation

To install the plugin, either build the code yourself or use the latest precompiled release.

### Install from Precompiled Binaries

1. Download the latest release from [Releases](https://github.com/your-github-username/cf-sample-plugin/releases).
2. Install the plugin with the `cf` command:

   ---
   cf install-plugin ./path/to/binary
   ---

### Manual Compilation

If you'd like to manually build the plugin, follow these steps:

1. Clone the repository:

   ---
   git clone https://github.com/your-github-username/cf-sample-plugin.git
   ---

2. Build the plugin for your platform:

   ---
   go build -o cf-sample-plugin
   ---

3. Install the plugin:

   ---
   cf install-plugin ./cf-sample-plugin
   ---

## Usage

Once installed, the following commands are available:

- **`cf hello`**: Prints a welcome message.
- **`cf list-apps`**: Lists all deployed applications.
- **`cf sample-command`**: Example command without arguments.
- **`cf sample-command-with-args arg1 arg2`**: Example command that takes arguments.

Example usage:

---
cf hello
---

---
cf list-apps
---

## Releasing with GoReleaser

This project uses [GoReleaser](https://goreleaser.com) to automate the build and release process. The `.goreleaser.yml` file is configured to generate binaries for multiple platforms (Linux, macOS, Windows) and automatically create GitHub releases.

### Creating a Release

To create a new release with GoReleaser, ensure you have the appropriate GitHub authentication set up (via GitHub CLI or environment variables), then run:

---
goreleaser release --rm-dist
---

This will:

- Build the binaries for multiple platforms.
- Create a GitHub release with the binaries and checksums.
- Use the Git tag as the version number.

## Contributing

Contributions are welcome! Follow the steps in the [CONTRIBUTING.md](./CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
