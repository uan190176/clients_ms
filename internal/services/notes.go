package services

import (
	"clients_ms/internal/api"
	hlp "clients_ms/internal/helpers"
	stg "clients_ms/internal/storage/notes"
	cfg "clients_ms/pkg/config"
	lgr "clients_ms/pkg/logger"
	"context"
	st "github.com/uan190176/statuses"
)

const notesTableName = "public.notes_reg_info"

// GetNotes - returns client notes
func (s *GrpcClientsServer) GetNotes(ctx context.Context, req *api.RequestNote) (*api.ResponseNotes, error) {

	lgr.LOG.Info("***")
	lgr.LOG.Info("-->> ", "services.GetNotes()")

	//Check auth token
	lgr.LOG.Info("_ACTION_: ", "checking auth token")
	if !hlp.AuthTokenIsValid(req.AuthToken) {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(101))
		return &api.ResponseNotes{
			Notes:  nil,
			Status: hlp.GetStatusForResponse(st.GetStatus(101)),
		}, nil
	}

	//Check users access rights
	lgr.LOG.Info("_ACTION_: ", "checking users access rights")
	allow, stat := hlp.GetUserAccessRightForTable(ctx, &api.RequestUsersAccessRightsForTable{
		AuthToken: cfg.CFG.MicroServices["users_ms"].Token,
		UserId:    req.AuthorId,
		TableName: notesTableName,
		Action:    api.UsersAccessRights_Actions_CAN_SELECT,
	})

	//error
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return &api.ResponseNotes{
			Notes:  nil,
			Status: hlp.GetStatusForResponse(stat),
		}, nil
	}

	//not allow
	if !allow {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(201))
		return &api.ResponseNotes{
			Notes:  nil,
			Status: hlp.GetStatusForResponse(st.GetStatus(201)),
		}, nil
	}

	// Work with storage
	lgr.LOG.Info("_ACTION_: ", "working with storage")
	notes, status := stg.GetNotes(ctx, req)

	lgr.LOG.Info("_RESULT_: ", status)
	lgr.LOG.Info("-->> ", "services.GetNotes()")
	lgr.LOG.Info("***")

	return &api.ResponseNotes{
		Notes:  notes,
		Status: hlp.GetStatusForResponse(status),
	}, nil
}

// CreateNote - creates new country
func (s *GrpcClientsServer) CreateNote(ctx context.Context, req *api.RequestNote) (*api.ResponseNotes, error) {

	lgr.LOG.Info("***")
	lgr.LOG.Info("-->> ", "services.CreateNote()")

	// Check auth token
	lgr.LOG.Info("_ACTION_: ", "checking auth token")
	if !hlp.AuthTokenIsValid(req.AuthToken) {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(101))
		return &api.ResponseNotes{
			Notes:  nil,
			Status: hlp.GetStatusForResponse(st.GetStatus(101)),
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
		return &api.ResponseNotes{
			Notes:  nil,
			Status: hlp.GetStatusForResponse(stat),
		}, nil
	}

	//not allow
	if !allow {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(202))
		return &api.ResponseNotes{
			Notes:  nil,
			Status: hlp.GetStatusForResponse(st.GetStatus(202)),
		}, nil
	}

	// Work with storage
	lgr.LOG.Info("_ACTION_: ", "working with storage")
	notes, status := stg.CreateNote(ctx, req)

	//Result
	lgr.LOG.Info("_RESULT_: ", status)
	lgr.LOG.Info("<<-- ", "services.CreateNote()")
	lgr.LOG.Info("***")
	return &api.ResponseNotes{
		Notes:  notes,
		Status: hlp.GetStatusForResponse(status),
	}, nil
}

// UpdateNote - updates country
func (s *GrpcClientsServer) UpdateNote(ctx context.Context, req *api.RequestNote) (*api.ResponseNotes, error) {

	lgr.LOG.Info("***")
	lgr.LOG.Info("-->> ", "services.UpdateNote()")

	// Check auth token
	lgr.LOG.Info("_ACTION_: ", "checking auth token")
	if !hlp.AuthTokenIsValid(req.AuthToken) {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(101))
		return &api.ResponseNotes{
			Notes:  nil,
			Status: hlp.GetStatusForResponse(st.GetStatus(101)),
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
		return &api.ResponseNotes{
			Notes:  nil,
			Status: hlp.GetStatusForResponse(stat),
		}, nil
	}

	//not allow
	if !allow {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(203))
		return &api.ResponseNotes{
			Notes:  nil,
			Status: hlp.GetStatusForResponse(st.GetStatus(203)),
		}, nil
	}

	// Work with storage
	lgr.LOG.Info("_ACTION_: ", "working with storage")
	notes, status := stg.UpdateNote(ctx, req)

	//Result
	lgr.LOG.Info("_RESULT_: ", status)
	lgr.LOG.Info("<<-- ", "services.UpdateNote()")
	lgr.LOG.Info("***")
	return &api.ResponseNotes{
		Notes:  notes,
		Status: hlp.GetStatusForResponse(status),
	}, nil
}

// UpdateNotesDeletionFlags - updates countries deletion flags
func (s *GrpcClientsServer) UpdateNotesDeletionFlags(ctx context.Context, req *api.RequestNotesDeletionFlags) (*api.ResponseNotes, error) {

	lgr.LOG.Info("***")
	lgr.LOG.Info("-->> ", "services.UpdateNotesDeletionFlags()")

	// Check auth token
	lgr.LOG.Info("_ACTION_: ", "checking auth token")
	if !hlp.AuthTokenIsValid(req.AuthToken) {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(101))
		return &api.ResponseNotes{
			Notes:  nil,
			Status: hlp.GetStatusForResponse(st.GetStatus(101)),
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
		return &api.ResponseNotes{
			Notes:  nil,
			Status: hlp.GetStatusForResponse(stat),
		}, nil
	}

	//not allow
	if !allow {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(204))
		return &api.ResponseNotes{
			Notes:  nil,
			Status: hlp.GetStatusForResponse(st.GetStatus(204)),
		}, nil
	}

	// Work with storage
	lgr.LOG.Info("_ACTION_: ", "working with storage")
	notes, status := stg.UpdateNotesDeletionFlags(ctx, req)

	//Result
	lgr.LOG.Info("_RESULT_: ", status)
	lgr.LOG.Info("<<-- ", "services.UpdateNotesDeletionFlags()")
	lgr.LOG.Info("***")
	return &api.ResponseNotes{
		Notes:  notes,
		Status: hlp.GetStatusForResponse(status),
	}, nil
}
