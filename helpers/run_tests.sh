#!bin/bash
python helpers/pair_steps.py
ginkgo -r --keepGoing src/github.com/bitpay/bitpay-go
