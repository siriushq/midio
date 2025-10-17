package madmin

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// LogInfo holds console log messages
type LogInfo struct {
	logEntry
	ConsoleMsg string
	NodeName   string `json:"node"`
	Err        error  `json:"-"`
}

// GetLogs - listen on console log messages.
func (adm AdminClient) GetLogs(ctx context.Context, node string, lineCnt int, logKind string) <-chan LogInfo {
	logCh := make(chan LogInfo, 1)

	// Only success, start a routine to start reading line by line.
	go func(logCh chan<- LogInfo) {
		defer close(logCh)
		urlValues := make(url.Values)
		urlValues.Set("node", node)
		urlValues.Set("limit", strconv.Itoa(lineCnt))
		urlValues.Set("logType", logKind)
		for {
			reqData := requestData{
				relPath:     adminAPIPrefix + "/log",
				queryValues: urlValues,
			}
			// Execute GET to call log handler
			resp, err := adm.executeMethod(ctx, http.MethodGet, reqData)
			if err != nil {
				closeResponse(resp)
				return
			}

			if resp.StatusCode != http.StatusOK {
				logCh <- LogInfo{Err: httpRespToErrorResponse(resp)}
				return
			}
			dec := json.NewDecoder(resp.Body)
			for {
				var info LogInfo
				if err = dec.Decode(&info); err != nil {
					break
				}
				select {
				case <-ctx.Done():
					return
				case logCh <- info:
				}
			}

		}
	}(logCh)

	// Returns the log info channel, for caller to start reading from.
	return logCh
}
