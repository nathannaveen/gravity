// Copyright 2021 Gravitational Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package environment

import (
	"github.com/gravitational/gravity/e/lib/ops/client"
	"github.com/gravitational/gravity/lib/localenv"

	"github.com/gravitational/trace"
)

// Local extends the LocalEnvironment from open-source
type Local struct {
	*localenv.LocalEnvironment
}

// ClusterOperator returns the enterprise cluster operator client
func (l *Local) ClusterOperator() (*client.Client, error) {
	operator, err := l.LocalEnvironment.SiteOperator()
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return client.New(operator), nil
}

// ClusterOperator returns the enterprise cluster operator client
func ClusterOperator() (*client.Client, error) {
	operator, err := localenv.ClusterOperator()
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return client.New(operator), nil
}

// GetCurrentOpsCenter returns the currently active key store entry
func GetCurrentOpsCenter(keyStoreDir string) (string, error) {
	keyStore, err := localenv.GetLocalKeyStore(keyStoreDir)
	if err != nil {
		return "", trace.Wrap(err)
	}
	current := keyStore.GetCurrentOpsCenter()
	if current == "" {
		return "", trace.NotFound("not logged in")
	}
	return current, nil
}