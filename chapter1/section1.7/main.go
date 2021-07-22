package main

import (
	"fmt"
)

func formletter(name string, weather string, snack string){
	fmt.Printf("Hello %s!\nThe weather today is %s.\nToday's snack will be %s.\n", name, weather, snack)
}

func main(){
	formletter(Name, Weather, Snack)
}
