package services

import (
	"clients_ms/internal/api"
	hlp "clients_ms/internal/helpers"
	stg "clients_ms/internal/storage/countries"
	cfg "clients_ms/pkg/config"
	lgr "clients_ms/pkg/logger"
	"context"
	st "github.com/uan190176/statuses"
)

const countriesTableName = "public.countries_ref"

// GetCountries - returns countries
func (s *GrpcClientsServer) GetCountries(ctx context.Context, req *api.RequestCountry) (*api.ResponseCountries, error) {

	lgr.LOG.Info("***")
	lgr.LOG.Info("-->> ", "services.GetCountries()")

	//Check auth token
	lgr.LOG.Info("_ACTION_: ", "checking auth token")
	if !hlp.AuthTokenIsValid(req.AuthToken) {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(101))
		return &api.ResponseCountries{
			Countries: nil,
			Status:    hlp.GetStatusForResponse(st.GetStatus(101)),
		}, nil
	}

	//Check users access rights
	lgr.LOG.Info("_ACTION_: ", "checking users access rights")
	allow, stat := hlp.GetUserAccessRightForTable(ctx, &api.RequestUsersAccessRightsForTable{
		AuthToken: cfg.CFG.MicroServices["users_ms"].Token,
		UserId:    req.AuthorId,
		TableName: countriesTableName,
		Action:    api.UsersAccessRights_Actions_CAN_SELECT,
	})

	//error
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return &api.ResponseCountries{
			Countries: nil,
			Status:    hlp.GetStatusForResponse(stat),
		}, nil
	}

	//not allow
	if !allow {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(201))
		return &api.ResponseCountries{
			Countries: nil,
			Status:    hlp.GetStatusForResponse(st.GetStatus(201)),
		}, nil
	}

	// Work with storage
	lgr.LOG.Info("_ACTION_: ", "working with storage")
	countries, status := stg.GetCountries(ctx, req)

	lgr.LOG.Info("_RESULT_: ", status)
	lgr.LOG.Info("-->> ", "services.GetCountries()")
	lgr.LOG.Info("***")

	return &api.ResponseCountries{
		Countries: countries,
		Status:    hlp.GetStatusForResponse(status),
	}, nil
}

// CreateCountry - creates new country
func (s *GrpcClientsServer) CreateCountry(ctx context.Context, req *api.RequestCountry) (*api.ResponseCountries, error) {

	lgr.LOG.Info("***")
	lgr.LOG.Info("-->> ", "services.CreateCountry()")

	// Check auth token
	lgr.LOG.Info("_ACTION_: ", "checking auth token")
	if !hlp.AuthTokenIsValid(req.AuthToken) {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(101))
		return &api.ResponseCountries{
			Countries: nil,
			Status:    hlp.GetStatusForResponse(st.GetStatus(101)),
		}, nil
	}

	lgr.LOG.Info("_ACTION_: ", "checking users access rights")
	allow, stat := hlp.GetUserAccessRightForTable(ctx, &api.RequestUsersAccessRightsForTable{
		AuthToken: cfg.CFG.MicroServices["users_ms"].Token,
		UserId:    req.AuthorId,
		TableName: countriesTableName,
		Action:    api.UsersAccessRights_Actions_CAN_INSERT,
	})

	//error
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return &api.ResponseCountries{
			Countries: nil,
			Status:    hlp.GetStatusForResponse(stat),
		}, nil
	}

	//not allow
	if !allow {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(202))
		return &api.ResponseCountries{
			Countries: nil,
			Status:    hlp.GetStatusForResponse(st.GetStatus(202)),
		}, nil
	}

	// Work with storage
	lgr.LOG.Info("_ACTION_: ", "working with storage")
	countries, status := stg.CreateCountry(ctx, req)

	//Result
	lgr.LOG.Info("_RESULT_: ", status)
	lgr.LOG.Info("<<-- ", "services.CreateCountry()")
	lgr.LOG.Info("***")
	return &api.ResponseCountries{
		Countries: countries,
		Status:    hlp.GetStatusForResponse(status),
	}, nil
}

