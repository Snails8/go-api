package util

func BoolToInt8(value bool) int8 {
	num := int8(0)
	if value {
		num = int8(1)
		return num
	}	
	return num
}

func BoolToText(value bool, trueText string, falseText string) string {
	text := falseText
	if value {
		text = trueText
		return text
	}	
	return text
}

