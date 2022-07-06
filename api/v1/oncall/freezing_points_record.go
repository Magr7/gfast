package oncall

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/oncall/model/entity"
)

type FreezingPointsSearchReq struct {
	g.Meta `path:"/freezing_points/list" tags:"冻结积分管理" method:"get" summary:"冻结积分列表"`
	commonApi.PageReq
}

type FreezingPointsSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	FreezingPointsList []*entity.FreezingPointsRecord `json:"freezingPointsList"`
}
