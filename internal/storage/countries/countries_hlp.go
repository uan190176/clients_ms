package countries

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
	//latDigOnly = "^[a-zA-Z0-9_]+$"
	rusSpace = "^[а-яА-Я\\s]+$"
)

//**************************************************************
// Check required fields
//**************************************************************

// CheckAndNormalizeRequiredFieldsCountryInsert - checked required fields to INSERT
func CheckAndNormalizeRequiredFieldsCountryInsert(req *api.RequestCountry) st.ResponseStatus {
	lgr.LOG.Info("-->> ", "countries.CheckAndNormalizeRequiredFieldsCountryInsert()")

	var stat st.ResponseStatus

	//AuthorId
	lgr.LOG.Info("_ACTION_: ", "checking AuthorId")
	if req.AuthorId == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " AuthorId", " must contain a value"))
		return st.GetStatus(401, " AuthorId", " must contain a value")
	}

	//CountryName
	lgr.LOG.Info("_ACTION_: ", "checking CountryName")
	if req.CountryName == "" {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " CountryName", " must contain a value"))
		return st.GetStatus(401, " CountryName", " must contain a value")
	}

	lgr.LOG.Info("_ACTION_: ", "normalize CountryName")
	req.CountryName, stat = st.NormalizeStringByTemplate(req.CountryName, rusSpace, true)
	if stat.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", stat)
		return stat
	}

	lgr.LOG.Info("<<-- ", "countries.CheckAndNormalizeRequiredFieldsCountryInsert()")
	return st.GetStatus(100)
}

// CheckAndNormalizeRequiredFieldsCountryUpdate - checked required fields to UPDATE
func CheckAndNormalizeRequiredFieldsCountryUpdate(req *api.RequestCountry) st.ResponseStatus {
	lgr.LOG.Info("-->> ", "countries.CheckAndNormalizeRequiredFieldsCountryUpdate()")

	//AuthorId
	lgr.LOG.Info("_ACTION_: ", "checking AuthorId")
	if req.AuthorId == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " AuthorId", " must contain a value"))
		return st.GetStatus(401, " AuthorId", " must contain a value")
	}

	//CountryId
	lgr.LOG.Info("_ACTION_: ", "checking CountryId")
	if req.CountryId == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " CountryId", "must contain a value"))
		return st.GetStatus(401, " CountryId", "must contain a value")
	}

	//CountryName || CountryComment
	if req.CountryIsDeleted == api.ClientsMS_Bool_IS_ANY {
		lgr.LOG.Info("_ACTION_: ", "checking CountryName or comment")
		if req.CountryName == "" && req.CountryComment == "" {
			lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " CountryName || CountryComment", "must contain a value"))
			return st.GetStatus(401, " CountryName || CountryComment", "must contain a value")
		}
	}

	//Normalize CountryName
	var stat st.ResponseStatus
	if req.CountryName != "" {
		lgr.LOG.Info("_ACTION_: ", "normalize CountryName")
		req.CountryName, stat = st.NormalizeStringByTemplate(req.CountryName, rusSpace, true)
		if stat.Code > 100 {
			lgr.LOG.Warn("_WARN_: ", stat)
			return stat
		}
	}

	lgr.LOG.Info("<<-- ", "countries.CheckAndNormalizeRequiredFieldsCountryUpdate()")
	return st.GetStatus(100)
}

// CheckRequiredFieldsCountriesDeletionFlagUpdate - checked required fields to UPDATE deletion flags
func CheckRequiredFieldsCountriesDeletionFlagUpdate(req *api.RequestCountriesDeletionFlags) st.ResponseStatus {
	lgr.LOG.Info("-->> ", "countries.CheckRequiredFieldsCountriesDeletionFlagUpdate()")

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

	//CountryId
	lgr.LOG.Info("_ACTION_: ", "checking CountryId")
	if len(req.CountryId) == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " CountryId", " must contain a value"))
		return st.GetStatus(401, " CountryId", " must contain a value")
	}

	lgr.LOG.Info("<<-- ", "countries.CheckRequiredFieldsCountriesDeletionFlagUpdate()")
	return st.GetStatus(100)
}

//**************************************************************
// Create queries
//**************************************************************

