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

## Benchmarks

1. Comparisons of **Read Latency** between Mongo and after using a Cache (Memcached).

```json
	memcached 	mongodb	
1	561.478µs	874.359µs
2	27.991374ms	3.377816ms
3	2.901262ms	4.669834ms
4	2.721016ms	3.289583ms
5	2.120171ms	76.469257ms

Avg 7.258ms			17.57ms
```
Around **200** percent decrease in read latency.

2. Reduction in Size of Docker image by around **900%**

**Orig**
```bash
gorubu_app          latest              9dbe1cf26c39        About an hour ago   1.49GB
```

**After using Multistage build**
```bash
gorubu_docker_final   latest              912d533a9a52        4 minutes ago       25.4MB
```