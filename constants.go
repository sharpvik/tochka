package tochka

const (
	BaseURL      = "https://enter.tochka.com/"
	ProdURI      = "uapi"
	Version      = "v1.0"
	SandboxURI   = "sandbox/v2"
	SandboxToken = "working_token"
)

var (
	ProdURL    = BaseURL + ProdURI
	SandboxURL = BaseURL + SandboxURI
)
