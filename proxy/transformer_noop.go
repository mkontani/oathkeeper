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

package proxy

import (
	"encoding/json"
	"github.com/ory/oathkeeper/driver/configuration"
	"github.com/pkg/errors"
	"net/http"

	"github.com/ory/oathkeeper/rule"
)

type TransformerNoop struct{c configuration.Provider}

func NewCredentialsIssuerNoOp(c configuration.Provider) *TransformerNoop {
	return &TransformerNoop{c:c}
}

func (a *TransformerNoop) GetID() string {
	return "noop"
}

func (a *TransformerNoop) Transform(r *http.Request, session *AuthenticationSession, config json.RawMessage, rl *rule.Rule) (http.Header, error) {
	return r.Header, nil
}

func (a *TransformerNoop) Validate() error {
	if !a.c.TransformerIDTokenIsEnabled() {
		return errors.WithStack(ErrAuthenticatorNotEnabled.WithReasonf("Transformer % is disabled per configuration.", a.GetID()))
	}

	return nil
}
