package healthz

import (
	"encoding/json"
	"net/http"
)

type healthCheckResult struct {
	Message string `json:"message,omitempty"`
}

func Handler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		result := healthCheckResult{
			Message: "ok",
		}

		data, _ := json.Marshal(result)
		_, _ = w.Write(data)
	}
}
