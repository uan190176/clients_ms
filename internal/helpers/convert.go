package helpers

import (
	"clients_ms/internal/api"
	"fmt"
)

func GetBool(num api.ClientsMS_Bool) bool {
	if num == api.ClientsMS_Bool_IS_TRUE {
		return true
	}
	return false
}

// Uint64ArrayToString - convert uint64 array to string
func Uint64ArrayToString(a []uint64) string {
	str := ""

	for i, v := range a {
		if i > 0 {
			str += ", "
		}
		str += fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("(%s)", str)
}
