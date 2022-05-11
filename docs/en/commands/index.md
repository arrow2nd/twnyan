# twnyan

twnyan is a simple Twitter client that runs in the terminal.

```
twnyan [Option] [Command]
```

## Example

```
# Tweet in your main account
twnyan tweet hello,nekochan!

# Tweet with @subaccount
twnyan -A subaccount tweet "I want to eat sushi..."
```

## Option

- `-A` `--account <username>`
  - Specify the account to be used for the operation (if omitted, the main account will be specified)
- `-H` `--help`
  - Display help with options

## Command

### Accounts

- [account](./account.md)

### Tweet

- [tweet](./tweet.md)
- [reply](./reply.md)
- [like](./like.md)
- [retweet](./retweet.md)
- [likert](./likert.md)
- [quote](./quote.md)
- [open](./open.md)

### User

- [follow](./follow.md)
- [block](./block.md)
- [mute](./mute.md)

### Timeline

- [timeline](./timeline.md)
- [stream](./stream.md)
- [mention](./mention.md)
- [list](./list.md)
- [user](./user.md)
- [search](./search.md)

### Other

- [clear](./clear.md)
- [help](./help.md)
- [version](./version.md)
- [exit](./exit.md)
