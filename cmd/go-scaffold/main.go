package main

import "fmt"

func main() {
	fmt.Print("Enter project name:")
	var projectName string
	fmt.Scanf("%s", &projectName)
	fmt.Printf("project name: %s", projectName)
}
