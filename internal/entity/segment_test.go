package entity_test

import (
	"testing"

	"github.com/AnatoliyBr/dynamic-user-segmentation-service/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestSegment_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		slug    string
		isValid bool
	}{
		{
			name:    "valid",
			slug:    "AVITO_DISCOUNT_30",
			isValid: true,
		},
		{
			name:    "mixedcase with whitespace",
			slug:    "avito_DISCOUNT  30 ",
			isValid: true,
		},
		{
			name:    "empty",
			slug:    "",
			isValid: false,
		},
		{
			name:    "invalid symbols",
			slug:    "AVITO ?#@*&%!",
			isValid: false,
		},
		{
			name:    "long slug",
			slug:    "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF",
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			seg := &entity.Segment{Slug: tc.slug}
			if tc.isValid {
				assert.NoError(t, seg.Validate())
			} else {
				assert.Error(t, seg.Validate())
			}
		})
	}
}
