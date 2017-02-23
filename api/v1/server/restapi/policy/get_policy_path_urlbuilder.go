package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// GetPolicyPathURL generates an URL for the get policy path operation
type GetPolicyPathURL struct {
	Path string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetPolicyPathURL) WithBasePath(bp string) *GetPolicyPathURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetPolicyPathURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetPolicyPathURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/policy/{path}"

	path := o.Path
	if path != "" {
		_path = strings.Replace(_path, "{path}", path, -1)
	} else {
		return nil, errors.New("Path is required on GetPolicyPathURL")
	}
	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v1beta"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetPolicyPathURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetPolicyPathURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetPolicyPathURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetPolicyPathURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetPolicyPathURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetPolicyPathURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}