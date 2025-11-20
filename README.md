# Packs Calculator

A Go service that calculates the optimal set of product packs needed to fulfill an order.
The goal is to:
1. Use only whole packs (no splitting).
2. Minimize overdelivery.
3. If several solutions have the same overdelivery, use the fewest packs.

The service exposes a small HTTP API and includes a simple UI for quick testing.

---

## Running

### Local (Go)
make run

### Docker Compose 
make up

UI & API available at:
http://localhost:80

---

## API

### POST /calculate
Request:
{ "amount": 251 }

 Response:
 {
   "packs": { "500": 1 }
 }

### GET /packsizes
Returns list of pack sizes.

### POST /packsizes
Updates pack sizes in memory.
 { "packSizes": [23, 31, 53] }

 ---

## Tests
make test

---

## Author
Norbert
https://github.com/nwlosinski/packsCalculator