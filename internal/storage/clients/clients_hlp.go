package clients

import (
	"clients_ms/internal/api"
	hlp "clients_ms/internal/helpers"
	cfg "clients_ms/pkg/config"
	lgr "clients_ms/pkg/logger"
	"fmt"
	st "github.com/uan190176/statuses"
	"strings"
)

const (
	russianSymbols = "[А-Яа-яЁё]+"
	emailTemplate  = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	phoneTemplate  = "^7\\d{10}$"
	//addressTemplate    = "^[а-яА-ЯёЁ0-9\\s\\.,#-]+$ \n"
	//postalCodeTemplate = "^\\d{6}$\n"
)

//**************************************************************
// Check required fields
//**************************************************************

// CheckAndNormalizeRequiredFieldsClientInsert - checks required fields to client INSERT
func CheckAndNormalizeRequiredFieldsClientInsert(req *api.RequestClient) st.ResponseStatus {

	lgr.LOG.Info("-->> ", "clients.CheckAndNormalizeRequiredFieldsClientInsert()")

	var (
		stat st.ResponseStatus
	)

	//AuthorId
	lgr.LOG.Info("_ACTION_: ", "checking AuthorId")
	if req.AuthorId == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " AuthorId", " must contain a value"))
		return st.GetStatus(401, " AuthorId", " must contain a value")
	}

	//FirstName
	lgr.LOG.Info("_ACTION_: ", "checking FirstName")
	if req.FirstName == "" {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " FirstName", " must contain a value"))
		return st.GetStatus(401, " FirstName", " must contain a value")
	} else {
		req.FirstName, stat = st.NormalizeStringByTemplate(req.FirstName, russianSymbols, true, false)
		if stat != st.GetStatus(100) {
			lgr.LOG.Warn("_WARN_: ", stat)
			stat.Description += " FirstName"
			return stat
		}
	}

	//MiddleName
	lgr.LOG.Info("_ACTION_: ", "checking MiddleName")
	if req.MiddleName == "" {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " MiddleName", " must contain a value"))
		return st.GetStatus(401, " MiddleName", " must contain a value")
	} else {
		req.MiddleName, stat = st.NormalizeStringByTemplate(req.MiddleName, russianSymbols, true, false)
		if stat != st.GetStatus(100) {
			lgr.LOG.Warn("_WARN_: ", stat)
			stat.Description += " MiddleName"
			return stat
		}
	}

	//LastName
	lgr.LOG.Info("_ACTION_: ", "checking LastName")
	if req.LastName == "" {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " LastName", " must contain a value"))
		return st.GetStatus(401, " LastName", " must contain a value")
	} else {
		req.LastName, stat = st.NormalizeStringByTemplate(req.LastName, russianSymbols, true, false)
		if stat != st.GetStatus(100) {
			lgr.LOG.Warn("_WARN_: ", stat)
			stat.Description += " LastName"
			return stat
		}
	}

	//Phone
	lgr.LOG.Info("_ACTION_: ", "checking Phone")
	if req.Phone == "" {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " Phone", " must contain a value"))
		return st.GetStatus(401, " Phone", " must contain a value")
	} else {
		req.Phone, stat = st.NormalizeStringByTemplate(req.Phone, phoneTemplate)
		if stat != st.GetStatus(100) {
			lgr.LOG.Warn("_WARN_: ", stat)
			stat.Description += " Phone"
			return stat
		}
	}

	//Email
	lgr.LOG.Info("_ACTION_: ", "checking Email")
	if req.Email == "" {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " Email", " must contain a value"))
		return st.GetStatus(401, " Email", " must contain a value")
	} else {
		req.Email, stat = st.NormalizeStringByTemplate(req.Email, emailTemplate)
		if stat != st.GetStatus(100) {
			lgr.LOG.Warn("_WARN_: ", stat)
			stat.Description += " Email"
			return stat
		}
	}

	////Address
	//lgr.LOG.Info("_ACTION_: ", "checking Address")
	//if req.Address == "" {
	//	lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " Address", " must contain a value"))
	//	return st.GetStatus(401, " Address", " must contain a value")
	//} else {
	//	req.Address, stat = st.NormalizeStringByTemplate(req.Address, addressTemplate)
	//	if stat != st.GetStatus(100) {
	//		lgr.LOG.Warn("_WARN_: ", stat)
	//		return stat
	//	}
	//}
	//
	////PostalCode
	//lgr.LOG.Info("_ACTION_: ", "checking PostalCode")
	//if req.PostalCode == "" {
	//	lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " PostalCode", " must contain a value"))
	//	return st.GetStatus(401, " PostalCode", " must contain a value")
	//} else {
	//	req.PostalCode, stat = st.NormalizeStringByTemplate(req.PostalCode, postalCodeTemplate)
	//	if stat != st.GetStatus(100) {
	//		lgr.LOG.Warn("_WARN_: ", stat)
	//		return stat
	//	}
	//}

	lgr.LOG.Info("<<-- ", "client.CheckAndNormalizeRequiredFieldsClientInsert()")
	return st.GetStatus(100)
}

