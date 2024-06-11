package list

// Code generated by jsonapi. DO NOT EDIT.
// Source: github.com/pentops/o5-test-app/internal/genclient/psm/list/v1/list

import ()

// QueryRequest Proto: psm.list.v1.QueryRequest
type QueryRequest struct {
	Searches []*Search `json:"searches,omitempty"`
	Sorts    []*Sort   `json:"sorts,omitempty"`
	Filters  []*Filter `json:"filters,omitempty"`
}

// Field_type Proto: psm.list.v1.Field.type
type Field_type struct {
	Value *string `json:"value,omitempty"`
	Range *Range  `json:"range,omitempty"`
}

func (s Field_type) OneofKey() string {
	if s.Value != nil {
		return "value"
	}
	if s.Range != nil {
		return "range"
	}
	return ""
}

func (s Field_type) Type() interface{} {
	if s.Value != nil {
		return s.Value
	}
	if s.Range != nil {
		return s.Range
	}
	return nil
}

// Range Proto: psm.list.v1.Range
type Range struct {
	Min string `json:"min,omitempty"`
	Max string `json:"max,omitempty"`
}

// And Proto: psm.list.v1.And
type And struct {
	Filters []*Filter `json:"filters,omitempty"`
}

// Or Proto: psm.list.v1.Or
type Or struct {
	Filters []*Filter `json:"filters,omitempty"`
}

// PageResponse Proto: psm.list.v1.PageResponse
type PageResponse struct {
	NextToken *string `json:"nextToken,omitempty"`
}

// PageRequest Proto: psm.list.v1.PageRequest
type PageRequest struct {
	Token    *string `json:"token,omitempty"`
	PageSize *int64  `json:"pageSize,omitempty"`
}

// Search Proto: psm.list.v1.Search
type Search struct {
	Field string `json:"field,omitempty"`
	Value string `json:"value,omitempty"`
}

// Sort Proto: psm.list.v1.Sort
type Sort struct {
	Field      string `json:"field,omitempty"`
	Descending bool   `json:"descending"`
}

// Filter Proto: psm.list.v1.Filter
type Filter struct {
	Type *Filter_type `json:"type,omitempty"`
}

// Filter_type Proto: psm.list.v1.Filter.type
type Filter_type struct {
	Field *Field `json:"field,omitempty"`
	And   *And   `json:"and,omitempty"`
	Or    *Or    `json:"or,omitempty"`
}

func (s Filter_type) OneofKey() string {
	if s.Field != nil {
		return "field"
	}
	if s.And != nil {
		return "and"
	}
	if s.Or != nil {
		return "or"
	}
	return ""
}

func (s Filter_type) Type() interface{} {
	if s.Field != nil {
		return s.Field
	}
	if s.And != nil {
		return s.And
	}
	if s.Or != nil {
		return s.Or
	}
	return nil
}

// Field Proto: psm.list.v1.Field
type Field struct {
	Name string      `json:"name,omitempty"`
	Type *Field_type `json:"type,omitempty"`
}
