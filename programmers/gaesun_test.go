package programmers

import (
	"fmt"
	"math"
	"sync"
)

type ProgressReq struct {
	Id       int
	Progress int
	Speed    int
}

type ProgressRes struct {
	Id  int
	Day int
}

func Example_solution_gaesun() {
	testProgress := []int{93, 30, 55}
	testSpeed := []int{1, 30, 5}
	fmt.Println(solution_gaesun(testProgress, testSpeed))
	testProgress = []int{95, 90, 99, 99, 80, 99}
	testSpeed = []int{1, 1, 1, 1, 1, 1}
	fmt.Println(solution_gaesun(testProgress, testSpeed))
	// Output:
	// .
}

func solution_gaesun(progresses []int, speeds []int) []int {
	leaveDays := make([]int, len(progresses))
	answer := make([]int, 1)
	for completedDay := range getCompletedDayGo(progresses, speeds) {
		leaveDays[completedDay.Id] = completedDay.Day
	}
	for len(leaveDays) > 0 {
		job := 0
		popDay := leaveDays[0]
		for _, day := range leaveDays {
			if popDay >= day {
				job++
			} else {
				break
			}
		}
		answer = append(answer, job)
		copy(leaveDays[0:], leaveDays[job:])
		leaveDays = leaveDays[:len(leaveDays)-job]
	}
	return answer[1:]
}

func getCompletedDayGo(progresses []int, speeds []int) <-chan ProgressRes {
	resChan := make(chan ProgressRes)

	var wg sync.WaitGroup
	wg.Add(len(progresses))

	for idx, progress := range progresses {
		req := ProgressReq{idx, progress, speeds[idx]}
		go func(pq ProgressReq) {
			defer wg.Done()
			resChan <- getCompletedDay(pq)
		}(req)
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()

	return resChan
}

func getCompletedDay(pq ProgressReq) ProgressRes {
	leaveProgress := 100 - pq.Progress
	return ProgressRes{pq.Id, int(math.Ceil(float64(leaveProgress) / float64(pq.Speed)))}
}
