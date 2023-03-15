package services

import (
	"clients_ms/internal/api"
	hlp "clients_ms/internal/helpers"
	stg "clients_ms/internal/storage/addresses"
	cfg "clients_ms/pkg/config"
	lgr "clients_ms/pkg/logger"
	"context"
	st "github.com/uan190176/statuses"
)

const addressesTableName = "public.addresses_ref"

// GetAddresses - returns clients addresses
func (s *GrpcClientsServer) GetAddresses(ctx context.Context, req *api.RequestAddress) (*api.ResponseAddresses, error) {

	lgr.LOG.Info("***")
	lgr.LOG.Info("-->> ", "services.GetAddresses()")

	//Check auth token
	lgr.LOG.Info("_ACTION_: ", "checking auth token")
	if !hlp.AuthTokenIsValid(req.AuthToken) {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(101))
		return &api.ResponseAddresses{
			Addresses: nil,
			Status:    hlp.GetStatusForResponse(st.GetStatus(101)),
		}, nil
	}

	//Check users access rights
	lgr.LOG.Info("_ACTION_: ", "checking users access rights")
	allow, stat := hlp.GetUserAccessRightForTable(ctx, &api.RequestUsersAccessRightsForTable{
		AuthToken: cfg.CFG.MicroServices["users_ms"].Token,
		UserId:    req.AuthorId,
		TableName: addressesTableName,
		Action:    api.UsersAccessRights_Actions_CAN_SELECT,
	})

	//error
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return &api.ResponseAddresses{
			Addresses: nil,
			Status:    hlp.GetStatusForResponse(stat),
		}, nil
	}

	//not allow
	if !allow {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(201))
		return &api.ResponseAddresses{
			Addresses: nil,
			Status:    hlp.GetStatusForResponse(st.GetStatus(201)),
		}, nil
	}

	// Work with storage
	lgr.LOG.Info("_ACTION_: ", "working with storage")
	addresses, status := stg.GetAddresses(ctx, req)

	lgr.LOG.Info("_RESULT_: ", status)
	lgr.LOG.Info("-->> ", "services.GetAddresses()")
	lgr.LOG.Info("***")

	return &api.ResponseAddresses{
		Addresses: addresses,
		Status:    hlp.GetStatusForResponse(status),
	}, nil
}

// CreateAddress - creates new address
func (s *GrpcClientsServer) CreateAddress(ctx context.Context, req *api.RequestAddress) (*api.ResponseAddresses, error) {

	lgr.LOG.Info("***")
	lgr.LOG.Info("-->> ", "services.CreateAddress()")

	// Check auth token
	lgr.LOG.Info("_ACTION_: ", "checking auth token")
	if !hlp.AuthTokenIsValid(req.AuthToken) {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(101))
		return &api.ResponseAddresses{
			Addresses: nil,
			Status:    hlp.GetStatusForResponse(st.GetStatus(101)),
		}, nil
	}

	lgr.LOG.Info("_ACTION_: ", "checking users access rights")
	allow, stat := hlp.GetUserAccessRightForTable(ctx, &api.RequestUsersAccessRightsForTable{
		AuthToken: cfg.CFG.MicroServices["users_ms"].Token,
		UserId:    req.AuthorId,
		TableName: notesTableName,
		Action:    api.UsersAccessRights_Actions_CAN_INSERT,
	})

	//error
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return &api.ResponseAddresses{
			Addresses: nil,
			Status:    hlp.GetStatusForResponse(stat),
		}, nil
	}

	//not allow
	if !allow {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(202))
		return &api.ResponseAddresses{
			Addresses: nil,
			Status:    hlp.GetStatusForResponse(st.GetStatus(202)),
		}, nil
	}

	// Work with storage
	lgr.LOG.Info("_ACTION_: ", "working with storage")
	addresses, status := stg.CreateAddress(ctx, req)

	//Result
	lgr.LOG.Info("_RESULT_: ", status)
	lgr.LOG.Info("<<-- ", "services.CreateAddress()")
	lgr.LOG.Info("***")
	return &api.ResponseAddresses{
		Addresses: addresses,
		Status:    hlp.GetStatusForResponse(status),
	}, nil
}

