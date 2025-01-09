// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	securityv1 "github.com/openshift/api/security/v1"
	internal "github.com/openshift/client-go/security/applyconfigurations/internal"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// RangeAllocationApplyConfiguration represents an declarative configuration of the RangeAllocation type for use
// with apply.
type RangeAllocationApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Range                            *string `json:"range,omitempty"`
	Data                             []byte  `json:"data,omitempty"`
}

// RangeAllocation constructs an declarative configuration of the RangeAllocation type for use with
// apply.
func RangeAllocation(name string) *RangeAllocationApplyConfiguration {
	b := &RangeAllocationApplyConfiguration{}
	b.WithName(name)
	b.WithKind("RangeAllocation")
	b.WithAPIVersion("security.openshift.io/v1")
	return b
}

// ExtractRangeAllocation extracts the applied configuration owned by fieldManager from
// rangeAllocation. If no managedFields are found in rangeAllocation for fieldManager, a
// RangeAllocationApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// rangeAllocation must be a unmodified RangeAllocation API object that was retrieved from the Kubernetes API.
// ExtractRangeAllocation provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractRangeAllocation(rangeAllocation *securityv1.RangeAllocation, fieldManager string) (*RangeAllocationApplyConfiguration, error) {
	return extractRangeAllocation(rangeAllocation, fieldManager, "")
}

// ExtractRangeAllocationStatus is the same as ExtractRangeAllocation except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractRangeAllocationStatus(rangeAllocation *securityv1.RangeAllocation, fieldManager string) (*RangeAllocationApplyConfiguration, error) {
	return extractRangeAllocation(rangeAllocation, fieldManager, "status")
}

func extractRangeAllocation(rangeAllocation *securityv1.RangeAllocation, fieldManager string, subresource string) (*RangeAllocationApplyConfiguration, error) {
	b := &RangeAllocationApplyConfiguration{}
	err := managedfields.ExtractInto(rangeAllocation, internal.Parser().Type("com.github.openshift.api.security.v1.RangeAllocation"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(rangeAllocation.Name)

	b.WithKind("RangeAllocation")
	b.WithAPIVersion("security.openshift.io/v1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *RangeAllocationApplyConfiguration) WithKind(value string) *RangeAllocationApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *RangeAllocationApplyConfiguration) WithAPIVersion(value string) *RangeAllocationApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *RangeAllocationApplyConfiguration) WithName(value string) *RangeAllocationApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *RangeAllocationApplyConfiguration) WithGenerateName(value string) *RangeAllocationApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *RangeAllocationApplyConfiguration) WithNamespace(value string) *RangeAllocationApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *RangeAllocationApplyConfiguration) WithUID(value types.UID) *RangeAllocationApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *RangeAllocationApplyConfiguration) WithResourceVersion(value string) *RangeAllocationApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *RangeAllocationApplyConfiguration) WithGeneration(value int64) *RangeAllocationApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *RangeAllocationApplyConfiguration) WithCreationTimestamp(value metav1.Time) *RangeAllocationApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *RangeAllocationApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *RangeAllocationApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *RangeAllocationApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *RangeAllocationApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *RangeAllocationApplyConfiguration) WithLabels(entries map[string]string) *RangeAllocationApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *RangeAllocationApplyConfiguration) WithAnnotations(entries map[string]string) *RangeAllocationApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *RangeAllocationApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *RangeAllocationApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *RangeAllocationApplyConfiguration) WithFinalizers(values ...string) *RangeAllocationApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *RangeAllocationApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithRange sets the Range field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Range field is set to the value of the last call.
func (b *RangeAllocationApplyConfiguration) WithRange(value string) *RangeAllocationApplyConfiguration {
	b.Range = &value
	return b
}

// WithData adds the given value to the Data field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Data field.
func (b *RangeAllocationApplyConfiguration) WithData(values ...byte) *RangeAllocationApplyConfiguration {
	for i := range values {
		b.Data = append(b.Data, values[i])
	}
	return b
}
