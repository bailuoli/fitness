package main

import "fitness/cmd"

func main() {
	defer cmd.Clean()
	cmd.Start()
}
