package service_test

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/oncall"
	"github.com/tiger1103/gfast/v3/internal/app/oncall/service"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func TestList(t *testing.T) {
	var (
		ctx context.Context
		req oncall.FreezingPointsSearchReq
	)

	req.PageNum = 1
	req.PageSize = 2
	res, err := service.FreezingPoints().List(ctx, &req)
	liberr.ErrIsNil(ctx, err, "错误")
	g.Dump(res)

}
