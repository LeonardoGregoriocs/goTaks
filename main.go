package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	f, err := os.OpenFile("tasks.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening or creating file")
		return
	}
	defer f.Close()

	for {
		ShowMenu()

		option, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		optionInt, err := strconv.Atoi(strings.TrimSpace(option))
		if err != nil {
			fmt.Println("Invalid option", err)
			return
		}

		switch optionInt {
		case 1:
			err := AddTask(*reader, f)
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			err := GetTask(f)
			if err != nil {
				fmt.Println(err)
			}
		case 0:
			fmt.Println("Bye")
			return
		default:
			fmt.Println("Invalid Option")
			continue
		}
	}

}

func ShowMenu() {
	fmt.Println("=== GoTask ===\n 1 - Add task\n 2 - List tasks\n 0 - Exit\n Choose an option:")
}

func AddTask(reader bufio.Reader, f *os.File) error {
	var task string

	fmt.Println("TaskName: ")

	task, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	if strings.TrimSpace(task) == "" {
		return errors.New("the task cannot be empty")
	}

	_, err = f.Write([]byte(task))
	if err != nil {
		return err
	}

	return nil
}

func GetTask(f *os.File) error {
	var count int = 1

	_, err := f.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Printf("%d - %s\n", count, scanner.Text())
		count++
	}

	if count == 1 {
		return errors.New("No tasks found.")
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
