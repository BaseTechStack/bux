# BUX Release Process

This document outlines the process for creating and publishing new releases of the BUX CLI.

## Prerequisites

- Git access to the BaseTechStack/bux repository
- GitHub CLI (gh) installed and authenticated
- Go installed for building binaries

## Release Steps

1. **Update CHANGELOG.md**

   Add a new section to the CHANGELOG.md file with the new version number and the changes included in the release.

   ```markdown
   ## [vX.Y.Z] - YYYY-MM-DD

   ### Added
   - New features

   ### Changed
   - Changes to existing functionality

   ### Fixed
   - Bug fixes

   ### Removed
   - Removed features
   ```

   Don't forget to add the link at the bottom of the file:

   ```markdown
   [vX.Y.Z]: https://github.com/BaseTechStack/bux/releases/tag/vX.Y.Z
   ```

2. **Run the Release Script**

   ```bash
   ./release.sh X.Y.Z
   ```

   This script will:
   - Create a new Git tag for the version
   - Build binaries for multiple platforms (darwin_amd64, darwin_arm64, linux_amd64, windows_amd64)
   - Create archives for each platform
   - Attempt to push the tag to GitHub
   - Attempt to create a GitHub release with the archives

   If you don't have push access to the GitHub repository, the script will still create the binaries locally.

3. **Manual GitHub Release (if needed)**

   If the release script couldn't push to GitHub, you'll need to manually:
   
   - Push the tag to GitHub:
     ```bash
     git push origin vX.Y.Z
     ```
   
   - Create a new release on GitHub:
     - Go to https://github.com/BaseTechStack/bux/releases/new
     - Select the tag you just pushed
     - Add a title and description (can be copied from CHANGELOG.md)
     - Upload the archive files generated by the release script
     - Publish the release

4. **Update Installation Scripts (if needed)**

   If there are changes to the installation process, update:
   - `install.sh` for Unix systems
   - `install.ps1` for Windows systems

## Versioning

This project follows [Semantic Versioning](https://semver.org/):

- **MAJOR version** when you make incompatible API changes
- **MINOR version** when you add functionality in a backwards compatible manner
- **PATCH version** when you make backwards compatible bug fixes

## File Naming Conventions

All release files should follow these naming conventions:

- Archives: `bux_${OS}_${ARCH}.tar.gz` or `bux_${OS}_${ARCH}.zip` for Windows
- Binary: `bux` (Unix) or `bux.exe` (Windows)
- Installation directory: `.bux` in the user's home directory

## Troubleshooting

- If a tag already exists, the release script will delete it locally before creating a new one
- If you can't push to GitHub, you may need to set up SSH keys or use a personal access token
- If the GitHub release creation fails, you can still manually create a release on GitHub
