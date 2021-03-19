# BenchMarking Shortened Url

## Setup and Working

You need to have [k6](https://k6.io/) installed. It's written in Go and scriptable in JS.

1. K6 works with the concept of Virtual Users (VU)
2. Each VU executes your script in a completely separate JS runtime, parallel to all of the other running VU. Code inside the default function is called VU code, and is run over and over, for as long as the test is running.
NOTE: Virtual Users are designed to act and behave like real users/browsers would. That is, they are capable of making multiple network connections in parallel, just like a real user in a browser would.
https://k6.io/docs/misc/glossary#virtual-users

3. Every virtual user (VU) performs the GET requests, in a continuous loop, as fast as it can.


### Benchmarking

*** For Shorten Url Endpoint ***
1. For 1VU, in 30s.
```bash
k6 run -d 5s -u 1 ./load_test_shorten_url.js

http_req_duration..........: avg=5ms     min=1.66ms med=3.15ms max=222.44ms p(90)=8.7ms   p(95)=12.5ms
http_reqs..................: 934     186.794732/s
```

> Conclusion
95% of our users got served a response in under 12.5ms.
In the 30 second test duration we served 934 responses, at a rate of ~187 requests per second (RPS).