package cache

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/joho/godotenv"
)

// NOTE:
// Try to establish a connection with Memcached container or Local Memcached

// EXPIRY_TIME - TTL for an item int cache
var EXPIRY_TIME int

func init() {
	dir, _ := os.Getwd()
	envFile := "variables.env"
	if strings.Contains(dir, "test") {
		envFile = "../variables.env"
	}

	if err := godotenv.Load(envFile); err != nil {
		log.Fatal("Error: No Environment File Found, cacheConnection.go", err)
	}

	// in seconds
	EXPIRY_TIME, _ = strconv.Atoi(os.Getenv("EXPIRATION_TIME"))
}

func tryMemcached(domain string) *memcache.Client {
	mc := memcache.New(domain)

	inputUrl := "https://stackoverflow.com/questions/58442596/golang-base64-to-hex-conversion"
	newUrl := "https://goRubu/MTAyNDE="

	err := mc.Set(&memcache.Item{
		Key:        newUrl,
		Value:      []byte(inputUrl),
		Expiration: int32(EXPIRY_TIME),
	})

	if err != nil {
		log.Printf("Err: %v, Domain: %v", err, domain)
		return nil
	}

	return mc
}

// CreateCon - Create Memcached Connection
// as this is called by mainService init function only once, hence checking whether local/docker memcache
// is up will happen only once.
func CreateCon() *memcache.Client {
	var cacheDomain = os.Getenv("MEMCACHED_DOMAIN_DOCKER")
	var client *memcache.Client

	client = tryMemcached(cacheDomain)

	if client == nil {
		log.Println("Connection Failed while trying to connect with Memcached Container")

		cacheDomain = os.Getenv("MEMCACHED_DOMAIN_LOCALHOST")
		client = tryMemcached(cacheDomain)
		if client == nil {
			log.Fatal("Connection Failed while Trying to connect with Local Memcached")
		}

		log.Println("Connected to Local Memcached!!")
	} else {
		log.Println("Connected to Memcached Container!!")
	}

	return client
}
