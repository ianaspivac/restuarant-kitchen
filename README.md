# Restuarant Kitchen

`Dinning-hall` and `Kitchen` are two servers communicating through HTTP requests.
`Dinning-hall` communicates through `waiters`, which recieve and send `<orders>` to `<kitchen>`.
`Dinning-hall` contains `tables` which generate and recieve `<orders>`.
`Kitchen` contains `cooks` which prepare the `orders, same as `Waiters` , these 2 groups run on multiple threads.

### Clone the repo
```shell
$ git clone https://github.com/ianaspivac/restuarant-kitchen
```

### Install Go 1.17 (or at least 1.15)
[Go Install Guide](https://golang.org/doc/install)

### Install the dependencies
```shell
$ go mod download
```
## Start the kitchen

### Simple start with default config path
```shell
go run main.go
```

### Build the project
```shell
go build -o <name>
```


### Hosting



**Address of kitchen port:**<br>
```json
"localhost:8081"
```


Address of dining hall

**Address of dining hall port:**<br>
```json
"localhost:8080"
```
