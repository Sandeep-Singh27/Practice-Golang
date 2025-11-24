package main
import "fmt"

const LogingToken string = "ajfiojfwoj" //This is actually public since start with uppercase letter in variable name

func main(){
	var username string = "sandeep";
	fmt.Println(username)
	fmt.Printf("variable is of type : %T\n\n",username)

	var isPrime bool = false
	fmt.Println(isPrime)
	fmt.Printf("variable is of type : %T\n\n",isPrime)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("variable is of type : %T\n\n", smallValue)

	var smallFloat float32 = 255.4555553422424242
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T\n\n", smallFloat)

	var bigFloat float64 = 255.4555553422424242
	fmt.Println(bigFloat)
	fmt.Printf("variable is of type: %T\n\n", bigFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T\n\n", anotherVariable)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("variable is of type: %T\n\n", anotherString)

	//implicit type
	var website = "google.com"
	fmt.Println(website)
	fmt.Printf("variable is of type: %T\n\n", website)

	//no var style
	numberOfUSer := 3000.0
	fmt.Println(numberOfUSer)
	fmt.Printf("variable is of type: %T\n\n", numberOfUSer)

	fmt.Println(LogingToken)
	fmt.Printf("variable is of type: %T\n\n", LogingToken)

}