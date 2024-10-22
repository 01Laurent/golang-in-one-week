package main

import (
	"fmt"
	// "strings"
	// "sort"
	
)

func main() {
	// booleans
	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	}else {
		fmt.Println("age is not less than 45")
	}

	names := []string{"Mario", "luigi", "yoshi", "peach", "bowser"}
	for index, value := range names {
		if index == 1 {
			fmt.Println("continuinig at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("breaking at pos", index)
			break
		}
		fmt.Printf("the value at pos %v is %v \n", index, value)
	}


	// loops
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value of x is:", i)
	// }

	// names := []string{"Mario", "luigi", "yoshi", "peach"}

	// for i := 0; i <len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for index, value := range names {
	// 	fmt.Printf("the value at index %v is %v \n ", index, value)
	// }

	// for _, value := range names {
	// 	fmt.Printf("the value at is %v \n ", value)
	// 	value = "newstring"
	// }
	// fmt.Println(names)

	// x := 0
	// for x < 5 {
	// 	fmt.Println("value of x is:", x)
	// 	x++
	// }


	// greeting := "Hello There friends"
	// fmt.Println(strings.Contains(greeting, "Hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "Hello", "hi"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "Th"))
	// fmt.Println(strings.Split(greeting, " "))

	// ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	// sort.Ints(ages)
	// fmt.Println(ages)
	
	// index := sort.SearchInts(ages, 30)
	// fmt.Println(index)

	// names := []string{"yoshi", "mario","peach", "bowser", "luigi"}
	// sort.Strings(names)
	// fmt.Println(names)

	// fmt.Println(sort.SearchStrings(names, "bowser"))


	// // strings
	// var nameOne string = "Laurent"
	// var nameTwo = "Legend"
	// var nameThree string

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameOne = "peach"
	// nameThree = "bowser"

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameFour := "yoshi"
	// fmt.Println(nameFour)

	// ints
	// var ageOne int = 20
	// var ageTwo int = 30
	// ageThree := 40

	// fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory
	// var numOne int8 = 25
	// var numTwo int8 = -128
	// var numThree uint8 = 128

	// float
	// var scoreOne float32 = -1.6
	// var scoreTwo float64 = 8765.76543456
	// scoreThree := 1.5

	// print
	// age := 35
	// name := "Laurent"
	// fmt.Print("hello, ")
	// fmt.Print("World! \n")
	// fmt.Print("new line \n")

	// println
	// fmt.Println("hello, Laurent")
	// fmt.Println("goodbye Laurent")
	// fmt.Println("My age is", age, "and my name is", name)

	// printf (formatted strings) %_ = format specifier
	// fmt.Printf("my age is %v and my name is %v \n", age, name)
	// fmt.Printf("my age is %q and my name is %q \n", age, name)
	// fmt.Printf("age is of type %T \n", age)
	// fmt.Printf("you scored %f points! \n", 225.55)
	// fmt.Printf("you scored %0.1f points! \n", 225.55)

	// sprintf (save formatted strings)
	// var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	// fmt.Println("the saved string is:", str)

	// var ages [3]int = [3]int{20, 25, 30}
	// var ages = [3]int{20, 25, 30}

	// names := [4]string{"Yoshi", "Mario", "Peach", "bowser"}
	// names[1] = "luigi"

	// fmt.Println(ages, len(ages))
	// fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	// var scores = []int{100, 50, 60}
	// scores[2] = 25
	// scores = append(scores, 85)
	// fmt.Println(scores, len(scores))

	// slice ranges
	// rangeOne := names[1:3]
	// rangeTwo := names[2:]
	// rangeThree := names[:3]
	// fmt.Println(rangeOne, rangeTwo, rangeThree)

	// rangeOne = append(rangeOne, "koopa")
	// fmt.Println(rangeOne)
}