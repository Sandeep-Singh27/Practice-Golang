package main

import "fmt"
import "os"
import "bufio"

func main(){
	welcome := "Welcome to User Input Program"
	fmt.Println(welcome)

//	reader := bufio.NewReader(os.Stdin) returns *bufio.Reader type object
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	fmt.Printf("Reader is of type %T\n",reader)
	fmt.Println("Enter your name :")

//	comma ok || comma err
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of input is %T \n ", input)
}