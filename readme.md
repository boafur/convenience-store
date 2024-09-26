# Convenience store

no dependencies were used outside of `testing_utls.go` file to log to stdout and read csv data

## usage

install go and run `go run .` to install dependencies and run the program

the `-file` flag can be used to specify the file that holds the test cases

file data must be in the same format as `tests.csv`

## test case definition

test cases are defined inside of `tests.csv` should be self explanatory

the **quarters**, **dimes**, **nickels**, **pennies** fields are how many of each respective coin you have

**changeneeded** is how much change is required

**assertisenoughchange** is whether this test case has enough change to cover the required amount

---
## prompt, [from slothbytes weekly](https://slothbytes.beehiiv.com/):

Given a **total due** and **an array** representing **the amount of change in your pocket**, determine **whether or not you are able to pay for the item.**

Change will always be represented in the following order: **quarters**, **dimes**, **nickels**, **pennies**.

Example: changeEnough([25, 20, 5, 0], 4.25) return true because:

25 quarters, 20 dimes, 5 nickels and 0 pennies gives you 6.25 + 2 + .25 + 0 = 8.50.

This means you can afford the item, so return true.

### Examples (already in tests.csv):
```
changeEnough([2, 100, 0, 0], 14.11) ➞ false

changeEnough([0, 0, 20, 5], 0.75) ➞ true

changeEnough([30, 40, 20, 5], 12.55) ➞ true

changeEnough([10, 0, 0, 50], 3.85) ➞ false

changeEnough([1, 0, 5, 219], 19.99) ➞ false
```