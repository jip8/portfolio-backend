package portfolio

import "errors"

var (
	// Experiences
	ErrExperienceIdIsRequired = 			errors.New("id is required")
	ErrExperienceTitleIsRequired = 			errors.New("title is required")
	ErrExperienceFunctionIsRequired = 		errors.New("function is required")
	ErrExperienceDescriptionIsRequired = 	errors.New("description is required")
	ErrExperienceInvalidDate = 				errors.New("invalid date")
)