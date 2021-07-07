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

// NetToVirtualGatewayLink Information about the attachment.
type NetToVirtualGatewayLink struct {
	// The ID of the Net to which the virtual gateway is attached.
	NetId *string `json:"NetId,omitempty"`
	// The state of the attachment (`attaching` \\| `attached` \\| `detaching` \\| `detached`).
	State *string `json:"State,omitempty"`
}

// NewNetToVirtualGatewayLink instantiates a new NetToVirtualGatewayLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetToVirtualGatewayLink() *NetToVirtualGatewayLink {
	this := NetToVirtualGatewayLink{}
	return &this
}

// NewNetToVirtualGatewayLinkWithDefaults instantiates a new NetToVirtualGatewayLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetToVirtualGatewayLinkWithDefaults() *NetToVirtualGatewayLink {
	this := NetToVirtualGatewayLink{}
	return &this
}

// GetNetId returns the NetId field value if set, zero value otherwise.
func (o *NetToVirtualGatewayLink) GetNetId() string {
	if o == nil || o.NetId == nil {
		var ret string
		return ret
	}
	return *o.NetId
}

// GetNetIdOk returns a tuple with the NetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetToVirtualGatewayLink) GetNetIdOk() (*string, bool) {
	if o == nil || o.NetId == nil {
		return nil, false
	}
	return o.NetId, true
}

// HasNetId returns a boolean if a field has been set.
func (o *NetToVirtualGatewayLink) HasNetId() bool {
	if o != nil && o.NetId != nil {
		return true
	}

	return false
}

// SetNetId gets a reference to the given string and assigns it to the NetId field.
func (o *NetToVirtualGatewayLink) SetNetId(v string) {
	o.NetId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *NetToVirtualGatewayLink) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetToVirtualGatewayLink) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *NetToVirtualGatewayLink) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *NetToVirtualGatewayLink) SetState(v string) {
	o.State = &v
}

func (o NetToVirtualGatewayLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetId != nil {
		toSerialize["NetId"] = o.NetId
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableNetToVirtualGatewayLink struct {
	value *NetToVirtualGatewayLink
	isSet bool
}

func (v NullableNetToVirtualGatewayLink) Get() *NetToVirtualGatewayLink {
	return v.value
}

func (v *NullableNetToVirtualGatewayLink) Set(val *NetToVirtualGatewayLink) {
	v.value = val
	v.isSet = true
}

func (v NullableNetToVirtualGatewayLink) IsSet() bool {
	return v.isSet
}

func (v *NullableNetToVirtualGatewayLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetToVirtualGatewayLink(val *NetToVirtualGatewayLink) *NullableNetToVirtualGatewayLink {
	return &NullableNetToVirtualGatewayLink{value: val, isSet: true}
}

func (v NullableNetToVirtualGatewayLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetToVirtualGatewayLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}