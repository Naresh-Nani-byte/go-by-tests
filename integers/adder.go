package main
import "fmt"

const(
	sum = "SUM"
	substraction = "SUB"
	multiplication = "MULTI"
)

func mathExpression(num1, num2 int, Exp string) (value int){

	switch Exp{
	case sum:
		value=num1+num2
	case substraction:
		value=num1-num2
	case multiplication:
		value=num1*num2 
	} 
	return value
}
 

func main(){
	fmt.Println(mathExpression(0,0,"SUM"))
}