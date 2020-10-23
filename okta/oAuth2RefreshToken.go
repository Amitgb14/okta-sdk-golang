/*
* Copyright 2018 - Present Okta, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

// AUTO-GENERATED!  DO NOT EDIT FILE DIRECTLY

package okta

import (
	"time"
)

type OAuth2RefreshToken struct {
	Embedded    interface{}  `json:"_embedded,omitempty"`
	Links       interface{}  `json:"_links,omitempty"`
	ClientId    string       `json:"clientId,omitempty"`
	Created     *time.Time   `json:"created,omitempty"`
	CreatedBy   *OAuth2Actor `json:"createdBy,omitempty"`
	ExpiresAt   *time.Time   `json:"expiresAt,omitempty"`
	Id          string       `json:"id,omitempty"`
	Issuer      string       `json:"issuer,omitempty"`
	LastUpdated *time.Time   `json:"lastUpdated,omitempty"`
	Scopes      []string     `json:"scopes,omitempty"`
	Status      string       `json:"status,omitempty"`
	UserId      string       `json:"userId,omitempty"`
}

func NewOAuth2RefreshToken() *OAuth2RefreshToken {
	return &OAuth2RefreshToken{}
}

func (a *OAuth2RefreshToken) IsApplicationInstance() bool {
	return true
}
