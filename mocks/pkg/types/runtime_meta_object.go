// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	pkgtypes "k8s.io/apimachinery/pkg/types"

	runtime "k8s.io/apimachinery/pkg/runtime"

	schema "k8s.io/apimachinery/pkg/runtime/schema"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RuntimeMetaObject is an autogenerated mock type for the RuntimeMetaObject type
type RuntimeMetaObject struct {
	mock.Mock
}

// DeepCopyObject provides a mock function with given fields:
func (_m *RuntimeMetaObject) DeepCopyObject() runtime.Object {
	ret := _m.Called()

	var r0 runtime.Object
	if rf, ok := ret.Get(0).(func() runtime.Object); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(runtime.Object)
		}
	}

	return r0
}

// GetAnnotations provides a mock function with given fields:
func (_m *RuntimeMetaObject) GetAnnotations() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

// GetClusterName provides a mock function with given fields:
func (_m *RuntimeMetaObject) GetClusterName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetCreationTimestamp provides a mock function with given fields:
func (_m *RuntimeMetaObject) GetCreationTimestamp() v1.Time {
	ret := _m.Called()

	var r0 v1.Time
	if rf, ok := ret.Get(0).(func() v1.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1.Time)
	}

	return r0
}

// GetDeletionGracePeriodSeconds provides a mock function with given fields:
func (_m *RuntimeMetaObject) GetDeletionGracePeriodSeconds() *int64 {
	ret := _m.Called()

	var r0 *int64
	if rf, ok := ret.Get(0).(func() *int64); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int64)
		}
	}

	return r0
}

// GetDeletionTimestamp provides a mock function with given fields:
func (_m *RuntimeMetaObject) GetDeletionTimestamp() *v1.Time {
	ret := _m.Called()

	var r0 *v1.Time
	if rf, ok := ret.Get(0).(func() *v1.Time); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Time)
		}
	}

	return r0
}

// GetFinalizers provides a mock function with given fields:
func (_m *RuntimeMetaObject) GetFinalizers() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// GetGenerateName provides a mock function with given fields:
func (_m *RuntimeMetaObject) GetGenerateName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetGeneration provides a mock function with given fields:
func (_m *RuntimeMetaObject) GetGeneration() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// GetLabels provides a mock function with given fields:
func (_m *RuntimeMetaObject) GetLabels() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

// GetManagedFields provides a mock function with given fields:
func (_m *RuntimeMetaObject) GetManagedFields() []v1.ManagedFieldsEntry {
	ret := _m.Called()

	var r0 []v1.ManagedFieldsEntry
	if rf, ok := ret.Get(0).(func() []v1.ManagedFieldsEntry); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v1.ManagedFieldsEntry)
		}
	}

	return r0
}

// GetName provides a mock function with given fields:
func (_m *RuntimeMetaObject) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetNamespace provides a mock function with given fields:
func (_m *RuntimeMetaObject) GetNamespace() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetObjectKind provides a mock function with given fields:
func (_m *RuntimeMetaObject) GetObjectKind() schema.ObjectKind {
	ret := _m.Called()

	var r0 schema.ObjectKind
	if rf, ok := ret.Get(0).(func() schema.ObjectKind); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(schema.ObjectKind)
		}
	}

	return r0
}

// GetOwnerReferences provides a mock function with given fields:
func (_m *RuntimeMetaObject) GetOwnerReferences() []v1.OwnerReference {
	ret := _m.Called()

	var r0 []v1.OwnerReference
	if rf, ok := ret.Get(0).(func() []v1.OwnerReference); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v1.OwnerReference)
		}
	}

	return r0
}

// GetResourceVersion provides a mock function with given fields:
func (_m *RuntimeMetaObject) GetResourceVersion() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetSelfLink provides a mock function with given fields:
func (_m *RuntimeMetaObject) GetSelfLink() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetUID provides a mock function with given fields:
func (_m *RuntimeMetaObject) GetUID() pkgtypes.UID {
	ret := _m.Called()

	var r0 pkgtypes.UID
	if rf, ok := ret.Get(0).(func() pkgtypes.UID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(pkgtypes.UID)
	}

	return r0
}

// SetAnnotations provides a mock function with given fields: annotations
func (_m *RuntimeMetaObject) SetAnnotations(annotations map[string]string) {
	_m.Called(annotations)
}

// SetClusterName provides a mock function with given fields: clusterName
func (_m *RuntimeMetaObject) SetClusterName(clusterName string) {
	_m.Called(clusterName)
}

// SetCreationTimestamp provides a mock function with given fields: timestamp
func (_m *RuntimeMetaObject) SetCreationTimestamp(timestamp v1.Time) {
	_m.Called(timestamp)
}

// SetDeletionGracePeriodSeconds provides a mock function with given fields: _a0
func (_m *RuntimeMetaObject) SetDeletionGracePeriodSeconds(_a0 *int64) {
	_m.Called(_a0)
}

// SetDeletionTimestamp provides a mock function with given fields: timestamp
func (_m *RuntimeMetaObject) SetDeletionTimestamp(timestamp *v1.Time) {
	_m.Called(timestamp)
}

// SetFinalizers provides a mock function with given fields: finalizers
func (_m *RuntimeMetaObject) SetFinalizers(finalizers []string) {
	_m.Called(finalizers)
}

// SetGenerateName provides a mock function with given fields: name
func (_m *RuntimeMetaObject) SetGenerateName(name string) {
	_m.Called(name)
}

// SetGeneration provides a mock function with given fields: generation
func (_m *RuntimeMetaObject) SetGeneration(generation int64) {
	_m.Called(generation)
}

// SetLabels provides a mock function with given fields: labels
func (_m *RuntimeMetaObject) SetLabels(labels map[string]string) {
	_m.Called(labels)
}

// SetManagedFields provides a mock function with given fields: managedFields
func (_m *RuntimeMetaObject) SetManagedFields(managedFields []v1.ManagedFieldsEntry) {
	_m.Called(managedFields)
}

// SetName provides a mock function with given fields: name
func (_m *RuntimeMetaObject) SetName(name string) {
	_m.Called(name)
}

// SetNamespace provides a mock function with given fields: namespace
func (_m *RuntimeMetaObject) SetNamespace(namespace string) {
	_m.Called(namespace)
}

// SetOwnerReferences provides a mock function with given fields: _a0
func (_m *RuntimeMetaObject) SetOwnerReferences(_a0 []v1.OwnerReference) {
	_m.Called(_a0)
}

// SetResourceVersion provides a mock function with given fields: version
func (_m *RuntimeMetaObject) SetResourceVersion(version string) {
	_m.Called(version)
}

// SetSelfLink provides a mock function with given fields: selfLink
func (_m *RuntimeMetaObject) SetSelfLink(selfLink string) {
	_m.Called(selfLink)
}

// SetUID provides a mock function with given fields: uid
func (_m *RuntimeMetaObject) SetUID(uid pkgtypes.UID) {
	_m.Called(uid)
}
