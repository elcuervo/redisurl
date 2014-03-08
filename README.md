# redisurl

Parses a redis url and returns a struct with the result

```go
url = redisurl.Parse("redis://username:s3cr37@somewhere:6380/9")
// url.Port
// 6380
// url.Database
// 9
// url.Host
// somewhere
// url.Password
// s3cr37
```
You can also use `.Url` to get a Redis compatible string.
