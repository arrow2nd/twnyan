# About the configuration file

The directory where the configuration files are placed is created as `.twnyan` directly under the home directory.

```
.twnyan
├── .cred.yaml
├── color.yaml
└── option.yaml
```

## .cred.yaml

A file containing account credentials.

## option.yaml

Optional settings file.

| 名前       | 説明                                |
| ---------- | ----------------------------------- |
| ConfigDir  | Path of the configuration directory |
| Counts     | Default number of acquisitions      |
| DateFormat | Date Format                         |
| TimeFormat | Time Format                         |

- Date and time formats are in the same format as the format string in the [time package](https://golang.org/pkg/time/#pkg-constants)

## color.yaml

Color setting file.

| 名前         | 説明                           |
| ------------ | ------------------------------ |
| Accent1      | Accent Color 1                 |
| Accent2      | Accent Color 2                 |
| Accent3      | Accent Color 3                 |
| Error        | Error Message Background Color |
| BoxForground | Text Color When Reversed       |
| Separator    | Separator                      |
| UserName     | User Name                      |
| ScreenName   | Screen Name                    |
| Reply        | Reply                          |
| Hashtag      | Hashtag                        |
| Favorite     | Like                           |
| Retweet      | Retweet                        |
| Verified     | Verified Account               |
| Protected    | Private Account                |
| Following    | Following                      |
| FollowedBy   | Followed By                    |
| Block        | Block                          |
| Mute         | Mute                           |
