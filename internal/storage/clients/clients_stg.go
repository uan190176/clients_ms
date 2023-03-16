package clients

import (
	"clients_ms/internal/api"
	cfg "clients_ms/pkg/config"
	lgr "clients_ms/pkg/logger"
	"context"
	fs "github.com/uan190176/fstorage"
	st "github.com/uan190176/statuses"
)

// GetClients - returns clients
func GetClients(ctx context.Context, req *api.RequestClient) ([]*api.Client, st.ResponseStatus) {

	lgr.LOG.Info("-->> ", "clients.GetClients()")

	// Vars
	var (
		s       *fs.Storage
		q       string
		clients []*api.Client
		err     error
	)

	//Generate query
	lgr.LOG.Info("_ACTION_: ", "generate query")
	q = GenQueryClientsSelect(req)

	//Create storage parameters
	lgr.LOG.Info("_ACTION_: ", "creating storage parameters")
	sp := &fs.StorageParams{
		Ctx:        ctx,
		Dsn:        cfg.CFG.Database.URL,
		AppName:    "ClientsMS::GetClients()",
		AutoCommit: true,
		AutoClose:  true,
	}

	//Create query parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	qp := &fs.QueryParams{
		Dest:        &clients,
		QueryString: q,
	}

	//Open storage
	lgr.LOG.Info("_ACTION_: ", "opening connection")
	s, err = fs.NewStorage(sp)
	if err != nil {
		lgr.LOG.Error("_ERR_: ", err)
		return nil, st.GetStatus(0, err.Error())
	}

	//Execute query
	lgr.LOG.Info("_ACTION_: ", "running query")
	err = s.RunQuery(qp)
	if err != nil {
		lgr.LOG.Error("_ERR_: ", err)
		return nil, st.GetStatus(0, err.Error())
	}

	//Check result
	if len(clients) > 0 {
		lgr.LOG.Info("_RESULT_: ", clients)
		lgr.LOG.Info("<<-- ", "clients.GetClients()")
		return clients, st.GetStatus(100)
	}

	lgr.LOG.Warn("_WARN_: ", st.GetStatus(301))
	return nil, st.GetStatus(301)
}

// CreateClient - creates new client
func CreateClient(ctx context.Context, req *api.RequestClient) ([]*api.Client, st.ResponseStatus) {

	lgr.LOG.Info("-->> ", "clients.CreateClient()")

	// Vars
	var (
		s       *fs.Storage
		q       string
		clients []*api.Client
		err     error
		stat    st.ResponseStatus
	)

	//Check required fields
	lgr.LOG.Info("_ACTION_: ", "checking required fields")
	stat = CheckAndNormalizeRequiredFieldsClientInsert(req)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return nil, stat
	}

	//Create query
	lgr.LOG.Info("_ACTION_: ", "generating query")
	q = CreateQueryClientsInsert(req)

	//Create storage parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	sp := &fs.StorageParams{
		Ctx:        ctx,
		Dsn:        cfg.CFG.Database.URL,
		AppName:    "ClientsMS::CreateClient()",
		AutoCommit: true,
		AutoClose:  true,
	}

	//Create query parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	qp := &fs.QueryParams{
		Dest:        &clients,
		QueryString: q,
	}

	//Open storage
	lgr.LOG.Info("_ACTION_: ", "opening connection")
	s, err = fs.NewStorage(sp)
	if err != nil {
		lgr.LOG.Error("_ERR_: ", err)
		return nil, st.GetStatus(0, err.Error())
	}

	//Execute query
	lgr.LOG.Info("_ACTION_: ", "running query")
	err = s.RunQuery(qp)
	if err != nil {
		lgr.LOG.Error("_ERR_: ", err)
		return nil, st.GetStatus(0, err.Error())
	}

	//Check result
	if len(clients) > 0 {
		lgr.LOG.Info("_RESULT_: ", clients)
		lgr.LOG.Info("<<-- ", "clients.CreateClient()")
		return clients, st.GetStatus(100)
	}

	lgr.LOG.Warn("_WARN_: ", st.GetStatus(301))
	return nil, st.GetStatus(301)
}

