package address_types

import (
	"clients_ms/internal/api"
	hlp "clients_ms/internal/helpers"
	lgr "clients_ms/pkg/logger"
	"fmt"
	st "github.com/uan190176/statuses"
	"strings"
)

const (
	// Templates for checking fields in request
	latDigOnly = "^[a-zA-Z0-9_]+$"
	rusSpace   = "^[а-яА-Я\\s]+$"
)

//**************************************************************
// Check required fields
//**************************************************************

// CheckAndNormaliseRequiredFieldsAddressTypeInsert - checks and normalize required fields
func CheckAndNormaliseRequiredFieldsAddressTypeInsert(req *api.RequestAddressType) st.ResponseStatus {
	lgr.LOG.Info("-->> ", "address_types.CheckAndNormaliseRequiredFieldsAddressTypeInsert()")

	var stat st.ResponseStatus

	//AuthorId
	lgr.LOG.Info("_ACTION_: ", "checking address type author id")
	if req.AuthorId == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, "AuthorId", "must contain a value"))
		return st.GetStatus(401, "AuthorId", "must contain a value")
	}

	// AddrTypeName
	lgr.LOG.Info("_ACTION_: ", "checking address type name")
	if req.AddrTypeName == "" {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, "AddrTypeName", "must contain a value"))
		return st.GetStatus(401, "AddrTypeName", "must contain a value")
	}

	lgr.LOG.Info("_ACTION_: ", "normalize address type name")
	req.AddrTypeName, stat = st.NormalizeStringByTemplate(req.AddrTypeName, rusSpace, true)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return stat
	}

	// AddrTypeCode
	if req.AddrTypeCode == "" {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, "AddrTypeCode", "must contain a value"))
		return st.GetStatus(401, "AddrTypeCode", "must contain a value")
	}

	lgr.LOG.Info("_ACTION_: ", "normalize address type code")
	req.AddrTypeCode, stat = st.NormalizeStringByTemplate(req.AddrTypeCode, latDigOnly, true, true)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return stat
	}

	lgr.LOG.Info("<<-- ", "address_types.CheckAndNormaliseRequiredFieldsAddressTypeInsert()")
	return st.GetStatus(100)
}

// CheckAndNormaliseRequiredFieldsAddressTypeUpdate - checks and normalize required fields
func CheckAndNormaliseRequiredFieldsAddressTypeUpdate(req *api.RequestAddressType) st.ResponseStatus {
	lgr.LOG.Info("-->> ", "address_types.CheckAndNormaliseRequiredFieldsAddressTypeUpdate()")

	var stat st.ResponseStatus
	var cnt int

	//AddrTypeId
	lgr.LOG.Info("_ACTION_: ", "checking address type id")
	if req.AddrTypeID == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, "AddrTypeID", "must contain a value"))
		return st.GetStatus(401, "AddrTypeID", "must contain a value")
	}

	//AuthorId
	lgr.LOG.Info("_ACTION_: ", "checking address type author id")
	if req.AuthorId == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, "AuthorId", "must contain a value"))
		return st.GetStatus(401, "AuthorId", "must contain a value")
	}

	if req.AddrTypeIsDeleted == api.ClientsMS_Bool_IS_ANY && req.AddrTypeNeedsNormalizing == api.ClientsMS_Bool_IS_ANY && req.AddrTypeNeedsCleaning == api.ClientsMS_Bool_IS_ANY {

		// AddrTypeName
		lgr.LOG.Info("_ACTION_: ", "checking address type name")
		if req.AddrTypeName != "" {
			lgr.LOG.Info("_ACTION_: ", "normalize address type name")
			req.AddrTypeName, stat = st.NormalizeStringByTemplate(req.AddrTypeName, rusSpace, true)
			if stat.Code > 100 {
				lgr.LOG.Warn("_WARN_: ", stat)
				return stat
			}
			cnt++
		}

		// AddrTypeCode
		if req.AddrTypeCode != "" {
			lgr.LOG.Info("_ACTION_: ", "normalize address type code")
			req.AddrTypeCode, stat = st.NormalizeStringByTemplate(req.AddrTypeCode, latDigOnly, true, true)
			if stat.Code > 100 {
				lgr.LOG.Warn("_WARN_: ", stat)
				return stat
			}
			cnt++
		}

		// AddrTypeComment
		if req.AddrTypeComment != "" {
			cnt++
		}

		// Check count
		if cnt == 0 {
			lgr.LOG.Warn("_WARN_: ", st.GetStatus(403))
			return st.GetStatus(403)
		}
	}

	lgr.LOG.Info("<<-- ", "address_types.CheckAndNormaliseRequiredFieldsAddressTypeUpdate()")
	return st.GetStatus(100)
}

