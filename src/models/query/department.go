package query

import validation "github.com/go-ozzo/ozzo-validation"

type DepartmentQuery struct {
	Limit   int64  `query:"limit"`
	Page    int64  `query:"page"`
	Keyword string `query:"keyword"`
}

func (query DepartmentQuery) ValidateDepartmentQuery() error {
	return validation.ValidateStruct(&query,
		validation.Field(
			&query.Limit,
			validation.Required,
			validation.Min(1).Error("Limit is > 0"),
			validation.Max(20).Error("Limit is < 20"),
		),
		validation.Field(
			&query.Page,
			validation.Required,
			validation.Min(1).Error("Limit is > 0"),
		),
	)
}
