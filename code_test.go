package morse

import (
	"fmt"
	"strings"
	"testing"
)

func TestCode(t *testing.T) {
	str := " A b C d 123 "
	upperStr := strings.ToUpper(str)

	codes := MorseCode(str)
	t.Logf("morse len: %d", len(codes))
	if s := codes.String(); s != upperStr {
		t.Errorf("morse [%s] != [%s]", s, upperStr)
	}

	sigCodes := SigMorseCode(str)
	t.Logf("sig len: %d", len(sigCodes))
	if s := sigCodes.String(); s != str {
		t.Errorf("sig morse [%s] != [%s]", s, str)
	}
	sig := sigCodes.Signal()
	t.Logf("sig string: %s", sig)
	t.Logf("sig len: %d", len(sig))

	parsedCodes, err := ParseSigMorseCode(fmt.Sprint(sig))
	if err != nil {
		t.Error(err)
	}
	if parsedCodes.String() != str {
		t.Errorf("[%s] != [%s]", parsedCodes, str)
	}
}
