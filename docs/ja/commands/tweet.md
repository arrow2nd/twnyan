# tweet

ツイートを投稿します。

```
twnyan tweet {[テキスト] [画像...] | <コマンド>}
twnyan tw {[テキスト] [画像...] | <コマンド>}
```

- パイプ入力にも対応しています (e.g. `echo "nyaan..." | twnyan tweet`)
- テキストを指定しない場合、「にゃーん」が投稿されます
- 画像のみの投稿も可能です (e.g. `tweet cat.png`)
- 画像が複数ある場合、半角スペースで区切って指定してください

## コマンド

### tweet multi

複数行のツイートを投稿します。

```
twnyan tweet multi [画像...]
twnyan tweet ml [画像...]
```

- 画像が複数ある場合、半角スペースで区切って指定してください
- 入力を終了する場合、文末にセミコロン `;` を入力してください
- 入力をキャンセルする場合、`:exit` を入力してください

### tweet remove

ツイートを削除します。

```
twnyan tweet remove <ツイート番号>...
twnyan tweet rm <ツイート番号>...
```

- ツイート番号が複数ある場合は、スペース区切りで指定してください
