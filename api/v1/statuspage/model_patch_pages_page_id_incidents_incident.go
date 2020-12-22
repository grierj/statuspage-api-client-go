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

// PatchPagesPageIdIncidentsIncident struct for PatchPagesPageIdIncidentsIncident
type PatchPagesPageIdIncidentsIncident struct {
	// Incident Name
	Name *string `json:"name,omitempty"`
	// The incident status. For realtime incidents, valid values are investigating, identified, monitoring, and resolved. For scheduled incidents, valid values are scheduled, in_progress, verifying, and completed.
	Status *string `json:"status,omitempty"`
	// value to override calculated impact value
	ImpactOverride *string `json:"impact_override,omitempty"`
	// The timestamp the incident is scheduled for.
	ScheduledFor *time.Time `json:"scheduled_for,omitempty"`
	// The timestamp the incident is scheduled until.
	ScheduledUntil *time.Time `json:"scheduled_until,omitempty"`
	// Controls whether to remind subscribers prior to scheduled incidents.
	ScheduledRemindPrior *bool `json:"scheduled_remind_prior,omitempty"`
	// Controls whether the incident is scheduled to automatically change to in progress.
	ScheduledAutoInProgress *bool `json:"scheduled_auto_in_progress,omitempty"`
	// Controls whether the incident is scheduled to automatically change to complete.
	ScheduledAutoCompleted *bool `json:"scheduled_auto_completed,omitempty"`
	// Attach a json object to the incident. All top-level values in the object must also be objects.
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	// Deliver notifications to subscribers if this is true. If this is false, create an incident without notifying customers.
	DeliverNotifications *bool `json:"deliver_notifications,omitempty"`
	// Controls whether send notification when scheduled maintenances auto transition to completed.
	AutoTransitionDeliverNotificationsAtEnd *bool `json:"auto_transition_deliver_notifications_at_end,omitempty"`
	// Controls whether send notification when scheduled maintenances auto transition to started.
	AutoTransitionDeliverNotificationsAtStart *bool `json:"auto_transition_deliver_notifications_at_start,omitempty"`
	// Controls whether change components status to under_maintenance once scheduled maintenance is in progress.
	AutoTransitionToMaintenanceState *bool `json:"auto_transition_to_maintenance_state,omitempty"`
	// Controls whether change components status to operational once scheduled maintenance completes.
	AutoTransitionToOperationalState *bool `json:"auto_transition_to_operational_state,omitempty"`
	// Controls whether tweet automatically when scheduled maintenance starts.
	AutoTweetAtBeginning *bool `json:"auto_tweet_at_beginning,omitempty"`
	// Controls whether tweet automatically when scheduled maintenance completes.
	AutoTweetOnCompletion *bool `json:"auto_tweet_on_completion,omitempty"`
	// Controls whether tweet automatically when scheduled maintenance is created.
	AutoTweetOnCreation *bool `json:"auto_tweet_on_creation,omitempty"`
	// Controls whether tweet automatically one hour before scheduled maintenance starts.
	AutoTweetOneHourBefore *bool `json:"auto_tweet_one_hour_before,omitempty"`
	// TimeStamp when incident was backfilled.
	BackfillDate *string `json:"backfill_date,omitempty"`
	// Controls whether incident is backfilled. If true, components cannot be specified.
	Backfilled *bool `json:"backfilled,omitempty"`
	// The initial message, created as the first incident update.
	Body *string `json:"body,omitempty"`
	Components *PostPagesPageIdIncidentsIncidentComponents `json:"components,omitempty"`
	// List of component_ids affected by this incident
	ComponentIds *[]string `json:"component_ids,omitempty"`
	// Same as :scheduled_auto_transition_in_progress. Controls whether the incident is scheduled to automatically change to in progress.
	ScheduledAutoTransition *bool `json:"scheduled_auto_transition,omitempty"`
}

// NewPatchPagesPageIdIncidentsIncident instantiates a new PatchPagesPageIdIncidentsIncident object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchPagesPageIdIncidentsIncident() *PatchPagesPageIdIncidentsIncident {
	this := PatchPagesPageIdIncidentsIncident{}
	var deliverNotifications bool = true
	this.DeliverNotifications = &deliverNotifications
	return &this
}

