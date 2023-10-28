package rest

import "net/http"

type HealthController struct{}

func (h *HealthController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
