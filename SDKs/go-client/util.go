package swagger


func clearEmptyParams(paramMap map[string][]string) {

	for key, value := range paramMap {
		if (len(value) == 1 && value[0] == "") {
			delete(paramMap, key)
		}
	}
}
