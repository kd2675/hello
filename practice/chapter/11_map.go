package chapter

import "fmt"

func Map() {
	sample := map[string]int{
		"a": 1,
		"b": 2,
	}

	fmt.Println(sample)

	m := make(map[int]string)

	m[0] = "zero"
	m[1] = "one"
	fmt.Println(m)

	fmt.Println(m[0])

	value, exists := m[2]
	fmt.Println(value, exists)

	for key, value := range m {
		fmt.Println(key, value)
	}
}
