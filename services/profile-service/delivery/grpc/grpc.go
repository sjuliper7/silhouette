package grpc

import "github.com/sjuliper7/silhouette/services/profile-service/usecase"

type ProfileServer struct {
	profileUc usecase.ProfileUsecase
}

func NewProfileServer(profileUsecase usecase.ProfileUsecase) ProfileServer {
	return ProfileServer{profileUc: profileUsecase}
}