// CheckRequiredFieldsAddressesTypesDeletionFlagsUpdate - checking required fields
func CheckRequiredFieldsAddressesTypesDeletionFlagsUpdate(req *api.RequestAddressesTypesDeletionFlags) st.ResponseStatus {
	lgr.LOG.Info("-->> ", "address_types.CheckRequiredFieldsAddressesTypesDeletionFlagsUpdate()")

	//IDs
	if len(req.AddrTypeID) == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, "AddrTypeID", "must contain array of IDs"))
		return st.GetStatus(401, "AddrTypeID", "must contain array of IDs")
	}

	//AuthorId
	lgr.LOG.Info("_ACTION_: ", "checking address type author id")
	if req.AuthorId == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, "AuthorId", "must contain a value"))
		return st.GetStatus(401, "AuthorId", "must contain a value")
	}

	//Isdeleted
	lgr.LOG.Info("_ACTION_: ", "checking address type isdeleted")
	if req.AddrTypeIsDeleted == api.ClientsMS_Bool_IS_ANY {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, "AddrTypeIsDeleted", "must contain a value true or false"))
		return st.GetStatus(401, "AuthorId", "must contain a value true or false")
	}
	lgr.LOG.Info("<<-- ", "address_types.CheckRequiredFieldsAddressesTypesDeletionFlagsUpdate()")
	return st.GetStatus(100)
}

//**************************************************************
// Create queries
//**************************************************************

// createQueryAddressesTypesSelect - returns query to SELECT
func createQueryAddressesTypesSelect() string {
	lgr.LOG.Info("-->> ", "address_types.CreateQueryAddressesTypesSelect()")
	q := fmt.Sprintf("SELECT %s FROM addresses_types", createViewAddressesTypes())
	lgr.LOG.Info("_QUERY_: ", q)
	lgr.LOG.Info("<<-- ", "address_types.CreateQueryAddressesTypesSelect()")
	return q
}

// CreateQueryAddressTypeInsert - returns query to INSERT
func CreateQueryAddressTypeInsert(req *api.RequestAddressType) (string, st.ResponseStatus) {
	lgr.LOG.Info("-->> ", "address_types.CreateQueryAddressTypeInsert()")

	//q := fmt.Sprintf("INSERT INTO addresses_types (name, code, comment, created_by)"+
	//	" VALUES ('%s', '%s', '%s', %d) RETURNING id", req.AddrTypeName, req.AddrTypeCode, req.AddrTypeComment, req.AuthorId)

	q := `INSERT INTO addresses_types (name, code, comment, created_by)
        VALUES ($1, $2, $3, $4) RETURNING `
	q += createViewAddressesTypes()

	lgr.LOG.Info("_QUERY_: ", q)
	lgr.LOG.Info("<<-- ", "address_types.CreateQueryAddressTypeInsert()")
	return q, st.GetStatus(100)
}

