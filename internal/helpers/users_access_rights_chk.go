package helpers

import (
	"clients_ms/internal/api"
	cfg "clients_ms/pkg/config"
	lgr "clients_ms/pkg/logger"
	"context"
	st "github.com/uan190176/statuses"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GetUserAccessRightForTable - check and return uar to table
func GetUserAccessRightForTable(ctx context.Context, req *api.RequestUsersAccessRightsForTable) (bool, st.ResponseStatus) {

	lgr.LOG.Info("-->> ", "helpers.GetUserAccessRightForTable()")

	//Open connection to microservice
	lgr.LOG.Info("_ACTION_: ", "opening grpc connection to microservice users_ms")
	conn, err := grpc.Dial(cfg.CFG.MicroServices["users_ms"].BindAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		lgr.LOG.Error("_ERR_: ", err)
		return false, st.GetStatus(0, err.Error())
	}
	defer func(conn *grpc.ClientConn) error {
		err = conn.Close()
		if err != nil {
			lgr.LOG.Error("_ERR_: ", err)
			return err
		}
		return nil
	}(conn)

	c := api.NewUsersServicesClient(conn)

	lgr.LOG.Info("_ACTION_: ", "calling method from microservice GetUserAccessRightForTable()")
	res, err := c.GetUserAccessRightForTable(ctx, req)
	if err != nil {
		lgr.LOG.Error("_ERR_: ", err)
		return false, st.GetStatus(0, err.Error())
	}
	if res == nil {
		lgr.LOG.Warn("_WARN_: ", st.GetStatus(205))
		return false, st.GetStatus(205)
	}

	//Check result
	if res.Status.Code > 100 {
		lgr.LOG.Warn("_WARN_: ", res.Status)
		return false, st.ResponseStatus{
			Code:        res.Status.Code,
			Description: res.Status.Description,
		}
	}

	lgr.LOG.Info("_RESULT_: ", res.UsersAccessRightsForTable)
	lgr.LOG.Info("<<-- ", "helpers.GetUserAccessRightForTable()")
	return res.UsersAccessRightsForTable.Result, st.GetStatus(100)
}
