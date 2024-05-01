package tools

func StrLimit(str string, length int) string {
	if len(str) < length {
		return str
	}

	return str[:length]
}
