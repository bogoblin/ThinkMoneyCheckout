# Shopping Cart Calculator

## How to run

This program accepts SKUs from standard input. Each line should contain a single SKU.
On end of file, the program will print the total price of the cart.

First, build the executable with:

```bash
go build
```

I've included a sample input file that you can test like so:

```bash
cat example.txt | ./ThinkMoneyCheckout
```

Alternatively you can type in items interactively, then press CTRL+D to send an EOF.

## Running tests

```bash
go test cart/cart_test.go
```