// NewPatchPagesPageIdIncidentsIncidentWithDefaults instantiates a new PatchPagesPageIdIncidentsIncident object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchPagesPageIdIncidentsIncidentWithDefaults() *PatchPagesPageIdIncidentsIncident {
	this := PatchPagesPageIdIncidentsIncident{}
	var deliverNotifications bool = true
	this.DeliverNotifications = &deliverNotifications
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncident) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncident) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncident) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchPagesPageIdIncidentsIncident) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncident) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncident) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncident) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PatchPagesPageIdIncidentsIncident) SetStatus(v string) {
	o.Status = &v
}

// GetImpactOverride returns the ImpactOverride field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncident) GetImpactOverride() string {
	if o == nil || o.ImpactOverride == nil {
		var ret string
		return ret
	}
	return *o.ImpactOverride
}

// GetImpactOverrideOk returns a tuple with the ImpactOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncident) GetImpactOverrideOk() (*string, bool) {
	if o == nil || o.ImpactOverride == nil {
		return nil, false
	}
	return o.ImpactOverride, true
}

// HasImpactOverride returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncident) HasImpactOverride() bool {
	if o != nil && o.ImpactOverride != nil {
		return true
	}

	return false
}

// SetImpactOverride gets a reference to the given string and assigns it to the ImpactOverride field.
func (o *PatchPagesPageIdIncidentsIncident) SetImpactOverride(v string) {
	o.ImpactOverride = &v
}

// GetScheduledFor returns the ScheduledFor field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncident) GetScheduledFor() time.Time {
	if o == nil || o.ScheduledFor == nil {
		var ret time.Time
		return ret
	}
	return *o.ScheduledFor
}

// GetScheduledForOk returns a tuple with the ScheduledFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncident) GetScheduledForOk() (*time.Time, bool) {
	if o == nil || o.ScheduledFor == nil {
		return nil, false
	}
	return o.ScheduledFor, true
}

// HasScheduledFor returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncident) HasScheduledFor() bool {
	if o != nil && o.ScheduledFor != nil {
		return true
	}

	return false
}

// SetScheduledFor gets a reference to the given time.Time and assigns it to the ScheduledFor field.
func (o *PatchPagesPageIdIncidentsIncident) SetScheduledFor(v time.Time) {
	o.ScheduledFor = &v
}

// GetScheduledUntil returns the ScheduledUntil field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncident) GetScheduledUntil() time.Time {
	if o == nil || o.ScheduledUntil == nil {
		var ret time.Time
		return ret
	}
	return *o.ScheduledUntil
}

// GetScheduledUntilOk returns a tuple with the ScheduledUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncident) GetScheduledUntilOk() (*time.Time, bool) {
	if o == nil || o.ScheduledUntil == nil {
		return nil, false
	}
	return o.ScheduledUntil, true
}

// HasScheduledUntil returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncident) HasScheduledUntil() bool {
	if o != nil && o.ScheduledUntil != nil {
		return true
	}

	return false
}

// SetScheduledUntil gets a reference to the given time.Time and assigns it to the ScheduledUntil field.
func (o *PatchPagesPageIdIncidentsIncident) SetScheduledUntil(v time.Time) {
	o.ScheduledUntil = &v
}

// GetScheduledRemindPrior returns the ScheduledRemindPrior field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncident) GetScheduledRemindPrior() bool {
	if o == nil || o.ScheduledRemindPrior == nil {
		var ret bool
		return ret
	}
	return *o.ScheduledRemindPrior
}

// GetScheduledRemindPriorOk returns a tuple with the ScheduledRemindPrior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncident) GetScheduledRemindPriorOk() (*bool, bool) {
	if o == nil || o.ScheduledRemindPrior == nil {
		return nil, false
	}
	return o.ScheduledRemindPrior, true
}

// HasScheduledRemindPrior returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncident) HasScheduledRemindPrior() bool {
	if o != nil && o.ScheduledRemindPrior != nil {
		return true
	}

	return false
}

