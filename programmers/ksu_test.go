package programmers

import (
	"fmt"
	"sort"
	"sync"
)

type CommandReq struct {
	Id      int
	Command []int
}

type CommandRes struct {
	Id  int
	Res int
}

func Example_solution_ksu() {
	testArray := []int{1, 5, 2, 6, 3, 7, 4}
	testCommands := [][]int{
		{2, 5, 3},
		{4, 4, 1},
		{1, 7, 3},
	}
	fmt.Println(solution_ksu(testArray, testCommands))
	// Output:
	// .
}

func solution_ksu(array []int, commands [][]int) []int {
	result := make([]int, len(commands))
	for ksuNumber := range getKsuNumberGo(array, commands) {
		result[ksuNumber.Id] = ksuNumber.Res
	}
	return result
}

func getKsuNumberGo(array []int, commands [][]int) <-chan CommandRes {
	commandChan := make(chan CommandRes)

	var wg sync.WaitGroup
	wg.Add(len(commands))

	for idx, command := range commands {
		commandReq := CommandReq{idx, command}
		go func(a []int, cq CommandReq) {
			defer wg.Done()
			commandChan <- getKsuNumber(a, cq)
		}(array, commandReq)
	}

	go func() {
		wg.Wait()
		close(commandChan)
	}()

	return commandChan
}

func getKsuNumber(array []int, req CommandReq) CommandRes {
	start := req.Command[0] - 1
	end := req.Command[1]
	pos := req.Command[2] - 1
	size := req.Command[1] - req.Command[0] + 1

	sliceArray := make([]int, size)
	copy(sliceArray, array[start:end])
	sort.Ints(sliceArray)
	return CommandRes{req.Id, sliceArray[pos]}
}
