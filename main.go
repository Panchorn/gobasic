package main

import (
	"fmt"
	"gobasic/customer"
	"gobasic/user"
	"unicode/utf8"
)

func main() {
    var x int = 2
    y := 5
	println("Hello world!")
	fmt.Printf("Hello %T\n", x)
	fmt.Printf("Hello %T\n", y)

	arr := []int{1,2}
	arr1 := arr[1]
	fmt.Printf("%#v\n", arr)
	fmt.Println(arr1)
	fmt.Println(len(arr))

	arr = append(arr, 4, 5)
	fmt.Printf("%#v\n", arr)
	fmt.Println(len(arr))

	text := "กบ"
	fmt.Printf("%#v %v\n", text, len(text))
	fmt.Printf("%#v %v\n", text, utf8.RuneCountInString(text))

	arrSlice := arr[2:4]
	fmt.Printf("%#v\n", arrSlice)

	countries := map[string]string{"th": "Thailand"}
	countries["en"] = "English"
	fmt.Printf("%#v\n", countries["th"])

	country, ok := countries["jp"]
	if ok {
        fmt.Printf("%v\n", country)
    } else {
        fmt.Println("Country not found")
    }

	for a:=0; a<len(arr); a++ {
		fmt.Println(arr[a])
	}
	for i, v := range arr {
        fmt.Printf("%d: %d\n", i, v)
    }
	b := 0
	for b < len(arr) {
		fmt.Printf("%v: %v\n", b, arr[b])
        b++
	}

	c, _, _, ok := sum(10, 30)
	if ok {
        fmt.Println(c)
    } else {
        fmt.Println("Error: Invalid input")
    }


	// anonymous func
	anonymousFunc := func (a, b int) int {
		return a + b
	}
	value := anonymousFunc(10, 2)
	fmt.Println(value)

	//func param
	addResult := cal(add)
	fmt.Println(addResult)
	subResult := cal(sub)
	fmt.Println(subResult)

	sumArr := sumArr([]int{10,2,3,4,5,6,70,80})
	fmt.Println(sumArr)

	sumArr2 := sumArr2(10,2,3,4,5,6,70,80)
	fmt.Println(sumArr2)

	fmt.Println(customer.Name)
	fmt.Println(customer.Hello())
	fmt.Println(user.Name)

	passValue := 10
	passValue2 := passValue
	fmt.Printf("passValue source %v, dest %v\n", passValue, passValue2)
	
	pointer := &passValue // copy pointer
	valueOfPointer := *pointer
	fmt.Println("passValue\n", &passValue)
	fmt.Println("passValue2\n", &passValue2)
	fmt.Println("pointer\n", pointer)
	fmt.Println("value of pointer\n", valueOfPointer)

	*pointer = 20
	fmt.Println("original of value\n", passValue)
	fmt.Println("passValue2\n", passValue2)
	fmt.Println("value of pointer\n", *pointer)

	passValue = 30
	fmt.Println("original of value\n", passValue)
	fmt.Println("passValue2\n", passValue2)
	fmt.Println("value of pointer\n", *pointer)

	var number int
	sumReference(&number)
	fmt.Println("number \n", number)

	userNon := user.User{
		Name: "Non",
	}
	userNon.SetAge(29)
	fmt.Println(userNon.Hello())
	fmt.Println(userNon)
	fmt.Printf("%#v", userNon.Name)
}

func sum(a, b int) (int, int, int, bool) {
	return a + b, a, b, true
}

// signature (int, int) int
func add(a, b int) int {
	return a + b
}

// signature (int, int) int
func sub(a, b int) int {
	return a - b
}

func hello(name string) string {
	return "hello " + name
}

func cal(f func(int, int) int) int {
	return f(50, 10)
}

func sumArr(a []int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}

func sumArr2(a ...int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}

func sumReference(result *int) {
	a := 20
	b := 30
	*result = a + b
}