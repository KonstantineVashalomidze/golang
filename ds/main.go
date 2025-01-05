package main

import (
	"ds/lists"
	"fmt"
)

func main() {
	var list *lists.ArrayList[int]
	list = lists.NewArrayList[int]()

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)

	fmt.Println(list.Size())

	removed, err := list.Remove(0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(removed)

	fmt.Println(list.Size())

	fmt.Println(list.String())
}
