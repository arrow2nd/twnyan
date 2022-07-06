# 🐈 twnyan

いつでも「にゃーん」したいねこのための Twitter クライアント

[![release](https://github.com/arrow2nd/twnyan/actions/workflows/release.yml/badge.svg)](https://github.com/arrow2nd/twnyan/actions/workflows/release.yml)
[![test](https://github.com/arrow2nd/twnyan/actions/workflows/test.yml/badge.svg)](https://github.com/arrow2nd/twnyan/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/arrow2nd/twnyan)](https://goreportcard.com/report/github.com/arrow2nd/twnyan)
![GitHub all releases](https://img.shields.io/github/downloads/arrow2nd/twnyan/total)
[![GitHub license](https://img.shields.io/github/license/arrow2nd/twnyan)](https://github.com/arrow2nd/twnyan/blob/main/LICENSE)

> **[English](README_EN.md)**

![twnyan](https://user-images.githubusercontent.com/44780846/106699506-612c0f80-6626-11eb-803e-332512822789.gif)

## 特徴

- いつでも「にゃーん」できる 🐾
- 疑似 UserStream モード
- マルチアカウント対応
- コマンドライン / 対話モード両対応
- とってもカラフル 🎨

## 動作条件

- Windows / macOS / Linux
- 絵文字が表示可能なターミナル

## インストール

### Homebrew

```sh
brew tap arrow2nd/tap
brew install twnyan
```

### Scoop

```
scoop bucket add arrow2nd https://github.com/arrow2nd/scoop-bucket.git
scoop install arrow2nd/twnyan
```

### Go

```sh
go install github.com/arrow2nd/twnyan@latest
```

### バイナリファイル

[Releases](https://github.com/arrow2nd/twnyan/releases) からお使いの環境にあったファイルをダウンロードしてください。

## アンインストール

### Homebrew

```sh
brew uninstall twnyan
rm -rf ~/.twnyan # 設定ファイル群を削除
```

### Go

```sh
go clean -i github.com/arrow2nd/twnyan
rm -rf $GOPATH/src/github.com/arrow2nd/twnyan
rm -rf ~/.twnyan
```

## 初期設定

![auth](https://user-images.githubusercontent.com/44780846/106747441-4a59dd00-6667-11eb-8248-3468cb39f7d1.png)

1. 初回起動時に表示される認証ページの URL にブラウザでアクセス
2. 画面に従って手順を進め、表示される PIN コードを twnyan に入力
3. 完了！ 😺

> 最初に追加したアカウントは **「メインアカウント」** として扱われます

## 使い方

### コマンドラインモード

![cmdline](https://user-images.githubusercontent.com/44780846/106699170-b287cf00-6625-11eb-8374-8565286db3e2.gif)

- 一部のコマンドは使用できません

### 対話モード

コマンドを指定しない場合、対話モードで起動します。

### にゃーん

![nyaan](https://user-images.githubusercontent.com/44780846/106699001-558c1900-6625-11eb-948e-6212ab0cba40.gif)

```
twnyan tw
```

ツイート文を省略すると「にゃーん」とツイートされます。

- リプライや引用リツイートでも同様の動作をします
- **画像が添付されている場合は「にゃーん」しません**

### 詳しい使い方

- [コマンド一覧](./docs/ja/commands/index.md)
- [設定ファイルについて](./docs/ja/options.md)
