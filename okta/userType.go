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
	"fmt"
	"time"
)

type UserTypeResource resource

type UserType struct {
	Links         interface{} `json:"_links,omitempty"`
	Created       *time.Time  `json:"created,omitempty"`
	CreatedBy     string      `json:"createdBy,omitempty"`
	Default       *bool       `json:"default,omitempty"`
	Description   string      `json:"description,omitempty"`
	DisplayName   string      `json:"displayName,omitempty"`
	Id            string      `json:"id,omitempty"`
	LastUpdated   *time.Time  `json:"lastUpdated,omitempty"`
	LastUpdatedBy string      `json:"lastUpdatedBy,omitempty"`
	Name          string      `json:"name,omitempty"`
}

// Updates an existing User Type
func (m *UserTypeResource) UpdateUserType(typeId string, body UserType) (*UserType, *Response, error) {
	url := fmt.Sprintf("/api/v1/meta/types/user/%v", typeId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var userType *UserType

	resp, err := m.client.requestExecutor.Do(req, &userType)
	if err != nil {
		return nil, resp, err
	}

	return userType, resp, nil
}

// Fetches a User Type by ID. The special identifier &#x60;default&#x60; may be used to fetch the default User Type.
func (m *UserTypeResource) GetUserType(typeId string) (*UserType, *Response, error) {
	url := fmt.Sprintf("/api/v1/meta/types/user/%v", typeId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var userType *UserType

	resp, err := m.client.requestExecutor.Do(req, &userType)
	if err != nil {
		return nil, resp, err
	}

	return userType, resp, nil
}

// Deletes a User Type permanently. This operation is not permitted for the default type, nor for any User Type that has existing users
func (m *UserTypeResource) DeleteUserType(typeId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/meta/types/user/%v", typeId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Fetches all User Types in your org
func (m *UserTypeResource) ListUserTypes() ([]*UserType, *Response, error) {
	url := fmt.Sprintf("/api/v1/meta/types/user")

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var userType []*UserType

	resp, err := m.client.requestExecutor.Do(req, &userType)
	if err != nil {
		return nil, resp, err
	}

	return userType, resp, nil
}

// Creates a new User Type. A default User Type is automatically created along with your org, and you may add another 9 User Types for a maximum of 10.
func (m *UserTypeResource) CreateUserType(body UserType) (*UserType, *Response, error) {
	url := fmt.Sprintf("/api/v1/meta/types/user")

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var userType *UserType

	resp, err := m.client.requestExecutor.Do(req, &userType)
	if err != nil {
		return nil, resp, err
	}

	return userType, resp, nil
}

// Replace an existing User Type
func (m *UserTypeResource) ReplaceUserType(typeId string, body UserType) (*UserType, *Response, error) {
	url := fmt.Sprintf("/api/v1/meta/types/user/%v", typeId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var userType *UserType

	resp, err := m.client.requestExecutor.Do(req, &userType)
	if err != nil {
		return nil, resp, err
	}

	return userType, resp, nil
}