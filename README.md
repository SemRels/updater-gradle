# updater-gradle

Gradle package updater plugin for SemRel.

Updates Gradle project versions in version.properties and build.gradle style files.

## Documentation

- SemRel docs (planned): <https://github.com/SemRels/semrel/tree/main/docs/plugins/updater-gradle>
- Plugin template: <https://github.com/SemRels/plugin-template>
- Registry: <https://registry.semrel.io>

## Repository Layout

~~~text
cmd/plugin/              Plugin entry point
internal/plugin/         Business logic scaffold
internal/grpc/           gRPC transport scaffold
proto/v1                 Symlink to the SemRel protobuf contract
.github/workflows/       CI, release, and security automation
~~~

## Development

~~~bash
go build ./cmd/plugin
go test ./...
~~~

## Configuration Example

~~~yaml
plugins:
  - name: updater-gradle
    type: updater
    config:
      property_file: gradle/version.properties
      property_name: version
      build_files:
        - build.gradle
        - build.gradle.kts
~~~

## Status

This repository is bootstrapped from SemRels/plugin-template and is ready for implementation.
