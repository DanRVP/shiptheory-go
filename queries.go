package shiptheory

import (
	"net/url"
	"reflect"
)

type queryBuilder struct {}

// Build a query string from a query struct
func (params queryBuilder) toQueryString() string {
	query := url.Values{}
	values := reflect.ValueOf(params)
	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		value := values.Field(i).Interface().(string)
		if value != "" {
			query.Add(types.Field(i).Tag.Get("url"), value)
		}

	}

	return query.Encode()
}

type ListShipmentsQuery struct {
	Created string `url:"created"`
	Modified string `url:"modified"`
	Status string `url:"status"`
	ChannelName string `url:"channel_name"`
	ChannelReferenceId string `url:"channel_reference_id"`
	ChannelReferenceId2 string `url:"channel_reference_id_2"`
}

