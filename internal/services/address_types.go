package services

import (
	"clients_ms/internal/api"
	hlp "clients_ms/internal/helpers"
	stg "clients_ms/internal/storage/addresses_types"
	cfg "clients_ms/pkg/config"
	lgr "clients_ms/pkg/logger"
	"context"
	st "github.com/uan190176/statuses"
)

const addressTypesTableName = "public.addresses_types_ref"

// GetAddressesTypes - returns address types
func (s *GrpcClientsServer) GetAddressesTypes(ctx context.Context, req *api.RequestAddressType) (*api.ResponseAddressesTypes, error) {

	lgr.LOG.Info("***")
	lgr.LOG.Info("-->> ", "services.GetAddressesTypes()")

	//Check auth token
	lgr.LOG.Info("_ACTION_: ", "checking auth token")
	if !hlp.AuthTokenIsValid(req.AuthToken) {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(101))
		return &api.ResponseAddressesTypes{
			AddressesTypes: nil,
			Status:         hlp.GetStatusForResponse(st.GetStatus(101)),
		}, nil
	}

	//Check UAR
	lgr.LOG.Info("_ACTION_: ", "checking users access rights")
	allow, stat := hlp.GetUserAccessRightForTable(ctx, &api.RequestUsersAccessRightsForTable{
		AuthToken: cfg.CFG.MicroServices["users_ms"].Token,
		UserId:    req.AuthorId,
		TableName: addressTypesTableName,
		Action:    api.UsersAccessRights_Actions_CAN_SELECT,
	})

	//error
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return &api.ResponseAddressesTypes{
			AddressesTypes: nil,
			Status:         hlp.GetStatusForResponse(stat),
		}, nil
	}

	//not allow
	if !allow {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(201))
		return &api.ResponseAddressesTypes{
			AddressesTypes: nil,
			Status:         hlp.GetStatusForResponse(st.GetStatus(201)),
		}, nil
	}

	// Work with storage
	lgr.LOG.Info("_ACTION_: ", "working with storage")
	addressesTypes, status := stg.GetAddressTypes(ctx, req)

	lgr.LOG.Info("_RESULT_: ", status)
	lgr.LOG.Info("-->> ", "services.GetAddressesTypes()")
	lgr.LOG.Info("***")

	return &api.ResponseAddressesTypes{
		AddressesTypes: addressesTypes,
		Status:         hlp.GetStatusForResponse(status),
	}, nil
}

// CreateAddressType - creates new address type
func (s *GrpcClientsServer) CreateAddressType(ctx context.Context, req *api.RequestAddressType) (*api.ResponseAddressesTypes, error) {

	lgr.LOG.Info("***")
	lgr.LOG.Info("-->> ", "services.CreateAddressType()")

	// Check auth token
	lgr.LOG.Info("_ACTION_: ", "checking auth token")
	if !hlp.AuthTokenIsValid(req.AuthToken) {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(101))
		return &api.ResponseAddressesTypes{
			AddressesTypes: nil,
			Status:         hlp.GetStatusForResponse(st.GetStatus(101)),
		}, nil
	}

	lgr.LOG.Info("_ACTION_: ", "checking users access rights")
	allow, stat := hlp.GetUserAccessRightForTable(ctx, &api.RequestUsersAccessRightsForTable{
		AuthToken: cfg.CFG.MicroServices["users_ms"].Token,
		UserId:    req.AuthorId,
		TableName: addressTypesTableName,
		Action:    api.UsersAccessRights_Actions_CAN_INSERT,
	})

	//error
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return &api.ResponseAddressesTypes{
			AddressesTypes: nil,
			Status:         hlp.GetStatusForResponse(stat),
		}, nil
	}

	//not allow
	if !allow {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(202))
		return &api.ResponseAddressesTypes{
			AddressesTypes: nil,
			Status:         hlp.GetStatusForResponse(st.GetStatus(202)),
		}, nil
	}

	// Work with storage
	lgr.LOG.Info("_ACTION_: ", "working with storage")
	addrTypes, status := stg.CreateAddressType(ctx, req)

	//Result
	lgr.LOG.Info("_RESULT_: ", status)
	lgr.LOG.Info("<<-- ", "services.CreateAddressType()")
	lgr.LOG.Info("***")
	return &api.ResponseAddressesTypes{
		AddressesTypes: addrTypes,
		Status:         hlp.GetStatusForResponse(status),
	}, nil
}

