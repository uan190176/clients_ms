package countries

import (
	"clients_ms/internal/api"
	cfg "clients_ms/pkg/config"
	"clients_ms/pkg/logger"
	"context"
	fs "github.com/uan190176/fstorage"
	st "github.com/uan190176/statuses"
)

// GetCountries - returns countries
func GetCountries(ctx context.Context, req *api.RequestCountry) ([]*api.Country, st.ResponseStatus) {

	lgr.LOG.Info("-->> ", "countries.GetCountries()")

	// Vars
	var (
		s         *fs.Storage
		q         string
		countries []*api.Country
		err       error
	)

	//Generate query
	lgr.LOG.Info("_ACTION_: ", "generate query")
	q = GenQueryCountriesSelect(req)

	//Create storage parameters
	lgr.LOG.Info("_ACTION_: ", "creating storage parameters")
	sp := &fs.StorageParams{
		Ctx:        ctx,
		Dsn:        cfg.CFG.Database.URL,
		AppName:    "ClientsMS::GetCountries()",
		AutoCommit: true,
		AutoClose:  true,
	}

	//Create query parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	qp := &fs.QueryParams{
		Dest:        &countries,
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
	if len(countries) > 0 {
		lgr.LOG.Info("_RESULT_: ", countries)
		lgr.LOG.Info("<<-- ", "countries.GetCountries()")
		return countries, st.GetStatus(100)
	}

	lgr.LOG.Warn("_WARN_: ", st.GetStatus(301))
	return nil, st.GetStatus(301)
}

// CreateCountry - creates new country
func CreateCountry(ctx context.Context, req *api.RequestCountry) ([]*api.Country, st.ResponseStatus) {

	lgr.LOG.Info("-->> ", "countries.CreateCountry()")

	// Vars
	var (
		s         *fs.Storage
		q         string
		countries []*api.Country
		err       error
		stat      st.ResponseStatus
	)

	//Check required fields
	lgr.LOG.Info("_ACTION_: ", "checking required fields")
	stat = CheckAndNormalizeRequiredFieldsCountryInsert(req)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return nil, stat
	}

	//Create query
	lgr.LOG.Info("_ACTION_: ", "generating query")
	q, stat = CreateQueryCountryInsert(req)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return nil, stat
	}

	//Create storage parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	sp := &fs.StorageParams{
		Ctx:        ctx,
		Dsn:        cfg.CFG.Database.URL,
		AppName:    "ClientsMS::CreateCountry()",
		AutoCommit: true,
		AutoClose:  true,
	}

	//Create query parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	qp := &fs.QueryParams{
		Dest:        &countries,
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
	if len(countries) > 0 {
		lgr.LOG.Info("_RESULT_: ", countries)
		lgr.LOG.Info("<<-- ", "countries.CreateCountry()")
		return countries, st.GetStatus(100)
	}

	lgr.LOG.Warn("_WARN_: ", st.GetStatus(301))
	return nil, st.GetStatus(301)
}

// UpdateCountry - updates country
func UpdateCountry(ctx context.Context, req *api.RequestCountry) ([]*api.Country, st.ResponseStatus) {

	lgr.LOG.Info("-->> ", "countries.UpdateCountry()")

	// Vars
	var (
		s         *fs.Storage
		q         string
		countries []*api.Country
		err       error
		stat      st.ResponseStatus
	)

	//Check required fields
	lgr.LOG.Info("_ACTION_: ", "checking required fields")
	stat = CheckAndNormalizeRequiredFieldsCountryUpdate(req)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return nil, stat
	}

	//Create query
	lgr.LOG.Info("_ACTION_: ", "generating query")
	q, stat = CreateQueryCountryUpdate(req)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return nil, stat
	}

	//Create storage parameters
	lgr.LOG.Info("_ACTION_: ", "creating storage parameters")
	sp := &fs.StorageParams{
		Ctx:        ctx,
		Dsn:        cfg.CFG.Database.URL,
		AppName:    "ClientsMS::UpdateCountry()",
		AutoCommit: true,
		AutoClose:  true,
	}

	//Create query parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	qp := &fs.QueryParams{
		Dest:        &countries,
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
	if len(countries) > 0 {
		lgr.LOG.Info("_RESULT_: ", countries)
		lgr.LOG.Info("<<-- ", "countries.UpdateCountry()")
		return countries, st.GetStatus(100)
	}

	lgr.LOG.Warn("_WARN_: ", st.GetStatus(301))
	return nil, st.GetStatus(301)
}

// UpdateCountriesDeletionFlags - updates countries deletion flags
func UpdateCountriesDeletionFlags(ctx context.Context, req *api.RequestCountriesDeletionFlags) ([]*api.Country, st.ResponseStatus) {

	lgr.LOG.Info("-->> ", "countries.UpdateCountriesDeletionFlags()")

	// Vars
	var (
		s         *fs.Storage
		q         string
		countries []*api.Country
		err       error
		stat      st.ResponseStatus
	)

	//Check required fields
	lgr.LOG.Info("_ACTION_: ", "checking required fields")
	stat = CheckRequiredFieldsCountriesDeletionFlagUpdate(req)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return nil, stat
	}

	//Create query
	lgr.LOG.Info("_ACTION_: ", "generating query")
	q = CreateQueryCountriesDeletionFlagsUpdate(req)

	//Create storage parameters
	lgr.LOG.Info("_ACTION_: ", "creating storage parameters")
	sp := &fs.StorageParams{
		Ctx:        ctx,
		Dsn:        cfg.CFG.Database.URL,
		AppName:    "ClientsMS::UpdateCountriesDeletionFlags()",
		AutoCommit: true,
		AutoClose:  true,
	}

	//Create query parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	qp := &fs.QueryParams{
		Dest:        &countries,
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
	if len(countries) > 0 {
		lgr.LOG.Info("_RESULT_: ", countries)
		lgr.LOG.Info("<<-- ", "countries.UpdateCountriesDeletionFlags()")
		return countries, st.GetStatus(100)
	}

	lgr.LOG.Warn("_WARN_: ", st.GetStatus(301))
	return nil, st.GetStatus(301)
}
