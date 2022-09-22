# Istio Retry Defaults

This is on top of Envoy behavior to know what values Istio makes for the retries, destination rules etc.

```sh
# dev profile enabled all stats
istioc 1.13.1
ki app httpbin,sleep

# confirm starting status, retry status as zero.
kex $(kpid sleep) -- curl localhost:15000/stats | grep 'httpbin.*retry'
cluster.outbound|8000||httpbin.default.svc.cluster.local.circuit_breakers.default.rq_retry_open: 0
cluster.outbound|8000||httpbin.default.svc.cluster.local.circuit_breakers.high.rq_retry_open: 0
cluster.outbound|8000||httpbin.default.svc.cluster.local.retry_or_shadow_abandoned: 0
cluster.outbound|8000||httpbin.default.svc.cluster.local.upstream_rq_retry: 0
cluster.outbound|8000||httpbin.default.svc.cluster.local.upstream_rq_retry_backoff_exponential: 0
cluster.outbound|8000||httpbin.default.svc.cluster.local.upstream_rq_retry_backoff_ratelimited: 0
cluster.outbound|8000||httpbin.default.svc.cluster.local.upstream_rq_retry_limit_exceeded: 0
cluster.outbound|8000||httpbin.default.svc.cluster.local.upstream_rq_retry_overflow: 0
cluster.outbound|8000||httpbin.default.svc.cluster.local.upstream_rq_retry_success: 0
vhost.httpbin.default.svc.cluster.local:8000.vcluster.other.upstream_rq_retry: 0
vhost.httpbin.default.svc.cluster.local:8000.vcluster.other.upstream_rq_retry_limit_exceeded: 0
vhost.httpbin.default.svc.cluster.local:8000.vcluster.other.upstream_rq_retry_overflow: 0
vhost.httpbin.default.svc.cluster.local:8000.vcluster.other.upstream_rq_retry_success: 0
```

```sh
# confirm again
kex $(kpid sleep) -- curl 'http://httpbin:8000/status/503' -v
kex $(kpid sleep) -- curl localhost:15000/stats | grep 'httpbin.*retry'
cluster.outbound|8000||httpbin.default.svc.cluster.local.circuit_breakers.default.rq_retry_open: 0
cluster.outbound|8000||httpbin.default.svc.cluster.local.circuit_breakers.high.rq_retry_open: 0
cluster.outbound|8000||httpbin.default.svc.cluster.local.retry.upstream_rq_503: 2
cluster.outbound|8000||httpbin.default.svc.cluster.local.retry.upstream_rq_5xx: 2
cluster.outbound|8000||httpbin.default.svc.cluster.local.retry.upstream_rq_completed: 2
cluster.outbound|8000||httpbin.default.svc.cluster.local.retry_or_shadow_abandoned: 0
cluster.outbound|8000||httpbin.default.svc.cluster.local.upstream_rq_retry: 2
cluster.outbound|8000||httpbin.default.svc.cluster.local.upstream_rq_retry_backoff_exponential: 2
cluster.outbound|8000||httpbin.default.svc.cluster.local.upstream_rq_retry_backoff_ratelimited: 0
cluster.outbound|8000||httpbin.default.svc.cluster.local.upstream_rq_retry_limit_exceeded: 1
cluster.outbound|8000||httpbin.default.svc.cluster.local.upstream_rq_retry_overflow: 0
cluster.outbound|8000||httpbin.default.svc.cluster.local.upstream_rq_retry_success: 0
vhost.httpbin.default.svc.cluster.local:8000.vcluster.other.upstream_rq_retry: 0
vhost.httpbin.default.svc.cluster.local:8000.vcluster.other.upstream_rq_retry_limit_exceeded: 0
vhost.httpbin.default.svc.cluster.local:8000.vcluster.other.upstream_rq_retry_overflow: 0
vhost.httpbin.default.svc.cluster.local:8000.vcluster.other.upstream_rq_retry_success: 0
```
