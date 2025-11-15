package service

import "context"

type HealthServiceStruct struct{}

func NewHealthService() *HealthServiceStruct {
	return &HealthServiceStruct{}
}

func (h *HealthServiceStruct) Ping(ctx context.Context) error {
	return nil
}
