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
twnyan -A subaccount tweet "I want to eat sushi..."
```

## フラグオプション

- `-A` `--account <ユーザ名>`
  - 使用するアカウントを指定 (指定しない場合、メインアカウントが使用されます)
- `-H` `--help`
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
