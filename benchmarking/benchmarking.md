# Benchmarks
NOTE: These load tests are done on a single machine having 8GB Ram, a Dual-Core Intel Core i5 Processor.

### Url Redirection (Load Tested Using Docker Containers)
[Graph - Url Redirection (no of Concurrent Users vs Response Time)](url_redirection_graph.png)

 k6 run --summary-trend-stats "min,avg,max,p(95),p(99),p(99.99)" -d 120s -u 200 ./load_test_url_redirection.js

100users, 120s
http_req_duration..........: min=2.75ms avg=190.44ms max=950.19ms p(95)=388.29ms p(99)=523.61ms p(99.99)=872.22ms

http_reqs..................: 62701   522.507805/s

200users, 120s
http_req_duration..........: min=2.23ms avg=372.33ms max=1.73s    p(95)=701.94ms p(99)=986.86ms p(99.99)=1.48s

http_reqs..................: 64160   534.666138/s

300users, 120s
 http_req_duration..........: min=2.58ms avg=561.86ms max=2.21s    p(95)=1.02s p(99)=1.41s   p(99.99)=2.1s
http_reqs..................: 63741   531.172736/s

400users, 120s
http_req_duration..........: min=4.15ms avg=795.82ms max=2.75s    p(95)=1.42s p(99)=1.77s    p(99.99)=2.47s
http_reqs..................: 59934   499.448018/s

500users, 120s
http_req_duration..........: min=3.52ms avg=1s      max=4.38s    p(95)=1.73s p(99)=2.56s    p(99.99)=4.29s
http_reqs..................: 59490   495.748899/s

600users, 120s


for 10 users, 120s
```bash
k6 run --summary-trend-stats "min,avg,max,p(95),p(99),p(99.99)" -d 120s -u 1000 ./load_test_url_redirection.js

http_req_duration..........: min=5.22ms avg=2.26s    max=11.75s   p(95)=4.62s p(99)=6.79s    p(99.99)=11.67s
http_reqs..................: 55678   434.924911/s
```

for 100 users, 120s
```bash
k6 run --summary-trend-stats "min,avg,max,p(95),p(99),p(99.99)" -d 120s -u 100 ./load_test_url_redirection.js

```

for 1000 users, 120s
```bash
k6 run --summary-trend-stats "min,avg,max,p(95),p(99),p(99.99)" -d 120s -u 1000 ./load_test_url_redirection.js

http_req_duration..........: min=5.22ms avg=2.26s    max=11.75s   p(95)=4.62s p(99)=6.79s    p(99.99)=11.67s
http_reqs..................: 55678   434.924911/s
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