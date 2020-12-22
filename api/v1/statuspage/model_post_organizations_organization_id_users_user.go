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

// PostOrganizationsOrganizationIdUsersUser struct for PostOrganizationsOrganizationIdUsersUser
type PostOrganizationsOrganizationIdUsersUser struct {
	// Email address for the team member
	Email *string `json:"email,omitempty"`
	// Password the team member uses to access the site
	Password *string `json:"password,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName *string `json:"last_name,omitempty"`
}

// NewPostOrganizationsOrganizationIdUsersUser instantiates a new PostOrganizationsOrganizationIdUsersUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostOrganizationsOrganizationIdUsersUser() *PostOrganizationsOrganizationIdUsersUser {
	this := PostOrganizationsOrganizationIdUsersUser{}
	return &this
}

// NewPostOrganizationsOrganizationIdUsersUserWithDefaults instantiates a new PostOrganizationsOrganizationIdUsersUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostOrganizationsOrganizationIdUsersUserWithDefaults() *PostOrganizationsOrganizationIdUsersUser {
	this := PostOrganizationsOrganizationIdUsersUser{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PostOrganizationsOrganizationIdUsersUser) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostOrganizationsOrganizationIdUsersUser) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PostOrganizationsOrganizationIdUsersUser) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PostOrganizationsOrganizationIdUsersUser) SetEmail(v string) {
	o.Email = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *PostOrganizationsOrganizationIdUsersUser) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostOrganizationsOrganizationIdUsersUser) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *PostOrganizationsOrganizationIdUsersUser) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *PostOrganizationsOrganizationIdUsersUser) SetPassword(v string) {
	o.Password = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *PostOrganizationsOrganizationIdUsersUser) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostOrganizationsOrganizationIdUsersUser) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *PostOrganizationsOrganizationIdUsersUser) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *PostOrganizationsOrganizationIdUsersUser) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *PostOrganizationsOrganizationIdUsersUser) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostOrganizationsOrganizationIdUsersUser) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *PostOrganizationsOrganizationIdUsersUser) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *PostOrganizationsOrganizationIdUsersUser) SetLastName(v string) {
	o.LastName = &v
}

func (o PostOrganizationsOrganizationIdUsersUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	return json.Marshal(toSerialize)
}

type NullablePostOrganizationsOrganizationIdUsersUser struct {
	value *PostOrganizationsOrganizationIdUsersUser
	isSet bool
}

func (v NullablePostOrganizationsOrganizationIdUsersUser) Get() *PostOrganizationsOrganizationIdUsersUser {
	return v.value
}

func (v *NullablePostOrganizationsOrganizationIdUsersUser) Set(val *PostOrganizationsOrganizationIdUsersUser) {
	v.value = val
	v.isSet = true
}

func (v NullablePostOrganizationsOrganizationIdUsersUser) IsSet() bool {
	return v.isSet
}

func (v *NullablePostOrganizationsOrganizationIdUsersUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostOrganizationsOrganizationIdUsersUser(val *PostOrganizationsOrganizationIdUsersUser) *NullablePostOrganizationsOrganizationIdUsersUser {
	return &NullablePostOrganizationsOrganizationIdUsersUser{value: val, isSet: true}
}

func (v NullablePostOrganizationsOrganizationIdUsersUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostOrganizationsOrganizationIdUsersUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


