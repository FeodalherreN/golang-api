package entities

// Response is used by the api to wrap all data it sends back
type Response struct {
	Data    interface{}
	Success bool
}
