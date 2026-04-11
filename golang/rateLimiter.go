/*
Create a simple backend API that limits how many requests a user can make.

🎯 Requirements
Create an API endpoint:
GET /api/request?userId=123
Logic:
Each user can make max 5 requests per minute
If limit exceeded → return error
✅ Expected Behavior
Allowed request:

	{
	  "success": true,
	  "message": "Request allowed"
	}

Rate limit exceeded:

	{
	  "success": false,
	  "message": "Rate limit exceeded"
	}

⚙️ Constraints
Use in-memory storage (no DB)
Time window: 1 minute
Must handle multiple users separately
*/

package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	mu       sync.Mutex
	requests map[string][]time.Time
	limit    int
	window   time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}
}

func (rl *RateLimiter) IsAllowed(userID string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	windowStart := now.Add(-rl.window)

	// old requests filter 
	var valid []time.Time
	for _, t := range rl.requests[userID] {
		if t.After(windowStart) {
			valid = append(valid, t)
		}
	}

	if len(valid) >= rl.limit {
		rl.requests[userID] = valid
		return false
	}

	rl.requests[userID] = append(valid, now)
	return true
}

func main() {
	rl := NewRateLimiter(5, time.Minute)

	http.HandleFunc("/api/request", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("userId")
		if userID == "" {
			http.Error(w, "userId required", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if rl.IsAllowed(userID) {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": true,
				"message": "Request allowed",
			})
		} else {
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"message": "Rate limit exceeded",
			})
		}
	})

	http.ListenAndServe(":8080", nil)
}
