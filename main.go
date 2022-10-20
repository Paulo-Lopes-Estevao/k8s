package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {

	/* animals := []string{"Colorado", "Utah", "Wisconsin", "Oregon"}

	fmt.Println(len(animals))

	sort.Slice(animals, func(i, j int) bool {
		fmt.Println(animals[i])
		fmt.Println("---------------------------")
		fmt.Println(animals[j])
		return len(animals[i]) < len(animals[j])
	})
	fmt.Println(animals) */

	/* var num = []int32{1, 2, 3, 4, 10, 11}
	dr := simpleArraySum2(num)
	fmt.Println(dr)

	a := []int32{17, 28, 30}
	b := []int32{99, 16, 8}
	Triplets := compareTriplets(a, b)
	fmt.Println(Triplets) */

	//fizzBuzz(15)
	// mod := ModifyString("oll123eH56")
	// println(mod)

	go Server()

}

func Server() {
	oddChan := make(chan int32)
	//evenChan := make(chan int32)
	//arr := []int32{1, 2, 3, 4, 5}
	var i int32
	go func() {
		for i = 0; i < 4; i++ {
			fmt.Println("NÃO", i)
			oddChan <- 32
		}
	}()

	od := <-oddChan
	fmt.Println("oddChan", od)

}

func simpleArraySum(ar []int32) int32 {
	var sum int32 = 0
	for k, v := range ar {
		if 0 < v && ar[k] <= 1000 {
			sum += v
		}
	}
	return sum
}

func simpleArraySum2(ar []int32) int32 {
	n := len(ar)
	var sum int32 = 0
	for i := 0; i < n; i++ {
		if 0 < n && ar[i] <= 1000 {
			sum += ar[i]
		}
	}
	return sum
}

func compareTriplets(a []int32, b []int32) []int32 {

	var Alice, Bob int32
	n := len(a)

	for i := 0; i < n; i++ {
		if 1 <= a[i] && a[i] <= 100 || 1 <= b[i] && b[i] <= 100 {
			if a[i] > b[i] {
				Alice++
			} else if a[i] < b[i] {
				Bob++
			}
		}
	}
	return []int32{Alice, Bob}
}

func fizzBuzz(n int32) {
	var i int32
	for i = 1; i <= n; i++ {
		tres := i % 3
		cinco := i % 5
		if tres == 0 && cinco == 0 {
			fmt.Println("FizzBuzz")
		} else if tres == 0 {
			fmt.Println("Fizz")
		} else if cinco == 0 {
			fmt.Println("Buzz")
		} else if tres != 0 && cinco != 0 {
			fmt.Println(i)
		}

	}
}

func ModifyString(str string) string {
	num := regexp.MustCompile(`\d`)
	result := strings.ReplaceAll(str, " ", "")
	str1 := num.ReplaceAllString(result, "")

	var reverse string
	for _, v := range str1 {
		reverse = string(v) + reverse
	}
	return reverse

}
