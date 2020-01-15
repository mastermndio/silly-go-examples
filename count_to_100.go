package main

import "fmt"

func main() {

	names := []string{"james", "aaron", "james", "tom", "brady"}

	for x := 0; x < len(names); x++ {
		if names[x] == "james" {
			fmt.Println("Hello James")
		}
	}
}
