# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed
- Do not show help on most error messages
- Use go modules for vendoring instead of dep
- Update golangci-lint to v1.17.1
- Update goreleaser to v0.110.0

## [0.2.0] - 2019-05-16
### Fixed
- Add sleep of 100ms to XMPP output to ensure message is sent before connection closes

## [0.2.0] - 2019-05-16
### Changed
- Use etcd instead of redis as storage

## [0.1.0] - 2019-05-13
### Fixed
- Filter invalid users for role (user id == 0)

### Changed
- More detailed errors for Redmine import

### Added
- XMPP output support

## [0.0.1] - 2019-05-13
### Added
- Import user mails for project members from Redmine
- Output events via mail
- Output events to Slack
