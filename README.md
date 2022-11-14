# ð twnyan

> **Warning**
>
> twnyanã¯ã¡ã³ããã³ã¹ã¢ã¼ãã§ãã
> [nekome](https://github.com/arrow2nd/nekome)ã¸ã®ç§»è¡ããããããã¾ãã

ãã¤ã§ããã«ãã¼ãããããã­ãã®ããã® Twitter ã¯ã©ã¤ã¢ã³ã

[![release](https://github.com/arrow2nd/twnyan/actions/workflows/release.yml/badge.svg)](https://github.com/arrow2nd/twnyan/actions/workflows/release.yml)
[![test](https://github.com/arrow2nd/twnyan/actions/workflows/test.yml/badge.svg)](https://github.com/arrow2nd/twnyan/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/arrow2nd/twnyan)](https://goreportcard.com/report/github.com/arrow2nd/twnyan)
![GitHub all releases](https://img.shields.io/github/downloads/arrow2nd/twnyan/total)
[![GitHub license](https://img.shields.io/github/license/arrow2nd/twnyan)](https://github.com/arrow2nd/twnyan/blob/main/LICENSE)

> **[English](README_EN.md)**

![twnyan](https://user-images.githubusercontent.com/44780846/106699506-612c0f80-6626-11eb-803e-332512822789.gif)

## ç¹å¾´

- ãã¤ã§ããã«ãã¼ããã§ãã ð¾
- çä¼¼ UserStream ã¢ã¼ã
- ãã«ãã¢ã«ã¦ã³ãå¯¾å¿
- ã³ãã³ãã©ã¤ã³ / å¯¾è©±ã¢ã¼ãä¸¡å¯¾å¿
- ã¨ã£ã¦ãã«ã©ãã« ð¨

## åä½æ¡ä»¶

- Windows / macOS / Linux
- çµµæå­ãè¡¨ç¤ºå¯è½ãªã¿ã¼ããã«

## ã¤ã³ã¹ãã¼ã«

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

### ãã¤ããªãã¡ã¤ã«

[Releases](https://github.com/arrow2nd/twnyan/releases)
ãããä½¿ãã®ç°å¢ã«ãã£ããã¡ã¤ã«ãããã³ã­ã¼ããã¦ãã ããã

## ã¢ã³ã¤ã³ã¹ãã¼ã«

### Homebrew

```sh
brew uninstall twnyan
rm -rf ~/.twnyan # è¨­å®ãã¡ã¤ã«ç¾¤ãåé¤
```

### Go

```sh
go clean -i github.com/arrow2nd/twnyan
rm -rf $GOPATH/src/github.com/arrow2nd/twnyan
rm -rf ~/.twnyan
```

## åæè¨­å®

![auth](https://user-images.githubusercontent.com/44780846/106747441-4a59dd00-6667-11eb-8248-3468cb39f7d1.png)

1. ååèµ·åæã«è¡¨ç¤ºãããèªè¨¼ãã¼ã¸ã® URL ã«ãã©ã¦ã¶ã§ã¢ã¯ã»ã¹
2. ç»é¢ã«å¾ã£ã¦æé ãé²ããè¡¨ç¤ºããã PIN ã³ã¼ãã twnyan ã«å¥å
3. å®äºï¼ ðº

> æåã«è¿½å ããã¢ã«ã¦ã³ãã¯ **ãã¡ã¤ã³ã¢ã«ã¦ã³ãã** ã¨ãã¦æ±ããã¾ã

## ä½¿ãæ¹

### ã³ãã³ãã©ã¤ã³ã¢ã¼ã

![cmdline](https://user-images.githubusercontent.com/44780846/106699170-b287cf00-6625-11eb-8374-8565286db3e2.gif)

- ä¸é¨ã®ã³ãã³ãã¯ä½¿ç¨ã§ãã¾ãã

### å¯¾è©±ã¢ã¼ã

ã³ãã³ããæå®ããªãå ´åãå¯¾è©±ã¢ã¼ãã§èµ·åãã¾ãã

### ã«ãã¼ã

![nyaan](https://user-images.githubusercontent.com/44780846/106699001-558c1900-6625-11eb-948e-6212ab0cba40.gif)

```
twnyan tw
```

ãã¤ã¼ãæãçç¥ããã¨ãã«ãã¼ããã¨ãã¤ã¼ãããã¾ãã

- ãªãã©ã¤ãå¼ç¨ãªãã¤ã¼ãã§ãåæ§ã®åä½ããã¾ã
- **ç»åãæ·»ä»ããã¦ããå ´åã¯ãã«ãã¼ãããã¾ãã**

### è©³ããä½¿ãæ¹

- [ã³ãã³ãä¸è¦§](./docs/ja/commands/index.md)
- [è¨­å®ãã¡ã¤ã«ã«ã¤ãã¦](./docs/ja/options.md)
