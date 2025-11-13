package chapter

import (
	"log"
	"os"
)

func Error() {
	_, err := os.Open("hoge.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}
