package service

import "context"

type Math interface {
	Sum(ctx context.Context, x, y int32) int32
	Sub(ctx context.Context, x, y int32) int32
}

type Service struct {
	Math
}

func NewService() *Service {
	return &Service{
		Math: NewMathService(),
	}
}

type MathService struct{}

func NewMathService() *MathService {
	return &MathService{}
}

func (s *MathService) Sum(ctx context.Context, x, y int32) int32 {
	return x + y
}

func (s *MathService) Sub(ctx context.Context, x, y int32) int32 {
	return x - y
}
