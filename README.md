# updater-gradle

[![Latest Release](https://img.shields.io/github/v/release/SemRels/updater-gradle?label=version&color=blue)](https://github.com/SemRels/updater-gradle/releases/latest)
[![CI](https://github.com/SemRels/updater-gradle/actions/workflows/ci.yml/badge.svg)](https://github.com/SemRels/updater-gradle/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/SemRels/updater-gradle)](LICENSE)
[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/SemRels/updater-gradle/badge)](https://scorecard.dev/viewer/?uri=github.com/SemRels/updater-gradle)
[![REUSE status](https://api.reuse.software/badge/github.com/SemRels/updater-gradle)](https://api.reuse.software/info/github.com/SemRels/updater-gradle)

Updates a Gradle version property.

This plugin is distributed as the standalone Go binary `semrel-plugin-updater-gradle`. Semrel executes the binary as a subprocess, provides plugin configuration through `SEMREL_PLUGIN_*` environment variables, provides release context through `SEMREL_*` environment variables, reads standard output, and treats exit code `0` as success and any non-zero exit code as failure. Install the binary in `~/.semrel/plugins/` or anywhere on your `$PATH`.

## Installation

### Binary

```bash
go install github.com/SemRels/updater-gradle/cmd/plugin@latest
```

### Docker

Pre-built, multi-platform images (linux/amd64, linux/arm64) are published to the GitHub Container Registry on every release:

```bash
docker pull ghcr.io/semrels/updater-gradle:latest
```

Images are signed with [cosign](https://github.com/sigstore/cosign) and include a full SBOM attestation. Verify the signature:

```bash
cosign verify ghcr.io/semrels/updater-gradle:latest \
  --certificate-identity-regexp 'https://github.com/SemRels/updater-gradle/.github/workflows/container.yml.*' \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com
```


## Configuration

```yaml
plugins:
  - uses: "@semrel/updater-gradle"
    phase: pre-tag
    args:
      file: "gradle.properties"
      key: "version"
```

## Options

Every option can be set as an entry under `args` in `.semrel.yaml`. semrel passes each entry to the plugin as the environment variable shown in the table, so you can also set that variable directly (for example in your CI environment) instead of using `args`.

| Argument | Environment variable | Required | Description | Default |
| --- | --- | --- | --- | --- |
| `file` | `SEMREL_PLUGIN_FILE` | Optional | Path to the Gradle properties file to update. | gradle.properties |
| `key` | `SEMREL_PLUGIN_KEY` | Optional | Property key that stores the version. | version |

## `SEMREL_*` release context used

| Variable | Description |
| --- | --- |
| `SEMREL_VERSION` | Resolved release version for the current run. |
| `SEMREL_NEXT_VERSION` | Next version computed by semrel for the release. |
| `SEMREL_DRY_RUN` | Whether semrel is running in dry-run mode. |

## Example behavior

The plugin updates the configured Gradle property to the new version and logs the file change.

## License

Apache-2.0
