package morse

import "strings"

type Code string

type Codes []Code

func (c Codes) String() string {
	s := make([]rune, 0, 100)
	for _, code := range c {
		if code != morseShort {
			s = append(s, rMorseMap[code])
		}
	}
	return string(s)
}

// morse code
const (
	dot  = "·"
	dash = "-"

	morseA Code = "·-"
	morseB Code = "-···"
	morseC Code = "-·-·"
	morseD Code = "-··"
	morseE Code = "·"
	morseF Code = "··-·"
	morseG Code = "--·"
	morseH Code = "····"
	morseI Code = "··"
	morseJ Code = "·---"
	morseK Code = "-·-"
	morseL Code = "·-··"
	morseM Code = "--"
	morseN Code = "-·"
	morseO Code = "---"
	morseP Code = "·--·"
	morseQ Code = "--·-"
	morseR Code = "·-·"
	morseS Code = "···"
	morseT Code = "-"
	morseU Code = "··-"
	morseV Code = "···-"
	morseW Code = "·--"
	morseX Code = "-··-"
	morseY Code = "-·--"
	morseZ Code = "--··"

	morseShort Code = "000"
	morseSpace Code = "0000000"

	morseNum0 Code = "-----"
	morseNum1 Code = "·----"
	morseNum2 Code = "··---"
	morseNum3 Code = "···--"
	morseNum4 Code = "····-"
	morseNum5 Code = "·····"
	morseNum6 Code = "-····"
	morseNum7 Code = "--···"
	morseNum8 Code = "---··"
	morseNum9 Code = "----·"
)

var morseMap = map[rune]Code{
	'A': morseA, 'B': morseB, 'C': morseC, 'D': morseD,
	'E': morseE, 'F': morseF, 'G': morseG, 'H': morseH,
	'I': morseI, 'J': morseJ, 'K': morseK, 'L': morseL,
	'M': morseM, 'N': morseN, 'O': morseO, 'P': morseP,
	'Q': morseQ, 'R': morseR, 'S': morseS, 'T': morseT,
	'U': morseU, 'V': morseV, 'W': morseW, 'X': morseX,
	'Y': morseY, 'Z': morseZ,

	'0': morseNum0, '1': morseNum1, '2': morseNum2,
	'3': morseNum3, '4': morseNum4, '5': morseNum5,
	'6': morseNum6, '7': morseNum7,
	'8': morseNum8, '9': morseNum9,

	' ': morseSpace,
}

var rMorseMap = map[Code]rune{
	morseA: 'A', morseB: 'B', morseC: 'C', morseD: 'D',
	morseE: 'E', morseF: 'F', morseG: 'G', morseH: 'H',
	morseI: 'I', morseJ: 'J', morseK: 'K', morseL: 'L',
	morseM: 'M', morseN: 'N', morseO: 'O', morseP: 'P',
	morseQ: 'Q', morseR: 'R', morseS: 'S', morseT: 'T',
	morseU: 'U', morseV: 'V', morseW: 'W', morseX: 'X',
	morseY: 'Y', morseZ: 'Z',

	morseNum0: '0', morseNum1: '1', morseNum2: '2',
	morseNum3: '3', morseNum4: '4', morseNum5: '5',
	morseNum6: '6', morseNum7: '7',
	morseNum8: '8', morseNum9: '9',

	morseSpace: ' ',
}

func MorseCode(input string) Codes {
	codes := make([]Code, 0, len(input)*2)

	for _, s := range strings.ToUpper(input) {
		if s == ' ' && len(codes) > 0 {
			codes = codes[:len(codes)-1]
		}
		codes = append(codes, morseMap[s])
		if s != ' ' {
			codes = append(codes, morseShort)
		}
	}

	return codes
}
