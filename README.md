# twnyan
[![arrow2nd](https://circleci.com/gh/arrow2nd/twnyan.svg?style=shield)](https://circleci.com/gh/arrow2nd/twnyan/tree/main)
[![Go Report Card](https://goreportcard.com/badge/github.com/arrow2nd/twnyan)](https://goreportcard.com/report/github.com/arrow2nd/twnyan)

This is a Twitter client by cats for cats🐾

> **[日本語](README_JP.md)**

## Features
- Multi-byte character support
- It's easy to tweet "にゃーん"
- ~Needlessly~ flexible color settings
- Support for interactive mode

## Screenshot
<image src="https://user-images.githubusercontent.com/44780846/102610268-3b2ed680-4170-11eb-8095-56811ee546b4.gif" width=70%>

## How to install
**(Recommended)**

```$ go get -u github.com/arrow2nd/twnyan```

### Use binary files
Download a zip file from ReleasePage that fits your environment and pass the Path through the binary file.

## Operating conditions
- Windows/Linux
- A terminal that can display emoji

### Remarks
- I have not been able to confirm that it works on a Mac
- To run in a WSL environment, you need to be able to use xdg-open

## initialization
<image src="https://user-images.githubusercontent.com/44780846/102592746-432e4c80-4157-11eb-8581-29a1f8f850c9.png" width=60%>

The first time you start the program, the browser will start and the authentication page will be displayed.

Follow the steps on the screen and enter the displayed PIN code into twnyan.

## Usage
```$ twnyan [command] [argument]```

If you omit the command ```$ twnyan```, it will start in interactive mode.

You can manipulate tweets (like, RT, etc.) by specifying the number of the tweet.

## Command list

<details>
<summary>Open</summary>

## tweet
```Alias: tw```

Manipulates tweets.
### tweet [text] [imagefile]
Post tweet.

| Arguments  | Hints                                                                            | Examples                            |
| ---------- | -------------------------------------------------------------------------------- | ----------------------------------- |
| text       | If there is no text and image file, the message will be posted with a "にゃーん" | ```tweet```                         |
| image file | If there is more than one, please separate them with a space                     | ```tweet 🍣 sushi1.png sushi2.png``` |

- You can also omit the text and just post an image (e.g. ```tweet cat.png```)

### tweet remove [\<tweetnumber\>]...
```Alias: rm```

Delete tweet.

| Arguments   | Hints                                                      | Examples               |
| ----------- | ---------------------------------------------------------- | ---------------------- |
| tweetNumber | Separate each tweet with a space if there is more than one | ```tweet remove 2 5``` |

## timeline
```Alias: tl```

Get a home timeline.
### timeline [counts]

| Arguments | Hints                                                                | Examples          |
| --------- | -------------------------------------------------------------------- | ----------------- |
| counts    | If you omit it, the default value is given in the configuration file | ```timeline 39``` |

## mention
```Alias: mt```

Get a Mentions to you.
### mention [counts]

| Arguments | Hints                                                                | Examples         |
| --------- | -------------------------------------------------------------------- | ---------------- |
| counts    | If you omit it, the default value is given in the configuration file | ```mention 20``` |

## list
```Alias: ls```

Get a timeline of the list.
### list [\<listname\>] [counts]

| Arguments | Hints                                                                        | Examples                      |
| --------- | ---------------------------------------------------------------------------- | ----------------------------- |
| list name | If you are running in interactive mode, you can complete it with the Tab key | ```list Cats```               |
| counts    | If you omit it, the default value is given in the configuration file         | ```list "Cat Gathering" 30``` |

## user
```Alias: ur```

Get a user timeline.
### user [<username/tweetnumber>] [counts]
Get a timeline of the specified user.

| Arguments            | Hints                                                                | Examples                          |
| -------------------- | -------------------------------------------------------------------- | --------------------------------- |
| username/tweetnumber | Either can be specified<br>The '@' in the username is optional       | ```user github```<br>```user 1``` |
| counts               | If you omit it, the default value is given in the configuration file | ```user twitter 15```             |

### user myuser [counts]
Get your own timeline.

| Arguments | Hints                                                                | Examples             |
| --------- | -------------------------------------------------------------------- | -------------------- |
| counts    | If you omit it, the default value is given in the configuration file | ```user myuser 50``` |

## search
```Alias: sh```

Searches for tweets tweets in the past 7 days.
### search [\<keyword\>] [counts]

| Arguments | Hints                                                                | Examples               |
| --------- | -------------------------------------------------------------------- | ---------------------- |
| keyword   | Please enclose any spaces in double quotes                           | ```search "cat dog"``` |
| counts    | If you omit it, the default value is given in the configuration file | ```search sushi 5```   |

## favorite
```Alias: fv, like```

Manipulate "like".
### favorite [\<tweetnumber\>]
Like tweet.

| Arguments   | Hints                                                      | Examples           |
| ----------- | ---------------------------------------------------------- | ------------------ |
| tweetnumber | Separate each tweet with a space if there is more than one | ```favorite 1 2``` |

### favorite remove [\<tweetnumber\>]
```Alias: rm```

UnLike tweet.

| Arguments   | Hints                                                      | Examples                  |
| ----------- | ---------------------------------------------------------- | ------------------------- |
| tweetnumber | Separate each tweet with a space if there is more than one | ```favorite remove 1 2``` |

## retweet
```Alias: rt```

Manipulate retweets.
### retweet [\<tweetnumber\>]...
Retweet tweet.

| Arguments   | Hints                                                      | Examples          |
| ----------- | ---------------------------------------------------------- | ----------------- |
| tweetnumber | Separate each tweet with a space if there is more than one | ```retweet 1 5``` |

### retweet quote [\<tweetnumber\>] [text] [imagefile]
```Alias: qt```

Quote tweet.

| Arguments   | Hints                                                                            | Examples                                      |
| ----------- | -------------------------------------------------------------------------------- | --------------------------------------------- |
| tweetnumber | Specify the number of the tweet to quote                                         | ```retweet quote 1 good!!!```                 |
| text        | If there is no text and image file, the message will be posted with a "にゃーん" | ```retweet quote 1```                         |
| imagefile   | If there is more than one, please separate them with a space                     | ```retweet quote 1 🍣 sushi1.png sushi2.png``` |

### retweet remove [\<tweetnumber\>]...
```Alias: rm```

UnRetweet tweet.

| Arguments   | Hints                                                      | Examples                 |
| ----------- | ---------------------------------------------------------- | ------------------------ |
| tweetnumber | Separate each tweet with a space if there is more than one | ```retweet remove 1 5``` |

## reply
```Alias: rp```

Post a reply.
### reply [\<tweetnumber\>] [text] [imagefile]

| Arguments   | Hints                                                                            | Examples                         |
| ----------- | -------------------------------------------------------------------------------- | -------------------------------- |
| tweetnumber | Specify the number of the tweet you want to reply to.                            | ```reply 1 meow```               |
| text        | If there is no text and image file, the message will be posted with a "にゃーん" | ```reply 1```                    |
| image file  | If there is more than one, please separate them with a space                     | ```reply 1 good!!! sushi1.png``` |

- You can also omit the text and just post an image. (e.g. ```reply 1 dog.png```)

## follow
```Alias: fw```

Performs a follow operation.
### follow [<username/tweetnumber>]
Follow user.

| Arguments            | Hints                                                          | Examples                              |
| -------------------- | -------------------------------------------------------------- | ------------------------------------- |
| username/tweetnumber | Either can be specified<br>The '@' in the username is optional | ```follow github```<br>```follow 1``` |

### follow remove [<username/tweetnumber>]
```Alias: rm```

Unfollow user.

| Arguments            | Hints                                                          | Examples                                               |
| -------------------- | -------------------------------------------------------------- | ------------------------------------------------------ |
| username/tweetnumber | Either can be specified<br>The '@' in the username is optional | ```follow remove arrow_2nd```<br>```follow remove 1``` |

## block
```Alias: bk```

Performs a block operation.
### block [<username/tweetnumber>]
Block user.

| Arguments            | Hints                                                          | Examples                               |
| -------------------- | -------------------------------------------------------------- | -------------------------------------- |
| username/tweetnumber | Either can be specified<br>The '@' in the username is optional | ```block arrow_2nd```<br>```block 1``` |

### block remove [<username/tweetnumber>]
```Alias: rm```

Unblock user.

| Arguments            | Hints                                                          | Examples                                             |
| -------------------- | -------------------------------------------------------------- | ---------------------------------------------------- |
| username/tweetnumber | Either can be specified<br>The '@' in the username is optional | ```block remove arrow_2nd```<br>```block remove 1``` |

## mute
```Alias: mu```

Performs a mute operation.
### mute [<username/tweetnumber>]
Mute user.

| Arguments            | Hints                                                          | Examples                             |
| -------------------- | -------------------------------------------------------------- | ------------------------------------ |
| username/tweetnumber | Either can be specified<br>The '@' in the username is optional | ```mute arrow_2nd```<br>```mute 1``` |

### mute remove [<username/tweetnumber>]
```Alias: rm```

Unmute user.

| Arguments            | Hints                                                          | Examples                                           |
| -------------------- | -------------------------------------------------------------- | -------------------------------------------------- |
| username/tweetnumber | Either can be specified<br>The '@' in the username is optional | ```mute remove arrow_2nd```<br>```mute remove 1``` |

## open
```Alias: op```

View the tweet in your browser.
### open [\<tweetnumber\>]

| Arguments   | Hints                                                          | Examples     |
| ----------- | -------------------------------------------------------------- | ------------ |
| tweetnumber | Specify the number of the tweet to be displayed in the browser | ```open 2``` |

## export
```Alias: ep```

Specify a file name excluding the file extension
Exports the currently displayed timeline or tweets to a file.
### export [\<format\>] [\<filename\>]

| Arguments | Hints                                                                                                      | Examples                 |
| --------- | ---------------------------------------------------------------------------------------------------------- | ------------------------ |
| format    | json/yaml can be specified<br>If you are running in interactive mode, you can complete it with the Tab key | ```export json tweets``` |
| filename  | Specify a file name excluding the file extension                                                           | ```export yaml test```   |

## config
Manipulation of configuration files.
### config reset
Regenerate the configuration file.
### config remove
Deletes the configuration file.

</details>

## Configuration Files
The configuration file is saved directly under your home directory as ```.twnyan.yaml```

<details>
<summary>Open</summary>

## ColorData
Color setting.

Specify it with a hexadecimal color code.

| Name      | Description                                |
| --------- | ------------------------------------------ |
| Accent1   | Background color of tweetNos, etc.         |
| Accent2   | Posting times for tweets, etc.             |
| Accent3   | Somewhere                                  |
| BoxFg     | Text color of tweetNos, etc.               |
| UserName  | Username                                   |
| UserID    | UserID                                     |
| Separator | Separator（--------）                      |
| Reply     | Reply ID and Reply Display                 |
| Hashtag   | Hashtag                                    |
| Fav       | Likes                                      |
| RT        | Number of retweets and display of retweets |
| Verified  | Authenticated Users                        |
| Protected | Protected Users                            |
| Follow    | Following and FollowedBy                   |
| Block     | Blocking                                   |
| Mute      | Muting                                     |
 
## DefaultData
The default value is set.

| Name       | Description               |
| ---------- | ------------------------- |
| Counts     | Default number of fetches |
| Prompt     | Prompt character          |
| DateFormat | Date Format               |
| TimeFormat | Time Format               |

The format is the same as the format string of the [time package](https://golang.org/pkg/time/#pkg-constants)

</details>
