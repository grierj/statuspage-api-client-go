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

// PostPagesPageIdIncidentTemplates Create a template
type PostPagesPageIdIncidentTemplates struct {
	Template *PostPagesPageIdIncidentTemplatesTemplate `json:"template,omitempty"`
}

// NewPostPagesPageIdIncidentTemplates instantiates a new PostPagesPageIdIncidentTemplates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostPagesPageIdIncidentTemplates() *PostPagesPageIdIncidentTemplates {
	this := PostPagesPageIdIncidentTemplates{}
	return &this
}

// NewPostPagesPageIdIncidentTemplatesWithDefaults instantiates a new PostPagesPageIdIncidentTemplates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostPagesPageIdIncidentTemplatesWithDefaults() *PostPagesPageIdIncidentTemplates {
	this := PostPagesPageIdIncidentTemplates{}
	return &this
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *PostPagesPageIdIncidentTemplates) GetTemplate() PostPagesPageIdIncidentTemplatesTemplate {
	if o == nil || o.Template == nil {
		var ret PostPagesPageIdIncidentTemplatesTemplate
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPagesPageIdIncidentTemplates) GetTemplateOk() (*PostPagesPageIdIncidentTemplatesTemplate, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *PostPagesPageIdIncidentTemplates) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given PostPagesPageIdIncidentTemplatesTemplate and assigns it to the Template field.
func (o *PostPagesPageIdIncidentTemplates) SetTemplate(v PostPagesPageIdIncidentTemplatesTemplate) {
	o.Template = &v
}

func (o PostPagesPageIdIncidentTemplates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	return json.Marshal(toSerialize)
}

type NullablePostPagesPageIdIncidentTemplates struct {
	value *PostPagesPageIdIncidentTemplates
	isSet bool
}

func (v NullablePostPagesPageIdIncidentTemplates) Get() *PostPagesPageIdIncidentTemplates {
	return v.value
}

func (v *NullablePostPagesPageIdIncidentTemplates) Set(val *PostPagesPageIdIncidentTemplates) {
	v.value = val
	v.isSet = true
}

func (v NullablePostPagesPageIdIncidentTemplates) IsSet() bool {
	return v.isSet
}

func (v *NullablePostPagesPageIdIncidentTemplates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostPagesPageIdIncidentTemplates(val *PostPagesPageIdIncidentTemplates) *NullablePostPagesPageIdIncidentTemplates {
	return &NullablePostPagesPageIdIncidentTemplates{value: val, isSet: true}
}

func (v NullablePostPagesPageIdIncidentTemplates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostPagesPageIdIncidentTemplates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


