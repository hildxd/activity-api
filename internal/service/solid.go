package service

import (
	v1 "activity-api/api/solid/v1"
	"activity-api/internal/biz"
	"context"
)

func (s *SolidService) ListSolid(ctx context.Context, req *v1.ListSolidRequest) (reply *v1.ListSolidReply, err error) {
	rv, err := s.sa.List(ctx)
	if err != nil {
		return nil, err
	}
	solids := make([]*v1.Solid, 0)
	for _, x := range rv {
		solids = append(solids, &v1.Solid{
			Id:         int64(x.ID),
			WechatName: x.WechatName,
			Date:       x.Date})
	}
	reply = &v1.ListSolidReply{
		List: solids,
	}
	return
}

func (s *SolidService) CreateSolid(ctx context.Context, req *v1.CreateSolidRequest) (reply *v1.CreateSolidReply, err error) {
	solid, err := s.sa.Create(ctx, &biz.SolidActivity{
		WechatName: req.WechatName,
		Date:       req.Date,
	})
	if err != nil {
		return nil, err
	}
	reply = &v1.CreateSolidReply{
		Solid: &v1.Solid{
			Id:         int64(solid.ID),
			WechatName: solid.WechatName,
			Date:       solid.Date,
		},
	}
	return
}
