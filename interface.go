package vox

import (
	"github.com/end-r/efp"
)

// ParseString ...
func ParseString(str string) (*Division, []string) {
	p, errs := efp.PrototypeFile("vox.efp")
	if errs != nil {
		return nil, errs
	}
	e, errs := p.ValidateString(str)
	if errs != nil {
		return nil, errs
	}
	return Vox(e)
}

// ParseBytes ...
func ParseBytes(bytes []byte) (*Division, []string) {
	p, errs := efp.PrototypeFile("vox.efp")
	if errs != nil {
		return nil, errs
	}
	e, errs := p.ValidateBytes(bytes)
	if errs != nil {
		return nil, errs
	}
	return Vox(e)
}

// ParseFile ...
func ParseFile(path string) (*Division, []string) {
	p, errs := efp.PrototypeFile("vox.efp")
	if errs != nil {
		return nil, errs
	}
	e, errs := p.ValidateFile(path)
	if errs != nil {
		return nil, errs
	}
	return Vox(e)
}
