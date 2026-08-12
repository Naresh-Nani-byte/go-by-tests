package main

import (
	"fmt"
)
const(
	spanish = "Spanish"
	malayalam  = "Malayalam"
	telugu = "Telugu"
	english = "English"
	spanishPrifix = "Hola, "
	teluguPrifix = "Namaste, "
	malayalamPrifix = "Namaste, "
	englishPrefix = "Hello, "
)

func Hello(name, lang string) string{
	if name == ""{
		name = "World"
	}
	return greetingPrifix(lang) + name
}

func greetingPrifix(language string) (prifix string){
	switch language{
	case spanish:
		prifix = spanishPrifix
	case telugu:
		prifix = teluguPrifix
	case malayalam:
		prifix = malayalamPrifix
	case english:
		prifix = englishPrefix
	}
	return
}

func main (){
	fmt.Println(Hello("",""))
}