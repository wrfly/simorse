package morse

import (
	"strings"
	"testing"
)

func TestCode(t *testing.T) {
	str := " hello world 000"
	upperStr := strings.ToUpper(str)

	codes := MorseCode(str)
	if s := codes.String(); s != upperStr {
		t.Errorf("morse [%s] != [%s]", s, upperStr)
	}

	sigCodes := SigMorseCode(str)
	if s := sigCodes.String(); s != upperStr {
		t.Errorf("sig morse [%s] != [%s]", s, upperStr)
	}
}
