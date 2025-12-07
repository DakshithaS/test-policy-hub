module example-consumer

go 1.25.3

require example.com/policies/rate-limiter v0.0.0

replace example.com/policies/rate-limiter => ../policy-catalog/policies/rate-limiter/v1.0.0/src
