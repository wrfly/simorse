package morse

import (
	"fmt"
	"strings"
	"testing"
)

func TestMorseCode(t *testing.T) {
	str := " A b C d 123 "
	upperStr := strings.ToUpper(str)

	codes := MorseCode(str)
	if s := codes.String(); s != upperStr {
		t.Errorf("morse [%s] != [%s]", s, upperStr)
	}
	for _, x := range codes {
		fmt.Printf("%v\n", x)
	}
}

func TestSigCode(t *testing.T) {
	str := "qwer1dsava 923r4qfasdsaa09snv "

	sigCodes := SigMorseCode(str)
	if s := sigCodes.String(); s != str {
		t.Errorf("sig morse [%s] != [%s]", s, str)
	}
	sig := sigCodes.Signal()
	fmt.Printf("sig string: %s\n", sig)
	fmt.Printf("sig len: %d\n", len(sig))

	parsedCodes, err := ParseSigMorseCode(fmt.Sprint(sig))
	if err != nil {
		t.Error(err)
	}
	if parsedCodes.String() != str {
		t.Errorf("[%s] != [%s]", parsedCodes, str)
	}
}
