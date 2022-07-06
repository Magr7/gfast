package service

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/oncall"
	"github.com/tiger1103/gfast/v3/internal/app/oncall/service/internal/dao"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

type IFreezingPoints interface {
	List(ctx context.Context, req *oncall.FreezingPointsSearchReq) (res *oncall.FreezingPointsSearchRes, err error)
}

type freezingPointsImpl struct {
}

var freezingPointsService = freezingPointsImpl{}

func FreezingPoints() IFreezingPoints {
	return &freezingPointsService
}

func (s *freezingPointsImpl) List(ctx context.Context, req *oncall.FreezingPointsSearchReq) (res *oncall.FreezingPointsSearchRes, err error) {
	res = new(oncall.FreezingPointsSearchRes)
	g.Try(func() {
		m := dao.FreezingPointsRecord.Ctx(ctx)
		if req != nil {
			if req.OrderBy != "" {
				m = m.Order(req.OrderBy)
			}
		}

		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取冻结积分失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		res.CurrentPage = req.PageNum
		err = m.Page(req.PageNum, req.PageSize).Scan(&res.FreezingPointsList)
		liberr.ErrIsNil(ctx, err, "获取冻结积分列表失败")
	})
	return res, nil
}
