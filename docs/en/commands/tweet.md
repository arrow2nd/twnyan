# tweet

Post tweet.

```
twnyan tweet {[Text] [Image...] | <Command>}
twnyan tw {[Text] [Image...] | <Command>}
```

- If no text is specified, "にゃーん" will be posted
- You can also submit images only (e.g. `tweet cat.png`)
- If there are multiple images, please specify them separated by a single space

## Command

### tweet multi

Post a multi-line tweet.

```
twnyan tweet multi [Image...]
twnyan tweet ml [Image...]
```

- If there are multiple images, please specify them separated by a single space
- To finish typing, type a semicolon `;` at the end of the sentence
- Enter `:exit` to cancel entry

### tweet remove

Remove tweet.

```
twnyan tweet remove [<Tweet Number>...]
twnyan tweet rm [<Tweet Number>...]
```

- If there are multiple tweet numbers, please specify them separated by spaces

### tweet help

Displays help.

```
twnyan tweet help
```
