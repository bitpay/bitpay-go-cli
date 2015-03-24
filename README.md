# BitPay Command Line Tool for Go

This is an example library for integration of the bitpay-go library with a go project. It provides a command line tool that can authenticate with a bitpay server, create invoices, and retrieve invoices.

## Installation and Configuration

Clone this repository and run `source helpers/envioro.sh`

## Basic Usage

The gobitpay tool allows authenticating with BitPay, creating invoices, and retrieving invoices.
  
### Pairing with Bitpay.com

Before pairing with BitPay.com, you'll need to log in to your BitPay account and navigate to /api-tokens. Generate a new pairing code and use it in the next step.

First you will want to create a client. Creating a client will save some variables in the ~/.bp/ folder. Only one client can be stored at one time. The stored values are a PEM file, the intended api endpoint, and the ssl security preference. To use the bitpay.com test server (test.bitpay.com) the command is:

    $ gobitpay new test

Once you have created the client preferences, log in to the endpoint you specificied (in this case test.bitpay.com) and navigate to `dashboard/merchant/api-tokens`. Create a new api token and copy the pairing code displayed. You will use this in the next step.

    $ gobitpay pair <pairingcode> 

This will save a token in the ~/.bp/folder. 

### Creating and retrieving invoices

If you have completed the previous step, you can create and retrieve invoices. 

    $ gobitpay createinvoice <price> <currency>
    
This returns an invoice id, which can be used in the next step:

    $ gobitpay getinvoice <invoiceid>

