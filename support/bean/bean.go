package bean

type Request struct {
	Type  string       `json:"type,omitempty"`
	Http  HttpRequest  `json:"http,omitempty"`
	SSH   SSHRequest   `json:"ssh,omitempty"`
	Shell ShellRequest `json:"shell,omitempty"`
}

type HttpRequest struct {
	Url     string            `json:"url,omitempty"`
	Method  string            `json:"method,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	Body    string            `json:"body,omitempty"`
}

type HttpResponse struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

type ShellRequest struct {
	Commands []string `json:"commands,omitempty"`
}

type ShellResponse struct {
	Results []string `json:"results,omitempty"`
}

type SSHRequest struct {
	ServerId string   `json:"serverId,omitempty"`
	Commands []string `json:"commands,omitempty"`
}

type SSHResponse struct {
	Results []string `json:"results,omitempty"`
}
