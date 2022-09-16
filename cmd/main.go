//https://gobyexample.com.ru/command-line-subcommands
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected subcommands. Use -h or --help for help")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "symbols":
		fmt.Println("This is symbol list")
	case "convert":
		fmt.Println("This is symbol list")
		convertCmd := flag.NewFlagSet("convert", flag.ExitOnError)
		fromVal := convertCmd.String("from", "USD", "convert -from USD -to RUB")
		toVal := convertCmd.String("to", "USD", "convert -from USD -to RUB")
		convertCmd.Parse(os.Args[2:])
		fmt.Println("From ", *fromVal, " to ", *toVal)
	case "history":
		fmt.Println("This is history")
	default:
		fmt.Printf("unexpected subcommand: \"%v\". Use -h or --help for help\n", os.Args[1])
		os.Exit(1)
	}
}
