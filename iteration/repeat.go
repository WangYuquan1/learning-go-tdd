package iteration

func Repeat(character string, repeatCount int) string {
	var repeatd string
	for i := 0; i < repeatCount; i++ {
		repeatd += character
	}
	return repeatd
}