// SetScheduledRemindPrior gets a reference to the given bool and assigns it to the ScheduledRemindPrior field.
func (o *PatchPagesPageIdIncidentsIncident) SetScheduledRemindPrior(v bool) {
	o.ScheduledRemindPrior = &v
}

// GetScheduledAutoInProgress returns the ScheduledAutoInProgress field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncident) GetScheduledAutoInProgress() bool {
	if o == nil || o.ScheduledAutoInProgress == nil {
		var ret bool
		return ret
	}
	return *o.ScheduledAutoInProgress
}

// GetScheduledAutoInProgressOk returns a tuple with the ScheduledAutoInProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncident) GetScheduledAutoInProgressOk() (*bool, bool) {
	if o == nil || o.ScheduledAutoInProgress == nil {
		return nil, false
	}
	return o.ScheduledAutoInProgress, true
}

// HasScheduledAutoInProgress returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncident) HasScheduledAutoInProgress() bool {
	if o != nil && o.ScheduledAutoInProgress != nil {
		return true
	}

	return false
}

// SetScheduledAutoInProgress gets a reference to the given bool and assigns it to the ScheduledAutoInProgress field.
func (o *PatchPagesPageIdIncidentsIncident) SetScheduledAutoInProgress(v bool) {
	o.ScheduledAutoInProgress = &v
}

// GetScheduledAutoCompleted returns the ScheduledAutoCompleted field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncident) GetScheduledAutoCompleted() bool {
	if o == nil || o.ScheduledAutoCompleted == nil {
		var ret bool
		return ret
	}
	return *o.ScheduledAutoCompleted
}

// GetScheduledAutoCompletedOk returns a tuple with the ScheduledAutoCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncident) GetScheduledAutoCompletedOk() (*bool, bool) {
	if o == nil || o.ScheduledAutoCompleted == nil {
		return nil, false
	}
	return o.ScheduledAutoCompleted, true
}

// HasScheduledAutoCompleted returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncident) HasScheduledAutoCompleted() bool {
	if o != nil && o.ScheduledAutoCompleted != nil {
		return true
	}

	return false
}

// SetScheduledAutoCompleted gets a reference to the given bool and assigns it to the ScheduledAutoCompleted field.
func (o *PatchPagesPageIdIncidentsIncident) SetScheduledAutoCompleted(v bool) {
	o.ScheduledAutoCompleted = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncident) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncident) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncident) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PatchPagesPageIdIncidentsIncident) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetDeliverNotifications returns the DeliverNotifications field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncident) GetDeliverNotifications() bool {
	if o == nil || o.DeliverNotifications == nil {
		var ret bool
		return ret
	}
	return *o.DeliverNotifications
}

// GetDeliverNotificationsOk returns a tuple with the DeliverNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncident) GetDeliverNotificationsOk() (*bool, bool) {
	if o == nil || o.DeliverNotifications == nil {
		return nil, false
	}
	return o.DeliverNotifications, true
}

// HasDeliverNotifications returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncident) HasDeliverNotifications() bool {
	if o != nil && o.DeliverNotifications != nil {
		return true
	}

	return false
}

// SetDeliverNotifications gets a reference to the given bool and assigns it to the DeliverNotifications field.
func (o *PatchPagesPageIdIncidentsIncident) SetDeliverNotifications(v bool) {
	o.DeliverNotifications = &v
}

// GetAutoTransitionDeliverNotificationsAtEnd returns the AutoTransitionDeliverNotificationsAtEnd field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncident) GetAutoTransitionDeliverNotificationsAtEnd() bool {
	if o == nil || o.AutoTransitionDeliverNotificationsAtEnd == nil {
		var ret bool
		return ret
	}
	return *o.AutoTransitionDeliverNotificationsAtEnd
}

// GetAutoTransitionDeliverNotificationsAtEndOk returns a tuple with the AutoTransitionDeliverNotificationsAtEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncident) GetAutoTransitionDeliverNotificationsAtEndOk() (*bool, bool) {
	if o == nil || o.AutoTransitionDeliverNotificationsAtEnd == nil {
		return nil, false
	}
	return o.AutoTransitionDeliverNotificationsAtEnd, true
}

