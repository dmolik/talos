// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ntp

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/talos-systems/talos/internal/app/machined/pkg/runtime"
	"github.com/talos-systems/talos/pkg/config/types/v1alpha1"
)

type NtpSuite struct {
	suite.Suite
}

func TestNtpSuite(t *testing.T) {
	suite.Run(t, new(NtpSuite))
}

func (suite *NtpSuite) TestQuery() {
	testServer := "time.cloudflare.com"
	// Create ntp client
	n, err := NewNTPClient(WithServer(testServer))
	suite.Assert().NoError(err)

	_, err = n.Query()
	suite.Assert().NoError(err)
}

func (suite *NtpSuite) TestNtpConfig() {
	server := "time.cloudflare.com"

	// Test unset config, single server config, multiple server config
	for _, conf := range []runtime.Configurator{&v1alpha1.Config{MachineConfig: &v1alpha1.MachineConfig{}}, sampleConfigSingleServer(), sampleConfigMultipleServers()} {
		// Check if ntp servers are defined
		// Support for only a single time server currently
		if len(conf.Machine().Time().Servers()) >= 1 {
			server = conf.Machine().Time().Servers()[0]
		}

		n, err := NewNTPClient(
			WithServer(server),
		)
		suite.Assert().NoError(err)
		suite.Assert().Equal(server, n.Server)
	}
}

func sampleConfigSingleServer() runtime.Configurator {
	return &v1alpha1.Config{
		MachineConfig: &v1alpha1.MachineConfig{
			MachineTime: &v1alpha1.TimeConfig{
				TimeServers: []string{"my.timeserver.org.biz.highfive"},
			},
		},
	}
}

func sampleConfigMultipleServers() runtime.Configurator {
	return &v1alpha1.Config{
		MachineConfig: &v1alpha1.MachineConfig{
			MachineTime: &v1alpha1.TimeConfig{
				TimeServers: []string{"my.timeserver.org.biz.highfive", "another.my.timeserver.gov.org.https.time"},
			},
		},
	}
}
