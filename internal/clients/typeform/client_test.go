//go:build integration
// +build integration

package typeform

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"testing"

	"github.com/krateoplatformops/provider-typeform/internal/clients"
	"github.com/krateoplatformops/provider-typeform/internal/helpers"
	"github.com/lucasepe/dotenv"
	"github.com/stretchr/testify/assert"
)

func TestCreateForm(t *testing.T) {
	if err := setupEnv(); err != nil {
		t.Fatal(err)
	}

	opts := ClientOpts{
		Token: os.Getenv("TYPEFORM_API_TOKEN"),
	}
	if len(os.Getenv("TYPEFORM_VERBOSE")) > 0 {
		opts.HTTPClient = clients.TracerHTTPClient()
	}

	val := &Form{
		Title: "Form di esempio",
		Type:  "form",
		Fields: []Field{
			{
				Type:  "dropdown",
				Title: "Scelta singola",
				Properties: Properties{
					Choices: []Choice{
						{Label: "AAA"}, {Label: "BBB"}, {Label: "CCC"},
					},
				},
				Validations: &Validations{
					Required: helpers.BoolPtr(true),
				},
				Layout: &Layout{
					Type: "split",
					Attachment: Attachment{
						Href: "https://images.typeform.com/images/QQgJMC6EqXen",
						Type: "image",
					},
				},
			},
			{
				Type:  "rating",
				Title: "Valutazione",
				Properties: Properties{
					Shape: helpers.StringPtr("heart"),
					Steps: helpers.IntPtr(5),
				},
				Validations: &Validations{
					Required: helpers.BoolPtr(true),
				},
			},
			{
				Type:  "multiple_choice",
				Title: " Scelta multipla",
				Properties: Properties{
					Choices: []Choice{
						{Label: "AAA"}, {Label: "BBB"}, {Label: "CCC"},
					},
				},
				Validations: &Validations{
					MinSelection: helpers.IntPtr(1),
					MaxSelection: helpers.IntPtr(2),
					Required:     helpers.BoolPtr(true),
				},
			},
			{
				Type:  "long_text",
				Title: " Testo libero",
				Validations: &Validations{
					MaxLength: helpers.IntPtr(200),
					Required:  helpers.BoolPtr(true),
				},
			},
		},
	}

	res, err := NewClient(opts).CreateForm(context.TODO(), val)
	//if err != nil {
	//	var reqErrs *Errors
	//	if errors.As(err, &reqErrs) {
	//		fmt.Println("ECCOLO!!")
	//	}
	//	t.Logf(err.Error())
	//}

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")

	t.Log(res.ID)
}

func TestGetForm(t *testing.T) {
	if err := setupEnv(); err != nil {
		t.Fatal(err)
	}

	opts := ClientOpts{
		Token: os.Getenv("TYPEFORM_API_TOKEN"),
	}
	if len(os.Getenv("TYPEFORM_VERBOSE")) > 0 {
		opts.HTTPClient = clients.TracerHTTPClient()
	}

	//formID := "dezzD5Cw"
	formID := "KMb3pAsq"

	res, err := NewClient(opts).GetForm(context.TODO(), formID)
	if err != nil {
		var reqErrs *Errors
		if !errors.As(err, &reqErrs) {
			t.Fatal(err)
		}

		if reqErrs.Code != "FORM_NOT_FOUND" {
			t.Fatal(reqErrs)
		}
	}

	if res != nil {
		t.Logf("Display link: %s", res.Links.Display)

		s, _ := json.MarshalIndent(res, "", " ")
		t.Logf("\n%s\n", s)
	}
}

func TestDeleteForm(t *testing.T) {
	if err := setupEnv(); err != nil {
		t.Fatal(err)
	}

	opts := ClientOpts{
		Token: os.Getenv("TYPEFORM_API_TOKEN"),
	}
	if len(os.Getenv("TYPEFORM_VERBOSE")) > 0 {
		opts.HTTPClient = clients.TracerHTTPClient()
	}

	formID := "kvE1U8Mw"

	err := NewClient(opts).DeleteForm(context.TODO(), formID)
	if err != nil {
		t.Fatal(err)
	}
}

func setupEnv() error {
	envMap, err := dotenv.FromFile("../../../.env")
	if err != nil {
		return err
	}
	dotenv.PutInEnv(envMap, false)
	return nil
}
