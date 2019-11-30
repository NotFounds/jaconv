# jaconv

## Description

A converter that Japanese Hiragana/Katakana and Romaji.

## Usage

```golang
text = "こんにちは"
roman, err := jaconv.Kana2roman(text)
if err != nil {
    log.Fatal(err)
}
fmt.Println(roman) // konnnitiwa
```

## Installation

```sh
go get github.com/NotFounds/jaconv
```

## License

MIT
