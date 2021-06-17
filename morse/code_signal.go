package morse

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
)

type SigCode uint8

func (c SigCode) String() string {
	return string(signalCodes[c])
}

type SigCodes []SigCode

// String returns the raw signal string
func (c SigCodes) String() (x string) {
	for _, code := range c {
		x = fmt.Sprintf("%s%s", x, code)
	}
	return
}

// Signal converts the SigCodes into signals, 01010011101
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

// Signal returns the syscall.Signal, USR1 and USR2
func (s Signal) Signal() syscall.Signal {
	if s {
		return syscall.SIGUSR1
	}
	return syscall.SIGUSR2
}

// ParseSignal can convert os.Signal to morse.Signal
func ParseSignal(s os.Signal) Signal {
	if s == syscall.SIGUSR1 {
		return true
	}
	return false
}

type Signals []Signal

func (s Signals) String() (x string) {
	for _, xx := range s {
		x = fmt.Sprintf("%s%s", x, xx)
	}
	return
}

// signal code
const signalCodes = "qwertyuiopasdfghjklzxcvbnm" +
	"QWERTYUIOPASDFGHJKLZXCVBNM1234567890 _"

// SigMorseCode convert string to SigCodes
func SigMorseCode(input string) SigCodes {
	codes := make(SigCodes, 0, len(input))

	for _, r := range input {
		index := strings.IndexRune(signalCodes, r)
		if index != -1 {
			codes = append(codes, SigCode(index))
		}
	}

	return codes
}

// ParseSigMorseCode convert 010101 to SigCodes
func ParseSigMorseCode(input string) (SigCodes, error) {
	if len(input)%6 != 0 {
		return nil, fmt.Errorf("bad input signal length %d", len(input))
	}
	sigCodes := SigCodes{}
	for i := 0; i < len(input); i += 6 {
		binaryStr := fmt.Sprintf("%s", input[i:i+6])
		u, err := strconv.ParseUint(binaryStr, 2, 8)
		if err != nil {
			return nil, err
		}
		sigCodes = append(sigCodes, SigCode(uint8(u)))
	}
	return sigCodes, nil
}
