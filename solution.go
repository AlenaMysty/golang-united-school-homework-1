package main

import (
	"fmt"
	"github.com/kyokomi/emoji"
)

func main() {
	worldMessage := GetMessage()
	fmt.Println(worldMessage)
}
func GetMessage() string {
	return emoji.Sprint("Hello :world_map:!")
}
