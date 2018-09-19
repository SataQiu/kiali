package kubernetes

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// stackdriver is one of the specific types defined by Istio as a Kubernetes extension.
// Preliminar rule is to define one file per type.
// types.go will collect common/shared types.
// This type is extracted from Istio Pilot as models used are not public and it doesn't make sense to fetch all
// Istio project as a dependency.
// Reference: https://github.com/istio/istio/blob/master/pilot/pkg/config/kube/crd/types.go

const stackdrivers = "stackdrivers"
const stackdriverType = "stackdriver"
const stackdriverLabel = "stackdriver"

// stackdriver is the generic Kubernetes API object wrapper
// stackdriver starts with lowercase as it maps a "kind":"stackdriver" Istio response.
type stackdriver struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata"`
	Spec               map[string]interface{} `json:"spec"`
}

// stackdriverList is the generic Kubernetes API list wrapper
// stackdriverList starts with lowercase as it maps a "kind":"stackdriverList" Istio response.
type stackdriverList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`
	Items            []stackdriver `json:"items"`
}

// GetSpec from a wrapper
func (in *stackdriver) GetSpec() map[string]interface{} {
	return in.Spec
}

// SetSpec for a wrapper
func (in *stackdriver) SetSpec(spec map[string]interface{}) {
	in.Spec = spec
}

// GetObjectMeta from a wrapper
func (in *stackdriver) GetObjectMeta() meta_v1.ObjectMeta {
	return in.ObjectMeta
}

// SetObjectMeta for a wrapper
func (in *stackdriver) SetObjectMeta(metadata meta_v1.ObjectMeta) {
	in.ObjectMeta = metadata
}

// GetItems from a wrapper
func (in *stackdriverList) GetItems() []IstioObject {
	out := make([]IstioObject, len(in.Items))
	for i := range in.Items {
		out[i] = &in.Items[i]
	}
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *stackdriver) DeepCopyInto(out *stackdriver) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new stackdriver.
func (in *stackdriver) DeepCopy() *stackdriver {
	if in == nil {
		return nil
	}
	out := new(stackdriver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *stackdriver) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}

	return nil
}

// DeepCopyIstioObject is an autogenerated deepcopy function, copying the receiver, creating a new IstioObject.
func (in *stackdriver) DeepCopyIstioObject() IstioObject {
	if c := in.DeepCopy(); c != nil {
		return c
	}

	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *stackdriverList) DeepCopyInto(out *stackdriverList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]stackdriver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleList.
func (in *stackdriverList) DeepCopy() *stackdriverList {
	if in == nil {
		return nil
	}
	out := new(stackdriverList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *stackdriverList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}

	return nil
}