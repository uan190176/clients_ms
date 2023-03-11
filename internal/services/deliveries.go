package services

import (
	"clients_ms/internal/api"
	hlp "clients_ms/internal/helpers"
	stg "clients_ms/internal/storage/deliveries"
	cfg "clients_ms/pkg/config"
	lgr "clients_ms/pkg/logger"
	"context"
	st "github.com/uan190176/statuses"
)

// GetDeliveries - returns deliveries
func (s *GrpcClientsServer) GetDeliveries(ctx context.Context, req *api.RequestDelivery) (*api.ResponseDeliveries, error) {

	lgr.LOG.Info("***")
	lgr.LOG.Info("-->> ", "services.GetDeliveries()")

	//Check auth token
	lgr.LOG.Info("_ACTION_: ", "checking auth token")
	if !hlp.AuthTokenIsValid(req.AuthToken) {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(101))
		return &api.ResponseDeliveries{
			Deliveries: nil,
			Status:     hlp.GetStatusForResponse(st.GetStatus(101)),
		}, nil
	}

	//Check users access rights
	lgr.LOG.Info("_ACTION_: ", "checking users access rights")
	allow, stat := hlp.GetUserAccessRightForTable(ctx, &api.RequestUsersAccessRightsForTable{
		AuthToken: cfg.CFG.MicroServices["users_ms"].Token,
		UserId:    req.AuthorId,
		TableName: "public.deliveries_ref",
		Action:    api.UsersAccessRights_Actions_CAN_SELECT,
	})

	//error
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return &api.ResponseDeliveries{
			Deliveries: nil,
			Status:     hlp.GetStatusForResponse(stat),
		}, nil
	}

	//not allow
	if !allow {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(201))
		return &api.ResponseDeliveries{
			Deliveries: nil,
			Status:     hlp.GetStatusForResponse(st.GetStatus(201)),
		}, nil
	}

	// Work with storage
	lgr.LOG.Info("_ACTION_: ", "working with storage")
	deliveries, status := stg.GetDeliveries(ctx, req)

	lgr.LOG.Info("_RESULT_: ", status)
	lgr.LOG.Info("-->> ", "services.GetDeliveries()")
	lgr.LOG.Info("***")

	return &api.ResponseDeliveries{
		Deliveries: deliveries,
		Status:     hlp.GetStatusForResponse(status),
	}, nil
}

// CreateDelivery - creates new country
func (s *GrpcClientsServer) CreateDelivery(ctx context.Context, req *api.RequestDelivery) (*api.ResponseDeliveries, error) {

	lgr.LOG.Info("***")
	lgr.LOG.Info("-->> ", "services.CreateDelivery()")

	// Check auth token
	lgr.LOG.Info("_ACTION_: ", "checking auth token")
	if !hlp.AuthTokenIsValid(req.AuthToken) {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(101))
		return &api.ResponseDeliveries{
			Deliveries: nil,
			Status:     hlp.GetStatusForResponse(st.GetStatus(101)),
		}, nil
	}

	lgr.LOG.Info("_ACTION_: ", "checking users access rights")
	allow, stat := hlp.GetUserAccessRightForTable(ctx, &api.RequestUsersAccessRightsForTable{
		AuthToken: cfg.CFG.MicroServices["users_ms"].Token,
		UserId:    req.AuthorId,
		TableName: "public.deliveries_ref",
		Action:    api.UsersAccessRights_Actions_CAN_INSERT,
	})

	//error
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return &api.ResponseDeliveries{
			Deliveries: nil,
			Status:     hlp.GetStatusForResponse(stat),
		}, nil
	}

	//not allow
	if !allow {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(202))
		return &api.ResponseDeliveries{
			Deliveries: nil,
			Status:     hlp.GetStatusForResponse(st.GetStatus(202)),
		}, nil
	}

	// Work with storage
	lgr.LOG.Info("_ACTION_: ", "working with storage")
	deliveries, status := stg.CreateDelivery(ctx, req)

	//Result
	lgr.LOG.Info("_RESULT_: ", status)
	lgr.LOG.Info("<<-- ", "services.CreateDelivery()")
	lgr.LOG.Info("***")
	return &api.ResponseDeliveries{
		Deliveries: deliveries,
		Status:     hlp.GetStatusForResponse(status),
	}, nil
}

