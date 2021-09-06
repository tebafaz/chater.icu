package middlewares

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type sessionCounter struct {
	mu       sync.RWMutex
	counter  int
	maxCount int
}

var sc *sessionCounter

func NewSessionCounter() {
	sc = &sessionCounter{}
}

func (scount *sessionCounter) increment() {
	scount.mu.Lock()
	defer scount.mu.Unlock()
	scount.counter++
}

func (scount *sessionCounter) decrement() {
	scount.mu.Lock()
	defer scount.mu.Unlock()
	scount.counter--
}

func (scount *sessionCounter) checkvalue() int {
	scount.mu.Lock()
	defer scount.mu.Unlock()
	return scount.counter
}

func SessionCounter(maxCount int) gin.HandlerFunc {
	sc.maxCount = maxCount
	return func(c *gin.Context) {
		sc.increment()
		c.Next()
		sc.decrement()
	}

}

func CheckSessionNumber(c *gin.Context) {
	if sc.checkvalue() > sc.maxCount {
		c.AbortWithStatus(http.StatusTooManyRequests)
	}
}
