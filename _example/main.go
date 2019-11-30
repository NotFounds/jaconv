package main

import (
	"fmt"

	"github.com/NotFounds/jaconv"
)

func main() {
	text := "こんにちは"
	romaji := jaconv.Kana2Romaji(text)
	fmt.Println(romaji) // konnichiha

	katakana := jaconv.Romaji2Kana(romaji)
	fmt.Println(katakana) // コンニチハ

	hiragana := jaconv.Katakana2Hiragana(katakana)
	fmt.Println(hiragana) // こんにちは
}
