package main //import github.com/oshankkumar/prettyjson

import (
	"log"
)

func main() {
	if err := NewPrettyJsonCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}

type WriterFunc func(b []byte) (int, error)

func (f WriterFunc) Write(b []byte) (int, error) {
	return f(b)
}
