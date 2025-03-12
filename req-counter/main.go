package reqcounter

import (
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	mu            sync.Mutex
	counter       map[int64]int
	totalRequests int
	lastRequest   time.Time
	checkTimer    *time.Timer
	surgeActive   bool
)

func init() {
	counter = make(map[int64]int)
	lastRequest = time.Now()
	surgeActive = false
}

// Handler function for requests
func handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Unix()

	mu.Lock()
	counter[currentTime]++
	totalRequests++
	lastRequest = time.Now()

	// If we weren't in a surge, we are now
	if !surgeActive {
		surgeActive = true
		log.Println("🌊 Request surge started")
		// Reset the timer if it exists
		if checkTimer != nil {
			checkTimer.Stop()
		}
	}

	// Reset the surge check timer
	resetSurgeCheckTimer()
	mu.Unlock()

	log.Printf("[request] %s | Path: %s | request-number: %d",
		r.Method, r.URL.Path, totalRequests)
}

// Resets the timer that checks for surge end
func resetSurgeCheckTimer() {
	if checkTimer != nil {
		checkTimer.Stop()
	}

	checkTimer = time.AfterFunc(5*time.Second, func() {
		mu.Lock()
		defer mu.Unlock()

		// If it's been 5 seconds since the last request and we were in a surge
		if time.Since(lastRequest) >= 5*time.Second && surgeActive {
			log.Printf("[surge ended] Total requests received: %d", totalRequests)
			surgeActive = false
			// reset the counter to track only per surge
			totalRequests = 0
			counter = make(map[int64]int)
		}
	})
}

// demoUse is a function that demonstrates how to use the reqcounter package
func DemoUse() {
	// Set up request handler
	http.HandleFunc("/", handler)

	// Start the backend server
	log.Println("🚀 Backend server running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
