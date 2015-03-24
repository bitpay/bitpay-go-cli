package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "gobitpay"
	app.Usage = "gobitpay <command> [opts]"
	app.Commands = []cli.Command{
		{
			Name:  "new",
			Usage: "bitpaygo new --[test|staging|custom <url> [--insecure]]",
			Subcommands: []cli.Command{
				{
					Name:    "live",
					Aliases: []string{"l"},
					Usage:   "sets up a live, aye",
					Action: func(c *cli.Context) {
						newClient("https://bitpay.com", false)
					},
				},
				{
					Name:  "test",
					Usage: "sets up a test, aye",
					Action: func(c *cli.Context) {
						newClient("https://test.bitpay.com", false)
					},
				},
				{
					Name:  "staging",
					Usage: "sets up a staging server, aye",
					Action: func(c *cli.Context) {
						newClient("https://staging.b-pay.net", false)
					},
				},
				{
					Name:  "custom",
					Usage: "sets up a custom server: ",
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:  "insecure",
							Usage: "new client will ignore ssl errors",
						},
					},
					Action: func(c *cli.Context) {
						apiuri := ""
						var insecure bool
						insecure = c.Bool("insecure")
						if len(c.Args()) > 0 {
							apiuri = c.Args()[0]
							newClient(apiuri, insecure)

						} else {
							println("custom server requires a server address")
						}
					},
				},
			},
		},
		{
			Name:  "pair",
			Usage: "bitpaygo pair",
			Action: func(c *cli.Context) {
				if len(c.Args()) > 0 {
					code := c.Args()[0]
					_, err := pairClient(code)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					println("Must provide a pairing code")
				}
			},
		},
		{
			Name:  "createinvoice",
			Usage: "bitpaygo createinvoice price currency <parameters>",
			Action: func(c *cli.Context) {
				if len(c.Args()) > 1 {
					price := c.Args()[0]
					currency := c.Args()[1]
					invId, err := createInvoice(price, currency)
					if err != nil {
						fmt.Println(err)
					} else {
						println(invId)
					}
				} else {
					println("Must provide a price and a currency")
				}
			},
		},
		{
			Name:  "getinvoice",
			Usage: "bitpaygo getinvoice <invoiceid>",
			Action: func(c *cli.Context) {
				if len(c.Args()) > 0 {
					invId := c.Args()[0]
					price, currency, err := getInvoice(invId)
					if err != nil {
						fmt.Println(err)
					} else {
						priceSt := fmt.Sprintf("%f", price)
						fmt.Println("Price: " + priceSt + ", Currency: " + currency)
					}
				} else {
					println("Must provide an invoice id")
				}
			},
		},
	}

	app.Run(os.Args)
}
