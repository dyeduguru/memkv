package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/dyeduguru/memkv/client"
	"os"
)

var (
	port   = flag.Int("port", 5555, "Port on which server is listening")
	host   = flag.String("host", "localhost", "Host on which server is listening")
	addKey = flag.Bool("add", false, "Add key value pair")
	getKey = flag.Bool("get", false, "Get value for key")
)

func main() {
	flag.Parse()
	client := client.New(*host, *port)
	input := bufio.NewScanner(os.Stdin)
	if *addKey {
		fmt.Print("Enter Key: ")
		input.Scan()
		key := input.Text()
		fmt.Print("Enter Value: ")
		input.Scan()
		value := input.Text()
		if err := client.AddKey(key, value); err != nil {
			panic(err)
		}
	}
	if *getKey {
		fmt.Print("Enter Key: ")
		input.Scan()
		key := input.Text()
		value, err := client.GetKey(key)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Value: %v\n", value)
	}
}
