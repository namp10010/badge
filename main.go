package main

import (
	"flag"
	"log"
)

func main() {
	c := flag.Float64("c", 70.0, "test coverage float")
	o := flag.String("o", "cover_badge.png", "output file")
	flag.Parse()

	err := drawBadge(*c, *o)
	if err != nil {
		log.Println(err)
	}
}
