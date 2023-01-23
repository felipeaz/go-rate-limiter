# go-rate-limiter
Rate Limiter is used in situations where it's required to limit the number of requests a service is receiving.

In this example, I created a simple endpoint that handles a request for 10 seconds, let's say that our service takes that
amount of time to process a single request. Each request will have an ID defined on the requestID middleware, that ID
will be used to identify which request got processed at the end.

The rate limiter will be 3 request per second, if the 4th request is sent, it'll be denied

## Testing
Go to the project root, execute the main go file then run `make deny` to simulate few denies or `make run` for a single request

![ezgif com-gif-maker](https://user-images.githubusercontent.com/32846823/214162095-8d2b34f6-ce55-4e95-b6ef-b559ef29753f.gif)
