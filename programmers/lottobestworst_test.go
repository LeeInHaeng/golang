package programmers

import "fmt"

func Example_solution_lottobestworst() {
	testLottos := []int{44, 1, 0, 0, 31, 25}
	testWinNums := []int{31, 10, 45, 1, 6, 19}
	fmt.Println(solution_lottobestworst(testLottos, testWinNums))
	// Output:
	// .
}

func solution_lottobestworst(lottos []int, win_nums []int) []int {
	lottoRes := map[int]int{
		6: 1,
		5: 2,
		4: 3,
		3: 4,
		2: 5,
		1: 6,
		0: 6,
	}
	cnt := 0
	zeroCnt := 0
	for _, lotto := range lottos {
		if lotto == 0 {
			zeroCnt++
		} else {
			for _, win := range win_nums {
				if lotto == win {
					cnt++
					break
				}
			}
		}
	}
	return []int{lottoRes[cnt+zeroCnt], lottoRes[cnt]}
}
