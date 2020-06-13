# Commands for Debugging

1. To check whether entries are being created in **Memcached** or not.
```bash
telnet localhost 11211
get https://goRubu/MTAyMTk=
```
this will give the original url, to corresponding to the shorten url.
If you get nothing, it means there is not entry in memcached.

2.  To run tests.
```bash
go test ./tests -v
go test ./tests -v - cover
```

**NOTE** - Normally these two commands work, but with "go 1.13" they are not working.
Hence use these
go test ./... -v -coverpkg=./... -coverprofile=cover.txt
go tool cover -html=cover.txt -o cover.html

# Benchmarks

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


