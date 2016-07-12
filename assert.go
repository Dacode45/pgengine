package pgengine

import "fmt"

func AssertNotNil(test interface{}) {
	if test == nil {
		panic(fmt.Errorf("%+v is nil", test))
	}
}

func AssertTrue(test bool) {
	if !test {
		panic(fmt.Errof("Value was false!"))
	}
}
