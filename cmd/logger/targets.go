package logger

// Target is the entity that we will receive
// a single log entry and Send it to the log target
//
//	e.g. Send the log to a http server
type Target interface {
	String() string
	Endpoint() string
	Validate() error
	Send(entry interface{}, errKind string) error
}

// Targets is the set of enabled loggers
var Targets = []Target{}

// AuditTargets is the list of enabled audit loggers
var AuditTargets = []Target{}

// AddAuditTarget adds a new audit logger target to the
// list of enabled loggers
func AddAuditTarget(t Target) error {
	if err := t.Validate(); err != nil {
		return err
	}

	AuditTargets = append(AuditTargets, t)
	return nil
}

// AddTarget adds a new logger target to the
// list of enabled loggers
func AddTarget(t Target) error {
	if err := t.Validate(); err != nil {
		return err
	}
	Targets = append(Targets, t)
	return nil
}
