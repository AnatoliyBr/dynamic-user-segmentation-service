package sqlrepository_test

import (
	"testing"

	"github.com/AnatoliyBr/dynamic-user-segmentation-service/internal/entity"
	"github.com/AnatoliyBr/dynamic-user-segmentation-service/internal/repository"
	"github.com/AnatoliyBr/dynamic-user-segmentation-service/internal/repository/sqlrepository"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestSegmentRepository_Create(t *testing.T) {
	db, teardown := sqlrepository.TestDB(t, testDatabaseURL)
	defer teardown("segments")

	r := sqlrepository.NewSegmentRepository(db)
	seg := &entity.Segment{Slug: "AVITO_DISCOUNT_30"}

	assert.NoError(t, r.Create(seg))
}

func TestSegmentRepository_FindBySlug(t *testing.T) {
	db, teardown := sqlrepository.TestDB(t, testDatabaseURL)
	defer teardown("segments")

	r := sqlrepository.NewSegmentRepository(db)
	seg1 := &entity.Segment{Slug: "AVITO_DISCOUNT_30"}
	_, err := r.FindBySlug(seg1.Slug)
	assert.EqualError(t, err, repository.ErrRecordNotFound.Error())

	r.Create(seg1)
	seg2, err := r.FindBySlug(seg1.Slug)
	assert.NoError(t, err)
	assert.NotNil(t, seg2)
}

func TestSegmentRepository_Delete(t *testing.T) {
	db, teardown := sqlrepository.TestDB(t, testDatabaseURL)
	defer teardown("segments")

	r := sqlrepository.NewSegmentRepository(db)
	seg := &entity.Segment{Slug: "AVITO_DISCOUNT_30"}
	r.Create(seg)
	err := r.Delete(seg)
	assert.NoError(t, err)
}

func TestSegmentRepository_AddUserToSegments(t *testing.T) {
	db, teardown := sqlrepository.TestDB(t, testDatabaseURL)
	defer teardown("users_with_segments", "segments")

	r := sqlrepository.NewSegmentRepository(db)

	userID := 1
	segList := []*entity.Segment{
		{Slug: "AVITO_DISCOUNT_30"},
		{Slug: "AVITO_DISCOUNT_50"},
	}

	err := r.AddUserToSegments(userID, segList)
	assert.Error(t, err)

	r.Create(segList[0])
	r.Create(segList[1])

	err = r.AddUserToSegments(userID, segList)
	assert.NoError(t, err)
}

func TestSegmentRepository_DeleteUserFromSegments(t *testing.T) {
	db, teardown := sqlrepository.TestDB(t, testDatabaseURL)
	defer teardown("users_with_segments", "segments")

	r := sqlrepository.NewSegmentRepository(db)

	userID := 1
	segList := []*entity.Segment{
		{Slug: "AVITO_DISCOUNT_30"},
		{Slug: "AVITO_DISCOUNT_50"},
	}

	r.Create(segList[0])
	r.Create(segList[1])
	r.AddUserToSegments(userID, segList)

	err := r.DeleteUserFromSegments(userID, segList)
	assert.NoError(t, err)
}

func TestSegmentRepository_FindByUser(t *testing.T) {
	db, teardown := sqlrepository.TestDB(t, testDatabaseURL)
	defer teardown("users_with_segments", "segments")

	r := sqlrepository.NewSegmentRepository(db)

	userID := 1
	segList1 := []*entity.Segment{
		{Slug: "AVITO_DISCOUNT_30"},
		{Slug: "AVITO_DISCOUNT_50"},
	}

	r.Create(segList1[0])
	r.Create(segList1[1])

	_, err := r.FindByUser(userID)
	assert.EqualError(t, err, repository.ErrRecordNotFound.Error())

	r.AddUserToSegments(userID, segList1)
	segList2, err := r.FindByUser(userID)
	assert.NoError(t, err)
	assert.NotNil(t, segList2)
}
