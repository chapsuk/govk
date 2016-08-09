package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/chapsuk/govk"
)

var (
	c = flag.String("c", "", "is client_id of vk application")
	s = flag.String("s", "", "is secret of vk application")
	v = flag.String("v", "5.53", "vk api version")
	// cmd
	cmd = flag.String("cmd", "", "vk method name")
	// params
	offset = flag.Int("offset", 0, "offset param")
	count  = flag.Int("count", 10, "count param")
	// database params
	needAll = flag.Bool("need-all", false, "database.getCountries needAll param")
	code    = flag.String("code", "", "databse.getCountries code param")
	// orders.get params
	ogt = flag.Int("ogt", 0, "orders.get enable test_mode")
)

func main() {
	flag.Parse()

	needAuth := needAuth(*cmd)
	cli := govk.NewClient(*c, *s, *v)
	if needAuth {
		if *c == "" || *s == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}
		err := cli.Auth()
		handleErr(err)
		log.Printf("\nGotten access_token: %s", cli.AccessToken)
	} else {
		log.Print("\nWithout auth")
	}

	switch *cmd {
	case "orders.get":
		res, err := cli.OrdersGet(*count, *offset, *ogt)
		handleErr(err)
		printOrdersGet(res)
	case "database.getCountries":
		res, err := cli.DatabaseGetCountries(*count, *offset, *needAll, *code)
		handleErr(err)
		printGetCountries(res)
	default:
		handleErr(fmt.Errorf("undefined method"))
	}
}

func needAuth(cmd string) bool {
	switch cmd {
	case "orders.get":
		return true
	case "database.getCountries":
		return false
	default:
		return false
	}
}

func printGetCountries(countries []govk.CountryResponse) {
	var buffer bytes.Buffer
	buffer.WriteString("\nResult database.getCountries method\n")

	for _, o := range countries {
		m := fmt.Sprintf("%+v\n", o)
		buffer.WriteString(m)
	}
	buffer.WriteString("\n")
	log.Print(buffer.String())
}

func printOrdersGet(orders []govk.OrderResponse) {
	var buffer bytes.Buffer
	buffer.WriteString("\nResult orders.get method\n")

	for _, o := range orders {
		m := fmt.Sprintf("%+v\n", o)
		buffer.WriteString(m)
	}
	buffer.WriteString("\n")
	log.Print(buffer.String())
}

func handleErr(err error) {
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
