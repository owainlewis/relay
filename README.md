# Relay

This library lets you write HTTP requests as JSON data and send them easily using a CLI.

## Why?

I love tools like Postman but I wanted an easier way to distribute HTTP request around
a team. Relay lets you treat HTTP requests for testing etc as pure data and store them
in version control with your project.

## Example

Let's define a HTTP request as pure JSON.

```json
{ "method": "POST",
  "url": "http://localhost:8080/protect",
  "headers": { "Content-Type": "application/json" },
  "body": [["owain@owainlewis.com", "email"]]
}

```

Now we can dispatch it using the CLI

```
relay request.json
```
