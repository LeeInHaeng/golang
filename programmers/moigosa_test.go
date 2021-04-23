package programmers

import "fmt"

type Person struct {
	Id  int
	Dap []int
}

type PersonRes struct {
	Id  int
	Cnt int
}

func Example_solution() {
	testCase := []int{1, 2, 3, 4, 5}
	fmt.Println(solution(testCase))
	// Output:
	// .
}

func solution(answers []int) []int {
	PersonChan := make(chan PersonRes)
	p1 := Person{1, []int{1, 2, 3, 4, 5}}
	p2 := Person{2, []int{2, 1, 3, 2, 4, 2, 5}}
	go func() {
		PersonChan <- getAnswerCnt(answers, p1)
	}()
	fmt.Println(<-PersonChan)
	go func() {
		PersonChan <- getAnswerCnt(answers, p2)
	}()
	fmt.Println(<-PersonChan)

	return []int{}
}

func getAnswerCnt(answers []int, person Person) PersonRes {
	cnt := 0
	recur := len(person.Dap)

	for idx, answer := range answers {
		if answer == person.Dap[idx%recur] {
			cnt++
		}
	}

	return PersonRes{person.Id, cnt}
}
