# simple-ratelimiter
An simple ratelimit middleware.

## Simple example
You can simply try the ratelimiter in a simple sample program.

### About this simple example
This example is an simple HTTP server, which listens HTTP requests and just returns 200(OK).

If the number of requests has reached the limit, it returns 409 (Too Many Requests),

### How to build
```
make
```

### How to run
To run a sample server, execute the command below.

```
./simple-ratelimiter
```

The server will be served at localhost:3000.

If you want to limit the number of request allowed per second, you can configure it via flags.
For example, you can limit to 100 per second,

```
./simple-ratelimiter --limit 100
```
