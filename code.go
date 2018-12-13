package morse

import (
	"fmt"
	"strconv"
	"strings"
)

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

type SigCode uint8

func (c SigCode) String() string {
	return string(rMorseSigMap[c])
}

type SigCodes []SigCode

func (c SigCodes) String() (x string) {
	for _, code := range c {
		x = fmt.Sprintf("%s%s", x, code)
	}
	return
}

func (c SigCodes) Signal() Signals {
	list := []Signal{}
	for _, r := range c {
		// 2^6=64 > 26*2+10
		for _, b := range fmt.Sprintf("%0.6b", r) {
			list = append(list, b == '1')
		}
	}
	return list
}

type Signal bool

func (s Signal) String() string {
	if s {
		return "1"
	}
	return "0"
}

type Signals []Signal

func (s Signals) String() (x string) {
	for _, xx := range s {
		x = fmt.Sprintf("%s%s", x, xx)
	}
	return
}

// signal code
const (
	sigSpace SigCode = iota

	// upper case
	sigA
	sigB
	sigC
	sigD
	sigE
	sigF
	sigG
	sigH
	sigI
	sigJ
	sigK
	sigL
	sigM
	sigN
	sigO
	sigP
	sigQ
	sigR
	sigS
	sigT
	sigU
	sigV
	sigW
	sigX
	sigY
	sigZ

	// lower case
	sigA_
	sigB_
	sigC_
	sigD_
	sigE_
	sigF_
	sigG_
	sigH_
	sigI_
	sigJ_
	sigK_
	sigL_
	sigM_
	sigN_
	sigO_
	sigP_
	sigQ_
	sigR_
	sigS_
	sigT_
	sigU_
	sigV_
	sigW_
	sigX_
	sigY_
	sigZ_

	sigNum0
	sigNum1
	sigNum2
	sigNum3
	sigNum4
	sigNum5
	sigNum6
	sigNum7
	sigNum8
	sigNum9
)

var morseSigMap = map[rune]SigCode{
	'A': sigA, 'B': sigB, 'C': sigC, 'D': sigD,
	'E': sigE, 'F': sigF, 'G': sigG, 'H': sigH,
	'I': sigI, 'J': sigJ, 'K': sigK, 'L': sigL,
	'M': sigM, 'N': sigN, 'O': sigO, 'P': sigP,
	'Q': sigQ, 'R': sigR, 'S': sigS, 'T': sigT,
	'U': sigU, 'V': sigV, 'W': sigW, 'X': sigX,
	'Y': sigY, 'Z': sigZ,

	'a': sigA_, 'b': sigB_, 'c': sigC_, 'd': sigD_,
	'e': sigE_, 'f': sigF_, 'g': sigG_, 'h': sigH_,
	'i': sigI_, 'j': sigJ_, 'k': sigK_, 'l': sigL_,
	'm': sigM_, 'n': sigN_, 'o': sigO_, 'p': sigP_,
	'q': sigQ_, 'r': sigR_, 's': sigS_, 't': sigT_,
	'u': sigU_, 'v': sigV_, 'w': sigW_, 'x': sigX_,
	'y': sigY_, 'z': sigZ_,

	'0': sigNum0, '1': sigNum1, '2': sigNum2,
	'3': sigNum3, '4': sigNum4, '5': sigNum5,
	'6': sigNum6, '7': sigNum7,
	'8': sigNum8, '9': sigNum9,

	' ': sigSpace,
}

var rMorseSigMap = map[SigCode]rune{
	sigA: 'A', sigB: 'B', sigC: 'C', sigD: 'D',
	sigE: 'E', sigF: 'F', sigG: 'G', sigH: 'H',
	sigI: 'I', sigJ: 'J', sigK: 'K', sigL: 'L',
	sigM: 'M', sigN: 'N', sigO: 'O', sigP: 'P',
	sigQ: 'Q', sigR: 'R', sigS: 'S', sigT: 'T',
	sigU: 'U', sigV: 'V', sigW: 'W', sigX: 'X',
	sigY: 'Y', sigZ: 'Z',

	sigA_: 'a', sigB_: 'b', sigC_: 'c', sigD_: 'd',
	sigE_: 'e', sigF_: 'f', sigG_: 'g', sigH_: 'h',
	sigI_: 'i', sigJ_: 'j', sigK_: 'k', sigL_: 'l',
	sigM_: 'm', sigN_: 'n', sigO_: 'o', sigP_: 'p',
	sigQ_: 'q', sigR_: 'r', sigS_: 's', sigT_: 't',
	sigU_: 'u', sigV_: 'v', sigW_: 'w', sigX_: 'x',
	sigY_: 'y', sigZ_: 'z',

	sigNum0: '0', sigNum1: '1', sigNum2: '2',
	sigNum3: '3', sigNum4: '4', sigNum5: '5',
	sigNum6: '6', sigNum7: '7',
	sigNum8: '8', sigNum9: '9',

	sigSpace: ' ',
}

// SigMorseCode convert string to sigCodes
func SigMorseCode(input string) SigCodes {
	codes := make([]SigCode, 0, len(input)*2)

	for _, s := range input {
		codes = append(codes, morseSigMap[s])
	}

	return codes
}

// ParseSigMorseCode convert 010101 to string
func ParseSigMorseCode(input string) (SigCodes, error) {
	if len(input)%6 != 0 {
		return nil, fmt.Errorf("bad input signal")
	}
	sigCodes := SigCodes{}
	for i := 0; i < len(input); i += 6 {
		binaryStr := fmt.Sprintf("%s", input[i:i+6])
		u, err := strconv.ParseUint(binaryStr, 2, 8)
		if err != nil {
			return nil, err
		}
		sigCodes = append(sigCodes, SigCode(u))
	}
	return sigCodes, nil
}