// HasAutoTransitionDeliverNotificationsAtEnd returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncident) HasAutoTransitionDeliverNotificationsAtEnd() bool {
	if o != nil && o.AutoTransitionDeliverNotificationsAtEnd != nil {
		return true
	}

	return false
}

// SetAutoTransitionDeliverNotificationsAtEnd gets a reference to the given bool and assigns it to the AutoTransitionDeliverNotificationsAtEnd field.
func (o *PatchPagesPageIdIncidentsIncident) SetAutoTransitionDeliverNotificationsAtEnd(v bool) {
	o.AutoTransitionDeliverNotificationsAtEnd = &v
}

// GetAutoTransitionDeliverNotificationsAtStart returns the AutoTransitionDeliverNotificationsAtStart field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncident) GetAutoTransitionDeliverNotificationsAtStart() bool {
	if o == nil || o.AutoTransitionDeliverNotificationsAtStart == nil {
		var ret bool
		return ret
	}
	return *o.AutoTransitionDeliverNotificationsAtStart
}

// GetAutoTransitionDeliverNotificationsAtStartOk returns a tuple with the AutoTransitionDeliverNotificationsAtStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncident) GetAutoTransitionDeliverNotificationsAtStartOk() (*bool, bool) {
	if o == nil || o.AutoTransitionDeliverNotificationsAtStart == nil {
		return nil, false
	}
	return o.AutoTransitionDeliverNotificationsAtStart, true
}

// HasAutoTransitionDeliverNotificationsAtStart returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncident) HasAutoTransitionDeliverNotificationsAtStart() bool {
	if o != nil && o.AutoTransitionDeliverNotificationsAtStart != nil {
		return true
	}

	return false
}

// SetAutoTransitionDeliverNotificationsAtStart gets a reference to the given bool and assigns it to the AutoTransitionDeliverNotificationsAtStart field.
func (o *PatchPagesPageIdIncidentsIncident) SetAutoTransitionDeliverNotificationsAtStart(v bool) {
	o.AutoTransitionDeliverNotificationsAtStart = &v
}

// GetAutoTransitionToMaintenanceState returns the AutoTransitionToMaintenanceState field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncident) GetAutoTransitionToMaintenanceState() bool {
	if o == nil || o.AutoTransitionToMaintenanceState == nil {
		var ret bool
		return ret
	}
	return *o.AutoTransitionToMaintenanceState
}

// GetAutoTransitionToMaintenanceStateOk returns a tuple with the AutoTransitionToMaintenanceState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncident) GetAutoTransitionToMaintenanceStateOk() (*bool, bool) {
	if o == nil || o.AutoTransitionToMaintenanceState == nil {
		return nil, false
	}
	return o.AutoTransitionToMaintenanceState, true
}

// HasAutoTransitionToMaintenanceState returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncident) HasAutoTransitionToMaintenanceState() bool {
	if o != nil && o.AutoTransitionToMaintenanceState != nil {
		return true
	}

	return false
}

// SetAutoTransitionToMaintenanceState gets a reference to the given bool and assigns it to the AutoTransitionToMaintenanceState field.
func (o *PatchPagesPageIdIncidentsIncident) SetAutoTransitionToMaintenanceState(v bool) {
	o.AutoTransitionToMaintenanceState = &v
}

// GetAutoTransitionToOperationalState returns the AutoTransitionToOperationalState field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncident) GetAutoTransitionToOperationalState() bool {
	if o == nil || o.AutoTransitionToOperationalState == nil {
		var ret bool
		return ret
	}
	return *o.AutoTransitionToOperationalState
}

// GetAutoTransitionToOperationalStateOk returns a tuple with the AutoTransitionToOperationalState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncident) GetAutoTransitionToOperationalStateOk() (*bool, bool) {
	if o == nil || o.AutoTransitionToOperationalState == nil {
		return nil, false
	}
	return o.AutoTransitionToOperationalState, true
}

// HasAutoTransitionToOperationalState returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncident) HasAutoTransitionToOperationalState() bool {
	if o != nil && o.AutoTransitionToOperationalState != nil {
		return true
	}

	return false
}

