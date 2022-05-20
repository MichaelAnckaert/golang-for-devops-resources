package main

import (
	"flag"
	"fmt"
)

func main () {
	var branch string

	flag.StringVar(&branch, "branch", "master", 
		  "Branch to deploy, default is 'master'")

	flag.Parse()
	
	fmt.Printf("Deploying branch %v\n", branch)
}
