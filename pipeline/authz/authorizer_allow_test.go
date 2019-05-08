/*
 * Copyright © 2017-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author       Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright  2017-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license  	   Apache-2.0
 */

package authz

import (
	"testing"

	"github.com/spf13/viper"

	"github.com/ory/oathkeeper/driver/configuration"
	"github.com/ory/oathkeeper/internal"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthorizerAllow(t *testing.T) {
	conf := internal.NewConfigurationWithDefaults()
	reg := internal.NewRegistry(conf)

	a, err := reg.PipelineAuthorizer("allow")
	require.NoError(t, err)
	assert.Equal(t, "allow", a.GetID())

	t.Run("case=never returns an error", func(t *testing.T) {
		require.NoError(t, a.Authorize(nil, nil, nil, nil))
	})

	t.Run("case=validate enabled/disabled", func(t *testing.T) {
		viper.Set(configuration.ViperKeyAuthorizerAllowIsEnabled, true)
		require.NoError(t, a.Validate())

		viper.Set(configuration.ViperKeyAuthorizerAllowIsEnabled, false)
		require.Error(t, a.Validate())
	})
}