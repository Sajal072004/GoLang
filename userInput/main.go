package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rationg for our Pizza:")

	//comma ok // err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating : " , input)

}
