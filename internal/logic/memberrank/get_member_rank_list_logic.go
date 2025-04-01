package memberrank

import (
	"context"
	"github.com/suyuan32/simple-admin-member-rpc/ent/memberrank"
	"github.com/suyuan32/simple-admin-member-rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-member-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-member-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberRankListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberRankListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberRankListLogic {
	return &GetMemberRankListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMemberRankListLogic) GetMemberRankList(in *mms.MemberRankListReq) (*mms.MemberRankListResp, error) {
	var predicates []predicate.MemberRank
	//if in.CreatedAt != nil {
	//	predicates = append(predicates, memberrank.CreatedAtGTE(time.UnixMilli(*in.CreatedAt)))
	//}
	//if in.UpdatedAt != nil {
	//	predicates = append(predicates, memberrank.UpdatedAtGTE(time.UnixMilli(*in.UpdatedAt)))
	//}
	if in.Name != nil {
		predicates = append(predicates, memberrank.NameContains(*in.Name))
	}
	result, err := l.svcCtx.DB.MemberRank.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &mms.MemberRankListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &mms.MemberRankInfo{
			Id:          &v.ID,
			CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:   pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Name:        &v.Name,
			Code:        &v.Code,
			Description: &v.Description,
			Remark:      &v.Remark,
		})
	}

	return resp, nil
}
