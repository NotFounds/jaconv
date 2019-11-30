package jaconv

import "testing"

func kana2romajiHelper(t *testing.T, text string, expected string) {
	romaji := Kana2Romaji(text)
	if romaji != expected {
		t.Errorf("\ngot:      %v\nexpected: %v", romaji, expected)
	}
}

func romaji2kanaHelper(t *testing.T, text string, expected string) {
	kana := Romaji2Kana(text)
	if kana != expected {
		t.Errorf("\ngot:      %v\nexpected: %v", kana, expected)
	}
}

func TestKana2romaji(t *testing.T) {
	kana2romajiHelper(t, "アンパンとどらやきのどちらがすきですか？", "ampantodorayakinodochiragasukidesuka？")
	kana2romajiHelper(t, "コンニチハ イイ tenki デスネ", "konnichiha ii tenki desune")
	kana2romajiHelper(t, "はっとり", "hattori")
	kana2romajiHelper(t, "ほんま", "homma")
	kana2romajiHelper(t, "ヒュウガ", "hyuga")
	kana2romajiHelper(t, "こうた", "kota")
	kana2romajiHelper(t, "オオノ", "ono")
	kana2romajiHelper(t, "たかとお", "takatoo")
	kana2romajiHelper(t, "イトウ", "ito")
}

func TestRomaji2kana(t *testing.T) {
	romaji2kanaHelper(t, "Konnichiha ii tenki desune", "コンニチハ イイ テンキ デスネ")
}

func BenchmarkKana2Romaji(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Kana2Romaji("アンパンとどらやきのどちらがすきですか？")
	}
}

func BenchmarkRomaji2Kana(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Romaji2Kana("Konnichiha ii tenki desune")
	}
}

func BenchmarkHiragana2Katakana(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Hiragana2Katakana("ひらがなをカタカナにへんかんします")
	}
}

func BenchmarkKatakana2Hiragana(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Katakana2Hiragana("ゼンカクヲハンカクニヘンカンシマス")
	}
}
