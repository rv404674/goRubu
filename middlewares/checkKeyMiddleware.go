package middlewares

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type inputUrl struct {
	Url string `json:"Url"`
}

type myResponseWriter struct {
	http.ResponseWriter
	buf *bytes.Buffer
}

type response struct {
	Writer   http.ResponseWriter
	UrlValue string
}

func (mrw *myResponseWriter) Write(p []byte) (int, error) {
	return mrw.buf.Write(p)
}

// GetUrlFromReq - NOTE when I read body, it becomes empty and I cannot read it twice.
// it happens because it is of type ReadCloser
// https://code-examples.net/en/q/2907302
func GetUrlFromReq(w http.ResponseWriter, r *http.Request) response {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
	}

	var input inputUrl
	_ = json.Unmarshal(body, &input)

	// Work / inspect body. You may even modify it!

	// And now set a new body, which will simulate the same data we read:
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	// Create a response wrapper:
	mrw := &myResponseWriter{
		ResponseWriter: w,
		buf:            &bytes.Buffer{},
	}

	// Call next handler, passing the response wrapper:
	// handler.ServeHTTP(mrw, r)
	return response{mrw, input.Url}

}

// CheckApiKey - check whether "Url" as a key exists in request body or not
func CheckApiKey(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// work for both "url" and "Url"
			log.Println("Request Validation")

			response := GetUrlFromReq(w, r)

			if response.UrlValue == "" {
				http.Error(w, "Missing Key", http.StatusBadRequest)
				log.Println("** Missing Key")
				return
			}

			h.ServeHTTP(w, r)
		})
}

// Logger - used for logging everything
func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			t1 := time.Now()
			log.Println(r.RequestURI)
			//defer log.Println(w.Header().Get("Body")) // this will run after service.UrlCreation is called
			h.ServeHTTP(w, r)
			t2 := time.Now()
			log.Printf("Logging [%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
			//log.Println(r.Response.StatusCode) - THIS WONT WORK
			// TODO find a way to get response status back
		})
}
