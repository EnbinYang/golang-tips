package main

import "fmt"

func processMap(data map[string]int) {
	fmt.Println("processMap")
}

func main() {
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Carol": 35,
	}
	ages["Dave"] = 40

	scores := make(map[string]int)
	scores["Alice"] = 99

	fmt.Println(ages["Alice"], scores["Alice"])

	delete(ages, "Dave")

	age, ok := ages["Alice"]
	if ok {
		fmt.Println("Alice's age is ", age)
	} else {
		fmt.Println("Alice not found")
	}

	for name, age := range ages {
		fmt.Println(name, age)
	}

	fmt.Println(len((ages)))

	processMap(scores)
}
