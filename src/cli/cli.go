package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter text (entering will cause containers to shut down): ")
	text, _ := reader.ReadString('\n')

	fmt.Println("You entered:", text)
}
