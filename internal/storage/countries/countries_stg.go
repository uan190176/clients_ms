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

	//Create query parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	qp := &fs.QueryParams{
		Ctx:         ctx,
		Dsn:         cfg.CFG.Database.URL,
		AppName:     "ClientsMS::GetCountries()",
		Dest:        &countries,
		QueryString: q,
		AutoClose:   true,
	}

	//Open storage
	lgr.LOG.Info("_ACTION_: ", "opening connection")
	s, err = fs.NewStorage(qp)
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

	//Create query parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	qp := &fs.QueryParams{
		Ctx:         ctx,
		Dsn:         cfg.CFG.Database.URL,
		AppName:     "ClientsMS::CreateCountry()",
		Dest:        &countries,
		QueryString: q,
		AutoClose:   true,
	}

	//Open storage
	lgr.LOG.Info("_ACTION_: ", "opening connection")
	s, err = fs.NewStorage(qp)
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

// UpdateCountry - creates new country
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

	//Create query parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	qp := &fs.QueryParams{
		Ctx:         ctx,
		Dsn:         cfg.CFG.Database.URL,
		AppName:     "ClientsMS::UpdateCountry()",
		Dest:        &countries,
		QueryString: q,
		AutoClose:   true,
	}

	//Open storage
	lgr.LOG.Info("_ACTION_: ", "opening connection")
	s, err = fs.NewStorage(qp)
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

// UpdateCountriesDeletionFlags - creates new country
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

	//Create query parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	qp := &fs.QueryParams{
		Ctx:         ctx,
		Dsn:         cfg.CFG.Database.URL,
		AppName:     "ClientsMS::UpdateCountriesDeletionFlags()",
		Dest:        &countries,
		QueryString: q,
		AutoClose:   true,
	}

	//Open storage
	lgr.LOG.Info("_ACTION_: ", "opening connection")
	s, err = fs.NewStorage(qp)
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
