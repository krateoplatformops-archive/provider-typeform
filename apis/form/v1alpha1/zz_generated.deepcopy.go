//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 KrateoPlatformOps.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Attachment) DeepCopyInto(out *Attachment) {
	*out = *in
	if in.Scale != nil {
		in, out := &in.Scale, &out.Scale
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Attachment.
func (in *Attachment) DeepCopy() *Attachment {
	if in == nil {
		return nil
	}
	out := new(Attachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Field) DeepCopyInto(out *Field) {
	*out = *in
	in.Properties.DeepCopyInto(&out.Properties)
	if in.Validations != nil {
		in, out := &in.Validations, &out.Validations
		*out = new(Validations)
		(*in).DeepCopyInto(*out)
	}
	if in.Layout != nil {
		in, out := &in.Layout, &out.Layout
		*out = new(Layout)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Field.
func (in *Field) DeepCopy() *Field {
	if in == nil {
		return nil
	}
	out := new(Field)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Form) DeepCopyInto(out *Form) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Form.
func (in *Form) DeepCopy() *Form {
	if in == nil {
		return nil
	}
	out := new(Form)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Form) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FormList) DeepCopyInto(out *FormList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Form, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FormList.
func (in *FormList) DeepCopy() *FormList {
	if in == nil {
		return nil
	}
	out := new(FormList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FormList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FormObservation) DeepCopyInto(out *FormObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FormObservation.
func (in *FormObservation) DeepCopy() *FormObservation {
	if in == nil {
		return nil
	}
	out := new(FormObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FormParams) DeepCopyInto(out *FormParams) {
	*out = *in
	if in.Fields != nil {
		in, out := &in.Fields, &out.Fields
		*out = make([]Field, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WelcomeScreens != nil {
		in, out := &in.WelcomeScreens, &out.WelcomeScreens
		*out = make([]WelcomeScreen, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ThankyouScreens != nil {
		in, out := &in.ThankyouScreens, &out.ThankyouScreens
		*out = make([]ThankyouScreen, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FormParams.
func (in *FormParams) DeepCopy() *FormParams {
	if in == nil {
		return nil
	}
	out := new(FormParams)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FormSpec) DeepCopyInto(out *FormSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FormSpec.
func (in *FormSpec) DeepCopy() *FormSpec {
	if in == nil {
		return nil
	}
	out := new(FormSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FormStatus) DeepCopyInto(out *FormStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FormStatus.
func (in *FormStatus) DeepCopy() *FormStatus {
	if in == nil {
		return nil
	}
	out := new(FormStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Layout) DeepCopyInto(out *Layout) {
	*out = *in
	in.Attachment.DeepCopyInto(&out.Attachment)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Layout.
func (in *Layout) DeepCopy() *Layout {
	if in == nil {
		return nil
	}
	out := new(Layout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Properties) DeepCopyInto(out *Properties) {
	*out = *in
	if in.AllowMultipleSelection != nil {
		in, out := &in.AllowMultipleSelection, &out.AllowMultipleSelection
		*out = new(bool)
		**out = **in
	}
	if in.AllowOtherChoice != nil {
		in, out := &in.AllowOtherChoice, &out.AllowOtherChoice
		*out = new(bool)
		**out = **in
	}
	if in.AlphabeticalOrder != nil {
		in, out := &in.AlphabeticalOrder, &out.AlphabeticalOrder
		*out = new(bool)
		**out = **in
	}
	if in.Choices != nil {
		in, out := &in.Choices, &out.Choices
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = new(int)
		**out = **in
	}
	if in.Shape != nil {
		in, out := &in.Shape, &out.Shape
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Properties.
func (in *Properties) DeepCopy() *Properties {
	if in == nil {
		return nil
	}
	out := new(Properties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThankYouScreenProperties) DeepCopyInto(out *ThankYouScreenProperties) {
	*out = *in
	if in.ShareIcons != nil {
		in, out := &in.ShareIcons, &out.ShareIcons
		*out = new(bool)
		**out = **in
	}
	if in.ShowButton != nil {
		in, out := &in.ShowButton, &out.ShowButton
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThankYouScreenProperties.
func (in *ThankYouScreenProperties) DeepCopy() *ThankYouScreenProperties {
	if in == nil {
		return nil
	}
	out := new(ThankYouScreenProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThankyouScreen) DeepCopyInto(out *ThankyouScreen) {
	*out = *in
	if in.Attachment != nil {
		in, out := &in.Attachment, &out.Attachment
		*out = new(Attachment)
		(*in).DeepCopyInto(*out)
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(ThankYouScreenProperties)
		(*in).DeepCopyInto(*out)
	}
	if in.Layout != nil {
		in, out := &in.Layout, &out.Layout
		*out = new(Layout)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThankyouScreen.
func (in *ThankyouScreen) DeepCopy() *ThankyouScreen {
	if in == nil {
		return nil
	}
	out := new(ThankyouScreen)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Validations) DeepCopyInto(out *Validations) {
	*out = *in
	if in.MaxLength != nil {
		in, out := &in.MaxLength, &out.MaxLength
		*out = new(int)
		**out = **in
	}
	if in.MaxSelection != nil {
		in, out := &in.MaxSelection, &out.MaxSelection
		*out = new(int)
		**out = **in
	}
	if in.MinSelection != nil {
		in, out := &in.MinSelection, &out.MinSelection
		*out = new(int)
		**out = **in
	}
	if in.Required != nil {
		in, out := &in.Required, &out.Required
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Validations.
func (in *Validations) DeepCopy() *Validations {
	if in == nil {
		return nil
	}
	out := new(Validations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WelcomeScreen) DeepCopyInto(out *WelcomeScreen) {
	*out = *in
	if in.Layout != nil {
		in, out := &in.Layout, &out.Layout
		*out = new(Layout)
		(*in).DeepCopyInto(*out)
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(WelcomeScreenProperties)
		(*in).DeepCopyInto(*out)
	}
	if in.Attachment != nil {
		in, out := &in.Attachment, &out.Attachment
		*out = new(Attachment)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WelcomeScreen.
func (in *WelcomeScreen) DeepCopy() *WelcomeScreen {
	if in == nil {
		return nil
	}
	out := new(WelcomeScreen)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WelcomeScreenProperties) DeepCopyInto(out *WelcomeScreenProperties) {
	*out = *in
	if in.ShowButton != nil {
		in, out := &in.ShowButton, &out.ShowButton
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WelcomeScreenProperties.
func (in *WelcomeScreenProperties) DeepCopy() *WelcomeScreenProperties {
	if in == nil {
		return nil
	}
	out := new(WelcomeScreenProperties)
	in.DeepCopyInto(out)
	return out
}
