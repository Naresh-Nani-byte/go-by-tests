package maps

import (
	"testing"
)
func TestSearch(t *testing.T){
	dictionary := map[string]string{
		"B1618151":"Naresh", 
		"B1618150":"Mohan",
		"R161851": "Mahesh",
		"R210112": "Praveen",
	}

	t.Run("get the value using the key", func(t *testing.T) {
		got := GetSearchVaule(dictionary, "B1618151")
		want := "Naresh"
		assetStrings(t, got, want)
	})

	t.Run("Let's try if the key is in the dictionary or not ", func(t *testing.T) {
		dictionary := Dictionary{"R210112": "Praveen"}
		_, err := dictionary.SearchKeyInDictionary("R21011234")
		if err == nil {
			t.Fatalf("nee an error but got nothing ")
		}
		assetStrings(t, err.Error(), ErrNotFound.Error())
	})

	t.Run("Add New word to dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.AddToDictionary("B161021", "Shiva")
		want := "Shiva"
		defination, err := dictionary.SearchKeyInDictionary("B161021")
		if err != nil {
			t.Fatalf("Couldn't able to add the value to dictionary")
		}
		assetStrings(t, defination, want)
	})

}

func assetStrings(t testing.TB, got, want string){
	t.Helper()

	if got != want {
		t.Errorf("got thi %q but want %q", got, want)
	}
}