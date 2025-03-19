package constants

var ContentType = struct {
	JSON string
	Form string
	CSV  string
}{
	JSON: "application/json",
	Form: "multipart/form-data",
	CSV:  "text/csv",
}
