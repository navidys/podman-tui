// Code generated by go generate; DO NOT EDIT.
package containers

import (
	"bufio"
	"io"
	"net/url"

	"github.com/containers/podman/v4/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *ExecStartAndAttachOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *ExecStartAndAttachOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithOutputStream set field OutputStream to given value
func (o *ExecStartAndAttachOptions) WithOutputStream(value io.WriteCloser) *ExecStartAndAttachOptions {
	o.OutputStream = &value
	return o
}

// GetOutputStream returns value of field OutputStream
func (o *ExecStartAndAttachOptions) GetOutputStream() io.WriteCloser {
	if o.OutputStream == nil {
		var z io.WriteCloser
		return z
	}
	return *o.OutputStream
}

// WithErrorStream set field ErrorStream to given value
func (o *ExecStartAndAttachOptions) WithErrorStream(value io.WriteCloser) *ExecStartAndAttachOptions {
	o.ErrorStream = &value
	return o
}

// GetErrorStream returns value of field ErrorStream
func (o *ExecStartAndAttachOptions) GetErrorStream() io.WriteCloser {
	if o.ErrorStream == nil {
		var z io.WriteCloser
		return z
	}
	return *o.ErrorStream
}

// WithInputStream set field InputStream to given value
func (o *ExecStartAndAttachOptions) WithInputStream(value bufio.Reader) *ExecStartAndAttachOptions {
	o.InputStream = &value
	return o
}

// GetInputStream returns value of field InputStream
func (o *ExecStartAndAttachOptions) GetInputStream() bufio.Reader {
	if o.InputStream == nil {
		var z bufio.Reader
		return z
	}
	return *o.InputStream
}

// WithAttachOutput set field AttachOutput to given value
func (o *ExecStartAndAttachOptions) WithAttachOutput(value bool) *ExecStartAndAttachOptions {
	o.AttachOutput = &value
	return o
}

// GetAttachOutput returns value of field AttachOutput
func (o *ExecStartAndAttachOptions) GetAttachOutput() bool {
	if o.AttachOutput == nil {
		var z bool
		return z
	}
	return *o.AttachOutput
}

// WithAttachError set field AttachError to given value
func (o *ExecStartAndAttachOptions) WithAttachError(value bool) *ExecStartAndAttachOptions {
	o.AttachError = &value
	return o
}

// GetAttachError returns value of field AttachError
func (o *ExecStartAndAttachOptions) GetAttachError() bool {
	if o.AttachError == nil {
		var z bool
		return z
	}
	return *o.AttachError
}

// WithAttachInput set field AttachInput to given value
func (o *ExecStartAndAttachOptions) WithAttachInput(value bool) *ExecStartAndAttachOptions {
	o.AttachInput = &value
	return o
}

// GetAttachInput returns value of field AttachInput
func (o *ExecStartAndAttachOptions) GetAttachInput() bool {
	if o.AttachInput == nil {
		var z bool
		return z
	}
	return *o.AttachInput
}