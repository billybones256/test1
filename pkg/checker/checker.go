package checker

import (
	"context"
	"task1/pkg/api"
)

//GRPCServer ...
type GRPCServer struct{}

//Square ...
func (server *GRPCServer) Check(ctx context.Context, req *api.CheckRequest) (*api.CheckResponse, error) {
	ret := int32(Checker(req.GetMessage()))
	return &api.CheckResponse{Answer: ret}, nil
}

// Checker ...
func Checker(message string) int {
	maxLen := 0
	curLen := 0
	for _, ch := range message {
		if ch == 'o' {
			curLen++
		} else {
			if curLen > maxLen {
				maxLen = curLen
			}
			curLen = 0
		}
	}
	if curLen > maxLen {
		return curLen
	}
	return maxLen
}
