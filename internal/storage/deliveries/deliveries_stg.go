package deliveries

import (
	"clients_ms/internal/api"
	cfg "clients_ms/pkg/config"
	lgr "clients_ms/pkg/logger"
	"context"
	fs "github.com/uan190176/fstorage"
	st "github.com/uan190176/statuses"
)

// GetDeliveries - returns deliveries
func GetDeliveries(ctx context.Context, req *api.RequestDelivery) ([]*api.Delivery, st.ResponseStatus) {

	lgr.LOG.Info("-->> ", "deliveries.GetCountries()")

	// Vars
	var (
		s          *fs.Storage
		q          string
		deliveries []*api.Delivery
		err        error
	)

	//Generate query
	lgr.LOG.Info("_ACTION_: ", "generate query")
	q = GenQueryDeliveriesSelect(req)

	//Create storage parameters
	lgr.LOG.Info("_ACTION_: ", "creating storage parameters")
	sp := &fs.StorageParams{
		Ctx:        ctx,
		Dsn:        cfg.CFG.Database.URL,
		AppName:    "ClientsMS::GetDeliveries()",
		AutoCommit: true,
		AutoClose:  true,
	}

	//Create query parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	qp := &fs.QueryParams{
		Dest:        &deliveries,
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
	if len(deliveries) > 0 {
		lgr.LOG.Info("_RESULT_: ", deliveries)
		lgr.LOG.Info("<<-- ", "deliveries.GetCountries()")
		return deliveries, st.GetStatus(100)
	}

	lgr.LOG.Warn("_WARN_: ", st.GetStatus(301))
	return nil, st.GetStatus(301)
}

// CreateDelivery - creates new delivery
func CreateDelivery(ctx context.Context, req *api.RequestDelivery) ([]*api.Delivery, st.ResponseStatus) {

	lgr.LOG.Info("-->> ", "deliveries.CreateDelivery()")

	// Vars
	var (
		s          *fs.Storage
		q          string
		deliveries []*api.Delivery
		err        error
		stat       st.ResponseStatus
	)

	//Check required fields
	lgr.LOG.Info("_ACTION_: ", "checking required fields")
	stat = CheckRequiredFieldsDeliveryInsert(req)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return nil, stat
	}

	//Create query
	lgr.LOG.Info("_ACTION_: ", "generating query")
	q, stat = CreateQueryDeliveryInsert(req)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return nil, stat
	}

	//Create storage parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	sp := &fs.StorageParams{
		Ctx:        ctx,
		Dsn:        cfg.CFG.Database.URL,
		AppName:    "ClientsMS::CreateDelivery()",
		AutoCommit: true,
		AutoClose:  true,
	}

	//Create query parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	qp := &fs.QueryParams{
		Dest:        &deliveries,
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
	if len(deliveries) > 0 {
		lgr.LOG.Info("_RESULT_: ", deliveries)
		lgr.LOG.Info("<<-- ", "deliveries.CreateDelivery()")
		return deliveries, st.GetStatus(100)
	}

	lgr.LOG.Warn("_WARN_: ", st.GetStatus(301))
	return nil, st.GetStatus(301)
}

// UpdateDelivery - updates delivery
func UpdateDelivery(ctx context.Context, req *api.RequestDelivery) ([]*api.Delivery, st.ResponseStatus) {

	lgr.LOG.Info("-->> ", "deliveries.UpdateDelivery()")

	// Vars
	var (
		s          *fs.Storage
		q          string
		deliveries []*api.Delivery
		err        error
		stat       st.ResponseStatus
	)

	//Check required fields
	lgr.LOG.Info("_ACTION_: ", "checking required fields")
	stat = CheckRequiredFieldsDeliveryUpdate(req)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return nil, stat
	}

	//Create query
	lgr.LOG.Info("_ACTION_: ", "generating query")
	q, stat = CreateQueryDeliveryUpdate(req)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return nil, stat
	}

	//Create storage parameters
	lgr.LOG.Info("_ACTION_: ", "creating storage parameters")
	sp := &fs.StorageParams{
		Ctx:        ctx,
		Dsn:        cfg.CFG.Database.URL,
		AppName:    "ClientsMS::UpdateDelivery()",
		AutoCommit: true,
		AutoClose:  true,
	}

	//Create query parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	qp := &fs.QueryParams{
		Dest:        &deliveries,
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
	if len(deliveries) > 0 {
		lgr.LOG.Info("_RESULT_: ", deliveries)
		lgr.LOG.Info("<<-- ", "deliveries.UpdateDelivery()")
		return deliveries, st.GetStatus(100)
	}

	lgr.LOG.Warn("_WARN_: ", st.GetStatus(301))
	return nil, st.GetStatus(301)
}

// UpdateDeliveriesDeletionFlags - updates deletion flags
func UpdateDeliveriesDeletionFlags(ctx context.Context, req *api.RequestDeliveriesDeletionFlags) ([]*api.Delivery, st.ResponseStatus) {

	lgr.LOG.Info("-->> ", "deliveries.UpdateDeliveriesDeletionFlags()")

	// Vars
	var (
		s          *fs.Storage
		q          string
		deliveries []*api.Delivery
		err        error
		stat       st.ResponseStatus
	)

	//Check required fields
	lgr.LOG.Info("_ACTION_: ", "checking required fields")
	stat = CheckRequiredFieldsDeliveriesDeletionFlagUpdate(req)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return nil, stat
	}

	//Create query
	lgr.LOG.Info("_ACTION_: ", "generating query")
	q = CreateQueryDeliveriesDeletionFlagsUpdate(req)

	//Create storage parameters
	lgr.LOG.Info("_ACTION_: ", "creating storage parameters")
	sp := &fs.StorageParams{
		Ctx:        ctx,
		Dsn:        cfg.CFG.Database.URL,
		AppName:    "ClientsMS::UpdateDeliveriesDeletionFlags()",
		AutoCommit: true,
		AutoClose:  true,
	}

	//Create query parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	qp := &fs.QueryParams{
		Dest:        &deliveries,
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
	if len(deliveries) > 0 {
		lgr.LOG.Info("_RESULT_: ", deliveries)
		lgr.LOG.Info("<<-- ", "deliveries.UpdateDeliveriesDeletionFlags()")
		return deliveries, st.GetStatus(100)
	}

	lgr.LOG.Warn("_WARN_: ", st.GetStatus(301))
	return nil, st.GetStatus(301)
}
