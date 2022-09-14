package main

import "fmt"

func main () {
	// fmt.Println ("Hello World") 

	// var age int = 12
	// var num float64 = 2.3456
	// var str = "Hello World"
	// fmt.Println(age, num, str)

	// var a = 10
	// var b = 5
	// var res int
	// res = a % b
	// fmt.Println("Result is", res)

	// const pi = 3.14
	// var web string = "itProger"
	// fmt.Println(len(web))
	// fmt.Println(web + " is cool website")

	// var num float64 = 4.3254357
	// fmt.Printf("%f \n", num);
	// fmt.Printf("%.2f \n", num);
	// fmt.Printf("%T \n", num);
	// var isDone bool = true
	// fmt.Printf("%t \n", isDone)

	// var age = 20
	// if age<5 {
	// 	fmt.Println("Вам пора в детский сад")
	// } else if age == 5 {
	// 	fmt.Println("Вам пора в прескул")
	// } else if (age>5) && (age<=18) {
	// 	var grade = age - 5
	// 	fmt.Println("Пора идти в", grade, "класс")
	// } else {
	// 	fmt.Println("Вам пора в университет")
	// }

	// var age = 10
	// switch age {
	// 	case 5: fmt.Println("Вам 5 лет")
	// 	case 15: fmt.Println("Вам 15 лет")
	// 	case 10: fmt.Println("Вам 10 лет")
	// 	default: fmt.Println("Вам неизвестно сколько лет")
	// }

	// var i = 1
	// for i<=10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for i := 0; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// var arr[3] int
	// arr[0] = 45
	// arr[1] = 97
	// arr[2] = 76 
	// fmt.Println(arr[1])
	// nums := [3]float64 {4.23, 5.23, 98.1}
	// for i, value := range nums {
	// 	fmt.Println(value, i)
	// }

	// ассоциативный массив
	// webSites := make(map[string]float64)
	// webSites["itProger"] = 0.8
	// webSites["yandex"] = 0.99
	// fmt.Println(webSites["itProger"])

	// var a = 3
	// var b = 2
	// var r int
	// r = summ(a, b)
	// a, b = summ(a, b)
	// fmt.Println(a, b)

	// замыкание
	// var num = 3
	// multiple := func() int {
	// 	num *= 2
	// 	return num
	// }
	// fmt.Println(multiple())

	// Откладование
	// defer two()
	// one()

	// указатели
	// var x = 0
	// pointer(x)
	// fmt.Println(x)//0
	// pointer_2(&x)
	// fmt.Println(x)//2

	// Структуры
	bob := Cats{"Bob", 7, 0.87}
	fmt.Println("Bob age is", bob.age)
	fmt.Println("Bob function is", bob.test())
}

// func summ (num_1 int, num_2 int) (int, int) {
// 	var res int
// 	res = num_1 + num_2
// 	return res, num_1 * num_2
// }

// func one () {
// 	fmt.Println("1")
// }

// func two () {
// 	fmt.Println("2")
// }

// func pointer (x int) {
// 	x = 2
// }
// func pointer_2 (x *int) {
// 	*x = 2
// }

type Cats struct {
	name string
	age int
	happiness float64
}

func(cat *Cats) test() float64 {
	return float64(cat.age) * cat.happiness
}
