/*
 * Statuspage API
 *
 * # Code of Conduct Please don't abuse the API, and please report all feature requests and issues to https://help.statuspage.io/help/contact-us-30  # Rate Limiting Each API token is limited to 1 request / second as measured on a 60 second rolling window. To get this limit increased, please contact us at https://help.statuspage.io/help/contact-us-30  # Basics  ## HTTPS It's required  ## URL Prefix In order to maintain version integrity into the future, the API is versioned. All calls currently begin with the following prefix:    https://api.statuspage.io/v1/  ## RESTful Interface Wherever possible, the API seeks to implement repeatable patterns with logical, representative URLs and descriptive HTTP verbs. Below are some examples and conventions you will see throughout the documentation.  * Collections are buckets: https://api.statuspage.io/v1/pages/asdf123/incidents.json * Elements have unique IDs: https://api.statuspage.io/v1/pages/asdf123/incidents/jklm456.json * GET will retrieve information about a collection/element * POST will create an element in a collection * PATCH will update a single element * PUT will replace a single element in a collection (rarely used) * DELETE will destroy a single element  ## Sending Data Information can be sent in the body as form urlencoded or JSON, but make sure the Content-Type header matches the body structure or the server gremlins will be angry.  All examples are provided in JSON format, however they can easily be converted to form encoding if required.  Some examples of how to convert things are below:      // JSON     {       \"incident\": {         \"name\": \"test incident\",         \"components\": [\"8kbf7d35c070\", \"vtnh60py4yd7\"]       }     }      // Form Encoded (using curl as an example):     curl -X POST https://api.statuspage.io/v1/example \\       -d \"incident[name]=test incident\" \\       -d \"incident[components][]=8kbf7d35c070\" \\       -d \"incident[components][]=vtnh60py4yd7\"  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)

type PatchPagesPageIdIncidentsIncident struct {
	// Incident Name
	Name string `json:"name,omitempty"`
	// The incident status. For realtime incidents, valid values are investigating, identified, monitoring, and resolved. For scheduled incidents, valid values are scheduled, in_progress, verifying, and completed.
	Status string `json:"status,omitempty"`
	// value to override calculated impact value
	ImpactOverride string `json:"impact_override,omitempty"`
	// The timestamp the incident is scheduled for.
	ScheduledFor time.Time `json:"scheduled_for,omitempty"`
	// The timestamp the incident is scheduled until.
	ScheduledUntil time.Time `json:"scheduled_until,omitempty"`
	// Controls whether to remind subscribers prior to scheduled incidents.
	ScheduledRemindPrior bool `json:"scheduled_remind_prior,omitempty"`
	// Controls whether the incident is scheduled to automatically change to in progress.
	ScheduledAutoInProgress bool `json:"scheduled_auto_in_progress,omitempty"`
	// Controls whether the incident is scheduled to automatically change to complete.
	ScheduledAutoCompleted bool `json:"scheduled_auto_completed,omitempty"`
	// Attach a json object to the incident. All top-level values in the object must also be objects.
	Metadata *interface{} `json:"metadata,omitempty"`
	// Deliver notifications to subscribers if this is true. If this is false, create an incident without notifying customers.
	DeliverNotifications bool `json:"deliver_notifications,omitempty"`
	// Controls whether send notification when scheduled maintenances auto transition to completed.
	AutoTransitionDeliverNotificationsAtEnd bool `json:"auto_transition_deliver_notifications_at_end,omitempty"`
	// Controls whether send notification when scheduled maintenances auto transition to started.
	AutoTransitionDeliverNotificationsAtStart bool `json:"auto_transition_deliver_notifications_at_start,omitempty"`
	// Controls whether change components status to under_maintenance once scheduled maintenance is in progress.
	AutoTransitionToMaintenanceState bool `json:"auto_transition_to_maintenance_state,omitempty"`
	// Controls whether change components status to operational once scheduled maintenance completes.
	AutoTransitionToOperationalState bool `json:"auto_transition_to_operational_state,omitempty"`
	// Controls whether tweet automatically when scheduled maintenance starts.
	AutoTweetAtBeginning bool `json:"auto_tweet_at_beginning,omitempty"`
	// Controls whether tweet automatically when scheduled maintenance completes.
	AutoTweetOnCompletion bool `json:"auto_tweet_on_completion,omitempty"`
	// Controls whether tweet automatically when scheduled maintenance is created.
	AutoTweetOnCreation bool `json:"auto_tweet_on_creation,omitempty"`
	// Controls whether tweet automatically one hour before scheduled maintenance starts.
	AutoTweetOneHourBefore bool `json:"auto_tweet_one_hour_before,omitempty"`
	// TimeStamp when incident was backfilled.
	BackfillDate string `json:"backfill_date,omitempty"`
	// Controls whether incident is backfilled. If true, components cannot be specified.
	Backfilled bool `json:"backfilled,omitempty"`
	// The initial message, created as the first incident update.
	Body string `json:"body,omitempty"`
	Components *PostPagesPageIdIncidentsIncidentComponents `json:"components,omitempty"`
	// List of component_ids affected by this incident
	ComponentIds []string `json:"component_ids,omitempty"`
	// Same as :scheduled_auto_transition_in_progress. Controls whether the incident is scheduled to automatically change to in progress.
	ScheduledAutoTransition bool `json:"scheduled_auto_transition,omitempty"`
}
