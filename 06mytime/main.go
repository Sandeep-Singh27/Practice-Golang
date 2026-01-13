package main

import "fmt"
import "time"

func main(){
	fmt.Println("Welcome to time study in golang")

	var presentTime time.Time = time.Now()
	fmt.Printf("%T\n",presentTime)
	fmt.Println(presentTime.Format("02-01-2006 Monday 15:04:05"))

	createdDate := time.Date(2020,time.September, 32, 19, 56, 32,00,time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("02/Jan/2006 Monday"))
}