# Go API client for openapi

# Code of Conduct Please don't abuse the API, and please report all feature requests and issues to https://help.statuspage.io/help/contact-us-30  # Rate Limiting Each API token is limited to 1 request / second as measured on a 60 second rolling window. To get this limit increased, please contact us at https://help.statuspage.io/help/contact-us-30  # Basics  ## HTTPS It's required  ## URL Prefix In order to maintain version integrity into the future, the API is versioned. All calls currently begin with the following prefix:    https://api.statuspage.io/v1/  ## RESTful Interface Wherever possible, the API seeks to implement repeatable patterns with logical, representative URLs and descriptive HTTP verbs. Below are some examples and conventions you will see throughout the documentation.  * Collections are buckets: https://api.statuspage.io/v1/pages/asdf123/incidents.json * Elements have unique IDs: https://api.statuspage.io/v1/pages/asdf123/incidents/jklm456.json * GET will retrieve information about a collection/element * POST will create an element in a collection * PATCH will update a single element * PUT will replace a single element in a collection (rarely used) * DELETE will destroy a single element  ## Sending Data Information can be sent in the body as form urlencoded or JSON, but make sure the Content-Type header matches the body structure or the server gremlins will be angry.  All examples are provided in JSON format, however they can easily be converted to form encoding if required.  Some examples of how to convert things are below:      // JSON     {       \"incident\": {         \"name\": \"test incident\",         \"components\": [\"8kbf7d35c070\", \"vtnh60py4yd7\"]       }     }      // Form Encoded (using curl as an example):     curl -X POST https://api.statuspage.io/v1/example \\       -d \"incident[name]=test incident\" \\       -d \"incident[components][]=8kbf7d35c070\" \\       -d \"incident[components][]=vtnh60py4yd7\"  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://help.statuspage.io/help/contact-us-30](https://help.statuspage.io/help/contact-us-30)

## Installation

