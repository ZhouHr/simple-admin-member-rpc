package token

import (
	"context"
	"github.com/suyuan32/simple-admin-member-rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-member-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-member-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTokenListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenListLogic {
	return &GetTokenListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTokenListLogic) GetTokenList(in *mms.TokenListReq) (*mms.TokenListResp, error) {
	var predicates []predicate.Token
	//if in.CreatedAt != nil {
	//	predicates = append(predicates, token.CreatedAtGTE(time.UnixMilli(*in.CreatedAt)))
	//}
	//if in.UpdatedAt != nil {
	//	predicates = append(predicates, token.UpdatedAtGTE(time.UnixMilli(*in.UpdatedAt)))
	//}
	//if in.Status != nil {
	//	predicates = append(predicates, token.StatusEQ(uint8(*in.Status)))
	//}
	result, err := l.svcCtx.DB.Token.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &mms.TokenListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &mms.TokenInfo{
			Id:        pointy.GetPointer(v.ID.String()),
			CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Status:    pointy.GetPointer(uint32(v.Status)),
			Uuid:      pointy.GetPointer(v.UUID.String()),
			Token:     &v.Token,
			Username:  &v.Username,
			Source:    &v.Source,
			ExpiredAt: pointy.GetPointer(v.ExpiredAt.UnixMilli()),
		})
	}

	return resp, nil
}
