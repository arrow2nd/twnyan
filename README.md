# twnyan
[![Go Report Card](https://goreportcard.com/badge/github.com/arrow2nd/twnyan)](https://goreportcard.com/report/github.com/arrow2nd/twnyan)

This is a Twitter client by cats for cats🐾

> **[日本語](README_JP.md)**

## Features
- Multi-byte character support
- It's easy to tweet "にゃーん"
- ~Needlessly~ flexible color settings
- Support for interactive mode

## Screenshot
![twnyan](https://user-images.githubusercontent.com/44780846/99259409-5058d280-285d-11eb-82f3-ba80065517be.gif)

## How to install
**(Recommended)**

```$ go get github.com/arrow2nd/twnyan```

### Use binary files
Download a zip file from ReleasePage that fits your environment and pass the Path through the binary file.

## Operating conditions
- Windows/Linux
- Emoji can be displayed.

### Remarks
- I have not been able to confirm that it works on a Mac
- To run in a WSL environment, you need to be able to use xdg-open

## Usage
```$ twnyan [command] [argument]```

If you omit the command ```$ twnyan```, it will start in interactive mode.

You can manipulate tweets (like, RT, etc.) by specifying the number of the tweet.
> Example: ```favorite 0```

## Command list

<details>
<summary>Open</summary>

## tweet
**tweet [Subcommand] [argument]**

Manipulates tweets.
> Alias: tw

| Subcommand | Alias | Description | Argument |
| -------- | -------- | -------- | -------- |
| none |  | The default value for the tweet is 10,000 words. | ```tweet [text] [image file]```.
| remove | rm | remove the tweet. | ```tweet remove [<tweet number>]```.

| Arguments | Hints | Examples |
| -------- | -------- | -------- |
| text | If there is no text and image file, the message will be posted with a "にゃーん" | ```tweet``` |
| image file | If there is more than one, please separate them with a space | ```tweet 🍣 sushi1.png sushi2.png``` |
| tweet number | Separate each tweet with a space if there is more than one | ```tweet remove 2 5``` |

- You can also omit the text and just post an image (e.g. ```tweet cat.png```)

## timeline
**timeline [counts]**

Displays the home timeline.
> Alias: tl

| Arguments | Hints | Examples |
| -------- | -------- | -------- |
| counts | If you omit it, the default value is given in the configuration file | ```timeline 39``` |

## mention
**twnyan mention [counts]**

Displays the Mentions to you.
> Alias: mt

| Arguments | Hints | Examples |
| -------- | -------- | -------- |
| counts | If you omit it, the default value is given in the configuration file | ```mention 20``` |

## list
**list [<list name>] [counts]**

Displays the timeline of the list.
> Alias: ls

| Arguments | Hints | Examples |
| -------- | -------- | -------- |
| list name | If you are running in interactive mode, you can complete it with the Tab key | ```list Cats```|
| counts | If you omit it, the default value is given in the configuration file | ```list "Cat Gathering" 30``` |

## user
**user [Subcommand] [argument]**

Displays the user timeline.
> Alias: ur

| Subcommand | Alias | Description | Argument |
| -------- | -------- | -------- | -------- |
| none |  | Displays the timeline of the specified user | ```user [userID] [counts]``` |
| number | num, no | Displays the timeline of the person who posted the specified tweet | ```user number [<tweet number>] [counts]``` |

| Arguments | Hints | Examples |
| -------- | -------- | -------- |
| userID | If you omit it, you will be specified | ```user``` |
| counts | If you omit it, the default value is given in the configuration file | ```user twitter 15``` |

- The '@' in the user ID is optional

## search
**search [<keyword>] [counts]**

Searches for tweets tweets in the past 7 days.
> Alias: sh

| Arguments | Hints | Examples |
| -------- | -------- | -------- |
| keyword | Please enclose any spaces in double quotes | ```search "cat dog"``` |
| counts | If you omit it, the default value is given in the configuration file | ```search sushi 5``` |

## favorite
**favorite  [Subcommand] [<tweet number>]**

Manipulate "like".
> Alias: like, fv

| Subcommand | Alias | Description |
| -------- | -------- | -------- |
| none |  | Like tweet |
| remove | rm | UnLike tweet |

| Arguments | Hints | Examples |
| -------- | -------- | -------- |
| tweet number | Separate each tweet with a space if there is more than one | ```favorite 1 2``` |

## retweet
**retweet [Subcommand] [<tweet number>]**

Manipulate retweets.
> Alias: rt

| Subcommand | Alias | Description | Argument |
| -------- | -------- | -------- | -------- |
| none |  | Retweet tweet | ```retweet [<tweet number>]``` |
| quote | qt | Quote tweet | ```retweet quote [<tweet number>] [text] [image file]``` |
| remove | rm | UnRetweet tweet | ```retweet remove [<tweet number>]``` |

| Arguments | Hints | Examples |
| -------- | -------- | -------- |
| text | If there is no text and image file, the message will be posted with a "にゃーん" | ```retweet quote 1``` |
| image file | If there is more than one, please separate them with a space | ```retweet quote 1 🐾 paw_pad.png footprints.png``` |
| tweet number | Separate each tweet with a space if there is more than one | ```retweet 1 5``` |

## reply
**reply [<tweet number>] [text] [image file]**

Post a reply.
> Alias: rp

| Arguments | Hints | Examples |
| -------- | -------- | -------- |
| text | If there is no text and image file, the message will be posted with a "にゃーん" | ```reply 1``` |
| image file | If there is more than one, please separate them with a space | ```reply good!!! sushi1.png sushi2.png``` |

- You can also omit the text and just post an image. (e.g. ```reply dog.png```)

## follow
**follow [Subcommand] [<tweet number / userID>]**

Performs a follow operation.
> Alias: fw

| Subcommand | Alias | Description |
| -------- | -------- | -------- |
| none | | Follow user |
| remove | rm | Unfollow user |

| Arguments | Hints | Examples |
| -------- | -------- | -------- |
| tweet number | Follow the author of the specified tweet | ```follow 1``` |
| userID | Follow users with the user ID you entered | ```follow arrow_2nd``` |

## block
**block [Subcommand] [<tweet number / userID>]**

Performs a block operation.
> Alias: bk

| Subcommand | Alias | Description |
| -------- | -------- | -------- |
| none | | Block user |
| remove | rm | Unblock user |

| Arguments | Hints | Examples |
| -------- | -------- | -------- |
| tweet number | Block the author of the specified tweet | ```block 1``` |
| userID | Block users with the user ID you entered | ```block arrow_2nd``` |

## mute
**mute [Subcommand] [<tweet number / userID>]**

Performs a mute operation.
> Alias: mu

| Subcommand | Alias | Description |
| -------- | -------- | -------- |
| none | | Mute user |
| remove | rm | Unmute user |

| Arguments | Hints | Examples |
| -------- | -------- | -------- |
| tweet number | Mute the author of the specified tweet | ```mute 1``` |
| userID | Mute users with the user ID you entered | ```mute arrow_2nd``` |

## open
**open [<tweet number>]**

View the tweet in your browser.
> Alias: op

## config
**config [<Subcommand>]**

Manipulation of configuration files

| Subcommand | Alias | Description |
| -------- | -------- | -------- |
| reset    | Regenerate the configuration file | ```config reset``` |
| remove   | Deletes the configuration file | ```config remove``` |

</details>

## Configuration Files
The configuration file is saved directly under your home directory as ```.twnyan.yaml```

<details>
<summary>Open</summary>

## ColorData
Color setting.

Specify it with a hexadecimal color code.

| Name | Description |
| -------- | -------- |
| Accent1 | Background color of tweet numbers, etc. |
| Accent2 | Posting times for tweets, etc. |
| Accent3 | Somewhere |
| BoxFg | Text color of tweet numbers, etc. |
| UserName | Username |
| UserID | UserID |
| Text | Tweet |
| Separator |Separator（--------） |
| Reply | Reply ID and Reply Display |
| Hashtag | Hashtag |
| Fav | Likes |
| RT | Number of retweets and display of retweets |
| Verified | Authenticated Users |
| Protected | Protected Users |
| Follow | Following and FollowedBy |
| Block | Blocking |
| Mute | Muting |
 
## DefaultData
The default value is set.

| Name | Description |
| -------- | -------- |
| Counts | Default number of fetches |
| Prompt | Prompt character |
| DateFormat | Date Format |
| TimeFormat | Time Format |

The format is the same as the format string of the [time package](https://golang.org/pkg/time/#pkg-constants)

</details>
