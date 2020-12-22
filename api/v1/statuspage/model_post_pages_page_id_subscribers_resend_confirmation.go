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

// PostPagesPageIdSubscribersResendConfirmation Resend confirmations to a list of subscribers
type PostPagesPageIdSubscribersResendConfirmation struct {
	// The array of subscriber codes to resend confirmations for, or \"all\" to resend confirmations to all subscribers. Only unconfirmed email subscribers will receive this notification.
	Subscribers string `json:"subscribers"`
}

// NewPostPagesPageIdSubscribersResendConfirmation instantiates a new PostPagesPageIdSubscribersResendConfirmation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostPagesPageIdSubscribersResendConfirmation(subscribers string, ) *PostPagesPageIdSubscribersResendConfirmation {
	this := PostPagesPageIdSubscribersResendConfirmation{}
	this.Subscribers = subscribers
	return &this
}

// NewPostPagesPageIdSubscribersResendConfirmationWithDefaults instantiates a new PostPagesPageIdSubscribersResendConfirmation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostPagesPageIdSubscribersResendConfirmationWithDefaults() *PostPagesPageIdSubscribersResendConfirmation {
	this := PostPagesPageIdSubscribersResendConfirmation{}
	return &this
}

// GetSubscribers returns the Subscribers field value
func (o *PostPagesPageIdSubscribersResendConfirmation) GetSubscribers() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Subscribers
}

// GetSubscribersOk returns a tuple with the Subscribers field value
// and a boolean to check if the value has been set.
func (o *PostPagesPageIdSubscribersResendConfirmation) GetSubscribersOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Subscribers, true
}

// SetSubscribers sets field value
func (o *PostPagesPageIdSubscribersResendConfirmation) SetSubscribers(v string) {
	o.Subscribers = v
}

func (o PostPagesPageIdSubscribersResendConfirmation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subscribers"] = o.Subscribers
	}
	return json.Marshal(toSerialize)
}

type NullablePostPagesPageIdSubscribersResendConfirmation struct {
	value *PostPagesPageIdSubscribersResendConfirmation
	isSet bool
}

func (v NullablePostPagesPageIdSubscribersResendConfirmation) Get() *PostPagesPageIdSubscribersResendConfirmation {
	return v.value
}

func (v *NullablePostPagesPageIdSubscribersResendConfirmation) Set(val *PostPagesPageIdSubscribersResendConfirmation) {
	v.value = val
	v.isSet = true
}

func (v NullablePostPagesPageIdSubscribersResendConfirmation) IsSet() bool {
	return v.isSet
}

func (v *NullablePostPagesPageIdSubscribersResendConfirmation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostPagesPageIdSubscribersResendConfirmation(val *PostPagesPageIdSubscribersResendConfirmation) *NullablePostPagesPageIdSubscribersResendConfirmation {
	return &NullablePostPagesPageIdSubscribersResendConfirmation{value: val, isSet: true}
}

func (v NullablePostPagesPageIdSubscribersResendConfirmation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostPagesPageIdSubscribersResendConfirmation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


