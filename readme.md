# goRubu

<img style="float: right;" width="600" src="./assets/goRubu.png"> 

This repo contains implementation of a **Url Shortner** written in [Go](https://golang.org/).

[![Build Status](https://travis-ci.com/rv404674/goRubu.svg?branch=master)](https://travis-ci.org/rv404674/goRubu)
[![Coverage Status](https://coveralls.io/repos/github/rv404674/goRubu/badge.svg?branch=master)](https://coveralls.io/github/rv404674/goRubu?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/rv404674/goRubu)](https://goreportcard.com/report/github.com/rv404674/goRubu)
[![Stars](https://img.shields.io/github/stars/rv404674/goRubu)](https://github.com/rv404674/goRubu/stargazers)
[![MIT License](https://img.shields.io/github/license/rv404674/goRubu)](https://github.com/rv404674/goRubu/blob/master/LICENSE)

## Contents

- [What is goRubu?](#what-is-gorubu-rocket)
- [Why goRubu?](#why-gorubu-dog)
- [BenchMarking](#benchmarking)
- [Running Server](#running-rerver-gear)
    - [Docker](#docker)
    - [Local](#local)
- [Api's](#apis-computer)
- [Monitoring](#monitoring-microscope)
- [Contributing](#contributing-beers)
- [Future Todo](#future-todo-notebook)
- [Maintainer](#maintainer-sunglasses)
- [License](#license-scroll)

## What is goRubu? :rocket:

1. A Url Shortner written in **Go**, with **MongoDb based** backend.
2. Supports Caching for Hot urls, with Memcached, using a LRU based eviction
strategy, and **write through type** of caching mechanism. Saw **[200%](/benchmarking/benchmarking.md)** decrease in Read Latency for URL redirection, after caching.
3. Used Travis CI for adding a **CI/CD** pipeline.
4. Dockerized the whole application. Used **Docker compose** for tying up different containers and **multi-stage build** for reducing the size of docker image by **[900%](/benchmarking/benchmarking.md)**.
4. **Prometheus and Grafana based monitoring**, to get an overall picture of the
system and application metrics.
5. Contains Api Validation and Logging Middlewares, along with Swagger based documentation

## Why goRubu? :dog:

Wanted to Learn Go and system design, by building a project. Hence goRubu.

## BenchMarking
> NOTE - Url Shortner is a read heavy system (read:write = 100:1), and these load tests are done on a single Machine (8Gb Ram, i5 Processor).

Check [this](/benchmarking/benchmarking.md) out for more info.

1. For **Url Redirection**, with **1000 Concurrent Users**, bombarding the app server for **2mins**:
```bash
http_req_duration..........: avg=2.39s    min=0s     med=2.27s max=7.51s    p(90)=3.65s p(95)=4.16s
http_reqs..................: 49587   413.224692/s
```

2. For **Url Shortening**, for same specs
```
http_req_duration..........: avg=8.6s     min=0s      med=8.6s  max=19.03s   p(90)=10.33s p(95)=10.69s
http_reqs..................: 13345   111.207074/s
```

## Running Server :gear:

### Docker

1. You need to have [docker](https://www.docker.com/) and **docker-compose** installed. After that just do
```bash
make docker
```
> Note - I haven't dockerized prometheus and grafana with it. Containerized goRubu and locally installed
> prometheus and grafana will work fine as prometheus is only listening to "local:8080/metrics".

Check the Api's Section afterwards.

### Local

### Prerequisites ✅

Ensure you have the following installed.
**[Mongodb](https://docs.mongodb.com/manual/)** 
**[Memcached](https://www.memcached.org/)**
**[Make](https://tutorialedge.net/golang/makefiles-for-go-developers/)**

On Macos, simply use thse
```bash
$ brew update
$ brew install mongodb/brew/mongodb-community
$ brew install memcached
$ brew install make
```

> **Note**: After that check whether these have been started, else you will get connection error.
```bash
$ brew services list
```

if any service is not up do
```bash
$ brew service start service_name
```

Then we need to download the tar files for prometheus, grafana and node exporter.
We will need to edit their config files, hence we are not using brew install for these. Also it becomes easy to run the server and do changes.

**For Grafana**
```bash
wget https://dl.grafana.com/oss/release/grafana-6.7.3.darwin-amd64.tar.gz
tar -zxvf grafana-6.7.3.darwin-amd64.tar.gz
mv grafana-6.7.3 /usr/local/bin
cd /usr/local/bin/grafana-6.7.3/bin
./grafana-server # this will run the grafana server
```

**Note** - By default prometheus server runs on 9090, and grafana on 3000.

Similarly download tar files for prometheus and node exporter and run there servers as well.

### Step to Run locally

1. Go to the dir where prometheus is installed and change the prometheus default .yml file to this one [new_yml](prometheus.yml), run the prometheus server, and node_exporter server.

2. do 
```bash 
$git clone https://github.com/rv404674/goRubu.git
```
3. cd in goRubu, and do
```bash
make deps
```
It will install all the go dependencies.

4. Then do 
```bash
make install
make execute
```
> **Note**: To see what these commands do check out this [makefile](Makefile)


## Api's :computer:

1. Hit **localhost:8080/all/shorten_url** with any url as key.
```json
{
	"Url": "https://www.redditgifts.com/exchanges/manage"
}
```

The endpoint will return a shortened URL.

2. Hit **http://localhost:8080/all/redirect** with the shortened url to get the original URL back.
```json
{
	"Url": "https://goRubu/MTAxMTA="
}
```

## Monitoring :microscope:

> Working:
1. Prometheus follows a Pull based Mechanism instead of Push Based.
2. goRubu exposes an HTTP endpoint exposing monitoring metrics.
3. Prometheus then periodically download the metrics. For UI, prometheus is used as a data source for Grafana.

<img style="float: left;" width="600" src="./assets/application_metrics.png"> 
<p align="left"> Grafana on Top of Prometheus </p>

## Contributing :beers:

Peformance Improvements, bug fixes, better design approaches are welcome. Please Discuss your change by raising an issue, beforehand.

## Future Todo :notebook:

1. Use Promtheus and Grafana as Docker Containers instead of Local Installation.
2. Use Kubernetes instead of Docker Compose.
3. Increase Coverage to more than 90%.
4. Add Monitoring in Prometheus for Number of 2xx, 3xx and 4xx.


# Maintainer :sunglasses:

[Rahul Verma Linkedin](https://www.linkedin.com/in/rahul-verma-8aa59b116/)
[Email](rv404674@gmail.com)

## License :scroll:

[MIT](LICENSE) © Rahul Verma