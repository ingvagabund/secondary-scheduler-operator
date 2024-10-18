/*
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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	operatorv1 "github.com/openshift/api/operator/v1"
	v1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// SecondarySchedulerSpecApplyConfiguration represents a declarative configuration of the SecondarySchedulerSpec type for use
// with apply.
type SecondarySchedulerSpecApplyConfiguration struct {
	v1.OperatorSpecApplyConfiguration `json:",inline"`
	SchedulerConfig                   *string `json:"schedulerConfig,omitempty"`
	SchedulerImage                    *string `json:"schedulerImage,omitempty"`
}

// SecondarySchedulerSpecApplyConfiguration constructs a declarative configuration of the SecondarySchedulerSpec type for use with
// apply.
func SecondarySchedulerSpec() *SecondarySchedulerSpecApplyConfiguration {
	return &SecondarySchedulerSpecApplyConfiguration{}
}

// WithManagementState sets the ManagementState field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ManagementState field is set to the value of the last call.
func (b *SecondarySchedulerSpecApplyConfiguration) WithManagementState(value operatorv1.ManagementState) *SecondarySchedulerSpecApplyConfiguration {
	b.ManagementState = &value
	return b
}

// WithLogLevel sets the LogLevel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LogLevel field is set to the value of the last call.
func (b *SecondarySchedulerSpecApplyConfiguration) WithLogLevel(value operatorv1.LogLevel) *SecondarySchedulerSpecApplyConfiguration {
	b.LogLevel = &value
	return b
}

// WithOperatorLogLevel sets the OperatorLogLevel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OperatorLogLevel field is set to the value of the last call.
func (b *SecondarySchedulerSpecApplyConfiguration) WithOperatorLogLevel(value operatorv1.LogLevel) *SecondarySchedulerSpecApplyConfiguration {
	b.OperatorLogLevel = &value
	return b
}

// WithUnsupportedConfigOverrides sets the UnsupportedConfigOverrides field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UnsupportedConfigOverrides field is set to the value of the last call.
func (b *SecondarySchedulerSpecApplyConfiguration) WithUnsupportedConfigOverrides(value runtime.RawExtension) *SecondarySchedulerSpecApplyConfiguration {
	b.UnsupportedConfigOverrides = &value
	return b
}

// WithObservedConfig sets the ObservedConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedConfig field is set to the value of the last call.
func (b *SecondarySchedulerSpecApplyConfiguration) WithObservedConfig(value runtime.RawExtension) *SecondarySchedulerSpecApplyConfiguration {
	b.ObservedConfig = &value
	return b
}

// WithSchedulerConfig sets the SchedulerConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SchedulerConfig field is set to the value of the last call.
func (b *SecondarySchedulerSpecApplyConfiguration) WithSchedulerConfig(value string) *SecondarySchedulerSpecApplyConfiguration {
	b.SchedulerConfig = &value
	return b
}

// WithSchedulerImage sets the SchedulerImage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SchedulerImage field is set to the value of the last call.
func (b *SecondarySchedulerSpecApplyConfiguration) WithSchedulerImage(value string) *SecondarySchedulerSpecApplyConfiguration {
	b.SchedulerImage = &value
	return b
}
