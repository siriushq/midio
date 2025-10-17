package madmin

// Args - defines the arguments for the API.
type logArgs struct {
	Bucket   string            `json:"bucket,omitempty"`
	Object   string            `json:"object,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

// Trace - defines the trace.
type logTrace struct {
	Message   string            `json:"message,omitempty"`
	Source    []string          `json:"source,omitempty"`
	Variables map[string]string `json:"variables,omitempty"`
}

// API - defines the api type and its args.
type logAPI struct {
	Name string   `json:"name,omitempty"`
	Args *logArgs `json:"args,omitempty"`
}

// Entry - defines fields and values of each log entry.
type logEntry struct {
	DeploymentID string    `json:"deploymentid,omitempty"`
	Level        string    `json:"level"`
	LogKind      string    `json:"errKind"`
	Time         string    `json:"time"`
	API          *logAPI   `json:"api,omitempty"`
	RemoteHost   string    `json:"remotehost,omitempty"`
	Host         string    `json:"host,omitempty"`
	RequestID    string    `json:"requestID,omitempty"`
	UserAgent    string    `json:"userAgent,omitempty"`
	Message      string    `json:"message,omitempty"`
	Trace        *logTrace `json:"error,omitempty"`
}
