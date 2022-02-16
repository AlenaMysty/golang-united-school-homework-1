package main

import (
	"fmt"
	"github.com/kyokomi/emoji"
)

func main() {
	worldMessage := emoji.Sprint("Hello :world_map:!")
	fmt.Println(worldMessage)
}
func GetMessage() string {
	return ""
}
