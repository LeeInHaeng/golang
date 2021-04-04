package myqueue

import "fmt"

func ExampleMyQueue() {
	myDataList := make([]int, 0)
	myDataList = MyQueueAppend(myDataList, 10)
	myDataList = MyQueueAppend(myDataList, 20)
	myDataList = MyQueueAppend(myDataList, 30)
	myDataList = MyQueueAppend(myDataList, 40)
	myDataList = MyQueueAppend(myDataList, 50)

	data, myDataList := GetMyQueue(myDataList)
	fmt.Println(data)
	data, myDataList = GetMyQueue(myDataList)
	fmt.Println(data)
	data, myDataList = GetMyQueue(myDataList)
	fmt.Println(data)

	myDataList = MyQueueAppend(myDataList, 100)
	myDataList = MyQueueAppend(myDataList, 110)
	myDataList = MyQueueAppend(myDataList, 120)

	data, myDataList = GetMyQueue(myDataList)
	fmt.Println(data)
	data, myDataList = GetMyQueue(myDataList)
	fmt.Println(data)

	// Output:
	// 10
	// 20
	// 30
	// 40
	// 50
}
