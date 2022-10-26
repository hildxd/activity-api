package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type SolidActivityForm struct {
	Date       string
	WechatName string
}

type SolidActivityRepo interface {
	List(ctx context.Context) ([]*SolidActivity, error)
	Create(ctx context.Context, solidActivity *SolidActivity) (*SolidActivity, error)
}

type SolidActivityUsecase struct {
	repo SolidActivityRepo
	log  *log.Helper
}

type SolidActivity struct {
	ID         uint
	WechatName string
	Date       string
}

func NewSolidActivityUsecase(repo SolidActivityRepo, logger log.Logger) *SolidActivityUsecase {
	return &SolidActivityUsecase{repo: repo, log: log.NewHelper(logger)}
}

// usecase implements the SolidActivityUsecase interface.
func (sc *SolidActivityUsecase) Create(ctx context.Context, solidActivity *SolidActivity) (*SolidActivity, error) {
	return sc.repo.Create(ctx, solidActivity)
}

func (sc *SolidActivityUsecase) List(ctx context.Context) ([]*SolidActivity, error) {
	return sc.repo.List(ctx)
}
