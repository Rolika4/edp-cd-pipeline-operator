//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDPipeline) DeepCopyInto(out *CDPipeline) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDPipeline.
func (in *CDPipeline) DeepCopy() *CDPipeline {
	if in == nil {
		return nil
	}
	out := new(CDPipeline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CDPipeline) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDPipelineList) DeepCopyInto(out *CDPipelineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CDPipeline, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDPipelineList.
func (in *CDPipelineList) DeepCopy() *CDPipelineList {
	if in == nil {
		return nil
	}
	out := new(CDPipelineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CDPipelineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDPipelineSpec) DeepCopyInto(out *CDPipelineSpec) {
	*out = *in
	if in.InputDockerStreams != nil {
		in, out := &in.InputDockerStreams, &out.InputDockerStreams
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Applications != nil {
		in, out := &in.Applications, &out.Applications
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ApplicationsToPromote != nil {
		in, out := &in.ApplicationsToPromote, &out.ApplicationsToPromote
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDPipelineSpec.
func (in *CDPipelineSpec) DeepCopy() *CDPipelineSpec {
	if in == nil {
		return nil
	}
	out := new(CDPipelineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDPipelineStatus) DeepCopyInto(out *CDPipelineStatus) {
	*out = *in
	in.LastTimeUpdated.DeepCopyInto(&out.LastTimeUpdated)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDPipelineStatus.
func (in *CDPipelineStatus) DeepCopy() *CDPipelineStatus {
	if in == nil {
		return nil
	}
	out := new(CDPipelineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Library) DeepCopyInto(out *Library) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Library.
func (in *Library) DeepCopy() *Library {
	if in == nil {
		return nil
	}
	out := new(Library)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QualityGate) DeepCopyInto(out *QualityGate) {
	*out = *in
	if in.AutotestName != nil {
		in, out := &in.AutotestName, &out.AutotestName
		*out = new(string)
		**out = **in
	}
	if in.BranchName != nil {
		in, out := &in.BranchName, &out.BranchName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QualityGate.
func (in *QualityGate) DeepCopy() *QualityGate {
	if in == nil {
		return nil
	}
	out := new(QualityGate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Source) DeepCopyInto(out *Source) {
	*out = *in
	out.Library = in.Library
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Source.
func (in *Source) DeepCopy() *Source {
	if in == nil {
		return nil
	}
	out := new(Source)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Stage) DeepCopyInto(out *Stage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Stage.
func (in *Stage) DeepCopy() *Stage {
	if in == nil {
		return nil
	}
	out := new(Stage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Stage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StageList) DeepCopyInto(out *StageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Stage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StageList.
func (in *StageList) DeepCopy() *StageList {
	if in == nil {
		return nil
	}
	out := new(StageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StageSpec) DeepCopyInto(out *StageSpec) {
	*out = *in
	if in.QualityGates != nil {
		in, out := &in.QualityGates, &out.QualityGates
		*out = make([]QualityGate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(Source)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StageSpec.
func (in *StageSpec) DeepCopy() *StageSpec {
	if in == nil {
		return nil
	}
	out := new(StageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StageStatus) DeepCopyInto(out *StageStatus) {
	*out = *in
	in.LastTimeUpdated.DeepCopyInto(&out.LastTimeUpdated)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StageStatus.
func (in *StageStatus) DeepCopy() *StageStatus {
	if in == nil {
		return nil
	}
	out := new(StageStatus)
	in.DeepCopyInto(out)
	return out
}
