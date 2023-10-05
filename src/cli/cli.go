package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Sending GET request to play-game...")
	res, err := http.Get("http://play-game:5001/")
	if err != nil {
		fmt.Println("The GET request to play-game threw an error:", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("The GET request to play-game threw an error:", err)
	} else {
		fmt.Println("The GET request to play-game returned:", string(body))
	}

	fmt.Print("Enter text (entering will cause containers to shut down): ")
	text, _ := reader.ReadString('\n')

	fmt.Println("You entered:", text)
}
