package notes

import (
	"clients_ms/internal/api"
	hlp "clients_ms/internal/helpers"
	lgr "clients_ms/pkg/logger"
	"fmt"
	st "github.com/uan190176/statuses"
	"strings"
)

//**************************************************************
// Check required fields
//**************************************************************

// CheckRequiredFieldsNotesSelect - checks required fields to notes INSERT
func CheckRequiredFieldsNotesSelect(req *api.RequestNote) st.ResponseStatus {
	lgr.LOG.Info("-->> ", "notes.CheckRequiredFieldsNotesSelect()")

	//ClientId
	lgr.LOG.Info("_ACTION_: ", "checking ClientId")
	if req.ClientId == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " ClientId", " must contain a value"))
		return st.GetStatus(401, " ClientId", " must contain a value")
	}

	lgr.LOG.Info("<<-- ", "notes.CheckRequiredFieldsNotesSelect()")
	return st.GetStatus(100)
}

// CheckRequiredFieldsNoteInsert - checks required fields to notes INSERT
func CheckRequiredFieldsNoteInsert(req *api.RequestNote) st.ResponseStatus {
	lgr.LOG.Info("-->> ", "notes.CheckRequiredFieldsNoteInsert()")

	//AuthorId
	lgr.LOG.Info("_ACTION_: ", "checking AuthorId")
	if req.AuthorId == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " AuthorId", " must contain a value"))
		return st.GetStatus(401, " AuthorId", " must contain a value")
	}

	//ClientId
	lgr.LOG.Info("_ACTION_: ", "checking ClientId")
	if req.ClientId == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " ClientId", " must contain a value"))
		return st.GetStatus(401, " ClientId", " must contain a value")
	}

	//NoteText
	lgr.LOG.Info("_ACTION_: ", "checking delivery name")
	if req.NoteText == "" {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " NoteText", " must contain a value"))
		return st.GetStatus(401, " NoteText", " must contain a value")
	}

	lgr.LOG.Info("<<-- ", "notes.CheckRequiredFieldsNoteInsert()")
	return st.GetStatus(100)
}

// CheckRequiredFieldsNoteUpdate - checks required fields to notes UPDATE
func CheckRequiredFieldsNoteUpdate(req *api.RequestNote) st.ResponseStatus {
	lgr.LOG.Info("-->> ", "notes.CheckRequiredFieldsNoteUpdate()")

	//AuthorId
	lgr.LOG.Info("_ACTION_: ", "checking AuthorId")
	if req.AuthorId == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " AuthorId", " must contain a value"))
		return st.GetStatus(401, " AuthorId", " must contain a value")
	}

	//ClientId
	lgr.LOG.Info("_ACTION_: ", "checking ClientId")
	if req.ClientId == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " ClientId", " must contain a value"))
		return st.GetStatus(401, " ClientId", " must contain a value")
	}

	//NoteText
	lgr.LOG.Info("_ACTION_: ", "checking delivery name")
	if req.NoteText == "" {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " NoteText", " must contain a value"))
		return st.GetStatus(401, " NoteText", " must contain a value")
	}

	lgr.LOG.Info("<<-- ", "notes.CheckRequiredFieldsNoteUpdate()")
	return st.GetStatus(100)
}

// CheckRequiredFieldsNotesDeletionFlagUpdate - checked required fields to UPDATE deletion flags
func CheckRequiredFieldsNotesDeletionFlagUpdate(req *api.RequestNotesDeletionFlags) st.ResponseStatus {
	lgr.LOG.Info("-->> ", "notes.CheckRequiredFieldsNotesDeletionFlagUpdate()")

	//AuthorId
	lgr.LOG.Info("_ACTION_: ", "checking AuthorId")
	if req.AuthorId == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " AuthorId", " must contain a value"))
		return st.GetStatus(401, " AuthorId", " must contain a value")
	}

	//IsDeleted
	lgr.LOG.Info("_ACTION_: ", "checking IsDeleted")
	if req.IsDeleted == api.ClientsMS_Bool_IS_ANY {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " IsDeleted", "must contain a value"))
		return st.GetStatus(401, " IsDeleted", " must contain a value")
	}

	//Ids
	lgr.LOG.Info("_ACTION_: ", "checking Ids")
	if len(req.Ids) == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " Ids", " must contain a value"))
		return st.GetStatus(401, " Ids", " must contain a value")
	}

	lgr.LOG.Info("<<-- ", "notes.CheckRequiredFieldsNotesDeletionFlagUpdate()")
	return st.GetStatus(100)
}

