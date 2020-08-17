// Code generated by goa v3.2.3, DO NOT EDIT.
//
// pipeline views
//
// Command:
// $ goa gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// EnduroStoredPipeline is the viewed result type that is projected based on a
// view.
type EnduroStoredPipeline struct {
	// Type to project
	Projected *EnduroStoredPipelineView
	// View to render
	View string
}

// EnduroStoredPipelineView is a type that runs validations on a projected type.
type EnduroStoredPipelineView struct {
	// Name of the collection
	ID *string
	// Name of the collection
	Name *string
}

var (
	// EnduroStoredPipelineMap is a map of attribute names in result type
	// EnduroStoredPipeline indexed by view name.
	EnduroStoredPipelineMap = map[string][]string{
		"default": []string{
			"id",
			"name",
		},
	}
)

// ValidateEnduroStoredPipeline runs the validations defined on the viewed
// result type EnduroStoredPipeline.
func ValidateEnduroStoredPipeline(result *EnduroStoredPipeline) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateEnduroStoredPipelineView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateEnduroStoredPipelineView runs the validations defined on
// EnduroStoredPipelineView using the "default" view.
func ValidateEnduroStoredPipelineView(result *EnduroStoredPipelineView) (err error) {
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.id", *result.ID, goa.FormatUUID))
	}
	return
}
