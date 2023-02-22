// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2023 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReleaseAuthor release author
//
// swagger:model releaseAuthor
type ReleaseAuthor struct {

	// avatar url
	AvatarURL string `json:"avatar_url,omitempty"`

	// events url
	EventsURL string `json:"events_url,omitempty"`

	// followers url
	FollowersURL string `json:"followers_url,omitempty"`

	// following url
	FollowingURL string `json:"following_url,omitempty"`

	// gists url
	GistsURL string `json:"gists_url,omitempty"`

	// gravatar id
	GravatarID string `json:"gravatar_id,omitempty"`

	// html url
	HTMLURL string `json:"html_url,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// login
	Login string `json:"login,omitempty"`

	// node id
	NodeID string `json:"node_id,omitempty"`

	// organizations url
	OrganizationsURL string `json:"organizations_url,omitempty"`

	// received events url
	ReceivedEventsURL string `json:"receivedEvents_url,omitempty"`

	// repos url
	ReposURL string `json:"repos_url,omitempty"`

	// site admin
	SiteAdmin bool `json:"site_admin,omitempty"`

	// starred url
	StarredURL string `json:"starred_url,omitempty"`

	// subscriptions url
	SubscriptionsURL string `json:"subscriptions_url,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this release author
func (m *ReleaseAuthor) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this release author based on context it is used
func (m *ReleaseAuthor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseAuthor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseAuthor) UnmarshalBinary(b []byte) error {
	var res ReleaseAuthor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}