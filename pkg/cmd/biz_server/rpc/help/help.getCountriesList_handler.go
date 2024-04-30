/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : help.getCountriesList_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/app/country"
)

//  help.getCountriesList#735787a8 lang_code:string hash:int = help.CountriesList;
//
func (s *HelpServiceImpl) HelpGetCountriesList(ctx context.Context, request *mtproto.TLHelpGetCountriesList) (*mtproto.Help_CountriesList, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("HelpGetCountriesList %v, request: %v", metadata.RpcMetaDataDebug(md), request)
	countryList := country.GetCountryList()
	return countryList, nil
}
