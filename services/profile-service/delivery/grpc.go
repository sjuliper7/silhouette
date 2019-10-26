package delivery

import "github.com/sjuliper7/silhouette/services/profile-service/usecase"

type ProfileServer struct {
	usecase usecase.ProfileUsecase
}

func NewProfileServer(profileUsecase usecase.ProfileUsecase) ProfileServer {
	return ProfileServer{usecase: profileUsecase}
}
