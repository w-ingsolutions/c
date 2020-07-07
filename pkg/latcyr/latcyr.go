package latcyr

import "fmt"

// Mapping of Latin characters to Cyrillic.
var LatinToCyrillic = map[string]string{
	"a":  "а",
	"A":  "А",
	"b":  "б",
	"B":  "Б",
	"v":  "в",
	"V":  "В",
	"g":  "г",
	"G":  "Г",
	"d":  "д",
	"D":  "Д",
	"đ":  "ђ",
	"Đ":  "Ђ",
	"e":  "е",
	"E":  "Е",
	"ž":  "ж",
	"Ž":  "Ж",
	"z":  "з",
	"Z":  "З",
	"i":  "и",
	"I":  "И",
	"j":  "ј",
	"J":  "Ј",
	"k":  "к",
	"K":  "К",
	"l":  "л",
	"L":  "Л",
	"lj": "љ",
	"Lj": "Љ",
	"LJ": "Љ",
	"m":  "м",
	"M":  "М",
	"n":  "н",
	"N":  "Н",
	"nj": "њ",
	"Nj": "Њ",
	"NJ": "Њ",
	"o":  "о",
	"O":  "О",
	"p":  "п",
	"P":  "П",
	"r":  "р",
	"R":  "Р",
	"s":  "с",
	"S":  "С",
	"t":  "т",
	"T":  "Т",
	"ć":  "ћ",
	"Ć":  "Ћ",
	"u":  "у",
	"U":  "У",
	"f":  "ф",
	"F":  "Ф",
	"h":  "х",
	"H":  "Х",
	"c":  "ц",
	"C":  "Ц",
	"č":  "ч",
	"Č":  "Ч",
	"dž": "џ",
	"Dž": "Џ",
	"DŽ": "Џ",
	"š":  "ш",
	"Š":  "Ш",
}

func C(s string, enable bool) (result string) {
	if enable {
		var caracter string
		for i, r := range s {
			caracter = string(r)
			if i != len(s)-1 {
				if c, exist := LatinToCyrillic[string(r)+string(s[i+1])]; exist {
					caracter = string(r) + string(s[i+1])
					fmt.Println("i", i)
					//r = last + r
					fmt.Println("last", c)
					fmt.Println("last + rr", string(r))
				}
			}

			if cyrillic, exist := LatinToCyrillic[caracter]; exist {
				fmt.Println("caracter", caracter)

				result += cyrillic
			} else {
				result += string(r)
			}
		}
	} else {
		result = s
	}
	return
}
