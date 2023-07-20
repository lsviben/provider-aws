//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Activity) DeepCopyInto(out *Activity) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Activity.
func (in *Activity) DeepCopy() *Activity {
	if in == nil {
		return nil
	}
	out := new(Activity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Activity) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivityInitParameters) DeepCopyInto(out *ActivityInitParameters) {
	*out = *in
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivityInitParameters.
func (in *ActivityInitParameters) DeepCopy() *ActivityInitParameters {
	if in == nil {
		return nil
	}
	out := new(ActivityInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivityList) DeepCopyInto(out *ActivityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Activity, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivityList.
func (in *ActivityList) DeepCopy() *ActivityList {
	if in == nil {
		return nil
	}
	out := new(ActivityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActivityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivityObservation) DeepCopyInto(out *ActivityObservation) {
	*out = *in
	if in.CreationDate != nil {
		in, out := &in.CreationDate, &out.CreationDate
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivityObservation.
func (in *ActivityObservation) DeepCopy() *ActivityObservation {
	if in == nil {
		return nil
	}
	out := new(ActivityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivityParameters) DeepCopyInto(out *ActivityParameters) {
	*out = *in
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivityParameters.
func (in *ActivityParameters) DeepCopy() *ActivityParameters {
	if in == nil {
		return nil
	}
	out := new(ActivityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivitySpec) DeepCopyInto(out *ActivitySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivitySpec.
func (in *ActivitySpec) DeepCopy() *ActivitySpec {
	if in == nil {
		return nil
	}
	out := new(ActivitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivityStatus) DeepCopyInto(out *ActivityStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivityStatus.
func (in *ActivityStatus) DeepCopy() *ActivityStatus {
	if in == nil {
		return nil
	}
	out := new(ActivityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingConfigurationInitParameters) DeepCopyInto(out *LoggingConfigurationInitParameters) {
	*out = *in
	if in.IncludeExecutionData != nil {
		in, out := &in.IncludeExecutionData, &out.IncludeExecutionData
		*out = new(bool)
		**out = **in
	}
	if in.Level != nil {
		in, out := &in.Level, &out.Level
		*out = new(string)
		**out = **in
	}
	if in.LogDestination != nil {
		in, out := &in.LogDestination, &out.LogDestination
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingConfigurationInitParameters.
func (in *LoggingConfigurationInitParameters) DeepCopy() *LoggingConfigurationInitParameters {
	if in == nil {
		return nil
	}
	out := new(LoggingConfigurationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingConfigurationObservation) DeepCopyInto(out *LoggingConfigurationObservation) {
	*out = *in
	if in.IncludeExecutionData != nil {
		in, out := &in.IncludeExecutionData, &out.IncludeExecutionData
		*out = new(bool)
		**out = **in
	}
	if in.Level != nil {
		in, out := &in.Level, &out.Level
		*out = new(string)
		**out = **in
	}
	if in.LogDestination != nil {
		in, out := &in.LogDestination, &out.LogDestination
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingConfigurationObservation.
func (in *LoggingConfigurationObservation) DeepCopy() *LoggingConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(LoggingConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingConfigurationParameters) DeepCopyInto(out *LoggingConfigurationParameters) {
	*out = *in
	if in.IncludeExecutionData != nil {
		in, out := &in.IncludeExecutionData, &out.IncludeExecutionData
		*out = new(bool)
		**out = **in
	}
	if in.Level != nil {
		in, out := &in.Level, &out.Level
		*out = new(string)
		**out = **in
	}
	if in.LogDestination != nil {
		in, out := &in.LogDestination, &out.LogDestination
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingConfigurationParameters.
func (in *LoggingConfigurationParameters) DeepCopy() *LoggingConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(LoggingConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateMachine) DeepCopyInto(out *StateMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateMachine.
func (in *StateMachine) DeepCopy() *StateMachine {
	if in == nil {
		return nil
	}
	out := new(StateMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StateMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateMachineInitParameters) DeepCopyInto(out *StateMachineInitParameters) {
	*out = *in
	if in.Definition != nil {
		in, out := &in.Definition, &out.Definition
		*out = new(string)
		**out = **in
	}
	if in.LoggingConfiguration != nil {
		in, out := &in.LoggingConfiguration, &out.LoggingConfiguration
		*out = make([]LoggingConfigurationInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RoleArn != nil {
		in, out := &in.RoleArn, &out.RoleArn
		*out = new(string)
		**out = **in
	}
	if in.RoleArnRef != nil {
		in, out := &in.RoleArnRef, &out.RoleArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RoleArnSelector != nil {
		in, out := &in.RoleArnSelector, &out.RoleArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TracingConfiguration != nil {
		in, out := &in.TracingConfiguration, &out.TracingConfiguration
		*out = make([]TracingConfigurationInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateMachineInitParameters.
func (in *StateMachineInitParameters) DeepCopy() *StateMachineInitParameters {
	if in == nil {
		return nil
	}
	out := new(StateMachineInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateMachineList) DeepCopyInto(out *StateMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StateMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateMachineList.
func (in *StateMachineList) DeepCopy() *StateMachineList {
	if in == nil {
		return nil
	}
	out := new(StateMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StateMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateMachineObservation) DeepCopyInto(out *StateMachineObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.CreationDate != nil {
		in, out := &in.CreationDate, &out.CreationDate
		*out = new(string)
		**out = **in
	}
	if in.Definition != nil {
		in, out := &in.Definition, &out.Definition
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LoggingConfiguration != nil {
		in, out := &in.LoggingConfiguration, &out.LoggingConfiguration
		*out = make([]LoggingConfigurationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RoleArn != nil {
		in, out := &in.RoleArn, &out.RoleArn
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TracingConfiguration != nil {
		in, out := &in.TracingConfiguration, &out.TracingConfiguration
		*out = make([]TracingConfigurationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateMachineObservation.
func (in *StateMachineObservation) DeepCopy() *StateMachineObservation {
	if in == nil {
		return nil
	}
	out := new(StateMachineObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateMachineParameters) DeepCopyInto(out *StateMachineParameters) {
	*out = *in
	if in.Definition != nil {
		in, out := &in.Definition, &out.Definition
		*out = new(string)
		**out = **in
	}
	if in.LoggingConfiguration != nil {
		in, out := &in.LoggingConfiguration, &out.LoggingConfiguration
		*out = make([]LoggingConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RoleArn != nil {
		in, out := &in.RoleArn, &out.RoleArn
		*out = new(string)
		**out = **in
	}
	if in.RoleArnRef != nil {
		in, out := &in.RoleArnRef, &out.RoleArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RoleArnSelector != nil {
		in, out := &in.RoleArnSelector, &out.RoleArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TracingConfiguration != nil {
		in, out := &in.TracingConfiguration, &out.TracingConfiguration
		*out = make([]TracingConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateMachineParameters.
func (in *StateMachineParameters) DeepCopy() *StateMachineParameters {
	if in == nil {
		return nil
	}
	out := new(StateMachineParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateMachineSpec) DeepCopyInto(out *StateMachineSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateMachineSpec.
func (in *StateMachineSpec) DeepCopy() *StateMachineSpec {
	if in == nil {
		return nil
	}
	out := new(StateMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateMachineStatus) DeepCopyInto(out *StateMachineStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateMachineStatus.
func (in *StateMachineStatus) DeepCopy() *StateMachineStatus {
	if in == nil {
		return nil
	}
	out := new(StateMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracingConfigurationInitParameters) DeepCopyInto(out *TracingConfigurationInitParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracingConfigurationInitParameters.
func (in *TracingConfigurationInitParameters) DeepCopy() *TracingConfigurationInitParameters {
	if in == nil {
		return nil
	}
	out := new(TracingConfigurationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracingConfigurationObservation) DeepCopyInto(out *TracingConfigurationObservation) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracingConfigurationObservation.
func (in *TracingConfigurationObservation) DeepCopy() *TracingConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(TracingConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracingConfigurationParameters) DeepCopyInto(out *TracingConfigurationParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracingConfigurationParameters.
func (in *TracingConfigurationParameters) DeepCopy() *TracingConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(TracingConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}
