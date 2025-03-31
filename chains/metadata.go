// SPDX-License-Identifier: MIT
//
// Copyright 2025 NASD Inc. All Rights Reserved.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package chains

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed metadata.yaml
var MetadataFile []byte

var (
	MetadataByDomain = make(map[uint32]Metadata)
	MetadataByName   = make(map[string]Metadata)
)

func init() {
	err := yaml.Unmarshal(MetadataFile, &MetadataByDomain)
	if err != nil {
		panic("unable to unmarshal hyperlane chain metadata file")
	}

	for _, metadata := range MetadataByName {
		MetadataByDomain[metadata.DomainId] = metadata
	}
}

// GetName returns the chain name of the provided domain.
func GetName(domain uint32) string {
	metadata, exists := MetadataByDomain[domain]
	if !exists {
		return "unknown"
	}

	return metadata.Name
}
