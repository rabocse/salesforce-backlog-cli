package main

import (
	"flag"
	"fmt"
	"os"
)

/*

Example of current execution

❯ ./cli sr -n=1234567
subcommand 'sr'
  service request number: 1234567
  tail: []


❯ ./cli lv -l="My Active Cases"
subcommand 'lv'
  listview: My Active Cases
  tail: []

*/

func main() {

	sr := flag.NewFlagSet("sr", flag.ExitOnError)
	n := sr.String("n", "", "Service Request Number")

	lv := flag.NewFlagSet("lv", flag.ExitOnError)
	l := lv.String("l", "", "Listview ")

	if len(os.Args) < 2 {
		fmt.Println("	sr - Service Request")
		fmt.Println("	lv - List View")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "sr":
		sr.Parse(os.Args[2:])
		fmt.Println("subcommand 'sr'")
		fmt.Println("  service request number:", *n)
		fmt.Println("  tail:", sr.Args())
	case "lv":
		lv.Parse(os.Args[2:])
		fmt.Println("subcommand 'lv'")
		fmt.Println("  listview:", *l)
		fmt.Println("  tail:", sr.Args())
	default:
		fmt.Println("sr - Service Request")
		fmt.Println("lv - List View (placeholder)")
		os.Exit(1)
	}
}
