# key value storage is okay

```sh
go run main.go
```

### set value 'content' to 'mykey'
```sh
curl -X POST 'http://127.0.0.1:8021/mykey' --data-raw 'content'
```

### get value of key 'mykey'
```sh
curl -X GET 'http://127.0.0.1:8021/mykey'
```

### delete value of key 'mykey'
```sh
curl -X DELETE 'http://127.0.0.1:8021/mykey'
```
