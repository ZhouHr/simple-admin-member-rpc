package oauthprovider

import (
	"context"
	"github.com/suyuan32/simple-admin-member-rpc/ent/oauthprovider"
	"github.com/suyuan32/simple-admin-member-rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-member-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-member-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-member-rpc/types/mms"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetOauthProviderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOauthProviderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOauthProviderListLogic {
	return &GetOauthProviderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOauthProviderListLogic) GetOauthProviderList(in *mms.OauthProviderListReq) (*mms.OauthProviderListResp, error) {
	var predicates []predicate.OauthProvider
	//if in.CreatedAt != nil {
	//	predicates = append(predicates, oauthprovider.CreatedAtGTE(time.UnixMilli(*in.CreatedAt)))
	//}
	//if in.UpdatedAt != nil {
	//	predicates = append(predicates, oauthprovider.UpdatedAtGTE(time.UnixMilli(*in.UpdatedAt)))
	//}
	if in.Name != nil {
		predicates = append(predicates, oauthprovider.NameContains(*in.Name))
	}
	result, err := l.svcCtx.DB.OauthProvider.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &mms.OauthProviderListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &mms.OauthProviderInfo{
			Id:           &v.ID,
			CreatedAt:    pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:    pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Name:         &v.Name,
			ClientId:     &v.ClientID,
			ClientSecret: &v.ClientSecret,
			RedirectUrl:  &v.RedirectURL,
			Scopes:       &v.Scopes,
			AuthUrl:      &v.AuthURL,
			TokenUrl:     &v.TokenURL,
			AuthStyle:    &v.AuthStyle,
			InfoUrl:      &v.InfoURL,
		})
	}

	return resp, nil
}
