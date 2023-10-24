package service

import (
	"context"
	"math/rand"

	pb "KratosDemo_01/api/verifyCode"
)

type VerifyCodeService struct {
	pb.UnimplementedVerifyCodeServer
}

func NewVerifyCodeService() *VerifyCodeService {
	return &VerifyCodeService{}
}

func (s *VerifyCodeService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeRequest) (*pb.GetVerifyCodeReply, error) {
	return &pb.GetVerifyCodeReply{
		Code: RandCode(int(req.Length), req.Type),
	}, nil
}

// 开放的被调用的方法，用于区分类型
func RandCode(l int, t pb.TYPE) string {
	switch t {
	case pb.TYPE_DEFAULT:
		fallthrough
	case pb.TYPE_DIGIT:
		return randCode("0123456789", l)
	case pb.TYPE_LETTER:
		return randCode("abcdefghijklmnopqrstuvwxyz", l)
	case pb.TYPE_MIXED:
		return randCode("0123456789abcdefghijklmnopqrstuvwxyz", l)
	default:

	}
	return ""
}

// 随机的核心方法(最简单的实现)
func randCode(chars string, l int) string {
	charsLen := len(chars)

	result := make([]byte, l)
	//根据目标长度进行循环
	for i := 0; i < l; i++ {
		//随机生成0-n的整形随机数，随机的下标
		randIndex := rand.Intn(charsLen)
		result[i] = chars[randIndex]
	}
	return string(result)
}
