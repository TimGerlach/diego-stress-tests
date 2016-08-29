package cli

import (
	"encoding/json"
	"errors"
	"time"

	"golang.org/x/net/context"

	"code.cloudfoundry.org/lager"
)

var (
	ErrNoDomains = errors.New("no shared domains returned by CC")
)

type SharedDomainResponse struct {
	Resources []struct {
		Entity struct {
			Name string `json:"name"`
		} `json:"entity"`
	} `json:"resources"`
}

func GetDefaultSharedDomain(ctx context.Context, cfClient CFClient) (string, error) {
	logger := ctx.Value("logger").(lager.Logger)
	logger = logger.Session("get-default-shared-domain")

	// cf curl to get shared domains
	out, err := cfClient.Cf(ctx, 30*time.Second, "curl", "/v2/shared_domains")
	if err != nil {
		return "", err
	}

	// parse response
	var f SharedDomainResponse
	err = json.Unmarshal(out, &f)
	if err != nil {
		return "", err
	}
	if len(f.Resources) > 0 {
		return f.Resources[0].Entity.Name, nil
	}

	return "", ErrNoDomains
}
