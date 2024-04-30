/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File :
 * @Desc :
 *
 */

package data_users

import (
	"database/sql"
	"time"
)

type Users struct {
	Id                int64          `json:"id,omitempty" db:"id" gorm:"primary key;AUTO_INCREMENT;not null;comment:ID" `
	Username          sql.NullString `json:"username,omitempty" db:"username" gorm:"unique;column:username;type:varchar(255) COLLATE utf8mb4_unicode_ci; COMMENT:-"`
	Phone             sql.NullString `json:"phone,omitempty" db:"phone" gorm:"unique;column:phone;type:varchar(255) COLLATE utf8mb4_unicode_ci; COMMENT:-" `
	AccessHash        int64          `json:"access_hash,omitempty" db:"access_hash" gorm:"column:access_hash;default:0; comment:-" `
	FirstName         string         `json:"first_name,omitempty" db:"first_name" gorm:"type:varchar(255) COLLATE utf8mb4_unicode_ci;default:'';not null; COMMENT:-"`
	LastName          string         `json:"last_name,omitempty" db:"last_name" gorm:"type:varchar(255) COLLATE utf8mb4_unicode_ci;default:'';not null; COMMENT:-"`
	Verified          int8           `json:"verified,omitempty" db:"verified" gorm:"default:0;not null; COMMENT:-"`
	Support           uint8          `json:"support,omitempty" db:"support" gorm:"default:0;not null; COMMENT:-"`
	Bot               uint8          `json:"bot,omitempty" db:"bot" gorm:"default:0;not null; COMMENT:-"`
	About             string         `json:"about,omitempty" db:"about" gorm:"default:'';not null;size:512;type:varchar(512) COLLATE utf8mb4_unicode_ci; COMMENT:-"`
	DaysTtl           int32          `json:"days_ttl,omitempty" db:"days_ttl" gorm:"not null; default:180; COMMENT:-"`
	Photos            string         `json:"photos,omitempty" db:"photos" gorm:"default:'';not null;size:1024;type:varchar(1024) COLLATE utf8mb4_unicode_ci; COMMENT:-"`
	Restricted        int8           `json:"restricted,omitempty" db:"restricted" gorm:"default:0;not null; COMMENT:-"`
	RestrictionReason string         `json:"restriction_reason,omitempty" db:"restriction_reason" gorm:"default:'';not null;size:128;type:varchar(128) COLLATE utf8mb4_unicode_ci; COMMENT:-"`
	Deleted           int8           `json:"deleted,omitempty" db:"deleted" gorm:"default:0;not null; COMMENT:-"`
	Password          string         `json:"password,omitempty" db:"password" gorm:"default:'';not null;size:512;type:varchar(512) COLLATE utf8mb4_unicode_ci; COMMENT:-"`
	CreatedAt         time.Time      `json:"created_at,omitempty" db:"created_at" sqlca:"readonly"  gorm:"not null;type:TIMESTAMP;default: CURRENT_TIMESTAMP;<-:create; COMMENT:-"`
	UpdatedAt         time.Time      `json:"updated_at,omitempty" db:"updated_at" sqlca:"readonly" gorm:"not null;type:TIMESTAMP;default: CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP; COMMENT:-"`
}

func (m *Users) GetUsername() string {
	return m.Username.String
}

func (m *Users) GetPhone() string {
	return m.Phone.String
}

func (m *Users) GetDeleted() bool {
	return m.Deleted != 0
}

func (m *Users) GetVerified() bool {
	return m.Verified != 0
}
func (m *Users) GetSupport() bool {
	return m.Support != 0
}
func (m *Users) GetBot() bool {
	return m.Bot != 0
}
func (m *Users) GetRestricted() bool {
	return m.Restricted != 0
}
