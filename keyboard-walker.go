package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

type Key struct {
	Row int
	Col int
}

func main() {
	start := flag.Int("s", 0, "Starting number")
	keyboardType := flag.String("k", "qwerty", "Keyboard type selection")
	finish := flag.Int("f", 5, "Finishing number")
	iterations := flag.Int("i", 2, "Number of iterations")

	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	keyboardLayout := getKeyboardLayout(*keyboardType)

	for i := 0; i < *iterations; i++ {
		walkLength := rand.Intn(*finish-*start+1) + *start
		walk := generateKeyboardWalk(keyboardLayout, walkLength)
		fmt.Println(walk)
	}
}

func getKeyboardLayout(keyboardType string) [][]string {
	qwerty := [][]string{
		{"`", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "-", "=", "+"},
		{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p", "[", "]", "\\"},
		{"a", "s", "d", "f", "g", "h", "j", "k", "l", ";", "'", "enter"},
		{"shift", "z", "x", "c", "v", "b", "n", "m", ",", ".", "/", "shift"},
		{"ctrl", "alt", "space", "alt", "ctrl"},
	}

	azerty := [][]string{
		{"²", "&", "é", "\"", "'", "(", "-", "è", "_", "ç", "à", ")", "=", "+"},
		{"a", "z", "e", "r", "t", "y", "u", "i", "o", "p", "^", "$", "\\"},
		{"q", "s", "d", "f", "g", "h", "j", "k", "l", "m", "ù", "*", "enter"},
		{"shift", "w", "x", "c", "v", "b", "n", ",", ";", ":", "!", "shift"},
		{"ctrl", "alt", "space", "alt", "ctrl"},
	}

	dvorak := [][]string{
		{"`", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "[", "]", "\\"},
		{"'", ",", ".", "p", "y", "f", "g", "c", "r", "l", "/", "=", "+"},
		{"a", "o", "e", "u", "i", "d", "h", "t", "n", "s", "-", "enter"},
		{"shift", ";", "q", "j", "k", "x", "b", "m", "w", "v", "z", "shift"},
		{"ctrl", "alt", "space", "alt", "ctrl"},
	}

	switch keyboardType {
	case "qwerty":
		return qwerty
	case "azerty":
		return azerty
	case "dvorak":
		return dvorak
	default:
		return qwerty
	}
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