// UpdateClient - creates new client
func UpdateClient(ctx context.Context, req *api.RequestClient) ([]*api.Client, st.ResponseStatus) {

	lgr.LOG.Info("-->> ", "clients.UpdateClient()")

	// Vars
	var (
		s       *fs.Storage
		q       string
		clients []*api.Client
		err     error
		stat    st.ResponseStatus
	)

	//Check required fields
	lgr.LOG.Info("_ACTION_: ", "checking required fields")
	stat = CheckAndNormalizeRequiredFieldsClientUpdate(req)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return nil, stat
	}

	//Create query
	lgr.LOG.Info("_ACTION_: ", "generating query")
	q = CreateQueryClientsUpdate(req)

	//Create storage parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	sp := &fs.StorageParams{
		Ctx:        ctx,
		Dsn:        cfg.CFG.Database.URL,
		AppName:    "ClientsMS::UpdateClient()",
		AutoCommit: true,
		AutoClose:  true,
	}

	//Create query parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	qp := &fs.QueryParams{
		Dest:        &clients,
		QueryString: q,
	}

	//Open storage
	lgr.LOG.Info("_ACTION_: ", "opening connection")
	s, err = fs.NewStorage(sp)
	if err != nil {
		lgr.LOG.Error("_ERR_: ", err)
		return nil, st.GetStatus(0, err.Error())
	}

	//Execute query
	lgr.LOG.Info("_ACTION_: ", "running query")
	err = s.RunQuery(qp)
	if err != nil {
		lgr.LOG.Error("_ERR_: ", err)
		return nil, st.GetStatus(0, err.Error())
	}

	//Check result
	if len(clients) > 0 {
		lgr.LOG.Info("_RESULT_: ", clients)
		lgr.LOG.Info("<<-- ", "clients.UpdateClient()")
		return clients, st.GetStatus(100)
	}

	lgr.LOG.Warn("_WARN_: ", st.GetStatus(301))
	return nil, st.GetStatus(301)
}

// UpdateClientsDeletionFlags - updates deletion flags
func UpdateClientsDeletionFlags(ctx context.Context, req *api.RequestClientsDeletionFlags) ([]*api.Client, st.ResponseStatus) {

	lgr.LOG.Info("-->> ", "clients.UpdateClientsDeletionFlags()")

	// Vars
	var (
		s       *fs.Storage
		q       string
		clients []*api.Client
		err     error
		stat    st.ResponseStatus
	)

	//Check required fields
	lgr.LOG.Info("_ACTION_: ", "checking required fields")
	stat = CheckRequiredFieldsClientsDeletionFlagUpdate(req)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return nil, stat
	}

	//Create query
	lgr.LOG.Info("_ACTION_: ", "generating query")
	q = CreateQueryClientsDeletionFlagsUpdate(req)

	//Create storage parameters
	lgr.LOG.Info("_ACTION_: ", "creating storage parameters")
	sp := &fs.StorageParams{
		Ctx:        ctx,
		Dsn:        cfg.CFG.Database.URL,
		AppName:    "ClientsMS::UpdateClientsDeletionFlags()",
		AutoCommit: true,
		AutoClose:  true,
	}

	//Create query parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	qp := &fs.QueryParams{
		Dest:        &clients,
		QueryString: q,
	}

	//Open storage
	lgr.LOG.Info("_ACTION_: ", "opening connection")
	s, err = fs.NewStorage(sp)
	if err != nil {
		lgr.LOG.Error("_ERR_: ", err)
		return nil, st.GetStatus(0, err.Error())
	}

	//Execute query
	lgr.LOG.Info("_ACTION_: ", "running query")
	err = s.RunQuery(qp)
	if err != nil {
		lgr.LOG.Error("_ERR_: ", err)
		return nil, st.GetStatus(0, err.Error())
	}

	//Check result
	if len(clients) > 0 {
		lgr.LOG.Info("_RESULT_: ", clients)
		lgr.LOG.Info("<<-- ", "clients.UpdateClientsDeletionFlags()")
		return clients, st.GetStatus(100)
	}

	lgr.LOG.Warn("_WARN_: ", st.GetStatus(301))
	return nil, st.GetStatus(301)
}
