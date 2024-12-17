# echo-server

This is a simple echo server that listens on port 8080 and echoes back any data it receives in a structured JSON response.

```shell
curl -X POST \
  -H "X-Forwarded-For: 192.168.1.10, 10.0.0.1" \
  -H "X-Forwarded-Proto: https" \
  -H "X-Forwarded-Host: proxy.example.com" \
  "http://localhost:8080/param1/param2/param3?query1=1&query2=2"
```

```json
{
  "host": {
    "hostname": "localhost",
    "ip": "192.168.1.193",
    "ips": []
  },
  "http": {
    "method": "POST",
    "baseUrl": "",
    "originalUrl": "/param1/param2/param3?query1=1&query2=2",
    "protocol": "http"
  },
  "request": {
    "params": {
      "0": "param1",
      "1": "param2",
      "2": "param3"
    },
    "query": {
      "query1": "1",
      "query2": "2"
    },
    "cookies": {},
    "body": {},
    "headers": {
      "accept": "*/*",
      "user-agent": "curl/8.7.1",
      "x-forwarded-for": "192.168.1.10, 10.0.0.1",
      "x-forwarded-host": "proxy.example.com",
      "x-forwarded-proto": "https"
    }
  },
  "environment": {
    "HOME": "/Users/user"
  }
}
```