// SetAutoTransitionToOperationalState gets a reference to the given bool and assigns it to the AutoTransitionToOperationalState field.
func (o *PatchPagesPageIdIncidentsIncident) SetAutoTransitionToOperationalState(v bool) {
	o.AutoTransitionToOperationalState = &v
}

// GetAutoTweetAtBeginning returns the AutoTweetAtBeginning field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncident) GetAutoTweetAtBeginning() bool {
	if o == nil || o.AutoTweetAtBeginning == nil {
		var ret bool
		return ret
	}
	return *o.AutoTweetAtBeginning
}

// GetAutoTweetAtBeginningOk returns a tuple with the AutoTweetAtBeginning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncident) GetAutoTweetAtBeginningOk() (*bool, bool) {
	if o == nil || o.AutoTweetAtBeginning == nil {
		return nil, false
	}
	return o.AutoTweetAtBeginning, true
}

// HasAutoTweetAtBeginning returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncident) HasAutoTweetAtBeginning() bool {
	if o != nil && o.AutoTweetAtBeginning != nil {
		return true
	}

	return false
}

// SetAutoTweetAtBeginning gets a reference to the given bool and assigns it to the AutoTweetAtBeginning field.
func (o *PatchPagesPageIdIncidentsIncident) SetAutoTweetAtBeginning(v bool) {
	o.AutoTweetAtBeginning = &v
}

// GetAutoTweetOnCompletion returns the AutoTweetOnCompletion field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncident) GetAutoTweetOnCompletion() bool {
	if o == nil || o.AutoTweetOnCompletion == nil {
		var ret bool
		return ret
	}
	return *o.AutoTweetOnCompletion
}

// GetAutoTweetOnCompletionOk returns a tuple with the AutoTweetOnCompletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncident) GetAutoTweetOnCompletionOk() (*bool, bool) {
	if o == nil || o.AutoTweetOnCompletion == nil {
		return nil, false
	}
	return o.AutoTweetOnCompletion, true
}

// HasAutoTweetOnCompletion returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncident) HasAutoTweetOnCompletion() bool {
	if o != nil && o.AutoTweetOnCompletion != nil {
		return true
	}

	return false
}

// SetAutoTweetOnCompletion gets a reference to the given bool and assigns it to the AutoTweetOnCompletion field.
func (o *PatchPagesPageIdIncidentsIncident) SetAutoTweetOnCompletion(v bool) {
	o.AutoTweetOnCompletion = &v
}

// GetAutoTweetOnCreation returns the AutoTweetOnCreation field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncident) GetAutoTweetOnCreation() bool {
	if o == nil || o.AutoTweetOnCreation == nil {
		var ret bool
		return ret
	}
	return *o.AutoTweetOnCreation
}

// GetAutoTweetOnCreationOk returns a tuple with the AutoTweetOnCreation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncident) GetAutoTweetOnCreationOk() (*bool, bool) {
	if o == nil || o.AutoTweetOnCreation == nil {
		return nil, false
	}
	return o.AutoTweetOnCreation, true
}

// HasAutoTweetOnCreation returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncident) HasAutoTweetOnCreation() bool {
	if o != nil && o.AutoTweetOnCreation != nil {
		return true
	}

	return false
}

// SetAutoTweetOnCreation gets a reference to the given bool and assigns it to the AutoTweetOnCreation field.
func (o *PatchPagesPageIdIncidentsIncident) SetAutoTweetOnCreation(v bool) {
	o.AutoTweetOnCreation = &v
}

// GetAutoTweetOneHourBefore returns the AutoTweetOneHourBefore field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncident) GetAutoTweetOneHourBefore() bool {
	if o == nil || o.AutoTweetOneHourBefore == nil {
		var ret bool
		return ret
	}
	return *o.AutoTweetOneHourBefore
}

// GetAutoTweetOneHourBeforeOk returns a tuple with the AutoTweetOneHourBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncident) GetAutoTweetOneHourBeforeOk() (*bool, bool) {
	if o == nil || o.AutoTweetOneHourBefore == nil {
		return nil, false
	}
	return o.AutoTweetOneHourBefore, true
}

