package myqueue

func MyQueueAppend(myqueue []int, n int) []int {
	return append(myqueue, n)
}

func GetMyQueue(myqueue []int) (int, []int) {
	if len(myqueue) < 1 {
		return myqueue[0], make([]int, 0)
	} else {
		return myqueue[0], myqueue[1:]
	}
}
