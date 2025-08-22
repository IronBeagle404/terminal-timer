package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
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

	fmt.Println("BEEP BEEP BEEP")
	Beep()
}

func Beep() {
	// Read the mp3 file
	fileBytes, err := os.ReadFile("./beep.mp3")
	if err != nil {
		panic("Failed to read beep.mp3 : " + err.Error())
	}

	// Convert bytes to reader object to use with mp3 decoder
	fileBytesReader := bytes.NewReader(fileBytes)

	// Decode file
	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
	if err != nil {
		panic("mp3.NewDecoder failed : " + err.Error())
	}

	// Setup Oto context options
	op := &oto.NewContextOptions{}

	// Default recommended value
	op.SampleRate = 44100

	// 2 channels for Stereo
	op.ChannelCount = 2

	// Source format corresponding to go-mp3's format, signed 16bit integers
	op.Format = oto.FormatSignedInt16LE

	// Create Oto context with the defined options
	otoCtx, readyChan, err := oto.NewContext(op)
	if err != nil {
		panic("oto.NewContext failed : " + err.Error())
	}

	// wait for the channel
	<-readyChan

	// Create a new player, paused by default
	player := otoCtx.NewPlayer(decodedMp3)

	// Start playing the sound, returning without waiting for it (async)
	player.Play()

	// Wait for the sound to finish playing
	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}

	// Close the sound
	err = player.Close()
	if err != nil {
		panic("player.Close failed : " + err.Error())
	}
}
