# Relay

[![wercker status](https://app.wercker.com/status/c209eca6ce0c52f92ca6ad091fa89117/s/master "wercker status")](https://app.wercker.com/project/byKey/c209eca6ce0c52f92ca6ad091fa89117)

Relay lets you write HTTP requests as easy to read, structured YAML and dispatch them easily using a CLI. 

The motivation for this library is to have a Postman like tool for sharing HTTP reqests in a team. Relay lets you treat HTTP requests as human readable data files and store them in version control with your project.

Relay provides a CLI for executing the request only and no GUI.

## Quickstart

Download the relevant binary from [here](https://github.com/owainlewis/relay/releases). Add to your path i.e /usr/local/bin

## Examples

Relay HTTP requests are defined as YAML with a simple structure. See the examples folder for more ideas.

A request definition can contain the following

| Field   | Required | Description                                     |
|---------|----------|-------------------------------------------------|
| url     | Yes      | The full URL including protocol                 |
| method  | Yes      | HTTP method as uppercase                        |
| body    | No       | Optional HTTP request body                      |
| query   | No       | Optional key value query params appended to URL |
| headers | No       | Key value HTTP headers                          |
|         |          |                                                 |

Here are some simple examples

#### Get request

This is the most basic HTTP request with only a method and URL defined.

```yaml
description: A simple GET request
request:
  method: GET
  url: https://requestb.in/1ead0f91
```

Dispatch this request from the CLI as follows

```
relay examples/get.yml
```

#### Post request

```yaml
description: A simple POST request with body
request:
  method: POST
  url: https://requestb.in/1ead0f91
  body: 'Hello, World!'
```

#### Get request with query params and HTTP headers

```yaml
request:
  method: GET
  url: https://httpbin.org/get
  query:
    foo: bar
    baz: qux
  headers:
    Content-Type: application/json
    Authorization: Bearer {{env "AUTH_TOKEN"}}
```

#### Custom variables

Many times you will want to inject values into the templates. For example dynamic URLs like GET /foo/:id etc. 

You can use a `-params 'a=b c=d'` flag format in the CLI. Here is an example of passing in custom variables

```yaml
description: Example using injected values via params
request:
  method: GET
  url: https://api.mysite.com/users/{{.id}}
  headers:
    Content-Type: application/json
    Authorization: Bearer {{.authToken}}
```

Now we can dispatch it using the CLI

```
relay examples/dynamic.yaml -params 'id=1 authToken=XXX'
```

## Functions

A selection of functions are provided to make life easier

### Environment variables

The `env` function will extract an environment variable. 

If the environment variable is not defined then an empty string is returned.

```yaml
description: A simple request example that makes use of environment vars
request:
  method: GET
  url: https://httpbin.org/{{.method}}
  headers:
    Content-Type: application/json
    Authorization: Bearer {{env "AUTH_TOKEN"}}
```

### Basic Auth

```yaml
description: A request with HTTP Basic Auth
request:
  method: GET
  url: https://requestb.in/1ead0f91
  headers:
    Authorization: Basic {{basic "USER" "PASS"}}
```

### Basic 64 Encoding

Use `b64encode` 

## Advanced (Timeouts etc)

You can set an explicit HTTP request timeout as in the following examples:

```yaml
description: A request with explicit request timeouts
request:
  method: GET
  url: https://requestb.in/1ead0f91
options:
  timeout: 20
```

## Roadmap

+ CLI option for casting request to CURL request
+ Support for proxies
