package main

import "testing"

func TestAdd(t *testing.T){
	t.Run("Add two number and get the result", func(t *testing.T) {
		got := mathExpression(2,2,"MULTI")
		want := 4
		asserMessage(t, got, want)
	})
	t.Run("send 0 and 0 let's see what our response", func(t *testing.T) {
		got := mathExpression(0,0,"SUM")
		want := 0
		asserMessage(t, got, want)
	})

	t.Run("let's give the math expression to do", func(t *testing.T) {
		got := mathExpression(6,5,"SUB")
		want := 1
		asserMessage(t,got, want)
	})
}

func asserMessage(t testing.TB, got, want int){
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}

}