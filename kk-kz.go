package ntw

import (
	"strings"
)

func init() {
	// register the Kazakh language
	Languages["kk-kz"] = Language{
		Name:    "Kazakh",
		Aliases: []string{"kk", "kz", "kk_KZ", "kazakh"},
		Flag:    "🇰🇿",

		IntegerToWords: IntegerToKkKz,
	}
}

func pluralKk(n int, words []string) string {
	// В казахском склонения как в русском нет, но структура оставлена для совместимости
	return words[0]
}

func IntegerToKkKz(input int) string {
	var kkUnits = []string{
		"",
		"бір",
		"екі",
		"үш",
		"төрт",
		"бес",
		"алты",
		"жеті",
		"сегіз",
		"тоғыз",
	}
	var kkTeens = []string{
		"он",
		"он бір",
		"он екі",
		"он үш",
		"он төрт",
		"он бес",
		"он алты",
		"он жеті",
		"он сегіз",
		"он тоғыз",
	}
	var kkTens = []string{
		"",
		"он",
		"жиырма",
		"отыз",
		"қырық",
		"елу",
		"алпыс",
		"жетпіс",
		"сексен",
		"тоқсан",
	}
	var kkHundreds = []string{
		"",
		"жүз",
		"екі жүз",
		"үш жүз",
		"төрт жүз",
		"бес жүз",
		"алты жүз",
		"жеті жүз",
		"сегіз жүз",
		"тоғыз жүз",
	}
	var kkMegas = [][]string{
		{"", "", ""},
		{"мың", "мың", "мың"},
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
			words = append(words, kkHundreds[hundreds])
		}

		if tens > 0 || units > 0 {
			switch tens {
			case 0:
				words = append(words, kkUnits[units])
			case 1:
				words = append(words, kkTeens[units])
			default:
				words = append(words, kkTens[tens])
				if units > 0 {
					words = append(words, kkUnits[units])
				}
			}
		}

		if idx >= 1 && idx < len(kkMegas) {
			mega := kkMegas[idx]
			words = append(words, pluralKk(tens*10+units, mega))
		}
	}

	return strings.Join(words, " ")
}
