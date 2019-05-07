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

package rule

import (
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"

	"github.com/ory/go-convenience/stringslice"
	"github.com/ory/herodot"
	"github.com/ory/oathkeeper/pipeline/authn"
	"github.com/ory/oathkeeper/pipeline/authz"
	"github.com/ory/oathkeeper/pipeline/mutate"
)

var methods = []string{
	"GET",
	"POST",
	"PUT",
	"HEAD",
	"DELETE",
	"PATCH",
	"OPTIONS",
	"TRACE",
	"CONNECT",
}

type validatorRegistry interface {
	authn.Registry
	authz.Registry
	mutate.Registry
}

type Validator interface {
	Validate(r *Rule) error
}

type ValidatorDefault struct {
	r validatorRegistry
}

func NewValidatorDefault(r validatorRegistry) *ValidatorDefault {
	return &ValidatorDefault{r: r}
}

func (v *ValidatorDefault) validateAuthenticators(r *Rule) error {
	if len(r.Authenticators) == 0 {
		return errors.WithStack(herodot.ErrInternalServerError.WithReason(`Value of "authenticators" must be set and can not be an empty array.`))
	}

	for k, a := range r.Authenticators {
		auth, err := v.r.PipelineAuthenticator(a.Handler)
		if err != nil {
			return herodot.ErrInternalServerError.WithReasonf(`Value "%s" of "authenticators[%d]" is not in list of supported authenticators: %v`, a.Handler, k, v.r.AvailablePipelineAuthenticators()).WithTrace(err).WithDebug(err.Error())
		}

		if err := auth.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (v *ValidatorDefault) validateAuthorizer(r *Rule) error {
	if r.Authorizer.Handler == "" {
		return errors.WithStack(herodot.ErrInternalServerError.WithReason(`Value of "authorizer.handler" can not be empty.`))
	}

	auth, err := v.r.PipelineAuthorizer(r.Authorizer.Handler)
	if err != nil {
		return errors.WithStack(herodot.ErrInternalServerError.WithReasonf(`Value "%s" of "authorizer.handler" is not in list of supported authorizers: %v`, r.Authorizer.Handler, v.r.AvailablePipelineAuthorizers()).WithTrace(err).WithDebug(err.Error()))
	}

	return auth.Validate()
}

func (v *ValidatorDefault) validateMutator(r *Rule) error {
	if r.Mutator.Handler == "" {
		return errors.WithStack(herodot.ErrInternalServerError.WithReason(`Value of "mutator.handler" can not be empty.`))
	}

	mutator, err := v.r.PipelineMutator(r.Mutator.Handler)
	if err != nil {
		return herodot.ErrInternalServerError.WithReasonf(`Value "%s" of "mutator.handler" is not in list of supported mutators: %v`, r.Mutator.Handler, v.r.AvailablePipelineMutators()).WithTrace(err).WithDebug(err.Error())
	}

	return mutator.Validate()
}

func (v *ValidatorDefault) Validate(r *Rule) error {
	if r.Match.URL == "" {
		return errors.WithStack(herodot.ErrInternalServerError.WithReasonf(`Value "%s" of "match.url" field is not a valid url.`, r.Match.URL))
	}

	for _, m := range r.Match.Methods {
		if !stringslice.Has(methods, m) {
			return errors.WithStack(herodot.ErrInternalServerError.WithReasonf(`Value "%s" of "match.methods" is not a valid HTTP method, valid methods are: %v`, m, methods))
		}
	}

	if r.Upstream.URL == "" {
		// Having no upstream URL is fine here because the judge does not need an upstream!
	} else if !govalidator.IsURL(r.Upstream.URL) {
		return errors.WithStack(herodot.ErrInternalServerError.WithReasonf(`Value "%s" of "upstream.url" is not a valid url.`, r.Upstream.URL))
	}

	if err := v.validateAuthenticators(r); err != nil {
		return err
	}

	if err := v.validateAuthorizer(r); err != nil {
		return err
	}

	if err := v.validateMutator(r); err != nil {
		return err
	}

	return nil
}
