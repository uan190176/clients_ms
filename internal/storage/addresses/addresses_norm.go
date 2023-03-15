package addresses

import (
	"clients_ms/internal/api"
	at "clients_ms/internal/storage/addresses_types"
	cfg "clients_ms/pkg/config"
	lgr "clients_ms/pkg/logger"
	"context"
	st "github.com/uan190176/statuses"
)

// GetValidationFlags - gets validation flags
func GetValidationFlags(ctx context.Context, req *api.RequestNormalizeAddress) (bool, bool, st.ResponseStatus) {
	lgr.LOG.Info("-->> ", "addresses.GetValidationFlags()")

	addressType, stat := at.GetAddressTypes(ctx, &api.RequestAddressType{
		AuthToken: cfg.CFG.Server.CliToken,
		Id:        req.AddressTypeId,
	})

	//Check result
	if len(addressType) == 0 || addressType == nil || stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return false, false, stat
	}

	lgr.LOG.Info("<<-- ", "addresses.GetValidationFlags()")
	return addressType[0].NeedsCleaning, addressType[0].NeedsNormalizing, st.GetStatus(100)
}
