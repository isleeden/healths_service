package entities

type ServiceCheck struct {
	id           int
	serviceId    int
	name         string
	state        string
	statusCode   int
	responseTime int
	endpointUrl  string
	created_at   int
}
