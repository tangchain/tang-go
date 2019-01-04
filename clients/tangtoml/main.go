package tangtoml

import "net/http"

// TangTomlMaxSize is the maximum size of tang.toml file
const TangTomlMaxSize = 5 * 1024

// WellKnownPath represents the url path at which the tang.toml file should
// exist to conform to the federation protocol.
const WellKnownPath = "/.well-known/tang.toml"

// DefaultClient is a default client using the default parameters
var DefaultClient = &Client{HTTP: http.DefaultClient}

// Client represents a client that is capable of resolving a Tang.toml file
// using the internet.
type Client struct {
	// HTTP is the http client used when resolving a Tang.toml file
	HTTP HTTP

	// UseHTTP forces the client to resolve against servers using plain HTTP.
	// Useful for debugging.
	UseHTTP bool
}

type ClientInterface interface {
	GetTangToml(domain string) (*Response, error)
	GetTangTomlByAddress(addy string) (*Response, error)
}

// HTTP represents the http client that a stellertoml resolver uses to make http
// requests.
type HTTP interface {
	Get(url string) (*http.Response, error)
}

// Response represents the results of successfully resolving a tang.toml file
type Response struct {
	AuthServer       string `toml:"AUTH_SERVER"`
	FederationServer string `toml:"FEDERATION_SERVER"`
	EncryptionKey    string `toml:"ENCRYPTION_KEY"`
	SigningKey       string `toml:"SIGNING_KEY"`
}

// GetTangToml returns tang.toml file for a given domain
func GetTangToml(domain string) (*Response, error) {
	return DefaultClient.GetTangToml(domain)
}

// GetTangTomlByAddress returns tang.toml file of a domain fetched from a
// given address
func GetTangTomlByAddress(addy string) (*Response, error) {
	return DefaultClient.GetTangTomlByAddress(addy)
}
