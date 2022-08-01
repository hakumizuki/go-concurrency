package main

import (
	"github.com/hakumizuki/go-concurrency/mutex"
	"github.com/hakumizuki/go-concurrency/wg"
)

func main() {
	wg.Main(false)
	mutex.Main(true)
}
