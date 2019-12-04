package channel

// HBServiceData ...
type HBServiceData struct {
	Service string `json:"service"`
	Detail  string `json:"detail"`
}

var servicesChData chan HBServiceData

// Initialize the data channel
func init() {
	servicesChData = make(chan HBServiceData, 20)
}

// GetServicesChData ...
func GetServicesChData() chan HBServiceData {
	return servicesChData
}
