goRubu

A URL Shortner written in GoLang, with a Mongo based backend.
Supports Caching for Hot URLs, with Memcached, using an LRU based eviction strategy, and write through a type of caching mechanism.
Prometheus and Grafana based monitoring, to get an overall picture of the system
Swagger based documentation.

1. Makefile - if someone wants to try out our project, can easy install all the packages.
do "make deps" - to install the listed go packages.