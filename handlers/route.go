package handlers

import (
	"fmt"
	"goRubu/middlewares"
	service "goRubu/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// HistPrometheus - was not able to directly create a prometheues middleware. Some type conversion problem. Google to
// find more.
type HistPrometheus struct {
	histogram *prometheus.HistogramVec
}

// Populate  - Created the histogram, with some keys
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

// FloatToString - to convert a float number to a string
func FloatToString(inputNum float64) string {
	return strconv.FormatFloat(inputNum, 'f', 6, 64)
}

// PrometheusMonitoring - This middleware will be use to address these things
// which REST endpoints are most used by consumers?
// how often?
// what are the response times?
func (histProme *HistPrometheus) PrometheusMonitoring(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			hist2 := histProme.histogram

			recordMetrics()
			t1 := time.Now()
			h.ServeHTTP(w, r)
			t2 := time.Now()

			hist2.WithLabelValues(
				FloatToString(t2.Sub(t1).Seconds()), // code
				r.URL.String(),                      // controller
				r.Method).Observe(float64(time.Since(t1)) / float64(time.Second))
		})
}

//Hellohandler - To test a basic handler
func Hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello. Go is an Awesome Language")
	time.Sleep(5 * time.Second)
}

//CreateUrlHandler - created new shortened Url
func CreateUrlHandler(w http.ResponseWriter, r *http.Request) {
	tempVar := middlewares.GetUrlFromReq(w, r)
	shortenedUrl := service.CreateShortenedUrl(tempVar.UrlValue)
	fmt.Fprintf(w, "Shortened Url: %s ", shortenedUrl)
	// type Response struct {
	// Message      string
	// ShortenedUrl string
	// }
	//json.NewEncoder(w).Encode(Response{Message: "Success", ShortenedUrl: shortenedUrl})
}

//RedirectionHandler - will return the original from which the shortened url was created
func RedirectionHandler(w http.ResponseWriter, r *http.Request) {
	orgUrl := service.UrlRedirection(middlewares.GetUrlFromReq(w, r).UrlValue)
	if orgUrl == "" {
		orgUrl = "This shortened Url doesn't exist in DB"
	}
	fmt.Fprintf(w, "Original Url: %s ", orgUrl)
}

//New - Add middlewares, endpoints to all the routes.
func New() http.Handler {
	//gorilla mux, supports addition of a middleware to a route.
	route := mux.NewRouter()
	// all denotes that all middlewares will be use
	// routes starting having prefix /all will have middlewares applied to them
	mainRoute := route.PathPrefix("/all").Subrouter()

	hist := HistPrometheus{}
	hist.Populate()
	// prometheusMonitoring ,CheckApiKey, Logger will be applied to all middlewares prefixed by all
	mainRoute.Use(hist.PrometheusMonitoring)
	mainRoute.Use(middlewares.CheckApiKey)
	mainRoute.Use(middlewares.Logger)

	// WILL expose default metrics, along with our custom metrics for go application
	// NOTE: Prometheus follows a Pull based Mechanism instead of Push Based.
	// Monitored applications exposes an HTTP endpoint exposing monitoring metrics.
	// Prometheus then periodically download the metrics.
	route.Handle("/metrics", promhttp.Handler()).Methods("GET")

	mainRoute.HandleFunc("/check", Hellohandler)
	mainRoute.HandleFunc("/shorten_url", CreateUrlHandler).Methods("POST")
	mainRoute.HandleFunc("/redirect", RedirectionHandler).Methods("POST")

	return route
}
