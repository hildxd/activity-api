package service

import (
	v1 "activity-api/api/solid/v1"
	"activity-api/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewSolidService)

type SolidService struct {
	v1.UnimplementedSolidServiceServer

	sa  *biz.SolidActivityUsecase
	log *log.Helper
}

func NewSolidService(sa *biz.SolidActivityUsecase, logger log.Logger) *SolidService {
	return &SolidService{sa: sa, log: log.NewHelper(logger)}
}
