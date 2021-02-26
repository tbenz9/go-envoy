package main

import (
	"flag"
	"fmt"
	"github.com/cloudkucooland/go-envoy"
	// "strconv"
)

func main() {
	var command, host string

	flag.Parse()
	args := flag.Args()
	argc := len(args)
	if argc == 0 {
		command = "unset"
	}
	if argc >= 1 {
		command = args[0]
	}
	if argc > 1 {
		host = args[1]
	}

	var e *envoy.Envoy
	if host == "" {
		fmt.Println("usage: encore command host")
		return
	}

	var err error
	e, err = envoy.New(host)
	if err != nil {
		panic(err)
	}

	switch command {
	case "prod":
		s, err := e.Production()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", s)
	case "now":
		n, err := e.Now()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Now : %f\n", n)
	case "today":
		t, err := e.Today()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Today: %f\n", t)
	case "home":
		s, err := e.Home()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", s)
	case "inventory":
		s, err := e.Inventory()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", s)
	case "info":
		s, err := e.Info()
		if err != nil {
			panic(err)
		}
		// fmt.Printf("%+v\n", s)
		fmt.Println("Serial Number: ", s.Device.Sn)
		fmt.Println("Part Number: ", s.Device.Pn)
		fmt.Println("Software Version: ", s.Device.Software)
	case "stream":
		/* s, err := e.Home()
		if err != nil {
			panic(err)
		} */
		fmt.Printf("working on it...\n")
	default:
		fmt.Println("Valid commands: prod, home, inventory, stream")
	}
}
