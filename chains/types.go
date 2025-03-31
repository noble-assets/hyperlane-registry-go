// SPDX-License-Identifier: MIT
//
// Copyright 2025 NASD Inc. All Rights Reserved.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package chains

// Metadata defines the required chain medata of the Hyperlane Registry.
type Metadata struct {
	ChainId  string `yaml:"chainId"`
	DomainId uint32 `yaml:"domainId"`
	Name     string `yaml:"name"`
	Protocol string `yaml:"protocol"`
	RpcUrls  []struct {
		HTTP string `yaml:"http"`
	} `yaml:"rpcUrls"`
}
