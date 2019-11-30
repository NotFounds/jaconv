# jaconv

![workflow status](https://github.com/NotFounds/jaconv/workflows/Go/badge.svg)

## Description

A converter that Japanese Hiragana/Katakana and Romaji.

## Usage

```golang
text = "こんにちは"
romaji := jaconv.Kana2Romaji(text)
fmt.Println(romaji) // konnnichiha
katakana := jaconv.Romaji2Kana(romaji)
fmt.Println(katakana) // コンニチハ
hiragana := jaconv.Katakana2Hiragana(katakana)
fmt.Println(hiragana) // こんにちは
```

## Installation

```sh
go get github.com/NotFounds/jaconv
```

## References

[ヘボン式ローマ字綴方表 - 外務省](https://www.ezairyu.mofa.go.jp/passport/hebon.html)

## License

MIT
