/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.10
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// LoadBalancerTag Information about the load balancer tag.
type LoadBalancerTag struct {
	// The key of the tag.
	Key *string `json:"Key,omitempty"`
	// The name of the load balancer.
	LoadBalancerName *string `json:"LoadBalancerName,omitempty"`
	// The value of the tag.
	Value *string `json:"Value,omitempty"`
}

// NewLoadBalancerTag instantiates a new LoadBalancerTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerTag() *LoadBalancerTag {
	this := LoadBalancerTag{}
	return &this
}

// NewLoadBalancerTagWithDefaults instantiates a new LoadBalancerTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerTagWithDefaults() *LoadBalancerTag {
	this := LoadBalancerTag{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *LoadBalancerTag) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerTag) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *LoadBalancerTag) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *LoadBalancerTag) SetKey(v string) {
	o.Key = &v
}

// GetLoadBalancerName returns the LoadBalancerName field value if set, zero value otherwise.
func (o *LoadBalancerTag) GetLoadBalancerName() string {
	if o == nil || o.LoadBalancerName == nil {
		var ret string
		return ret
	}
	return *o.LoadBalancerName
}

// GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerTag) GetLoadBalancerNameOk() (*string, bool) {
	if o == nil || o.LoadBalancerName == nil {
		return nil, false
	}
	return o.LoadBalancerName, true
}

// HasLoadBalancerName returns a boolean if a field has been set.
func (o *LoadBalancerTag) HasLoadBalancerName() bool {
	if o != nil && o.LoadBalancerName != nil {
		return true
	}

	return false
}

// SetLoadBalancerName gets a reference to the given string and assigns it to the LoadBalancerName field.
func (o *LoadBalancerTag) SetLoadBalancerName(v string) {
	o.LoadBalancerName = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *LoadBalancerTag) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerTag) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *LoadBalancerTag) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *LoadBalancerTag) SetValue(v string) {
	o.Value = &v
}

func (o LoadBalancerTag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.LoadBalancerName != nil {
		toSerialize["LoadBalancerName"] = o.LoadBalancerName
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableLoadBalancerTag struct {
	value *LoadBalancerTag
	isSet bool
}

func (v NullableLoadBalancerTag) Get() *LoadBalancerTag {
	return v.value
}

func (v *NullableLoadBalancerTag) Set(val *LoadBalancerTag) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerTag) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerTag(val *LoadBalancerTag) *NullableLoadBalancerTag {
	return &NullableLoadBalancerTag{value: val, isSet: true}
}

func (v NullableLoadBalancerTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}