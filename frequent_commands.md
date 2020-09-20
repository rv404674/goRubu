# Commands for Debugging

1. To check whether entries are being created in **Memcached** or not.
```bash
telnet localhost 11211
get https://goRubu/MTAyMTk=
```
This will give the original url, corresponding to the shortened url.


2. To check Indexes in Mongo.
```bash
db.collection.getIndexes()
```

3. To run tests.
```bash
go test ./tests -v
go test ./tests -v - cover
```

**NOTE** - Normally these two commands work, but with "go 1.13", second command (the cover one) is not working.

I have integrated [coveralls.io](https://coveralls.io/) with goRubu.
or you can use these.

```bash
go test ./... -v -coverpkg=./... -coverprofile=cover.txt
go tool cover -html=cover.txt -o cover.html
```

## Docker 

1. This will give the logs of a particular container.
```bash
docker logs container_id
```

2. Get into the container and see what is happening.
```bash
docker exec -it container_id /bin/sh
```

3. Build an image with some name("gorubuimage") from dockerfile
```bash
docker build -t gorubuimage .
```
