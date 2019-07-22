﻿// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package cix

import (
    "github.com/google/uuid"
    "time"
)

type BuildFrameworkDetectionType string

type buildFrameworkDetectionTypeValuesType struct {
    Shallow BuildFrameworkDetectionType
    Full BuildFrameworkDetectionType
}

var BuildFrameworkDetectionTypeValues = buildFrameworkDetectionTypeValuesType{
    Shallow: "shallow",
    Full: "full",
}

type ConfigurationFile struct {
    // The content of the file.
    Content *string `json:"content,omitempty"`
    // Indicates if the content is base64 encoded.
    IsBase64Encoded *bool `json:"isBase64Encoded,omitempty"`
    // The full path of the file, relative to the root of the repository.
    Path *string `json:"path,omitempty"`
}

type CreatedResources struct {
    Resources *map[string]interface{} `json:"resources,omitempty"`
}

// This class is used to create a pipeline connection within the team project provided. If the team project does not exist, it will be created.
type CreatePipelineConnectionInputs struct {
    // The team project settings for an existing team project or for a new team project.
    Project *TeamProject `json:"project,omitempty"`
    // This dictionary contains information that is specific to the provider. This data is opaque to the rest of the Pipelines infrastructure and does NOT contribute to the resources Token. The format of the string and its contents depend on the implementation of the provider.
    ProviderData *map[string]string `json:"providerData,omitempty"`
    // The external source provider id for which the connection is being made.
    ProviderId *string `json:"providerId,omitempty"`
    // If provided, this will be the URL returned with the PipelineConnection. This will override any other redirect URL that would have been generated for the connection.
    RedirectUrl *string `json:"redirectUrl,omitempty"`
    // Where the request to create the pipeline originated (such as 'GitHub Marketplace' or 'Azure DevOps')
    RequestSource *string `json:"requestSource,omitempty"`
}

type DetectedBuildFramework struct {
    // List of build targets discovered for the framework to act upon.
    BuildTargets *[]DetectedBuildTarget `json:"buildTargets,omitempty"`
    // The unique identifier of the build framework.
    Id *string `json:"id,omitempty"`
    // Additional detected settings for the build framework.
    Settings *map[string]string `json:"settings,omitempty"`
    // The version of the framework if it can be determined from the sources.
    Version *string `json:"version,omitempty"`
}

type DetectedBuildTarget struct {
    Path *string `json:"path,omitempty"`
    Settings *map[string]string `json:"settings,omitempty"`
    Type *string `json:"type,omitempty"`
}

// Contains information about the progress or result of an async operation.
type Operation struct {
    // Unique identifier for the operation.
    Id *uuid.UUID `json:"id,omitempty"`
    // Unique identifier for the plugin.
    PluginId *uuid.UUID `json:"pluginId,omitempty"`
    // The current status of the operation.
    Status *OperationStatus `json:"status,omitempty"`
    // URL to get the full operation object.
    Url *string `json:"url,omitempty"`
    // Links to other related objects.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Detailed messaged about the status of an operation.
    DetailedMessage *string `json:"detailedMessage,omitempty"`
    // Result message for an operation.
    ResultMessage *string `json:"resultMessage,omitempty"`
    // URL to the operation result.
    ResultUrl *OperationResultReference `json:"resultUrl,omitempty"`
}

// Reference for an async operation.
type OperationReference struct {
    // Unique identifier for the operation.
    Id *uuid.UUID `json:"id,omitempty"`
    // Unique identifier for the plugin.
    PluginId *uuid.UUID `json:"pluginId,omitempty"`
    // The current status of the operation.
    Status *OperationStatus `json:"status,omitempty"`
    // URL to get the full operation object.
    Url *string `json:"url,omitempty"`
}

type OperationResultReference struct {
    // URL to the operation result.
    ResultUrl *string `json:"resultUrl,omitempty"`
}

// The status of an operation.
type OperationStatus string

type operationStatusValuesType struct {
    NotSet OperationStatus
    Queued OperationStatus
    InProgress OperationStatus
    Cancelled OperationStatus
    Succeeded OperationStatus
    Failed OperationStatus
}

var OperationStatusValues = operationStatusValuesType{
    // The operation does not have a status set.
    NotSet: "notSet",
    // The operation has been queued.
    Queued: "queued",
    // The operation is in progress.
    InProgress: "inProgress",
    // The operation was cancelled by the user.
    Cancelled: "cancelled",
    // The operation completed successfully.
    Succeeded: "succeeded",
    // The operation completed with a failure.
    Failed: "failed",
}

type PipelineConnection struct {
    // The account id that contains the team project for the connection.
    AccountId *uuid.UUID `json:"accountId,omitempty"`
    // Deprecated: This property is no longer filled in and will be removed soon.
    DefinitionId *int `json:"definitionId,omitempty"`
    // This is the URL that the user should be taken to in order to continue setup.
    RedirectUrl *string `json:"redirectUrl,omitempty"`
    // The service endpoint that was created for the connection.
    ServiceEndpointId *uuid.UUID `json:"serviceEndpointId,omitempty"`
    // The team project that contains the definition for the connection.
    TeamProjectId *uuid.UUID `json:"teamProjectId,omitempty"`
}

