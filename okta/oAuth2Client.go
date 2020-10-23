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

import ()

type OAuth2Client struct {
	Links      interface{} `json:"_links,omitempty"`
	ClientId   string      `json:"client_id,omitempty"`
	ClientName string      `json:"client_name,omitempty"`
	ClientUri  string      `json:"client_uri,omitempty"`
	LogoUri    string      `json:"logo_uri,omitempty"`
}

func NewOAuth2Client() *OAuth2Client {
	return &OAuth2Client{}
}

func (a *OAuth2Client) IsApplicationInstance() bool {
	return true
}
