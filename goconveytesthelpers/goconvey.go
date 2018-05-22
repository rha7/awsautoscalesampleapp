package goconveytesthelpers

import (
	"fmt"
	"regexp"
)

func ShouldMatchString(actual interface{}, expected ...interface{}) string {
	strActual, ok := actual.(string)
	if !ok {
		return fmt.Sprintf("invalid actual (not a string) value [%#v]", actual)
	}
	if len(expected) != 1 {
		return fmt.Sprintf("missing or more than one expected (regular expression) value [%#v]", expected)
	}
	strExpected, ok := expected[0].(string)
	if !ok {
		return fmt.Sprintf("invalid expected (not a string) value [%#v]", expected[0])
	}

	rx, err := regexp.Compile(strExpected)
	if err != nil {
		return fmt.Sprintf("invalid expected (invalid regular expression) value [%#v]", strExpected)
	}

	if rx.MatchString(strActual) {
		return "" // we good
	} else {
		return fmt.Sprintf("actual (%#v) does not match expression (%#v)", strActual, strExpected)
	}
}
