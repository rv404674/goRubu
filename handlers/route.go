package handlers

import (
	"fmt"
	"goRubu/middlewares"
	service "goRubu/services"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Response struct {
	Message      string
	ShortenedUrl string
}

type RedirectionResp struct {
	Message     string
	OriginalUrl string
}

type HistPrometheus struct {
	histogram *prometheus.HistogramVec
}

// Initialize it somewhere
func (histProme *HistPrometheus) Populate() {
	histProme.histogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_durations_histogram_seconds",
		Help:    "HTTP request latency distributions.",
		Buckets: prometheus.ExponentialBuckets(0.0001, 1.5, 36),
	}, []string{"total_time_taken", "controller", "action"})

	prometheus.MustRegister(histProme.histogram)
}

// used to check number of incoming req
func recordMetrics() {
	reqProcessed.Inc()
}

var (
	reqProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "total_requests",
		Help: "The total number of Incoming requests",
	})
)

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

// This middleware will be use to address these things
// which REST endpoints are most used by consumers?
// how often?
// what are the response times?

func (histProme *HistPrometheus) PrometheusMonitoring(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			hist2 := histProme.histogram

			recordMetrics()

			t1 := time.Now()
			log.Println(r.RequestURI)
			//defer log.Println(w.Header().Get("Body")) // this will run after service.UrlCreation is called
			h.ServeHTTP(w, r)
			t2 := time.Now()
			log.Printf("Logging [%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
			//log.Println(r.Response.StatusCode) - THIS WONT WORK
			// TODO find a way to get response status back
			hist2.WithLabelValues(
				FloatToString(t2.Sub(t1).Seconds()), // code
				r.URL.String(),                      // controller
				r.Method).Observe(float64(time.Since(t1)) / float64(time.Second))
		})
}

//To test a basic handler
func Hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello. Go is an Awesome Language")
	time.Sleep(5 * time.Second)
}

//created new shortened Url
func CreateUrlHandler(w http.ResponseWriter, r *http.Request) {
	temp_var := middlewares.GetUrlFromReq(w, r)
	shortenedUrl := service.CreateShortenedUrl(temp_var.UrlValue)
	fmt.Fprintf(w, "Shortened Url: %s ", shortenedUrl)
	//json.NewEncoder(w).Encode(Response{Message: "Success", ShortenedUrl: shortenedUrl})
}

func RedirectionHandler(w http.ResponseWriter, r *http.Request) {
	orgUrl := service.UrlRedirection(middlewares.GetUrlFromReq(w, r).UrlValue)
	if orgUrl == "" {
		orgUrl = "This shortened Url doesn't exist in DB"
	}
	fmt.Fprintf(w, "Original Url: %s ", orgUrl)
}

func New() http.Handler {
	//gorilla mux, supports addition of a middleware to a route.
	route := mux.NewRouter()
	// all denotes that all middlewares will be use
	// routes starting having prefix /all will have middlewares applied to them
	main_route := route.PathPrefix("/all").Subrouter()

	hist := HistPrometheus{}
	hist.Populate()
	// prometheusMonitoring and CheckApiKey will be applied to all middlewares prefixed by
	// /all
	main_route.Use(hist.PrometheusMonitoring)
	main_route.Use(middlewares.CheckApiKey)

	// WILL expose default metrics for go application
	// also as we won't be passing "url" in body, so to prevent missing key from
	// being returned from checkApiKey middleware, we wont prefix this with all
	route.Handle("/metrics", promhttp.Handler()).Methods("GET")

	main_route.HandleFunc("/check", Hellohandler)
	main_route.HandleFunc("/shorten_url", CreateUrlHandler).Methods("POST")
	main_route.HandleFunc("/redirect", RedirectionHandler).Methods("POST")

	return route
}
