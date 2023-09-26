/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'DeleteMedia'
type DeleteMediaParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is associated with the Media resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteMediaParams) SetPathAccountSid(PathAccountSid string) *DeleteMediaParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete the Media resource.
func (c *ApiService) DeleteMedia(MessageSid string, Sid string, params *DeleteMediaParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/Messages/{MessageSid}/Media/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"MessageSid"+"}", MessageSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Optional parameters for the method 'FetchMedia'
type FetchMediaParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) associated with the Media resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchMediaParams) SetPathAccountSid(PathAccountSid string) *FetchMediaParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch a single Media resource associated with a specific Message resource
func (c *ApiService) FetchMedia(MessageSid string, Sid string, params *FetchMediaParams) (*ApiV2010Media, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Messages/{MessageSid}/Media/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"MessageSid"+"}", MessageSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010Media{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListMedia'
type ListMediaParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is associated with the Media resources.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// Only include Media resources that were created on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read Media that were created on this date. You can also specify an inequality, such as `StartTime<=YYYY-MM-DD`, to read Media that were created on or before midnight of this date, and `StartTime>=YYYY-MM-DD` to read Media that were created on or after midnight of this date.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// Only include Media resources that were created on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read Media that were created on this date. You can also specify an inequality, such as `StartTime<=YYYY-MM-DD`, to read Media that were created on or before midnight of this date, and `StartTime>=YYYY-MM-DD` to read Media that were created on or after midnight of this date.
	DateCreatedBefore *time.Time `json:"DateCreated&lt;,omitempty"`
	// Only include Media resources that were created on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read Media that were created on this date. You can also specify an inequality, such as `StartTime<=YYYY-MM-DD`, to read Media that were created on or before midnight of this date, and `StartTime>=YYYY-MM-DD` to read Media that were created on or after midnight of this date.
	DateCreatedAfter *time.Time `json:"DateCreated&gt;,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListMediaParams) SetPathAccountSid(PathAccountSid string) *ListMediaParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListMediaParams) SetDateCreated(DateCreated time.Time) *ListMediaParams {
	params.DateCreated = &DateCreated
	return params
}
func (params *ListMediaParams) SetDateCreatedBefore(DateCreatedBefore time.Time) *ListMediaParams {
	params.DateCreatedBefore = &DateCreatedBefore
	return params
}
func (params *ListMediaParams) SetDateCreatedAfter(DateCreatedAfter time.Time) *ListMediaParams {
	params.DateCreatedAfter = &DateCreatedAfter
	return params
}
func (params *ListMediaParams) SetPageSize(PageSize int) *ListMediaParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListMediaParams) SetLimit(Limit int) *ListMediaParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Media records from the API. Request is executed immediately.
func (c *ApiService) PageMedia(MessageSid string, params *ListMediaParams, pageToken, pageNumber string) (*ListMediaResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Messages/{MessageSid}/Media.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"MessageSid"+"}", MessageSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.DateCreated != nil {
		data.Set("DateCreated", fmt.Sprint((*params.DateCreated).Format(time.RFC3339)))
	}
	if params != nil && params.DateCreatedBefore != nil {
		data.Set("DateCreated<", fmt.Sprint((*params.DateCreatedBefore).Format(time.RFC3339)))
	}
	if params != nil && params.DateCreatedAfter != nil {
		data.Set("DateCreated>", fmt.Sprint((*params.DateCreatedAfter).Format(time.RFC3339)))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMediaResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Media records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListMedia(MessageSid string, params *ListMediaParams) ([]ApiV2010Media, error) {
	response, errors := c.StreamMedia(MessageSid, params)

	records := make([]ApiV2010Media, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Media records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamMedia(MessageSid string, params *ListMediaParams) (chan ApiV2010Media, chan error) {
	if params == nil {
		params = &ListMediaParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ApiV2010Media, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageMedia(MessageSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamMedia(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamMedia(response *ListMediaResponse, params *ListMediaParams, recordChannel chan ApiV2010Media, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.MediaList
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListMediaResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListMediaResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListMediaResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMediaResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
