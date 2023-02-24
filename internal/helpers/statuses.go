package helpers

import (
	"clients_ms/internal/api"
	st "github.com/uan190176/statuses"
)

// GetStatusForResponse - returns instance of status for response
func GetStatusForResponse(status st.ResponseStatus) *api.ClientsMS_Status {
	return &api.ClientsMS_Status{
		Code:        status.Code,
		Description: status.Description,
	}
}