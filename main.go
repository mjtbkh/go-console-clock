package main

import (
	"fmt"
	"time"
)

func main() {

	type placeholder [5]string

	zero := placeholder{
		"🟩🟩🟩",
		"🟩  🟩",
		"🟩  🟩",
		"🟩  🟩",
		"🟩🟩🟩",
	}

	one := placeholder{
		"🟩🟩  ",
		"  🟩  ",
		"  🟩  ",
		"  🟩  ",
		"🟩🟩🟩",
	}

	two := placeholder{
		"🟩🟩🟩",
		"    🟩",
		"🟩🟩🟩",
		"🟩    ",
		"🟩🟩🟩",
	}

	three := placeholder{
		"🟩🟩🟩",
		"    🟩",
		"🟩🟩🟩",
		"    🟩",
		"🟩🟩🟩",
	}

	four := placeholder{
		"🟩  🟩",
		"🟩  🟩",
		"🟩🟩🟩",
		"    🟩",
		"    🟩",
	}

	five := placeholder{
		"🟩🟩🟩",
		"🟩    ",
		"🟩🟩🟩",
		"    🟩",
		"🟩🟩🟩",
	}

	six := placeholder{
		"🟩🟩🟩",
		"🟩    ",
		"🟩🟩🟩",
		"🟩  🟩",
		"🟩🟩🟩",
	}

	seven := placeholder{
		"🟩🟩🟩",
		"    🟩",
		"    🟩",
		"    🟩",
		"    🟩",
	}

	eight := placeholder{
		"🟩🟩🟩",
		"🟩  🟩",
		"🟩🟩🟩",
		"🟩  🟩",
		"🟩🟩🟩",
	}

	nine := placeholder{
		"🟩🟩🟩",
		"🟩  🟩",
		"🟩🟩🟩",
		"    🟩",
		"🟩🟩🟩",
	}

	digits := [...]placeholder{zero, one, two, three, four, five, six, seven, eight, nine}

	seperator := placeholder{
		"      ",
		"  🟩  ",
		"      ",
		"  🟩  ",
		"      ",
	}

	for {

		// clear the console
		fmt.Print("\033[2J")

		// move the cursor to top left
		fmt.Print("\033[H")

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		// [8][5]string -> [8]placeholder
		clock_array := [...]placeholder{
			digits[hour/10], digits[hour%10],
			seperator,
			digits[min/10], digits[min%10],
			seperator,
			digits[sec/10], digits[sec%10],
		}

		for line := range clock_array[0] {
			for digit := range clock_array {
				fmt.Print(clock_array[digit][line], "\t")
			}
			fmt.Println()
		}

		time.Sleep(time.Second)

	}

}
