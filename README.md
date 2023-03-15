# Settler

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/kevinrobayna/settler)
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/kevinrobayna/Settler/test.yml)
![GitHub](https://img.shields.io/github/license/kevinrobayna/settler)

Library to calculate debt when a group wants to split expenses

## Caveat

This is not a real library this is something I created to understand how splitting expenses works, since my friends and
I use SettleUp when we do a trip to the snow!.

## Concepts

To understand how this library works we need first to understand a basic concept of accounting, which is the
[double-entry bookkeeping system](https://en.wikipedia.org/wiki/Double-entry_bookkeeping).

> Double-entry bookkeeping is a method of bookkeeping that relies on two-sided accounting entries to maintain financial
> information. Every entry to an account requires a corresponding and opposite entry to a different account.

For example if you paid someone 100 then if you kept track of this in a [ledger](https://en.wikipedia.org/wiki/Ledger)

|          | Debit | Credit |
|----------|-------|--------|
| Cash     | 100   |        |
| Transfer |       | 100    |

At any point of the bookkeeping you can trust that if you sum all the debits minus the credits the result should be 0.

## Expense Splitting

Knowing what a double-entry bookkeeping system is we can now understand how to split expenses.

Let's assume we have the group had the following expenses:

* Person A paid 100 for the hotel room for person A and B
* Person C paid 60 for the lunch for person A, B and C
* Person B paid 300 for the rental of the snow equipment for person A, B and C

The "ledger" table would look like this:

|                | A   | B   | C   |
|----------------|-----|-----|-----|
| Hotel          | 50  | 50  | 0   |
| Lunch          | 20  | 20  | 20  |
| Snow Equipment | 100 | 100 | 100 |

The keep eyes around you have noticed that we are missing the other side of the system. For this I've introduced
signs (`+` and `-`) to represents credits and debits.

|                    | A    | B    | C   |
|--------------------|------|------|-----|
| Hotel (C)          | -100 | 0    | 0   |
| Hotel              | 50   | 50   | 0   |
| Lunch (C)          | 0    | 0    | -60 |
| Lunch              | 20   | 20   | 20  |
| Snow Equipment (C) | 0    | -300 | 0   |
| Snow Equipment     | 100  | 100  | 100 |

Now if we add everything, columns and rows, we get the 0 we were expecting. Let's now add all the rows together and see
what people owe to each other.

|       | A   | B    | C   |
|-------|-----|------|-----|
| Debts | 70  | -130 | 60  |

So this means that person B is owed 130 from A and C. So person A and C would need to transfer in total 130 and the
table would look like this:

|            | A   | B    | C   |
|------------|-----|------|-----|
| Debts      | 70  | -130 | 60  |
| Transfer A | -70 | 70   | 0   |
| Transfer C | 0   | 60   | -60 |

Here we see that Person A is Debited 70 which means person B is credited 70 and person C is debited 60. Which means
that everything comes back to a balance position.

