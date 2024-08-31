package shiptheory

type ListShipmentsQuery struct {
	Created string `url:"created,omitempty"`
	Modified string `url:"modified,omitempty"`
	Status string `url:"status,omitempty"`
	ChannelName string `url:"channel_name,omitempty"`
	ChannelReferenceId string `url:"channel_reference_id,omitempty"`
	ChannelReferenceId2 string `url:"channel_reference_id_2,omitempty"`
}
