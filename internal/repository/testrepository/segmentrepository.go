package testrepository

import (
	"github.com/AnatoliyBr/dynamic-user-segmentation-service/internal/entity"
	"github.com/AnatoliyBr/dynamic-user-segmentation-service/internal/repository"
)

type Pair struct {
	userID int
	segID  int
}

type SegmentRepository struct {
	segments          map[int]*entity.Segment
	usersWithSegments map[Pair]*entity.Segment
}

func NewSegmentRepository() *SegmentRepository {
	return &SegmentRepository{
		segments:          make(map[int]*entity.Segment),
		usersWithSegments: make(map[Pair]*entity.Segment),
	}
}

func (r *SegmentRepository) Create(seg *entity.Segment) error {
	if err := seg.Validate(); err != nil {
		return err
	}

	seg.SegID = len(r.segments) + 1
	r.segments[seg.SegID] = seg

	return nil
}

func (r *SegmentRepository) FindBySlug(slug string) (*entity.Segment, error) {
	for _, seg := range r.segments {
		if seg.Slug == slug {
			return seg, nil
		}
	}
	return nil, repository.ErrRecordNotFound
}

func (r *SegmentRepository) Delete(seg *entity.Segment) error {
	if _, ok := r.segments[seg.SegID]; !ok {
		return repository.ErrRecordNotFound
	}
	delete(r.segments, seg.SegID)

	for key := range r.usersWithSegments {
		if key.segID == seg.SegID {
			delete(r.usersWithSegments, key)
		}
	}
	return nil
}

func (r *SegmentRepository) AddUserToSegments(userID int, segList []*entity.Segment) error {
	for _, seg := range segList {
		if _, ok := r.segments[seg.SegID]; !ok {
			return repository.ErrRecordNotFound
		}
		r.usersWithSegments[Pair{userID: userID, segID: seg.SegID}] = seg
	}
	return nil
}

func (r *SegmentRepository) DeleteUserFromSegments(userID int, segList []*entity.Segment) error {
	for _, segDel := range segList {
		for key := range r.usersWithSegments {
			if key.segID == segDel.SegID && key.userID == userID {
				delete(r.usersWithSegments, key)
			}
		}
	}
	return nil
}

func (r *SegmentRepository) FindByUser(userID int) ([]*entity.Segment, error) {
	segList := make([]*entity.Segment, 0)

	for key, seg := range r.usersWithSegments {
		if key.userID == userID {
			segList = append(segList, seg)
		}
	}

	if len(segList) > 0 {
		return segList, nil
	} else {
		return nil, repository.ErrRecordNotFound
	}
}
