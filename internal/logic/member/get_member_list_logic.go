package member

import (
	"context"
	"github.com/suyuan32/simple-admin-member-rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-member-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-member-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberListLogic {
	return &GetMemberListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMemberListLogic) GetMemberList(in *mms.MemberListReq) (*mms.MemberListResp, error) {
	var predicates []predicate.Member
	//if in.CreatedAt != nil {
	//	predicates = append(predicates, member.CreatedAtGTE(time.UnixMilli(*in.CreatedAt)))
	//}
	//if in.UpdatedAt != nil {
	//	predicates = append(predicates, member.UpdatedAtGTE(time.UnixMilli(*in.UpdatedAt)))
	//}
	//if in.Status != nil {
	//	predicates = append(predicates, member.StatusEQ(uint8(*in.Status)))
	//}
	result, err := l.svcCtx.DB.Member.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &mms.MemberListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &mms.MemberInfo{
			Id:        pointy.GetPointer(v.ID.String()),
			CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Status:    pointy.GetPointer(uint32(v.Status)),
			Username:  &v.Username,
			Password:  &v.Password,
			Nickname:  &v.Nickname,
			RankId:    &v.RankID,
			Mobile:    &v.Mobile,
			Email:     &v.Email,
			Avatar:    &v.Avatar,
			WechatId:  &v.WechatOpenID,
			ExpiredAt: pointy.GetUnixMilliPointer(v.ExpiredAt.UnixMilli()),
		})
	}

	return resp, nil
}
