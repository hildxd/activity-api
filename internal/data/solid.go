package data

import (
	"activity-api/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type SolidActivity struct {
	gorm.Model
	Date       string `gorm:"size:200"`
	WechatName string `gorm:"size:200"`
}

type solidActivityRepo struct {
	data *Data
	log  *log.Helper
}

func NewSolidRepo(data *Data, logger log.Logger) biz.SolidActivityRepo {
	return &solidActivityRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s *solidActivityRepo) List(ctx context.Context) (rv []*biz.SolidActivity, err error) {
	var solidActivitys []SolidActivity
	result := s.data.db.Find(&solidActivitys)
	if result.Error != nil {
		return nil, result.Error
	}
	rv = make([]*biz.SolidActivity, len(solidActivitys))
	for i, x := range solidActivitys {
		rv[i] = &biz.SolidActivity{
			ID:         x.ID,
			Date:       x.Date,
			WechatName: x.WechatName,
		}
	}
	return rv, result.Error
}

func (s *solidActivityRepo) Create(ctx context.Context, solidActivity *biz.SolidActivity) (*biz.SolidActivity, error) {
	po := SolidActivity{
		Date:       solidActivity.Date,
		WechatName: solidActivity.WechatName,
	}
	result := s.data.db.Create(&po)
	if result.Error != nil {
		return nil, result.Error
	}

	return &biz.SolidActivity{
		ID:         po.ID,
		Date:       po.Date,
		WechatName: po.WechatName,
	}, nil
}
