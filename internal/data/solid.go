package data

import (
	"activity-api/internal/biz"
	"context"
	"errors"

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

func findSolidWithDate(ctx context.Context, db *gorm.DB, date string) (*SolidActivity, error) {
	var solid SolidActivity
	result := db.Where("date = ?", date).First(&solid)
	if result.Error != nil {
		return nil, result.Error
	}
	return &solid, nil
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
	if solidActivity == nil {
		return nil, errors.New("request body is nil")
	}
	item, _ := findSolidWithDate(ctx, s.data.db, solidActivity.Date)
	if item == nil {

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
	return nil, errors.New("date already exists")
}
