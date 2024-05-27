//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 Loft Labs Inc.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	v1 "k8s.io/api/admissionregistration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Conditions) DeepCopyInto(out *Conditions) {
	{
		in := &in
		*out = make(Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Conditions.
func (in Conditions) DeepCopy() Conditions {
	if in == nil {
		return nil
	}
	out := new(Conditions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JsPolicy) DeepCopyInto(out *JsPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JsPolicy.
func (in *JsPolicy) DeepCopy() *JsPolicy {
	if in == nil {
		return nil
	}
	out := new(JsPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JsPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JsPolicyBundle) DeepCopyInto(out *JsPolicyBundle) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JsPolicyBundle.
func (in *JsPolicyBundle) DeepCopy() *JsPolicyBundle {
	if in == nil {
		return nil
	}
	out := new(JsPolicyBundle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JsPolicyBundle) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JsPolicyBundleList) DeepCopyInto(out *JsPolicyBundleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JsPolicyBundle, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JsPolicyBundleList.
func (in *JsPolicyBundleList) DeepCopy() *JsPolicyBundleList {
	if in == nil {
		return nil
	}
	out := new(JsPolicyBundleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JsPolicyBundleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JsPolicyBundleSpec) DeepCopyInto(out *JsPolicyBundleSpec) {
	*out = *in
	if in.Bundle != nil {
		in, out := &in.Bundle, &out.Bundle
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JsPolicyBundleSpec.
func (in *JsPolicyBundleSpec) DeepCopy() *JsPolicyBundleSpec {
	if in == nil {
		return nil
	}
	out := new(JsPolicyBundleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JsPolicyBundleStatus) DeepCopyInto(out *JsPolicyBundleStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JsPolicyBundleStatus.
func (in *JsPolicyBundleStatus) DeepCopy() *JsPolicyBundleStatus {
	if in == nil {
		return nil
	}
	out := new(JsPolicyBundleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JsPolicyList) DeepCopyInto(out *JsPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JsPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JsPolicyList.
func (in *JsPolicyList) DeepCopy() *JsPolicyList {
	if in == nil {
		return nil
	}
	out := new(JsPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JsPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JsPolicySpec) DeepCopyInto(out *JsPolicySpec) {
	*out = *in
	if in.Operations != nil {
		in, out := &in.Operations, &out.Operations
		*out = make([]v1.OperationType, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.APIGroups != nil {
		in, out := &in.APIGroups, &out.APIGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.APIVersions != nil {
		in, out := &in.APIVersions, &out.APIVersions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(v1.ScopeType)
		**out = **in
	}
	if in.FailurePolicy != nil {
		in, out := &in.FailurePolicy, &out.FailurePolicy
		*out = new(v1.FailurePolicyType)
		**out = **in
	}
	if in.MatchPolicy != nil {
		in, out := &in.MatchPolicy, &out.MatchPolicy
		*out = new(v1.MatchPolicyType)
		**out = **in
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ObjectSelector != nil {
		in, out := &in.ObjectSelector, &out.ObjectSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.TimeoutSeconds != nil {
		in, out := &in.TimeoutSeconds, &out.TimeoutSeconds
		*out = new(int32)
		**out = **in
	}
	if in.ViolationPolicy != nil {
		in, out := &in.ViolationPolicy, &out.ViolationPolicy
		*out = new(ViolationPolicyType)
		**out = **in
	}
	if in.AuditPolicy != nil {
		in, out := &in.AuditPolicy, &out.AuditPolicy
		*out = new(AuditPolicyType)
		**out = **in
	}
	if in.AuditLogSize != nil {
		in, out := &in.AuditLogSize, &out.AuditLogSize
		*out = new(int32)
		**out = **in
	}
	if in.Dependencies != nil {
		in, out := &in.Dependencies, &out.Dependencies
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JsPolicySpec.
func (in *JsPolicySpec) DeepCopy() *JsPolicySpec {
	if in == nil {
		return nil
	}
	out := new(JsPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JsPolicyStatus) DeepCopyInto(out *JsPolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JsPolicyStatus.
func (in *JsPolicyStatus) DeepCopy() *JsPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(JsPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JsPolicyViolations) DeepCopyInto(out *JsPolicyViolations) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JsPolicyViolations.
func (in *JsPolicyViolations) DeepCopy() *JsPolicyViolations {
	if in == nil {
		return nil
	}
	out := new(JsPolicyViolations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JsPolicyViolations) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JsPolicyViolationsList) DeepCopyInto(out *JsPolicyViolationsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JsPolicyViolations, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JsPolicyViolationsList.
func (in *JsPolicyViolationsList) DeepCopy() *JsPolicyViolationsList {
	if in == nil {
		return nil
	}
	out := new(JsPolicyViolationsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JsPolicyViolationsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JsPolicyViolationsSpec) DeepCopyInto(out *JsPolicyViolationsSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JsPolicyViolationsSpec.
func (in *JsPolicyViolationsSpec) DeepCopy() *JsPolicyViolationsSpec {
	if in == nil {
		return nil
	}
	out := new(JsPolicyViolationsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JsPolicyViolationsStatus) DeepCopyInto(out *JsPolicyViolationsStatus) {
	*out = *in
	if in.Violations != nil {
		in, out := &in.Violations, &out.Violations
		*out = make([]PolicyViolation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JsPolicyViolationsStatus.
func (in *JsPolicyViolationsStatus) DeepCopy() *JsPolicyViolationsStatus {
	if in == nil {
		return nil
	}
	out := new(JsPolicyViolationsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyViolation) DeepCopyInto(out *PolicyViolation) {
	*out = *in
	if in.RequestInfo != nil {
		in, out := &in.RequestInfo, &out.RequestInfo
		*out = new(RequestInfo)
		**out = **in
	}
	if in.UserInfo != nil {
		in, out := &in.UserInfo, &out.UserInfo
		*out = new(UserInfo)
		**out = **in
	}
	in.Timestamp.DeepCopyInto(&out.Timestamp)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyViolation.
func (in *PolicyViolation) DeepCopy() *PolicyViolation {
	if in == nil {
		return nil
	}
	out := new(PolicyViolation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestInfo) DeepCopyInto(out *RequestInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestInfo.
func (in *RequestInfo) DeepCopy() *RequestInfo {
	if in == nil {
		return nil
	}
	out := new(RequestInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserInfo) DeepCopyInto(out *UserInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserInfo.
func (in *UserInfo) DeepCopy() *UserInfo {
	if in == nil {
		return nil
	}
	out := new(UserInfo)
	in.DeepCopyInto(out)
	return out
}