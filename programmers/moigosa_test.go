package programmers

import (
	"fmt"
	"sort"
	"sync"
)

type Person struct {
	Id  int
	Dap []int
}

type PersonRes struct {
	Id  int
	Cnt int
}

func Example_solution_moigosa() {
	testCase := []int{1, 2, 3, 4, 5}
	fmt.Println(solution_moigosa(testCase))
	testCase = []int{1, 3, 2, 4, 2}
	fmt.Println(solution_moigosa(testCase))
	testCase = []int{0, 0, 0, 0, 0}
	fmt.Println(solution_moigosa(testCase))
	testCase = []int{5}
	fmt.Println(solution_moigosa(testCase))
	testCase = []int{5, 5, 5, 1, 4, 1}
	fmt.Println(solution_moigosa(testCase))
	testCase = []int{1, 1, 1, 1}
	fmt.Println(solution_moigosa(testCase))
	// Output:
	// .
}

func solution_moigosa(answers []int) []int {
	bestCnt := 0
	result := make([]int, 1)
	for personCnt := range getAnswerCntGo(answers) {
		fmt.Println(personCnt)
		if bestCnt < personCnt.Cnt {
			result = make([]int, 1)
			result[0] = personCnt.Id
			bestCnt = personCnt.Cnt
		} else if bestCnt == personCnt.Cnt {
			result = append(result, personCnt.Id)
		}
	}
	if bestCnt == 0 {
		return []int{1, 2, 3}
	}
	sort.Ints(result)
	return result
}

func getAnswerCntGo(answers []int) <-chan PersonRes {
	PersonChan := make(chan PersonRes)
	persons := []Person{
		Person{1, []int{1, 2, 3, 4, 5}},
		Person{2, []int{2, 1, 2, 3, 2, 4, 2, 5}},
		Person{3, []int{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}},
	}

	var wg sync.WaitGroup
	wg.Add(len(persons))

	for _, person := range persons {
		go func(p Person) {
			defer wg.Done()
			PersonChan <- getAnswerCnt(answers, p)
		}(person)
	}

	go func() {
		wg.Wait()
		close(PersonChan)
	}()

	return PersonChan
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
