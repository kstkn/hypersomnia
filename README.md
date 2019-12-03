# Hypersomnia

Hypersomnia is a web tool to debug your [micro](https://github.com/micro/go-micro) services RPC endpoints. 
Created as a user-friendly replacement of [micro web](https://github.com/micro/micro/tree/master/web).

![](https://i.imgur.com/VJHYUx6.png)

## Install
```
go get -u github.com/gietos/hypersomnia
```

## Usage

Default registry is consul on localhost:8500. If you use it, you can simply start with:
```
hypersomnia
```

To use mdns registry:
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

This is possible if there's available instance of micro web. To configure, pass a map with 
environment name as a key and micro web url as value:

```
HYPERSOMNIA_ENVIRONMENTS="dev:https://web-dev.example.com;stage:https://web-stage.example.com" hypersomnia 
```

You will see configured environments in environment switcher:

![](https://i.imgur.com/642Ycdq.png)

## Quick access to external system to search by correlation id

For each environment you can set up an URL template to open quickly by pressing correlation id badge.
Special placeholder `_correlationId_` will be replaced and URL will open.
It can be log collector, like kibana. To set up URLs, press environment settings button next to environment selector:

![](https://i.imgur.com/yQSRj4O.png)

Settings popup will open:

![](https://i.imgur.com/xGu7ffp.png)

Now to quickly open external URL, click correlation id badge on top of response block:

![](https://i.imgur.com/K3L9Ze7.png) 

## Development

To ship hypersomnia to the target system we have to include HTML template into binary. We use `go generate` for that, so
after any changes to `templates/index.html` don't forget to run `go generate` and commit `templates/index.go`
 
