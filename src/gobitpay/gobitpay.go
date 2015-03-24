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
			Name:        "new",
			Usage:       "gobitpay new [test|staging|custom <url> [--insecure]]",
			Description: "Options:\n\tlive:\tWrites the bitpay livenet uri, a pem code, and a security preference to files stored in ~/.bp/\n\ttest:\tWrites the bitpay testnet uri, a pem code, and a security preference to files stored in ~/.bp/\n\tstaging:\tWrites the bitpay staging testnet uri, a pem code, and a security preference to files stored in ~/.bp/\n\tcustom:\tWrites a custom uri, a pem code, and a security preference to files stored in ~/.bp/\n\t\tExample: gobitpay custom https://me.bp:8088 --insecure\n",
			Subcommands: []cli.Command{
				{
					Name:    "live",
					Aliases: []string{"l"},
					Usage:   "gobitpay live",
					Action: func(c *cli.Context) {
						newClient("https://bitpay.com", false)
					},
				},
				{
					Name:    "test",
					Aliases: []string{"t"},
					Usage:   "gobitpay test",
					Action: func(c *cli.Context) {
						newClient("https://test.bitpay.com", false)
					},
				},
				{
					Name:    "staging",
					Aliases: []string{"s"},
					Usage:   "gobitpay staging",
					Action: func(c *cli.Context) {
						newClient("https://staging.b-pay.net", false)
					},
				},
				{
					Name:    "custom",
					Aliases: []string{"c"},
					Usage:   "gobitpay custom <uri> [--insecure]",
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:  "insecure, i",
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
			Name:        "pair",
			Aliases:     []string{"p"},
			Usage:       "gobitpay pair <pairingCode>",
			Description: "Creates a client from values stored in ~/.bp, then retrieves a token and saves\n   it in ~/.bp/tokens.json.\n\n   Example: gobitpay pair 9zqAzwY\n",
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
			Name:        "createinvoice",
			Aliases:     []string{"ci"},
			Usage:       "gobitpay createinvoice price currency",
			Description: "create an invoice on the server using the provided price and currency.\n   Returns the invoiceid of the created invoice, which can be used by the getinvoice command\n\n   Example: gobitpay createinvoice 123.32 EUR\n",
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
			Name:        "getinvoice",
			Aliases:     []string{"gi"},
			Usage:       "gobitpay getinvoice <invoiceid>",
			Description: "retrieves an invoice from the server using the provided invoiceid\n   Returns the price and currency of the invoice.\n\n   Example: gobitpay getinvoice 1234erewRU8\n",
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
