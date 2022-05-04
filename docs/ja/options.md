# 設定ファイルについて

設定ファイルが置かれるディレクトリは、ホームディレクトリ直下に `.twnyan` として作成されます。

```
.twnyan
├── .cred.yaml
├── color.yaml
└── option.yaml
```

## .cred.yaml

アカウントの認証情報が含まれたファイルです。

## option.yaml

オプション設定のファイルです。

| 名前       | 説明                   |
| ---------- | ---------------------- |
| ConfigDir  | 設定ディレクトリのパス |
| Counts     | デフォルトの取得件数   |
| DateFormat | 日付のフォーマット     |
| TimeFormat | 時刻のフォーマット     |

- 日付、時刻のフォーマットは[time パッケージのフォーマット文字列](https://golang.org/pkg/time/#pkg-constants)と同じ書式です

## color.yaml

色設定のファイルです。

| 名前         | 説明                   |
| ------------ | ---------------------- |
| Accent1      | アクセント 1           |
| Accent2      | アクセント 2           |
| Accent3      | アクセント 3           |
| Error        | エラーメッセージ背景色 |
| BoxForground | 反転時の文字色         |
| Separator    | セパレータ             |
| UserName     | ユーザ名               |
| ScreenName   | スクリーンネーム       |
| Reply        | リプライ表示           |
| Hashtag      | ハッシュタグ           |
| Favorite     | いいね表示             |
| Retweet      | リツイート表示         |
| Verified     | 認証済みアカウント     |
| Protected    | 鍵アカウント           |
| Following    | フォロー中表示         |
| FollowedBy   | 被フォロー表示         |
| Block        | ブロック表示           |
| Mute         | ミュート表示           |
