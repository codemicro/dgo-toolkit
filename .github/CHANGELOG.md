# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.8.3] - 2021-08-23
### Fixed
* Fixed `RestrictionByRole` (was producing inverted output, sort of?)

## [0.8.2] - 2021-08-23
### Changed
* `func RestrictionByRole(roleId string) CommandRestriction` changed to `func RestrictionByRole(roleIds ...string) CommandRestriction`

## [0.8.1] - 2021-07-03
### Fixed
* Ensure temporary reaction and message handlers have unique IDs
* Loop variable `n` no longer captured by the goroutine used when running temporary message handlers

## [0.8.0] - 2021-05-16
### Added
* `Command` to `route.MessageContext`
### Fixed
* Temporary message handlers now have correctly cloned contexts

##[0.7.1] - 2021-05-05
### Fixed
* Case (in)sensitive modes are now correctly applied when parsing commands

## [0.7.0] - 2021-04-20
### Added
* Temporary message handlers 

## [0.6.0] - 2021-04-02
### Added
* Middleware
* `route.(*MessageContext).SendErrorMessage` and `route.Kit.UserErrorFunc` 

## [0.5.1] - 2021-03-19

## [0.5.0] - 2021-03-19
### Added
* `HasRestrictions` to `CommandInfo`
* `Category` to `Command` and `CommandInfo`
* `kit.NewPagination`

## [0.4.0] - 2021-03-19
### Changed
* Command overloading now needs to be explicitly enabled

## [0.3.3] - 2021-03-17
### Added
* `Raw` field to `MessageContext`

## [0.3.2] - 2021-03-16
### Changed
* Renamed `DiscordSnowflakeType` -> `DiscordSnowflake`
### Added
* `ChannelMention`

## [0.3.1] - 2021-03-15 
### Added
* `DiscordSnowflakeType` argument type
### Changed
* Improve command parsing error messages 

## [0.3.0] - 2021-03-14
### Added
* Support for command overloading
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

[Unreleased]: https://github.com/codemicro/dgo-toolkit/compare/v0.8.2...HEAD
[0.8.2]: https://github.com/codemicro/dgo-toolkit/compare/v0.8.1...v0.8.2
[0.8.1]: https://github.com/codemicro/dgo-toolkit/compare/v0.8.0...v0.8.1
[0.8.0]: https://github.com/codemicro/dgo-toolkit/compare/v0.7.1...v0.8.0
[0.7.1]: https://github.com/codemicro/dgo-toolkit/compare/v0.7.0...v0.7.1
[0.7.0]: https://github.com/codemicro/dgo-toolkit/compare/v0.6.0...v0.7.0
[0.6.0]: https://github.com/codemicro/dgo-toolkit/compare/v0.5.1...v0.6.0
[0.5.1]: https://github.com/codemicro/dgo-toolkit/compare/v0.5.0...v0.5.1
[0.5.0]: https://github.com/codemicro/dgo-toolkit/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/codemicro/dgo-toolkit/compare/v0.3.3...v0.4.0
[0.3.3]: https://github.com/codemicro/dgo-toolkit/compare/v0.3.2...v0.3.3
[0.3.2]: https://github.com/codemicro/dgo-toolkit/compare/v0.3.1...v0.3.2
[0.3.1]: https://github.com/codemicro/dgo-toolkit/compare/v0.3.0...v0.3.1
[0.3.0]: https://github.com/codemicro/dgo-toolkit/compare/v0.2.4...v0.3.0
[0.2.4]: https://github.com/codemicro/dgo-toolkit/compare/v0.2.3...v0.2.4
[0.2.3]: https://github.com/codemicro/dgo-toolkit/compare/v0.2.2...v0.2.3
[0.2.2]: https://github.com/codemicro/dgo-toolkit/compare/v0.2.1...v0.2.2
[0.2.1]: https://github.com/codemicro/dgo-toolkit/compare/v0.2.0...v0.2.1
[0.2.0]: https://github.com/codemicro/dgo-toolkit/compare/v0.1.3...v0.2.0
[0.1.3]: https://github.com/codemicro/dgo-toolkit/compare/v0.1.2...v0.1.3
[0.1.2]: https://github.com/codemicro/dgo-toolkit/compare/v0.1.1...v0.1.2
[0.1.1]: https://github.com/codemicro/dgo-toolkit/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/codemicro/dgo-toolkit/releases/tag/v0.1.0
