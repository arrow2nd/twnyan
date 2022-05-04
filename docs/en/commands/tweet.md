# tweet

Post a tweet.

```
twnyan tweet {[text] [image...] | <command>}
twnyan tw {[text] [image...] | <command>}
```

- Pipe input is also supported (e.g. `echo "nyaan..." | twnyan tweet`)
- If no text is specified, "にゃーん" will be posted
- You can also submit images only (e.g. `tweet cat.png`)
- If there are multiple images, please specify them separated by a single space

## Command

### tweet multi

Post a multi-line tweet.

```
twnyan tweet multi [image...]
twnyan tweet ml [image...]
```

- If there are multiple images, please specify them separated by a single space
- To finish typing, type a semicolon `;` at the end of the sentence
- Enter `:exit` to cancel entry

### tweet remove

Remove a tweet.

```
twnyan tweet remove <tweet-number>...
twnyan tweet rm <tweet-number>...
```

- If there are multiple tweet numbers, please specify them separated by spaces
