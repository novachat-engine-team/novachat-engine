/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : auth.exportAuthorization_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/hack"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/pkg/util"
	"novachat_engine/service/constants"
)

/*
	## Authorization Transfer
	The following methods can be used to eliminate the need for users to enter the code from a text message every time:

	```
	auth.exportedAuthorization#df969c2d id:int bytes:bytes = auth.ExportedAuthorization;
	auth.authorization#f6b673a4 expires:int user:User = auth.Authorization;
	---functions---
	auth.importAuthorization#e3ef9613 id:int bytes:bytes = auth.Authorization;
	auth.exportAuthorization#e5bfffcd dc_id:int = auth.ExportedAuthorization;
	```

	auth.exportAuthorization must be executed in the current DC (the DC with which a connection has already been established),
	passing in dc_id as the value for the new DC.
	The method should return the user identifier and a long string of random data.
	An import operation can be performed at the new DC by sending it what was received.
	Queries requiring authorization can then be successfully executed in the new DC.
*/

//  auth.exportAuthorization#e5bfffcd dc_id:int = auth.ExportedAuthorization;
//
func (s *AuthServiceImpl) AuthExportAuthorization(ctx context.Context, request *mtproto.TLAuthExportAuthorization) (*mtproto.Auth_ExportedAuthorization, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("AuthExportAuthorization %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	exported := mtproto.NewTLAuthExportedAuthorization(&mtproto.Auth_ExportedAuthorization{
		Id:    constants.PeerTypeFromChannelIDType(md.UserId).ToInt32(),
		Bytes: hack.Bytes(util.Int32ToString(request.DcId)),
	})

	//log.Debugf("AuthExportAuthorization - reply: %s", log.JsonDebugData(exported))
	return exported.To_Auth_ExportedAuthorization(), nil
}
