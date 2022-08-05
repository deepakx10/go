// trunk-ignore(gofmt)
package main

import "fmt"

func main() {
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	V, ok := m["hello"]
	fmt.Println(V, ok)

	V, ok = m["world"]
	fmt.Println(V, ok)

	V, ok = m["goodbye"]
	fmt.Println(V, ok)
}