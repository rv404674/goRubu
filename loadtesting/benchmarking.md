# BenchMarking Shortened Url

## Setup and Working

You need to have [k6](https://k6.io/) installed. It's written in Go and scriptable in JS.

1. K6 works with the concept of Virtual Users (VU)
2. Each VU executes your script in a completely separate JS runtime, parallel to all of the other running VU.Code inside the default function is called VU code, and is run over and over, for as long as the test is running.
3. Every virtual user (VU) performs the GET requests, in a continuous loop, as fast as it can.


### Benchmarking

*** For Shorten Url Endpoint ***