// CreateQueryAddressTypeUpdate - returns query to UPDATE
func CreateQueryAddressTypeUpdate(req *api.RequestAddressType) (string, st.ResponseStatus) {
	lgr.LOG.Info("-->> ", "address_types.CreateQueryAddressTypeUpdate()")

	sFields := ""
	sValues := ""
	cnt := 0

	if req.AddrTypeName != "" {
		sFields += "name, "
		sValues += fmt.Sprintf("'%s', ", req.AddrTypeName)
		cnt++
	}

	if req.AddrTypeCode != "" {
		sFields += "code, "
		sValues += fmt.Sprintf("'%s', ", req.AddrTypeCode)
		cnt++
	}

	if req.AddrTypeComment != "" {
		sFields += "comment, "
		sValues += fmt.Sprintf("'%s', ", req.AddrTypeComment)
		cnt++
	}

	// Is deleted
	if req.AddrTypeIsDeleted == api.ClientsMS_Bool_IS_TRUE || req.AddrTypeIsDeleted == api.ClientsMS_Bool_IS_FALSE {
		sFields += "isdeleted, "
		if req.AddrTypeIsDeleted == api.ClientsMS_Bool_IS_TRUE {
			sValues += fmt.Sprintf("%t, ", true)
		} else {
			sValues += fmt.Sprintf("%t, ", false)
		}
		cnt++
	}

	// Needs normalizing
	if req.AddrTypeNeedsNormalizing == api.ClientsMS_Bool_IS_TRUE || req.AddrTypeNeedsNormalizing == api.ClientsMS_Bool_IS_FALSE {
		sFields += "needs_normalizing, "
		if req.AddrTypeNeedsNormalizing == api.ClientsMS_Bool_IS_TRUE {
			sValues += fmt.Sprintf("%t, ", true)
		} else {
			sValues += fmt.Sprintf("%t, ", false)
		}
		cnt++
	}

	// Needs cleaning
	if req.AddrTypeNeedsCleaning == api.ClientsMS_Bool_IS_TRUE || req.AddrTypeNeedsCleaning == api.ClientsMS_Bool_IS_FALSE {
		sFields += "needs_cleaning, "
		if req.AddrTypeNeedsCleaning == api.ClientsMS_Bool_IS_TRUE {
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
		q = "UPDATE addresses_types SET (" + sFields + ") = (" + sValues + ")"
	} else {
		q = "UPDATE addresses_types SET " + sFields + " = " + sValues
	}

	q += fmt.Sprintf(" WHERE id = %v RETURNING ", req.AddrTypeID)
	q += createViewAddressesTypes()

	lgr.LOG.Info("_QUERY_: ", q)

	lgr.LOG.Info("<<-- ", "address_types.CreateQueryAddressTypeUpdate()")
	return q, st.GetStatus(100)
}

// CreateQueryAddressesTypesDeletionFlagsUpdate - returns query to DELETE
func CreateQueryAddressesTypesDeletionFlagsUpdate(req *api.RequestAddressesTypesDeletionFlags) string {
	lgr.LOG.Info("-->> ", "CreateQueryAddressesTypesDeletionFlagsUpdate()")
	q := fmt.Sprintf("UPDATE addresses_types SET isdeleted = %t, changed_by = %d WHERE id IN %s RETURNING ",
		hlp.GetBool(req.AddrTypeIsDeleted), req.AuthorId, hlp.Uint64ArrayToString(req.AddrTypeID))
	q += createViewAddressesTypes()
	lgr.LOG.Info("_QUERY_: ", q)
	lgr.LOG.Info("<<-- ", "CreateQueryAddressesTypesDeletionFlagsUpdate()")
	return q
}

//**************************************************************
// Create view
//**************************************************************

// createViewAddressesTypes - returns view for result of queries INSERT/UPDATE
func createViewAddressesTypes() string {
	return `id AS address_type_id,
    		name AS address_type_name,
    		code AS address_type_code,
    		concat(comment, '') AS address_type_comment,
    		needs_normalizing AS address_type_needs_normalizing,
    		needs_cleaning AS address_type_needs_cleaning,
    		to_char(created, 'YYYY-MM-DD HH24:MI:SS'::text) AS address_type_created,
    		created_by AS address_type_created_by,
			( SELECT u.fullname FROM users u WHERE u.id = created_by ) AS address_type_created_by_name,
    		to_char(changed, 'YYYY-MM-DD HH24:MI:SS'::text) AS address_type_changed,
    		changed_by AS address_type_changed_by,
    		( SELECT u.fullname FROM users u WHERE u.id = changed_by ) AS address_type_changed_by_name,
			isdeleted AS address_type_isdeleted`
}

//**************************************************************
// Add WHERE to queries
//**************************************************************

// addWhereToQuerySelectAddressesTypes - adding where to query select from addresses_types_full
func addWhereToQuerySelectAddressesTypes(q string, req *api.RequestAddressType) string {
	lgr.LOG.Info("-->> ", "address_types.AddWhereToQuerySelectAddressesTypes()")
	cnt := 0

	// ID
	if req.AddrTypeID != 0 {
		q += fmt.Sprintf(" WHERE address_type_id IN (%v)", req.AddrTypeID)
		cnt++
	}

	// CODE
	if req.AddrTypeCode != "" {
		if cnt > 0 {
			q += " AND address_type_code ILIKE '%" + req.AddrTypeCode + "%'"
		} else {
			q += " WHERE address_type_code ILIKE '%" + req.AddrTypeCode + "%'"
		}
		cnt++
	}

	// NAME
	if req.AddrTypeName != "" {
		if cnt > 0 {
			q += " AND address_type_name ILIKE '%" + req.AddrTypeName + "%'"
		} else {
			q += " WHERE address_type_name ILIKE '%" + req.AddrTypeName + "%'"
		}
		cnt++
	}

	// COMMENT
	if req.AddrTypeComment != "" {
		if cnt > 0 {
			q += " AND address_type_comment ILIKE '%" + req.AddrTypeComment + "%'"
		} else {
			q += " WHERE address_type_comment ILIKE '%" + req.AddrTypeComment + "%'"
		}
		cnt++
	}

	// Normalising
	if req.AddrTypeNeedsNormalizing == api.ClientsMS_Bool_IS_TRUE {
		if cnt > 0 {
			q += " AND address_type_needs_normalizing = true"
		} else {
			q += " WHERE address_type_needs_normalizing = true"
		}
		cnt++
	} else if req.AddrTypeNeedsNormalizing == api.ClientsMS_Bool_IS_FALSE {
		if cnt > 0 {
			q += " AND address_type_needs_normalizing = false"
		} else {
			q += " WHERE address_type_needs_normalizing = false"
		}
		cnt++
	}

	// DaData cleaning
	if req.AddrTypeNeedsCleaning == api.ClientsMS_Bool_IS_TRUE {
		if cnt > 0 {
			q += " AND address_type_needs_cleaning = true"
		} else {
			q += " WHERE address_type_needs_cleaning = true"
		}
		cnt++
	} else if req.AddrTypeNeedsCleaning == api.ClientsMS_Bool_IS_FALSE {
		if cnt > 0 {
			q += " AND address_type_needs_cleaning = false"
		} else {
			q += " WHERE address_type_needs_cleaning = false"
		}
		cnt++
	}

	// Isdeleted
	if req.AddrTypeIsDeleted == api.ClientsMS_Bool_IS_TRUE {
		if cnt > 0 {
			q += " AND address_type_isdeleted = true"
		} else {
			q += " WHERE address_type_isdeleted = true"
		}
		cnt++
	} else if req.AddrTypeIsDeleted == api.ClientsMS_Bool_IS_FALSE {
		if cnt > 0 {
			q += " AND address_type_isdeleted = false"
		} else {
			q += " WHERE address_type_isdeleted = false"
		}
		cnt++
	}

	lgr.LOG.Info("_QUERY_: ", q)
	lgr.LOG.Info("<<-- ", "address_types.AddWhereToQuerySelectAddressesTypes()")
	return q
}

//**************************************************************
// Add ORDER to queries
//**************************************************************

// addOrderToQuerySelectAddressesTypes - adding order to query select from addresses_types_full
func addOrderToQuerySelectAddressesTypes(q string, req *api.RequestAddressType) string {
	lgr.LOG.Info("-->> ", "address_types.AddOrderToQuerySelectAddressesTypes()")
	strOrder := ""
	switch req.AddressTypesOrder {
	case api.ClientsMS_AddressTypesOrder_BY_ADDR_TYPE_ID:
		strOrder = " ORDER BY address_type_id"
	case api.ClientsMS_AddressTypesOrder_BY_ADDR_TYPE_NAME:
		strOrder = " ORDER BY address_type_name"
	case api.ClientsMS_AddressTypesOrder_BY_ADDR_TYPE_CODE:
		strOrder = " ORDER BY address_type_code"
	case api.ClientsMS_AddressTypesOrder_BY_ADDR_TYPE_CREATED:
		strOrder = " ORDER BY address_type_created"
	case api.ClientsMS_AddressTypesOrder_BY_ADDR_TYPE_CHANGED:
		strOrder = " ORDER BY address_type_changed"
	default:
		strOrder = " ORDER BY address_type_id"
	}

	strDirection := ""
	switch req.AddressTypesOrderType {
	case api.ClientsMS_OrderType_ASC:
		strDirection = " ASC"
	case api.ClientsMS_OrderType_DESC:
		strDirection = " DESC"
	default:
		strDirection = " ASC"
	}

	q += strOrder + strDirection

	lgr.LOG.Info("_QUERY_: ", q)
	lgr.LOG.Info("<<-- ", "address_types.AddOrderToQuerySelectAddressesTypes()")
	return q
}

func GenQueryAddressesTypesSelect(req *api.RequestAddressType) string {
	lgr.LOG.Info("_ACTION_: ", "generating query")
	q := createQueryAddressesTypesSelect()

	lgr.LOG.Info("_ACTION_: ", "adding WHERE to query")
	q = addWhereToQuerySelectAddressesTypes(q, req)

	lgr.LOG.Info("_ACTION_: ", "adding ORDER to query")
	q = addOrderToQuerySelectAddressesTypes(q, req)

	lgr.LOG.Info("_QUERY_: ", q)

	return q
}
