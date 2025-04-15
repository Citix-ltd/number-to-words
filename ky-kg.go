package ntw

import (
	"strings"
)

func init() {
	// register the Kyrgyz language
	Languages["ky-kg"] = Language{
		Name:    "Kyrgyz",
		Aliases: []string{"ky", "kg", "ky_KG", "kyrgyz"},
		Flag:    "🇰🇬",

		IntegerToWords: IntegerToKyKg,
	}
}

func pluralKy(n int, words []string) string {
	n %= 100
	if n > 19 {
		n %= 10
	}
	if n == 1 {
		return words[0]
	} else if n >= 2 && n <= 4 {
		return words[1]
	}
	return words[2]
}

func IntegerToKyKg(input int) string {
	var kyUnits = []string{
		"",
		"бир",
		"эки",
		"үч",
		"төрт",
		"беш",
		"алты",
		"жети",
		"сегиз",
		"тогуз",
	}
	var kyTeens = []string{
		"он",
		"он бир",
		"он эки",
		"он үч",
		"он төрт",
		"он беш",
		"он алты",
		"он жети",
		"он сегиз",
		"он тогуз",
	}
	var kyTens = []string{
		"",
		"он",
		"жыйырма",
		"отуз",
		"кырк",
		"элүү",
		"алтымыш",
		"жетимиш",
		"сексен",
		"токсон",
	}
	var kyHundreds = []string{
		"",
		"жүз",
		"эки жүз",
		"үч жүз",
		"төрт жүз",
		"беш жүз",
		"алты жүз",
		"жети жүз",
		"сегиз жүз",
		"тогуз жүз",
	}
	var kyMegas = [][]string{
		{"", "", ""},
		{"миң", "миң", "миң"},
		{"миллион", "миллион", "миллион"},
		{"миллиард", "миллиард", "миллиард"},
		{"триллион", "триллион", "триллион"},
		{"квадриллион", "квадриллион", "квадриллион"},
		{"квинтиллион", "квинтиллион", "квинтиллион"},
	}

	var words []string

	if input < 0 {
		words = append(words, "минус")
		input *= -1
	}

	triplets := integerToTriplets(input)

	if len(triplets) == 0 {
		return "нөл"
	}

	for idx := len(triplets) - 1; idx >= 0; idx-- {
		triplet := triplets[idx]

		if triplet == 0 {
			continue
		}

		hundreds := triplet / 100 % 10
		tens := triplet / 10 % 10
		units := triplet % 10

		if hundreds > 0 {
			words = append(words, kyHundreds[hundreds])
		}

		if tens > 0 || units > 0 {
			switch tens {
			case 0:
				words = append(words, kyUnits[units])
			case 1:
				words = append(words, kyTeens[units])
			default:
				words = append(words, kyTens[tens])
				if units > 0 {
					words = append(words, kyUnits[units])
				}
			}
		}

		if idx >= 1 && idx < len(kyMegas) {
			mega := kyMegas[idx]
			tensTotal := tens*10 + units
			words = append(words, pluralKy(tensTotal, mega))
		}
	}

	return strings.Join(words, " ")
}
