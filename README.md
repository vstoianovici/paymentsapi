[![GoDoc](https://godoc.org/github.com/vstoianovici/wservice?status.svg)](https://godoc.org/github.com/vstoianovici/wservice) [![Go Report Card](https://goreportcard.com/badge/github.com/vstoianovici/wservice)](https://goreportcard.com/report/github.com/vstoianovici/wservice) [![Build Status](https://travis-ci.org/vstoianovici/wservice.svg?branch=master)](https://travis-ci.org/vstoianovici/wservice)
# PaymentsAPI

`PaymentsAPI` provides a simple payments RESTful API based on HTTP server implemented in Go (and [gokit.io](https://gokit.io)) that deals with Json format files to provide CRUD functionality against a Postgresql database.
I used Gokit to help in separating implementation concerns by employing an onion layered model, where at the very core we have our use cases or bussines domain (source code dependencies can only point inward) and then wrapping that with other functionality layers.

Here are the main requirements on whcih the design is based:

- Fetching a single payment resources
- Creating, updating and deleting a payment resouce
- Listing a collection of payment resources (a list of payments might look like http://mockbin.org/bin/41ca3269-d8c4-4063-9fd5-f306814ff03f.)
- Persisting resource state (to a database)

## Design

The prototype is divided into 4 modules:

1. The HTTP server and router module.
This is where the HTTP server is launched and managed and the rounting for the various endpoints is provided.

2. The gokit wrapper module
Gokit is used in order to implement a decorator pattern where concerns such as logging, monitoring, transport, circuit breaking are sepparated and do not introduce any dependencies in the core functioanlity code.

3. The core Payments Service functionality
This is where the core logic resides. The CRUD functionality and database model are managed and implemented at this level.

4. The Postgres database layer that employs Gorm to talk to Postgres
Gorm is used to facilitate the interaction with the Postgresql database. Since this is a fin-tech application where a delicate balance between data integrity and reliability on one hand and high performance and scalibility on the other, needs to be achieved the design decision was to emplpoy a typical sql-type database (for structured data) such as Postgres as it also guarantees ACID operations.

A few opensource frameworks and libraries were used in the implementation of this project:
 - Gokit - https://github.com/go-kit/kit - separation of concerns for designing microservices
 - Gorm - https://github.com/jinzhu/gorm - layer that facilitates interaction with the DB from Go
 - Gorilla/Mux - github.com/gorilla/mux" - for http routing
 - Viper - "github.com/spf13/viper" - for reading configuration files
 - A few librabries dedicated to testing scenarios such as Assert, Mock and GoMocket: "github.com/stretchr/testify/assert", "github.com/stretchr/testify/mock" and "github.com/Selvatico/go-mocket"
 
 This design did not address implementing a client to run against the API. In development cUrl commands were used to run against the server as can be seen below:

$ curl "http://localhost:8080/v1/payments/"
[]

$ curl -X POST --data-binary @payment1.json "http://localhost:8080/v1/payments/"
{"created_id":"2e1f6c5d-3965-489e-a156-6f0e7d482c9e"}

$ curl "http://localhost:8080/v1/payments/"
[{"id":"2e1f6c5d-3965-489e-a156-6f0e7d482c9e","type":"Payment","version":0,"organisation_id":"743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb","attributes":{"amount":"130.21","beneficiary_party":{"account_number":"31926819","bank_id":"403000","bank_id_code":"GBDSC","account_name":"W Owens","account_number_code":"BBAN","address":"1 The Beneficiary Localtown SE2","name":"Wilfred Jeremiah Owens","account_type":0},"charges_information":{"bearer_code":"SHAR","sender_charges":[{"amount":"5.00","currency":"GBP"},{"amount":"10.00","currency":"USD"}],"receiver_charges_amount":"1.00","receiver_charges_currency":"USD"},"currency":"GBP","debtor_party":{"account_number":"GB29XABC10161234567801","bank_id":"203301","bank_id_code":"GBDSC","account_name":"EJ Brown Black","account_number_code":"IBAN","address":"10 Debtor Crescent Sourcetown NE1","name":"Emelia Jane Brown"},"end_to_end_reference":"Wil piano Jan","fx":{"contract_reference":"FX123","exchange_rate":"2.00000","original_amount":"200.42","original_currency":"USD"},"numeric_reference":"1002001","payment_id":"123456789012345678","payment_purpose":"Paying for goods/services","payment_scheme":"FPS","payment_type":"Credit","processing_date":"2017-01-18","reference":"Payment for Em's piano lessons","scheme_payment_sub_type":"InternetBanking","scheme_payment_type":"ImmediatePayment","sponsor_party":{"account_number":"56781234","bank_id":"123123","bank_id_code":"GBDSC"}}}]

$ curl -X PUT --data-binary @payment0.json "http://localhost:8080/v1/payments/2e1f6c5d-3965-489e-a156-6f0e7d482c9e" 
{"updated_id":"2e1f6c5d-3965-489e-a156-6f0e7d482c9e"}

$ curl "http://localhost:8080/v1/payments/"
[{"id":"2e1f6c5d-3965-489e-a156-6f0e7d482c9e","type":"Payment","version":0,"organisation_id":"743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb","attributes":{"amount":"100.21","beneficiary_party":{"account_number":"31926819","bank_id":"403000","bank_id_code":"GBDSC","account_name":"W Owens","account_number_code":"BBAN","address":"1 The Beneficiary Localtown SE2","name":"Wilfred Jeremiah Owens","account_type":0},"charges_information":{"bearer_code":"SHAR","sender_charges":[{"amount":"5.00","currency":"GBP"},{"amount":"10.00","currency":"USD"}],"receiver_charges_amount":"1.00","receiver_charges_currency":"USD"},"currency":"GBP","debtor_party":{"account_number":"GB29XABC10161234567801","bank_id":"203301","bank_id_code":"GBDSC","account_name":"EJ Brown Black","account_number_code":"IBAN","address":"10 Debtor Crescent Sourcetown NE1","name":"Emelia Jane Brown"},"end_to_end_reference":"Wil piano Jan","fx":{"contract_reference":"FX123","exchange_rate":"2.00000","original_amount":"200.42","original_currency":"USD"},"numeric_reference":"1002001","payment_id":"123456789012345678","payment_purpose":"Paying for goods/services","payment_scheme":"FPS","payment_type":"Credit","processing_date":"2017-01-18","reference":"Payment for Em's piano lessons","scheme_payment_sub_type":"InternetBanking","scheme_payment_type":"ImmediatePayment","sponsor_party":{"account_number":"56781234","bank_id":"123123","bank_id_code":"GBDSC"}}}]

$ curl -X POST --data-binary @payment1.json "http://localhost:8080/v1/payments/"{"created_id":"d0f2bc35-7778-4e0a-a285-0618545c438f"}

$ curl "http://localhost:8080/v1/payments/"[{"id":"2e1f6c5d-3965-489e-a156-6f0e7d482c9e","type":"Payment","version":0,"organisation_id":"743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb","attributes":{"amount":"100.21","beneficiary_party":{"account_number":"31926819","bank_id":"403000","bank_id_code":"GBDSC","account_name":"W Owens","account_number_code":"BBAN","address":"1 The Beneficiary Localtown SE2","name":"Wilfred Jeremiah Owens","account_type":0},"charges_information":{"bearer_code":"SHAR","sender_charges":[{"amount":"5.00","currency":"GBP"},{"amount":"10.00","currency":"USD"}],"receiver_charges_amount":"1.00","receiver_charges_currency":"USD"},"currency":"GBP","debtor_party":{"account_number":"GB29XABC10161234567801","bank_id":"203301","bank_id_code":"GBDSC","account_name":"EJ Brown Black","account_number_code":"IBAN","address":"10 Debtor Crescent Sourcetown NE1","name":"Emelia Jane Brown"},"end_to_end_reference":"Wil piano Jan","fx":{"contract_reference":"FX123","exchange_rate":"2.00000","original_amount":"200.42","original_currency":"USD"},"numeric_reference":"1002001","payment_id":"123456789012345678","payment_purpose":"Paying for goods/services","payment_scheme":"FPS","payment_type":"Credit","processing_date":"2017-01-18","reference":"Payment for Em's piano lessons","scheme_payment_sub_type":"InternetBanking","scheme_payment_type":"ImmediatePayment","sponsor_party":{"account_number":"56781234","bank_id":"123123","bank_id_code":"GBDSC"}}},{"id":"d0f2bc35-7778-4e0a-a285-0618545c438f","type":"Payment","version":0,"organisation_id":"743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb","attributes":{"amount":"130.21","beneficiary_party":{"account_number":"31926819","bank_id":"403000","bank_id_code":"GBDSC","account_name":"W Owens","account_number_code":"BBAN","address":"1 The Beneficiary Localtown SE2","name":"Wilfred Jeremiah Owens","account_type":0},"charges_information":{"bearer_code":"SHAR","sender_charges":[{"amount":"5.00","currency":"GBP"},{"amount":"10.00","currency":"USD"}],"receiver_charges_amount":"1.00","receiver_charges_currency":"USD"},"currency":"GBP","debtor_party":{"account_number":"GB29XABC10161234567801","bank_id":"203301","bank_id_code":"GBDSC","account_name":"EJ Brown Black","account_number_code":"IBAN","address":"10 Debtor Crescent Sourcetown NE1","name":"Emelia Jane Brown"},"end_to_end_reference":"Wil piano Jan","fx":{"contract_reference":"FX123","exchange_rate":"2.00000","original_amount":"200.42","original_currency":"USD"},"numeric_reference":"1002001","payment_id":"123456789012345678","payment_purpose":"Paying for goods/services","payment_scheme":"FPS","payment_type":"Credit","processing_date":"2017-01-18","reference":"Payment for Em's piano lessons","scheme_payment_sub_type":"InternetBanking","scheme_payment_type":"ImmediatePayment","sponsor_party":{"account_number":"56781234","bank_id":"123123","bank_id_code":"GBDSC"}}}]

$ curl -X DELETE "http://localhost:8080/v1/payments/d0f2bc35-7778-4e0a-a285-0618545c438f"{"DeletedAt":"2019-04-22T11:45:26.089166Z"}

$ curl "http://localhost:8080/v1/payments/"
[{"id":"2e1f6c5d-3965-489e-a156-6f0e7d482c9e","type":"Payment","version":0,"organisation_id":"743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb","attributes":{"amount":"100.21","beneficiary_party":{"account_number":"31926819","bank_id":"403000","bank_id_code":"GBDSC","account_name":"W Owens","account_number_code":"BBAN","address":"1 The Beneficiary Localtown SE2","name":"Wilfred Jeremiah Owens","account_type":0},"charges_information":{"bearer_code":"SHAR","sender_charges":[{"amount":"5.00","currency":"GBP"},{"amount":"10.00","currency":"USD"}],"receiver_charges_amount":"1.00","receiver_charges_currency":"USD"},"currency":"GBP","debtor_party":{"account_number":"GB29XABC10161234567801","bank_id":"203301","bank_id_code":"GBDSC","account_name":"EJ Brown Black","account_number_code":"IBAN","address":"10 Debtor Crescent Sourcetown NE1","name":"Emelia Jane Brown"},"end_to_end_reference":"Wil piano Jan","fx":{"contract_reference":"FX123","exchange_rate":"2.00000","original_amount":"200.42","original_currency":"USD"},"numeric_reference":"1002001","payment_id":"123456789012345678","payment_purpose":"Paying for goods/services","payment_scheme":"FPS","payment_type":"Credit","processing_date":"2017-01-18","reference":"Payment for Em's piano lessons","scheme_payment_sub_type":"InternetBanking","scheme_payment_type":"ImmediatePayment","sponsor_party":{"account_number":"56781234","bank_id":"123123","bank_id_code":"GBDSC"}}}]

## Get started with docker


Get the source code:

```
$ go get -u github.com/vstoianovici/paymentsapi
```

Build the environment from the `docker-compose.yml` file in the root (`gowebapp` and `postgresdb` will be deployed):

```
$ docker-compose up -d
```
The `postgresdb` will already have a database called `Postgres` that has the `Accounts` and `Transfers` tables. The `Account` table will look something like this:

<img width="505" alt="Screenshot 2019-03-22 at 22 53 23" src="https://user-images.githubusercontent.com/26381671/54855623-c959fc80-4cff-11e9-8b92-c0b507c8bc18.png">

while the `Transfers` table will be empty:

<img width="755" alt="Screenshot 2019-03-22 at 22 54 17" src="https://user-images.githubusercontent.com/26381671/54855611-bcd5a400-4cff-11e9-9dd4-a7f8438ff2c1.png">

One can visualize both tables by accessing the follwing links:

- The `Account` table: http://127.0.0.1:8080/accounts

- The `Transfers` table: http://127.0.0.1:8080/transfers

- Additionally one can see the `Metrics & Instrumentation` endpoint here: http://127.0.0.1:8080/metrics

At this point one could proceed to the Runtime section from below and focus on the sections about the curl commands that allow you to interact with the service (ignore the part where `wservice` is launched, as that is only for those who don't use docker to deploy the service)


## Get started with building the Go binary and deploying a postgres DB

Get the source code:

```
$ go get -u github.com/vstoianovici/wservice
```

Build the binary by running the following command in the root:

```
$ make build
```

If the build succeeds, the resulting binary named "wService" should be found in the `/cmd` directory.

To build the Postgres db as a Docker container run these 2 commands:

```
docker build -t postgresdb -f ./Dockerfile_postgres .
```
followed by

```
docker run --rm --name postgresdb -e POSTGRES_PASSWORD=password -d -p 5432:5432 postgresdb
```
In case Postgres is installed in any other way other than the ones described above the user needs to create a database named `Postgres` and run the SQL queries contained in the `init-user-db.sh script` from the `/docker` folder to create the tables `Accounts` and `Transfers` as portrayed in the screenshots from the section "Get strated with Docker" from above.


Addtionally there is a `postgresql.cfg` file contained in `/cmd` that is used for configuring the connection between the Go webapp and the Postgres db. The content is pretty self-explanatory:

```yaml
sqlDriver : postgres,
sqlHost : 127.0.0.1,
sqlPort : 5432,
sqlUser : postgres,
sqlPassword : password,
sqlDbName : postgres,
sslmode : disable,
accountsTable : Accounts,
transfersTable : Transfers
```

Run the tests:

```
$ make test
```

Feel free to explore the `Makefile` available in the root directory.

### Runtime

- Otherwise, once `wService` is built and ready for runtime it can run (/cmd/wService) without any parameters (default should be fine) but there is the option of passing in a different port or a different `postgres.cfg` file (skip this step, if you are deploying with docker-compose, and continue to the curl commands bellow):

```
$ ./wService -h
time=2019-03-22T20:40:18.099917Z tag=start msg="created logger"
Usage of ./wService:
  -file string
        Path of postgresql config file to be parsed. (default "./postgresql.cfg")
  -port int
        Port on which the server will listen and serve. (default 8080)
```

- At runtime the easiest way to actually create fund transfers is to run a curl command against the `submittransfer` API endpoint such as the following:

```
curl  -d'{"from":"bob123","to":"alice456","amount":"20"}' "127.0.0.1:8080/submittransfer"
```

- The other touchpoints of the API to visualize the accounts' balance (`/accounts`,), the already submitted transactions (`/transfers`) and the metrics data (`/metrics`) are, as mentioned earlier reachable with:
```
curl "127.0.0.1:8080/transfers"
```
```
curl "127.0.0.1:8080/accounts"
```
```
curl "127.0.0.1:8080/metics"
```


### Build your own wallet

Anybody can use this resource as a library to create their own implementation of a micro Wallet Service as long as they mimic what is being done in `/cmd/main.go`

For the future, a nice feature to implement would be a gRPC endpoint in addition to the Json over HTTP REST API so that the wallet can be a service in a microservice architecture solution.

### Contribute

Contributions to this project are welcome, though please file an issue before starting work on anything major.
The next step in the evolution of this product would be a gRPC Transport wrapper to allow for optimum inter-process commuinication

### License

The MIT License (MIT) - see the LICENSE file for more details
