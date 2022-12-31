# 🐈 twnyan

> **Warning**
>
> twnyan is no longer in development. Recommend to migrate to
> [nekome](https://github.com/arrow2nd/nekome).

A colorful Twitter client that runs in a terminal

[![release](https://github.com/arrow2nd/twnyan/actions/workflows/release.yml/badge.svg)](https://github.com/arrow2nd/twnyan/actions/workflows/release.yml)
[![test](https://github.com/arrow2nd/twnyan/actions/workflows/test.yml/badge.svg)](https://github.com/arrow2nd/twnyan/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/arrow2nd/twnyan)](https://goreportcard.com/report/github.com/arrow2nd/twnyan)
![GitHub all releases](https://img.shields.io/github/downloads/arrow2nd/twnyan/total)
[![GitHub license](https://img.shields.io/github/license/arrow2nd/twnyan)](https://github.com/arrow2nd/twnyan/blob/main/LICENSE)

> **[日本語](README.md)**

![twnyan](https://user-images.githubusercontent.com/44780846/106699506-612c0f80-6626-11eb-803e-332512822789.gif)

## Features

- You can "にゃーん" anytime 🐾
- Pseudo UserStream mode
- Multi-account support
- Both command line and interactive mode
- Very colorful 🎨

## System requirements

- Windows / macOS / Linux
- A terminal that can display emoji

## Install

### Homebrew

```sh
brew tap arrow2nd/tap
brew install twnyan
```

### Scoop

```
scoop bucket add arrow2nd https://github.com/arrow2nd/scoop-bucket.git
scoop install arrow2nd/twnyan
```

### Go

```sh
go install github.com/arrow2nd/twnyan@latest
```

### Use binary files

Download the latest version of the file for your environment from
[Releases](https://github.com/arrow2nd/twnyan/releases).

## Uninstall

### Homebrew

```sh
brew uninstall twnyan
rm -rf ~/.twnyan # Delete configuration files
```

### Go

```sh
go clean -i github.com/arrow2nd/twnyan
rm -rf $GOPATH/src/github.com/arrow2nd/twnyan
rm -rf ~/.twnyan
```

## Initialization

![auth](https://user-images.githubusercontent.com/44780846/106747441-4a59dd00-6667-11eb-8248-3468cb39f7d1.png)

1. Access the authentication page that appears the first time you start the
   program
2. Follow the steps on the screen and enter the PIN code that appears on the
   screen into twnyan
3. done! 😺

> The first account added is treated as the "main account"

## Usage

### Command line mode

![cmdline](https://user-images.githubusercontent.com/44780846/106699170-b287cf00-6625-11eb-8374-8565286db3e2.gif)

- Some commands are not available

### Interactive mode

If no command is specified, it starts in interactive mode.

### にゃーん

![nyaan](https://user-images.githubusercontent.com/44780846/106699001-558c1900-6625-11eb-948e-6212ab0cba40.gif)

```
twnyan tw
```

If the tweet sentence is omitted, it will be tweeted as "にゃーん".

- Replies and quoted retweets work the same way
- If an image is attached, it will not "にゃーん"

> "にゃーん" is the Japanese word for a cat's meow

### Detailed usage

- [Command List](./docs/en/commands/index.md)
- [About the configuration file](./docs/en/options.md)
