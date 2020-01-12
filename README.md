# simple-ratelimiter
A simple ratelimit middleware.

## Simple example
You can simply try the ratelimiter in a simple sample program.

### About this simple example
This example is a simple HTTP server, which listens to HTTP requests and just returns 200(OK).

If the number of requests has reached the limit, it returns 409 (Too Many Requests),

### How to build
```
cd examples
make
```

### How to run
To run a sample server, execute the command below.

```
cd examples
./simple-server
```

The server will be served at localhost:3000.

If you want to limit the number of requests allowed per second, you can configure it via flags.

For example, you can limit to 100 per second,

```
cd examples
./simple-server --limit 100
```