//**************************************************************
// Create queries
//**************************************************************

// createQueryNotesSelect - returns query to SELECT
func createQueryNotesSelect() string {
	lgr.LOG.Info("-->> ", "notes.createQueryNotesSelect()")
	q := fmt.Sprintf("SELECT %s FROM clients_notes_reg_info", createViewNotes())
	lgr.LOG.Info("_QUERY_: ", q)
	lgr.LOG.Info("<<-- ", "notes.createQueryNotesSelect()")
	return q
}

// CreateQueryNoteInsert - returns query to INSERT
func CreateQueryNoteInsert(req *api.RequestNote) (string, st.ResponseStatus) {
	lgr.LOG.Info("-->> ", "notes.CreateQueryNoteInsert()")

	q := fmt.Sprintf("INSERT INTO clients_notes_reg_info (client_id, note_text, created_by) VALUES (%d, '%s', %d) RETURNING ", req.ClientId, req.NoteText, req.AuthorId)

	q += createViewNotes()

	lgr.LOG.Info("_QUERY_: ", q)
	lgr.LOG.Info("<<-- ", "notes.CreateQueryNoteInsert()")
	return q, st.GetStatus(100)
}

// CreateQueryNoteUpdate - returns query to UPDATE
func CreateQueryNoteUpdate(req *api.RequestNote) (string, st.ResponseStatus) {
	lgr.LOG.Info("-->> ", "notes.CreateQueryNoteUpdate()")

	sFields := ""
	sValues := ""
	cnt := 0

	//NoteText
	if req.NoteText != "" {
		sFields += "note_text, "
		sValues += fmt.Sprintf("'%s', ", req.NoteText)
		cnt++
	}

	//AuthorId
	if req.AuthorId != 0 {
		sFields += "updated_by, "
		sValues += fmt.Sprintf("%d, ", req.AuthorId)
		cnt++
	}

	// Is deleted
	if req.IsDeleted == api.ClientsMS_Bool_IS_TRUE || req.IsDeleted == api.ClientsMS_Bool_IS_FALSE {
		sFields += "isdeleted, "
		if req.IsDeleted == api.ClientsMS_Bool_IS_TRUE {
			sValues += fmt.Sprintf("%t, ", true)
		} else {
			sValues += fmt.Sprintf("%t, ", false)
		}
		cnt++
	}

	sFields = strings.TrimSuffix(sFields, ", ")
	sValues = strings.TrimSuffix(sValues, ", ")

	q := ""
	if cnt > 1 {
		q = "UPDATE clients_notes_reg_info SET (" + sFields + ") = (" + sValues + ")"
	} else {
		q = "UPDATE clients_notes_reg_info SET " + sFields + " = " + sValues
	}

	q += fmt.Sprintf(" WHERE id = %d RETURNING ", req.Id)
	q += createViewNotes()

	lgr.LOG.Info("_QUERY_: ", q)

	lgr.LOG.Info("<<-- ", "notes.CreateQueryNoteUpdate()")
	return q, st.GetStatus(100)
}

// CreateQueryNotesDeletionFlagsUpdate - returns query to UPDATE notes deletion flags
func CreateQueryNotesDeletionFlagsUpdate(req *api.RequestNotesDeletionFlags) string {
	lgr.LOG.Info("-->> ", "notes.CreateQueryNotesDeletionFlagsUpdate()")
	q := fmt.Sprintf("UPDATE clients_notes_reg_info SET isdeleted = %t, updated_by = %d WHERE id IN %s RETURNING ",
		hlp.GetBool(req.IsDeleted), req.AuthorId, hlp.Uint64ArrayToString(req.Ids))
	q += createViewNotes()
	lgr.LOG.Info("_QUERY_: ", q)
	lgr.LOG.Info("<<-- ", "notes.CreateQueryNotesDeletionFlagsUpdate()")
	return q
}

