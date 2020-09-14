// Copyright 2018 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package name

type Options struct {
	Strict   bool // weak by default
	Insecure bool // secure by default
}

func makeOptions(opts ...Option) Options {
	opt := Options{}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

// Option is a functional option for name parsing.
type Option func(*Options)

// StrictValidation is an Option that requires image references to be fully
// specified; i.e. no defaulting for registry (dockerhub), repo (library),
// or tag (latest).
func StrictValidation(opts *Options) {
	opts.Strict = true
}

// WeakValidation is an Option that sets defaults when parsing names, see
// StrictValidation.
func WeakValidation(opts *Options) {
	opts.Strict = false
}

// Insecure is an Option that allows image references to be fetched without TLS.
func Insecure(opts *Options) {
	opts.Insecure = true
}