// CheckAndNormalizeRequiredFieldsClientUpdate - checks required fields to client UPDATE
func CheckAndNormalizeRequiredFieldsClientUpdate(req *api.RequestClient) st.ResponseStatus {

	lgr.LOG.Info("-->> ", "clients.CheckAndNormalizeRequiredFieldsClientUpdate()")
	var stat st.ResponseStatus

	//Id
	lgr.LOG.Info("_ACTION_: ", "checking Id")
	if req.Id == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " Id", " must contain a value"))
		return st.GetStatus(401, " Id", " must contain a value")
	}

	//AuthorId
	lgr.LOG.Info("_ACTION_: ", "checking AuthorId")
	if req.AuthorId == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " AuthorId", " must contain a value"))
		return st.GetStatus(401, " AuthorId", " must contain a value")
	}

	stat = st.GetStatus(100)
	lgr.LOG.Info("<<-- ", "client.CheckAndNormalizeRequiredFieldsClientUpdate()")
	return stat
}

// CheckRequiredFieldsClientsDeletionFlagUpdate - checked required fields to UPDATE deletion flags
func CheckRequiredFieldsClientsDeletionFlagUpdate(req *api.RequestClientsDeletionFlags) st.ResponseStatus {
	lgr.LOG.Info("-->> ", "clients.CheckRequiredFieldsClientsDeletionFlagUpdate()")

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

	lgr.LOG.Info("<<-- ", "clients.CheckRequiredFieldsClientsDeletionFlagUpdate()")
	return st.GetStatus(100)
}

// **************************************************************
// Create queries
// **************************************************************

// createQueryClientsSelect - create query for clients SELECT
func createQueryClientsSelect(req *api.RequestClient) string {
	lgr.LOG.Info("-->> ", "clients.createQueryClientsSelect()")
	var q string
	if req.Address == "" {
		q = fmt.Sprintf("SELECT %s FROM clients_ref c", createViewClientsWhitLastAddress())
	} else {
		q = fmt.Sprintf("SELECT %s FROM clients_ref c", createViewClientsWhitFilterByAddress())
	}

	lgr.LOG.Info("_QUERY_: ", hlp.DeleteTabsAndNewLinesSymbols(q))
	lgr.LOG.Info("<<-- ", "clients.createQueryClientsSelect()")
	return q
}

// CreateQueryClientsInsert - create query for clients INSERT
func CreateQueryClientsInsert(req *api.RequestClient) string {

	lgr.LOG.Info("-->> ", "clients.CreateQueryClientsInsert()")

	q := fmt.Sprintf("INSERT INTO clients_ref AS c (firstname, middlename, lastname, phone, email, comment, created_by) "+
		"VALUES ('%s', '%s', '%s', '%s', '%s', '%s', %d) RETURNING ", req.FirstName, req.MiddleName, req.LastName, req.Phone, req.Email, req.Comment, req.AuthorId)
	q += createViewClients()

	lgr.LOG.Info("_QUERY_: ", hlp.DeleteTabsAndNewLinesSymbols(q))
	lgr.LOG.Info("<<-- ", "clients.CreateQueryClientsInsert()")
	return q
}

