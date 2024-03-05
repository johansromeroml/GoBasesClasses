package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	customerData, err := os.Open("clase8/ejer01/customers.txt")

	defer func() {
		fmt.Println("Execution finished")
	}()

	defer func(f *os.File) {
		fmt.Print("\nClosing the file ...")
		f.Close()
		fmt.Println(" File Closed")
	}(customerData)

	if err != nil {
		panic(err)
	}

	customerDataContent, err := io.ReadAll(customerData)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n-- File Content: ")
	fmt.Println(string(customerDataContent))
	fmt.Println("\n-- End of file content")
}
