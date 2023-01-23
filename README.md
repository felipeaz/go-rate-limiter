# go-rate-limiter
Rate Limiter is used in situations where it's required to limit the number of requests a service is receiving.

In this example, I created a simple endpoint that handles a request for 10 seconds, let's say that our service takes that
amount of time to process a single request. Each request will have an ID defined on the requestID middleware, that ID
will be used to identify which request got processed at the end.

The rate limiter will be 3 request per second, if the 4th request is sent, it'll be denied
