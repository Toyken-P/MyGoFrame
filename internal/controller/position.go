package controller

import (
	"context"
	"myGoFrame/api/backend"

	"myGoFrame/internal/model"
	"myGoFrame/internal/service"
)

// Position 手工位管理
var Position = cPosition{}

type cPosition struct{}

func (a *cPosition) Create(ctx context.Context, req *backend.PositionReq) (res *backend.PositionRes, err error) {
	out, err := service.Position().Create(ctx, model.PositionCreateInput{
		PositionCreateUpdateBase: model.PositionCreateUpdateBase{
			PicUrl:    req.PicUrl,
			Link:      req.Link,
			GoodsName: req.GoodsName,
			GoodsId:   req.GoodsId,
			Sort:      req.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.PositionRes{PositionId: out.PositionId}, nil
}

func (a *cPosition) Delete(ctx context.Context, req *backend.PositionDeleteReq) (res *backend.PositionDeleteRes, err error) {
	err = service.Position().Delete(ctx, req.Id)
	return
}

func (a *cPosition) Update(ctx context.Context, req *backend.PositionUpdateReq) (res *backend.PositionUpdateRes, err error) {
	err = service.Position().Update(ctx, model.PositionUpdateInput{
		Id: req.Id,
		PositionCreateUpdateBase: model.PositionCreateUpdateBase{
			PicUrl:    req.PicUrl,
			Link:      req.Link,
			GoodsName: req.GoodsName,
			GoodsId:   req.GoodsId,
			Sort:      req.Sort,
		},
	})
	return
}

func (a *cPosition) List(ctx context.Context, req *backend.PositionGetListCommonReq) (res *backend.PositionGetListCommonRes, err error) {
	getListRes, err := service.Position().GetList(ctx, model.PositionGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &backend.PositionGetListCommonRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}
