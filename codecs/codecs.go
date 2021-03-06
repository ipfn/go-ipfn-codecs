// Copyright © 2017-2018 The IPFN Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package codecs registers IPFN codecs in CID codec types.
package codecs

type target struct {
	Codecs     map[string]uint64
	CodecToStr map[uint64]string
}

var targets []target

// Codecs - Maps the name of a codec to its type.
var Codecs = make(map[string]uint64)

// CodecToStr - Maps the numeric codec to its name.
var CodecToStr = make(map[uint64]string)

// Register - Registers codecs in remote cids package.
func Register(codecs map[string]uint64) {
	for _, target := range targets {
		for name, codec := range codecs {
			target.Codecs[name] = codec
			target.CodecToStr[codec] = name
		}
	}
}

// RegisterTarget - Registers codecs in remote cids package.
func RegisterTarget(codecs map[string]uint64, codecToStr map[uint64]string) {
	targets = append(targets, target{
		Codecs:     codecs,
		CodecToStr: codecToStr,
	})
	for name, codec := range Codecs {
		codecs[name] = codec
		codecToStr[codec] = name
	}
}

func init() {
	// this one should be always before
	// subsequent calls to RegisterTarget
	// are using these maps to clone later
	RegisterTarget(Codecs, CodecToStr)
	// this one is how we would import it to our package
	// codecs.RegisterTarget(cid.Codecs, cid.CodecToStr)
}
