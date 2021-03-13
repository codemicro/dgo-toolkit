# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Changed
* Switch method used to detect if a channel is a direct message
  * Doesn't involve an API call, much faster than the old method 

## [0.2.4] - 2021-03-13
### Added
* `kit.AllowDirectMessages`
### Changed
* Order of execution in `Duration` argument parsing for better error messages

## [0.2.3] - 2021-03-13
### Added
* `Duration` argument type

## [0.2.2] - 2021-03-12
### Changed
* Export `Kit` in `CommonContext`

## [0.2.1] - 2021-03-12
### Changed
* Allow `acceptFunc` and `rejectFunc` in `NewConfirmation` to be `nil` 

## [0.2.0] - 2021-03-12
### Added
* `NewConfirmation`
* `*kit.AddTemporaryReaction`, `*kit.RemoveTemporaryReaction`

## [0.1.3] - 2021-03-08
### Added
* `URL` argument type

## [0.1.2] - 2021-03-08
### Added
* `*messageContext.SendMessageString` and `*messageContext.SendMessageEmbed`

## [0.1.1] - 2021-03-08
### Added
* `kit.AllowBots`
* `kit.DefaultAllowedMentions`
### Changed
* Bot now ignores own messages and reactions

## [0.1.0] - 2021-03-08
### Added
* Everything (initial release)

[Unreleased]: https://github.com/codemicro/dgo-toolbox/compare/v0.2.2...HEAD
[0.2.2]: https://github.com/codemicro/dgo-toolbox/compare/v0.2.1...v0.2.2
[0.2.1]: https://github.com/codemicro/dgo-toolbox/compare/v0.2.0...v0.2.1
[0.2.0]: https://github.com/codemicro/dgo-toolbox/compare/v0.1.3...v0.2.0
[0.1.3]: https://github.com/codemicro/dgo-toolbox/compare/v0.1.2...v0.1.3
[0.1.2]: https://github.com/codemicro/dgo-toolbox/compare/v0.1.1...v0.1.2
[0.1.1]: https://github.com/codemicro/dgo-toolbox/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/codemicro/dgo-toolbox/releases/tag/v0.1.0
