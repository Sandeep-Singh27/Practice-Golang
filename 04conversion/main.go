package main

import "fmt"
import "bufio"
import "os"
import "strconv"
import "strings"

func main(){
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 and 5")

	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Printf("Thanks of the rating %s", input)

	//suppose for some reason you want to add +1 to the given rating
	//we cannot directly add that because input is of type string

	addedRating, err := strconv.ParseFloat(strings.TrimSpace(input),64)
	
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("After adding +1 to it ",addedRating+1)
	}

}