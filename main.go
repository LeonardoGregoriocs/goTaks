package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var task string

	f, err := os.OpenFile("tasks.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening or creating file")
		return
	}
	defer f.Close()

	reader := bufio.NewReader(os.Stdin)
	task, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	if strings.TrimSpace(task) == "" {
		fmt.Println("The task cannot be empty.")
		return
	}

	_, err = f.Write([]byte(task))
	if err != nil {
		panic(err)
	}

}
