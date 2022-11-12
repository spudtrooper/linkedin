// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package linkedin

import "fmt"

type LoginOption struct {
	f func(*loginOptionImpl)
	s string
}

func (o LoginOption) String() string { return o.s }

type LoginOptions interface {
	SeleniumHead() bool
	HasSeleniumHead() bool
	SeleniumVerbose() bool
	HasSeleniumVerbose() bool
}

func LoginSeleniumHead(seleniumHead bool) LoginOption {
	return LoginOption{func(opts *loginOptionImpl) {
		opts.has_seleniumHead = true
		opts.seleniumHead = seleniumHead
	}, fmt.Sprintf("linkedin.LoginSeleniumHead(bool %+v)", seleniumHead)}
}
func LoginSeleniumHeadFlag(seleniumHead *bool) LoginOption {
	return LoginOption{func(opts *loginOptionImpl) {
		if seleniumHead == nil {
			return
		}
		opts.has_seleniumHead = true
		opts.seleniumHead = *seleniumHead
	}, fmt.Sprintf("linkedin.LoginSeleniumHead(bool %+v)", seleniumHead)}
}

func LoginSeleniumVerbose(seleniumVerbose bool) LoginOption {
	return LoginOption{func(opts *loginOptionImpl) {
		opts.has_seleniumVerbose = true
		opts.seleniumVerbose = seleniumVerbose
	}, fmt.Sprintf("linkedin.LoginSeleniumVerbose(bool %+v)", seleniumVerbose)}
}
func LoginSeleniumVerboseFlag(seleniumVerbose *bool) LoginOption {
	return LoginOption{func(opts *loginOptionImpl) {
		if seleniumVerbose == nil {
			return
		}
		opts.has_seleniumVerbose = true
		opts.seleniumVerbose = *seleniumVerbose
	}, fmt.Sprintf("linkedin.LoginSeleniumVerbose(bool %+v)", seleniumVerbose)}
}

type loginOptionImpl struct {
	seleniumHead        bool
	has_seleniumHead    bool
	seleniumVerbose     bool
	has_seleniumVerbose bool
}

func (l *loginOptionImpl) SeleniumHead() bool       { return l.seleniumHead }
func (l *loginOptionImpl) HasSeleniumHead() bool    { return l.has_seleniumHead }
func (l *loginOptionImpl) SeleniumVerbose() bool    { return l.seleniumVerbose }
func (l *loginOptionImpl) HasSeleniumVerbose() bool { return l.has_seleniumVerbose }

func makeLoginOptionImpl(opts ...LoginOption) *loginOptionImpl {
	res := &loginOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeLoginOptions(opts ...LoginOption) LoginOptions {
	return makeLoginOptionImpl(opts...)
}
