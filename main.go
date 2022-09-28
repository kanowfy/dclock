package main

import (
	"fmt"
	"strconv"
	"time"
)

var digits = [][]string{
	{" ┏━┓ ", "  ╻  ", " ┏━┓ ", " ┏━┓ ", " ╻ ╻ ", " ┏━┓ ", " ┏━┓ ", " ┏━┓ ", " ┏━┓ ", " ┏━┓ ", "   "},
	{" ┃ ┃ ", "  ┃  ", "   ┃ ", "   ┃ ", " ┃ ┃ ", " ┃   ", " ┃   ", "   ┃ ", " ┃ ┃ ", " ┃ ┃ ", " ╻ "},
	{" ┃ ┃ ", "  ┃  ", " ┏━┛ ", "  ━┫ ", " ┗━┫ ", " ┗━┓ ", " ┣━┓ ", "   ┃ ", " ┣━┫ ", " ┗━┫ ", "   "},
	{" ┃ ┃ ", "  ┃  ", " ┃   ", "   ┃ ", "   ┃ ", "   ┃ ", " ┃ ┃ ", "   ┃ ", " ┃ ┃ ", "   ┃ ", " ╻ "},
	{" ┗━┛ ", "  ╹  ", " ┗━━ ", " ┗━┛ ", "   ╹ ", " ┗━┛ ", " ┗━┛ ", "   ╹ ", " ┗━┛ ", " ┗━┛ ", "   "},
}

func main() {
	// Clear the screen
	fmt.Print("\x1b[2J")
	// Move cursor top left of the screen
	fmt.Print("\x1b[;H")
	// Hide cursor
	fmt.Print("\x1b[?25l")
	for {
		currentTime := time.Now()
		currentTimeFormatted := currentTime.Format("15:04:05")
		for _, row := range digits {
			for i := 0; i < len(currentTimeFormatted); i++ {
				c := string(currentTimeFormatted[i])
				switch c {
				case ":":
					fmt.Printf("%s", row[10])
					break
				default:
					idx, _ := strconv.Atoi(c)
					fmt.Printf("%s", row[idx])
				}
			}
			fmt.Print("\n")
		}
		time.Sleep(time.Millisecond * 999)
		// Move cursor up
		fmt.Print("\x1b[5A")
	}
}
