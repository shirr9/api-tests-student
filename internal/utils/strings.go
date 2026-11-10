package utils

import "unicode"

// MixCase чередует верхний и нижний регистры букв.
func MixCase(text string) string {
	runes := []rune(text)
	useUpper := true

	for i, r := range runes {
		lower := unicode.ToLower(r)
		upper := unicode.ToUpper(r)
		if lower == upper {
			continue
		}

		if useUpper {
			runes[i] = upper
		} else {
			runes[i] = lower
		}
		useUpper = !useUpper
	}

	return string(runes)
}