// UpdateAddress - updates address
func (s *GrpcClientsServer) UpdateAddress(ctx context.Context, req *api.RequestAddress) (*api.ResponseAddresses, error) {

	lgr.LOG.Info("***")
	lgr.LOG.Info("-->> ", "services.UpdateAddress()")

	// Check auth token
	lgr.LOG.Info("_ACTION_: ", "checking auth token")
	if !hlp.AuthTokenIsValid(req.AuthToken) {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(101))
		return &api.ResponseAddresses{
			Addresses: nil,
			Status:    hlp.GetStatusForResponse(st.GetStatus(101)),
		}, nil
	}

	lgr.LOG.Info("_ACTION_: ", "checking users access rights")
	allow, stat := hlp.GetUserAccessRightForTable(ctx, &api.RequestUsersAccessRightsForTable{
		AuthToken: cfg.CFG.MicroServices["users_ms"].Token,
		UserId:    req.AuthorId,
		TableName: notesTableName,
		Action:    api.UsersAccessRights_Actions_CAN_UPDATE,
	})

	//error
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return &api.ResponseAddresses{
			Addresses: nil,
			Status:    hlp.GetStatusForResponse(stat),
		}, nil
	}

	//not allow
	if !allow {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(203))
		return &api.ResponseAddresses{
			Addresses: nil,
			Status:    hlp.GetStatusForResponse(st.GetStatus(203)),
		}, nil
	}

	// Work with storage
	lgr.LOG.Info("_ACTION_: ", "working with storage")
	addresses, status := stg.UpdateAddress(ctx, req)

	//Result
	lgr.LOG.Info("_RESULT_: ", status)
	lgr.LOG.Info("<<-- ", "services.UpdateAddress()")
	lgr.LOG.Info("***")
	return &api.ResponseAddresses{
		Addresses: addresses,
		Status:    hlp.GetStatusForResponse(status),
	}, nil
}

// UpdateAddressesDeletionFlags - updates notes deletion flags
func (s *GrpcClientsServer) UpdateAddressesDeletionFlags(ctx context.Context, req *api.RequestAddressesDeletionFlags) (*api.ResponseAddresses, error) {

	lgr.LOG.Info("***")
	lgr.LOG.Info("-->> ", "services.UpdateAddressesDeletionFlags()")

	// Check auth token
	lgr.LOG.Info("_ACTION_: ", "checking auth token")
	if !hlp.AuthTokenIsValid(req.AuthToken) {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(101))
		return &api.ResponseAddresses{
			Addresses: nil,
			Status:    hlp.GetStatusForResponse(st.GetStatus(101)),
		}, nil
	}

	lgr.LOG.Info("_ACTION_: ", "checking users access rights")
	allow, stat := hlp.GetUserAccessRightForTable(ctx, &api.RequestUsersAccessRightsForTable{
		AuthToken: cfg.CFG.MicroServices["users_ms"].Token,
		UserId:    req.AuthorId,
		TableName: notesTableName,
		Action:    api.UsersAccessRights_Actions_CAN_DELETE,
	})

	//error
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return &api.ResponseAddresses{
			Addresses: nil,
			Status:    hlp.GetStatusForResponse(stat),
		}, nil
	}

	//not allow
	if !allow {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(204))
		return &api.ResponseAddresses{
			Addresses: nil,
			Status:    hlp.GetStatusForResponse(st.GetStatus(204)),
		}, nil
	}

	// Work with storage
	lgr.LOG.Info("_ACTION_: ", "working with storage")
	addresses, status := stg.UpdateAddressesDeletionFlags(ctx, req)

	//Result
	lgr.LOG.Info("_RESULT_: ", status)
	lgr.LOG.Info("<<-- ", "services.UpdateAddressesDeletionFlags()")
	lgr.LOG.Info("***")
	return &api.ResponseAddresses{
		Addresses: addresses,
		Status:    hlp.GetStatusForResponse(status),
	}, nil
}
