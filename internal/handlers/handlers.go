package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"yandex-metrics/internal/storage/memory"
	"yandex-metrics/pkg/utils"
)

var store = memory.NewMemoryStorage()

func GaugeHandler(w http.ResponseWriter, r *http.Request) {
	var metricsName = chi.URLParam(r, "name")
	var metricsValue = chi.URLParam(r, "value")

	if metricsName == "" {
		utils.MetricNameNotFound(w, metricsName)
		return
	}

	valueToFloat, err := strconv.ParseFloat(metricsValue, 64)
	if err != nil {
		_ = fmt.Errorf("[GaugeHandler] error parse int in url params")
		utils.MetricBad(w, metricsName)
		return
	}

	store.AddGauge(metricsName, valueToFloat)

	utils.PositiveAnswerToUser(w, utils.OK)
}

func CounterHandler(w http.ResponseWriter, r *http.Request) {
	var metricsName = chi.URLParam(r, "name")
	var metricsValue = chi.URLParam(r, "value")

	if metricsName == "" {
		utils.MetricNameNotFound(w, metricsName)
		return
	}

	valueToInt, err := strconv.ParseInt(metricsValue, 0, 64)
	if err != nil {
		_ = fmt.Errorf("[CounterHandler] error parse int in url params")
		utils.MetricBad(w, metricsName)
		return
	}

	store.AddCounter(metricsName, valueToInt)

	utils.PositiveAnswerToUser(w, utils.OK)
}

func BadRequestHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(utils.BadRequest)
}
