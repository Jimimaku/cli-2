# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres
to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

### 0.35.0

We are introducing a set of new environment management commands that will
eventually replace `state activate`. The intend behind this is to make the
use-cases currently covered by the activate command more explicit, so that users
have more control over their workflow.

In short; we're introducing the following commands:

- *checkout* - Checkout the given project and setup its runtime
    - A checkout is required before you can use any of the following commands
- *use* - Use the given project runtime as the default for your system
    - *reset* - Reset your default project runtime (this also resets the project
      configured via `state activate --default`)
    - *show* - Show your default project runtime
- *shell* - Starts a shell/prompt for the given project runtime (equivalent of
  virtualenv)
- *switch* - Switch to a branch or commit

All of the above commands are currently marked as unstable, meaning you cannot
use them unless you opt-in to unstable commands with
`state config set optin.unstable true`.
This is to give us time to test and improve the commands without necessarily
ensuring backward compatibility. These commands have been thoroughly tested, but
since they are new bugs are still more likely than with stable commands.

Note that `state activate` will still be available for the foreseeable future.

### Added

- Added new environment management commands (see above for details)
    - Added `state checkout` command.
    - Added `state use` command.
    - Added `state use reset` command.
    - Added `state use show` command.
    - Added `state shell` command.
    - Added `state switch` command.
- Added `state export env` command - Export the environment variables associated
  with your runtime.
- Added `state deploy uninstall` command for reverting a `state deploy`.
- Added `state update unlock` command, which undoes what `state update lock`
  does.
