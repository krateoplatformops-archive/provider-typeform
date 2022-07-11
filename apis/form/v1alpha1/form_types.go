package v1alpha1

import (
	"reflect"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Properties describe a form field attributes.
type Properties struct {
	// +optional
	// AllowMultipleSelection: true to allow respondents to select more than one answer choice.
	// Available for types: multiple_choice
	AllowMultipleSelection *bool `json:"allowMultipleSelection,omitempty"`

	// AllowOtherChoice: true to include an "Other" option so respondents can enter a different
	// answer choice from those listed; false to limit answer choices to those listed.
	// Available for types: multiple_choice
	// +optional
	AllowOtherChoice *bool `json:"allowOtherChoice,omitempty"`

	// AlphabeticalOrder: true if question should list dropdown answer choices in
	// alphabetical order; false if question should list dropdown answer choices in
	// the order they're listed in the "choices" array.
	// Available for types: dropdown.
	// +optional
	AlphabeticalOrder *bool `json:"alphabeticalOrder,omitempty"`

	// Choices: answer choices.
	// Available for types: dropdown, multiple_choice
	// +optional
	Choices []string `json:"choices,omitempty"`

	// Description: Question or instruction to display for the field.
	// +optional
	Description *string `json:"description,omitempty"`

	// Steps: Number of steps in the scale's range.
	// Minimum is 5 and maximum is 11.
	// Available fortypes: rating.
	// +optional
	Steps *int `json:"steps,omitempty"`

	// Shape: Shape to display on the scale's steps.
	// Valid values: cat, circle, cloud, crown, dog, droplet, flag, heart, lightbulb, pencil, skull,
	// star, thunderbolt, tick, trophy, up, user. Default: star
	// Available for types: rating types.
	// +optional
	Shape *string `json:"shape,omitempty"`
}

// Validations define specific field validation criteria.
type Validations struct {
	// MaxLength: maximum number of characters allowed in the answer.
	// Available for types: long_text, short_text.
	// +optional
	MaxLength *int `json:"maxLength,omitempty"`

	// MaxSelection: maximum selections allowed in the answer, must be a positive integer.
	// Available for types: multiple_choice.
	// +optional
	MaxSelection *int `json:"maxSelection,omitempty"`

	// MinSelection: minimum selections allowed in the answer, must be a positive integer.
	// Available for types: multiple_choice
	// +optional
	MinSelection *int `json:"minSelection,omitempty"`

	// Required: true if respondents must provide an answer. Otherwise, false.
	// Available for types: dropdown, long_text, multiple_choice, rating.
	// +optional
	Required *bool `json:"required,omitempty"`
}

// Field describe a form widget.
type Field struct {
	// Type: the widget type.
	// Valid values: dropdown, long_text, multiple_choice, rating.
	Type string `json:"type"`

	// Title: Unique name you assign to the field on this form.
	Title string `json:"title"`

	// Ref: Readable name you can use to reference the field.
	// +optional
	Ref *string `json:"ref,omitempty"`

	// Properties: specific properties for this field.
	// +optional
	Properties Properties `json:"properties,omitempty"`

	// Validations: specific properties for this field.
	// +optional
	Validations *Validations `json:"validations,omitempty"`

	// Layout: allows you to position attachment.
	// +optional
	Layout *Layout `json:"layout,omitempty"`
}

// Attachment allows you to display images and videos.
// Available for welcome and thank you screens, as well as all fields.
type Attachment struct {
	// URL for the image or video you want to display.
	// Images must already exist in your account
	Href string `json:"href,omitempty"`

	// Type of attachment. Valid values: image, video
	Type string `json:"type,omitempty"`

	// Optional parameter for responsively scaling videos.
	// Valid values: 0.4, 0.6, 0.8, 1
	// +optional
	Scale *int `json:"scale,omitempty"`
}

// Allows you to position attachment.
// Available in all fields, Welcome and Thank you screens.
type Layout struct {
	Attachment Attachment `json:"attachment,omitempty"`

	// Position of media for split and float layouts.
	// Valid values: left, right.
	Placement string `json:"placement,omitempty"`

	// Type of layout. Valid values: split, wallpaper, float.
	Type string `json:"type,omitempty"`
}

type WelcomeScreen struct {
	Ref        string                   `json:"ref,omitempty"`
	Title      string                   `json:"title"`
	Layout     *Layout                  `json:"layout,omitempty"`
	Properties *WelcomeScreenProperties `json:"properties,omitempty"`
	Attachment *Attachment              `json:"attachment,omitempty"`
}

type WelcomeScreenProperties struct {
	ButtonText  string `json:"buttonText,omitempty"`
	Description string `json:"description,omitempty"`
	ShowButton  *bool  `json:"showButton,omitempty"`
}

type ThankyouScreen struct {
	Title      string                    `json:"title"`
	Ref        string                    `json:"ref,omitempty"`
	Attachment *Attachment               `json:"attachment,omitempty"`
	Properties *ThankYouScreenProperties `json:"properties,omitempty"`
	Layout     *Layout                   `json:"layout,omitempty"`
}

type ThankYouScreenProperties struct {
	ButtonMode  string `json:"buttonMode,omitempty"`
	ButtonText  string `json:"buttonText,omitempty"`
	RedirectURL string `json:"redirectUrl,omitempty"`
	ShareIcons  *bool  `json:"shareIcons,omitempty"`
	ShowButton  *bool  `json:"showButton,omitempty"`
}

// FormParams are the configurable fields of a form instance.
type FormParams struct {
	// Title: to use for this form.
	Title string `json:"title"`

	// Fields: list of form widgets.
	Fields []Field `json:"fields,omitempty"`

	// WelcomeScreens: list of form welcome screens.
	WelcomeScreens []WelcomeScreen `json:"welcomeScreens,omitempty"`

	// ThankyouScreens: list of form thank you screens.
	ThankyouScreens []ThankyouScreen `json:"thankyouScreens,omitempty"`
}

// FormObservation are the observable fields of a Form.
type FormObservation struct {
	ID         string `json:"id"`
	DisplayURL string `json:"displayUrl"`
}

// A FormSpec defines the desired state of a Form.
type FormSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       FormParams `json:"forProvider"`
}

// A FormStatus represents the observed state of a Form.
type FormStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          FormObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Form is a typeform API type.
// +kubebuilder:printcolumn:name="ID",type="string",JSONPath=".status.atProvider.id"
// +kubebuilder:printcolumn:name="DISPLAY_URL",type="string",JSONPath=".status.atProvider.displayUrl"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status",priority=1
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status",priority=1
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,krateo,typeform}
type Form struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FormSpec   `json:"spec"`
	Status FormStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FormList contains a list of Form
type FormList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Form `json:"items"`
}

// Form type metadata.
var (
	FormKind             = reflect.TypeOf(Form{}).Name()
	FormGroupKind        = schema.GroupKind{Group: Group, Kind: FormKind}.String()
	FormKindAPIVersion   = FormKind + "." + SchemeGroupVersion.String()
	FormGroupVersionKind = SchemeGroupVersion.WithKind(FormKind)
)

func init() {
	SchemeBuilder.Register(&Form{}, &FormList{})
}
