// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	Id        uint64      `json:"id"         ` //
	Name      string      `json:"name"       ` //
	Body      string      `json:"body"       ` //
	Password  string      `json:"password"   ` //
	IsForbid  int         `json:"is_forbid"  ` //
	CreatedAt *gtime.Time `json:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updated_at" ` //
}