- Runtime artifacts are now cached, speeding up runtime setup and reducing
  network traffic.
    - The cache is capped at 500mb. This can be overridden with
      the `ACTIVESTATE_ARTIFACT_CACHE_SIZE_MB` environment variable (value is
      MB's as an int).

### Changed

- State tool will now error out when passed superfluous arguments (
  eg. `state activate name/space superfluos-arg`).
- The installer will no longer show debug error messages.
- We now start the background service automatically when you boot your machine.
- State tool now configures all compatible shells that were found on the users
  system.
- We now report how far ahead / behind you are from your branch when
  running `state show`.

### Fixed

- Fixed State Tool being unusable on M1 Macs running Ventura.
- Fixed `~/.cshrc` not being respected when using `tcsh`.
- Fixed `-v` flag not working when using `install.sh` to install State Tool.
- Fixed state tool background service closing prematurely.
- Fixed bash scripts on Windows using the wrong path format.
- Fixed a variety of missing/wrong localisation issues.
- Fixed `state invite` resulting with response code error message.
- Fixed various issues where running with `--non-interactive` would not have
  the desired behavior.
- Fixed `state config set` accepting invalid values for booleans.
- Fixed `state exec` not respecting the `--path` flag.
- Fixed issue where PYTHONPATH would be set up with a temp directory on macOS.
    - This still worked as expected in the end, but is obviously awkward.
- Fixed panic when running `state secrets get` without a project.
- Fixed issue where `state learn` would give an unhelpful error when it could
  not reach the browser.
- Fixed `state show` not working for private projects.
- Fixed variables as arguments to executors (eg. python3.exe) not being expanded
  properly.
- Fixed state tool interpreting `-v` flag when its passed through `state run` or
  `state exec` but not intended for the state tool.
- Fixed State Tool being added to PATH multiple times.
- Fixed unstable commands reporting `--help` info when passed invalid
  arguments, instead of saying the command is unstable and you should opt in.
- Fixed `state uninstall` with a non-existent package reporting the wrong error.

### 0.34.1

### Changed

* The `state use` command has been marked unstable.

### Fixed

* Fixed issue where activating a second project with an identical name to the
  first would instead activate the first project.
* Fixed issue where error output was sometimes missing important details about
  what went wrong.
* Fixed issue where build errors were incorrectly reported.
* Fixed issue where service could not run due to filepath size limits on macOS.
* Fixed issue where passing a relative path to `state activate --path` would
  sometimes not resolve to the correct path.
* Fixed issues where installer would sometimes give the update user experience.

### 0.34.0

### Added

* We've started flagging commands as stable and unstable, and by default will
  only support execution of stable commands. To run unstable commands you must
  first opt-in to them using `state config set optin.unstable true`.
* We've added a new `state use <orgname/project>` command, which will allow you
  configure the given project as the default runtime on your system.
* Automatic updates can now be disabled with `state config set autoupdate false`
  .
* On Windows we now add an Uninstall shortcut to the start menu.
* Analytics can now also be disabled with an environment variable:
  `ACTIVESTATE_CLI_DISABLE_ANALYTICS=true`.

### Changed

* The state-svc (our background daemon) has seen significant improvements to its
  start / stop behavior. Primarily intended to improve the reliability of our
  update process.
    * As a result our minimum Windows version required to run the state tool is
      now *Windows 10 Build 17134 (Codename Redstone 4)*.
* The State tool will now error out when it can't communicate with the
  state-svc.
  Preventing the user from running into much more vague errors as a result of
  the
  missing daemon.
* `state config` can now only act on valid config keys.
* A number of error messages have been improved to give a better idea of how the
  user can remedy the error.
* Our installer has been optimized to use a smaller file size and reduce the
  number of processes as part of the installation.

### Fixed

* Fixed issue where variables in command line arguments were not properly
  interpolated. Causing the command to receive an empty value rather than
  the variable name.
* Fixed issue where `state clean uninstall` would fail to clean up the
  environment.
* Fixed issue where `state activate --branch` would sometimes error out.
* Various issues leading to corrupt, miss-placed, or error-prone installation
  directories.
* Fixed issue where the State Tool installation directory was added to PATH
  multiple times.
* Fixed issue where calling `state clean cache` with `--non-interactive`
  did not clean the cache.
* Fixed issue where `state history` would fail if history had an author that is
  no longer a member of the organization.
* Fixed issue where automated tools and integrations (including Komodo IDE)
  could not get the list of organizations for the authenticated user due to a
  backwards incompatible change.
* Fixed cases of missing localization.

### Removed

* The `--replace` flag has been dropped from `state activate`, its use-case has
  been addressed by `state pull --set-project`.

## 0.33.0

### Added

* Authentication now uses your browser for a more secure and transparent
  authentication process.

    * The old behavior is still available as well, and use-cases where you
      provide
      the api key or credentials in the command are unaffected.

* Added a new `state config` command, which can be used to change behavior of
  the State Tool itself.

    * Currently can be used to disable analytics and error reporting, eg.

  ```bash
  state config set report.analytics false # Turns off analytics
  state config set report.errors false # Turns off error reporting
  ```

### Fixed

* Fixed issue where temporary files were not cleaned up in a timely manner.
* Fixed issue where the `state-svc` process would not be shut down correctly.
* Fixed issue where `state clean uninstall` would say it succeeded but the State
  Tool would still be installed.

### Changed

* Several performance enhancements have been made affecting all parts of the
  State Tool.
* Activating an already activated project won't error out anymore.
* The local project is no longer affected if `state install` fails.

### Removed

* The `-c` flag has been removed from `state activate` as this is now handled
  by `state exec`.

## 0.32.2

### Fixed

* Fixed issue where auto-update could not complete for certain older versions

## 0.32.1

### Fixed

* Fixed issue that could sometimes cause recursion in our logging

## 0.32.0

### Added

* Added PPM and PIP shims to help educate people about the State Tool.
* Added support for Ruby projects

## 0.31.1

### Fixed

* Fixed issue where a failed solve was reported incorrectly.

## 0.31.0

### Changed

* More progress indicators are now given when sourcing runtimes and installing
  packages.
* Package operations are now much faster
* Binary sizes have been significantly reduced
* You no longer need to start a new shell when installing the State Tool (
  provided you're running an interactive session)

## 0.30.7

### Fixed

* Fixed issue where environment would not always be sourced properly

## 0.30.6

### Fixed

* Fixed issue where certain runtime executables could not be resolved

## 0.30.5

### Changed

* Recursion has been disabled while we improve the mechanic for a future version

## 0.30.4

### Fixed

* Fixed recursion issue when running certain State Tool commands

## 0.30.3

### Changed

* Enriched the installer with analytics to allow us to diagnose installation
  failures

## 0.30.2

### Fixed

* Fixed issue where State Tool sometimes could not identify its service daemon

## 0.30.1

### Fixed

* Fixed issue where our analytics events would send the full executable paths

## 0.30.0

### Added

* New Command: `state learn`. Opens
  the [State Tool Cheat Sheet](https://platform.activestate.com/state-tool-cheat-sheet)
  in your browser.

### Changed

* The install and activate user experience have been overhauled to be much more
  concise and avoid unnecessary prompts.
* Several performance enhancements have been made. Note that some of these will
  require at least one more release before they can realise their potential.
* Running `state update` will now immediately perform the update, rather than
  deferring it to a background process.
* State Tool should now attempt to use the latest version available for a given
  language, when initializing a project.

### Fixed

* Fixed issue where on macOS the `state` executable would sometimes not be added
  to your PATH.
* Resolved issue where `state exec` or certain invocations of the language
  runtime could lead to recursion errors.
* Fixed issues where sometimes State Tool would say it have a new version
  available when it didn't.

## 0.29.5

### Fixed

- Fixed race condition in anonymized analytics

## 0.29.4

### Changed

- Improved error reporting to help direct stability improvements

## 0.29.3

### Fixed

- Fixed race condition that could lead to logs being written to stderr

## 0.29.2

### Fixed

- Uninstalling no longer leaves a stale executable

## 0.29.1

### Fixed

- Auto updating from earlier versions no longer results in error

## 0.29.0

### Added

- Package management is now performed only locally, meaning you have
  to `state push` your changes back to your project when you are ready to save
  them.
- Enhanced error reporting when attempting package operations on an out of sync
  project ([PR #1353](https://github.com/ActiveState/cli/pull/1353))
- Enhanced error reporting for errors that occur when cloning a project's
  associated git
  repository ([PR #1351](https://github.com/ActiveState/cli/pull/1351))
- The State Tool now comes with a preview of the ActiveState Desktop
  application, which facilitates shortcuts to commonly used actions, including
  activating your projects.
- You can now switch to specific State Tool versions by
  running `state update --set-version <version>` ([PR #1385](https://github.com/ActiveState/cli/pull/1385))

### Changed

- Enhanced error reporting for errors that occur when cloning a project's
  associated git
  repository ([PR #1351](https://github.com/ActiveState/cli/pull/1351))

### Removed

- We no longer produce 32bit Windows builds of the State Tool

### Fixed

- Removed unwanted output (eg. `%!s(<nil>)`) when running scripts
  ([PR #1354](https://github.com/ActiveState/cli/pull/1354))
- Fixed issue where `state clean uninstall` would not remove expected files on
  Windows ([PR #1349](https://github.com/ActiveState/cli/pull/1349))
- Fixed a rare case where the configuration file can get corrupted when two
  processes access it
  simultaneously.  ([PR #1370] (https://github.com/ActiveState/cli/pull/1370))

## 0.28.1

### Fixed

* Fixed package installs / uninstalls not using the
  cache ([PR #1331](https://github.com/ActiveState/cli/pull/1331))

## 0.28.0

### Changed

- New runtimes are installed in parallel and 2-4 times faster.
  ([PR #1275](https://github.com/ActiveState/cli/pull/1275))

### Fixed

- `state push` updates project name in activestate.yaml.
  ([PR1297](https://github.com/ActiveState/cli/pull/1297))

## 0.27.1

### Fixed

- Fixed issue where `state uninstall` would not completely remove package files
  ([PR #1304](https://github.com/ActiveState/cli/pull/1304))

## 0.27.0

### Added

- New system tray executable for the Windows platform
  ([PR #1285](https://github.com/ActiveState/cli/pull/1285))

### Changed

- Enhanced error reporting for errors that happened early on in the application
  logic ([PR #1280](https://github.com/ActiveState/cli/pull/1280))
- Updated name of `state cve` command to `state security`. Aliased `state cve`
  to `state security` ([PR #1286](https://github.com/ActiveState/cli/pull/1286))

### Fixed

- Fixed issue where `state push` would fail on existing projects.
  ([PR #1287](https://github.com/ActiveState/cli/pull/1287))

## 0.26.0

### Added

- New command `state cve open <cve-id>` opens the National Vulnerability
  Database entry for the given
  CVE ([PR #1269](https://github.com/ActiveState/cli/pull/1269))

### Fixed

- Fixed issue where `state deploy` would fail without the `--path` flag
  ([PR #1270](https://github.com/ActiveState/cli/pull/1270))

## 0.25.1

### Fixed

- Fixed issue where `state pull` would not pull in the latest
  changes ([PR #1272](https://github.com/ActiveState/cli/pull/1272))

## 0.25.0

**Warning:** This update will force a change to your activestate.yaml which is
incompatible with earlier state tool versions. As long as everyone on your
project updates their state tool there should be no interruption to your
workflow.

### Added

- New command `state cve` allows for reviewing security vulnerabilities on your
  project ([PR #1209](https://github.com/ActiveState/cli/pull/1209))
- You can now specify a package version when calling `state info`,
  eg. `state info <name>@<version>` ([PR #1201](https://github.com/ActiveState/cli/pull/1201))
- You can now specify a new project name by
  running `state pull --set-project OWNER/NAME` (primarily for converting
  headless projects) ([PR #1198](https://github.com/ActiveState/cli/pull/1198))
- You can now switch between update channels
  via `state update --set-channel` ([PR #1190](https://github.com/ActiveState/cli/pull/1190))
- State tool will now provide instructions on how to get out of a detached
  state ([PR #1249](https://github.com/ActiveState/cli/pull/1249))
- State tool now supports branches via flags in `state activate` and
  the `state branch` subcommand. See `state branch --help` for more information.

### Changed

- Activating a new project non-interactively no longer makes that project "
  default" (you can pass the `--default` flag for this
  use-case) ([PR #1210](https://github.com/ActiveState/cli/pull/1210))
- The user experience of `state secrets` is now consistent with the rest of the
  State Tool ([PR #1197](https://github.com/ActiveState/cli/pull/1197))
- `state import` now updates your runtime, so you don't need to re-activate
  after importing
  anymore ([PR #1241](https://github.com/ActiveState/cli/pull/1241))

### Fixed

- Progressbar sometimes hangs while waiting for build to
  complete ([PR #1218](https://github.com/ActiveState/cli/pull/1218))
- Fixed issue where some unicode characters were not printed
  properly ([PR #1207](https://github.com/ActiveState/cli/pull/1207))
- Prompts for default project should now only happen once per
  project ([PR #1210](https://github.com/ActiveState/cli/pull/1210))
- Fixed issue where `state activate` sometimes used the wrong
  activestate.yaml ([PR #1194](https://github.com/ActiveState/cli/pull/1194))
- Fixed issue where `state info owner/name` would fail if not currently in a
  project directory ([PR #1255](https://github.com/ActiveState/cli/pull/1255))
- Fixed issue where running tooling from the global default project with
  the `-v` flag would spew out state tool debug
  info ([PR #1239](https://github.com/ActiveState/cli/pull/1239))
- Fixed issue where sometimes perl/python is still pointing at the system
  install after
  activation ([PR #1238](https://github.com/ActiveState/cli/pull/1238))
- Fix issue where state tool sometimes throws "panic" errors when updating the
  configuration ([PR #1232](https://github.com/ActiveState/cli/pull/1232))
- Fix issue where `state activate` sometimes throws a "
  panic" ([PR #1229](https://github.com/ActiveState/cli/pull/1229))

### Deprecated

- The `--replace` flag for `state activate` is now deprecated in favour
  of `state pull --set-project`
