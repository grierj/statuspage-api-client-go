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
)

// PostPagesPageIdPageAccessUsersPageAccessUser struct for PostPagesPageIdPageAccessUsersPageAccessUser
type PostPagesPageIdPageAccessUsersPageAccessUser struct {
	// IDP login user id. Key is typically \"uid\".
	ExternalLogin *string `json:"external_login,omitempty"`
	Email *string `json:"email,omitempty"`
	PageAccessGroupIds *[]string `json:"page_access_group_ids,omitempty"`
	SubscribeToComponents *bool `json:"subscribe_to_components,omitempty"`
}

// NewPostPagesPageIdPageAccessUsersPageAccessUser instantiates a new PostPagesPageIdPageAccessUsersPageAccessUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostPagesPageIdPageAccessUsersPageAccessUser() *PostPagesPageIdPageAccessUsersPageAccessUser {
	this := PostPagesPageIdPageAccessUsersPageAccessUser{}
	return &this
}

// NewPostPagesPageIdPageAccessUsersPageAccessUserWithDefaults instantiates a new PostPagesPageIdPageAccessUsersPageAccessUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostPagesPageIdPageAccessUsersPageAccessUserWithDefaults() *PostPagesPageIdPageAccessUsersPageAccessUser {
	this := PostPagesPageIdPageAccessUsersPageAccessUser{}
	return &this
}

// GetExternalLogin returns the ExternalLogin field value if set, zero value otherwise.
func (o *PostPagesPageIdPageAccessUsersPageAccessUser) GetExternalLogin() string {
	if o == nil || o.ExternalLogin == nil {
		var ret string
		return ret
	}
	return *o.ExternalLogin
}

// GetExternalLoginOk returns a tuple with the ExternalLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPagesPageIdPageAccessUsersPageAccessUser) GetExternalLoginOk() (*string, bool) {
	if o == nil || o.ExternalLogin == nil {
		return nil, false
	}
	return o.ExternalLogin, true
}

// HasExternalLogin returns a boolean if a field has been set.
func (o *PostPagesPageIdPageAccessUsersPageAccessUser) HasExternalLogin() bool {
	if o != nil && o.ExternalLogin != nil {
		return true
	}

	return false
}

// SetExternalLogin gets a reference to the given string and assigns it to the ExternalLogin field.
func (o *PostPagesPageIdPageAccessUsersPageAccessUser) SetExternalLogin(v string) {
	o.ExternalLogin = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PostPagesPageIdPageAccessUsersPageAccessUser) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPagesPageIdPageAccessUsersPageAccessUser) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PostPagesPageIdPageAccessUsersPageAccessUser) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PostPagesPageIdPageAccessUsersPageAccessUser) SetEmail(v string) {
	o.Email = &v
}

// GetPageAccessGroupIds returns the PageAccessGroupIds field value if set, zero value otherwise.
func (o *PostPagesPageIdPageAccessUsersPageAccessUser) GetPageAccessGroupIds() []string {
	if o == nil || o.PageAccessGroupIds == nil {
		var ret []string
		return ret
	}
	return *o.PageAccessGroupIds
}

// GetPageAccessGroupIdsOk returns a tuple with the PageAccessGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPagesPageIdPageAccessUsersPageAccessUser) GetPageAccessGroupIdsOk() (*[]string, bool) {
	if o == nil || o.PageAccessGroupIds == nil {
		return nil, false
	}
	return o.PageAccessGroupIds, true
}

// HasPageAccessGroupIds returns a boolean if a field has been set.
func (o *PostPagesPageIdPageAccessUsersPageAccessUser) HasPageAccessGroupIds() bool {
	if o != nil && o.PageAccessGroupIds != nil {
		return true
	}

	return false
}

// SetPageAccessGroupIds gets a reference to the given []string and assigns it to the PageAccessGroupIds field.
func (o *PostPagesPageIdPageAccessUsersPageAccessUser) SetPageAccessGroupIds(v []string) {
	o.PageAccessGroupIds = &v
}

// GetSubscribeToComponents returns the SubscribeToComponents field value if set, zero value otherwise.
func (o *PostPagesPageIdPageAccessUsersPageAccessUser) GetSubscribeToComponents() bool {
	if o == nil || o.SubscribeToComponents == nil {
		var ret bool
		return ret
	}
	return *o.SubscribeToComponents
}

// GetSubscribeToComponentsOk returns a tuple with the SubscribeToComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPagesPageIdPageAccessUsersPageAccessUser) GetSubscribeToComponentsOk() (*bool, bool) {
	if o == nil || o.SubscribeToComponents == nil {
		return nil, false
	}
	return o.SubscribeToComponents, true
}

// HasSubscribeToComponents returns a boolean if a field has been set.
func (o *PostPagesPageIdPageAccessUsersPageAccessUser) HasSubscribeToComponents() bool {
	if o != nil && o.SubscribeToComponents != nil {
		return true
	}

	return false
}

// SetSubscribeToComponents gets a reference to the given bool and assigns it to the SubscribeToComponents field.
func (o *PostPagesPageIdPageAccessUsersPageAccessUser) SetSubscribeToComponents(v bool) {
	o.SubscribeToComponents = &v
}

func (o PostPagesPageIdPageAccessUsersPageAccessUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalLogin != nil {
		toSerialize["external_login"] = o.ExternalLogin
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.PageAccessGroupIds != nil {
		toSerialize["page_access_group_ids"] = o.PageAccessGroupIds
	}
	if o.SubscribeToComponents != nil {
		toSerialize["subscribe_to_components"] = o.SubscribeToComponents
	}
	return json.Marshal(toSerialize)
}

type NullablePostPagesPageIdPageAccessUsersPageAccessUser struct {
	value *PostPagesPageIdPageAccessUsersPageAccessUser
	isSet bool
}

func (v NullablePostPagesPageIdPageAccessUsersPageAccessUser) Get() *PostPagesPageIdPageAccessUsersPageAccessUser {
	return v.value
}

func (v *NullablePostPagesPageIdPageAccessUsersPageAccessUser) Set(val *PostPagesPageIdPageAccessUsersPageAccessUser) {
	v.value = val
	v.isSet = true
}

func (v NullablePostPagesPageIdPageAccessUsersPageAccessUser) IsSet() bool {
	return v.isSet
}

func (v *NullablePostPagesPageIdPageAccessUsersPageAccessUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostPagesPageIdPageAccessUsersPageAccessUser(val *PostPagesPageIdPageAccessUsersPageAccessUser) *NullablePostPagesPageIdPageAccessUsersPageAccessUser {
	return &NullablePostPagesPageIdPageAccessUsersPageAccessUser{value: val, isSet: true}
}

func (v NullablePostPagesPageIdPageAccessUsersPageAccessUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostPagesPageIdPageAccessUsersPageAccessUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


