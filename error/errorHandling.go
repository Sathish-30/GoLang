package error

import (
	"log"
	"os"
)

func CheckError() {
	_, err := os.Open("error.txt")
	if err != nil {
		log.Fatal(err)
	}
}
