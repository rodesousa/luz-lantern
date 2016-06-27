# Shard


## User

Check user state

```
- user:
    name : name of user [string]
```

check if root is present 
```
- user:
    name : root
```

## Curl

Try a ws

```
- curl:
    url : url [string]
    name : give a name to the test [string]
```

check a url of google
```
- curl:
    url: "http://www.google.fr" 
    name: "google testings"
```

## Ping

Try to resolve a url
```
- ping:
    url : url [string]
    name : give a name to the test [string]
```

try to resolve google.com
```
- ping:
    url: "http://www.google.fr"
    name: "Resolve Google"
```
