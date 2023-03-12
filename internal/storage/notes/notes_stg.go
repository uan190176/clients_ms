package notes

import (
	"clients_ms/internal/api"
	cfg "clients_ms/pkg/config"
	lgr "clients_ms/pkg/logger"
	"context"
	fs "github.com/uan190176/fstorage"
	st "github.com/uan190176/statuses"
)

// GetNotes - returns client notes
func GetNotes(ctx context.Context, req *api.RequestNote) ([]*api.Note, st.ResponseStatus) {

	lgr.LOG.Info("-->> ", "notes.GetNotes()")

	// Vars
	var (
		s     *fs.Storage
		q     string
		notes []*api.Note
		err   error
	)

	//Check required fields
	lgr.LOG.Info("_ACTION_: ", "checking required fields")
	stat := CheckRequiredFieldsNotesSelect(req)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return nil, stat
	}

	//Generate query
	lgr.LOG.Info("_ACTION_: ", "generate query")
	q = GenQueryNotesSelect(req)

	//Create storage parameters
	lgr.LOG.Info("_ACTION_: ", "creating storage parameters")
	sp := &fs.StorageParams{
		Ctx:        ctx,
		Dsn:        cfg.CFG.Database.URL,
		AppName:    "ClientsMS::GetNotes()",
		AutoCommit: true,
		AutoClose:  true,
	}

	//Create query parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	qp := &fs.QueryParams{
		Dest:        &notes,
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
	if len(notes) > 0 {
		lgr.LOG.Info("_RESULT_: ", notes)
		lgr.LOG.Info("<<-- ", "notes.GetNotes()")
		return notes, st.GetStatus(100)
	}

	lgr.LOG.Warn("_WARN_: ", st.GetStatus(301))
	return nil, st.GetStatus(301)
}

// CreateNote - creates new note
func CreateNote(ctx context.Context, req *api.RequestNote) ([]*api.Note, st.ResponseStatus) {

	lgr.LOG.Info("-->> ", "notes.CreateNote()")

	// Vars
	var (
		s     *fs.Storage
		q     string
		notes []*api.Note
		err   error
		stat  st.ResponseStatus
	)

	//Check required fields
	lgr.LOG.Info("_ACTION_: ", "checking required fields")
	stat = CheckRequiredFieldsNoteInsert(req)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return nil, stat
	}

	//Create query
	lgr.LOG.Info("_ACTION_: ", "generating query")
	q, stat = CreateQueryNoteInsert(req)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return nil, stat
	}

	//Create storage parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	sp := &fs.StorageParams{
		Ctx:        ctx,
		Dsn:        cfg.CFG.Database.URL,
		AppName:    "ClientsMS::CreateNote()",
		AutoCommit: true,
		AutoClose:  true,
	}

	//Create query parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	qp := &fs.QueryParams{
		Dest:        &notes,
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
	if len(notes) > 0 {
		lgr.LOG.Info("_RESULT_: ", notes)
		lgr.LOG.Info("<<-- ", "notes.CreateNote()")
		return notes, st.GetStatus(100)
	}

	lgr.LOG.Warn("_WARN_: ", st.GetStatus(301))
	return nil, st.GetStatus(301)
}

// UpdateNote - updates note
func UpdateNote(ctx context.Context, req *api.RequestNote) ([]*api.Note, st.ResponseStatus) {

	lgr.LOG.Info("-->> ", "notes.UpdateNote()")

	// Vars
	var (
		s     *fs.Storage
		q     string
		notes []*api.Note
		err   error
		stat  st.ResponseStatus
	)

	//Check required fields
	lgr.LOG.Info("_ACTION_: ", "checking required fields")
	stat = CheckRequiredFieldsNoteUpdate(req)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return nil, stat
	}

	//Create query
	lgr.LOG.Info("_ACTION_: ", "generating query")
	q, stat = CreateQueryNoteUpdate(req)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return nil, stat
	}

	//Create storage parameters
	lgr.LOG.Info("_ACTION_: ", "creating storage parameters")
	sp := &fs.StorageParams{
		Ctx:        ctx,
		Dsn:        cfg.CFG.Database.URL,
		AppName:    "ClientsMS::UpdateNote()",
		AutoCommit: true,
		AutoClose:  true,
	}

	//Create query parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	qp := &fs.QueryParams{
		Dest:        &notes,
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
	if len(notes) > 0 {
		lgr.LOG.Info("_RESULT_: ", notes)
		lgr.LOG.Info("<<-- ", "notes.UpdateNote()")
		return notes, st.GetStatus(100)
	}

	lgr.LOG.Warn("_WARN_: ", st.GetStatus(301))
	return nil, st.GetStatus(301)
}

// UpdateNotesDeletionFlags - updates deletion flags
func UpdateNotesDeletionFlags(ctx context.Context, req *api.RequestNotesDeletionFlags) ([]*api.Note, st.ResponseStatus) {

	lgr.LOG.Info("-->> ", "notes.UpdateNotesDeletionFlags()")

	// Vars
	var (
		s     *fs.Storage
		q     string
		notes []*api.Note
		err   error
		stat  st.ResponseStatus
	)

	//Check required fields
	lgr.LOG.Info("_ACTION_: ", "checking required fields")
	stat = CheckRequiredFieldsNotesDeletionFlagUpdate(req)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return nil, stat
	}

	//Create query
	lgr.LOG.Info("_ACTION_: ", "generating query")
	q = CreateQueryNotesDeletionFlagsUpdate(req)

	//Create storage parameters
	lgr.LOG.Info("_ACTION_: ", "creating storage parameters")
	sp := &fs.StorageParams{
		Ctx:        ctx,
		Dsn:        cfg.CFG.Database.URL,
		AppName:    "ClientsMS::UpdateNotesDeletionFlags()",
		AutoCommit: true,
		AutoClose:  true,
	}

	//Create query parameters
	lgr.LOG.Info("_ACTION_: ", "creating query parameters")
	qp := &fs.QueryParams{
		Dest:        &notes,
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
	if len(notes) > 0 {
		lgr.LOG.Info("_RESULT_: ", notes)
		lgr.LOG.Info("<<-- ", "notes.UpdateNotesDeletionFlags()")
		return notes, st.GetStatus(100)
	}

	lgr.LOG.Warn("_WARN_: ", st.GetStatus(301))
	return nil, st.GetStatus(301)
}
