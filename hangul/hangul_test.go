package hangul

import "fmt"

func ExampleHasConsonantSuffix() {
	fmt.Println(HasConsonantSuffix("Go 언어"))
	fmt.Println(HasConsonantSuffix("그럼"))
	fmt.Println(HasConsonantSuffix("우리 밥 먹고 합시다."))
	// Output:
	// .
}

func Example_array() {
	fruits := [4]string{"사과", "바나나", "토마토", "한라봉"}
	for _, fruit := range fruits {
		if HasConsonantSuffix(fruit) == true {
			fmt.Printf("%s은 맛있다.\n", fruit)
		} else {
			fmt.Printf("%s는 맛있다.\n", fruit)
		}
	}
	// Output:
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 토마토는 맛있다.
	// 한라봉은 맛있다.
}
