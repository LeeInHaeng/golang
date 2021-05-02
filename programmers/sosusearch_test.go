package programmers

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func Example_solution_sosusearch() {
	testNumbers := "17"
	fmt.Println(solution_sosusearch(testNumbers))
	testNumbers = "011"
	fmt.Println(solution_sosusearch(testNumbers))
	// Output:
	// .
}

func solution_sosusearch(numbers string) int {
	answer := 0
	for res := range getSosuSearchGo(numbers) {
		answer = res
	}
	return answer
}

func getSosuSearchGo(numbers string) <-chan int {

	chanCnt := make(chan int)
	cnt := 0
	var wg sync.WaitGroup
	wg.Add(2)

	var allNums []int
	var removeDuplNums map[int]int
	var allPrimes []int

	go func() {
		defer wg.Done()
		getAllNumsRecur(&allNums, "", numbers)
	}()

	go func() {
		defer wg.Done()
		allPrimes = getPrime(numbers)
	}()

	go func() {
		wg.Wait()
		removeDuplNums = getRemoveDuplNums(allNums)
		for _, val := range removeDuplNums {
			if allPrimes[val] == 0 {
				cnt++
			}
		}
		chanCnt <- cnt
		close(chanCnt)
	}()

	return chanCnt
}

func getRemoveDuplNums(req []int) map[int]int {
	mapNums := map[int]int{}
	for _, num := range req {
		_, exist := mapNums[num]
		if exist == false {
			mapNums[num] = num
		}
	}
	return mapNums
}

func getAllNumsRecur(res *[]int, primNum string, extraNum string) {
	splitNumbers := strings.Split(extraNum, "")
	for idx, snum := range splitNumbers {
		inum, _ := strconv.Atoi(primNum + snum)
		*res = append(*res, inum)
		getAllNumsRecur(res, primNum+snum, extraNum[:idx]+extraNum[idx+1:])
	}
}

func getPrime(numbers string) []int {
	splitNumbers := strings.Split(numbers, "")
	sort.Sort(sort.Reverse(sort.StringSlice(splitNumbers)))
	stringNumber := ""
	for _, stringNum := range splitNumbers {
		stringNumber += stringNum
	}
	maxNum, _ := strconv.Atoi(stringNumber)

	primes := make([]int, maxNum+1)
	for i := 0; i < maxNum+1; i++ {
		if i == 0 || i == 1 {
			primes[i] = 1
		} else {
			if primes[i] == 0 {
				for j := 2; j*i < maxNum+1; j++ {
					primes[j*i] = 1
				}
			}
		}
	}

	return primes
}
