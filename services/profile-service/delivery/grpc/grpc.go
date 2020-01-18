package grpc

import "github.com/sjuliper7/silhouette/services/profile-service/usecase"

type ProfileServer struct {
	profileUsecase usecase.ProfileUseCase
}

func NewProfileServer(profileCase usecase.ProfileUseCase) ProfileServer {
	return ProfileServer{profileUsecase: profileCase}
}
