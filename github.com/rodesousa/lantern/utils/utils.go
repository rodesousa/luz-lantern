package utils

func GetBool(toAnalyse interface{}, defaultValue bool) bool {
	if toAnalyse == nil {
		return defaultValue
	}
	return toAnalyse.(bool)
}


func ByteToString(c []byte) string {
	n := -1
	for i, b := range c {
		if b == 0 {
			break
		}
		n = i
	}
	return string(c[:n+1])
}