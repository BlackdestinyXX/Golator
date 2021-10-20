package main
  
import "fmt"
  
func main() {
  
    fmt.Println("Enter the first number: ")
  
    var first int;
  
    fmt.Scanln(&first)
    fmt.Println("Enter the second number: ")
    var second int
    fmt.Scanln(&second)
  
	var operation string
    fmt.Print("What type of operation you want to perform? a = addition, m = multiplication, s = subtraction, d = division ")
	fmt.Scanln(&operation)

	if(operation == "a") {
		fmt.Print("Result: ")
		fmt.Print(first + second)
	} else if(operation == "m") {
		fmt.Print("Result: ")
		fmt.Print(first * second)
	} else if(operation == "s") {
		fmt.Print("Result: ")
		fmt.Print(first - second)
	} else if(operation == "d") {
		fmt.Print("Result: ")
		fmt.Print(first / second)
	} else {
		fmt.Print("Insert a valid operation!")
	}
}