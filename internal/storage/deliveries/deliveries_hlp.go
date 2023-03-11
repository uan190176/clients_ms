package deliveries

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

// CheckRequiredFieldsDeliveryInsert - checks required fields to deliveries reference INSERT
func CheckRequiredFieldsDeliveryInsert(req *api.RequestDelivery) st.ResponseStatus {
	lgr.LOG.Info("-->> ", "deliveries.CheckRequiredFieldsDeliveriesInsert()")

	//AuthorId
	lgr.LOG.Info("_ACTION_: ", "checking AuthorId")
	if req.AuthorId == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " AuthorId", " must contain a value"))
		return st.GetStatus(401, " AuthorId", " must contain a value")
	}

	//DeliveryName
	lgr.LOG.Info("_ACTION_: ", "checking delivery name")
	if req.Name == "" {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " Name", " must contain a value"))
		return st.GetStatus(401, " Name", " must contain a value")
	}

	lgr.LOG.Info("<<-- ", "deliveries.CheckRequiredFieldsDeliveriesInsert()")
	return st.GetStatus(100)
}

// CheckRequiredFieldsDeliveryUpdate - checks required fields to deliveries reference UPDATE
func CheckRequiredFieldsDeliveryUpdate(req *api.RequestDelivery) st.ResponseStatus {
	lgr.LOG.Info("-->> ", "countries.CheckAndNormalizeRequiredFieldsCountryUpdate()")

	//AuthorId
	lgr.LOG.Info("_ACTION_: ", "checking AuthorId")
	if req.AuthorId == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " AuthorId", " must contain a value"))
		return st.GetStatus(401, " AuthorId", " must contain a value")
	}

	//DeliveryId
	lgr.LOG.Info("_ACTION_: ", "checking delivery id")
	if req.Id == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " Id", "must contain a value"))
		return st.GetStatus(401, " Id", "must contain a value")
	}

	//Name || Comment
	if req.IsDeleted == api.ClientsMS_Bool_IS_ANY {
		lgr.LOG.Info("_ACTION_: ", "checking delivery name or comment")
		if req.Name == "" && req.Comment == "" {
			lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " Name || Comment", "must contain a value"))
			return st.GetStatus(401, " Name || Comment", "must contain a value")
		}
	}

	lgr.LOG.Info("<<-- ", "countries.CheckAndNormalizeRequiredFieldsCountryUpdate()")
	return st.GetStatus(100)
}

// CheckRequiredFieldsDeliveriesDeletionFlagUpdate - checked required fields to UPDATE deletion flags
func CheckRequiredFieldsDeliveriesDeletionFlagUpdate(req *api.RequestDeliveriesDeletionFlags) st.ResponseStatus {
	lgr.LOG.Info("-->> ", "deliveries.CheckRequiredFieldsDeliveriesDeletionFlagUpdate()")

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

	lgr.LOG.Info("<<-- ", "deliveries.CheckRequiredFieldsDeliveriesDeletionFlagUpdate()")
	return st.GetStatus(100)
}

//**************************************************************
// Create queries
//**************************************************************

// createQueryDeliveriesSelect - returns query to SELECT
func createQueryDeliveriesSelect() string {
	lgr.LOG.Info("-->> ", "deliveries.createQueryDeliveriesSelect()")
	q := fmt.Sprintf("SELECT %s FROM deliveries_ref", createViewDeliveries())
	lgr.LOG.Info("_QUERY_: ", q)
	lgr.LOG.Info("<<-- ", "deliveries.createQueryDeliveriesSelect()")
	return q
}

// CreateQueryDeliveryInsert - returns query to INSERT
func CreateQueryDeliveryInsert(req *api.RequestDelivery) (string, st.ResponseStatus) {
	lgr.LOG.Info("-->> ", "deliveries.CreateQueryDeliveryInsert()")

	q := fmt.Sprintf("INSERT INTO deliveries_ref (name, comment, created_by) VALUES ('%s', '%s', %d) RETURNING ", req.Name, req.Comment, req.AuthorId)

	q += createViewDeliveries()

	lgr.LOG.Info("_QUERY_: ", q)
	lgr.LOG.Info("<<-- ", "deliveries.CreateQueryDeliveryInsert()")
	return q, st.GetStatus(100)
}

// CreateQueryDeliveryUpdate - returns query to UPDATE
func CreateQueryDeliveryUpdate(req *api.RequestDelivery) (string, st.ResponseStatus) {
	lgr.LOG.Info("-->> ", "deliveries.CreateQueryDeliveryUpdate()")

	sFields := ""
	sValues := ""
	cnt := 0

	if req.Name != "" {
		sFields += "name, "
		sValues += fmt.Sprintf("'%s', ", req.Name)
		cnt++
	}
	if req.Comment != "" {
		sFields += "comment, "
		sValues += fmt.Sprintf("'%s', ", req.Comment)
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
		q = "UPDATE deliveries_ref SET (" + sFields + ") = (" + sValues + ")"
	} else {
		q = "UPDATE deliveries_ref SET " + sFields + " = " + sValues
	}

	q += fmt.Sprintf(" WHERE id = %d RETURNING ", req.Id)
	q += createViewDeliveries()

	lgr.LOG.Info("_QUERY_: ", q)

	lgr.LOG.Info("<<-- ", "deliveries.CreateQueryDeliveryUpdate()")
	return q, st.GetStatus(100)
}

