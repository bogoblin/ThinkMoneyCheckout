# Think Money Task

Please develop a solution to the Supermarket Checkout problem detailed here:

------------

Implement the code for a checkout system that handles pricing schemes such as "pineapples cost 50, three pineapples cost 130."

Implement the code for a supermarket checkout that calculates the total price of a number of items. In a normal supermarket, things are identified using Stock Keeping Units, or SKUs. In our store, we’ll use individual letters of the alphabet (A, B, C, and so on). Our goods are priced individually. In addition, some items are multi-priced: buy n of them, and they’ll cost you y pence. For example, item A might cost 50 individually, but this week we have a special offer: buy three As and they’ll cost you 130. In fact the prices are:

| SKU | Unit Price | Special Price |
|-----|------------|---------------|
| A   | 50         | 3 for 130     |
| B   | 30         | 2 for 45      |
| C   | 20         |               |
| D   | 15         |               |

The checkout accepts items in any order, so that if we scan a B, an A, and another B, we’ll recognize the two Bs and price them at 45 (for a total price so far of 95). The implementation should consider if the pricing model may change frequently.

------------

We use this technical challenge as part of the recruitment process to give you the opportunity to demonstrate your skills and ability. We will review the submission to decide if we will proceed to the next stage of the process. In the interview we review your code with you, either as a discussion exercise, or often as a pairing exercise where we will ask you to extend your code with additional requirements.

Please submit your solution in Go Lang.

When submitting your kata we are specifically looking for:

- Test driven development (TDD)
- Small "baby" steps
- Frequent commits to a repository on GitHub (so we can see how you got to the solution, considering the first 2 points above)
- A README file if your solution has any specific setup instructions