// UpdateAddressType - updates address type
func (s *GrpcClientsServer) UpdateAddressType(ctx context.Context, req *api.RequestAddressType) (*api.ResponseAddressesTypes, error) {

	lgr.LOG.Info("***")
	lgr.LOG.Info("-->> ", "services.UpdateAddressType()")

	// Check auth token
	lgr.LOG.Info("_ACTION_: ", "checking auth token")
	if !hlp.AuthTokenIsValid(req.AuthToken) {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(101))
		return &api.ResponseAddressesTypes{
			AddressesTypes: nil,
			Status:         hlp.GetStatusForResponse(st.GetStatus(101)),
		}, nil
	}

	// Check uars
	lgr.LOG.Info("_ACTION_: ", "checking users access rights")
	allow, stat := hlp.GetUserAccessRightForTable(ctx, &api.RequestUsersAccessRightsForTable{
		AuthToken: cfg.CFG.MicroServices["users_ms"].Token,
		UserId:    req.AuthorId,
		TableName: addressTypesTableName,
		Action:    api.UsersAccessRights_Actions_CAN_UPDATE,
	})

	//error
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return &api.ResponseAddressesTypes{
			AddressesTypes: nil,
			Status:         hlp.GetStatusForResponse(stat),
		}, nil
	}

	//not allow
	if !allow {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(203))
		return &api.ResponseAddressesTypes{
			AddressesTypes: nil,
			Status:         hlp.GetStatusForResponse(st.GetStatus(203)),
		}, nil
	}

	// Work with storage
	lgr.LOG.Info("_ACTION_: ", "working with storage")
	addrTypes, status := stg.UpdateAddressType(ctx, req)

	//Result
	lgr.LOG.Info("_RESULT_: ", status)
	lgr.LOG.Info("<<-- ", "services.UpdateAddressType()")
	lgr.LOG.Info("***")
	return &api.ResponseAddressesTypes{
		AddressesTypes: addrTypes,
		Status:         hlp.GetStatusForResponse(status),
	}, nil
}

// UpdateAddressTypeDeletionFlags - mark/unmark addresses types as deleted
func (s *GrpcClientsServer) UpdateAddressTypeDeletionFlags(ctx context.Context, req *api.RequestAddressesTypesDeletionFlags) (*api.ResponseAddressesTypes, error) {
	lgr.LOG.Info("***")
	lgr.LOG.Info("-->> ", "services.UpdateDeletedFlagAddressType()")

	// Check auth token
	lgr.LOG.Info("_ACTION_: ", "checking auth token")
	if !hlp.AuthTokenIsValid(req.AuthToken) {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(101))
		return &api.ResponseAddressesTypes{
			AddressesTypes: nil,
			Status:         hlp.GetStatusForResponse(st.GetStatus(101)),
		}, nil
	}

	lgr.LOG.Info("_ACTION_: ", "checking users access rights")
	allow, stat := hlp.GetUserAccessRightForTable(ctx, &api.RequestUsersAccessRightsForTable{
		AuthToken: cfg.CFG.MicroServices["users_ms"].Token,
		UserId:    req.AuthorId,
		TableName: addressTypesTableName,
		Action:    api.UsersAccessRights_Actions_CAN_DELETE,
	})

	//error
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return &api.ResponseAddressesTypes{
			AddressesTypes: nil,
			Status:         hlp.GetStatusForResponse(stat),
		}, nil
	}

	//not allow
	if !allow {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(204))
		return &api.ResponseAddressesTypes{
			AddressesTypes: nil,
			Status:         hlp.GetStatusForResponse(st.GetStatus(204)),
		}, nil
	}

	// Work with storage
	lgr.LOG.Info("_ACTION_: ", "working with storage")
	addrTypes, status := stg.UpdateAddressTypeDeletionFlags(ctx, req)

	//Result
	lgr.LOG.Info("_RESULT_: ", status)
	lgr.LOG.Info("<<-- ", "services.UpdateDeletedFlagAddressType()")
	lgr.LOG.Info("***")
	return &api.ResponseAddressesTypes{
		AddressesTypes: addrTypes,
		Status:         hlp.GetStatusForResponse(status),
	}, nil
}
