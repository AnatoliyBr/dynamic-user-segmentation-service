package usecase_test

import (
	"testing"

	"github.com/AnatoliyBr/dynamic-user-segmentation-service/internal/entity"
	"github.com/AnatoliyBr/dynamic-user-segmentation-service/internal/repository"
	"github.com/AnatoliyBr/dynamic-user-segmentation-service/internal/repository/testrepository"
	"github.com/AnatoliyBr/dynamic-user-segmentation-service/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestAppUseCase_SegmentCreate(t *testing.T) {
	r := testrepository.NewSegmentRepository()
	uc := usecase.NewAppUseCase(r)
	seg := &entity.Segment{Slug: "AVITO_DISCOUNT_30"}

	assert.NoError(t, uc.SegmentCreate(seg))
}

func TestAppUseCase_SegmentFindBySlug(t *testing.T) {
	r := testrepository.NewSegmentRepository()
	uc := usecase.NewAppUseCase(r)
	seg1 := &entity.Segment{Slug: "AVITO_DISCOUNT_30"}
	_, err := uc.SegmentFindBySlug(seg1.Slug)
	assert.EqualError(t, err, repository.ErrRecordNotFound.Error())

	uc.SegmentCreate(seg1)
	seg2, err := uc.SegmentFindBySlug(seg1.Slug)
	assert.NoError(t, err)
	assert.NotNil(t, seg2)
}

func TestAppUseCase_SegmentDelete(t *testing.T) {
	r := testrepository.NewSegmentRepository()
	uc := usecase.NewAppUseCase(r)

	userID := 1
	segList := []*entity.Segment{
		{Slug: "AVITO_DISCOUNT_30"},
		{Slug: "AVITO_DISCOUNT_50"},
		{Slug: "AVITO_VOICE_MESSAGES"},
	}

	uc.SegmentCreate(segList[0])
	uc.SegmentCreate(segList[1])
	uc.AddUserToSegments(userID, segList[0:2])

	err := uc.SegmentDelete(segList[0])
	assert.NoError(t, err)

	err = uc.SegmentDelete(segList[2])
	assert.EqualError(t, err, repository.ErrRecordNotFound.Error())
}

func TestAppUseCase_AddUserToSegments(t *testing.T) {
	r := testrepository.NewSegmentRepository()
	uc := usecase.NewAppUseCase(r)

	userID := 1
	segList := []*entity.Segment{
		{Slug: "AVITO_DISCOUNT_30"},
		{Slug: "AVITO_DISCOUNT_50"},
	}

	err := uc.AddUserToSegments(userID, segList)
	assert.EqualError(t, err, repository.ErrRecordNotFound.Error())

	uc.SegmentCreate(segList[0])
	uc.SegmentCreate(segList[1])

	err = uc.AddUserToSegments(userID, segList)
	assert.NoError(t, err)
}

func TestAppUseCase_DeleteUserFromSegments(t *testing.T) {
	r := testrepository.NewSegmentRepository()
	uc := usecase.NewAppUseCase(r)

	userID := 1
	segList := []*entity.Segment{
		{Slug: "AVITO_DISCOUNT_30"},
		{Slug: "AVITO_DISCOUNT_50"},
	}

	uc.SegmentCreate(segList[0])
	uc.SegmentCreate(segList[1])
	uc.AddUserToSegments(userID, segList)

	err := uc.DeleteUserFromSegments(userID, segList)
	assert.NoError(t, err)
}

func TestAppUseCase_SegmentFindByUser(t *testing.T) {
	r := testrepository.NewSegmentRepository()
	uc := usecase.NewAppUseCase(r)

	userID := 1
	segList1 := []*entity.Segment{
		{Slug: "AVITO_DISCOUNT_30"},
		{Slug: "AVITO_DISCOUNT_50"},
	}

	uc.SegmentCreate(segList1[0])
	uc.SegmentCreate(segList1[1])

	_, err := uc.SegmentFindByUser(userID)
	assert.EqualError(t, err, repository.ErrRecordNotFound.Error())

	uc.AddUserToSegments(userID, segList1)
	segList2, err := uc.SegmentFindByUser(userID)
	assert.NoError(t, err)
	assert.NotNil(t, segList2)
}
