package services

import (
	"clients_ms/internal/api"
	hlp "clients_ms/internal/helpers"
	lgr "clients_ms/pkg/logger"
	"context"
	st "github.com/uan190176/statuses"
)

// GetMiscDataToCreateAndUpdateClientAddress - returns misc data to create and update clients address
func (s *GrpcClientsServer) GetMiscDataToCreateAndUpdateClientAddress(ctx context.Context, req *api.RequestMiscData) (*api.ResponseMiscData, error) {

	lgr.LOG.Info("***")
	lgr.LOG.Info("-->> ", "services.GetMiscDataToCreateAndUpdateClientAddress()")

	//Check auth token
	lgr.LOG.Info("_ACTION_: ", "checking auth token")
	if !hlp.AuthTokenIsValid(req.AuthToken) {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(101))
		return &api.ResponseMiscData{
			Countries:      nil,
			Deliveries:     nil,
			AddressesTypes: nil,
			Status:         hlp.GetStatusForResponse(st.GetStatus(101)),
		}, nil
	}

	//Get countries
	lgr.LOG.Info("_ACTION_: ", "getting countries")
	respCountries, _ := s.GetCountries(ctx, &api.RequestCountry{
		AuthToken: req.AuthToken,
		AuthorId:  req.AuthorId,
	})
	countries := respCountries.Countries
	countriesStatus := respCountries.Status

	//Get deliveries
	lgr.LOG.Info("_ACTION_: ", "getting deliveries")
	respDeliveries, _ := s.GetDeliveries(ctx, &api.RequestDelivery{
		AuthToken: req.AuthToken,
		AuthorId:  req.AuthorId,
	})
	deliveries := respDeliveries.Deliveries
	deliveriesStatus := respDeliveries.Status

	//Get addresses types
	lgr.LOG.Info("_ACTION_: ", "getting addresses types")
	respAddressesTypes, _ := s.GetAddressesTypes(ctx, &api.RequestAddressType{
		AuthToken: req.AuthToken,
		AuthorId:  req.AuthorId,
	})
	addressesTypes := respAddressesTypes.AddressesTypes
	AddressesTypesStatus := respAddressesTypes.Status

	resultStatus := st.ResponseStatus{}

	//Check statuses
	if countriesStatus.Code > 100 {
		resultStatus.Code = countriesStatus.Code
		resultStatus.Description = countriesStatus.Description
	}
	if deliveriesStatus.Code > 100 {
		resultStatus.Code = deliveriesStatus.Code
		resultStatus.Description += " " + deliveriesStatus.Description
	}
	if AddressesTypesStatus.Code > 100 {
		resultStatus.Code = AddressesTypesStatus.Code
		resultStatus.Description += " " + AddressesTypesStatus.Description
	}

	lgr.LOG.Info("<<-- ", "services.GetMiscDataToCreateAndUpdateClientAddress()")
	return &api.ResponseMiscData{
		Countries:      countries,
		Deliveries:     deliveries,
		AddressesTypes: addressesTypes,
		Status:         hlp.GetStatusForResponse(resultStatus),
	}, nil
}
