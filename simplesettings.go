// Copyright 2012, Bryan Matsuo. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*  Filename:    simplesettings.go
 *  Author:      Bryan Matsuo <bryan.matsuo [at] gmail.com>
 *  Created:     2012-11-25 16:44:51.995426 -0800 PST
 *  Description: Main source file in go-simplesettings
 */

// Package simplesettings defines a type for managing configuration settings
// with arbitrary structure.
package simplesettings

import (
	"fmt"
)

type Map map[string]interface{}

type Interface interface {
	Get(key string) Interface
	String() (string, error)
	MustString() string
	Bytes() ([]byte, error)
	MustBytes() []byte
	Int() (int, error)
	MustInt() int
	Int64() (int64, error)
	MustInt64() int64
	Uint() (uint, error)
	MustUint() uint
	Uint64() (uint64, error)
	MustUint64() uint64
	Float64() (float64, error)
	MustFloat64() float64
	Value() interface{}
}

func New(v interface{}) Interface {
	return Settings{
		data: v,
	}
}

type Settings struct {
	path []string
	data interface{}
}

func (s Settings) Get(key string) Interface {
	path := append(s.path, key)
	switch s.data.(type) {
	case Map:
		return Settings{
			path: path,
			data: s.data.(Map)[key],
		}
	case map[string]interface{}:
		return Settings{
			path: path,
			data: s.data.(map[string]interface{})[key],
		}
	}
	return Settings{
		path: path,
		data: nil,
	}
}

func (s Settings) String() (string, error) {
	if str, ok := s.data.(string); ok {
		return str, nil
	}
	return "", fmt.Errorf("invalid setting at path: %v", s.path)
}

func (s Settings) MustString() string {
	str, err := s.String()
	if err == nil {
		return str
	}
	panic(err.Error())
}

func (s Settings) Bytes() ([]byte, error) {
	switch s.data.(type) {
	case string:
		return []byte(s.data.(string)), nil
	case []byte:
		return s.data.([]byte), nil
	}
	return nil, fmt.Errorf("invalid setting at path: %v", s.path)
}

func (s Settings) MustBytes() []byte {
	p, err := s.Bytes()
	if err == nil {
		return p
	}
	panic(err.Error())
}

func (s Settings) Float64() (float64, error) {
	if str, ok := s.data.(float64); ok {
		return str, nil
	}
	return 0, fmt.Errorf("invalid setting at path: %v", s.path)
}

func (s Settings) MustFloat64() float64 {
	f, err := s.Float64()
	if err == nil {
		return f
	}
	panic(err.Error())
}

func (s Settings) Int64() (int64, error) {
	if str, ok := s.data.(int64); ok {
		return str, nil
	}
	return 0, fmt.Errorf("invalid setting at path: %v", s.path)
}

func (s Settings) MustInt64() int64 {
	i, err := s.Int64()
	if err == nil {
		return i
	}
	panic(err.Error())
}

func (s Settings) Int() (int, error) {
	if str, ok := s.data.(int); ok {
		return str, nil
	}
	return 0, fmt.Errorf("invalid setting at path: %v", s.path)
}

func (s Settings) MustInt() int {
	i, err := s.Int()
	if err == nil {
		return i
	}
	panic(err.Error())
}

func (s Settings) Uint64() (uint64, error) {
	if str, ok := s.data.(uint64); ok {
		return str, nil
	}
	return 0, fmt.Errorf("invalid setting at path: %v", s.path)
}

func (s Settings) MustUint64() uint64 {
	i, err := s.Uint64()
	if err == nil {
		return i
	}
	panic(err.Error())
}

func (s Settings) Uint() (uint, error) {
	if str, ok := s.data.(uint); ok {
		return str, nil
	}
	return 0, fmt.Errorf("invalid setting at path: %v", s.path)
}

func (s Settings) MustUint() uint {
	i, err := s.Uint()
	if err == nil {
		return i
	}
	panic(err.Error())
}

func (s Settings) Value() interface{} {
	return s.data
}
