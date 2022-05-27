# go-metrics-app

This is a hello-world golang application with one extra feature:
- it exposes a prometheus server on port 9000 and includes go metrics as well as one custom metric.
Open `localhost:9000/` to check it.

The main.go file will continuously run incrementing one custom prometheus counter metric `hello_total_count` every 5 seconds


There are also tests for the prometheus metrics written in Ginkgo framework that can be run with:

`go test ./...`

Build Image
```
docker build . -t go-test-app
```
