package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	mp1 := map[string]Person{
		"4603678901": {
			Name: "Ivanov",
			Age:  20,
		},
		"4509678543": {
			Name: "Petrov",
			Age:  30,
		},
	}
	var num string
	fmt.Scanf("%s", &num)

	if person, ok := mp1[num]; ok {
		fmt.Printf("key: %s value %+v\n", num, person)
	} else {
		fmt.Println("not found")
	}
}
