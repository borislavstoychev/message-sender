/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Pricing
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"net/url"
	"strings"
)

// Optional parameters for the method 'FetchTrunkingNumber'
type FetchTrunkingNumberParams struct {
	// The origination phone number, in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, for which to fetch the origin-based voice pricing information. E.164 format consists of a + followed by the country code and subscriber number.
	OriginationNumber *string `json:"OriginationNumber,omitempty"`
}

func (params *FetchTrunkingNumberParams) SetOriginationNumber(OriginationNumber string) *FetchTrunkingNumberParams {
	params.OriginationNumber = &OriginationNumber
	return params
}

// Fetch pricing information for a specific destination and, optionally, origination phone number.
func (c *ApiService) FetchTrunkingNumber(DestinationNumber string, params *FetchTrunkingNumberParams) (*PricingV2TrunkingNumber, error) {
	path := "/v2/Trunking/Numbers/{DestinationNumber}"
	path = strings.Replace(path, "{"+"DestinationNumber"+"}", DestinationNumber, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.OriginationNumber != nil {
		data.Set("OriginationNumber", *params.OriginationNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &PricingV2TrunkingNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
