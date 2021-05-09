package main

import "fmt"

func main() {
	studentMarks := map[string]int{
		"Sneh":   850,
		"Saurav": 841,
		"Rahul":  864,
	}
	studentMarks["Abhishek"] = 970
	studentMarks["Rahul"] = 777
	fmt.Print(studentMarks)
}

