package madmin

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/siriushq/midio/pkg/bandwidth"
)

// Report includes the bandwidth report or the error encountered.
type Report struct {
	Report bandwidth.Report `json:"report"`
	Err    error            `json:"error,omitempty"`
}

// GetBucketBandwidth - Gets a channel reporting bandwidth measurements for replication buckets. If no buckets
// generate replication traffic an empty map is returned in the report until traffic is seen.
func (adm *AdminClient) GetBucketBandwidth(ctx context.Context, buckets ...string) <-chan Report {
	queryValues := url.Values{}
	ch := make(chan Report)
	if len(buckets) > 0 {
		queryValues.Set("buckets", strings.Join(buckets, ","))
	}

	reqData := requestData{
		relPath:     adminAPIPrefix + "/bandwidth",
		queryValues: queryValues,
	}
	resp, err := adm.executeMethod(ctx, http.MethodGet, reqData)
	if err != nil {
		defer closeResponse(resp)
		ch <- Report{bandwidth.Report{}, err}
		return ch
	}
	if resp.StatusCode != http.StatusOK {
		ch <- Report{bandwidth.Report{}, httpRespToErrorResponse(resp)}
		return ch
	}

	dec := json.NewDecoder(resp.Body)

	go func(ctx context.Context, ch chan<- Report, resp *http.Response) {
		defer func() {
			closeResponse(resp)
			close(ch)
		}()
		for {
			var report bandwidth.Report

			if err = dec.Decode(&report); err != nil {
				ch <- Report{bandwidth.Report{}, err}
				return
			}
			select {
			case <-ctx.Done():
				return
			case ch <- Report{Report: report, Err: err}:
			}
		}
	}(ctx, ch, resp)
	return ch
}
