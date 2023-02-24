package helpers

import cfg "clients_ms/pkg/config"

// AuthTokenIsValid - check token
func AuthTokenIsValid(authToken string) bool {
	return authToken == cfg.CFG.Server.CliToken
}