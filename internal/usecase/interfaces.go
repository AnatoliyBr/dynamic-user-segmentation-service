package usecase

import "github.com/AnatoliyBr/dynamic-user-segmentation-service/internal/entity"

type UseCase interface {
	SegmentCreate(*entity.Segment) error
	SegmentFindBySlug(string) (*entity.Segment, error)
	SegmentDelete(*entity.Segment) error
	AddUserToSegments(int, []*entity.Segment) error
	DeleteUserFromSegments(int, []*entity.Segment) error
	SegmentFindByUser(int) ([]*entity.Segment, error)
}
