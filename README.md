# unboot

> A simple CLI to shutdown or restart macOS.

Regular macOS `shutdown` or `reboot` commands act as a forced shutdown, resulting in apps and windows being restored upon the next login.

This CLI offers a simple alternative to `shutdown` or `reboot` which shut down macOS a little more _gracefully_.


## Installation

Via [Homebrew](https://brew.sh/):
```sh
# Install:
brew install echocrow/tap/unboot
# Update:
brew upgrade echocrow/tap/unboot
```


## Examples

Shutdown:
```sh
unboot
```
Restart/Reboot:
```sh
unboot -r
```
