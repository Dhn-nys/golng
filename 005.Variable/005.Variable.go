package main

import "fmt"

func main ()  {

	// first tutorial
	// use var and name variable and type data. 
	var name string

	name = "Eko Kurniawan"
	fmt.Println(name)

	name = "Eko Khannedy"
	fmt.Println(name)

	// second tutorial
	// not use the type data

	var namefriend = "Budi"
	fmt.Println(namefriend)

	var age = 30
	fmt.Println(age)

	// you can edit the type data int32 to int8 but you must typing the formation variable

	var age2 int8 = 31
	fmt.Println(age2)

	// third tutorial
	// no used var you can use ":=" to change var

	country := "Indonesia"
	fmt.Println(country)

	// fourth tutorial
	// to clean variable you must use the code

	var (
		firstname = "Eko"
		lastname = "Kurniawan Khannedy"
	)

	fmt.Println(firstname)
	fmt.Println(lastname)
	
}
