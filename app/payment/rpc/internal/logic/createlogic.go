package logic

import (
	"context"

	"github.com/mirage208/gomall/app/order/rpc/pb/order"
	"github.com/mirage208/gomall/app/payment/model"
	"github.com/mirage208/gomall/app/payment/rpc/internal/svc"
	"github.com/mirage208/gomall/app/payment/rpc/pb/payment"
	"github.com/mirage208/gomall/app/user/rpc/pb/user"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *payment.CreateRequest) (*payment.CreateResponse, error) {
	// check user existence
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.Uid,
	})
	if err != nil {
		return nil, err
	}

	// check order existence
	_, err = l.svcCtx.OrderRpc.Detail(l.ctx, &order.DetailRequest{
		Id: in.Oid,
	})
	if err != nil {
		return nil, err
	}

	// check if payment already exists
	_, err = l.svcCtx.PaymentModel.FindOneByOid(l.ctx, in.Oid)
	if err == nil {
		return nil, status.Error(100, "订单已创建支付")
	}

	newPay := model.Pay{
		Uid:    in.Uid,
		Oid:    in.Oid,
		Amount: in.Amount,
		Source: 0,
		Status: 0,
	}

	res, err := l.svcCtx.PaymentModel.Insert(l.ctx, &newPay)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newPay.Id, err = res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &payment.CreateResponse{
		Id: newPay.Id,
	}, nil
}
