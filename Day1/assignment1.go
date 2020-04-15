package main

import "fmt"

func main() {
	var number int

	fmt.Println("0. Say Hello to User. ")
	fmt.Println("1. Print Asterisk(*) Pyramid. ")
	fmt.Println("2. Print Numbers Pyramid. ")
	fmt.Println("3. Find Factorial. ")
	fmt.Println("4. Check Prime Number. ")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Enter your choice (0 to 4) from above mentioned options:")
	fmt.Scan(&number)

	
	switch {
	case number == 0:
		var userName string
		fmt.Print("Enter UserName :")
		fmt.Scan(&userName)
		sayHello(userName)
	
	case number == 1:
		var asteriskRows int
		fmt.Print("Enter number of rows :")
		fmt.Scan(&asteriskRows)
		drawAsteriskPyramid(asteriskRows)
		
	case number == 2:
		var numberRows int
		fmt.Print("Enter number of rows :")
		fmt.Scan(&numberRows)
		drawNumbersPyramid(numberRows)
		
	case number == 3:
		var n int
		fmt.Print("Enter any positive integer : ")
		fmt.Scan(&n)
		fmt.Print("Factorial is: ", factorial(n))
		
	case number == 4:
		var number int
		fmt.Print("Enter number :")
		fmt.Scan(&number)
		checkPrime(number)
		
	case number > 4:
		fmt.Println("Invalid Choice......")
		fmt.Println("Please try again ...!!!")
	}
}

func drawAsteriskPyramid(rows int) {
	for i := 0; i < rows; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func drawNumbersPyramid(rows int) {
	var number int = 1
	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(number)
			fmt.Print(" ")
			number = number + 1
		}
		fmt.Println()
	}
}

func factorial(n int) int {
	var factVal int = 1
	if n < 0 {
		fmt.Print("Enter number is not valid.")
	} else {
		for i := 1; i <= n; i++ {
			factVal *= i
		}

	}
	return factVal
}

func checkPrime(number int) {
	isPrime := true
	if number == 0 || number == 1 {
		fmt.Printf(" %d is not a  prime number\n", number)
	} else {
		for i := 2; i <= number/2; i++ {
			if number%i == 0 {
				fmt.Printf(" %d is not a  prime number\n", number)
				isPrime = false
				break
			}
		}
		if isPrime == true {
			fmt.Printf(" %d is a prime number\n", number)
		}
	}
}

func sayHello(user string) {
	fmt.Println("Hello "+ user)
}