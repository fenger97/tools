package main

import "gologs/examples"

func main() {
	sl := examples.NewSeelogger()
	sl.Start()
}
