package programmers

import (
	"fmt"
)

type Printer struct {
	Id       int
	Priority int
}

func Example_solution_printer() {
	testPri := []int{2, 1, 3, 2}
	testLoc := 2
	fmt.Println(solution_printer(testPri, testLoc))
	testPri = []int{1, 1, 9, 1, 1, 1}
	testLoc = 0
	fmt.Println(solution_printer(testPri, testLoc))
	// Output:
	// .
}

func solution_printer(priorities []int, location int) int {
	printer := make([]Printer, len(priorities))
	for idx, p := range priorities {
		printer[idx] = Printer{idx, p}
	}

	answer := 0
	printerLen := len(printer)
	for true {
		loopEnd := true
		for i := 1; i < printerLen; i++ {
			if printer[0].Priority < printer[i].Priority {
				popPrinter := printer[0]
				copy(printer[0:], printer[1:])
				printer[printerLen-1] = popPrinter
				loopEnd = false
				break
			}
		}
		if loopEnd == true {
			answer++
			if printer[0].Id == location {
				break
			}
			copy(printer[0:], printer[1:printerLen])
			printerLen--
			printer = printer[:printerLen]
		}
	}

	return answer
}
