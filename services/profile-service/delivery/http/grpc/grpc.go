package grpc

import "github.com/sjuliper7/silhouette/services/profile-service/usecase"

//ProfileService ...
type ProfileService struct {
	profileUsecase usecase.ProfileUseCase
}

//NewProfileServer ...
func NewProfileServer(profileCase usecase.ProfileUseCase) ProfileService {
	return ProfileService{profileUsecase: profileCase}
}
