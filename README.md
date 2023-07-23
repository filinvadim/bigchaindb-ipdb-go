# bigchaindb-go
Unofficial Go BigchainDB driver.

This is an implementation of a BigchainDB client written in golang. It follows the requirements set out in [BEP-13](https://github.com/bigchaindb/BEPs/tree/master/13).

## TO-DO

* Add tests (WIP)
* Implement many-to-many transfer (WIP)
* Add documentation and more detailed example package
* Benchmark tests

## Compatibility

This driver is compatible with BigchainDB server v2.0 and above.

## Content

* [Installation](#installation)
    * [Example: Create a transaction](#example-create-a-transaction)
* [Documentation](#bigchaindb-documentation)
* [Authors](#authors)
* [License](#license)

## Installation

Run `go get -u github.com/filinvadim/bigchaindb-go`

## Example: Create a transaction

```go
        kp, _ := apiv1.NewKeyPair()
        cli, err := apiv1.NewRESTClientV1(context.Background(), "0.0.0.0:9984", kp)
        if err != nil {
            log.Fatalln(err)
        }
    
        meta := map[string]interface{}{"test": "test"}
        amount := 1.0
        _, err = cli.CreateTx(apiv1.ModeAsync, amount, meta, nil)
        if err != nil {
            log.Fatalln("create", err)
        }
```

## Example: Transfer

```go
        kp2, _ := apiv1.NewKeyPair()
        _, err = cli.TransferOneToOne(apiv1.ModeAsync, amount, kp2, previousCommittedTx, meta)
        if err != nil {
            log.Fatalln("transfer", err)
        }
```

## Example: Stream listening

```go
        stream := make(chan models.ValidTransactionResponse, 1)
        wsCli, err := apiv1.NewWSClient(context.Background(), "0.0.0.0:9985", log.Printf)
        if err != nil {
        log.Fatalln(err)
        }
        err = wsCli.SubscribeStream(stream)
        if err != nil {
        log.Fatalln(err)
        }
        for m := range stream {
            fmt.Printf("message %#v", m)
        }
```

## BigchainDB Documentation

- [HTTP API Reference](https://docs.bigchaindb.com/projects/server/en/latest/http-client-server-api.html)
- [The Transaction Model](https://docs.bigchaindb.com/projects/server/en/latest/data-models/transaction-model.html?highlight=crypto%20conditions)
- [Inputs and Outputs](https://docs.bigchaindb.com/projects/server/en/latest/data-models/inputs-outputs.html)
- [Asset Transfer](https://docs.bigchaindb.com/projects/py-driver/en/latest/usage.html#asset-transfer)
- [All BigchainDB Documentation](https://docs.bigchaindb.com/)

## Authors

* Vadim Filin
* ...

## License

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```