# quote

> Only available in interactive mode

Quote a tweet.

```
twnyan quote {<tweet-number> [text] []... | <command>}
twnyan qt {<tweet-number> [text] [image]... | <command>}
```

- If no text is specified, "にゃーん" will be posted
- If there are multiple images, please specify them separated by a single space

## Command

### quote multi

Post a multi-line quote retweet.

```
twnyan quote multi [image]...
twnyan quote ml [image]...
```

- If there are multiple images, please specify them separated by a single space
- To finish typing, type a semicolon `;` at the end of the sentence
- Enter `:exit` to cancel entry