// **************************************************************
// Create view
// **************************************************************
func createViewNotes() string {
	return `id AS note_id,
			client_id AS note_client_id,
    		note_text AS note_text,
    		to_char(created_at, 'YYYY-MM-DD HH24:MI:SS'::text) AS note_created_at,
    		created_by AS note_created_by,
            ( SELECT u.fullname FROM users_ref u WHERE u.id = created_by ) AS note_created_by_name,
    		to_char(updated_at, 'YYYY-MM-DD HH24:MI:SS'::text) AS note_updated_at,
    		updated_by AS note_updated_by,
    		( SELECT u.fullname FROM users_ref u WHERE u.id = updated_by ) AS note_updated_by_name,
    		isdeleted AS note_isdeleted`
}

//**************************************************************
// Add WHERE to queries
//**************************************************************

// addWhereToQueryNotesSelect - added WHERE to query SELECT
func addWhereToQueryNotesSelect(q string, req *api.RequestNote) string {
	lgr.LOG.Info("-->> ", "notes.addWhereToQueryNotesSelect()")
	cnt := 0

	// ClientID
	if req.ClientId != 0 {
		q += fmt.Sprintf(" WHERE client_id = %d", req.ClientId)
		cnt++
	}

	// Isdeleted
	switch req.IsDeleted {
	case api.ClientsMS_Bool_IS_TRUE:
		if cnt > 0 {
			q += " AND isdeleted = true"
		} else {
			q += " WHERE isdeleted = true"
		}
		cnt++
	case api.ClientsMS_Bool_IS_FALSE:
		if cnt > 0 {
			q += " AND isdeleted = false"
		} else {
			q += " WHERE isdeleted = false"
		}
		cnt++
	}

	lgr.LOG.Info("_QUERY_: ", q)
	lgr.LOG.Info("<<-- ", "notes.addWhereToQueryNotesSelect()")
	return q
}

//**************************************************************
// Add ORDER to queries
//**************************************************************

// addOrderToQueryNotesSelect - added ORDER to query SELECT
func addOrderToQueryNotesSelect(q string, req *api.RequestNote) string {

	lgr.LOG.Info("-->> ", "notes.addOrderToQueryNotesSelect()")

	strOrder := ""
	switch req.Order {
	case api.ClientsMS_NotesOrder_BY_NOTE_ID:
		strOrder = " ORDER BY note_id"
	case api.ClientsMS_NotesOrder_BY_NOTE_CREATED_AT:
		strOrder = " ORDER BY note_created_at"
	case api.ClientsMS_NotesOrder_BY_NOTE_UPDATED_AT:
		strOrder = " ORDER BY note_updated_at"
	default:
		strOrder = " ORDER BY note_created_at"
	}

	strDirection := ""
	switch req.OrderType {
	case api.ClientsMS_OrderType_ASC:
		strDirection = " ASC"
	case api.ClientsMS_OrderType_DESC:
		strDirection = " DESC"
	default:
		strDirection = " DESC"
	}

	q += strOrder + strDirection
	lgr.LOG.Info("_QUERY_: ", q)

	lgr.LOG.Info("<<-- ", "notes.addOrderToQueryNotesSelect()")
	return q
}

//**************************************************************
// Generating a queries
//**************************************************************

// GenQueryNotesSelect - generates a query for clients notes select
func GenQueryNotesSelect(req *api.RequestNote) string {

	lgr.LOG.Info("-->> ", "notes.GenQueryNotesSelect()")
	var q string

	lgr.LOG.Info("_ACTION_: ", "generating query")
	q = createQueryNotesSelect()

	lgr.LOG.Info("_ACTION_: ", "adding WHERE to query")
	q = addWhereToQueryNotesSelect(q, req)

	lgr.LOG.Info("_ACTION_: ", "adding ORDER to query")
	q = addOrderToQueryNotesSelect(q, req)

	lgr.LOG.Info("_QUERY_: ", q)
	lgr.LOG.Info("<<-- ", "notes.GenQueryNotesSelect()")
	return q
}
