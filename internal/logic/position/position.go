package position

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/frame/g"
	"myGoFrame/internal/dao"
	"myGoFrame/internal/model"
	"myGoFrame/internal/model/entity"
	"myGoFrame/internal/service"
)

type sPosition struct{}

func init() {
	service.RegisterPosition(New())
}

func New() *sPosition {
	return &sPosition{}
}

// Create 创建内容
func (s *sPosition) Create(ctx context.Context, in model.PositionCreateInput) (out model.PositionCreateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.PositionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.PositionCreateOutput{PositionId: uint(lastInsertID)}, err
}

func (s *sPosition) Delete(ctx context.Context, id uint) error {
	return dao.PositionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除轮播图
		_, err := dao.PositionInfo.Ctx(ctx).Where(g.Map{
			dao.PositionInfo.Columns().Id: id,
		}).Delete()
		return err
	})
}

// Update 修改
func (s *sPosition) Update(ctx context.Context, in model.PositionUpdateInput) error {
	return dao.PositionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		_, err := dao.PositionInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.PositionInfo.Columns().Id).
			Where(dao.PositionInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 查询内容列表
func (s *sPosition) GetList(ctx context.Context, in model.PositionGetListInput) (out *model.PositionGetListOutput, err error) {
	var (
		m = dao.PositionInfo.Ctx(ctx)
	)
	out = &model.PositionGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 分页查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式
	listModel = listModel.OrderDesc(dao.PositionInfo.Columns().Sort)

	// 执行查询
	var list []*entity.PositionInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	// Position
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
