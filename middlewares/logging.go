package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/justinas/alice"
	"github.com/sirupsen/logrus"
)

type LoggingMw struct {
	next http.Handler
}

func NewLoggingMw() alice.Constructor {
	return func(next http.Handler) http.Handler {
		return NewLogging(next)
	}
}

// Called once for each middleware chain
//
func NewLogging(next http.Handler) *LoggingMw {
	return &LoggingMw{next: next}
}

// This should be the first Middleware in the chain
//
func (mw *LoggingMw) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Printf("In LoggingMw::ServeHTTP: PRE\n")

	txnId := uuid.New().String()
	startTime := time.Now()

	rw.Header().Set("X-Txn-ID", txnId)

	mw.next.ServeHTTP(rw, r)

	logrus.WithFields(
		logrus.Fields{
			"method": r.Method,
			"proto":  r.Proto,
			"host":   r.Host,
			"remote": r.RemoteAddr,
			"start":  time.Now().Sub(startTime),
			"path":   r.URL.String(),
			"txnid":  txnId,
		},
	).Info("thing")

	fmt.Printf("In LoggingMw::ServeHTTP : POST\n")
}
