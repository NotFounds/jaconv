package jaconv

import (
	"strings"

	conversion "github.com/NotFounds/jaconv/lib"
)

// Romaji2Kana return string
func Romaji2Kana(text string) string {
	lower := strings.ToLower(text)
	return conversion.Romaji2kanaTable.Replace(lower)
}

// Kana2Romaji retune string
func Kana2Romaji(text string) string {
	katakana := Hiragana2Katakana(text)
	converted := []rune(conversion.Kana2RomajiTable.Replace(katakana))
	lastIdx := len(converted) - 1
	romaji := []rune{}
	for i, c := range converted {
		switch c {
		case 'ン':
			if i == lastIdx {
				romaji = append(romaji, 'n')
				continue
			}
			nextChara := converted[i+1]
			switch nextChara {
			case 'b', 'm', 'p':
				romaji = append(romaji, 'm')
			default:
				romaji = append(romaji, 'n')
			}

		case 'ッ':
			if i == lastIdx {
				romaji = append(romaji, []rune("ltu")...)
				continue
			}
			nextChara := converted[i+1]
			if nextChara == 'c' {
				romaji = append(romaji, 't')
			} else {
				romaji = append(romaji, converted[i+1])
			}

		case 'ー':
			// ignore

		default:
			if i > 0 && (c == 'u' || c == 'o') {
				prevChara := converted[i-1]
				switch c {
				case 'u':
					if prevChara == 'u' || prevChara == 'o' {
						continue
					}
				case 'o':
					if prevChara == 'o' && i != lastIdx {
						continue
					}
				}
			}
			romaji = append(romaji, c)
		}
	}
	return string(romaji)
}

// Hiragana2Katakana return string
func Hiragana2Katakana(text string) string {
	return conversion.Hiragana2Katakana.Replace(text)
}

// Katakana2Hiragana return string
func Katakana2Hiragana(text string) string {
	return conversion.Katakana2Hiragana.Replace(text)
}
