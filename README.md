# Hypersomnia

Hypersomnia is a web tool to debug your [micro](https://github.com/micro/go-micro) services RPC endpoints. 
Created as a user-friendly replacement of [micro web](https://github.com/micro/micro/tree/master/web).

## Install
```
go get -u github.com/gietos/hypersomnia
```

## Usage

Default registry is mdns. If you use it, you can simply start with:
```
hypersomnia
```

To use consul registry:
```
HYPERSOMNIA_REGISTRY=consul hypersomnia
```

If default port 8083 is used by some other application, or if you don't want to bind on public interface you can use different setting of your choice:
```
HYPERSOMNIA_ADDR=127.0.0.1:31337 hypersomnia
```

## Reference values from other request's responses

In request you can reference value from other request's response by using `Response.Endpoint(<jsonpath>[,int])` 

```
{
    "someField": "Response.SomeService.SomeEndpoint($.items[0].id)" 
}
```

To insert integer instead of string:

```
{
    "someField": "Response.SomeService.SomeEndpoint($.items[0].id,int)" 
}
```

## Interact with services in the cloud

This is possible if there's available instance of micro web dashboard. To configure, pass a map with 
environment name as a key and micro web dashboard url as value:

```
HYPERSOMNIA_ENVIRONMENTS="dev:https://web-dev.example.com;stage:https://web-stage.example.com" hypersomnia 
```

You will see configured environments in environment switcher:

![](https://i.imgur.com/FF5UxOt.png)

## Development

To ship hypersomnia to the target system we have to include HTML template into binary. We use `go generate` for that, so
after any changes to `templates/index.html` don't forget to run `go generate` and commit `templates/index.go`
 