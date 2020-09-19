# Benchmarks

## Concurrency Control

Using a global mutex, on request level.
```bash
mutex.lock()
CS
mutex.unlock()
```

### Url Redirection (Load Tested Using Docker Containers)

[Graph - Url Redirection (no of Concurrent Users vs Response Time)](url_redirection_graph.png)

for 10 users, 120s
```bash
k6 run -d 120s -u 10 ./load_test_url_redirection.js
http_req_duration..........: avg=21.98ms min=2.21ms med=17.89ms max=527.23ms p(90)=39.74ms p(95)=50.01ms
http_reqs..................: 53675   447.291485/s
```

for 100 users, 120s
```bash
k6 run -d 120s -u 100 ./load_test_url_redirection.js
http_req_duration..........: avg=161.85ms min=2.81ms med=149.48ms max=956.05ms p(90)=284.61ms p(95)=331.39ms
http_reqs..................: 73807   615.056893/s
```

for 1000 users, 120s
```bash
k6 run -d 120s -u 1000 ./load_test_url_redirection.js
http_req_duration..........: avg=2.39s    min=0s     med=2.27s max=7.51s    p(90)=3.65s p(95)=4.16s
http_reqs..................: 49587   413.224692/s
```

### Url Shortening (Load Tested using Docker Containers)

for 10 users, 120s
```bash
k6 run -d 120s -u 10 ./load_test_shorten_url.js
http_req_duration..........: avg=85.44ms min=7.95ms med=83.36ms max=263.38ms p(90)=102.51ms p(95)=111.68ms
http_reqs..................: 13991   116.591622/s
```

for 100 users, 120s
```bash
k6 run -d 120s -u 100 ./load_test_shorten_url.js
http_req_duration..........: avg=901.78ms min=11.5ms  med=848.51ms max=2.21s   p(90)=1.07s p(95)=1.41s
http_reqs..................: 13262   110.516569/s
```

for 1000 users, 120s
```bash
k6 run -d 120s -u 1000 ./load_test_shorten_url.js
http_req_duration..........: avg=8.6s     min=0s      med=8.6s  max=19.03s   p(90)=10.33s p(95)=10.69s
http_reqs..................: 13345   111.207074/s`
```

## After Using Caching

Comparisons of **Read Latency** between Mongo and after using a Cache (Memcached).

```json
	memcached 	mongodb	
1	561.478µs	874.359µs
2	27.991374ms	3.377816ms
3	2.901262ms	4.669834ms
4	2.721016ms	3.289583ms
5	2.120171ms	76.469257ms

Avg 7.258ms			17.57ms
```
> NOTE - Around **200** percent decrease in read latency.

## Using MultiStage Build in Docker

Reduction in Size of Docker image by around **900%**

**Orig**
```bash
gorubu_app          latest              9dbe1cf26c39        About an hour ago   1.49GB
```

**After using Multistage build**
```bash
gorubu_docker_final   latest              912d533a9a52        4 minutes ago       25.4MB
```