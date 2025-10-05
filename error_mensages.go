package portfolio

import "errors"

var (
	// General
	ErrInvalidRequestBody  = errors.New("invalid request body")
	ErrInvalidIDFormat     = errors.New("invalid ID format")
	ErrNotFound            = errors.New("not found")
	ErrInvalidLimitFormat  = errors.New("invalid 'limit' format")
	ErrInvalidOffsetFormat = errors.New("invalid 'offset' format")

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

	// Projects
	ErrProjectIdIsRequired        	= errors.New("id is required")
	ErrProjectTitleIsRequired     	= errors.New("title is required")
	ErrProjectInvalidDate        	= errors.New("invalid date")

	// Articles
	ErrArticleIdIsRequired        	= errors.New("id is required")
	ErrArticleTitleIsRequired     	= errors.New("title is required")
	ErrArticleInvalidDate        	= errors.New("invalid date")

	// Contacts
	ErrContactIdIsRequired        	= errors.New("id is required")
	ErrContactLinkIsRequired     	= errors.New("link is required")
	ErrContactPlataformIsRequired 	= errors.New("plataform is required")
)
