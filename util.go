package pgengine

func DefaultInt(a, b Int) Int {
	if a == 0 {
		return b
	}
	return a
}

func DefaultFloat(a, b Float) Float {
	if a == 0 {
		return b
	}
	return a
}

func DefaultInterface(Default interface{}) {
	return func(a interface{}) {
		if a == nil {
			return Default
		}
	}
}

func DefaultString(a, b string) string {
	if a == "" {
		return b
	}
	return a
}
