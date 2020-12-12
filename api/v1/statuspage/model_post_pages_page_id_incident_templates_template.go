/*
 * Statuspage API
 *
 * # Code of Conduct Please don't abuse the API, and please report all feature requests and issues to https://help.statuspage.io/help/contact-us-30  # Rate Limiting Each API token is limited to 1 request / second as measured on a 60 second rolling window. To get this limit increased, please contact us at https://help.statuspage.io/help/contact-us-30  # Basics  ## HTTPS It's required  ## URL Prefix In order to maintain version integrity into the future, the API is versioned. All calls currently begin with the following prefix:    https://api.statuspage.io/v1/  ## RESTful Interface Wherever possible, the API seeks to implement repeatable patterns with logical, representative URLs and descriptive HTTP verbs. Below are some examples and conventions you will see throughout the documentation.  * Collections are buckets: https://api.statuspage.io/v1/pages/asdf123/incidents.json * Elements have unique IDs: https://api.statuspage.io/v1/pages/asdf123/incidents/jklm456.json * GET will retrieve information about a collection/element * POST will create an element in a collection * PATCH will update a single element * PUT will replace a single element in a collection (rarely used) * DELETE will destroy a single element  ## Sending Data Information can be sent in the body as form urlencoded or JSON, but make sure the Content-Type header matches the body structure or the server gremlins will be angry.  All examples are provided in JSON format, however they can easily be converted to form encoding if required.  Some examples of how to convert things are below:      // JSON     {       \"incident\": {         \"name\": \"test incident\",         \"components\": [\"8kbf7d35c070\", \"vtnh60py4yd7\"]       }     }      // Form Encoded (using curl as an example):     curl -X POST https://api.statuspage.io/v1/example \\       -d \"incident[name]=test incident\" \\       -d \"incident[components][]=8kbf7d35c070\" \\       -d \"incident[components][]=vtnh60py4yd7\"  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PostPagesPageIdIncidentTemplatesTemplate struct {
	// Name of the template, as shown in the list on the \"Templates\" tab of the \"Incidents\" page
	Name string `json:"name,omitempty"`
	// Title to be applied to the incident or maintenance when selecting this template
	Title string `json:"title,omitempty"`
	// The initial message, created as the first incident or maintenance update.
	Body string `json:"body,omitempty"`
	// Identifier of Template Group this template belongs to
	GroupId string `json:"group_id,omitempty"`
	// The status the incident or maintenance should transition to when selecting this template
	UpdateStatus string `json:"update_status,omitempty"`
	// Whether the \"tweet update\" checkbox should be selected when selecting this template
	ShouldTweet bool `json:"should_tweet,omitempty"`
	// Whether the \"deliver notifications\" checkbox should be selected when selecting this template
	ShouldSendNotifications bool `json:"should_send_notifications,omitempty"`
	// List of component_ids affected by this incident
	ComponentIds []string `json:"component_ids,omitempty"`
}