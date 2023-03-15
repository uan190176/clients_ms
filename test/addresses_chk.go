package main

import (
	"clients_ms/internal/api"
	at "clients_ms/internal/storage/addresses_types"
	cfg "clients_ms/pkg/config"
	lgr "clients_ms/pkg/logger"
	ps "clients_ms/pkg/postservice"
	"context"
	st "github.com/uan190176/statuses"
	"strings"
)

// GetValidationFlags - gets validation flags
func GetValidationFlags(ctx context.Context, req *api.RequestAddress) (bool, bool, st.ResponseStatus) {
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

func NormalizeAddress(ctx context.Context, req *api.RequestAddress) (string, string, bool, st.ResponseStatus) {
	lgr.LOG.Info("-->> ", "addresses.NormalizeAddress()")

	//Check needs normalization
	needsCleaning, needsNormalizing, stat := GetValidationFlags(ctx, req)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return req.Address, "", false, stat
	}

	//TODO cleaning
	_ = needsCleaning

	//Normalize
	if needsNormalizing {
		normAddr, err := ps.GetNormalizedAddress(req.Address)
		if err != nil {
			lgr.LOG.Error("_ERR_: ", err)
			return req.Address, "", false, st.GetStatus(999, err.Error())
		}
		if addressIsApproved(normAddr.QualityCode, normAddr.ValidationCode) {
			return req.Address, normAddr.AddressJson, true, st.GetStatus(100)
		}
	}
	return req.Address, "", false, st.GetStatus(999, "address is not approved")
}

func addressIsApproved(qCode, vCode string) bool {
	lgr.LOG.Info("-->> ", "addresses.addressIsApproved()")
	validCodes := strings.Split(cfg.CFG.PostAPI.ValidValidationCodes, ",")
	qualCodes := strings.Split(cfg.CFG.PostAPI.ValidQualityCodes, ",")
	if contains(validCodes, vCode) && contains(qualCodes, qCode) {
		return true
	}
	lgr.LOG.Info("<<-- ", "addresses.addressIsApproved()")
	return false
}

func contains(s []string, e string) bool {
	e = strings.TrimSpace(e)

	for _, a := range s {
		a = strings.TrimSpace(a)
		if a == e {
			return true
		}
	}
	return false
}
