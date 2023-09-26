/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Events
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// EventsV1Sink struct for EventsV1Sink
type EventsV1Sink struct {
	// The date that this Sink was created, given in ISO 8601 format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this Sink was updated, given in ISO 8601 format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// A human readable description for the Sink
	Description *string `json:"description,omitempty"`
	// A 34 character string that uniquely identifies this Sink.
	Sid *string `json:"sid,omitempty"`
	// The information required for Twilio to connect to the provided Sink encoded as JSON.
	SinkConfiguration *interface{} `json:"sink_configuration,omitempty"`
	SinkType          *string      `json:"sink_type,omitempty"`
	Status            *string      `json:"status,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
	// Contains a dictionary of URL links to nested resources of this Sink.
	Links *map[string]interface{} `json:"links,omitempty"`
}