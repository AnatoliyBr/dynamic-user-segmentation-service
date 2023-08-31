package usecase

import (
	"github.com/AnatoliyBr/dynamic-user-segmentation-service/internal/entity"
	"github.com/AnatoliyBr/dynamic-user-segmentation-service/internal/repository"
)

type AppUseCase struct {
	segmentRepository repository.SegmentRepository
}

func NewAppUseCase(r repository.SegmentRepository) *AppUseCase {
	return &AppUseCase{
		segmentRepository: r,
	}
}

func (uc *AppUseCase) SegmentCreate(seg *entity.Segment) error {
	return uc.segmentRepository.Create(seg)
}

func (uc *AppUseCase) SegmentFindBySlug(slug string) (*entity.Segment, error) {
	return uc.segmentRepository.FindBySlug(slug)
}

func (uc *AppUseCase) SegmentDelete(seg *entity.Segment) error {
	return uc.segmentRepository.Delete(seg)
}

func (uc *AppUseCase) AddUserToSegments(userID int, segList []*entity.Segment) error {
	return uc.segmentRepository.AddUserToSegments(userID, segList)
}

func (uc *AppUseCase) DeleteUserFromSegments(userID int, segList []*entity.Segment) error {
	return uc.segmentRepository.DeleteUserFromSegments(userID, segList)
}

func (uc *AppUseCase) SegmentFindByUser(userID int) ([]*entity.Segment, error) {
	return uc.segmentRepository.FindByUser(userID)
}
