### Httpparser
A simple concurrent http parser that fetches response from multiple servers and returns an MD5 hash for them.

### How to setup?
- Make sure you have GoLang installed on your machine.
- Clone the repository run these commands on terminal.
```
cd Httpparser
export GOPATH="$PWD"
```

### How to build?
```
go install httpparser
```

### How to run?
```
./bin/httpparser -concurrency (level of concurrency) url1 url2 url3 ...
```
### Correct Usage
```
./bin/httpparser -concurrency 10 http://www.google.com http://www.yahoo.com

Output
2020/04/23 20:17:00 http://www.google.com a47e04ea7e27a988dff4741ee8b9d247
2020/04/23 20:17:01 http://www.yahoo.com 486748ef76abab71edf87b9899f6b56b
```

### Incorrect Usage
```
./bin/httpparser -concurrency 10 wrongUrl http://www.google.com

Output
2020/04/23 20:23:14 Could not parse resonse for url: wrongUrl error: Get "wrongUrl": unsupported protocol scheme ""
2020/04/23 20:23:16 http://www.yahoo.com 7c1660ba78fd94f955a7cec273c2c8e4
```

How to run tests?
````
go test ./...
````

### Note
- The concurreny flag can be ignored, by default 10 threads are spwaned.
- The maximum allowed concurreny is 100.
- The minimum allowed concurency is 1.
