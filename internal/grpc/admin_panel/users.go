package adminpanelgrpc

import (
	"context"
	adminpanelv1 "github.com/curtrika/UMetrika_server/pkg/proto/admin_panel/v1"
)

func (s *serverAPI) CreateUser(ctx context.Context, req *adminpanelv1.CreateUserRequest) (*adminpanelv1.CreateUserResponse, error) {
	userModel := s.cvt.UserToModel(req.User)
	newUser, err := s.adminPanel.CreateUser(ctx, *userModel)
	if err != nil {
		return nil, err
	}

	return &adminpanelv1.CreateUserResponse{
		NewUser: s.cvt.ModelToUser(newUser),
	}, nil
}
