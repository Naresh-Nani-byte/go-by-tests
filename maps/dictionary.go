package maps

type dictionary map[string]string



func GetSearchVaule(d dictionary, key string) string { 
	return d[key]
}

func SearchKeyInDictionary(d dictionary, key string) bool {
	var ok bool
	_, ok = d[key]
	return  ok
}