// createQueryCountriesSelect - returns query to SELECT
func createQueryCountriesSelect() string {
	lgr.LOG.Info("-->> ", "countries.CreateQueryCountriesSelect()")
	q := fmt.Sprintf("SELECT %s FROM countries", createViewCountries())
	lgr.LOG.Info("_QUERY_: ", q)
	lgr.LOG.Info("<<-- ", "countries.CreateQueryCountriesSelect()")
	return q
}

// CreateQueryCountryInsert - returns query to INSERT
func CreateQueryCountryInsert(req *api.RequestCountry) (string, st.ResponseStatus) {
	lgr.LOG.Info("-->> ", "countries.CreateQueryCountryInsert()")

	q := `INSERT INTO countries (name, comment, created_by)
        	VALUES ($1, $2, $3) RETURNING `

	q += createViewCountries()

	lgr.LOG.Info("_QUERY_: ", q)
	lgr.LOG.Info("<<-- ", "countries.CreateQueryCountryInsert()")
	return q, st.GetStatus(100)
}

// CreateQueryCountryUpdate - returns query to UPDATE
func CreateQueryCountryUpdate(req *api.RequestCountry) (string, st.ResponseStatus) {
	lgr.LOG.Info("-->> ", "countries.CreateQueryCountryUpdate()")

	sFields := ""
	sValues := ""
	cnt := 0

	if req.CountryName != "" {
		sFields += "name, "
		sValues += fmt.Sprintf("'%s', ", req.CountryName)
		cnt++
	}
	if req.CountryComment != "" {
		sFields += "comment, "
		sValues += fmt.Sprintf("'%s', ", req.CountryComment)
		cnt++
	}
	// Is deleted
	if req.CountryIsDeleted == api.ClientsMS_Bool_IS_TRUE || req.CountryIsDeleted == api.ClientsMS_Bool_IS_FALSE {
		sFields += "isdeleted, "
		if req.CountryIsDeleted == api.ClientsMS_Bool_IS_TRUE {
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
		q = "UPDATE countries SET (" + sFields + ") = (" + sValues + ")"
	} else {
		q = "UPDATE countries SET " + sFields + " = " + sValues
	}

	q += fmt.Sprintf(" WHERE id = %d RETURNING ", req.CountryId)
	q += createViewCountries()

	lgr.LOG.Info("_QUERY_: ", q)

	lgr.LOG.Info("<<-- ", "countries.CreateQueryCountryUpdate()")
	return q, st.GetStatus(100)
}

// CreateQueryCountriesDeletionFlagsUpdate - returns query to UPDATE countries deletion flags
func CreateQueryCountriesDeletionFlagsUpdate(req *api.RequestCountriesDeletionFlags) string {
	lgr.LOG.Info("-->> ", "CreateQueryCountriesDeletionFlagsUpdate()")
	q := fmt.Sprintf("UPDATE countries SET isdeleted = %t, changed_by = %d WHERE id IN %s RETURNING ",
		hlp.GetBool(req.IsDeleted), req.AuthorId, hlp.Uint64ArrayToString(req.CountryId))
	q += createViewCountries()
	lgr.LOG.Info("_QUERY_: ", q)
	lgr.LOG.Info("<<-- ", "CreateQueryCountriesDeletionFlagsUpdate()")
	return q
}

// **************************************************************
// Create view
// **************************************************************
func createViewCountries() string {
	return `id AS country_id,
    			name AS country_name,
    			concat(comment, '') AS country_comment,
    			to_char(created, 'YYYY-MM-DD HH24:MI:SS'::text) AS country_created,
    			created_by AS country_created_by,
                ( SELECT u.fullname FROM users u WHERE u.id = created_by ) AS country_created_by_name,
    			to_char(changed, 'YYYY-MM-DD HH24:MI:SS'::text) AS country_changed,
    			changed_by AS country_changed_by,
    			( SELECT u.fullname FROM users u WHERE u.id = changed_by ) AS country_changed_by_name,
    			isdeleted AS country_isdeleted`
}

//**************************************************************
// Add WHERE to queries
//**************************************************************

// addWhereToQueryCountriesSelect - added WHERE to query SELECT
func addWhereToQueryCountriesSelect(q string, req *api.RequestCountry) string {
	lgr.LOG.Info("-->> ", "countries.AddWhereToQueryCountriesSelect()")
	cnt := 0

	// ID
	if req.CountryId != 0 {
		q += fmt.Sprintf(" WHERE country_id = %d", req.CountryId)
		cnt++
	}

	// NAME
	if req.CountryName != "" {
		if cnt > 0 {
			q += " AND country_name ILIKE '%" + req.CountryName + "%'"
		} else {
			q += " WHERE country_name ILIKE '%" + req.CountryName + "%'"
		}
		cnt++
	}
	// COMMENT
	if req.CountryComment != "" {
		if cnt > 0 {
			q += " AND country_comment ILIKE '%" + req.CountryComment + "%'"
		} else {
			q += " WHERE country_comment ILIKE '%" + req.CountryComment + "%'"
		}
		cnt++
	}

	// Isdeleted

	switch req.CountryIsDeleted {
	case api.ClientsMS_Bool_IS_TRUE:
		if cnt > 0 {
			q += " AND country_isdeleted = true"
		} else {
			q += " WHERE country_isdeleted = true"
		}
		cnt++
	case api.ClientsMS_Bool_IS_FALSE:
		if cnt > 0 {
			q += " AND country_isdeleted = false"
		} else {
			q += " WHERE country_isdeleted = false"
		}
		cnt++
	}

	lgr.LOG.Info("_QUERY_: ", q)
	lgr.LOG.Info("<<-- ", "countries.AddWhereToQueryCountriesSelect()")
	return q
}

//**************************************************************
// Add ORDER to queries
//**************************************************************

// addOrderToQueryCountriesSelect - added ORDER to query SELECT
func addOrderToQueryCountriesSelect(q string, req *api.RequestCountry) string {

	lgr.LOG.Info("-->> ", "countries.AddOrderToQueryCountriesSelect()")

	strOrder := ""
	switch req.CountryOrder {
	case api.ClientsMS_CountryOrder_BY_COUNTRY_ID:
		strOrder = " ORDER BY country_id"
	case api.ClientsMS_CountryOrder_BY_COUNTRY_NAME:
		strOrder = " ORDER BY country_name"
	case api.ClientsMS_CountryOrder_BY_COUNTRY_CREATED:
		strOrder = " ORDER BY country_created"
	case api.ClientsMS_CountryOrder_BY_COUNTRY_CHANGED:
		strOrder = " ORDER BY country_changed"
	default:
		strOrder = " ORDER BY country_id"
	}

	strDirection := ""
	switch req.CountryOrderType {
	case api.ClientsMS_OrderType_ASC:
		strDirection = " ASC"
	case api.ClientsMS_OrderType_DESC:
		strDirection = " DESC"
	default:
		strDirection = " ASC"
	}

	q += strOrder + strDirection
	lgr.LOG.Info("_QUERY_: ", q)

	lgr.LOG.Info("<<-- ", "countries.AddOrderToQueryCountriesSelect()")
	return q
}

//**************************************************************
// Generating a queries
//**************************************************************

// GenQueryCountriesSelect - generates a query for countries select
func GenQueryCountriesSelect(req *api.RequestCountry) string {

	lgr.LOG.Info("-->> ", "countries.GenQueryCountriesSelect()")
	var q string

	lgr.LOG.Info("_ACTION_: ", "generating query")
	q = createQueryCountriesSelect()

	lgr.LOG.Info("_ACTION_: ", "adding WHERE to query")
	q = addWhereToQueryCountriesSelect(q, req)

	lgr.LOG.Info("_ACTION_: ", "adding ORDER to query")
	q = addOrderToQueryCountriesSelect(q, req)

	lgr.LOG.Info("_QUERY_: ", q)
	lgr.LOG.Info("<<-- ", "countries.GenQueryCountriesSelect()")
	return q
}

//**************************************************************
// Converting a request to slice of interfaces
//**************************************************************

// ConvertToInterfaceSlice - converting a request to slice of interfaces
func ConvertToInterfaceSlice(req *api.RequestCountriesDeletionFlags) [][]interface{} {
	var result [][]interface{}
	for _, countryId := range req.CountryId {
		result = append(result, []interface{}{countryId, req.IsDeleted, req.AuthorId})
	}
	return result
}
