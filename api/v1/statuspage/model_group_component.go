/*
 * Statuspage API
 *
 * # Code of Conduct Please don't abuse the API, and please report all feature requests and issues to https://help.statuspage.io/help/contact-us-30  # Rate Limiting Each API token is limited to 1 request / second as measured on a 60 second rolling window. To get this limit increased, please contact us at https://help.statuspage.io/help/contact-us-30  # Basics  ## HTTPS It's required  ## URL Prefix In order to maintain version integrity into the future, the API is versioned. All calls currently begin with the following prefix:    https://api.statuspage.io/v1/  ## RESTful Interface Wherever possible, the API seeks to implement repeatable patterns with logical, representative URLs and descriptive HTTP verbs. Below are some examples and conventions you will see throughout the documentation.  * Collections are buckets: https://api.statuspage.io/v1/pages/asdf123/incidents.json * Elements have unique IDs: https://api.statuspage.io/v1/pages/asdf123/incidents/jklm456.json * GET will retrieve information about a collection/element * POST will create an element in a collection * PATCH will update a single element * PUT will replace a single element in a collection (rarely used) * DELETE will destroy a single element  ## Sending Data Information can be sent in the body as form urlencoded or JSON, but make sure the Content-Type header matches the body structure or the server gremlins will be angry.  All examples are provided in JSON format, however they can easily be converted to form encoding if required.  Some examples of how to convert things are below:      // JSON     {       \"incident\": {         \"name\": \"test incident\",         \"components\": [\"8kbf7d35c070\", \"vtnh60py4yd7\"]       }     }      // Form Encoded (using curl as an example):     curl -X POST https://api.statuspage.io/v1/example \\       -d \"incident[name]=test incident\" \\       -d \"incident[components][]=8kbf7d35c070\" \\       -d \"incident[components][]=vtnh60py4yd7\"  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// GroupComponent Get a component group
type GroupComponent struct {
	// Component Group Identifier
	Id *string `json:"id,omitempty"`
	PageId *string `json:"page_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Components *[]string `json:"components,omitempty"`
	Position *int32 `json:"position,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewGroupComponent instantiates a new GroupComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupComponent() *GroupComponent {
	this := GroupComponent{}
	return &this
}

// NewGroupComponentWithDefaults instantiates a new GroupComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupComponentWithDefaults() *GroupComponent {
	this := GroupComponent{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupComponent) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupComponent) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupComponent) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GroupComponent) SetId(v string) {
	o.Id = &v
}

// GetPageId returns the PageId field value if set, zero value otherwise.
func (o *GroupComponent) GetPageId() string {
	if o == nil || o.PageId == nil {
		var ret string
		return ret
	}
	return *o.PageId
}

// GetPageIdOk returns a tuple with the PageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupComponent) GetPageIdOk() (*string, bool) {
	if o == nil || o.PageId == nil {
		return nil, false
	}
	return o.PageId, true
}

// HasPageId returns a boolean if a field has been set.
func (o *GroupComponent) HasPageId() bool {
	if o != nil && o.PageId != nil {
		return true
	}

	return false
}

// SetPageId gets a reference to the given string and assigns it to the PageId field.
func (o *GroupComponent) SetPageId(v string) {
	o.PageId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GroupComponent) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupComponent) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupComponent) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupComponent) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GroupComponent) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupComponent) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroupComponent) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroupComponent) SetDescription(v string) {
	o.Description = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *GroupComponent) GetComponents() []string {
	if o == nil || o.Components == nil {
		var ret []string
		return ret
	}
	return *o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupComponent) GetComponentsOk() (*[]string, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *GroupComponent) HasComponents() bool {
	if o != nil && o.Components != nil {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []string and assigns it to the Components field.
func (o *GroupComponent) SetComponents(v []string) {
	o.Components = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *GroupComponent) GetPosition() int32 {
	if o == nil || o.Position == nil {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupComponent) GetPositionOk() (*int32, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *GroupComponent) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *GroupComponent) SetPosition(v int32) {
	o.Position = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GroupComponent) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupComponent) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GroupComponent) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *GroupComponent) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GroupComponent) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupComponent) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GroupComponent) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *GroupComponent) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o GroupComponent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PageId != nil {
		toSerialize["page_id"] = o.PageId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableGroupComponent struct {
	value *GroupComponent
	isSet bool
}

func (v NullableGroupComponent) Get() *GroupComponent {
	return v.value
}

func (v *NullableGroupComponent) Set(val *GroupComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupComponent(val *GroupComponent) *NullableGroupComponent {
	return &NullableGroupComponent{value: val, isSet: true}
}

func (v NullableGroupComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


