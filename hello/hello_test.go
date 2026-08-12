package main

import "testing"

func TestHello(t *testing.T){
	t.Run("saying hello to the people", func (t *testing.T)  {
		got := Hello("Chris","English")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)

	})

	t.Run("Say hello to the world when we pass empty string", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Giving the Language with string", func (t *testing.T)  {
		got := Hello("Elodie","Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say it in the Telugu", func(t *testing.T) {
		got := Hello("Naresh", "Telugu")
		want := "Namaste, Naresh"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say it in the Malayalam ", func(t *testing.T) {
		got :=  Hello("Ramya", "Malayalam")
		want := "Namaste, Ramya"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string){
	t.Helper()
	if got != want{
		t.Errorf("got %q want %q", got, want)
	}
}
