# my-golang-api-template

Base on [Echo](https://github.com/labstack/echo)

## Installtion

```Bash
$> go get -u github.com/labstack/echo/
```

## Test

```Bash
$> curl localhost:8081/health
{"time":"2019-10-07T18:39:42.082643+08:00","id":"","remote_ip":"127.0.0.1","host":"127.0.0.1:8081","method":"GET","uri":"/health","user_agent":"curl/7.54.0","status":200,"error":"","latency":14981,"latency_human":"14.981µs","bytes_in":0,"bytes_out":2}
```