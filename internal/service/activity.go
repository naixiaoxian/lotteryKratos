package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"lotteryKratos/internal/biz"
	"time"

	pb "lotteryKratos/api/activity/v1"
)

type ActivityService struct {
	pb.UnimplementedActivityServer
	log      *log.Helper
	activity *biz.ActivityDomain
}

func NewActivityService(activity *biz.ActivityDomain, logger log.Logger) *ActivityService {
	return &ActivityService{
		log:      log.NewHelper(logger),
		activity: activity,
	}
}

func (s *ActivityService) CreateActivity(ctx context.Context, req *pb.CreateActivityRequest) (*pb.CreateActivityReply, error) {
	s.log.Infof("input data %v", req)
	beginTime, _ := time.ParseInLocation("2006-01-02 15:04:05", req.BeginDateTime, time.Local)
	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", req.EndDateTime, time.Local)
	ret, err := s.activity.Create(ctx, &biz.Activity{
		Name:          req.ActivityName,
		Desc:          req.ActivityDesc,
		BeginDateTime: beginTime,
		EndDateTime:   endTime,
		StockCount:    0,
		TakeCount:     0,
		State:         0,
		Creator:       req.Creator,
	})
	return &pb.CreateActivityReply{
		Activity: &pb.Activity{
			ActivityName:  ret.Name,
			ActivityDesc:  ret.Desc,
			BeginDateTime: ret.BeginDateTime.String(),
			EndDateTime:   ret.EndDateTime.String(),
			StockCount:    ret.StockCount,
			TakeCount:     ret.TakeCount,
			State:         ret.State,
			Creator:       ret.Creator,
			CreateTime:    ret.CreateTime.String(),
			UpdateTime:    ret.UpdateTime.String(),
		},
	}, err
}
func (s *ActivityService) UpdateActivity(ctx context.Context, req *pb.UpdateActivityRequest) (*pb.UpdateActivityReply, error) {
	return &pb.UpdateActivityReply{}, nil
}
func (s *ActivityService) DeleteActivity(ctx context.Context, req *pb.DeleteActivityRequest) (*pb.DeleteActivityReply, error) {
	return &pb.DeleteActivityReply{}, nil
}
func (s *ActivityService) GetActivity(ctx context.Context, req *pb.GetActivityRequest) (*pb.GetActivityReply, error) {
	s.log.Infof("input data %v", req)
	ret, err := s.activity.Get(ctx, req.Id)
	return &pb.GetActivityReply{
		Activity: &pb.Activity{
			ActivityName:  ret.Name,
			ActivityDesc:  ret.Desc,
			BeginDateTime: ret.BeginDateTime.String(),
			EndDateTime:   ret.EndDateTime.String(),
			StockCount:    ret.StockCount,
			TakeCount:     ret.TakeCount,
			State:         ret.State,
			Creator:       ret.Creator,
			CreateTime:    ret.CreateTime.String(),
			UpdateTime:    ret.UpdateTime.String(),
		},
	}, err
}
func (s *ActivityService) ListActivity(ctx context.Context, req *pb.ListActivityRequest) (*pb.ListActivityReply, error) {
	return &pb.ListActivityReply{}, nil
}
