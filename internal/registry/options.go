package registry

import "github.com/marcsauter/terraform-registry/internal/registry/provider"

// Option function to configure the Terraform registry
type Option func(*Registry)

// WithHTTPListen configures the http listen address.
func WithHTTPListen(listenAddr string) Option {
	return func(reg *Registry) {
		reg.listenAddr = listenAddr
	}
}

// WithProviderBackend sets the provider backend to use.
func WithProviderBackend(b provider.Backend) Option {
	return func(reg *Registry) {
		reg.providerBackend = b
	}
}
