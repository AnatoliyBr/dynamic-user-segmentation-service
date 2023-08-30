package sqlrepository

import (
	"database/sql"

	"github.com/AnatoliyBr/dynamic-user-segmentation-service/internal/entity"
	"github.com/AnatoliyBr/dynamic-user-segmentation-service/internal/repository"
)

type SegmentRepository struct {
	db *sql.DB
}

func NewSegmentRepository(db *sql.DB) *SegmentRepository {
	return &SegmentRepository{
		db: db,
	}
}

func (r *SegmentRepository) Create(seg *entity.Segment) error {
	if err := seg.Validate(); err != nil {
		return err
	}

	return r.db.QueryRow(
		"INSERT INTO segments (slug) VALUES ($1) RETURNING seg_id",
		seg.Slug,
	).Scan(&seg.SegID)
}

func (r *SegmentRepository) FindBySlug(slug string) (*entity.Segment, error) {
	seg := &entity.Segment{}
	if err := r.db.QueryRow(
		"SELECT seg_id, slug FROM segments WHERE slug = $1",
		slug,
	).Scan(
		&seg.SegID,
		&seg.Slug,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, repository.ErrRecordNotFound
		}
		return nil, err
	}
	return seg, nil
}

func (r *SegmentRepository) Delete(seg *entity.Segment) error {
	_, err := r.db.Exec(
		"DELETE FROM segments WHERE slug = $1",
		seg.Slug)
	if err != nil {
		return err
	}
	return nil
}

func (r *SegmentRepository) AddUserToSegments(userID int, segList []*entity.Segment) error {
	stmt, err := r.db.Prepare(
		"INSERT INTO users_with_segments (user_id, seg_id) VALUES ($1, $2)")
	if err != nil {
		return err
	}

	for _, seg := range segList {
		_, err := stmt.Exec(userID, seg.SegID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *SegmentRepository) DeleteUserFromSegments(userID int, segList []*entity.Segment) error {
	stmt, err := r.db.Prepare(
		"DELETE FROM users_with_segments WHERE user_id = $1 AND seg_id = $2")
	if err != nil {
		return err
	}

	for _, seg := range segList {
		_, err := stmt.Exec(userID, seg.SegID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *SegmentRepository) FindByUser(userID int) ([]*entity.Segment, error) {
	segList := make([]*entity.Segment, 0)

	rows, err := r.db.Query(
		"SELECT seg_id, slug FROM segments WHERE seg_id IN (SELECT seg_id FROM users_with_segments WHERE user_id = $1)",
		userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var seg_id int
		var slug string

		err := rows.Scan(&seg_id, &slug)
		if err != nil {
			return nil, err
		}
		segList = append(segList, &entity.Segment{SegID: seg_id, Slug: slug})
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(segList) > 0 {
		return segList, nil
	} else {
		return nil, repository.ErrRecordNotFound
	}
}
