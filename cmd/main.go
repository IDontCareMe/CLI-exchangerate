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
		tips := "convert -from USD -to RUB 100"
		convertCmd := flag.NewFlagSet("convert", flag.ExitOnError)
		fromVal := convertCmd.String("from", "USD", tips)
		toVal := convertCmd.String("to", "USD", tips)
		convertCmd.Parse(os.Args[2:])
		fmt.Println("From", *fromVal, "to", *toVal, convertCmd.Args()[0])
	case "history":
		tips := "history -from USD -to RUB -data_from 20220101 -data_to 20220201 100"
		historyCmd := flag.NewFlagSet("history", flag.ExitOnError)
		fromVal := historyCmd.String("from", "USD", tips)
		toVal := historyCmd.String("to", "USD", tips)
		fromDate := historyCmd.String("data_from", "", tips)
		toDate := historyCmd.String("data_to", "", tips)
		historyCmd.Parse(os.Args[2:])
		fmt.Println("From", *fromVal, "to", *toVal, "date from", *fromDate, "date to", *toDate, historyCmd.Args()[0])
	case "-h", "--help":
		fmt.Println("help will be soon")
	default:
		fmt.Printf("unexpected subcommand: \"%v\". Use -h or --help for help\n", os.Args[1])
		os.Exit(1)
	}
}
