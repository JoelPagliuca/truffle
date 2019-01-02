```
 _   _                                     _ _
| \ | | ___   ___ ___  _ __ ___  _ __ ___ (_) |_
|  \| |/ _ \ / __/ _ \| '_ ` _ \| '_ ` _ \| | __|
| |\  | (_) | (_| (_) | | | | | | | | | | | | |_
|_| \_|\___/ \___\___/|_| |_| |_|_| |_| |_|_|\__|
```
Utility to help you **avoid** committing secrets to repositories

## Usage
```sh
truffle -i path/to/your/project
```
Now just use either
```
// truffle
# truffle
```
in any lines you shouldn't commit.

e.g.
```go
var devKey = "15f3440f-1cff-475c-84e0-4b716bb9e3cb" // truffle
```

Any `git commit`s that include these marked lines will be blocked

## Installation
```sh
go get github.com/JoelPagliuca/truffle

# verify installation success
truffle -h
```
Or grab a binary from [here](https://github.com/JoelPagliuca/truffle/releases/latest)

## Development
Using `makefile` for most things.

Set the `VERBOSE` in `main.go` if you want some logging.

I have a test that builds the binary, installs it into a git project, then tries to commit something with a `truffle` tag.

```
build-linux           linux compilation
build-macos           macos compilation
build                 Builds an executable for host platform
clean                 Cleanup any build binaries or packages
cross                 Builds the cross-compiled binaries
help                  Print this message and exit
test-setup            Setup the test project for testing
test                  Run some tests against a test git project
```

## Going forward
* Hit me with an Issue if you find anything wrong or want a feature
* PRs welcome :)