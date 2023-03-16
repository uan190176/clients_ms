package addresses

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

// CheckRequiredFieldsAddressesSelect - checks required fields to address SELECT
func CheckRequiredFieldsAddressesSelect(req *api.RequestAddress) st.ResponseStatus {
	lgr.LOG.Info("-->> ", "addresses.CheckRequiredFieldsAddressesSelect()")

	//ClientId
	lgr.LOG.Info("_ACTION_: ", "checking ClientId")
	if req.ClientId == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " ClientId", " must contain a value"))
		return st.GetStatus(401, " ClientId", " must contain a value")
	}

	lgr.LOG.Info("<<-- ", "addresses.CheckRequiredFieldsAddressesSelect()")
	return st.GetStatus(100)
}

// CheckRequiredFieldsAddressInsert - checks required fields to address INSERT
func CheckRequiredFieldsAddressInsert(req *api.RequestAddress) st.ResponseStatus {
	lgr.LOG.Info("-->> ", "addresses.CheckRequiredFieldsAddressInsert()")

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

	//Address
	lgr.LOG.Info("_ACTION_: ", "checking Address")
	if req.Address == "" {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " Address", " must contain a value"))
		return st.GetStatus(401, " Address", " must contain a value")
	}

	//DeliveryId
	lgr.LOG.Info("_ACTION_: ", "checking DeliveryId")
	if req.DeliveryId == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " DeliveryId", " must contain a value"))
		return st.GetStatus(401, " DeliveryId", " must contain a value")
	}

	//AddressTypeId
	lgr.LOG.Info("_ACTION_: ", "checking AddressTypeId")
	if req.AddressTypeId == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " AddressTypeId", " must contain a value"))
		return st.GetStatus(401, " AddressTypeId", " must contain a value")
	}

	//CountryId
	lgr.LOG.Info("_ACTION_: ", "checking CountryId")
	if req.CountryId == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " CountryId", " must contain a value"))
		return st.GetStatus(401, " CountryId", " must contain a value")
	}

	lgr.LOG.Info("<<-- ", "addresses.CheckRequiredFieldsAddressInsert()")
	return st.GetStatus(100)
}

// CheckRequiredFieldsAddressUpdate - checks required fields to address UPDATE
func CheckRequiredFieldsAddressUpdate(req *api.RequestAddress) st.ResponseStatus {
	lgr.LOG.Info("-->> ", "addresses.CheckRequiredFieldsAddressUpdate()")

	//AuthorId
	lgr.LOG.Info("_ACTION_: ", "checking AuthorId")
	if req.AuthorId == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " AuthorId", " must contain a value"))
		return st.GetStatus(401, " AuthorId", " must contain a value")
	}

	//Id
	lgr.LOG.Info("_ACTION_: ", "checking Id")
	if req.Id == 0 {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(401, " Id", " must contain a value"))
		return st.GetStatus(401, " Id", " must contain a value")
	}

	lgr.LOG.Info("<<-- ", "addresses.CheckRequiredFieldsAddressUpdate()")
	return st.GetStatus(100)
}

// CheckRequiredFieldsAddressesDeletionFlagUpdate - checked required fields to UPDATE deletion flags
func CheckRequiredFieldsAddressesDeletionFlagUpdate(req *api.RequestAddressesDeletionFlags) st.ResponseStatus {
	lgr.LOG.Info("-->> ", "addresses.CheckRequiredFieldsAddressesDeletionFlagUpdate()")

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

	lgr.LOG.Info("<<-- ", "addresses.CheckRequiredFieldsAddressesDeletionFlagUpdate()")
	return st.GetStatus(100)
}

//**************************************************************
// Create queries
//**************************************************************

// createQueryAddressesSelect - returns query to SELECT
func createQueryAddressesSelect() string {
	lgr.LOG.Info("-->> ", "addresses.createQueryAddressesSelect()")
	q := fmt.Sprintf("SELECT %s FROM addresses_ref", createViewAddresses())
	lgr.LOG.Info("_QUERY_: ", hlp.DeleteTabsAndNewLinesSymbols(q))
	lgr.LOG.Info("<<-- ", "addresses.createQueryAddressesSelect()")
	return q
}

