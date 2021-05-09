package programmers

import "fmt"

type Dejin struct {
	Id int
}

func Example_solution_dejin() {
	testN := 8
	testA := 4
	testB := 7
	fmt.Println(solution_dejin(testN, testA, testB))
	// Output:
	// .
}

func solution(n int, a int, b int) int {
	// 시간초과
	return solution_dejin(n, a, b)
}

func solution_dejin(n int, a int, b int) int {
	answer := 0
	var persons []Dejin
	for i := 1; i < n+1; i++ {
		persons = append(persons, Dejin{i})
	}

	loop := true
	for loop == true {
		answer++
		var winners []Dejin
		for i := 0; i < len(persons)-1; i += 2 {
			dualPerson := persons[i : i+2]
			if (dualPerson[0].Id == a && dualPerson[1].Id == b) || (dualPerson[1].Id == b && dualPerson[0].Id == a) {
				loop = false
				break
			} else if dualPerson[0].Id == a {
				winners = append(winners, dualPerson[0])
			} else if dualPerson[0].Id == b {
				winners = append(winners, dualPerson[0])
			} else if dualPerson[1].Id == a {
				winners = append(winners, dualPerson[1])
			} else {
				winners = append(winners, dualPerson[1])
			}
		}
		persons = winners
	}

	return answer
}
