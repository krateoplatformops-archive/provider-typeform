package typeform

import "strings"

type Settings struct {
	IsPublic               *bool   `json:"is_public,omitempty"`
	RedirectAfterSubmitURL *string `json:"redirect_after_submit_url,omitempty"`
	ShowProgressBar        *bool   `json:"show_progress_bar,omitempty"`
	ShowTimeToComplete     *bool   `json:"show_time_to_complete,omitempty"`
}

type Form struct {
	ID              string           `json:"id,omitempty"`
	Title           string           `json:"title"`
	Type            string           `json:"type"`
	Settings        *Settings        `json:"settings,omitempty"`
	WelcomeScreens  []WelcomeScreen  `json:"welcome_screens,omitempty"`
	ThankyouScreens []ThankyouScreen `json:"thankyou_screens,omitempty"`
	Theme           *struct {
		Href string `json:"href"`
	} `json:"theme,omitempty"`
	Fields []Field `json:"fields,omitempty"`
	Links  *struct {
		Display string `json:"display"`
	} `json:"_links,omitempty"`
}

type Field struct {
	ID          string       `json:"id,omitempty"`
	Type        string       `json:"type"`
	Title       string       `json:"title"`
	Ref         string       `json:"ref,omitempty"`
	Properties  Properties   `json:"properties,omitempty"`
	Validations *Validations `json:"validations,omitempty"`
	Layout      *Layout      `json:"layout,omitempty"`
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

type Attachment struct {
	Href  string `json:"href,omitempty"`
	Type  string `json:"type,omitempty"`
	Scale *int   `json:"scale,omitempty"`
}

type Layout struct {
	Attachment Attachment `json:"attachment,omitempty"`
	Placement  string     `json:"placement,omitempty"`
	Type       string     `json:"type,omitempty"`
}

type WelcomeScreen struct {
	Ref        string                   `json:"ref,omitempty"`
	Title      string                   `json:"title"`
	Layout     *Layout                  `json:"layout,omitempty"`
	Properties *WelcomeScreenProperties `json:"properties,omitempty"`
	Attachment *Attachment              `json:"attachment,omitempty"`
}

type WelcomeScreenProperties struct {
	ButtonText  string `json:"button_text,omitempty"`
	Description string `json:"description,omitempty"`
	ShowButton  *bool  `json:"show_button,omitempty"`
}

type ThankyouScreen struct {
	Title      string                    `json:"title"`
	Ref        string                    `json:"ref,omitempty"`
	Attachment *Attachment               `json:"attachment,omitempty"`
	Properties *ThankYouScreenProperties `json:"properties,omitempty"`
	Layout     *Layout                   `json:"layout,omitempty"`
}

type ThankYouScreenProperties struct {
	ButtonMode  string `json:"button_mode,omitempty"`
	ButtonText  string `json:"button_text,omitempty"`
	RedirectURL string `json:"redirect_url,omitempty"`
	ShareIcons  *bool  `json:"share_icons,omitempty"`
	ShowButton  *bool  `json:"show_button,omitempty"`
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
