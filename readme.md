# Envoy Retries and Outlier Detection

- `app/`, go program.
  1. `foo` app, `--name=foo`, ports: 8080, 80801.
  2. `bar` app, `--name=bar`, ports: 9090, 9091
- `envoy-client.yaml`, envoy client, listening on the port 7000.
  - entry point, http server, different routes.
  - `/foo`
  - `/bar`
  - `/nonexists`
  - `/gateway`.
- `envoy-gateway.yaml`, envoy gateawy, 6000.

App by itself.

/hello, /reset, /delay, /code/503.