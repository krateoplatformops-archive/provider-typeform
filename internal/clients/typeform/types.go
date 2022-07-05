package typeform

import "strings"

type Form struct {
	ID     string  `json:"id,omitempty"`
	Title  string  `json:"title"`
	Type   string  `json:"type"`
	Fields []Field `json:"fields,omitempty"`
	Links  *struct {
		Display string `json:"display"`
	} `json:"_links,omitempty"`
}

type Field struct {
	ID          string       `json:"id,omitempty"`
	Type        string       `json:"type"`
	Title       string       `json:"title"`
	Properties  Properties   `json:"properties,omitempty"`
	Validations *Validations `json:"validations,omitempty"`
}

type Properties struct {
	AllowMultipleSelection *bool    `json:"allow_multiple_selection,omitempty"`
	AllowOtherChoice       *bool    `json:"allow_other_choice,omitempty"`
	AlphabeticalOrder      *bool    `json:"alphabetical_order,omitempty"`
	ButtonText             string   `json:"button_text,omitempty"`
	Choices                []Choice `json:"choices,omitempty"`
	Description            *string  `json:"description,omitempty"`
	HideMarks              *bool    `json:"hide_marks,omitempty"`
	Labels                 *Labels  `json:"labels,omitempty"`
	Randomize              *bool    `json:"randomize,omitempty"`
	StartAtOne             *bool    `json:"start_at_one,omitempty"`
	Steps                  *int     `json:"steps,omitempty"`
	Shape                  *string  `json:"shape,omitempty"`
}

type Choice struct {
	Label string `json:"label"`
}

type Labels struct {
	Center string `json:"center,omitempty"`
	Left   string `json:"left,omitempty"`
	Right  string `json:"right,omitempty"`
}

type Validations struct {
	MaxLength    *int  `json:"max_length,omitempty"`
	MaxSelection *int  `json:"max_selection,omitempty"`
	MaxValue     *int  `json:"max_value,omitempty"`
	MinSelection *int  `json:"min_selection,omitempty"`
	MinValue     *int  `json:"min_value,omitempty"`
	Required     *bool `json:"required,omitempty"`
}

type Errors struct {
	Code        string `json:"code"`
	Description string `json:"description,omitempty"`
	Details     []struct {
		Code        string `json:"code,omitempty"`
		Description string `json:"description,omitempty"`
		Field       string `json:"field,omitempty"`
	} `json:"details,omitempty"`
}

func (e *Errors) Error() string {
	sb := new(strings.Builder)
	sb.WriteString(e.Code)
	if e.Description != "" {
		sb.WriteString(": ")
		sb.WriteString(e.Description)
	}
	return sb.String()
}
