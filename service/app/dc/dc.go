/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/4/19 17:46
 * @File : dc.go
 */

package dc

import "novachat_engine/mtproto"

const DefaultDc int32 = 2
const DefaultCountry string = "US"

func InstallDCOption(optionPath string) {

}

func GetNearestDc(thisDC int32) *mtproto.NearestDc {
	_ = thisDC

	n := mtproto.NewTLNearestDc(nil)
	n.SetThisDc(DefaultDc)
	n.SetNearestDc(DefaultDc)
	n.SetCountry(DefaultCountry)

	return n.To_NearestDc()
}

func GetDCByApiID(apiId int32, apiHash string) int32 {
	return DefaultDc
}