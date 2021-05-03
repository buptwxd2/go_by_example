package main

import "fmt"

func main()  {
	//m := make(map[string]int)
	//
	//m["a"] = 3
	//fmt.Println(m)

	m := map[string]string {"name": "alex"}
	fmt.Println(m)
	v, exist := m["age"]
	fmt.Println(v, exist)

	m["age"] = "30"
	fmt.Println(m)

	delete(m, "age2")
	fmt.Println(m)

	for k,v := range m{
		fmt.Println("Key is", k, "Value is", v)
	}
	fmt.Println("Length of m is", len(m))

	text := "Hello中文"
	bs := []byte (text)
	fmt.Println(bs)
	for i,c := range text{
		fmt.Println("Index", i, "is", c)
	}

}
