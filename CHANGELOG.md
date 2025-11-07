# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Added
- [#14] Extend Backup api to get backups and restores
- [#15] get & set backup-schedule
- [#15] get retention-policy 
- [#17] create backups
- [#20] create restores
- [#21] delete restores

## [v1.5.0] - 2025-08-01
### Changed
- `CreateSupportArchiveRequest` now has two formats (Legacy and Common) (#11)
  - Legacy is the old format used by `cesappd` in Classic CES
  - Common is the new more general format supported by Multinode as well as hopefully future generations of the EcoSystem

## [v1.4.0] - 2024-09-18
### Changed
- Relicense to AGPL-3.0-only

## [v1.3.0] - 2024-05-21
### Added
- Add log level to api (#7)

## [v1.2.0] - 2024-05-21
### Added
- Add rpc to set log level for specific dogu in logging API which also restarts the dogu when changed (#5)

## [v1.1.0] - 2024-05-02
### Added
- Add new query API for logs (#3)

## [v1.0.0] - 2024-04-25
### Changed
- extracted grpc api definitions from k8s-ces-control and cesappd (#1)
