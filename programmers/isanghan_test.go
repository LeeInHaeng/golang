package programmers

import (
	"fmt"
	"strings"
	"sync"
)

type StringReq struct {
	Idx int
	Req string
}

func Example_solution_isanghan() {
	testCase := "try hello world"
	fmt.Println(solution_isanghan(testCase))
	// Output:
	// .
}

func solution_isanghan(s string) string {
	splitedString := make([]string, len(strings.Split(s, " ")))
	for StringReqChan := range convertStringGo(s) {
		splitedString[StringReqChan.Idx] = StringReqChan.Req
	}
	return strings.Join(splitedString, " ")
}

func convertStringGo(s string) <-chan StringReq {
	splitSpaceString := strings.Split(s, " ")
	convertedChan := make(chan StringReq)

	var wg sync.WaitGroup
	wg.Add(len(splitSpaceString))

	for idx, splited := range splitSpaceString {
		sr := StringReq{idx, splited}
		go func(req StringReq) {
			defer wg.Done()
			convertedChan <- convertString(req)
		}(sr)
	}

	go func() {
		wg.Wait()
		close(convertedChan)
	}()

	return convertedChan
}

func convertString(sr StringReq) StringReq {
	converted := ""

	splitString := strings.Split(sr.Req, "")
	for idx, ss := range splitString {
		if idx%2 == 0 {
			converted += strings.ToUpper(ss)
		} else {
			converted += strings.ToLower(ss)
		}
	}

	return StringReq{sr.Idx, converted}
}
