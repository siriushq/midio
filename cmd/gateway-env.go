package cmd

type gatewaySSE []string

const (
	// GatewaySSES3 is set when SSE-S3 encryption needed on both gateway and backend
	gatewaySSES3 = "S3"
	// GatewaySSEC is set when SSE-C encryption needed on both gateway and backend
	gatewaySSEC = "C"
)

func (sse gatewaySSE) SSES3() bool {
	for _, v := range sse {
		if v == gatewaySSES3 {
			return true
		}
	}
	return false
}

func (sse gatewaySSE) SSEC() bool {
	for _, v := range sse {
		if v == gatewaySSEC {
			return true
		}
	}
	return false
}

func (sse gatewaySSE) IsSet() bool {
	return sse.SSES3() || sse.SSEC()
}
