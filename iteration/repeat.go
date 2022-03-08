package iteration

const repeatCount = 5

func Repeat(character string) string {
	var repeatd string
	for i := 0; i < repeatCount; i++ {
		repeatd += character
	}
	return repeatd
}
