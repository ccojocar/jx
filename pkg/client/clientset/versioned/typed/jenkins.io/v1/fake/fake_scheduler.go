// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	jenkinsiov1 "github.com/jenkins-x/jx/v2/pkg/apis/jenkins.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSchedulers implements SchedulerInterface
type FakeSchedulers struct {
	Fake *FakeJenkinsV1
	ns   string
}

var schedulersResource = schema.GroupVersionResource{Group: "jenkins.io", Version: "v1", Resource: "schedulers"}

var schedulersKind = schema.GroupVersionKind{Group: "jenkins.io", Version: "v1", Kind: "Scheduler"}

// Get takes name of the scheduler, and returns the corresponding scheduler object, and an error if there is any.
func (c *FakeSchedulers) Get(name string, options v1.GetOptions) (result *jenkinsiov1.Scheduler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(schedulersResource, c.ns, name), &jenkinsiov1.Scheduler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.Scheduler), err
}

// List takes label and field selectors, and returns the list of Schedulers that match those selectors.
func (c *FakeSchedulers) List(opts v1.ListOptions) (result *jenkinsiov1.SchedulerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(schedulersResource, schedulersKind, c.ns, opts), &jenkinsiov1.SchedulerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &jenkinsiov1.SchedulerList{ListMeta: obj.(*jenkinsiov1.SchedulerList).ListMeta}
	for _, item := range obj.(*jenkinsiov1.SchedulerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested schedulers.
func (c *FakeSchedulers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(schedulersResource, c.ns, opts))

}

// Create takes the representation of a scheduler and creates it.  Returns the server's representation of the scheduler, and an error, if there is any.
func (c *FakeSchedulers) Create(scheduler *jenkinsiov1.Scheduler) (result *jenkinsiov1.Scheduler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(schedulersResource, c.ns, scheduler), &jenkinsiov1.Scheduler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.Scheduler), err
}

// Update takes the representation of a scheduler and updates it. Returns the server's representation of the scheduler, and an error, if there is any.
func (c *FakeSchedulers) Update(scheduler *jenkinsiov1.Scheduler) (result *jenkinsiov1.Scheduler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(schedulersResource, c.ns, scheduler), &jenkinsiov1.Scheduler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.Scheduler), err
}

// Delete takes name of the scheduler and deletes it. Returns an error if one occurs.
func (c *FakeSchedulers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(schedulersResource, c.ns, name), &jenkinsiov1.Scheduler{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSchedulers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(schedulersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &jenkinsiov1.SchedulerList{})
	return err
}

// Patch applies the patch and returns the patched scheduler.
func (c *FakeSchedulers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *jenkinsiov1.Scheduler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(schedulersResource, c.ns, name, data, subresources...), &jenkinsiov1.Scheduler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.Scheduler), err
}
