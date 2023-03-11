package storage

import (
	cnf "clients_ms/pkg/config"
	"context"
	fs "github.com/uan190176/fstorage"
)

// PingDatabase - check connection to database
func PingDatabase(ctx context.Context) bool {

	// Create config to connect
	sp := &fs.StorageParams{
		Ctx:     ctx,
		Dsn:     cnf.CFG.Database.URL,
		AppName: "ClientsMS::Ping()",
	}

	// Open connection to database
	c, err := fs.NewStorage(sp)
	if err != nil {
		return false
	}

	err = c.Ping(ctx)
	if err != nil {
		e := c.Release()
		if e != nil {
			return false
		}
		return false
	}

	e := c.Release()
	if e != nil {
		return false
	}
	return true
}