Install the following dependencies:
```
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:
```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.statuspage.io/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ComponentGroupsApi* | [**DeletePagesPageIdComponentGroupsId**](docs/ComponentGroupsApi.md#deletepagespageidcomponentgroupsid) | **Delete** /pages/{page_id}/component-groups/{id} | Delete a component group
*ComponentGroupsApi* | [**GetPagesPageIdComponentGroups**](docs/ComponentGroupsApi.md#getpagespageidcomponentgroups) | **Get** /pages/{page_id}/component-groups | Get a list of component groups
*ComponentGroupsApi* | [**GetPagesPageIdComponentGroupsId**](docs/ComponentGroupsApi.md#getpagespageidcomponentgroupsid) | **Get** /pages/{page_id}/component-groups/{id} | Get a component group
*ComponentGroupsApi* | [**GetPagesPageIdComponentGroupsIdUptime**](docs/ComponentGroupsApi.md#getpagespageidcomponentgroupsiduptime) | **Get** /pages/{page_id}/component-groups/{id}/uptime | Get uptime data for a component group
*ComponentGroupsApi* | [**PatchPagesPageIdComponentGroupsId**](docs/ComponentGroupsApi.md#patchpagespageidcomponentgroupsid) | **Patch** /pages/{page_id}/component-groups/{id} | Update a component group
*ComponentGroupsApi* | [**PostPagesPageIdComponentGroups**](docs/ComponentGroupsApi.md#postpagespageidcomponentgroups) | **Post** /pages/{page_id}/component-groups | Create a component group
*ComponentGroupsApi* | [**PutPagesPageIdComponentGroupsId**](docs/ComponentGroupsApi.md#putpagespageidcomponentgroupsid) | **Put** /pages/{page_id}/component-groups/{id} | Update a component group
*ComponentsApi* | [**DeletePagesPageIdComponentsComponentId**](docs/ComponentsApi.md#deletepagespageidcomponentscomponentid) | **Delete** /pages/{page_id}/components/{component_id} | Delete a component
*ComponentsApi* | [**DeletePagesPageIdComponentsComponentIdPageAccessGroups**](docs/ComponentsApi.md#deletepagespageidcomponentscomponentidpageaccessgroups) | **Delete** /pages/{page_id}/components/{component_id}/page_access_groups | Remove page access groups from a component
*ComponentsApi* | [**DeletePagesPageIdComponentsComponentIdPageAccessUsers**](docs/ComponentsApi.md#deletepagespageidcomponentscomponentidpageaccessusers) | **Delete** /pages/{page_id}/components/{component_id}/page_access_users | Remove page access users from component
*ComponentsApi* | [**GetPagesPageIdComponents**](docs/ComponentsApi.md#getpagespageidcomponents) | **Get** /pages/{page_id}/components | Get a list of components
*ComponentsApi* | [**GetPagesPageIdComponentsComponentId**](docs/ComponentsApi.md#getpagespageidcomponentscomponentid) | **Get** /pages/{page_id}/components/{component_id} | Get a component
*ComponentsApi* | [**GetPagesPageIdComponentsComponentIdUptime**](docs/ComponentsApi.md#getpagespageidcomponentscomponentiduptime) | **Get** /pages/{page_id}/components/{component_id}/uptime | Get uptime data for a component
*ComponentsApi* | [**PatchPagesPageIdComponentsComponentId**](docs/ComponentsApi.md#patchpagespageidcomponentscomponentid) | **Patch** /pages/{page_id}/components/{component_id} | Update a component
*ComponentsApi* | [**PostPagesPageIdComponents**](docs/ComponentsApi.md#postpagespageidcomponents) | **Post** /pages/{page_id}/components | Create a component
*ComponentsApi* | [**PostPagesPageIdComponentsComponentIdPageAccessGroups**](docs/ComponentsApi.md#postpagespageidcomponentscomponentidpageaccessgroups) | **Post** /pages/{page_id}/components/{component_id}/page_access_groups | Add page access groups to a component
*ComponentsApi* | [**PostPagesPageIdComponentsComponentIdPageAccessUsers**](docs/ComponentsApi.md#postpagespageidcomponentscomponentidpageaccessusers) | **Post** /pages/{page_id}/components/{component_id}/page_access_users | Add page access users to a component
*ComponentsApi* | [**PutPagesPageIdComponentsComponentId**](docs/ComponentsApi.md#putpagespageidcomponentscomponentid) | **Put** /pages/{page_id}/components/{component_id} | Update a component
*IncidentPostmortemApi* | [**DeletePagesPageIdIncidentsIncidentIdPostmortem**](docs/IncidentPostmortemApi.md#deletepagespageidincidentsincidentidpostmortem) | **Delete** /pages/{page_id}/incidents/{incident_id}/postmortem | Delete Postmortem
*IncidentPostmortemApi* | [**GetPagesPageIdIncidentsIncidentIdPostmortem**](docs/IncidentPostmortemApi.md#getpagespageidincidentsincidentidpostmortem) | **Get** /pages/{page_id}/incidents/{incident_id}/postmortem | Get Postmortem
*IncidentPostmortemApi* | [**PutPagesPageIdIncidentsIncidentIdPostmortem**](docs/IncidentPostmortemApi.md#putpagespageidincidentsincidentidpostmortem) | **Put** /pages/{page_id}/incidents/{incident_id}/postmortem | Create Postmortem
*IncidentPostmortemApi* | [**PutPagesPageIdIncidentsIncidentIdPostmortemPublish**](docs/IncidentPostmortemApi.md#putpagespageidincidentsincidentidpostmortempublish) | **Put** /pages/{page_id}/incidents/{incident_id}/postmortem/publish | Publish Postmortem
*IncidentPostmortemApi* | [**PutPagesPageIdIncidentsIncidentIdPostmortemRevert**](docs/IncidentPostmortemApi.md#putpagespageidincidentsincidentidpostmortemrevert) | **Put** /pages/{page_id}/incidents/{incident_id}/postmortem/revert | Revert Postmortem
*IncidentSubscribersApi* | [**DeletePagesPageIdIncidentsIncidentIdSubscribersSubscriberId**](docs/IncidentSubscribersApi.md#deletepagespageidincidentsincidentidsubscriberssubscriberid) | **Delete** /pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id} | Unsubscribe an incident subscriber
*IncidentSubscribersApi* | [**GetPagesPageIdIncidentsIncidentIdSubscribers**](docs/IncidentSubscribersApi.md#getpagespageidincidentsincidentidsubscribers) | **Get** /pages/{page_id}/incidents/{incident_id}/subscribers | Get a list of incident subscribers
*IncidentSubscribersApi* | [**GetPagesPageIdIncidentsIncidentIdSubscribersSubscriberId**](docs/IncidentSubscribersApi.md#getpagespageidincidentsincidentidsubscriberssubscriberid) | **Get** /pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id} | Get an incident subscriber
*IncidentSubscribersApi* | [**PostPagesPageIdIncidentsIncidentIdSubscribers**](docs/IncidentSubscribersApi.md#postpagespageidincidentsincidentidsubscribers) | **Post** /pages/{page_id}/incidents/{incident_id}/subscribers | Create an incident subscriber
*IncidentSubscribersApi* | [**PostPagesPageIdIncidentsIncidentIdSubscribersSubscriberIdResendConfirmation**](docs/IncidentSubscribersApi.md#postpagespageidincidentsincidentidsubscriberssubscriberidresendconfirmation) | **Post** /pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}/resend_confirmation | Resend confirmation to an incident subscriber
*IncidentUpdatesApi* | [**PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId**](docs/IncidentUpdatesApi.md#patchpagespageidincidentsincidentidincidentupdatesincidentupdateid) | **Patch** /pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id} | Update a previous incident update
*IncidentUpdatesApi* | [**PutPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId**](docs/IncidentUpdatesApi.md#putpagespageidincidentsincidentidincidentupdatesincidentupdateid) | **Put** /pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id} | Update a previous incident update
*IncidentsApi* | [**DeletePagesPageIdIncidentsIncidentId**](docs/IncidentsApi.md#deletepagespageidincidentsincidentid) | **Delete** /pages/{page_id}/incidents/{incident_id} | Delete an incident
*IncidentsApi* | [**GetPagesPageIdIncidents**](docs/IncidentsApi.md#getpagespageidincidents) | **Get** /pages/{page_id}/incidents | Get a list of incidents
*IncidentsApi* | [**GetPagesPageIdIncidentsActiveMaintenance**](docs/IncidentsApi.md#getpagespageidincidentsactivemaintenance) | **Get** /pages/{page_id}/incidents/active_maintenance | Get a list of active maintenances
*IncidentsApi* | [**GetPagesPageIdIncidentsIncidentId**](docs/IncidentsApi.md#getpagespageidincidentsincidentid) | **Get** /pages/{page_id}/incidents/{incident_id} | Get an incident
*IncidentsApi* | [**GetPagesPageIdIncidentsScheduled**](docs/IncidentsApi.md#getpagespageidincidentsscheduled) | **Get** /pages/{page_id}/incidents/scheduled | Get a list of scheduled incidents
*IncidentsApi* | [**GetPagesPageIdIncidentsUnresolved**](docs/IncidentsApi.md#getpagespageidincidentsunresolved) | **Get** /pages/{page_id}/incidents/unresolved | Get a list of unresolved incidents
*IncidentsApi* | [**GetPagesPageIdIncidentsUpcoming**](docs/IncidentsApi.md#getpagespageidincidentsupcoming) | **Get** /pages/{page_id}/incidents/upcoming | Get a list of upcoming incidents
*IncidentsApi* | [**PatchPagesPageIdIncidentsIncidentId**](docs/IncidentsApi.md#patchpagespageidincidentsincidentid) | **Patch** /pages/{page_id}/incidents/{incident_id} | Update an incident
*IncidentsApi* | [**PostPagesPageIdIncidents**](docs/IncidentsApi.md#postpagespageidincidents) | **Post** /pages/{page_id}/incidents | Create an incident
*IncidentsApi* | [**PutPagesPageIdIncidentsIncidentId**](docs/IncidentsApi.md#putpagespageidincidentsincidentid) | **Put** /pages/{page_id}/incidents/{incident_id} | Update an incident
*MetricProvidersApi* | [**DeletePagesPageIdMetricsProvidersMetricsProviderId**](docs/MetricProvidersApi.md#deletepagespageidmetricsprovidersmetricsproviderid) | **Delete** /pages/{page_id}/metrics_providers/{metrics_provider_id} | Delete a metric provider
*MetricProvidersApi* | [**GetPagesPageIdMetricsProviders**](docs/MetricProvidersApi.md#getpagespageidmetricsproviders) | **Get** /pages/{page_id}/metrics_providers | Get a list of metric providers
*MetricProvidersApi* | [**GetPagesPageIdMetricsProvidersMetricsProviderId**](docs/MetricProvidersApi.md#getpagespageidmetricsprovidersmetricsproviderid) | **Get** /pages/{page_id}/metrics_providers/{metrics_provider_id} | Get a metric provider
*MetricProvidersApi* | [**GetPagesPageIdMetricsProvidersMetricsProviderIdMetrics**](docs/MetricProvidersApi.md#getpagespageidmetricsprovidersmetricsprovideridmetrics) | **Get** /pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics | List metrics for a metric provider
*MetricProvidersApi* | [**PatchPagesPageIdMetricsProvidersMetricsProviderId**](docs/MetricProvidersApi.md#patchpagespageidmetricsprovidersmetricsproviderid) | **Patch** /pages/{page_id}/metrics_providers/{metrics_provider_id} | Update a metric provider
*MetricProvidersApi* | [**PostPagesPageIdMetricsProviders**](docs/MetricProvidersApi.md#postpagespageidmetricsproviders) | **Post** /pages/{page_id}/metrics_providers | Create a metric provider
*MetricProvidersApi* | [**PostPagesPageIdMetricsProvidersMetricsProviderIdMetrics**](docs/MetricProvidersApi.md#postpagespageidmetricsprovidersmetricsprovideridmetrics) | **Post** /pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics | Create a metric for a metric provider
*MetricProvidersApi* | [**PutPagesPageIdMetricsProvidersMetricsProviderId**](docs/MetricProvidersApi.md#putpagespageidmetricsprovidersmetricsproviderid) | **Put** /pages/{page_id}/metrics_providers/{metrics_provider_id} | Update a metric provider
*MetricsApi* | [**DeletePagesPageIdMetricsMetricId**](docs/MetricsApi.md#deletepagespageidmetricsmetricid) | **Delete** /pages/{page_id}/metrics/{metric_id} | Delete a metric
*MetricsApi* | [**DeletePagesPageIdMetricsMetricIdData**](docs/MetricsApi.md#deletepagespageidmetricsmetriciddata) | **Delete** /pages/{page_id}/metrics/{metric_id}/data | Reset data for a metric
*MetricsApi* | [**GetPagesPageIdMetrics**](docs/MetricsApi.md#getpagespageidmetrics) | **Get** /pages/{page_id}/metrics | Get a list of metrics
*MetricsApi* | [**GetPagesPageIdMetricsMetricId**](docs/MetricsApi.md#getpagespageidmetricsmetricid) | **Get** /pages/{page_id}/metrics/{metric_id} | Get a metric
*MetricsApi* | [**GetPagesPageIdMetricsProvidersMetricsProviderIdMetrics**](docs/MetricsApi.md#getpagespageidmetricsprovidersmetricsprovideridmetrics) | **Get** /pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics | List metrics for a metric provider
*MetricsApi* | [**PatchPagesPageIdMetricsMetricId**](docs/MetricsApi.md#patchpagespageidmetricsmetricid) | **Patch** /pages/{page_id}/metrics/{metric_id} | Update a metric
*MetricsApi* | [**PostPagesPageIdMetricsData**](docs/MetricsApi.md#postpagespageidmetricsdata) | **Post** /pages/{page_id}/metrics/data | Add data points to metrics
*MetricsApi* | [**PostPagesPageIdMetricsMetricIdData**](docs/MetricsApi.md#postpagespageidmetricsmetriciddata) | **Post** /pages/{page_id}/metrics/{metric_id}/data | Add data to a metric
*MetricsApi* | [**PostPagesPageIdMetricsProvidersMetricsProviderIdMetrics**](docs/MetricsApi.md#postpagespageidmetricsprovidersmetricsprovideridmetrics) | **Post** /pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics | Create a metric for a metric provider
*MetricsApi* | [**PutPagesPageIdMetricsMetricId**](docs/MetricsApi.md#putpagespageidmetricsmetricid) | **Put** /pages/{page_id}/metrics/{metric_id} | Update a metric
*PageAccessGroupComponentsApi* | [**DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents**](docs/PageAccessGroupComponentsApi.md#deletepagespageidpageaccessgroupspageaccessgroupidcomponents) | **Delete** /pages/{page_id}/page_access_groups/{page_access_group_id}/components | Delete components for a page access group
*PageAccessGroupComponentsApi* | [**DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponentsComponentId**](docs/PageAccessGroupComponentsApi.md#deletepagespageidpageaccessgroupspageaccessgroupidcomponentscomponentid) | **Delete** /pages/{page_id}/page_access_groups/{page_access_group_id}/components/{component_id} | Remove a component from a page access group
*PageAccessGroupComponentsApi* | [**GetPagesPageIdPageAccessGroupsPageAccessGroupIdComponents**](docs/PageAccessGroupComponentsApi.md#getpagespageidpageaccessgroupspageaccessgroupidcomponents) | **Get** /pages/{page_id}/page_access_groups/{page_access_group_id}/components | List components for a page access group
*PageAccessGroupComponentsApi* | [**PatchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents**](docs/PageAccessGroupComponentsApi.md#patchpagespageidpageaccessgroupspageaccessgroupidcomponents) | **Patch** /pages/{page_id}/page_access_groups/{page_access_group_id}/components | Add components to page access group
*PageAccessGroupComponentsApi* | [**PostPagesPageIdPageAccessGroupsPageAccessGroupIdComponents**](docs/PageAccessGroupComponentsApi.md#postpagespageidpageaccessgroupspageaccessgroupidcomponents) | **Post** /pages/{page_id}/page_access_groups/{page_access_group_id}/components | Replace components for a page access group
*PageAccessGroupComponentsApi* | [**PutPagesPageIdPageAccessGroupsPageAccessGroupIdComponents**](docs/PageAccessGroupComponentsApi.md#putpagespageidpageaccessgroupspageaccessgroupidcomponents) | **Put** /pages/{page_id}/page_access_groups/{page_access_group_id}/components | Add components to page access group
*PageAccessGroupsApi* | [**DeletePagesPageIdPageAccessGroupsPageAccessGroupId**](docs/PageAccessGroupsApi.md#deletepagespageidpageaccessgroupspageaccessgroupid) | **Delete** /pages/{page_id}/page_access_groups/{page_access_group_id} | Remove a page access group
*PageAccessGroupsApi* | [**GetPagesPageIdPageAccessGroups**](docs/PageAccessGroupsApi.md#getpagespageidpageaccessgroups) | **Get** /pages/{page_id}/page_access_groups | Get a list of page access groups
*PageAccessGroupsApi* | [**GetPagesPageIdPageAccessGroupsPageAccessGroupId**](docs/PageAccessGroupsApi.md#getpagespageidpageaccessgroupspageaccessgroupid) | **Get** /pages/{page_id}/page_access_groups/{page_access_group_id} | Get a page access group
*PageAccessGroupsApi* | [**PatchPagesPageIdPageAccessGroupsPageAccessGroupId**](docs/PageAccessGroupsApi.md#patchpagespageidpageaccessgroupspageaccessgroupid) | **Patch** /pages/{page_id}/page_access_groups/{page_access_group_id} | Update a page access group
*PageAccessGroupsApi* | [**PostPagesPageIdPageAccessGroups**](docs/PageAccessGroupsApi.md#postpagespageidpageaccessgroups) | **Post** /pages/{page_id}/page_access_groups | Create a page access group
*PageAccessGroupsApi* | [**PutPagesPageIdPageAccessGroupsPageAccessGroupId**](docs/PageAccessGroupsApi.md#putpagespageidpageaccessgroupspageaccessgroupid) | **Put** /pages/{page_id}/page_access_groups/{page_access_group_id} | Update a page access group
*PageAccessUserComponentsApi* | [**DeletePagesPageIdPageAccessUsersPageAccessUserIdComponents**](docs/PageAccessUserComponentsApi.md#deletepagespageidpageaccessuserspageaccessuseridcomponents) | **Delete** /pages/{page_id}/page_access_users/{page_access_user_id}/components | Remove components for page access user
*PageAccessUserComponentsApi* | [**DeletePagesPageIdPageAccessUsersPageAccessUserIdComponentsComponentId**](docs/PageAccessUserComponentsApi.md#deletepagespageidpageaccessuserspageaccessuseridcomponentscomponentid) | **Delete** /pages/{page_id}/page_access_users/{page_access_user_id}/components/{component_id} | Remove component for page access user
*PageAccessUserComponentsApi* | [**GetPagesPageIdPageAccessUsersPageAccessUserIdComponents**](docs/PageAccessUserComponentsApi.md#getpagespageidpageaccessuserspageaccessuseridcomponents) | **Get** /pages/{page_id}/page_access_users/{page_access_user_id}/components | Get components for page access user
*PageAccessUserComponentsApi* | [**PatchPagesPageIdPageAccessUsersPageAccessUserIdComponents**](docs/PageAccessUserComponentsApi.md#patchpagespageidpageaccessuserspageaccessuseridcomponents) | **Patch** /pages/{page_id}/page_access_users/{page_access_user_id}/components | Add components for page access user
*PageAccessUserComponentsApi* | [**PostPagesPageIdPageAccessUsersPageAccessUserIdComponents**](docs/PageAccessUserComponentsApi.md#postpagespageidpageaccessuserspageaccessuseridcomponents) | **Post** /pages/{page_id}/page_access_users/{page_access_user_id}/components | Replace components for page access user
*PageAccessUserComponentsApi* | [**PutPagesPageIdPageAccessUsersPageAccessUserIdComponents**](docs/PageAccessUserComponentsApi.md#putpagespageidpageaccessuserspageaccessuseridcomponents) | **Put** /pages/{page_id}/page_access_users/{page_access_user_id}/components | Add components for page access user
*PageAccessUserMetricsApi* | [**DeletePagesPageIdPageAccessUsersPageAccessUserIdMetrics**](docs/PageAccessUserMetricsApi.md#deletepagespageidpageaccessuserspageaccessuseridmetrics) | **Delete** /pages/{page_id}/page_access_users/{page_access_user_id}/metrics | Delete metrics for page access user
*PageAccessUserMetricsApi* | [**DeletePagesPageIdPageAccessUsersPageAccessUserIdMetricsMetricId**](docs/PageAccessUserMetricsApi.md#deletepagespageidpageaccessuserspageaccessuseridmetricsmetricid) | **Delete** /pages/{page_id}/page_access_users/{page_access_user_id}/metrics/{metric_id} | Delete metric for page access user
*PageAccessUserMetricsApi* | [**GetPagesPageIdPageAccessUsersPageAccessUserIdMetrics**](docs/PageAccessUserMetricsApi.md#getpagespageidpageaccessuserspageaccessuseridmetrics) | **Get** /pages/{page_id}/page_access_users/{page_access_user_id}/metrics | Get metrics for page access user
*PageAccessUserMetricsApi* | [**PatchPagesPageIdPageAccessUsersPageAccessUserIdMetrics**](docs/PageAccessUserMetricsApi.md#patchpagespageidpageaccessuserspageaccessuseridmetrics) | **Patch** /pages/{page_id}/page_access_users/{page_access_user_id}/metrics | Add metrics for page access user
*PageAccessUserMetricsApi* | [**PostPagesPageIdPageAccessUsersPageAccessUserIdMetrics**](docs/PageAccessUserMetricsApi.md#postpagespageidpageaccessuserspageaccessuseridmetrics) | **Post** /pages/{page_id}/page_access_users/{page_access_user_id}/metrics | Replace metrics for page access user
*PageAccessUserMetricsApi* | [**PutPagesPageIdPageAccessUsersPageAccessUserIdMetrics**](docs/PageAccessUserMetricsApi.md#putpagespageidpageaccessuserspageaccessuseridmetrics) | **Put** /pages/{page_id}/page_access_users/{page_access_user_id}/metrics | Add metrics for page access user
*PageAccessUsersApi* | [**DeletePagesPageIdPageAccessUsersPageAccessUserId**](docs/PageAccessUsersApi.md#deletepagespageidpageaccessuserspageaccessuserid) | **Delete** /pages/{page_id}/page_access_users/{page_access_user_id} | Delete page access user
*PageAccessUsersApi* | [**GetPagesPageIdPageAccessUsers**](docs/PageAccessUsersApi.md#getpagespageidpageaccessusers) | **Get** /pages/{page_id}/page_access_users | Get a list of page access users
*PageAccessUsersApi* | [**GetPagesPageIdPageAccessUsersPageAccessUserId**](docs/PageAccessUsersApi.md#getpagespageidpageaccessuserspageaccessuserid) | **Get** /pages/{page_id}/page_access_users/{page_access_user_id} | Get page access user
*PageAccessUsersApi* | [**PatchPagesPageIdPageAccessUsersPageAccessUserId**](docs/PageAccessUsersApi.md#patchpagespageidpageaccessuserspageaccessuserid) | **Patch** /pages/{page_id}/page_access_users/{page_access_user_id} | Update page access user
*PageAccessUsersApi* | [**PostPagesPageIdPageAccessUsers**](docs/PageAccessUsersApi.md#postpagespageidpageaccessusers) | **Post** /pages/{page_id}/page_access_users | Add a page access user
*PageAccessUsersApi* | [**PutPagesPageIdPageAccessUsersPageAccessUserId**](docs/PageAccessUsersApi.md#putpagespageidpageaccessuserspageaccessuserid) | **Put** /pages/{page_id}/page_access_users/{page_access_user_id} | Update page access user
*PagesApi* | [**GetPages**](docs/PagesApi.md#getpages) | **Get** /pages | Get a list of pages
*PagesApi* | [**GetPagesPageId**](docs/PagesApi.md#getpagespageid) | **Get** /pages/{page_id} | Get a page
*PagesApi* | [**PatchPagesPageId**](docs/PagesApi.md#patchpagespageid) | **Patch** /pages/{page_id} | Update a page
*PagesApi* | [**PutPagesPageId**](docs/PagesApi.md#putpagespageid) | **Put** /pages/{page_id} | Update a page
*StatusEmbedConfigApi* | [**GetPagesPageIdStatusEmbedConfig**](docs/StatusEmbedConfigApi.md#getpagespageidstatusembedconfig) | **Get** /pages/{page_id}/status_embed_config | Get status embed config settings
*StatusEmbedConfigApi* | [**PatchPagesPageIdStatusEmbedConfig**](docs/StatusEmbedConfigApi.md#patchpagespageidstatusembedconfig) | **Patch** /pages/{page_id}/status_embed_config | Update status embed config settings
*StatusEmbedConfigApi* | [**PutPagesPageIdStatusEmbedConfig**](docs/StatusEmbedConfigApi.md#putpagespageidstatusembedconfig) | **Put** /pages/{page_id}/status_embed_config | Update status embed config settings
*SubscribersApi* | [**DeletePagesPageIdSubscribersSubscriberId**](docs/SubscribersApi.md#deletepagespageidsubscriberssubscriberid) | **Delete** /pages/{page_id}/subscribers/{subscriber_id} | Unsubscribe a subscriber
*SubscribersApi* | [**GetPagesPageIdSubscribers**](docs/SubscribersApi.md#getpagespageidsubscribers) | **Get** /pages/{page_id}/subscribers | Get a list of subscribers
*SubscribersApi* | [**GetPagesPageIdSubscribersCount**](docs/SubscribersApi.md#getpagespageidsubscriberscount) | **Get** /pages/{page_id}/subscribers/count | Get a count of subscribers by type
*SubscribersApi* | [**GetPagesPageIdSubscribersHistogramByState**](docs/SubscribersApi.md#getpagespageidsubscribershistogrambystate) | **Get** /pages/{page_id}/subscribers/histogram_by_state | Get a histogram of subscribers by type and then state
*SubscribersApi* | [**GetPagesPageIdSubscribersSubscriberId**](docs/SubscribersApi.md#getpagespageidsubscriberssubscriberid) | **Get** /pages/{page_id}/subscribers/{subscriber_id} | Get a subscriber
*SubscribersApi* | [**GetPagesPageIdSubscribersUnsubscribed**](docs/SubscribersApi.md#getpagespageidsubscribersunsubscribed) | **Get** /pages/{page_id}/subscribers/unsubscribed | Get a list of unsubscribed subscribers
*SubscribersApi* | [**PatchPagesPageIdSubscribersSubscriberId**](docs/SubscribersApi.md#patchpagespageidsubscriberssubscriberid) | **Patch** /pages/{page_id}/subscribers/{subscriber_id} | Update a subscriber
*SubscribersApi* | [**PostPagesPageIdSubscribers**](docs/SubscribersApi.md#postpagespageidsubscribers) | **Post** /pages/{page_id}/subscribers | Create a subscriber
*SubscribersApi* | [**PostPagesPageIdSubscribersReactivate**](docs/SubscribersApi.md#postpagespageidsubscribersreactivate) | **Post** /pages/{page_id}/subscribers/reactivate | Reactivate a list of subscribers
*SubscribersApi* | [**PostPagesPageIdSubscribersResendConfirmation**](docs/SubscribersApi.md#postpagespageidsubscribersresendconfirmation) | **Post** /pages/{page_id}/subscribers/resend_confirmation | Resend confirmations to a list of subscribers
*SubscribersApi* | [**PostPagesPageIdSubscribersSubscriberIdResendConfirmation**](docs/SubscribersApi.md#postpagespageidsubscriberssubscriberidresendconfirmation) | **Post** /pages/{page_id}/subscribers/{subscriber_id}/resend_confirmation | Resend confirmation to a subscriber
*SubscribersApi* | [**PostPagesPageIdSubscribersUnsubscribe**](docs/SubscribersApi.md#postpagespageidsubscribersunsubscribe) | **Post** /pages/{page_id}/subscribers/unsubscribe | Unsubscribe a list of subscribers
*TemplatesApi* | [**GetPagesPageIdIncidentTemplates**](docs/TemplatesApi.md#getpagespageidincidenttemplates) | **Get** /pages/{page_id}/incident_templates | Get a list of templates
*TemplatesApi* | [**PostPagesPageIdIncidentTemplates**](docs/TemplatesApi.md#postpagespageidincidenttemplates) | **Post** /pages/{page_id}/incident_templates | Create a template
*UsersApi* | [**DeleteOrganizationsOrganizationIdUsersUserId**](docs/UsersApi.md#deleteorganizationsorganizationidusersuserid) | **Delete** /organizations/{organization_id}/users/{user_id} | Delete a user
*UsersApi* | [**GetOrganizationsOrganizationIdUsers**](docs/UsersApi.md#getorganizationsorganizationidusers) | **Get** /organizations/{organization_id}/users | Get a list of users
*UsersApi* | [**PostOrganizationsOrganizationIdUsers**](docs/UsersApi.md#postorganizationsorganizationidusers) | **Post** /organizations/{organization_id}/users | Create a user


## Documentation For Models

 - [Component](docs/Component.md)
 - [ComponentGroupUptime](docs/ComponentGroupUptime.md)
 - [ComponentGroupUptimeRelatedEvents](docs/ComponentGroupUptimeRelatedEvents.md)
 - [ComponentUptime](docs/ComponentUptime.md)
 - [ComponentUptimeRelatedEvents](docs/ComponentUptimeRelatedEvents.md)
 - [DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents](docs/DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents.md)
 - [DeletePagesPageIdPageAccessUsersPageAccessUserIdMetrics](docs/DeletePagesPageIdPageAccessUsersPageAccessUserIdMetrics.md)
 - [ErrorEntity](docs/ErrorEntity.md)
 - [GroupComponent](docs/GroupComponent.md)
 - [Incident](docs/Incident.md)
 - [IncidentTemplate](docs/IncidentTemplate.md)
 - [IncidentUpdate](docs/IncidentUpdate.md)
 - [Metric](docs/Metric.md)
 - [MetricAddResponse](docs/MetricAddResponse.md)
 - [MetricAddResponseMetricId](docs/MetricAddResponseMetricId.md)
 - [MetricsProvider](docs/MetricsProvider.md)
 - [Page](docs/Page.md)
 - [PageAccessGroup](docs/PageAccessGroup.md)
 - [PageAccessUser](docs/PageAccessUser.md)
 - [PatchPages](docs/PatchPages.md)
 - [PatchPagesPage](docs/PatchPagesPage.md)
 - [PatchPagesPageIdComponentGroups](docs/PatchPagesPageIdComponentGroups.md)
 - [PatchPagesPageIdComponents](docs/PatchPagesPageIdComponents.md)
 - [PatchPagesPageIdIncidents](docs/PatchPagesPageIdIncidents.md)
 - [PatchPagesPageIdIncidentsIncident](docs/PatchPagesPageIdIncidentsIncident.md)
 - [PatchPagesPageIdIncidentsIncidentIdIncidentUpdates](docs/PatchPagesPageIdIncidentsIncidentIdIncidentUpdates.md)
 - [PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate](docs/PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate.md)
 - [PatchPagesPageIdMetrics](docs/PatchPagesPageIdMetrics.md)
 - [PatchPagesPageIdMetricsMetric](docs/PatchPagesPageIdMetricsMetric.md)
 - [PatchPagesPageIdMetricsProviders](docs/PatchPagesPageIdMetricsProviders.md)
 - [PatchPagesPageIdMetricsProvidersMetricsProvider](docs/PatchPagesPageIdMetricsProvidersMetricsProvider.md)
 - [PatchPagesPageIdPageAccessGroups](docs/PatchPagesPageIdPageAccessGroups.md)
 - [PatchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents](docs/PatchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents.md)
 - [PatchPagesPageIdPageAccessUsersPageAccessUserIdMetrics](docs/PatchPagesPageIdPageAccessUsersPageAccessUserIdMetrics.md)
 - [PatchPagesPageIdStatusEmbedConfig](docs/PatchPagesPageIdStatusEmbedConfig.md)
 - [PatchPagesPageIdStatusEmbedConfigStatusEmbedConfig](docs/PatchPagesPageIdStatusEmbedConfigStatusEmbedConfig.md)
 - [PatchPagesPageIdSubscribers](docs/PatchPagesPageIdSubscribers.md)
 - [PostOrganizationsOrganizationIdUsers](docs/PostOrganizationsOrganizationIdUsers.md)
 - [PostOrganizationsOrganizationIdUsersUser](docs/PostOrganizationsOrganizationIdUsersUser.md)
 - [PostPagesPageIdComponentGroups](docs/PostPagesPageIdComponentGroups.md)
 - [PostPagesPageIdComponentGroupsComponentGroup](docs/PostPagesPageIdComponentGroupsComponentGroup.md)
 - [PostPagesPageIdComponents](docs/PostPagesPageIdComponents.md)
 - [PostPagesPageIdComponentsComponent](docs/PostPagesPageIdComponentsComponent.md)
 - [PostPagesPageIdIncidentTemplates](docs/PostPagesPageIdIncidentTemplates.md)
 - [PostPagesPageIdIncidentTemplatesTemplate](docs/PostPagesPageIdIncidentTemplatesTemplate.md)
 - [PostPagesPageIdIncidents](docs/PostPagesPageIdIncidents.md)
 - [PostPagesPageIdIncidentsIncident](docs/PostPagesPageIdIncidentsIncident.md)
 - [PostPagesPageIdIncidentsIncidentComponents](docs/PostPagesPageIdIncidentsIncidentComponents.md)
 - [PostPagesPageIdIncidentsIncidentIdSubscribers](docs/PostPagesPageIdIncidentsIncidentIdSubscribers.md)
 - [PostPagesPageIdIncidentsIncidentIdSubscribersSubscriber](docs/PostPagesPageIdIncidentsIncidentIdSubscribersSubscriber.md)
 - [PostPagesPageIdMetricsData](docs/PostPagesPageIdMetricsData.md)
 - [PostPagesPageIdMetricsMetricIdData](docs/PostPagesPageIdMetricsMetricIdData.md)
 - [PostPagesPageIdMetricsMetricIdDataData](docs/PostPagesPageIdMetricsMetricIdDataData.md)
 - [PostPagesPageIdMetricsProviders](docs/PostPagesPageIdMetricsProviders.md)
 - [PostPagesPageIdMetricsProvidersMetricsProvider](docs/PostPagesPageIdMetricsProvidersMetricsProvider.md)
 - [PostPagesPageIdMetricsProvidersMetricsProviderIdMetrics](docs/PostPagesPageIdMetricsProvidersMetricsProviderIdMetrics.md)
 - [PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric](docs/PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric.md)
 - [PostPagesPageIdPageAccessGroups](docs/PostPagesPageIdPageAccessGroups.md)
 - [PostPagesPageIdPageAccessGroupsPageAccessGroup](docs/PostPagesPageIdPageAccessGroupsPageAccessGroup.md)
 - [PostPagesPageIdPageAccessGroupsPageAccessGroupIdComponents](docs/PostPagesPageIdPageAccessGroupsPageAccessGroupIdComponents.md)
 - [PostPagesPageIdPageAccessUsers](docs/PostPagesPageIdPageAccessUsers.md)
 - [PostPagesPageIdPageAccessUsersPageAccessUser](docs/PostPagesPageIdPageAccessUsersPageAccessUser.md)
 - [PostPagesPageIdPageAccessUsersPageAccessUserIdMetrics](docs/PostPagesPageIdPageAccessUsersPageAccessUserIdMetrics.md)
 - [PostPagesPageIdSubscribers](docs/PostPagesPageIdSubscribers.md)
 - [PostPagesPageIdSubscribersReactivate](docs/PostPagesPageIdSubscribersReactivate.md)
 - [PostPagesPageIdSubscribersResendConfirmation](docs/PostPagesPageIdSubscribersResendConfirmation.md)
 - [PostPagesPageIdSubscribersSubscriber](docs/PostPagesPageIdSubscribersSubscriber.md)
 - [PostPagesPageIdSubscribersUnsubscribe](docs/PostPagesPageIdSubscribersUnsubscribe.md)
 - [Postmortem](docs/Postmortem.md)
 - [PutPages](docs/PutPages.md)
 - [PutPagesPageIdComponentGroups](docs/PutPagesPageIdComponentGroups.md)
 - [PutPagesPageIdComponents](docs/PutPagesPageIdComponents.md)
 - [PutPagesPageIdIncidents](docs/PutPagesPageIdIncidents.md)
 - [PutPagesPageIdIncidentsIncidentIdIncidentUpdates](docs/PutPagesPageIdIncidentsIncidentIdIncidentUpdates.md)
 - [PutPagesPageIdIncidentsIncidentIdPostmortem](docs/PutPagesPageIdIncidentsIncidentIdPostmortem.md)
 - [PutPagesPageIdIncidentsIncidentIdPostmortemPostmortem](docs/PutPagesPageIdIncidentsIncidentIdPostmortemPostmortem.md)
 - [PutPagesPageIdIncidentsIncidentIdPostmortemPublish](docs/PutPagesPageIdIncidentsIncidentIdPostmortemPublish.md)
 - [PutPagesPageIdIncidentsIncidentIdPostmortemPublishPostmortem](docs/PutPagesPageIdIncidentsIncidentIdPostmortemPublishPostmortem.md)
 - [PutPagesPageIdMetrics](docs/PutPagesPageIdMetrics.md)
 - [PutPagesPageIdMetricsProviders](docs/PutPagesPageIdMetricsProviders.md)
 - [PutPagesPageIdPageAccessGroups](docs/PutPagesPageIdPageAccessGroups.md)
 - [PutPagesPageIdPageAccessGroupsPageAccessGroupIdComponents](docs/PutPagesPageIdPageAccessGroupsPageAccessGroupIdComponents.md)
 - [PutPagesPageIdPageAccessUsersPageAccessUserIdMetrics](docs/PutPagesPageIdPageAccessUsersPageAccessUserIdMetrics.md)
 - [PutPagesPageIdStatusEmbedConfig](docs/PutPagesPageIdStatusEmbedConfig.md)
 - [SingleMetricAddResponse](docs/SingleMetricAddResponse.md)
 - [StatusEmbedConfig](docs/StatusEmbedConfig.md)
 - [Subscriber](docs/Subscriber.md)
 - [SubscriberCountByState](docs/SubscriberCountByState.md)
 - [SubscriberCountByType](docs/SubscriberCountByType.md)
 - [SubscriberCountByTypeAndState](docs/SubscriberCountByTypeAndState.md)
 - [User](docs/User.md)


## Documentation For Authorization

## api_key
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author


