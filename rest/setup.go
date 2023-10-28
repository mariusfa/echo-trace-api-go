package rest

import "net/http"

func SetupHandlers(healthController HealthController) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthController.HealthCheck)
	return mux
}
