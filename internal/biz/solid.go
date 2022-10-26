package biz

import "github.com/go-kratos/kratos/v2/log"

type SolidActivityForm struct {
	Date       string
	WechatName string
}

type SolidActivityRepo interface {
}

type SolidActivityUsecase struct {
	repo SolidActivityRepo
	log  *log.Helper
}

func NewSolidActivityUsecase(repo SolidActivityRepo, logger log.Logger) *SolidActivityUsecase {
	return &SolidActivityUsecase{repo: repo, log: log.NewHelper(logger)}
}
