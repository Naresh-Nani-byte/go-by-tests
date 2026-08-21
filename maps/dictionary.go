package maps

import "errors"


type dictionary map[string]string
type Dictionary map[string]string
var ErrNotFound = errors.New("Could not find the word you were looking for")


func GetSearchVaule(d dictionary, key string) string { 
	return d[key]
}

func (d Dictionary) SearchKeyInDictionary(key string) (string, error) {
	defination, ok := d[key]
	if !ok{
		return "", ErrNotFound
	}
	return defination, nil
}

func (d Dictionary) AddToDictionary(word, defination string){
	d[word] = defination
}