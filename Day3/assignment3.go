package main

import (
	"fmt"
	"net/http"
)

func main() {

	//slice
	studentNames := []string{"Mayur", "Nitesh", "Rajesh", "Karan", "Sumit", "Ravi", "Sagar", "Suraj"}
	fmt.Println(len(studentNames))
	studentNames[5] = "Mahesh"

	//defer statements
	defer fmt.Println("End of function")
	defer func(slice []string, s int) []string {
		fmt.Println("Elements removed from slices")
		return append(slice[:s], slice[s+1:]...)

	}(studentNames, 0)

	var newSlice = make([]string, len(studentNames))
	copy(newSlice, studentNames)
	newSlice[6] = "Kiran"
	fmt.Println(newSlice)

	// map
	student := map[int]string{
		10: "Siddhant",
		14: "Arti",
		15: "Mayur",
	}

	fmt.Println(student)
	student[15] = "Human"
	delete(student, 10)
	fmt.Println(student)

	//Http get request
	resp, err := http.Get("http://google.com/")

	if err != nil {
		fmt.Println("Error is :")
		fmt.Println(err)
	} else {
		fmt.Println("Date: " + resp.Header.Get("Date"))
	}

}