package httpserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AnatoliyBr/dynamic-user-segmentation-service/internal/entity"
	"github.com/AnatoliyBr/dynamic-user-segmentation-service/internal/repository/testrepository"
	"github.com/AnatoliyBr/dynamic-user-segmentation-service/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleHello(t *testing.T) {
	r := testrepository.NewSegmentRepository()
	uc := usecase.NewAppUseCase(r)
	s := NewServer(NewConfig(), uc)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)

	s.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotNil(t, rec.Body)
}

func TestServer_HandleSegmentsCreate(t *testing.T) {
	r := testrepository.NewSegmentRepository()
	uc := usecase.NewAppUseCase(r)
	s := NewServer(NewConfig(), uc)

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"slug": "AVITO_DISCOUNT_30",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid symbols",
			payload: map[string]string{
				"slug": "?#@*&%!",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/seg", b)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleSegmentsDelete(t *testing.T) {
	r := testrepository.NewSegmentRepository()
	uc := usecase.NewAppUseCase(r)
	s := NewServer(NewConfig(), uc)

	userID := 1
	segList := []*entity.Segment{{Slug: "AVITO_DISCOUNT_30"}}

	s.uc.SegmentCreate(segList[0])
	s.uc.AddUserToSegments(userID, segList)

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"slug": "AVITO_DISCOUNT_30",
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid payload",
			payload:      "",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "seg not found",
			payload: map[string]string{
				"slug": "AVITO_DISCOUNT_50",
			},
			expectedCode: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodDelete, "/seg", b)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleSegmentsUpdateUser(t *testing.T) {
	type request struct {
		SlugListAdd []string `json:"slug_list_add"`
		SlugListDel []string `json:"slug_list_del"`
		UserID      int      `json:"user_id"`
	}

	r := testrepository.NewSegmentRepository()
	uc := usecase.NewAppUseCase(r)
	s := NewServer(NewConfig(), uc)

	userID := 1
	segList := []*entity.Segment{
		{Slug: "AVITO_VOICE_MESSAGES"},
		{Slug: "AVITO_PERFORMANCE_VAS"},
		{Slug: "AVITO_DISCOUNT_30"},
		{Slug: "AVITO_DISCOUNT_50"},
	}

	for _, seg := range segList {
		s.uc.SegmentCreate(seg)
	}

	s.uc.AddUserToSegments(userID, segList[0:2])

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: &request{
				SlugListAdd: []string{
					"AVITO_DISCOUNT_30",
					"AVITO_DISCOUNT_50",
				},
				SlugListDel: []string{
					"AVITO_VOICE_MESSAGES",
					"AVITO_PERFORMANCE_VAS",
				},
				UserID: userID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "empty del list",
			payload: &request{
				SlugListAdd: []string{
					"AVITO_DISCOUNT_30",
					"AVITO_DISCOUNT_50",
				},
				SlugListDel: []string{},
				UserID:      userID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "new user",
			payload: &request{
				SlugListAdd: []string{
					"AVITO_DISCOUNT_30",
					"AVITO_DISCOUNT_50",
				},
				SlugListDel: []string{
					"AVITO_VOICE_MESSAGES",
					"AVITO_PERFORMANCE_VAS",
				},
				UserID: userID + 1,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid payload",
			payload:      "",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "add seg not found",
			payload: &request{
				SlugListAdd: []string{
					"NOT_FOUND",
				},
				SlugListDel: []string{},
				UserID:      userID,
			},
			expectedCode: http.StatusNotFound,
		},
		{
			name: "del seg not found",
			payload: &request{
				SlugListAdd: []string{},
				SlugListDel: []string{
					"NOT_FOUND",
				},
				UserID: userID,
			},
			expectedCode: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPut, "/seg", b)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleSegmentsGetByUser(t *testing.T) {
	r := testrepository.NewSegmentRepository()
	uc := usecase.NewAppUseCase(r)
	s := NewServer(NewConfig(), uc)

	userID := 1
	segList := []*entity.Segment{{Slug: "AVITO_DISCOUNT_30"}}

	s.uc.SegmentCreate(segList[0])
	s.uc.AddUserToSegments(userID, segList)

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]int{
				"user_id": 1,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid payload",
			payload:      "",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "seg not found",
			payload: map[string]int{
				"user_id": 2,
			},
			expectedCode: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodGet, "/seg", b)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
