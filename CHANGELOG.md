# Change Log

## [Unreleased]

## [1.9.5] - 2022-12-31

### Security

- 依存関係を更新

## [1.9.4] - 2022-08-02

### Fixed

- 特定のツイートにおいて、ハッシュタグのハイライト処理に失敗する

## [1.9.3] - 2022-07-06

### Added

- Scoop での配布を開始

### Changed

- Windows 向けファイルの圧縮形式を zip に変更

### Security

- 依存関係を更新

## [1.9.2] - 2022-06-17

### Fixed

- 特定のツイートにおいて、ハッシュタグのハイライト処理に失敗する

### Security

- 依存関係を更新

## [1.9.1] - 2022-05-26

### Security

- 依存関係を更新

## [1.9.0] - 2022-05-17

### Added

- コマンドラインモードで使用できないコマンドがコマンドラインから呼び出された際、エラーメッセージを表示

### Changed

- オプションフラグのショートハンドを変更

## [1.8.2] - 2022-05-10

### Fixed

- 一定の条件を満たすハッシュタグ付きツイートを表示するとエラーが発生する

## [1.8.1] - 2022-05-04

### Changed

- `account list` の表示形式を変更

## [1.8.0] - 2022-05-04

### Security

- 依存関係を更新

### Added

- マルチアカウントに対応 [#15](https://github.com/arrow2nd/twnyan/issues/15)
- `tweet` コマンドをパイプ入力に対応 [#16](https://github.com/arrow2nd/twnyan/issues/16)

### Changed

- ヘルプ文を変更
- ドキュメントを整備

### Fixed

- 一部ハッシュタグがハイライトされない
- ツイートが無いユーザのタイムラインを表示するとエラーになる

## [1.7.2] - 2022-03-29

### Security

- 依存関係を更新

## [1.7.1] - 2021-12-14

### Fixed

- ツイート番号を指定してのユーザ操作ができない

## [1.7.0] - 2021-12-12

### Added

- いいねと RT を同時に行う likert コマンドを追加

## [1.6.0] - 2021-11-14

### Changed

- エラーメッセージの出力先を変更

## [1.5.2] - 2021-11-08

### Changed

- 自動リリースの環境を変更

## [1.5.1] - 2021-11-08

### Added

- Homebrew でのインストールに対応

### Fixed

- ハッシュタグのハイライトを修正

## [1.5.0] - 2021-08-12

### Changed

- multi コマンドの仕様を変更

## [1.4.6] - 2021-07-01

### Fixed

- 今日の投稿されたツイートに日付が表示される

## [1.4.5] - 2021-04-27

### Fixed

- アンダースコア・一部の漢字を含むハッシュタグが正しくハイライトされない

## [1.4.4] - 2021-04-19

### Fixed

- ハイライトされないハッシュタグがある

## [1.4.3] - 2021-04-13

### Fixed

- コマンドが見つからなかったときの挙動を修正

## [1.4.2] - 2021-04-13

### Fixed

- stream コマンド実行中にレート制限に引っかかった時の挙動を修正

## [1.4.1] - 2021-04-13

### Fixed

- デバッグログの消し忘れを修正

### Changed

- アカウント種類の表示形式を変更

## [1.4.0] - 2021-04-12

### Added

- UserStream を再現する stream コマンドを追加

## [1.3.0] - 2021-04-03

### Added

- 複数行入力をキャンセルする機能を追加

## [1.2.5] - 2021-03-22

### Fixed

- ハイライトされないハッシュタグがある

## [1.2.4] - 2021-03-13

### Fixed

- Go 1.16.x で go install が失敗する

## [1.2.3] - 2021-02-12

### Fixed

- ハッシュタグ等が途中で改行されるとハイライトされない
- アップロード処理を修正

## [1.2.2] - 2021-02-04

### Added

- quote コマンドを追加
- 複数行のツイートに画像添付機能を追加
- リプライ、引用 RT も複数行に対応

### Fixed

- ヘルプ文を修正

## [1.2.1] - 2021-02-03

### Fixed

- Windows 環境で起動できない
- リツイートの取り消しに失敗する
- API のエラーメッセージがパースできない

### Changed

- API 制限時にエラーメッセージを表示するよう変更

## [1.2.0] - 2021-02-03

### Added

- 複数行のツイートに対応

### Removed

- Export コマンドを削除
- Config コマンドを削除

### Changed

- 表示のデザインを変更
- 設定ファイルのパス、形式を変更

## [1.1.1] - 2020-12-30

### Added

- 設定ファイル操作時の確認を追加

### Changed

- デフォルト色を変更
- 一部デザインを変更

## [1.1.0] - 2020-12-18

### Added

- json/yaml 形式でのファイル出力機能

### Removed

- 形式を指定して取得したデータを標準出力する機能

### Changed

- コマンドのヘルプ文
- システムメッセージの表示形式

## [1.0.2] - 2020-12-17

### Fixed

- メンションのハイライトが正しくない

## [1.0.1] - 2020-12-17

### Fixed

- ハッシュタグのハイライトが正しくない

## [1.0.0] - 2020-12-03

### Fixed

- テキスト色が反映されない

### Removed

- .twnyan.yaml 内の色の設定からテキストの項目を削除

## [0.0.4] - 2020-12-01

### Fixed

- ハッシュタグ・メンションのハイライトが正しくない

## [0.0.3] - 2020-11-20

### Added

- 形式を指定して取得したデータを標準出力する機能

### Fixed

- ハッシュタグのハイライトが正しくない

## [0.0.2] - 2020-11-19

### Fixed

- 日本語環境での絵文字入力時の挙動がおかしい
- `go get`でインストールした際、正しい依存パッケージが取得できない

## 0.0.1 - 2020-11-17

- リリースしました！

[unreleased]: https://github.com/arrow2nd/twnyan/compare/v1.9.5...HEAD
[1.9.5]: https://github.com/arrow2nd/twnyan/compare/v1.9.4...v1.9.5
[1.9.4]: https://github.com/arrow2nd/twnyan/compare/v1.9.3...v1.9.4
[1.9.3]: https://github.com/arrow2nd/twnyan/compare/v1.9.2...v1.9.3
[1.9.2]: https://github.com/arrow2nd/twnyan/compare/v1.9.1...v1.9.2
[1.9.1]: https://github.com/arrow2nd/twnyan/compare/v1.9.0...v1.9.1
[1.9.0]: https://github.com/arrow2nd/twnyan/compare/v1.8.2...v1.9.0
[1.8.2]: https://github.com/arrow2nd/twnyan/compare/v1.8.1...v1.8.2
[1.8.1]: https://github.com/arrow2nd/twnyan/compare/v1.8.0...v1.8.1
[1.8.0]: https://github.com/arrow2nd/twnyan/compare/v1.7.2...v1.8.0
[1.7.2]: https://github.com/arrow2nd/twnyan/compare/v1.7.1...v1.7.2
[1.7.1]: https://github.com/arrow2nd/twnyan/compare/v1.7.0...v1.7.1
[1.7.0]: https://github.com/arrow2nd/twnyan/compare/v1.6.0...v1.7.0
[1.6.0]: https://github.com/arrow2nd/twnyan/compare/v1.5.2...v1.6.0
[1.5.2]: https://github.com/arrow2nd/twnyan/compare/v1.5.1...v1.5.2
[1.5.1]: https://github.com/arrow2nd/twnyan/compare/v1.5.0...v1.5.1
[1.5.0]: https://github.com/arrow2nd/twnyan/compare/v1.4.6...v1.5.0
[1.4.6]: https://github.com/arrow2nd/twnyan/compare/v1.4.5...v1.4.6
[1.4.5]: https://github.com/arrow2nd/twnyan/compare/v1.4.4...v1.4.5
[1.4.4]: https://github.com/arrow2nd/twnyan/compare/v1.4.3...v1.4.4
[1.4.3]: https://github.com/arrow2nd/twnyan/compare/v1.4.2...v1.4.3
[1.4.2]: https://github.com/arrow2nd/twnyan/compare/v1.4.1...v1.4.2
[1.4.1]: https://github.com/arrow2nd/twnyan/compare/v1.4.0...v1.4.1
[1.4.0]: https://github.com/arrow2nd/twnyan/compare/v1.3.0...v1.4.0
[1.3.0]: https://github.com/arrow2nd/twnyan/compare/v1.2.5...v1.3.0
[1.2.5]: https://github.com/arrow2nd/twnyan/compare/v1.2.4...v1.2.5
[1.2.4]: https://github.com/arrow2nd/twnyan/compare/v1.2.3...v1.2.4
[1.2.3]: https://github.com/arrow2nd/twnyan/compare/v1.2.2...v1.2.3
[1.2.2]: https://github.com/arrow2nd/twnyan/compare/v1.2.1...v1.2.2
[1.2.1]: https://github.com/arrow2nd/twnyan/compare/v1.2.0...v1.2.1
[1.2.0]: https://github.com/arrow2nd/twnyan/compare/v1.1.1...v1.2.0
[1.1.1]: https://github.com/arrow2nd/twnyan/compare/v1.1.0...v1.1.1
[1.1.0]: https://github.com/arrow2nd/twnyan/compare/v1.0.2...v1.1.0
[1.0.2]: https://github.com/arrow2nd/twnyan/compare/v1.0.1...v1.0.2
[1.0.1]: https://github.com/arrow2nd/twnyan/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/arrow2nd/twnyan/compare/v0.0.4...v1.0.0
[0.0.4]: https://github.com/arrow2nd/twnyan/compare/v0.0.3...v0.0.4
[0.0.3]: https://github.com/arrow2nd/twnyan/compare/v0.0.2...v0.0.3
[0.0.2]: https://github.com/arrow2nd/twnyan/compare/v0.0.1...v0.0.2
