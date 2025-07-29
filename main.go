package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Error : Not enough arguments. Run program with the help flag to display some guidance.")
		return
	}

	if os.Args[1] == "help" || strings.Count(os.Args[1], ":") != 2 {
		fmt.Println("To setup a timer, you can simply run the program followed by the duration, like this : go run . 1:30:0 (this will setup a timer of 1 hour and 30 minutes)\nPlease format timer duration as [hours:minutes:seconds]\nYou can format the numbers in single or double digits, as you prefer")
		return
	}

	input := strings.Split(os.Args[1], ":")

	for _, data := range input {
		dataInt, err := strconv.Atoi(data)
		if err != nil {
			fmt.Printf("Error : %v\n", err)
			return
		}
		if len(data) > 2 || dataInt < 0 {
			fmt.Println("Error : wrong input format")
			return
		}
	}

	hours, _ := strconv.Atoi(input[0])
	minutes, _ := strconv.Atoi(input[1])
	seconds, _ := strconv.Atoi(input[2])

	if minutes > 59 || seconds > 59 {
		fmt.Println("Error : wrong input format")
		return
	}

	fmt.Printf("You have set a timer for %d hour(s), %d minute(s), and %d second(s)\n", hours, minutes, seconds)

	totalTimerValue := seconds + minutes*60 + hours*60*60

	fmt.Printf("%02d:%02d:%02d\n", hours, minutes, seconds)
	for x := totalTimerValue; x > 0; x-- {
		seconds--

		if seconds < 0 {
			minutes--
			seconds = 59
		}

		if minutes < 0 {
			hours--
			minutes = 59
		}

		time.Sleep(1 * time.Second)
		fmt.Printf("%02d:%02d:%02d\n", hours, minutes, seconds)
	}

	for {
		fmt.Printf("BEEP BEEP BEEP\n")
		time.Sleep(1 * time.Second)
	}
}