// HasAutoTweetOneHourBefore returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncident) HasAutoTweetOneHourBefore() bool {
	if o != nil && o.AutoTweetOneHourBefore != nil {
		return true
	}

	return false
}

// SetAutoTweetOneHourBefore gets a reference to the given bool and assigns it to the AutoTweetOneHourBefore field.
func (o *PatchPagesPageIdIncidentsIncident) SetAutoTweetOneHourBefore(v bool) {
	o.AutoTweetOneHourBefore = &v
}

// GetBackfillDate returns the BackfillDate field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncident) GetBackfillDate() string {
	if o == nil || o.BackfillDate == nil {
		var ret string
		return ret
	}
	return *o.BackfillDate
}

// GetBackfillDateOk returns a tuple with the BackfillDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncident) GetBackfillDateOk() (*string, bool) {
	if o == nil || o.BackfillDate == nil {
		return nil, false
	}
	return o.BackfillDate, true
}

// HasBackfillDate returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncident) HasBackfillDate() bool {
	if o != nil && o.BackfillDate != nil {
		return true
	}

	return false
}

// SetBackfillDate gets a reference to the given string and assigns it to the BackfillDate field.
func (o *PatchPagesPageIdIncidentsIncident) SetBackfillDate(v string) {
	o.BackfillDate = &v
}

// GetBackfilled returns the Backfilled field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncident) GetBackfilled() bool {
	if o == nil || o.Backfilled == nil {
		var ret bool
		return ret
	}
	return *o.Backfilled
}

// GetBackfilledOk returns a tuple with the Backfilled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncident) GetBackfilledOk() (*bool, bool) {
	if o == nil || o.Backfilled == nil {
		return nil, false
	}
	return o.Backfilled, true
}

// HasBackfilled returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncident) HasBackfilled() bool {
	if o != nil && o.Backfilled != nil {
		return true
	}

	return false
}

// SetBackfilled gets a reference to the given bool and assigns it to the Backfilled field.
func (o *PatchPagesPageIdIncidentsIncident) SetBackfilled(v bool) {
	o.Backfilled = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncident) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncident) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncident) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *PatchPagesPageIdIncidentsIncident) SetBody(v string) {
	o.Body = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncident) GetComponents() PostPagesPageIdIncidentsIncidentComponents {
	if o == nil || o.Components == nil {
		var ret PostPagesPageIdIncidentsIncidentComponents
		return ret
	}
	return *o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncident) GetComponentsOk() (*PostPagesPageIdIncidentsIncidentComponents, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncident) HasComponents() bool {
	if o != nil && o.Components != nil {
		return true
	}

	return false
}

// SetComponents gets a reference to the given PostPagesPageIdIncidentsIncidentComponents and assigns it to the Components field.
func (o *PatchPagesPageIdIncidentsIncident) SetComponents(v PostPagesPageIdIncidentsIncidentComponents) {
	o.Components = &v
}

// GetComponentIds returns the ComponentIds field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncident) GetComponentIds() []string {
	if o == nil || o.ComponentIds == nil {
		var ret []string
		return ret
	}
	return *o.ComponentIds
}

// GetComponentIdsOk returns a tuple with the ComponentIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncident) GetComponentIdsOk() (*[]string, bool) {
	if o == nil || o.ComponentIds == nil {
		return nil, false
	}
	return o.ComponentIds, true
}

// HasComponentIds returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncident) HasComponentIds() bool {
	if o != nil && o.ComponentIds != nil {
		return true
	}

	return false
}

// SetComponentIds gets a reference to the given []string and assigns it to the ComponentIds field.
func (o *PatchPagesPageIdIncidentsIncident) SetComponentIds(v []string) {
	o.ComponentIds = &v
}

// GetScheduledAutoTransition returns the ScheduledAutoTransition field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncident) GetScheduledAutoTransition() bool {
	if o == nil || o.ScheduledAutoTransition == nil {
		var ret bool
		return ret
	}
	return *o.ScheduledAutoTransition
}

// GetScheduledAutoTransitionOk returns a tuple with the ScheduledAutoTransition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncident) GetScheduledAutoTransitionOk() (*bool, bool) {
	if o == nil || o.ScheduledAutoTransition == nil {
		return nil, false
	}
	return o.ScheduledAutoTransition, true
}

