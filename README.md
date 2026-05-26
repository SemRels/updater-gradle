# updater-gradle

Updates a Gradle version property.

This plugin is distributed as the standalone Go binary `semrel-plugin-updater-gradle`. Semrel executes the binary as a subprocess, provides plugin configuration through `SEMREL_PLUGIN_*` environment variables, provides release context through `SEMREL_*` environment variables, reads standard output, and treats exit code `0` as success and any non-zero exit code as failure. Install the binary in `~/.semrel/plugins/` or anywhere on your `$PATH`.

## Installation

```bash
go install github.com/SemRels/updater-gradle/cmd/plugin@latest
```

## Configuration

```yaml
plugins:
  - name: updater-gradle
    path: ~/.semrel/plugins/semrel-plugin-updater-gradle
    env:
      SEMREL_PLUGIN_FILE: "gradle.properties"
      SEMREL_PLUGIN_KEY: "version"
```

## `SEMREL_PLUGIN_*` variables

| Name | Required | Description | Default |
| --- | --- | --- | --- |
| `SEMREL_PLUGIN_FILE` | Optional | Path to the Gradle properties file to update. | gradle.properties |
| `SEMREL_PLUGIN_KEY` | Optional | Property key that stores the version. | version |

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
