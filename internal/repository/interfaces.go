package repository

import "github.com/AnatoliyBr/dynamic-user-segmentation-service/internal/entity"

type SegmentRepository interface {
	Create(*entity.Segment) error
	FindBySlug(string) (*entity.Segment, error)
	Delete(*entity.Segment) error
	AddUserToSegments(int, []*entity.Segment) error
	DeleteUserFromSegments(int, []*entity.Segment) error
	FindByUser(int) ([]*entity.Segment, error)
}
