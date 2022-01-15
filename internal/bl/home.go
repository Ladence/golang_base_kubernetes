package bl

import (
	"github.com/Ladence/golang_base_kubernetes/internal/api"
	"github.com/Ladence/golang_base_kubernetes/internal/version"
)

func Home() api.HomeResponse {
	return api.HomeResponse{
		BuildTime: version.BuildTime,
		Revision:  version.Commit,
		Release:   version.Release,
	}
}
