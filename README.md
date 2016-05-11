# Relay

This library lets you write HTTP requests as JSON data and send them easily using a CLI.

## TODO

Collections (relay run mycollection.json)

## Why?

I love tools like Postman but I wanted an easier way to distribute HTTP request around
a team. Relay lets you treat HTTP requests for testing etc as pure data and store them
in version control with your project.

## Building from source

```
git clone git@github.com:owainlewis/relay.git && cd relay
go build

# Execute a static request file

./relay run examples/get.json
```

## Examples

Replay HTTP requests are defined as pure JSON with a simple structure.

We can give a name to our requests to provide a human readable reference.

```json
{   "name": "Example GET request",
    "request": {
      "method": "GET",
      "url": "https://httpbin.org/get",
      "headers": {
          "Content-Type": "application/json",
          "X-Foo": "bar"
      }
    }
}
```

Now we can dispatch it using the CLI

```
relay request.json
```

Here are some more examples

```json
{   "name": "Example POST request",
    "request": {
      "method": "POST",
      "url": "https://httpbin.org/post",
      "headers": {
          "Content-Type": "application/json",
      },
      "body": { "Hello" : "World" }
    }
}
```

## In depth

The following fields are supported when defining relay HTTP requests

Method  : A HTTP request method (GET, POST, DELETE, PATCH etc)
Url     : The full URL to a request i.e http://google.com?foo=bar
Headers : A map of HTTP headers
Body    : An optional request body. Note that you can define the request body as pure JSON
