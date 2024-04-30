/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : messages.reorderStickerSets_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	syncClient "novachat_engine/pkg/cmd/sync/rpc_client"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/constants"
	"sort"
)

//  messages.reorderStickerSets#78337739 flags:# masks:flags.0?true order:Vector<long> = Bool;
//
func (s *MessagesServiceImpl) MessagesReorderStickerSets(ctx context.Context, request *mtproto.TLMessagesReorderStickerSets) (*mtproto.Bool, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("MessagesReorderStickerSets %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	if len(request.Order) > 0 {
		sort.SliceStable(request.Order, func(i, j int) bool {
			return request.Order[i] < request.Order[j]
		})
		_, err := syncClient.GetSyncClientById(md.UserId).ReqSyncUpdate(ctx, &syncClient.SyncUpdate{
			UserId: md.UserId,
			Updates: mtproto.NewTLUpdates(&mtproto.Updates{
				Updates: []*mtproto.Update{mtproto.NewTLUpdateStickerSetsOrder(&mtproto.Update{
					OrderBB2D20171: request.Order,
					Masks:          request.Masks,
				}).To_Update()},
			}).To_Updates(),
			Push:     false,
			PeerType: constants.PeerTypeUser.ToInt32(),
		})
		if err != nil {
			log.Warnf("MessagesReorderStickerSets error:%s", err.Error())
		}
	}
	return mtproto.ToMTBool(true), nil
}
