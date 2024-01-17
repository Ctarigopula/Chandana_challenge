package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("cards.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	lines, _ := strconv.Atoi(line)
	if lines > 0 && lines < 100 {
		for i := 0; i < lines; i++ {
			// Read and print each line
			scanner.Scan()
			card := scanner.Text()
			if valid(card) == true {
				fmt.Println("Valid")
			} else {
				fmt.Println("In-Valid")
			}
		}
	}
	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func valid(card string) (output bool) {
	// Remove hyphen "-" from given card number
	number := regexp.MustCompile(`[^0-9]+`).ReplaceAllString(card, "")
	if len(number) != 16 {
		return false
	}
	if card[:1] != "4" && card[:1] != "5" && card[:1] != "6" {
		return false
	}
	// Check if it may have digits in groups of 4, separated by one hyphen "-"
	if strings.Contains(card, "-") {
		groups := strings.Split(card, "-")
		for _, group := range groups {
			if len(group) != 4 {
				return false
			}
		}
	}
	// Check if it does not have four or more consecutive repeated digits
	for i := 0; i < len(number)-4; i++ {
		if number[i] == number[i+1] && number[i] == number[i+2] && number[i] == number[i+3] {
			return false
		}
	}
	return true
}