// HasScheduledAutoTransition returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncident) HasScheduledAutoTransition() bool {
	if o != nil && o.ScheduledAutoTransition != nil {
		return true
	}

	return false
}

// SetScheduledAutoTransition gets a reference to the given bool and assigns it to the ScheduledAutoTransition field.
func (o *PatchPagesPageIdIncidentsIncident) SetScheduledAutoTransition(v bool) {
	o.ScheduledAutoTransition = &v
}

func (o PatchPagesPageIdIncidentsIncident) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ImpactOverride != nil {
		toSerialize["impact_override"] = o.ImpactOverride
	}
	if o.ScheduledFor != nil {
		toSerialize["scheduled_for"] = o.ScheduledFor
	}
	if o.ScheduledUntil != nil {
		toSerialize["scheduled_until"] = o.ScheduledUntil
	}
	if o.ScheduledRemindPrior != nil {
		toSerialize["scheduled_remind_prior"] = o.ScheduledRemindPrior
	}
	if o.ScheduledAutoInProgress != nil {
		toSerialize["scheduled_auto_in_progress"] = o.ScheduledAutoInProgress
	}
	if o.ScheduledAutoCompleted != nil {
		toSerialize["scheduled_auto_completed"] = o.ScheduledAutoCompleted
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.DeliverNotifications != nil {
		toSerialize["deliver_notifications"] = o.DeliverNotifications
	}
	if o.AutoTransitionDeliverNotificationsAtEnd != nil {
		toSerialize["auto_transition_deliver_notifications_at_end"] = o.AutoTransitionDeliverNotificationsAtEnd
	}
	if o.AutoTransitionDeliverNotificationsAtStart != nil {
		toSerialize["auto_transition_deliver_notifications_at_start"] = o.AutoTransitionDeliverNotificationsAtStart
	}
	if o.AutoTransitionToMaintenanceState != nil {
		toSerialize["auto_transition_to_maintenance_state"] = o.AutoTransitionToMaintenanceState
	}
	if o.AutoTransitionToOperationalState != nil {
		toSerialize["auto_transition_to_operational_state"] = o.AutoTransitionToOperationalState
	}
	if o.AutoTweetAtBeginning != nil {
		toSerialize["auto_tweet_at_beginning"] = o.AutoTweetAtBeginning
	}
	if o.AutoTweetOnCompletion != nil {
		toSerialize["auto_tweet_on_completion"] = o.AutoTweetOnCompletion
	}
	if o.AutoTweetOnCreation != nil {
		toSerialize["auto_tweet_on_creation"] = o.AutoTweetOnCreation
	}
	if o.AutoTweetOneHourBefore != nil {
		toSerialize["auto_tweet_one_hour_before"] = o.AutoTweetOneHourBefore
	}
	if o.BackfillDate != nil {
		toSerialize["backfill_date"] = o.BackfillDate
	}
	if o.Backfilled != nil {
		toSerialize["backfilled"] = o.Backfilled
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.ComponentIds != nil {
		toSerialize["component_ids"] = o.ComponentIds
	}
	if o.ScheduledAutoTransition != nil {
		toSerialize["scheduled_auto_transition"] = o.ScheduledAutoTransition
	}
	return json.Marshal(toSerialize)
}

type NullablePatchPagesPageIdIncidentsIncident struct {
	value *PatchPagesPageIdIncidentsIncident
	isSet bool
}

func (v NullablePatchPagesPageIdIncidentsIncident) Get() *PatchPagesPageIdIncidentsIncident {
	return v.value
}

func (v *NullablePatchPagesPageIdIncidentsIncident) Set(val *PatchPagesPageIdIncidentsIncident) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchPagesPageIdIncidentsIncident) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchPagesPageIdIncidentsIncident) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchPagesPageIdIncidentsIncident(val *PatchPagesPageIdIncidentsIncident) *NullablePatchPagesPageIdIncidentsIncident {
	return &NullablePatchPagesPageIdIncidentsIncident{value: val, isSet: true}
}

func (v NullablePatchPagesPageIdIncidentsIncident) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchPagesPageIdIncidentsIncident) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


