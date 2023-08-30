package entity

import (
	"regexp"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Segment struct {
	SegID int    `json:"seg_id"`
	Slug  string `json:"slug"`
}

func (s *Segment) Validate() error {
	s.Slug = strings.Join(strings.Fields(s.Slug), "_")
	s.Slug = strings.ToUpper(s.Slug)

	return validation.ValidateStruct(
		s,
		validation.Field(
			&s.Slug,
			validation.Required,
			validation.Match(regexp.MustCompile(`\w`)),
			validation.Length(0, 50),
		),
	)
}
