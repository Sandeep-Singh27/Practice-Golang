package main
import "fmt"

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
}