// CreateQueryClientsUpdate - create query for clients UPDATE
func CreateQueryClientsUpdate(req *api.RequestClient) string {

	lgr.LOG.Info("-->> ", "clients.CreateQueryClientsUpdate()")
	sFields := ""
	sValues := ""
	cnt := 0

	//AuthorId
	if req.AuthorId > 0 {
		sFields += "updated_by, "
		sValues += fmt.Sprintf("%d, ", req.AuthorId)
		cnt++
	}

	//FirstName
	if req.FirstName != "" {
		sFields += "firstname, "
		sValues += fmt.Sprintf("'%s', ", req.FirstName)
		cnt++
	}

	//MiddleName
	if req.MiddleName != "" {
		sFields += "middlename, "
		sValues += fmt.Sprintf("'%s', ", req.MiddleName)
		cnt++
	}

	//LastName
	if req.LastName != "" {
		sFields += "lastname, "
		sValues += fmt.Sprintf("'%s', ", req.LastName)
		cnt++
	}

	//Phone
	if req.Phone != "" {
		sFields += "phone, "
		sValues += fmt.Sprintf("'%s', ", req.Phone)
		cnt++
	}

	//Email
	if req.Email != "" {
		sFields += "email, "
		sValues += fmt.Sprintf("'%s', ", req.Email)
		cnt++
	}

	//Comment
	if req.Comment != "" {
		sFields += "comment, "
		sValues += fmt.Sprintf("'%s', ", req.Comment)
		cnt++
	}

	//IsBad
	if req.IsBad != api.ClientsMS_Bool_IS_ANY {
		sFields += "isbad, "
		if req.IsBad == api.ClientsMS_Bool_IS_TRUE {
			sValues += fmt.Sprintf("%t, ", true)
		} else {
			sValues += fmt.Sprintf("%t, ", false)
		}
	}

	//IsDeleted
	if req.IsDeleted != api.ClientsMS_Bool_IS_ANY {
		sFields += "isdeleted, "
		if req.IsDeleted == api.ClientsMS_Bool_IS_TRUE {
			sValues += fmt.Sprintf("%t, ", true)
		} else {
			sValues += fmt.Sprintf("%t, ", false)
		}
	}

	sFields = strings.TrimSuffix(sFields, ", ")
	sValues = strings.TrimSuffix(sValues, ", ")

	var q string
	if cnt > 1 {
		q = "UPDATE clients_ref AS c SET (" + sFields + ") = (" + sValues + ")"
	} else {
		q = "UPDATE clients_ref AS c SET " + sFields + " = " + sValues
	}

	q += fmt.Sprintf(" WHERE id = %d RETURNING ", req.Id)
	q += createViewClients()

	lgr.LOG.Info("_QUERY_: ", hlp.DeleteTabsAndNewLinesSymbols(q))
	lgr.LOG.Info("<<-- ", "clients.CreateQueryClientsUpdate()")
	return q
}

// CreateQueryClientsDeletionFlagsUpdate - returns query to UPDATE notes deletion flags
func CreateQueryClientsDeletionFlagsUpdate(req *api.RequestClientsDeletionFlags) string {
	lgr.LOG.Info("-->> ", "clients.CreateQueryClientsDeletionFlagsUpdate()")
	q := fmt.Sprintf("UPDATE clients_ref AS c SET isdeleted = %t, updated_by = %d WHERE id IN %s RETURNING ",
		hlp.GetBool(req.IsDeleted), req.AuthorId, hlp.Uint64ArrayToString(req.Ids))
	q += createViewClients()
	lgr.LOG.Info("_QUERY_: ", hlp.DeleteTabsAndNewLinesSymbols(q))
	lgr.LOG.Info("<<-- ", "clients.CreateQueryClientsDeletionFlagsUpdate()")
	return q
}

// **************************************************************
// Create views
// **************************************************************

// createViewClients - create view for clients whit last address
func createViewClients() string {
	return `c.id AS client_id,
			c.firstname AS client_firstname,
			c.middlename AS client_middlename,
			c.lastname AS client_lastname,
			c.fullname AS client_fullname,
			c.regnumber AS client_regnumber,
			c.phone AS client_phone,
			c.email AS client_email,
			concat(c.comment, '') AS client_comment,
			to_char(c.created_at, 'YYYY-MM-DD HH24:MI:SS'::text) AS client_created_at,
			c.created_by AS client_created_by,
			( SELECT u.fullname FROM users_ref u WHERE u.id = created_by ) AS client_created_by_name,
			to_char(c.updated_at, 'YYYY-MM-DD HH24:MI:SS'::text) AS client_updated_at,
			c.updated_by AS client_updated_by,
			( SELECT u.fullname FROM users_ref u WHERE u.id = updated_by ) AS client_updated_by_name,
			c.isdeleted AS client_isdeleted,
			c.isbad AS client_isbad`
}

