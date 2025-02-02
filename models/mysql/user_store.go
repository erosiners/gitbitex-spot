// Copyright 2019 GitBitEx.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mysql

import (
	"github.com/gitbitex/gitbitex-spot/models"
	"github.com/jinzhu/gorm"
	"time"
)

func (s *Store) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := s.db.Raw("SELECT * FROM g_user WHERE email=?", email).Scan(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}

func (s *Store) AddUser(user *models.User) error {
	user.CreatedAt = time.Now()
	return s.db.Create(user).Error
}
