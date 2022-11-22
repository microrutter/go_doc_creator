package utils

import (
	"log"
)

func Check(e error, log *log.Logger) {
	if e != nil {
		log.Print(e)
	}
}
