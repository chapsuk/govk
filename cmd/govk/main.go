package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/chapsuk/govk"
	"log"
	"os"
)

func main() {
	c := flag.String("c", "", "is client_id of vk application")
	s := flag.String("s", "", "is secret of vk application")
	v := flag.String("v", "5.52", "vk api version")
	i := flag.Int("i", 10, "orders.get size param")
	o := flag.Int("o", 0, "offset param")
	t := flag.Int("t", 0, "enable test_mode")
	flag.Parse()

	if *c == "" || *s == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	cli := govk.NewClient(*c, *s, *v)

	err := cli.Auth()
	handleErr(err)
	log.Printf("\nGotten access_token: %s", cli.AccessToken)

	res, err := cli.OrdersGet(*i, *o, *t)
	handleErr(err)

	prettyPrint(res)
}

func prettyPrint(orders []govk.OrderResponse) {
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
