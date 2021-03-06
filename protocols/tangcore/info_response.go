package tangcore

// InfoResponse is the json response returned from tang-core's /info
// endpoint.
type InfoResponse struct {
	Info struct {
		Build           string `json:"build"`
		Network         string `json:"network"`
		ProtocolVersion int    `json:"protocol_version"`
		State           string `json:"state"`

		// TODO: all the other fields
	}
}

// IsSynced returns a boolean indicating whether tangcore is synced with the
// network.
func (resp *InfoResponse) IsSynced() bool {
	return resp.Info.State == "Synced!"
}
