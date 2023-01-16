// Code generated by Validator v0.1.4. DO NOT EDIT.

package douyin_user

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = (*regexp.Regexp)(nil)
	_ = time.Nanosecond
)

func (p *UserRegisterRequest) IsValid() error {
	return nil
}
func (p *UserRegisterResponse) IsValid() error {
	return nil
}
func (p *UserLoginRequest) IsValid() error {
	return nil
}
func (p *UserLoginResponse) IsValid() error {
	return nil
}
func (p *UserRequest) IsValid() error {
	return nil
}
func (p *UserResponse) IsValid() error {
	if p.User != nil {
		if err := p.User.IsValid(); err != nil {
			return fmt.Errorf("filed User not valid, %w", err)
		}
	}
	return nil
}
func (p *User) IsValid() error {
	return nil
}