// createViewClientsWhitLastAddress - create view for clients whit last address
func createViewClientsWhitLastAddress() string {
	return `c.id AS client_id,
			c.firstname AS client_firstname,
			c.middlename AS client_middlename,
			c.lastname AS client_lastname,
			c.fullname AS client_fullname,
			c.regnumber AS client_regnumber,
			c.phone AS client_phone,
			c.email AS client_email,
			concat(c.comment, '') AS client_comment,
			to_char(c.created_at, 'YYYY-MM-DD HH24:MI:SS'::text) AS client_created_at,
			c.created_by AS client_created_by,
			( SELECT u.fullname FROM users_ref u WHERE u.id = created_by ) AS client_created_by_name,
			to_char(c.updated_at, 'YYYY-MM-DD HH24:MI:SS'::text) AS client_updated_at,
			c.updated_by AS client_updated_by,
			( SELECT u.fullname FROM users_ref u WHERE u.id = updated_by ) AS client_updated_by_name,
			c.isdeleted AS client_isdeleted,
			c.isbad AS client_isbad,
			c.last_address_id AS client_last_address_id,
			( SELECT a.address FROM addresses_ref a WHERE a.id = last_address_id ) AS client_last_address`
}

// createViewClientsWhitFilterByAddress - create view for clients whit filter by address
func createViewClientsWhitFilterByAddress() string {
	return `c.id AS client_id,
			c.firstname AS client_firstname,
			c.middlename AS client_middlename,
			c.lastname AS client_lastname,
			c.fullname AS client_fullname,
			c.regnumber AS client_regnumber,
			c.phone AS client_phone,
			c.email AS client_email,
			concat(c.comment, '') AS client_comment,
			to_char(c.created_at, 'YYYY-MM-DD HH24:MI:SS'::text) AS client_created_at,
			c.created_by AS client_created_by,
			( SELECT u.fullname FROM users_ref u WHERE u.id = created_by ) AS client_created_by_name,
			to_char(c.updated_at, 'YYYY-MM-DD HH24:MI:SS'::text) AS client_updated_at,
			c.updated_by AS client_updated_by,
			( SELECT u.fullname FROM users_ref u WHERE u.id = updated_by ) AS client_updated_by_name,
			c.isdeleted AS client_isdeleted,
			c.isbad AS client_isbad,
			c.last_address_id AS client_last_address_id,
			( SELECT a.address FROM addresses_ref a WHERE a.id = last_address_id ) AS client_last_address,
			ad.address AS client_found_address`
}

//**************************************************************
// Add WHERE to queries
//**************************************************************

func addWhereToQueryClientsSelect(q string, req *api.RequestClient) string {
	lgr.LOG.Info("-->> ", "clients.addWhereToQueryClientsSelect()")
	cnt := 0

	// Address
	if req.Address != "" {
		q += " LEFT JOIN addresses_ref ad ON ad.client_id = c.id"
		q += " WHERE ad.address ILIKE '%" + req.Address + "%'"
		cnt++
	}

	// ClientID
	if req.Id != 0 {
		if cnt > 0 {
			q += fmt.Sprintf(" AND c.id = %d", req.Id)
		} else {
			q += fmt.Sprintf(" WHERE c.id = %d", req.Id)
		}
		cnt++
	}

	// Firstname
	if req.FirstName != "" {
		if cnt > 0 {
			q += " AND c.firstname ILIKE '%" + req.FirstName + "%'"
		} else {
			q += " WHERE c.firstname ILIKE '%" + req.FirstName + "%'"
		}
		cnt++
	}

	// Middlename
	if req.MiddleName != "" {
		if cnt > 0 {
			q += " AND c.middlename ILIKE '%" + req.MiddleName + "%'"
		} else {
			q += " WHERE c.middlename ILIKE '%" + req.MiddleName + "%'"
		}
		cnt++
	}

	// Lastname
	if req.LastName != "" {
		if cnt > 0 {
			q += " AND c.lastname ILIKE '%" + req.LastName + "%'"
		} else {
			q += " WHERE c.lastname ILIKE '%" + req.LastName + "%'"
		}
		cnt++
	}

	// Fullname
	if req.FullName != "" {
		if cnt > 0 {
			q += " AND c.fullname ILIKE '%" + req.FullName + "%'"
		} else {
			q += " WHERE c.fullname ILIKE '%" + req.FullName + "%'"
		}
		cnt++
	}

	// Regnumber
	if req.RegNumber != 0 {
		if cnt > 0 {
			q += " AND c.regnumber = " + fmt.Sprintf("%d", req.RegNumber)
		} else {
			q += " WHERE c.regnumber = " + fmt.Sprintf("%d", req.RegNumber)
		}
		cnt++
	}

	// Phone
	if req.Phone != "" {
		if cnt > 0 {
			q += " AND c.phone ILIKE '%" + req.Phone + "%'"
		} else {
			q += " WHERE c.phone ILIKE '%" + req.Phone + "%'"
		}
		cnt++
	}

	// Email
	if req.Email != "" {
		if cnt > 0 {
			q += " AND c.email ILIKE '%" + req.Email + "%'"
		} else {
			q += " WHERE c.email ILIKE '%" + req.Email + "%'"
		}
		cnt++
	}

	// Comment
	if req.Comment != "" {
		if cnt > 0 {
			q += " AND c.comment ILIKE '%" + req.Comment + "%'"
		} else {
			q += " WHERE c.comment ILIKE '%" + req.Comment + "%'"
		}
		cnt++
	}

	// IsBad
	switch req.IsBad {
	case api.ClientsMS_Bool_IS_TRUE:
		if cnt > 0 {
			q += " AND c.isbad = true"
		} else {
			q += " WHERE c.isbad = true"
		}
		cnt++
	case api.ClientsMS_Bool_IS_FALSE:
		if cnt > 0 {
			q += " AND c.isbad = false"
		} else {
			q += " WHERE c.isbad = false"
		}
	}

	// IsDeleted
	switch req.IsDeleted {
	case api.ClientsMS_Bool_IS_TRUE:
		if cnt > 0 {
			q += " AND c.isdeleted = true"
		} else {
			q += " WHERE c.isdeleted = true"
		}
		cnt++
	case api.ClientsMS_Bool_IS_FALSE:
		if cnt > 0 {
			q += " AND c.isdeleted = false"
		} else {
			q += " WHERE c.isdeleted = false"
		}
		cnt++
	}

	lgr.LOG.Info("<<-- ", "clients.addWhereToQueryClientsSelect()")
	return hlp.DeleteTabsAndNewLinesSymbols(q)
}

