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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MobileSecurityService) DeepCopyInto(out *MobileSecurityService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MobileSecurityService.
func (in *MobileSecurityService) DeepCopy() *MobileSecurityService {
	if in == nil {
		return nil
	}
	out := new(MobileSecurityService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MobileSecurityService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MobileSecurityServiceBind) DeepCopyInto(out *MobileSecurityServiceBind) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MobileSecurityServiceBind.
func (in *MobileSecurityServiceBind) DeepCopy() *MobileSecurityServiceBind {
	if in == nil {
		return nil
	}
	out := new(MobileSecurityServiceBind)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MobileSecurityServiceBind) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MobileSecurityServiceBindList) DeepCopyInto(out *MobileSecurityServiceBindList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MobileSecurityServiceBind, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MobileSecurityServiceBindList.
func (in *MobileSecurityServiceBindList) DeepCopy() *MobileSecurityServiceBindList {
	if in == nil {
		return nil
	}
	out := new(MobileSecurityServiceBindList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MobileSecurityServiceBindList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MobileSecurityServiceBindSpec) DeepCopyInto(out *MobileSecurityServiceBindSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MobileSecurityServiceBindSpec.
func (in *MobileSecurityServiceBindSpec) DeepCopy() *MobileSecurityServiceBindSpec {
	if in == nil {
		return nil
	}
	out := new(MobileSecurityServiceBindSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MobileSecurityServiceBindStatus) DeepCopyInto(out *MobileSecurityServiceBindStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MobileSecurityServiceBindStatus.
func (in *MobileSecurityServiceBindStatus) DeepCopy() *MobileSecurityServiceBindStatus {
	if in == nil {
		return nil
	}
	out := new(MobileSecurityServiceBindStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MobileSecurityServiceDB) DeepCopyInto(out *MobileSecurityServiceDB) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MobileSecurityServiceDB.
func (in *MobileSecurityServiceDB) DeepCopy() *MobileSecurityServiceDB {
	if in == nil {
		return nil
	}
	out := new(MobileSecurityServiceDB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MobileSecurityServiceDB) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MobileSecurityServiceDBList) DeepCopyInto(out *MobileSecurityServiceDBList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MobileSecurityServiceDB, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MobileSecurityServiceDBList.
func (in *MobileSecurityServiceDBList) DeepCopy() *MobileSecurityServiceDBList {
	if in == nil {
		return nil
	}
	out := new(MobileSecurityServiceDBList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MobileSecurityServiceDBList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MobileSecurityServiceDBSpec) DeepCopyInto(out *MobileSecurityServiceDBSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MobileSecurityServiceDBSpec.
func (in *MobileSecurityServiceDBSpec) DeepCopy() *MobileSecurityServiceDBSpec {
	if in == nil {
		return nil
	}
	out := new(MobileSecurityServiceDBSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MobileSecurityServiceDBStatus) DeepCopyInto(out *MobileSecurityServiceDBStatus) {
	*out = *in
	in.DeploymentStatus.DeepCopyInto(&out.DeploymentStatus)
	in.ServiceStatus.DeepCopyInto(&out.ServiceStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MobileSecurityServiceDBStatus.
func (in *MobileSecurityServiceDBStatus) DeepCopy() *MobileSecurityServiceDBStatus {
	if in == nil {
		return nil
	}
	out := new(MobileSecurityServiceDBStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MobileSecurityServiceList) DeepCopyInto(out *MobileSecurityServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MobileSecurityService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MobileSecurityServiceList.
func (in *MobileSecurityServiceList) DeepCopy() *MobileSecurityServiceList {
	if in == nil {
		return nil
	}
	out := new(MobileSecurityServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MobileSecurityServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MobileSecurityServiceSpec) DeepCopyInto(out *MobileSecurityServiceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MobileSecurityServiceSpec.
func (in *MobileSecurityServiceSpec) DeepCopy() *MobileSecurityServiceSpec {
	if in == nil {
		return nil
	}
	out := new(MobileSecurityServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MobileSecurityServiceStatus) DeepCopyInto(out *MobileSecurityServiceStatus) {
	*out = *in
	in.DeploymentStatus.DeepCopyInto(&out.DeploymentStatus)
	in.ServiceStatus.DeepCopyInto(&out.ServiceStatus)
	in.IngressStatus.DeepCopyInto(&out.IngressStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MobileSecurityServiceStatus.
func (in *MobileSecurityServiceStatus) DeepCopy() *MobileSecurityServiceStatus {
	if in == nil {
		return nil
	}
	out := new(MobileSecurityServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MobileSecurityServiceUnbind) DeepCopyInto(out *MobileSecurityServiceUnbind) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MobileSecurityServiceUnbind.
func (in *MobileSecurityServiceUnbind) DeepCopy() *MobileSecurityServiceUnbind {
	if in == nil {
		return nil
	}
	out := new(MobileSecurityServiceUnbind)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MobileSecurityServiceUnbind) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MobileSecurityServiceUnbindList) DeepCopyInto(out *MobileSecurityServiceUnbindList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MobileSecurityServiceUnbind, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MobileSecurityServiceUnbindList.
func (in *MobileSecurityServiceUnbindList) DeepCopy() *MobileSecurityServiceUnbindList {
	if in == nil {
		return nil
	}
	out := new(MobileSecurityServiceUnbindList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MobileSecurityServiceUnbindList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MobileSecurityServiceUnbindSpec) DeepCopyInto(out *MobileSecurityServiceUnbindSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MobileSecurityServiceUnbindSpec.
func (in *MobileSecurityServiceUnbindSpec) DeepCopy() *MobileSecurityServiceUnbindSpec {
	if in == nil {
		return nil
	}
	out := new(MobileSecurityServiceUnbindSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MobileSecurityServiceUnbindStatus) DeepCopyInto(out *MobileSecurityServiceUnbindStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MobileSecurityServiceUnbindStatus.
func (in *MobileSecurityServiceUnbindStatus) DeepCopy() *MobileSecurityServiceUnbindStatus {
	if in == nil {
		return nil
	}
	out := new(MobileSecurityServiceUnbindStatus)
	in.DeepCopyInto(out)
	return out
}
