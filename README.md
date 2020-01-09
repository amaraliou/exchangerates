# Simple Exchange Rate API

This is a coding exercise I completed where I am asked to implement a simple API capable of doing the following:

- Have an endpoint that calls ​https://exchangeratesapi.io/​ to get the latest exchange rates for the base currencies of GBP and USD;
- It should return the value of 1 GBP or 1 USD in euros;
- It should check that value against the historic rate for the last week and make a naive recommendation as to whether this is a good time to exchange money or not.

## Endpoints

Method - Endpoint | What does it do
------------------|----------------
GET /rates | Returns the latest EUR conversion rate with GBP as a base (gives the conversion of 1 GBP in EUR) and suggests whether you should buy or not.
GET /rates?base=XXX | Returns the latest EUR conversion rate with the given currency as a base (gives the conversion of 1 XXX in EUR) and suggests whether you should buy or not. For this exercise, it supports only GBP and USD as base.

## How to run

- Clone this repo with `git clone https://github.com/amaraliou/exchangerates` and `cd` into it;
- Run `go mod download` on the terminal;
- Run `make run` on the terminal;
- Go to `localhost:8080` to test the endpoints.

## How to deploy

### Prerequisites

- Heroku CLI
- Docker

### Instructions

- Clone this repo with `git clone https://github.com/amaraliou/exchangerates` and `cd` into it;
- Login into Heroku by running `heroku login`;
- Login into the Heroku container registry by running `heroku container:login`;
- Create a Heroku instance by running `heroku create name-of-your-app`;
- Push your app to the container by running `heroku container:push -a name-of-your-app web`;
- Release your app by running `heroku container:release -a name-of-your-app web`;
- Now go to `https://name-of-your-app.herokuapp.com` to test the endpoints (you can try `https://curve-exchangerates.herokuapp.com` which is live).

## Tests

Run `make test`