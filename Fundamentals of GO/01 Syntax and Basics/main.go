package main

import "fmt"

// MAIN FUNCTION-------------------------------------

func main(){

	variables()

	
}


// FUNCTION FOR CONSTANT---------------------------------------------

func constant(){
	
}


//FUNCTION FOR VARIABLES ------------------------------------

func variables(){

	// var keyword is used to declare variables
	var a int  
	a = 11
	fmt.Println(a)


	// without writing data type of variable we can define it
	var v = 17
	fmt.Println(v)


	// using ':=' keyword we can declare and assign value to variable
	z := 9
	fmt.Println(z)

	
	// We can also define multiple variables in same line
	var1, var2 := 1, "Akshat"
	fmt.Println(var1,var2)

	var v1, v2 = 88, "hello"
	fmt.Println(v1,v2)

 }
