package main

import (
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"time"
)

type Key struct {
	Row int
	Col int
}

type KeyboardLayout struct {
	Name   string
	Layout [][]string
}

var layouts = []KeyboardLayout{
	{
		Name: "qwerty",
		Layout: [][]string{
			{"`", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "-", "=", "+"},
			{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p", "[", "]", "\\"},
			{"a", "s", "d", "f", "g", "h", "j", "k", "l", ";", "'", "enter"},
			{"shift", "z", "x", "c", "v", "b", "n", "m", ",", ".", "/", "shift"},
			{"ctrl", "alt", "space", "alt", "ctrl"},
		},
	},
	{
		Name: "azerty",
		Layout: [][]string{
			{"²", "&", "é", "\"", "'", "(", "-", "è", "_", "ç", "à", ")", "=", "+"},
			{"a", "z", "e", "r", "t", "y", "u", "i", "o", "p", "^", "$", "\\"},
			{"q", "s", "d", "f", "g", "h", "j", "k", "l", "m", "ù", "*", "enter"},
			{"shift", "w", "x", "c", "v", "b", "n", ",", ";", ":", "!", "shift"},
			{"ctrl", "alt", "space", "alt", "ctrl"},
		},
	},
	{
		Name: "dvorak",
		Layout: [][]string{
			{"`", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "[", "]", "\\"},
			{"'", ",", ".", "p", "y", "f", "g", "c", "r", "l", "/", "=", "+"},
			{"a", "o", "e", "u", "i", "d", "h", "t", "n", "s", "-", "enter"},
			{"shift", ";", "q", "j", "k", "x", "b", "m", "w", "v", "z", "shift"},
			{"ctrl", "alt", "space", "alt", "ctrl"},
		},
	},
}

func main() {
	start := flag.Int("s", 0, "Starting number")
	keyboardType := flag.String("k", "qwerty", "Keyboard type selection")
	finish := flag.Int("f", 5, "Finishing number")
	iterations := flag.Int("i", 2, "Number of iterations")

	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	keyboardLayout, err := getKeyboardLayout(*keyboardType)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < *iterations; i++ {
		walkLength := rand.Intn(*finish-*start+1) + *start
		walk := generateKeyboardWalk(keyboardLayout.Layout, walkLength)
		fmt.Println(walk)
	}
}

func getKeyboardLayout(keyboardType string) (KeyboardLayout, error) {
	for _, layout := range layouts {
		if layout.Name == keyboardType {
			return layout, nil
		}
	}
	return KeyboardLayout{}, errors.New("Invalid keyboard layout")
}

func generateKeyboardWalk(layout [][]string, length int) string {
	keyPositions := make(map[string]Key)
	for rowIndex, row := range layout {
		for colIndex, key := range row {
			keyPositions[key] = Key{Row: rowIndex, Col: colIndex}
		}
	}

	currentKey := layout[rand.Intn(len(layout))][rand.Intn(len(layout[0]))]

	walk := make([]byte, length)
	for i := 0; i < length; i++ {
		walk[i] = currentKey[0]

		currentPosition := keyPositions[currentKey]

		nextRow := rand.Intn(3) - 1
		nextCol := rand.Intn(3) - 1

		newRow := (currentPosition.Row + nextRow + len(layout)) % len(layout)
		newCol := (currentPosition.Col + nextCol + len(layout[newRow])) % len(layout[newRow])

		currentKey = layout[newRow][newCol]
	}

	return string(walk)
}