// UpdateDelivery - updates country
func (s *GrpcClientsServer) UpdateDelivery(ctx context.Context, req *api.RequestDelivery) (*api.ResponseDeliveries, error) {

	lgr.LOG.Info("***")
	lgr.LOG.Info("-->> ", "services.UpdateDelivery()")

	// Check auth token
	lgr.LOG.Info("_ACTION_: ", "checking auth token")
	if !hlp.AuthTokenIsValid(req.AuthToken) {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(101))
		return &api.ResponseDeliveries{
			Deliveries: nil,
			Status:     hlp.GetStatusForResponse(st.GetStatus(101)),
		}, nil
	}

	lgr.LOG.Info("_ACTION_: ", "checking users access rights")
	allow, stat := hlp.GetUserAccessRightForTable(ctx, &api.RequestUsersAccessRightsForTable{
		AuthToken: cfg.CFG.MicroServices["users_ms"].Token,
		UserId:    req.AuthorId,
		TableName: "public.deliveries_ref",
		Action:    api.UsersAccessRights_Actions_CAN_UPDATE,
	})

	//error
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return &api.ResponseDeliveries{
			Deliveries: nil,
			Status:     hlp.GetStatusForResponse(stat),
		}, nil
	}

	//not allow
	if !allow {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(203))
		return &api.ResponseDeliveries{
			Deliveries: nil,
			Status:     hlp.GetStatusForResponse(st.GetStatus(203)),
		}, nil
	}

	// Work with storage
	lgr.LOG.Info("_ACTION_: ", "working with storage")
	deliveries, status := stg.UpdateDelivery(ctx, req)

	//Result
	lgr.LOG.Info("_RESULT_: ", status)
	lgr.LOG.Info("<<-- ", "services.UpdateDelivery()")
	lgr.LOG.Info("***")
	return &api.ResponseDeliveries{
		Deliveries: deliveries,
		Status:     hlp.GetStatusForResponse(status),
	}, nil
}

// UpdateDeliveriesDeletionFlags - updates countries deletion flags
func (s *GrpcClientsServer) UpdateDeliveriesDeletionFlags(ctx context.Context, req *api.RequestDeliveriesDeletionFlags) (*api.ResponseDeliveries, error) {

	lgr.LOG.Info("***")
	lgr.LOG.Info("-->> ", "services.UpdateDeliveriesDeletionFlags()")

	// Check auth token
	lgr.LOG.Info("_ACTION_: ", "checking auth token")
	if !hlp.AuthTokenIsValid(req.AuthToken) {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(101))
		return &api.ResponseDeliveries{
			Deliveries: nil,
			Status:     hlp.GetStatusForResponse(st.GetStatus(101)),
		}, nil
	}

	lgr.LOG.Info("_ACTION_: ", "checking users access rights")
	allow, stat := hlp.GetUserAccessRightForTable(ctx, &api.RequestUsersAccessRightsForTable{
		AuthToken: cfg.CFG.MicroServices["users_ms"].Token,
		UserId:    req.AuthorId,
		TableName: "public.deliveries_ref",
		Action:    api.UsersAccessRights_Actions_CAN_DELETE,
	})

	//error
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return &api.ResponseDeliveries{
			Deliveries: nil,
			Status:     hlp.GetStatusForResponse(stat),
		}, nil
	}

	//not allow
	if !allow {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(204))
		return &api.ResponseDeliveries{
			Deliveries: nil,
			Status:     hlp.GetStatusForResponse(st.GetStatus(204)),
		}, nil
	}

	// Work with storage
	lgr.LOG.Info("_ACTION_: ", "working with storage")
	countries, status := stg.UpdateDeliveriesDeletionFlags(ctx, req)

	//Result
	lgr.LOG.Info("_RESULT_: ", status)
	lgr.LOG.Info("<<-- ", "services.UpdateDeliveriesDeletionFlags()")
	lgr.LOG.Info("***")
	return &api.ResponseDeliveries{
		Deliveries: countries,
		Status:     hlp.GetStatusForResponse(status),
	}, nil
}
