# Relay

This library lets you write HTTP requests as structured Yaml and send them easily using a CLI.

## Why?

I love tools like Postman but I wanted an easier way to distribute HTTP request around
a team. Relay lets you treat HTTP requests for testing etc as pure data and store them
in version control with your project.

## Examples

Relay HTTP requests are defined as YAML with a simple structure.

A request definition can contain the following

* url 
* method 
* query
* headers

We can give a name to our requests to provide a human readable reference.

```yaml
description: A simple request example
request:
  method: GET
  url: https://httpbin.org/{{.method}}
  query:
    foo: bar
    baz: qux
  headers:
    Content-Type: application/json
    Authorization: Bearer {{env "AUTH_TOKEN"}}
```

Now we can dispatch it using the CLI

```
relay examples/get.yaml -params 'method=get'
```

### Special functions

#### Accessing environment variables

#### Accessing params

### Todo

+ CLI option for casting request to CURL request
+ Pass options like proxy etc

## Building from source

```
git clone git@github.com:owainlewis/relay.git && cd relay
go build
mv relay /usr/local/bin/

# Execute a static request file

relay examples/get.yaml
```
