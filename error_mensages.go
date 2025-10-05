package portfolio

import "errors"

var (
	// General
	ErrInvalidRequestBody  = errors.New("Invalid request body")
	ErrInvalidIDFormat     = errors.New("Invalid ID format")
	ErrNotFound            = errors.New("Not found")
	ErrInvalidLimitFormat  = errors.New("Invalid 'limit' format")
	ErrInvalidOffsetFormat = errors.New("Invalid 'offset' format")

	// Experiences
	ErrExperienceIdIsRequired        	= errors.New("id is required")
	ErrExperienceTitleIsRequired     	= errors.New("title is required")
	ErrExperienceFunctionIsRequired  	= errors.New("function is required")
	ErrExperienceDescriptionIsRequired 	= errors.New("description is required")
	ErrExperienceInvalidDate        	= errors.New("invalid date")

	// Courses
	ErrCourseIdIsRequired        	= errors.New("id is required")
	ErrCourseTitleIsRequired     	= errors.New("title is required")
	ErrCourseInvalidDate        	= errors.New("invalid date")
)
