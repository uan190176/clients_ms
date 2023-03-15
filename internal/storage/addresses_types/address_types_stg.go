package addresses_types

import (
	"clients_ms/internal/api"
	cfg "clients_ms/pkg/config"
	lgr "clients_ms/pkg/logger"
	"context"
	fs "github.com/uan190176/fstorage"
	st "github.com/uan190176/statuses"
)

// GetAddressTypes - returns countries
func GetAddressTypes(ctx context.Context, req *api.RequestAddressType) ([]*api.AddressType, st.ResponseStatus) {

	lgr.LOG.Info("-->> ", "address_types.GetAddressTypes()")

	// Vars
	var (
		s         *fs.Storage
		q         string
		addrTypes []*api.AddressType
		err       error
	)

	//Generate query
	lgr.LOG.Info("_ACTION_: ", "generate query")
	q = GenQueryAddressesTypesSelect(req)

	//Create storage parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	sp := &fs.StorageParams{
		Ctx:        ctx,
		Dsn:        cfg.CFG.Database.URL,
		AppName:    "ClientsMS::GetAddressTypes()",
		AutoCommit: true,
		AutoClose:  true,
	}

	//Create query parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	qp := &fs.QueryParams{
		Dest:        &addrTypes,
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
	if len(addrTypes) > 0 {
		lgr.LOG.Info("_RESULT_: ", addrTypes)
		lgr.LOG.Info("<<-- ", "address_types.GetAddressTypes()")
		return addrTypes, st.GetStatus(100)
	}

	lgr.LOG.Warn("_WARN_: ", st.GetStatus(301))
	return nil, st.GetStatus(301)
}

// CreateAddressType - creates new address type
func CreateAddressType(ctx context.Context, req *api.RequestAddressType) ([]*api.AddressType, st.ResponseStatus) {

	lgr.LOG.Info("-->> ", "address_types.CreateAddressType()")

	// Vars
	var (
		s         *fs.Storage
		q         string
		addrTypes []*api.AddressType
		err       error
		stat      st.ResponseStatus
	)

	//Check required fields
	lgr.LOG.Info("_ACTION_: ", "checking required fields")
	stat = CheckAndNormaliseRequiredFieldsAddressTypeInsert(req)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return nil, stat
	}

	//Create query
	lgr.LOG.Info("_ACTION_: ", "generating query")
	q, stat = CreateQueryAddressTypeInsert()
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return nil, stat
	}

	//Create storage parameters
	lgr.LOG.Info("_ACTION_: ", "creating storage parameters")
	sp := &fs.StorageParams{
		Ctx:        ctx,
		Dsn:        cfg.CFG.Database.URL,
		AppName:    "ClientsMS::CreateAddressType()",
		AutoCommit: true,
		AutoClose:  true,
	}

	//Create query parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	qp := &fs.QueryParams{
		Dest:        &addrTypes,
		QueryString: q,
		Args:        []interface{}{req.Name, req.Code, req.Comment, req.AuthorId},
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
	if len(addrTypes) > 0 {
		lgr.LOG.Info("_RESULT_: ", addrTypes)
		lgr.LOG.Info("<<-- ", "address_types.CreateAddressType()")
		return addrTypes, st.GetStatus(100)
	}

	lgr.LOG.Warn("_WARN_: ", st.GetStatus(301))
	return nil, st.GetStatus(301)
}

// UpdateAddressType - update address type
func UpdateAddressType(ctx context.Context, req *api.RequestAddressType) ([]*api.AddressType, st.ResponseStatus) {

	lgr.LOG.Info("-->> ", "address_types.UpdateAddressType()")

	// Vars
	var (
		s         *fs.Storage
		q         string
		addrTypes []*api.AddressType
		err       error
		stat      st.ResponseStatus
	)

	//Check required fields
	lgr.LOG.Info("_ACTION_: ", "checking required fields")
	stat = CheckAndNormaliseRequiredFieldsAddressTypeUpdate(req)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return nil, stat
	}

	//Create query
	lgr.LOG.Info("_ACTION_: ", "generating query")
	q, stat = CreateQueryAddressTypeUpdate(req)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return nil, stat
	}

	//Create storage parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	sp := &fs.StorageParams{
		Ctx:        ctx,
		Dsn:        cfg.CFG.Database.URL,
		AppName:    "ClientsMS::UpdateAddressType()",
		AutoCommit: true,
		AutoClose:  true,
	}

	//Create query parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	qp := &fs.QueryParams{
		Dest:        &addrTypes,
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
	if len(addrTypes) > 0 {
		lgr.LOG.Info("_RESULT_: ", addrTypes)
		lgr.LOG.Info("<<-- ", "address_types.UpdateAddressType()")
		return addrTypes, st.GetStatus(100)
	}

	lgr.LOG.Warn("_WARN_: ", st.GetStatus(301))
	return nil, st.GetStatus(301)
}

// UpdateAddressTypeDeletionFlags - mark/unmark addresses types as deleted in database
func UpdateAddressTypeDeletionFlags(ctx context.Context, req *api.RequestAddressesTypesDeletionFlags) ([]*api.AddressType, st.ResponseStatus) {
	lgr.LOG.Info("-->> ", "address_types.UpdateAddressTypeDeletionFlags()")

	// Vars
	var (
		s         *fs.Storage
		q         string
		addrTypes []*api.AddressType
		err       error
		stat      st.ResponseStatus
	)

	//Check required fields
	lgr.LOG.Info("_ACTION_: ", "checking required fields")
	stat = CheckRequiredFieldsAddressesTypesDeletionFlagsUpdate(req)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return nil, stat
	}

	//Create query
	lgr.LOG.Info("_ACTION_: ", "generating query")
	q = CreateQueryAddressesTypesDeletionFlagsUpdate(req)

	//Create storage parameters
	lgr.LOG.Info("_ACTION_: ", "creating storage parameters")
	sp := &fs.StorageParams{
		Ctx:        ctx,
		Dsn:        cfg.CFG.Database.URL,
		AppName:    "ClientsMS::UpdateAddressTypeDeletionFlags()",
		AutoCommit: true,
		AutoClose:  true,
	}

	//Create query parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	qp := &fs.QueryParams{
		Dest:        &addrTypes,
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
	if len(addrTypes) > 0 {
		lgr.LOG.Info("_RESULT_: ", addrTypes)
		lgr.LOG.Info("<<-- ", "address_types.UpdateAddressType()")
		return addrTypes, st.GetStatus(100)
	}

	lgr.LOG.Warn("_WARN_: ", st.GetStatus(301))
	return nil, st.GetStatus(301)
}
