package handle

func HandleString(str string) string {
	str = "'" + str + "'"
	return str
}
func InterfaceToString(i interface{}) string {
	res, _ := i.([]byte)

	return string(res)
}