type ProjectVisibility string

type projectVisibilityValuesType struct {
    Private ProjectVisibility
    Public ProjectVisibility
}

var ProjectVisibilityValues = projectVisibilityValuesType{
    // The project is only visible to users with explicit access.
    Private: "private",
    // The project is visible to all.
    Public: "public",
}

// The class to represent a collection of REST reference links.
type ReferenceLinks struct {
    // The readonly view of the links.  Because Reference links are readonly, we only want to expose them as read only.
    Links *map[string]interface{} `json:"links,omitempty"`
}

type ResourceCreationParameter struct {
    ResourceToCreate *interface{} `json:"resourceToCreate,omitempty"`
    Type *string `json:"type,omitempty"`
}

// Represents a Team Project object.
type TeamProject struct {
    // Project abbreviation.
    Abbreviation *string `json:"abbreviation,omitempty"`
    // Url to default team identity image.
    DefaultTeamImageUrl *string `json:"defaultTeamImageUrl,omitempty"`
    // The project's description (if any).
    Description *string `json:"description,omitempty"`
    // Project identifier.
    Id *uuid.UUID `json:"id,omitempty"`
    // Project last update time.
    LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
    // Project name.
    Name *string `json:"name,omitempty"`
    // Project revision.
    Revision *uint64 `json:"revision,omitempty"`
    // Project state.
    State *interface{} `json:"state,omitempty"`
    // Url to the full version of the object.
    Url *string `json:"url,omitempty"`
    // Project visibility.
    Visibility *ProjectVisibility `json:"visibility,omitempty"`
    // The links to other objects related to this object.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Set of capabilities this project has (such as process template & version control).
    Capabilities *map[string]map[string]string `json:"capabilities,omitempty"`
    // The shallow ref to the default team.
    DefaultTeam *WebApiTeamRef `json:"defaultTeam,omitempty"`
}

// Represents a shallow reference to a TeamProject.
type TeamProjectReference struct {
    // Project abbreviation.
    Abbreviation *string `json:"abbreviation,omitempty"`
    // Url to default team identity image.
    DefaultTeamImageUrl *string `json:"defaultTeamImageUrl,omitempty"`
    // The project's description (if any).
    Description *string `json:"description,omitempty"`
    // Project identifier.
    Id *uuid.UUID `json:"id,omitempty"`
    // Project last update time.
    LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
    // Project name.
    Name *string `json:"name,omitempty"`
    // Project revision.
    Revision *uint64 `json:"revision,omitempty"`
    // Project state.
    State *interface{} `json:"state,omitempty"`
    // Url to the full version of the object.
    Url *string `json:"url,omitempty"`
    // Project visibility.
    Visibility *ProjectVisibility `json:"visibility,omitempty"`
}

type Template struct {
    Assets *[]TemplateAsset `json:"assets,omitempty"`
    Content *string `json:"content,omitempty"`
    DataSourceBindings *[]TemplateDataSourceBinding `json:"dataSourceBindings,omitempty"`
    Description *string `json:"description,omitempty"`
    IconUrl *string `json:"iconUrl,omitempty"`
    Id *string `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    Parameters *[]TemplateParameterDefinition `json:"parameters,omitempty"`
    RecommendedWeight *int `json:"recommendedWeight,omitempty"`
}

type TemplateAsset struct {
    Content *string `json:"content,omitempty"`
    Description *string `json:"description,omitempty"`
    DestinationPath *string `json:"destinationPath,omitempty"`
    Path *string `json:"path,omitempty"`
    Type *string `json:"type,omitempty"`
}

type TemplateDataSourceBinding struct {
    DataSourceName *string `json:"dataSourceName,omitempty"`
    EndpointParameterName *string `json:"endpointParameterName,omitempty"`
    Parameters *map[string]string `json:"parameters,omitempty"`
    ResultTemplate *string `json:"resultTemplate,omitempty"`
    Target *string `json:"target,omitempty"`
}

type TemplateParameterDefinition struct {
    DefaultValue *string `json:"defaultValue,omitempty"`
    DisplayName *string `json:"displayName,omitempty"`
    Name *string `json:"name,omitempty"`
    PossibleValues *[]string `json:"possibleValues,omitempty"`
    Required *bool `json:"required,omitempty"`
    Type *string `json:"type,omitempty"`
}

type TemplateParameters struct {
    Tokens *map[string]interface{} `json:"tokens,omitempty"`
}

type WebApiTeamRef struct {
    // Team (Identity) Guid. A Team Foundation ID.
    Id *uuid.UUID `json:"id,omitempty"`
    // Team name
    Name *string `json:"name,omitempty"`
    // Team REST API Url
    Url *string `json:"url,omitempty"`
}
