package maps

import "testing"

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
		got := SearchKeyInDictionary(dictionary, "R210112")
		want := true
		assertBooleanValue(t, got, want)
	})

}

func assertBooleanValue(t testing.TB, got, want bool){
	t.Helper()
	if got != want{
		t.Errorf("got %v but want %v", got, want)
	}
}

func assetStrings(t testing.TB, got, want string){
	t.Helper()

	if got != want {
		t.Errorf("got thi %s but want %s", got, want)
	}
}