/*
Copyright 2022 The Kubernetes Authors.

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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1beta1 "sigs.k8s.io/kueue/apis/kueue/v1beta1"
	kueuev1beta1 "sigs.k8s.io/kueue/client-go/applyconfiguration/kueue/v1beta1"
)

// FakeAdmissionChecks implements AdmissionCheckInterface
type FakeAdmissionChecks struct {
	Fake *FakeKueueV1beta1
	ns   string
}

var admissionchecksResource = v1beta1.SchemeGroupVersion.WithResource("admissionchecks")

var admissionchecksKind = v1beta1.SchemeGroupVersion.WithKind("AdmissionCheck")

// Get takes name of the admissionCheck, and returns the corresponding admissionCheck object, and an error if there is any.
func (c *FakeAdmissionChecks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.AdmissionCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(admissionchecksResource, c.ns, name), &v1beta1.AdmissionCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AdmissionCheck), err
}

// List takes label and field selectors, and returns the list of AdmissionChecks that match those selectors.
func (c *FakeAdmissionChecks) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.AdmissionCheckList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(admissionchecksResource, admissionchecksKind, c.ns, opts), &v1beta1.AdmissionCheckList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.AdmissionCheckList{ListMeta: obj.(*v1beta1.AdmissionCheckList).ListMeta}
	for _, item := range obj.(*v1beta1.AdmissionCheckList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested admissionChecks.
func (c *FakeAdmissionChecks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(admissionchecksResource, c.ns, opts))

}

// Create takes the representation of a admissionCheck and creates it.  Returns the server's representation of the admissionCheck, and an error, if there is any.
func (c *FakeAdmissionChecks) Create(ctx context.Context, admissionCheck *v1beta1.AdmissionCheck, opts v1.CreateOptions) (result *v1beta1.AdmissionCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(admissionchecksResource, c.ns, admissionCheck), &v1beta1.AdmissionCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AdmissionCheck), err
}

// Update takes the representation of a admissionCheck and updates it. Returns the server's representation of the admissionCheck, and an error, if there is any.
func (c *FakeAdmissionChecks) Update(ctx context.Context, admissionCheck *v1beta1.AdmissionCheck, opts v1.UpdateOptions) (result *v1beta1.AdmissionCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(admissionchecksResource, c.ns, admissionCheck), &v1beta1.AdmissionCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AdmissionCheck), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAdmissionChecks) UpdateStatus(ctx context.Context, admissionCheck *v1beta1.AdmissionCheck, opts v1.UpdateOptions) (*v1beta1.AdmissionCheck, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(admissionchecksResource, "status", c.ns, admissionCheck), &v1beta1.AdmissionCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AdmissionCheck), err
}

// Delete takes name of the admissionCheck and deletes it. Returns an error if one occurs.
func (c *FakeAdmissionChecks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(admissionchecksResource, c.ns, name, opts), &v1beta1.AdmissionCheck{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAdmissionChecks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(admissionchecksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.AdmissionCheckList{})
	return err
}

// Patch applies the patch and returns the patched admissionCheck.
func (c *FakeAdmissionChecks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.AdmissionCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(admissionchecksResource, c.ns, name, pt, data, subresources...), &v1beta1.AdmissionCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AdmissionCheck), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied admissionCheck.
func (c *FakeAdmissionChecks) Apply(ctx context.Context, admissionCheck *kueuev1beta1.AdmissionCheckApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.AdmissionCheck, err error) {
	if admissionCheck == nil {
		return nil, fmt.Errorf("admissionCheck provided to Apply must not be nil")
	}
	data, err := json.Marshal(admissionCheck)
	if err != nil {
		return nil, err
	}
	name := admissionCheck.Name
	if name == nil {
		return nil, fmt.Errorf("admissionCheck.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(admissionchecksResource, c.ns, *name, types.ApplyPatchType, data), &v1beta1.AdmissionCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AdmissionCheck), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeAdmissionChecks) ApplyStatus(ctx context.Context, admissionCheck *kueuev1beta1.AdmissionCheckApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.AdmissionCheck, err error) {
	if admissionCheck == nil {
		return nil, fmt.Errorf("admissionCheck provided to Apply must not be nil")
	}
	data, err := json.Marshal(admissionCheck)
	if err != nil {
		return nil, err
	}
	name := admissionCheck.Name
	if name == nil {
		return nil, fmt.Errorf("admissionCheck.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(admissionchecksResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1beta1.AdmissionCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AdmissionCheck), err
}