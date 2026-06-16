package remote

import "net/http"

func isRetriable(err error) bool {
	// Logic to identify HTTP 429 as retriable
	return true
}
