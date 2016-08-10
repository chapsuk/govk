package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/chapsuk/govk"
)

var (
	version = "5.53"
	// server credentials
	c = flag.String("c", "", "is client_id of vk application")
	s = flag.String("s", "", "is secret of vk application")
	// cmd
	cmd = flag.String("cmd", "", "vk method name")
	// params
	language    = flag.Int("lang", 0, "response language")
	offset      = flag.Int("offset", 0, "offset param")
	count       = flag.Int("count", 1, "count param")
	userID      = flag.Int("user", 0, "user_id")
	accessToken = flag.String("access_token", "", "user access_token")
	// database params
	needAll   = flag.Bool("need-all", false, "database.getCountries needAll param")
	code      = flag.String("code", "", "databse.getCountries code param")
	countryID = flag.Int("country", 1, "database.getCities country_id param")
	regionID  = flag.Int("region", 0, "database.getCities region_id param")
	query     = flag.String("query", "", "database.getCities q param")
	cityIds   = flag.String("city-ids", "", "database.getCitiesById city_ids param")
	// orders.get params
	ogt = flag.Int("ogt", 0, "orders.get enable test_mode")
)

func main() {
	flag.Parse()

	needAuth := needAuth(*cmd)
	cli := govk.NewClient(*c, *s, version, *language)
	if needAuth {
		if *c == "" || *s == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}
		err := cli.Auth()
		handleErr(err)
		log.Printf("\nGotten access_token: %s", cli.AccessToken)
	} else {
		log.Print("\nWithout server auth")
	}

	if *cityIds != "" && *count == 1 {
		*count = len(strings.Split(*cityIds, ","))
	}

	r := make([]interface{}, *count)
	switch *cmd {
	case "orders.get":
		res, err := cli.OrdersGet(*count, *offset, *ogt)
		handleErr(err)
		for k, v := range res {
			r[k] = v
		}
		printResult(*cmd, r)
	case "database.getCountries":
		res, err := cli.DatabaseGetCountries(*count, *offset, *needAll, *code)
		handleErr(err)
		for k, v := range res {
			r[k] = v
		}
		printResult(*cmd, r)
	case "database.getCities":
		res, err := cli.DatabaseGetCities(*count, *offset, *needAll, *countryID, *regionID, *query)
		handleErr(err)
		for k, v := range res {
			r[k] = v
		}
		printResult(*cmd, r)
	case "database.getCitiesById":
		res, err := cli.DatabaseGetCitiesByID(*cityIds)
		handleErr(err)
		for k, v := range res {
			r[k] = v
		}
		printResult(*cmd, r)
	case "users.isAppUser":
		res, err := cli.UserIsAppUser(*userID, *accessToken)
		handleErr(err)
		r[0] = res
		printResult(*cmd, r)
	default:
		handleErr(fmt.Errorf("undefined method"))
	}
}

func needAuth(cmd string) bool {
	switch cmd {
	case "orders.get":
		return true
	case "users.isAppUser":
		return false
	case "database.getCountries", "database.getCities", "database.getCitiesById":
		return false
	default:
		return false
	}
}

func printResult(method string, res []interface{}) {
	var buffer bytes.Buffer
	buffer.WriteString("\nResult ")
	buffer.WriteString(method)
	buffer.WriteString(" method\n")

	for _, o := range res {
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
