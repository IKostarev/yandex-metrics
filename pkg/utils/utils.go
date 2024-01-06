package utils

import (
	"fmt"
	"net/http"
)

func MetricNameNotFound(w http.ResponseWriter, metricName string) {
	fmt.Printf("[MetricNameNotFound] metric %s is empty\n", metricName)
	w.WriteHeader(NotFound)
}

func MetricBad(w http.ResponseWriter, metricName string) {
	fmt.Printf("[MetricBad] metric %s is bad\n", metricName)
	w.WriteHeader(BadRequest)
}

func PositiveAnswerToUser(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(statusCode)
}
