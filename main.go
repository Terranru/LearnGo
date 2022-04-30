package main

import "fmt"

func main() {
	i := 12
	var name string = "Denis"

	Great(i, name)
	fio("Loktionov", "Denis", "Vladimirovich")
	a, b := 4, 5
	fmt.Printf("Сумма %v + %v = %v\n", a, b, Sum(a, b))
	summa, multi := MultiSum(a, b)
	fmt.Printf("Сумма %v Произведение %v\n", summa, multi)

	div := func(x, y int) float32 {
		return float32(x) / float32(y)
	}
	fmt.Println(div(10, 3))
}

func Great(i int, name string) {
	println("Hello!")
	fmt.Printf("%T\n", i)
	fmt.Printf("Hi, %s\n", name)

}

func fio(fam, name, surname string) {
	fmt.Printf("Your name is %s %s %s\n", name, surname, fam)
}

func Sum(a, b int) int {
	return a + b
}

func MultiSum(a, b int) (int, int) {
	return a + b, a * b
}