// CreateQueryDeliveriesDeletionFlagsUpdate - returns query to UPDATE countries deletion flags
func CreateQueryDeliveriesDeletionFlagsUpdate(req *api.RequestDeliveriesDeletionFlags) string {
	lgr.LOG.Info("-->> ", "deliveries.CreateQueryDeliveriesDeletionFlagsUpdate()")
	q := fmt.Sprintf("UPDATE deliveries_ref SET isdeleted = %t, updated_by = %d WHERE id IN %s RETURNING ",
		hlp.GetBool(req.IsDeleted), req.AuthorId, hlp.Uint64ArrayToString(req.Ids))
	q += createViewDeliveries()
	lgr.LOG.Info("_QUERY_: ", q)
	lgr.LOG.Info("<<-- ", "deliveries.CreateQueryDeliveriesDeletionFlagsUpdate()")
	return q
}

// **************************************************************
// Create view
// **************************************************************
func createViewDeliveries() string {
	return `id AS delivery_id,
    			name AS delivery_name,
    			concat(comment, '') AS delivery_comment,
    			to_char(created_at, 'YYYY-MM-DD HH24:MI:SS'::text) AS delivery_created_at,
    			created_by AS delivery_created_by,
                ( SELECT u.fullname FROM users_ref u WHERE u.id = created_by ) AS delivery_created_by_name,
    			to_char(updated_at, 'YYYY-MM-DD HH24:MI:SS'::text) AS delivery_updated_at,
    			updated_by AS delivery_updated_by,
    			( SELECT u.fullname FROM users_ref u WHERE u.id = updated_by ) AS delivery_updated_by_name,
    			isdeleted AS delivery_isdeleted`
}

//**************************************************************
// Add WHERE to queries
//**************************************************************

// addWhereToQueryDeliveriesSelect - added WHERE to query SELECT
func addWhereToQueryDeliveriesSelect(q string, req *api.RequestDelivery) string {
	lgr.LOG.Info("-->> ", "deliveries.addWhereToQueryDeliveriesSelect()")
	cnt := 0

	// ID
	if req.Id != 0 {
		q += fmt.Sprintf(" WHERE delivery_id = %d", req.Id)
		cnt++
	}

	// NAME
	if req.Name != "" {
		if cnt > 0 {
			q += " AND delivery_name ILIKE '%" + req.Name + "%'"
		} else {
			q += " WHERE delivery_name ILIKE '%" + req.Name + "%'"
		}
		cnt++
	}
	// COMMENT
	if req.Comment != "" {
		if cnt > 0 {
			q += " AND delivery_comment ILIKE '%" + req.Comment + "%'"
		} else {
			q += " WHERE delivery_comment ILIKE '%" + req.Comment + "%'"
		}
		cnt++
	}

	// Isdeleted

	switch req.IsDeleted {
	case api.ClientsMS_Bool_IS_TRUE:
		if cnt > 0 {
			q += " AND delivery_isdeleted = true"
		} else {
			q += " WHERE delivery_isdeleted = true"
		}
		cnt++
	case api.ClientsMS_Bool_IS_FALSE:
		if cnt > 0 {
			q += " AND delivery_isdeleted = false"
		} else {
			q += " WHERE delivery_isdeleted = false"
		}
		cnt++
	}

	lgr.LOG.Info("_QUERY_: ", q)
	lgr.LOG.Info("<<-- ", "deliveries.addWhereToQueryDeliveriesSelect()")
	return q
}

//**************************************************************
// Add ORDER to queries
//**************************************************************

// addOrderToQueryDeliveriesSelect - added ORDER to query SELECT
func addOrderToQueryDeliveriesSelect(q string, req *api.RequestDelivery) string {

	lgr.LOG.Info("-->> ", "deliveries.addOrderToQueryDeliveriesSelect()")

	strOrder := ""
	switch req.Order {
	case api.ClientsMS_DeliveriesOrder_BY_DELIVERY_ID:
		strOrder = " ORDER BY delivery_id"
	case api.ClientsMS_DeliveriesOrder_BY_DELIVERY_NAME:
		strOrder = " ORDER BY delivery_name"
	case api.ClientsMS_DeliveriesOrder_BY_DELIVERY_CREATED_AT:
		strOrder = " ORDER BY delivery_created"
	case api.ClientsMS_DeliveriesOrder_BY_DELIVERY_UPDATED_AT:
		strOrder = " ORDER BY delivery_changed"
	default:
		strOrder = " ORDER BY delivery_id"
	}

	strDirection := ""
	switch req.OrderType {
	case api.ClientsMS_OrderType_ASC:
		strDirection = " ASC"
	case api.ClientsMS_OrderType_DESC:
		strDirection = " DESC"
	default:
		strDirection = " ASC"
	}

	q += strOrder + strDirection
	lgr.LOG.Info("_QUERY_: ", q)

	lgr.LOG.Info("<<-- ", "deliveries.addOrderToQueryDeliveriesSelect()")
	return q
}

//**************************************************************
// Generating a queries
//**************************************************************

// GenQueryDeliveriesSelect - generates a query for countries select
func GenQueryDeliveriesSelect(req *api.RequestDelivery) string {

	lgr.LOG.Info("-->> ", "deliveries.GenQueryDeliveriesSelect()")
	var q string

	lgr.LOG.Info("_ACTION_: ", "generating query")
	q = createQueryDeliveriesSelect()

	lgr.LOG.Info("_ACTION_: ", "adding WHERE to query")
	q = addWhereToQueryDeliveriesSelect(q, req)

	lgr.LOG.Info("_ACTION_: ", "adding ORDER to query")
	q = addOrderToQueryDeliveriesSelect(q, req)

	lgr.LOG.Info("_QUERY_: ", q)
	lgr.LOG.Info("<<-- ", "deliveries.GenQueryDeliveriesSelect()")
	return q
}
