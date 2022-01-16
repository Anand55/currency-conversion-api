# currency-conversion-api

This is a currency converter app, which will calculate the amount given from one currency to another.

## Prerequisites

Need to have docker and docker compose installed on your system. If not, you can install it using below doc:

```bash
Docker https://docs.docker.com/engine/install/
Docker compose https://docs.docker.com/compose/install/
```

If you wish to run app locally, you will need to have golang installed on your system. If not installed, follow the steps in below doc to install it.

```bash
https://golang.org/doc/install
```

## Getting Started

Clone the git repository in your system and then cd into project root directory

```bash
$ git clone git@github.com:Anand55/currency-conversion-api.git
$ cd currency-conversion-api
```

### OR

Unzip the `currency-conversion-api.zip` file.

## 1. Steps to run via docker-compose

```bash
$ cd currency-conversion-api
```

After this, you can run the following command

```bash
$ docker-compose up
```

This should start the service at `localhost:8080`.

Now, go to postman and hit **_create_** endpont URL. This endpoint takes email as an body.

```bash
http://localhost:8080/create
```

with body:

```bash
{
    "email":"passbase@gmail.com"
}
```

This will return with a response which contains API access key. See below,

```bash
{
  "accesskey": "6c9a0c7f25d59bacadabcb61add62ed31312fa3c"
}
```

This key will be used to access the **_convert_** endpoint.

Now, using access key from above response, you can hit convert endpoint as below

```bash
http://localhost:8080/convert?key=6c9a0c7f25d59bacadabcb61add62ed31312fa3c
```

with body:

```bash
{
    "from": "EUR",
    "to": "USD",
    "amount": 1
}
```

You will get the response as below.

```bash
{
  "from": "EUR",
  "to": "USD",
  "amount": 1,
  "result": "1.000000 in EUR is equals to 1.141611 in USD"
}
```

This service is not just limited to EUR <-> USD and vice versa.

For example, If you request with below body.

```bash
{
    "from": "EUR",
    "to": "INR",
    "amount": 1
}
```

You will get response

```bash
{
  "from": "EUR",
  "to": "INR",
  "amount": 1,
  "result": "1.000000 in EUR is equals to 84.920459 in INR"
}
```

**_Note: You cannot access convert endpoint without access key._**

Convert endpoint offers 15 hits per user. After that you will see limit exceeded message.

## 2. Steps to run locally.

Make sure you have golang installed.

```bash
$ cd currency-conversion-api
```

You need to expose two environment variables. So run

```bash
$ export REDIS_ADDR=localhost:6379
$ export FIXER_KEY=d244f7d57ffda45bc7b3b39e1ae75d0d
```

**_Note: I have created the fixer account for testing purposes so you can use above fixer key instead of creating your own._**

Also, you need to start redis locally. You can start it using docker.

```bash
$ docker run -p 6379:6379 redis
```

Now, you can run

```bash
$ go run cmd/main.go
```

This will start the service. Follow the same steps as mentioned above for hitting endpoints.

## Design

I have used simple directory structure.

```
cmd: which contains config package and main file.
```

```
db : which contains redis package.
```

```
domain : which contains business logic, convert package.
```

```
handler : which contains endpoint handlers.
```

```
middleware : which contains middlewares. Auth middleware is implemented in this service.
```

```
routes : where you can register routes.
```

I am using redis to store access key and number of time user has hit the convert endpoint.

This helps me understand,

1. If user is registered or not.
2. If access key is valid or not.
3. If user has exceeded the number of hits.

I have unit tested the business logic.