// CreateQueryAddressInsert - returns query to INSERT
func CreateQueryAddressInsert(req *api.RequestAddress) (string, st.ResponseStatus) {
	lgr.LOG.Info("-->> ", "addresses.CreateQueryAddressInsert()")

	q := fmt.Sprintf("INSERT INTO addresses_ref (address, address_json, isquality_by_dadata, isvalid_by_post, delivery_id, address_type_id, client_id, country_id, comment, created_by, delayed_normalization) "+
		"VALUES ('%s', '%s', %v, %v, %d, %d, %d, %d, '%s', %d, %t) RETURNING ", req.Address, req.AddressJson, req.IsQualityByDadata, req.IsValidByPost, req.DeliveryId, req.AddressTypeId, req.ClientId, req.CountryId, req.Comment, req.AuthorId, req.DelayedNormalization)

	q += createViewAddresses()

	lgr.LOG.Info("_QUERY_: ", q)
	lgr.LOG.Info("<<-- ", "addresses.CreateQueryAddressInsert()")
	return q, st.GetStatus(100)
}

// CreateQueryAddressUpdate - returns query to UPDATE
func CreateQueryAddressUpdate(req *api.RequestAddress) (string, st.ResponseStatus) {
	lgr.LOG.Info("-->> ", "addresses.CreateQueryAddressUpdate()")

	sFields := ""
	sValues := ""
	cnt := 0

	//AuthorId
	if req.AuthorId != 0 {
		sFields += "updated_by, "
		sValues += fmt.Sprintf("%d, ", req.AuthorId)
		cnt++
	}

	//Address
	if req.Address != "" {
		sFields += "address, "
		sValues += fmt.Sprintf("'%s', ", req.Address)
		cnt++
	}

	//DeliveryId
	if req.DeliveryId != 0 {
		sFields += "delivery_id, "
		sValues += fmt.Sprintf("%d, ", req.DeliveryId)
		cnt++
	}

	//AddressTypeId
	if req.AddressTypeId != 0 {
		sFields += "address_type_id, "
		sValues += fmt.Sprintf("%d, ", req.AddressTypeId)
		cnt++
	}

	//ClientId
	if req.ClientId != 0 {
		sFields += "client_id, "
		sValues += fmt.Sprintf("%d, ", req.ClientId)
		cnt++
	}

	//CountryId
	if req.CountryId != 0 {
		sFields += "country_id, "
		sValues += fmt.Sprintf("%d, ", req.CountryId)
		cnt++
	}

	//Comment
	if req.Comment != "" {
		sFields += "comment, "
		sValues += fmt.Sprintf("'%s', ", req.Comment)
		cnt++
	}

	//PostalCode
	if req.PostCode != "" {
		sFields += "postal_code, "
		sValues += fmt.Sprintf("'%s', ", req.PostCode)
		cnt++
	}

	//Isdeleted
	if req.IsDeleted == api.ClientsMS_Bool_IS_TRUE || req.IsDeleted == api.ClientsMS_Bool_IS_FALSE {
		sFields += "isdeleted, "
		if req.IsDeleted == api.ClientsMS_Bool_IS_TRUE {
			sValues += fmt.Sprintf("%t, ", true)
		} else {
			sValues += fmt.Sprintf("%t, ", false)
		}
		cnt++
	}

	//IsLast
	if req.IsLast == api.ClientsMS_Bool_IS_TRUE || req.IsLast == api.ClientsMS_Bool_IS_FALSE {
		sFields += "islast, "
		if req.IsLast == api.ClientsMS_Bool_IS_TRUE {
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
		q = "UPDATE addresses_ref SET (" + sFields + ") = (" + sValues + ")"
	} else {
		q = "UPDATE addresses_ref SET " + sFields + " = " + sValues
	}

	q += fmt.Sprintf(" WHERE id = %d RETURNING ", req.Id)
	q += createViewAddresses()

	lgr.LOG.Info("_QUERY_: ", hlp.DeleteTabsAndNewLinesSymbols(q))

	lgr.LOG.Info("<<-- ", "addresses.CreateQueryAddressUpdate()")
	return q, st.GetStatus(100)
}

// CreateQueryAddressesDeletionFlagsUpdate - returns query to UPDATE notes deletion flags
func CreateQueryAddressesDeletionFlagsUpdate(req *api.RequestAddressesDeletionFlags) string {
	lgr.LOG.Info("-->> ", "addresses.CreateQueryAddressesDeletionFlagsUpdate()")
	q := fmt.Sprintf("UPDATE addresses_ref SET isdeleted = %t, updated_by = %d WHERE id IN %s RETURNING ",
		hlp.GetBool(req.IsDeleted), req.AuthorId, hlp.Uint64ArrayToString(req.Ids))
	q += createViewAddresses()
	lgr.LOG.Info("_QUERY_: ", hlp.DeleteTabsAndNewLinesSymbols(q))
	lgr.LOG.Info("<<-- ", "addresses.CreateQueryAddressesDeletionFlagsUpdate()")
	return q
}

// **************************************************************
// Create view
// **************************************************************
func createViewAddresses() string {
	return `id AS address_id,
			client_id AS address_client_id,
    		delivery_id AS address_delivery_id,
			( SELECT d.name FROM deliveries_ref d WHERE d.id = delivery_id ) AS address_delivery_name,
			address AS address_address,
			address_type_id AS address_address_type_id,
			( SELECT at.name FROM addresses_types_ref at WHERE at.id = address_type_id ) AS address_address_type_name,
			concat(postal_code, '') AS address_postal_code,
			concat(timezone, '') AS address_timezone,
			country_id AS address_country_id,
			( SELECT c.name FROM countries_ref c WHERE c.id = country_id ) AS address_country_name,
			isquality_by_dadata AS address_isquality_by_dadata,
			isvalid_by_post AS address_isvalid_by_post,
			islast AS address_islast,
    		to_char(created_at, 'YYYY-MM-DD HH24:MI:SS'::text) AS address_created_at,
    		created_by AS address_created_by,
            ( SELECT u.fullname FROM users_ref u WHERE u.id = created_by ) AS address_created_by_name,
    		to_char(updated_at, 'YYYY-MM-DD HH24:MI:SS'::text) AS address_updated_at,
    		updated_by AS address_updated_by,
    		( SELECT u.fullname FROM users_ref u WHERE u.id = updated_by ) AS address_updated_by_name,
    		isdeleted AS address_isdeleted,
			concat(address_json, '') AS address_address_json,
			concat(comment, '') AS address_comment,
			delayed_normalization AS address_delayed_normalization`
}

//**************************************************************
// Add WHERE to queries
//**************************************************************

// addWhereToQueryAddressesSelect - added WHERE to query SELECT
func addWhereToQueryAddressesSelect(q string, req *api.RequestAddress) string {
	lgr.LOG.Info("-->> ", "addresses.addWhereToQueryAddressesSelect()")
	cnt := 0

	// ClientID
	if req.ClientId != 0 {
		q += fmt.Sprintf(" WHERE client_id = %d", req.ClientId)
		cnt++
	}

	// Address
	if req.Address != "" {
		if cnt > 0 {
			q += " AND address ILIKE '%" + req.Address + "%'"
		} else {
			q += " WHERE address ILIKE '%" + req.Address + "%'"
		}
		cnt++
	}

	// Country
	if req.CountryId != 0 {
		if cnt > 0 {
			q += " AND country_id = " + fmt.Sprintf("%d", req.CountryId)
		} else {
			q += " WHERE country_id = " + fmt.Sprintf("%d", req.CountryId)
		}
		cnt++
	}

	// AddressType
	if req.AddressTypeId != 0 {
		if cnt > 0 {
			q += " AND address_type_id = " + fmt.Sprintf("%d", req.AddressTypeId)
		} else {
			q += " WHERE address_type_id = " + fmt.Sprintf("%d", req.AddressTypeId)
		}
		cnt++
	}

	// Delivery
	if req.DeliveryId != 0 {
		if cnt > 0 {
			q += " AND delivery_id = " + fmt.Sprintf("%d", req.DeliveryId)
		} else {
			q += " WHERE delivery_id = " + fmt.Sprintf("%d", req.DeliveryId)
		}
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

	// IsqualityByDadata
	switch req.IsQualityByDadata {
	case api.ClientsMS_Bool_IS_TRUE:
		if cnt > 0 {
			q += " AND isquality_by_dadata = true"
		} else {
			q += " WHERE isquality_by_dadata = true"
		}
		cnt++
	case api.ClientsMS_Bool_IS_FALSE:
		if cnt > 0 {
			q += " AND isquality_by_dadata = false"
		} else {
			q += " WHERE isquality_by_dadata = false"
		}
		cnt++
	}

	// IsvalidByPost
	switch req.IsValidByPost {
	case api.ClientsMS_Bool_IS_TRUE:
		if cnt > 0 {
			q += " AND isvalid_by_post = true"
		} else {
			q += " WHERE isvalid_by_post = true"
		}
		cnt++
	case api.ClientsMS_Bool_IS_FALSE:
		if cnt > 0 {
			q += " AND isvalid_by_post = false"
		} else {
			q += " WHERE isvalid_by_post = false"
		}
		cnt++
	}

	// Islast
	switch req.IsLast {
	case api.ClientsMS_Bool_IS_TRUE:
		if cnt > 0 {
			q += " AND islast = true"
		} else {
			q += " WHERE islast = true"
		}
		cnt++
	}

	lgr.LOG.Info("_QUERY_: ", hlp.DeleteTabsAndNewLinesSymbols(q))
	lgr.LOG.Info("<<-- ", "addresses.addWhereToQueryAddressesSelect()")
	return q
}

//**************************************************************
// Add ORDER to queries
//**************************************************************

// addOrderToQueryAddressesSelect - added ORDER to query SELECT
func addOrderToQueryAddressesSelect(q string, req *api.RequestAddress) string {

	lgr.LOG.Info("-->> ", "addresses.addOrderToQueryAddressesSelect()")

	strOrder := ""
	switch req.Order {
	case api.ClientsMS_AddressesOrder_BY_ADDRESS_ID:
		strOrder = " ORDER BY id"
	case api.ClientsMS_AddressesOrder_BY_ADDRESS_ADDRESS:
		strOrder = " ORDER BY address"
	case api.ClientsMS_AddressesOrder_BY_ADDRESS_CREATED_AT:
		strOrder = " ORDER BY created_at"
	case api.ClientsMS_AddressesOrder_BY_ADDRESS_UPDATED_AT:
		strOrder = " ORDER BY updated_at"
	default:
		strOrder = " ORDER BY created_at"
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

	lgr.LOG.Info("<<-- ", "addresses.addOrderToQueryAddressesSelect()")
	return q
}

//**************************************************************
// Generating a queries
//**************************************************************

// GenQueryAddressesSelect - generates a query for clients addresses select
func GenQueryAddressesSelect(req *api.RequestAddress) string {

	lgr.LOG.Info("-->> ", "addresses.GenQueryAddressesSelect()")
	var q string

	lgr.LOG.Info("_ACTION_: ", "generating query")
	q = createQueryAddressesSelect()

	lgr.LOG.Info("_ACTION_: ", "adding WHERE to query")
	q = addWhereToQueryAddressesSelect(q, req)

	lgr.LOG.Info("_ACTION_: ", "adding ORDER to query")
	q = addOrderToQueryAddressesSelect(q, req)

	lgr.LOG.Info("_QUERY_: ", hlp.DeleteTabsAndNewLinesSymbols(q))
	lgr.LOG.Info("<<-- ", "addresses.GenQueryAddressesSelect()")
	return q
}
