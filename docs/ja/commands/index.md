# twnyan

twnyan はターミナルで動く、シンプルな Twitter クライアントです。

```
twnyan [フラグオプション] [コマンド]
```

## 例

```
# メインアカウントでツイート
twnyan tweet hello,nekochan!

# @subaccount でツイート
twnyan --account subaccount tweet "I want to eat sushi..."
```

## フラグオプション

- `-a` `--account <ユーザ名>`
  - 操作に使用するアカウントを指定 (省略した場合、メインアカウントが指定されます)
- `-h` `--help`
  - オプションを含むヘルプを表示

## コマンド

### アカウント

- [account](./account.md)

### ツイート

- [tweet](./tweet.md)
- [reply](./reply.md)
- [like](./like.md)
- [retweet](./retweet.md)
- [likert](./likert.md)
- [quote](./quote.md)
- [open](./open.md)

### ユーザ

- [follow](./follow.md)
- [block](./block.md)
- [mute](./mute.md)

### タイムライン

- [timeline](./timeline.md)
- [stream](./stream.md)
- [mention](./mention.md)
- [list](./list.md)
- [user](./user.md)
- [search](./search.md)

### その他

- [clear](./clear.md)
- [help](./help.md)
- [version](./version.md)
- [exit](./exit.md)
