// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

package v1alpha1

import (
	tkexv1alpha1 "github.com/Tencent/bk-bcs/bcs-k8s/kubernetes/common/bcs-hook/apis/tkex/v1alpha1"
	inplaceupdate "github.com/Tencent/bk-bcs/bcs-k8s/kubernetes/common/update/inplaceupdate"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryPause) DeepCopyInto(out *CanaryPause) {
	*out = *in
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryPause.
func (in *CanaryPause) DeepCopy() *CanaryPause {
	if in == nil {
		return nil
	}
	out := new(CanaryPause)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryStatus) DeepCopyInto(out *CanaryStatus) {
	*out = *in
	if in.PauseStartTime != nil {
		in, out := &in.PauseStartTime, &out.PauseStartTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryStatus.
func (in *CanaryStatus) DeepCopy() *CanaryStatus {
	if in == nil {
		return nil
	}
	out := new(CanaryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryStep) DeepCopyInto(out *CanaryStep) {
	*out = *in
	if in.Partition != nil {
		in, out := &in.Partition, &out.Partition
		*out = new(int32)
		**out = **in
	}
	if in.Pause != nil {
		in, out := &in.Pause, &out.Pause
		*out = new(CanaryPause)
		(*in).DeepCopyInto(*out)
	}
	if in.Hook != nil {
		in, out := &in.Hook, &out.Hook
		*out = new(tkexv1alpha1.HookStep)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryStep.
func (in *CanaryStep) DeepCopy() *CanaryStep {
	if in == nil {
		return nil
	}
	out := new(CanaryStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryStrategy) DeepCopyInto(out *CanaryStrategy) {
	*out = *in
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]CanaryStep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryStrategy.
func (in *CanaryStrategy) DeepCopy() *CanaryStrategy {
	if in == nil {
		return nil
	}
	out := new(CanaryStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameStatefulSet) DeepCopyInto(out *GameStatefulSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameStatefulSet.
func (in *GameStatefulSet) DeepCopy() *GameStatefulSet {
	if in == nil {
		return nil
	}
	out := new(GameStatefulSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GameStatefulSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameStatefulSetCondition) DeepCopyInto(out *GameStatefulSetCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameStatefulSetCondition.
func (in *GameStatefulSetCondition) DeepCopy() *GameStatefulSetCondition {
	if in == nil {
		return nil
	}
	out := new(GameStatefulSetCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameStatefulSetList) DeepCopyInto(out *GameStatefulSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GameStatefulSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameStatefulSetList.
func (in *GameStatefulSetList) DeepCopy() *GameStatefulSetList {
	if in == nil {
		return nil
	}
	out := new(GameStatefulSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GameStatefulSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameStatefulSetPreDeleteUpdateStrategy) DeepCopyInto(out *GameStatefulSetPreDeleteUpdateStrategy) {
	*out = *in
	if in.Hook != nil {
		in, out := &in.Hook, &out.Hook
		*out = new(tkexv1alpha1.HookStep)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameStatefulSetPreDeleteUpdateStrategy.
func (in *GameStatefulSetPreDeleteUpdateStrategy) DeepCopy() *GameStatefulSetPreDeleteUpdateStrategy {
	if in == nil {
		return nil
	}
	out := new(GameStatefulSetPreDeleteUpdateStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameStatefulSetPreInplaceUpdateStrategy) DeepCopyInto(out *GameStatefulSetPreInplaceUpdateStrategy) {
	*out = *in
	if in.Hook != nil {
		in, out := &in.Hook, &out.Hook
		*out = new(tkexv1alpha1.HookStep)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameStatefulSetPreInplaceUpdateStrategy.
func (in *GameStatefulSetPreInplaceUpdateStrategy) DeepCopy() *GameStatefulSetPreInplaceUpdateStrategy {
	if in == nil {
		return nil
	}
	out := new(GameStatefulSetPreInplaceUpdateStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameStatefulSetSpec) DeepCopyInto(out *GameStatefulSetSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.Template.DeepCopyInto(&out.Template)
	if in.VolumeClaimTemplates != nil {
		in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
		*out = make([]corev1.PersistentVolumeClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.UpdateStrategy.DeepCopyInto(&out.UpdateStrategy)
	in.PreDeleteUpdateStrategy.DeepCopyInto(&out.PreDeleteUpdateStrategy)
	in.PreInplaceUpdateStrategy.DeepCopyInto(&out.PreInplaceUpdateStrategy)
	if in.RevisionHistoryLimit != nil {
		in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameStatefulSetSpec.
func (in *GameStatefulSetSpec) DeepCopy() *GameStatefulSetSpec {
	if in == nil {
		return nil
	}
	out := new(GameStatefulSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameStatefulSetStatus) DeepCopyInto(out *GameStatefulSetStatus) {
	*out = *in
	if in.CollisionCount != nil {
		in, out := &in.CollisionCount, &out.CollisionCount
		*out = new(int32)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]GameStatefulSetCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(string)
		**out = **in
	}
	if in.PauseConditions != nil {
		in, out := &in.PauseConditions, &out.PauseConditions
		*out = make([]tkexv1alpha1.PauseCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CurrentStepIndex != nil {
		in, out := &in.CurrentStepIndex, &out.CurrentStepIndex
		*out = new(int32)
		**out = **in
	}
	in.Canary.DeepCopyInto(&out.Canary)
	if in.PreDeleteHookConditions != nil {
		in, out := &in.PreDeleteHookConditions, &out.PreDeleteHookConditions
		*out = make([]tkexv1alpha1.PreDeleteHookCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}

	if in.PreInplaceHookConditions != nil {
		in, out := &in.PreInplaceHookConditions, &out.PreInplaceHookConditions
		*out = make([]tkexv1alpha1.PreInplaceHookCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameStatefulSetStatus.
func (in *GameStatefulSetStatus) DeepCopy() *GameStatefulSetStatus {
	if in == nil {
		return nil
	}
	out := new(GameStatefulSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameStatefulSetUpdateStrategy) DeepCopyInto(out *GameStatefulSetUpdateStrategy) {
	*out = *in
	if in.RollingUpdate != nil {
		in, out := &in.RollingUpdate, &out.RollingUpdate
		*out = new(RollingUpdateStatefulSetStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.InPlaceUpdateStrategy != nil {
		in, out := &in.InPlaceUpdateStrategy, &out.InPlaceUpdateStrategy
		*out = new(inplaceupdate.InPlaceUpdateStrategy)
		**out = **in
	}
	if in.CanaryStrategy != nil {
		in, out := &in.CanaryStrategy, &out.CanaryStrategy
		*out = new(CanaryStrategy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameStatefulSetUpdateStrategy.
func (in *GameStatefulSetUpdateStrategy) DeepCopy() *GameStatefulSetUpdateStrategy {
	if in == nil {
		return nil
	}
	out := new(GameStatefulSetUpdateStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollingUpdateStatefulSetStrategy) DeepCopyInto(out *RollingUpdateStatefulSetStrategy) {
	*out = *in
	if in.Partition != nil {
		in, out := &in.Partition, &out.Partition
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollingUpdateStatefulSetStrategy.
func (in *RollingUpdateStatefulSetStrategy) DeepCopy() *RollingUpdateStatefulSetStrategy {
	if in == nil {
		return nil
	}
	out := new(RollingUpdateStatefulSetStrategy)
	in.DeepCopyInto(out)
	return out
}
