# twnyan 🐈

A colorful Twitter client that runs in a terminal

[![release](https://github.com/arrow2nd/twnyan/actions/workflows/release.yml/badge.svg)](https://github.com/arrow2nd/twnyan/actions/workflows/release.yml)
[![arrow2nd](https://circleci.com/gh/arrow2nd/twnyan.svg?style=shield)](https://circleci.com/gh/arrow2nd/twnyan/tree/main)
[![Go Report Card](https://goreportcard.com/badge/github.com/arrow2nd/twnyan)](https://goreportcard.com/report/github.com/arrow2nd/twnyan)
![GitHub all releases](https://img.shields.io/github/downloads/arrow2nd/twnyan/total)
[![GitHub license](https://img.shields.io/github/license/arrow2nd/twnyan)](https://github.com/arrow2nd/twnyan/blob/main/LICENSE.txt)

> **[日本語](README.md)**

![twnyan](https://user-images.githubusercontent.com/44780846/106699506-612c0f80-6626-11eb-803e-332512822789.gif)

## Features

- pseudo-UserStream mode
- Multi-byte character support
- Interactive mode support
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

### Go

```sh
go install github.com/arrow2nd/twnyan@latest
```

### Use binary files

Download the latest version of the file for your environment from [Releases](https://github.com/arrow2nd/twnyan/releases).

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

1. Access the authentication page that appears the first time you start the program
2. Follow the steps on the screen and enter the PIN code that appears on the screen into twnyan
3. done! 😺

## Usage

### Command line mode

![cmdline](https://user-images.githubusercontent.com/44780846/106699170-b287cf00-6625-11eb-8374-8565286db3e2.gif)

`twnyan [command] [argument]`

Also, some commands cannot be used in command line mode.

### にゃーん

![nyaan](https://user-images.githubusercontent.com/44780846/106699001-558c1900-6625-11eb-948e-6212ab0cba40.gif)

`twnyan tw`

If you omit the tweet text, it will be tweeted as "にゃーん".

The same goes for replies and quote retweets.

**If an image is attached, the tweet will not "にゃーん".**

> "にゃーん" is the Japanese word for a cat's meow.

### Interactive mode

`twnyan`

You can like or RT the tweets on the timeline.

You can do this by specifying the tweetnumber displayed at the top of the tweet.

## Command list

<details>
<summary>Open</summary>

## tweet

`Alias: tw`

### tweet [text] [imagefile]

Post a tweet.

| Arguments  | Hints                                                                            | Examples                         |
| ---------- | -------------------------------------------------------------------------------- | -------------------------------- |
| text       | If there is no text and image file, the message will be posted with a "にゃーん" | `tweet`                          |
| image file | If there is more than one, please separate them with a space                     | `tweet 🍣 sushi1.png sushi2.png` |

- You can also omit the text and just post an image (e.g. `tweet cat.png`)

### tweet multi

`Alias: ml`

Post a multi-line tweet.

| Arguments  | Hints                                                        | Examples                      |
| ---------- | ------------------------------------------------------------ | ----------------------------- |
| image file | If there is more than one, please separate them with a space | `tweet multi dog.png cat.png` |

- To finish typing, type a semicolon `;` at the end of the sentence
- To cancel, input `:exit`.

### tweet remove [\<tweetnumber\>]...

`Alias: rm`

Delete a tweet.

| Arguments   | Hints                                                      | Examples           |
| ----------- | ---------------------------------------------------------- | ------------------ |
| tweetNumber | Separate each tweet with a space if there is more than one | `tweet remove 2 5` |

## timeline

`Alias: tl`

### timeline [counts]

Get a home timeline.

| Arguments | Hints                                                                | Examples      |
| --------- | -------------------------------------------------------------------- | ------------- |
| counts    | If you omit it, the default value is given in the configuration file | `timeline 39` |

## stream

`Alias: st`

It first accumulates tweets from the home timeline for a minute, then displays the tweets with a one-minute delay, just like the UserStream API.

**Ctrl+C** to exit.

## mention

`Alias: mt`

### mention [counts]

Get a Mentions to you.

| Arguments | Hints                                                                | Examples     |
| --------- | -------------------------------------------------------------------- | ------------ |
| counts    | If you omit it, the default value is given in the configuration file | `mention 20` |

## list

`Alias: ls`

### list [\<listname\>] [counts]

Get a timeline of the list.

| Arguments | Hints                                                                        | Examples                  |
| --------- | ---------------------------------------------------------------------------- | ------------------------- |
| list name | If you are running in interactive mode, you can complete it with the Tab key | `list Cats`               |
| counts    | If you omit it, the default value is given in the configuration file         | `list "Cat Gathering" 30` |

## user

`Alias: ur`

### user [<username/tweetnumber>] [counts]

Get a timeline of the specified user.

| Arguments            | Hints                                                                | Examples                  |
| -------------------- | -------------------------------------------------------------------- | ------------------------- |
| username/tweetnumber | Either can be specified<br>The '@' in the username is optional       | `user github`<br>`user 1` |
| counts               | If you omit it, the default value is given in the configuration file | `user twitter 15`         |

### user own [counts]

Get your own timeline.

| Arguments | Hints                                                                | Examples      |
| --------- | -------------------------------------------------------------------- | ------------- |
| counts    | If you omit it, the default value is given in the configuration file | `user own 50` |

## search

`Alias: sh`

### search [\<keyword\>] [counts]

Searches for tweets tweets in the past 7 days.

| Arguments | Hints                                                                | Examples           |
| --------- | -------------------------------------------------------------------- | ------------------ |
| keyword   | Please enclose any spaces in double quotes                           | `search "cat dog"` |
| counts    | If you omit it, the default value is given in the configuration file | `search sushi 5`   |

## like

`Alias: lk, fv`

### like [\<tweetnumber\>]

Like a tweet.

| Arguments   | Hints                                                      | Examples   |
| ----------- | ---------------------------------------------------------- | ---------- |
| tweetnumber | Separate each tweet with a space if there is more than one | `like 1 2` |

### like remove [\<tweetnumber\>]

`Alias: rm`

UnLike a tweet.

| Arguments   | Hints                                                      | Examples          |
| ----------- | ---------------------------------------------------------- | ----------------- |
| tweetnumber | Separate each tweet with a space if there is more than one | `like remove 1 2` |

## retweet

`Alias: rt`

### retweet [\<tweetnumber\>]...

Retweet a tweet.

| Arguments   | Hints                                                      | Examples      |
| ----------- | ---------------------------------------------------------- | ------------- |
| tweetnumber | Separate each tweet with a space if there is more than one | `retweet 1 5` |

### retweet remove [\<tweetnumber\>]...

`Alias: rm`

UnRetweet tweet.

| Arguments   | Hints                                                      | Examples             |
| ----------- | ---------------------------------------------------------- | -------------------- |
| tweetnumber | Separate each tweet with a space if there is more than one | `retweet remove 1 5` |

## quote

`Alias: qt`

### quote [\<tweetnumber\>] [text] [imagefile]

Quote a tweet.

| Arguments   | Hints                                                                            | Examples                           |
| ----------- | -------------------------------------------------------------------------------- | ---------------------------------- |
| tweetnumber | Specify the number of the tweet to quote                                         | `quote 1 good!!!`                  |
| text        | If there is no text and image file, the message will be posted with a "にゃーん" | `quote 1`                          |
| imagefile   | If there is more than one, please separate them with a space                     | `quote 1 🍣 sushi1.png sushi2.png` |

### quote multi

`Alias: ml`

Post a multi-line quote retweet.

| Arguments  | Hints                                                        | Examples              |
| ---------- | ------------------------------------------------------------ | --------------------- |
| image file | If there is more than one, please separate them with a space | `quote multi cat.png` |

- To finish typing, type a semicolon `;` at the end of the sentence
- To cancel, input `:exit`.

## reply

`Alias: rp`

### reply [\<tweetnumber\>] [text] [imagefile]

Post a reply.

| Arguments   | Hints                                                                            | Examples                     |
| ----------- | -------------------------------------------------------------------------------- | ---------------------------- |
| tweetnumber | Specify the number of the tweet you want to reply to.                            | `reply 1 meow`               |
| text        | If there is no text and image file, the message will be posted with a "にゃーん" | `reply 1`                    |
| image file  | If there is more than one, please separate them with a space                     | `reply 1 good!!! sushi1.png` |

- You can also omit the text and just post an image. (e.g. `reply 1 dog.png`)

### reply multi

`Alias: ml`

Post a multi-line reply.

| Arguments  | Hints                                                        | Examples              |
| ---------- | ------------------------------------------------------------ | --------------------- |
| image file | If there is more than one, please separate them with a space | `reply multi cat.png` |

- To finish typing, type a semicolon `;` at the end of the sentence
- To cancel, input `:exit`.

## follow

`Alias: fw`

### follow [<username/tweetnumber>]

Follow a user.

| Arguments            | Hints                                                          | Examples                      |
| -------------------- | -------------------------------------------------------------- | ----------------------------- |
| username/tweetnumber | Either can be specified<br>The '@' in the username is optional | `follow github`<br>`follow 1` |

### follow remove [<username/tweetnumber>]

`Alias: rm`

Unfollow a user.

| Arguments            | Hints                                                          | Examples                                       |
| -------------------- | -------------------------------------------------------------- | ---------------------------------------------- |
| username/tweetnumber | Either can be specified<br>The '@' in the username is optional | `follow remove arrow_2nd`<br>`follow remove 1` |

## block

`Alias: bk`

### block [<username/tweetnumber>]

Block a user.

| Arguments            | Hints                                                          | Examples                       |
| -------------------- | -------------------------------------------------------------- | ------------------------------ |
| username/tweetnumber | Either can be specified<br>The '@' in the username is optional | `block arrow_2nd`<br>`block 1` |

### block remove [<username/tweetnumber>]

`Alias: rm`

Unblock a user.

| Arguments            | Hints                                                          | Examples                                     |
| -------------------- | -------------------------------------------------------------- | -------------------------------------------- |
| username/tweetnumber | Either can be specified<br>The '@' in the username is optional | `block remove arrow_2nd`<br>`block remove 1` |

## mute

`Alias: mu`

### mute [<username/tweetnumber>]

Mute a user.

| Arguments            | Hints                                                          | Examples                     |
| -------------------- | -------------------------------------------------------------- | ---------------------------- |
| username/tweetnumber | Either can be specified<br>The '@' in the username is optional | `mute arrow_2nd`<br>`mute 1` |

### mute remove [<username/tweetnumber>]

`Alias: rm`

Unmute a user.

| Arguments            | Hints                                                          | Examples                                   |
| -------------------- | -------------------------------------------------------------- | ------------------------------------------ |
| username/tweetnumber | Either can be specified<br>The '@' in the username is optional | `mute remove arrow_2nd`<br>`mute remove 1` |

## open

`Alias: op`

### open [\<tweetnumber\>]

View the tweet in your browser.

| Arguments   | Hints                                                          | Examples |
| ----------- | -------------------------------------------------------------- | -------- |
| tweetnumber | Specify the number of the tweet to be displayed in the browser | `open 2` |

## clear

Initialize the screen.

## help

Displays help.

You can also use `[command] help` to display help for a command.

## exit

Exits the interactive mode.

</details>

## Configuration directory

The configuration directory will be created directly under your home directory as `.twnyan`

<details>
<summary>Open</summary>

### .cred.yaml

A file of authentication information.

### option.yaml

A file of option setting.

| 名前       | 説明                                |
| ---------- | ----------------------------------- |
| ConfigDir  | Path of the configuration directory |
| Counts     | Default number of acquisitions      |
| DateFormat | Date Format                         |
| TimeFormat | Time Format                         |

- The format of the date and time is the same as the format string of the [time package](https://golang.org/pkg/time/#pkg-constants)

### color.yaml

A file of color settings.

| 名前         | 説明                                           |
| ------------ | ---------------------------------------------- |
| Accent1      | Accent Color 1                                 |
| Accent2      | Accent Color 2                                 |
| Accent3      | Accent Color 3                                 |
| Error        | Background color of error messages             |
| BoxForground | Text color when reversing                      |
| Separator    | Separator                                      |
| UserName     | User Name                                      |
| ScreenName   | Screen Name                                    |
| Reply        | Indication of reply, color of user replying to |
| Hashtag      | Hashtag                                        |
| Favorite     | Display of likes, color of the number of likes |
| Retweet      | Display of Retweet, color of retweet count     |
| Verified     | Verified account                               |
| Protected    | Private account                                |
| Following    | Folloing                                       |
| FollowedBy   | Followed by                                    |
| Block        | Block                                          |
| Mute         | Mute                                           |

</details>