// UpdateCountry - updates country
func (s *GrpcClientsServer) UpdateCountry(ctx context.Context, req *api.RequestCountry) (*api.ResponseCountries, error) {

	lgr.LOG.Info("***")
	lgr.LOG.Info("-->> ", "services.UpdateCountry()")

	// Check auth token
	lgr.LOG.Info("_ACTION_: ", "checking auth token")
	if !hlp.AuthTokenIsValid(req.AuthToken) {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(101))
		return &api.ResponseCountries{
			Countries: nil,
			Status:    hlp.GetStatusForResponse(st.GetStatus(101)),
		}, nil
	}

	lgr.LOG.Info("_ACTION_: ", "checking users access rights")
	allow, stat := hlp.GetUserAccessRightForTable(ctx, &api.RequestUsersAccessRightsForTable{
		AuthToken: cfg.CFG.MicroServices["users_ms"].Token,
		UserId:    req.AuthorId,
		TableName: countriesTableName,
		Action:    api.UsersAccessRights_Actions_CAN_UPDATE,
	})

	//error
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return &api.ResponseCountries{
			Countries: nil,
			Status:    hlp.GetStatusForResponse(stat),
		}, nil
	}

	//not allow
	if !allow {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(203))
		return &api.ResponseCountries{
			Countries: nil,
			Status:    hlp.GetStatusForResponse(st.GetStatus(203)),
		}, nil
	}

	// Work with storage
	lgr.LOG.Info("_ACTION_: ", "working with storage")
	countries, status := stg.UpdateCountry(ctx, req)

	//Result
	lgr.LOG.Info("_RESULT_: ", status)
	lgr.LOG.Info("<<-- ", "services.UpdateCountry()")
	lgr.LOG.Info("***")
	return &api.ResponseCountries{
		Countries: countries,
		Status:    hlp.GetStatusForResponse(status),
	}, nil
}

// UpdateCountriesDeletionFlags - updates countries deletion flags
func (s *GrpcClientsServer) UpdateCountriesDeletionFlags(ctx context.Context, req *api.RequestCountriesDeletionFlags) (*api.ResponseCountries, error) {

	lgr.LOG.Info("***")
	lgr.LOG.Info("-->> ", "services.UpdateCountriesDeletionFlags()")

	// Check auth token
	lgr.LOG.Info("_ACTION_: ", "checking auth token")
	if !hlp.AuthTokenIsValid(req.AuthToken) {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(101))
		return &api.ResponseCountries{
			Countries: nil,
			Status:    hlp.GetStatusForResponse(st.GetStatus(101)),
		}, nil
	}

	lgr.LOG.Info("_ACTION_: ", "checking users access rights")
	allow, stat := hlp.GetUserAccessRightForTable(ctx, &api.RequestUsersAccessRightsForTable{
		AuthToken: cfg.CFG.MicroServices["users_ms"].Token,
		UserId:    req.AuthorId,
		TableName: countriesTableName,
		Action:    api.UsersAccessRights_Actions_CAN_DELETE,
	})

	//error
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return &api.ResponseCountries{
			Countries: nil,
			Status:    hlp.GetStatusForResponse(stat),
		}, nil
	}

	//not allow
	if !allow {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(204))
		return &api.ResponseCountries{
			Countries: nil,
			Status:    hlp.GetStatusForResponse(st.GetStatus(204)),
		}, nil
	}

	// Work with storage
	lgr.LOG.Info("_ACTION_: ", "working with storage")
	countries, status := stg.UpdateCountriesDeletionFlags(ctx, req)

	//Result
	lgr.LOG.Info("_RESULT_: ", status)
	lgr.LOG.Info("<<-- ", "services.UpdateCountriesDeletionFlags()")
	lgr.LOG.Info("***")
	return &api.ResponseCountries{
		Countries: countries,
		Status:    hlp.GetStatusForResponse(status),
	}, nil
}
