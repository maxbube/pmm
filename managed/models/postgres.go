// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package models

// PGParams represent postgreSQL server params.
type PGParams struct {
	// Addr represent postgresql connection address.
	Addr string
	// DBName represent postgresql database name.
	DBName string
	// DBUsername represent postgresql database username.
	DBUsername string
	// DBPassword represent postgresql database user password.
	DBPassword string
	// SSLMode represent postgresql database ssl mode.
	SSLMode string
	// SSLCAPath represent postgresql database root ssl CA cert path.
	SSLCAPath string
	// SSLKeyPath represent postgresql database user ssl key path.
	SSLKeyPath string
	// SSLCertPath represent postgresql database user ssl cert path.
	SSLCertPath string
}
