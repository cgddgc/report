// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// AssetsType is the golang structure for table assets_type.
type AssetsType struct {
	Id             uint        `orm:"id,primary"         json:"id"`              //
	TypeName       string      `orm:"type_name"          json:"type_name"`       //
	Attribution    string      `orm:"attribution,unique" json:"attribution"`     //
	Department     string      `orm:"department"         json:"department"`      //
	AssetsUsername string      `orm:"assets_username"    json:"assets_username"` //
	Subdomain      string      `orm:"subdomain"          json:"subdomain"`       //
	IntranetIp     string      `orm:"intranet_ip"        json:"intranet_ip"`     //
	PublicIp       string      `orm:"public_ip"          json:"public_ip"`       //
	CreateAt       *gtime.Time `orm:"create_at"          json:"create_at"`       //
}