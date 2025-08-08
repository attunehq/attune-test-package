# Attune Test Package Suite

A collection of test packages for Attune CLI smoke tests, including multiple fish-themed variants.

## Packages

This repository contains four identical test packages with different names:

- **attune-test-package** - The original test package
- **tuna-test-package** - Tuna-themed variant
- **salmon-test-package** - Salmon-themed variant  
- **cod-test-package** - Cod-themed variant

## Usage

All packages have identical functionality with the same command-line interface:

```bash
# Basic usage (replace with any package name)
attune-test-package
tuna-test-package
salmon-test-package
cod-test-package

# Show version
attune-test-package --version

# Show system info
attune-test-package --info
```

## Project Structure

```
cmd/
├── attune-test-package/main.go
├── tuna-test-package/main.go
├── salmon-test-package/main.go
└── cod-test-package/main.go
```

## Releasing

This project uses [GoReleaser](https://goreleaser.com/) to automate the release process. New versions are released by creating and pushing Git tags. **All four packages are built and released together.**

### Prerequisites

1. Install GoReleaser:
   ```bash
   # macOS
   brew install goreleaser
   
   # Or download from https://github.com/goreleaser/goreleaser/releases
   ```

2. Ensure you have push access to the repository and appropriate GitHub permissions.

### Release Process

1. **Ensure your local repository is up to date:**
   ```bash
   git checkout main
   git pull origin main
   ```

2. **Create and push a new tag:**
   ```bash
   # Create a new tag (replace v1.0.0 with your desired version)
   git tag v1.0.0
   
   # Push the tag to trigger the release
   git push origin v1.0.0
   ```

3. **Run GoReleaser locally (optional for testing):**
   ```bash
   # Test the release process without publishing
   goreleaser release --snapshot --clean
   
   # Or run a full release (requires proper Git tag)
   goreleaser release --clean
   ```

### What Gets Released

The release process will automatically build and release **all four packages**:
- Build binaries for Linux, Windows, and macOS (amd64 and arm64)
- Create separate compressed archives for each package (tar.gz for Unix, zip for Windows)
- Generate separate .deb packages for each package on Debian/Ubuntu
- Create a GitHub release with changelog
- Upload all artifacts to the GitHub release

### Release Artifacts

For **each of the four packages**, the following artifacts are created:
- **Binaries**: Platform-specific executables
- **Archives**: Compressed archives named `{package-name}_{OS}_{ARCH}.{ext}`
- **Packages**: `.deb` packages for Debian-based systems
- **Checksums**: SHA256 checksums for all artifacts

Example artifacts for a single release:
- `attune-test-package_Linux_x86_64.tar.gz`
- `tuna-test-package_Darwin_arm64.tar.gz`
- `salmon-test-package_Windows_x86_64.zip`
- `cod-test-package_1.0.0_amd64.deb`
- And many more...