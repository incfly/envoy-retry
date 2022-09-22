# Envoy Retries and Outlier Detection

This repo demonstrates how Envoy retries and outlier detection behave in real
practice.

## TL;DR

- Show you how to verify the retry and outlier detection via some stats and logging.
- Istio by default retries on the application code 503, because of retriable-code: 503 is configured.
- TODO:

outlier detection confirmation
Relationship of outlier detection and retry
  retry only use the healthy ep and failed results affecting outlier detection?;
single endpoint panic mode; gateway different endpoint exp;
Istio DR config.


## Configuration

The high level structure of the repo:

- `app/`, go program, contains a sample Golang HTTP server. Some arguments
  1. `--http=8080` port specifies the HTTP server port.
  1. `--tcp=3000` specifies the TCP ports.
  1. `/code/<200|503>` allows to specify the pre-known response from the server.
- `envoy-client.yaml`, serve as entry point for the testing, listening on port 7000
- `envoy-gateawy.yaml`, serve as the second hop of the traffic flow. 

## Setup

```shell
# Start the envoy client
func-e run -c ./envoy-client.yaml -l "trace"

# Optional, start the envoy gateway.
func-e run -c ./envoy-gateway.yaml --base-id 1

# Start the application with name as foo.
go run ./main.go --id foo --http 8080 --tcp=3000

# Optional, start the application with name as bar.
go run ./main.go --id bar --http=8081
```

## Verification Info

We need to check whether retries, and outlier detection happened or not.

For retries, `curl localhost:15000/stats | grep "<cluster-name>.*upstream.*re"`.
We can also verifies the retries since every request is logged in the Golang HTTP
server.

For outlier detection, `curl localhost:15000/stats -s | grep '<cluster-name>.*outlier'`.
Moreover, the envoy bootstrap config also specify the
`cluster_manager.outlier_detection`. This makes
envoy output a JSON formatted log line when an ejection happens.