//**************************************************************
// Add ORDER to queries
//**************************************************************

// addOrderToQueryClientsSelect - added ORDER to query SELECT
func addOrderToQueryClientsSelect(q string, req *api.RequestClient) string {

	lgr.LOG.Info("-->> ", "clients.addOrderToQueryClientsSelect()")

	strOrder := ""
	switch req.Order {
	case api.ClientsMS_ClientsOrder_BY_CLIENT_ID:
		strOrder = " ORDER BY c.id"
	case api.ClientsMS_ClientsOrder_BY_CLIENT_NAME:
		strOrder = " ORDER BY c.firstname, c.middlename, c.lastname"
	case api.ClientsMS_ClientsOrder_BY_CLIENT_CREATED_AT:
		strOrder = " ORDER BY c.created_at"
	case api.ClientsMS_ClientsOrder_BY_CLIENT_UPDATED_AT:
		strOrder = " ORDER BY c.updated_at"
	default:
		strOrder = " ORDER BY c.created_at"
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
	lgr.LOG.Info("_QUERY_: ", hlp.DeleteTabsAndNewLinesSymbols(q))

	lgr.LOG.Info("<<-- ", "clients.addOrderToQueryClientsSelect()")
	return q
}

//**************************************************************
// Add LIMIT to queries
//**************************************************************

// addLimitToQueryClientsSelect - added LIMIT to query SELECT
func addLimitToQueryClientsSelect(q string) string {
	lgr.LOG.Info("-->> ", "clients.addLimitToQueryClientsSelect()")

	q += " LIMIT " + fmt.Sprintf("%d", cfg.CFG.Database.Limit)

	lgr.LOG.Info("_QUERY_: ", hlp.DeleteTabsAndNewLinesSymbols(q))
	lgr.LOG.Info("<<-- ", "clients.addLimitToQueryClientsSelect()")
	return q
}

//**************************************************************
// Generating a queries
//**************************************************************

// GenQueryClientsSelect - generates a query for clients select
func GenQueryClientsSelect(req *api.RequestClient) string {

	lgr.LOG.Info("-->> ", "clients.GenQueryClientsSelect()")
	var q string

	lgr.LOG.Info("_ACTION_: ", "generating query")
	q = createQueryClientsSelect(req)

	lgr.LOG.Info("_ACTION_: ", "adding WHERE to query")
	q = addWhereToQueryClientsSelect(q, req)

	lgr.LOG.Info("_ACTION_: ", "adding ORDER to query")
	q = addOrderToQueryClientsSelect(q, req)

	lgr.LOG.Info("_ACTION_: ", "adding LIMIT to query")
	q = addLimitToQueryClientsSelect(q)

	lgr.LOG.Info("_QUERY_: ", hlp.DeleteTabsAndNewLinesSymbols(q))
	lgr.LOG.Info("<<-- ", "clients.GenQueryClientsSelect()")
	return q
}
