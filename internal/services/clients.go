package services

import (
	"clients_ms/internal/api"
	hlp "clients_ms/internal/helpers"
	stg "clients_ms/internal/storage/clients"
	cfg "clients_ms/pkg/config"
	lgr "clients_ms/pkg/logger"
	"context"
	st "github.com/uan190176/statuses"
)

const clientsTableName = "public.clients_ref"

// GetClients - returns clients
func (s *GrpcClientsServer) GetClients(ctx context.Context, req *api.RequestClient) (*api.ResponseClients, error) {

	lgr.LOG.Info("***")
	lgr.LOG.Info("-->> ", "services.GetClients()")

	//Check auth token
	lgr.LOG.Info("_ACTION_: ", "checking auth token")
	if !hlp.AuthTokenIsValid(req.AuthToken) {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(101))
		return &api.ResponseClients{
			Clients: nil,
			Status:  hlp.GetStatusForResponse(st.GetStatus(101)),
		}, nil
	}

	//Check users access rights
	lgr.LOG.Info("_ACTION_: ", "checking users access rights")
	allow, stat := hlp.GetUserAccessRightForTable(ctx, &api.RequestUsersAccessRightsForTable{
		AuthToken: cfg.CFG.MicroServices["users_ms"].Token,
		UserId:    req.AuthorId,
		TableName: clientsTableName,
		Action:    api.UsersAccessRights_Actions_CAN_SELECT,
	})

	//error
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return &api.ResponseClients{
			Clients: nil,
			Status:  hlp.GetStatusForResponse(stat),
		}, nil
	}

	//not allow
	if !allow {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(201))
		return &api.ResponseClients{
			Clients: nil,
			Status:  hlp.GetStatusForResponse(st.GetStatus(201)),
		}, nil
	}

	// Work with storage
	lgr.LOG.Info("_ACTION_: ", "working with storage")
	clients, status := stg.GetClients(ctx, req)

	lgr.LOG.Info("_RESULT_: ", status)
	lgr.LOG.Info("-->> ", "services.GetClients()")
	lgr.LOG.Info("***")

	return &api.ResponseClients{
		Clients: clients,
		Status:  hlp.GetStatusForResponse(status),
	}, nil
}

// CreateClient - creates new client
func (s *GrpcClientsServer) CreateClient(ctx context.Context, req *api.RequestClient) (*api.ResponseClients, error) {

	lgr.LOG.Info("***")
	lgr.LOG.Info("-->> ", "services.CreateClient()")

	// Check auth token
	lgr.LOG.Info("_ACTION_: ", "checking auth token")
	if !hlp.AuthTokenIsValid(req.AuthToken) {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(101))
		return &api.ResponseClients{
			Clients: nil,
			Status:  hlp.GetStatusForResponse(st.GetStatus(101)),
		}, nil
	}

	// Check users access rights
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
		return &api.ResponseClients{
			Clients: nil,
			Status:  hlp.GetStatusForResponse(stat),
		}, nil
	}

	//not allow
	if !allow {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(202))
		return &api.ResponseClients{
			Clients: nil,
			Status:  hlp.GetStatusForResponse(st.GetStatus(202)),
		}, nil
	}

	// Work with storage
	lgr.LOG.Info("_ACTION_: ", "working with storage")
	clients, status := stg.CreateClient(ctx, req)

	//Result
	lgr.LOG.Info("_RESULT_: ", status)
	lgr.LOG.Info("<<-- ", "services.CreateClient()")
	lgr.LOG.Info("***")
	return &api.ResponseClients{
		Clients: clients,
		Status:  hlp.GetStatusForResponse(status),
	}, nil
}

// UpdateClient - updates client
func (s *GrpcClientsServer) UpdateClient(ctx context.Context, req *api.RequestClient) (*api.ResponseClients, error) {

	lgr.LOG.Info("***")
	lgr.LOG.Info("-->> ", "services.UpdateClient()")

	// Check auth token
	lgr.LOG.Info("_ACTION_: ", "checking auth token")
	if !hlp.AuthTokenIsValid(req.AuthToken) {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(101))
		return &api.ResponseClients{
			Clients: nil,
			Status:  hlp.GetStatusForResponse(st.GetStatus(101)),
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
		return &api.ResponseClients{
			Clients: nil,
			Status:  hlp.GetStatusForResponse(stat),
		}, nil
	}

	//not allow
	if !allow {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(203))
		return &api.ResponseClients{
			Clients: nil,
			Status:  hlp.GetStatusForResponse(st.GetStatus(203)),
		}, nil
	}

	// Work with storage
	lgr.LOG.Info("_ACTION_: ", "working with storage")
	clients, status := stg.UpdateClient(ctx, req)

	//Result
	lgr.LOG.Info("_RESULT_: ", status)
	lgr.LOG.Info("<<-- ", "services.UpdateClient()")
	lgr.LOG.Info("***")
	return &api.ResponseClients{
		Clients: clients,
		Status:  hlp.GetStatusForResponse(status),
	}, nil
}

// UpdateClientsDeletionFlags - updates notes deletion flags
func (s *GrpcClientsServer) UpdateClientsDeletionFlags(ctx context.Context, req *api.RequestClientsDeletionFlags) (*api.ResponseClients, error) {

	lgr.LOG.Info("***")
	lgr.LOG.Info("-->> ", "services.UpdateClientsDeletionFlags()")

	// Check auth token
	lgr.LOG.Info("_ACTION_: ", "checking auth token")
	if !hlp.AuthTokenIsValid(req.AuthToken) {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(101))
		return &api.ResponseClients{
			Clients: nil,
			Status:  hlp.GetStatusForResponse(st.GetStatus(101)),
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
		return &api.ResponseClients{
			Clients: nil,
			Status:  hlp.GetStatusForResponse(stat),
		}, nil
	}

	//not allow
	if !allow {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(204))
		return &api.ResponseClients{
			Clients: nil,
			Status:  hlp.GetStatusForResponse(st.GetStatus(204)),
		}, nil
	}

	// Work with storage
	lgr.LOG.Info("_ACTION_: ", "working with storage")
	clients, status := stg.UpdateClientsDeletionFlags(ctx, req)

	//Result
	lgr.LOG.Info("_RESULT_: ", status)
	lgr.LOG.Info("<<-- ", "services.UpdateClientsDeletionFlags()")
	lgr.LOG.Info("***")
	return &api.ResponseClients{
		Clients: clients,
		Status:  hlp.GetStatusForResponse(status),
	}, nil
}
