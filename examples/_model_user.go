// Code generated. DO NOT EDIT.
package model

import (
	"github.com/golang/glog"
)

// UserOptions
func NewUserOptions() *UserOptions {
	model := new(UserOptions)
	return model
}

type UserOptions struct {
	ModelAbstract
	// FullName
	FullName string
	// FirstName
	FirstName string
	// LastName
	LastName string
}

func (model UserOptions) TransformTo(out interface{}) error {
	switch out.(type) {
	default:
		glog.Errorf("Not supported type %v", out)
		return ErrNotSupported
	}
	return nil
}

func (model *UserOptions) TransformFrom(in interface{}) error {
	switch in.(type) {
	default:
		glog.Errorf("Not supported type %v", in)
		return ErrNotSupported
	}
	return nil

}

//
// Helpful functions
//

func (u UserOptions) Maps() map[string]interface{} {
	maps := u.ModelAbstract.Maps()
	// FullName
	maps["fullname"] = u.FullName
	// FirstName
	maps["firstname"] = u.FirstName
	// LastName
	maps["lastname"] = u.LastName
	return maps
}

// Fields extract of fields from map
func (u UserOptions) Fields(fields ...string) ([]string, []interface{}) {
	return ExtractFieldsFromMap(u.Maps(), fields...)
}

// FromJson data as []byte or io.Reader
func (u *UserOptions) FromJson(data interface{}) error {
	return FromJson(u, data)
}

// User
func NewUser() *User {
	model := new(User)
	return model
}

type User struct {
	ModelAbstract
	// UserId
	UserId int64
	// Name	login
	Name string
	// Email
	Email string
	// Options
	Options UserOptions
}

func (model User) TransformTo(out interface{}) error {
	switch out.(type) {
	case *UserShortInfo:
		dto := out.(*UserShortInfo)
		dto.Name = model.Name
		dto.FullName = model.Options.FullName
	default:
		glog.Errorf("Not supported type %v", out)
		return ErrNotSupported
	}
	return nil
}

func (model *User) TransformFrom(in interface{}) error {
	switch in.(type) {
	case *UserUpdateEmail:
		dto := in.(*UserUpdateEmail)
		model.Email = dto.Email
	default:
		glog.Errorf("Not supported type %v", in)
		return ErrNotSupported
	}
	return nil

}

//
// Helpful functions
//

func (u User) Maps() map[string]interface{} {
	maps := u.ModelAbstract.Maps()
	// UserId
	maps["userid"] = u.UserId
	// Name	login
	maps["name"] = u.Name
	// Email
	maps["email"] = u.Email
	// Options
	maps["options"] = u.Options
	return maps
}

// Fields extract of fields from map
func (u User) Fields(fields ...string) ([]string, []interface{}) {
	return ExtractFieldsFromMap(u.Maps(), fields...)
}

// FromJson data as []byte or io.Reader
func (u *User) FromJson(data interface{}) error {
	return FromJson(u, data)
}

func (User) TableName() string {
	return "users"
}

// model
// UserUpdateEmail
func NewUserUpdateEmail() *UserUpdateEmail {
	model := new(UserUpdateEmail)
	return model
}

type UserUpdateEmail struct {
	// Email
	Email string `json:"email" validator:"required,email" `
}

func (model UserUpdateEmail) TransformTo(out interface{}) error {
	switch out.(type) {
	default:
		glog.Errorf("Not supported type %v", out)
		return ErrNotSupported
	}
	return nil
}

func (model *UserUpdateEmail) TransformFrom(in interface{}) error {
	switch in.(type) {
	default:
		glog.Errorf("Not supported type %v", in)
		return ErrNotSupported
	}
	return nil

}

//
// Helpful functions
//

func (u UserUpdateEmail) Maps() map[string]interface{} {
	return map[string]interface{}{
		// Email
		"email": &u.Email,
	}
}

// Fields extract of fields from map
func (u UserUpdateEmail) Fields(fields ...string) ([]string, []interface{}) {
	return ExtractFieldsFromMap(u.Maps(), fields...)
}

// FromJson data as []byte or io.Reader
func (u *UserUpdateEmail) FromJson(data interface{}) error {
	return FromJson(u, data)
}

// UserShortInfo
func NewUserShortInfo() *UserShortInfo {
	model := new(UserShortInfo)
	return model
}

type UserShortInfo struct {
	// Name
	Name string `json:"login" `
	// FullName
	FullName string
}

func (model UserShortInfo) TransformTo(out interface{}) error {
	switch out.(type) {
	default:
		glog.Errorf("Not supported type %v", out)
		return ErrNotSupported
	}
	return nil
}

func (model *UserShortInfo) TransformFrom(in interface{}) error {
	switch in.(type) {
	default:
		glog.Errorf("Not supported type %v", in)
		return ErrNotSupported
	}
	return nil

}

//
// Helpful functions
//

func (u UserShortInfo) Maps() map[string]interface{} {
	return map[string]interface{}{
		// Name
		"name": &u.Name,
		// FullName
		"fullname": &u.FullName,
	}
}

// Fields extract of fields from map
func (u UserShortInfo) Fields(fields ...string) ([]string, []interface{}) {
	return ExtractFieldsFromMap(u.Maps(), fields...)
}

// FromJson data as []byte or io.Reader
func (u *UserShortInfo) FromJson(data interface{}) error {
	return FromJson(u, data)
}
