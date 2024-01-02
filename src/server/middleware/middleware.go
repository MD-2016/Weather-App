package middleware

import (
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type siteVisitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var siteVisitors = make(map[string]*siteVisitor)
var mu sync.Mutex

func init() {
	go cleanOutSiteVisitors()
}

func getSiteVisitor(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	v, exists := siteVisitors[ip]

	if !exists {
		limiter := rate.NewLimiter(1, 3)
		siteVisitors[ip] = &siteVisitor{limiter, time.Now()}
		return limiter
	}

	v.lastSeen = time.Now()
	return v.limiter
}

func cleanOutSiteVisitors() {
	for {
		time.Sleep(time.Minute)

		mu.Lock()
		for ip, v := range siteVisitors {
			if time.Since(v.lastSeen) > 3*time.Minute {
				delete(siteVisitors, ip)
			}
		}
		mu.Unlock()
	}
}

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		limiter := getSiteVisitor(ip)
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
