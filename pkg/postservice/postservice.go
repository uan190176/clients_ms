package postservice

import (
	cfg "clients_ms/pkg/config"
	lgr "clients_ms/pkg/logger"
	"errors"
	ps "github.com/uan190176/postservice"
)

type NormalizedAddress struct {
	Address        string `json:"address"`
	AddressJson    string `json:"address_json"`
	QualityCode    string `json:"quality_code"`
	ValidationCode string `json:"validation_code"`
}

// GetNormalizedAddress - get normalized address and quality/validation code
func GetNormalizedAddress(address string) (*NormalizedAddress, error) {
	lgr.LOG.Info("-->> ", "postservice.GetNormalizedAddress()")

	if address == "" {
		return nil, errors.New("address is empty")
	}

	addr := ps.Address{
		Id:              "1",
		OriginalAddress: address,
	}
	addresses := []ps.Address{
		addr,
	}

	cred := ps.PostCredentials{
		URL:   cfg.CFG.PostAPI.URL,
		Token: cfg.CFG.PostAPI.PostToken,
		Auth:  cfg.CFG.PostAPI.PostAuth,
	}

	p := ps.NewPostService(cred)

	res, err := p.GetNormalizedAddresses(addresses)
	if err != nil {
		lgr.LOG.Error("_ERR_: ", err)
		return nil, err
	}

	if len(res) == 0 {
		lgr.LOG.Error("_ERR_: ", errors.New("no normalized address found"))
		return nil, errors.New("no normalized address found")
	}

	normAddr := res[0]
	lgr.LOG.Info("<<-- ", "postservice.GetNormalizedAddress()")
	return &NormalizedAddress{
		Address:        normAddr.Address.OriginalAddress,
		AddressJson:    normAddr.FullJson,
		QualityCode:    normAddr.QualityCode.Code,
		ValidationCode: normAddr.ValidationCode.Code,
	}, nil
}
