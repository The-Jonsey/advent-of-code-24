package common

import (
	"log"
	"os"
)

func GetInput(subDir string) string {
	dat, err := os.ReadFile(subDir + "/input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	return string(dat)
}
