package logger

import "log"

func Listen(ch chan string) {
	for {
		log.Println(<-ch)
	}
}
