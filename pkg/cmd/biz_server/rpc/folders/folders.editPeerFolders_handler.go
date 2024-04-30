/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : folders.editPeerFolders_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
)

//  folders.editPeerFolders#6847d0ab folder_peers:Vector<InputFolderPeer> = Updates;
//
func (s *FoldersServiceImpl) FoldersEditPeerFolders(ctx context.Context, request *mtproto.TLFoldersEditPeerFolders) (*mtproto.Updates, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("FoldersEditPeerFolders %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	panic("FoldersEditPeerFolders")

	//seqPts, err := pts_sequence.GetPtsSequence().NextPts(md.UserId, 0)
	//if err != nil {
	//    log.Errorf("FoldersEditPeerFolders %v, request: %v error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    return nil, err
	//}
	//
	//includePeerList := make([]*dbproxy_pb.ConversationPeer, 0, len(request.FolderPeers))
	//for _, v := range request.FolderPeers {
	//    if v == nil || v.Peer == nil {
	//        continue
	//    }
	//
	//    inputPeer := input.MakeInputPeer(v.Peer)
	//    includePeerList = append(includePeerList, &dbproxy_pb.ConversationPeer{
	//        UserId:   md.UserId,
	//        PeerId:   inputPeer.GetPeerId(),
	//        PeerType: inputPeer.GetPeerType().ToInt32(),
	//    })
	//}
	//_, err = dbproxyClient.GetDBProxyClient().ConversationEditFolder(context.Background(), &dbproxy_pb.ConversationEditFolder{
	//    FolderId:        constants.ArchivedFolderType,
	//    UserId:          md.UserId,
	//    IncludePeer:     includePeerList,
	//})
	//if err != nil {
	//    log.Errorf("FoldersEditPeerFolders %v, request: %v ConversationEditFolder error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    return nil, err
	//}
	//
	////  updateFolderPeers#19360dc0 folder_peers:Vector<FolderPeer> pts:int pts_count:int = Update;
	//folderPeers := make([]*mtproto.FolderPeer, 0, len(request.FolderPeers))
	//for _, v := range request.FolderPeers {
	//    if v.Peer.ClassName == mtproto.ClassInputPeerEmpty {
	//        continue
	//    }
	//    folderPeers = append(folderPeers, mtproto.NewTLFolderPeer(&mtproto.FolderPeer{
	//        Peer:                 input.MakeInputPeer(v.Peer).ToPeer(),
	//        FolderId:             v.FolderId,
	//    }).To_FolderPeer())
	//}
	//update := mtproto.NewTLUpdateFolderPeers(&mtproto.Update{
	//    FolderPeers: folderPeers,
	//    PtsCount: 1,
	//    Pts: seqPts.ToInt32(),
	//})
	//
	//updates := mtproto.NewTLUpdateShort(&mtproto.Updates{
	//    Update:         update.To_Update(),
	//    Date:           int32(time.Now().Unix()),
	//})
	//
	//go func() {
	//    _, err = syncClient.GetSyncClient().PushUpdates(context.Background(), &sync_pb.PushUpdates{
	//        UserId:               md.UserId,
	//        ExcludeAuthKeyIdList: []int64{md.AuthKeyId},
	//        Updates:              updates.To_Updates(),
	//    })
	//    if err != nil {
	//        log.Errorf("FoldersEditPeerFolders %v, request: %v PushUpdates error:%s", metadata.RpcMetaDataDebug(md), request, err.Error())
	//    }
	//}()
	//
	//return updates.To_Updates(), nil
}
