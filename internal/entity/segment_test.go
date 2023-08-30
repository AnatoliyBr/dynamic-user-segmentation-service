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
			name:    "uppercase",
			slug:    "AVITO_DISCOUNT_30",
			isValid: true,
		},
		{
			name:    "lowercase",
			slug:    "avito_discount_30",
			isValid: true,
		},
		{
			name:    "mixedcase",
			slug:    "avito_DISCOUNT_30",
			isValid: true,
		},
		{
			name:    "with whitespace",
			slug:    "AVITO    DISCOUNT 30 ",
			isValid: true,
		},
		{
			name:    "invalid symbols",
			slug:    "?#@*&%!",
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
