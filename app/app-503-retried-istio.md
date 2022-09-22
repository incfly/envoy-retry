

```
curl localhost:7000/foo/code-503/retry-istio -v
curl localhost:15000/stats  | grep 'foo.*retry'
cluster.foo.circuit_breakers.default.rq_retry_open: 0
cluster.foo.circuit_breakers.high.rq_retry_open: 0
cluster.foo.retry.upstream_rq_503: 3
cluster.foo.retry.upstream_rq_5xx: 3
cluster.foo.retry.upstream_rq_completed: 3
cluster.foo.retry_or_shadow_abandoned: 0
cluster.foo.upstream_rq_retry: 3
cluster.foo.upstream_rq_retry_backoff_exponential: 3
cluster.foo.upstream_rq_retry_backoff_ratelimited: 0
cluster.foo.upstream_rq_retry_limit_exceeded: 1
cluster.foo.upstream_rq_retry_overflow: 0
cluster.foo.upstream_rq_retry_success: 0
```