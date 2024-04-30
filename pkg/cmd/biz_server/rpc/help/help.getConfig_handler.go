/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/04/13 16:46
 * @File : help.getConfig_handler.go
 * @Desc :
 *
 */

package rpc

import (
	"context"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/service/app/help"
	"time"
)

//  help.getConfig#c4f9186b = Config;
//
func (s *HelpServiceImpl) HelpGetConfig(ctx context.Context, request *mtproto.TLHelpGetConfig) (*mtproto.Config, error) {
	md := metadata.RpcMetaDataFromContext(ctx)
	log.Infof("HelpGetConfig %v, request: %v", metadata.RpcMetaDataDebug(md), request)

	conf := help.GetConfig()
	now := int32(time.Now().Unix())
	conf.Date = now
	conf.Expires = now + help.ExpiresTimeout
	return conf, nil
}

/*
"TL_config" : {
          "autoupdate_url_prefix" : "",
          "call_connect_timeout_ms" : 30000,
          "call_packet_timeout_ms" : 10000,
          "call_receive_timeout_ms" : 20000,
          "call_ring_timeout_ms" : 90000,
          "caption_length_max" : 1024,
          "channels_read_media_period" : 604800,
          "chat_size_max" : 200,
          "date" : 1622269883,
          "dc_options" : [
             {
                "cdn" : false,
                "flags" : 0,
                "id" : 1,
                "ip_address" : "149.154.175.60",
                "ipv6" : false,
                "isStatic" : false,
                "media_only" : false,
                "port" : 443,
                "tcpo_only" : false
             },
             {
                "cdn" : false,
                "flags" : 16,
                "id" : 1,
                "ip_address" : "149.154.175.53",
                "ipv6" : false,
                "isStatic" : true,
                "media_only" : false,
                "port" : 443,
                "tcpo_only" : false
             },
             {
                "cdn" : false,
                "flags" : 1,
                "id" : 1,
                "ip_address" : "2001:0b28:f23d:f001:0000:0000:0000:000a",
                "ipv6" : true,
                "isStatic" : false,
                "media_only" : false,
                "port" : 443,
                "tcpo_only" : false
             },
             {
                "cdn" : false,
                "flags" : 0,
                "id" : 2,
                "ip_address" : "149.154.167.51",
                "ipv6" : false,
                "isStatic" : false,
                "media_only" : false,
                "port" : 443,
                "tcpo_only" : false
             },
             {
                "cdn" : false,
                "flags" : 16,
                "id" : 2,
                "ip_address" : "149.154.167.51",
                "ipv6" : false,
                "isStatic" : true,
                "media_only" : false,
                "port" : 443,
                "tcpo_only" : false
             },
             {
                "cdn" : false,
                "flags" : 2,
                "id" : 2,
                "ip_address" : "149.154.167.151",
                "ipv6" : false,
                "isStatic" : false,
                "media_only" : true,
                "port" : 443,
                "tcpo_only" : false
             },
             {
                "cdn" : false,
                "flags" : 1,
                "id" : 2,
                "ip_address" : "2001:067c:04e8:f002:0000:0000:0000:000a",
                "ipv6" : true,
                "isStatic" : false,
                "media_only" : false,
                "port" : 443,
                "tcpo_only" : false
             },
             {
                "cdn" : false,
                "flags" : 3,
                "id" : 2,
                "ip_address" : "2001:067c:04e8:f002:0000:0000:0000:000b",
                "ipv6" : true,
                "isStatic" : false,
                "media_only" : true,
                "port" : 443,
                "tcpo_only" : false
             },
             {
                "cdn" : false,
                "flags" : 0,
                "id" : 3,
                "ip_address" : "149.154.175.100",
                "ipv6" : false,
                "isStatic" : false,
                "media_only" : false,
                "port" : 443,
                "tcpo_only" : false
             },
             {
                "cdn" : false,
                "flags" : 16,
                "id" : 3,
                "ip_address" : "149.154.175.100",
                "ipv6" : false,
                "isStatic" : true,
                "media_only" : false,
                "port" : 443,
                "tcpo_only" : false
             },
             {
                "cdn" : false,
                "flags" : 1,
                "id" : 3,
                "ip_address" : "2001:0b28:f23d:f003:0000:0000:0000:000a",
                "ipv6" : true,
                "isStatic" : false,
                "media_only" : false,
                "port" : 443,
                "tcpo_only" : false
             },
             {
                "cdn" : false,
                "flags" : 0,
                "id" : 4,
                "ip_address" : "149.154.167.92",
                "ipv6" : false,
                "isStatic" : false,
                "media_only" : false,
                "port" : 443,
                "tcpo_only" : false
             },
             {
                "cdn" : false,
                "flags" : 16,
                "id" : 4,
               "ip_address" : "149.154.167.92",
                "ipv6" : false,
                "isStatic" : true,
                "media_only" : false,
                "port" : 443,
                "tcpo_only" : false
             },
             {
                "cdn" : false,
                "flags" : 1,
                "id" : 4,
                "ip_address" : "2001:067c:04e8:f004:0000:0000:0000:000a",
                "ipv6" : true,
                "isStatic" : false,
                "media_only" : false,
                "port" : 443,
                "tcpo_only" : false
             },
             {
                "cdn" : false,
                "flags" : 2,
                "id" : 4,
                "ip_address" : "149.154.166.120",
                "ipv6" : false,
                "isStatic" : false,
                "media_only" : true,
                "port" : 443,
                "tcpo_only" : false
             },
             {
                "cdn" : false,
                "flags" : 3,
                "id" : 4,
                "ip_address" : "2001:067c:04e8:f004:0000:0000:0000:000b",
                "ipv6" : true,
                "isStatic" : false
                "media_only" : true,
                "port" : 443,
                "tcpo_only" : false
             },
             {
                "cdn" : false,
                "flags" : 1,
                "id" : 5,
                "ip_address" : "2001:0b28:f23f:f005:0000:0000:0000:000a",
                "ipv6" : true,
                "isStatic" : false,
                "media_only" : false,
                "port" : 443,
                "tcpo_only" : false
             },
             {
                "cdn" : false,
                "flags" : 0,
                "id" : 5,
                "ip_address" : "91.108.56.124",
                "ipv6" : false,
                "isStatic" : false,
                "media_only" : false,
                "port" : 443,
                "tcpo_only" : false
             },
             {
                "cdn" : false,
                "flags" : 16,
                "id" : 5,
                "ip_address" : "91.108.56.124",
                "ipv6" : false,
                "isStatic" : true,
                "media_only" : false,
                "port" : 443,
                "tcpo_only" : false
             }
          ],
          "c_options.size" : "19",
          "dc_txt_domain_name" : "apv3.stel.com",
          "edit_time_limit" : 172800,
          "expires" : 1622273149,
          "flags" : 3660,
          "forwarded_count_max" : 100,
          "gif_search_username" : "gif",
          "img_search_username" : "bing",
          "lang_pack_version" : 6799494,
          "me_url_prefix" : "https://t.me/",
          "megagroup_size_max" : 200000,
          "message_length_max" : 4096,
          "notify_cloud_delay_ms" : 30000,
          "notify_default_delay_ms" : 1500,
          "offline_blur_timeout_ms" : 5000,
          "offline_idle_timeout_ms" : 30000,
          "online_cloud_timeout_ms" : 300000,
          "online_update_period_ms" : 210000,
          "pinned_dialogs_count_max" : 5,
          "push_chat_limit" : 2,
          "push_chat_period_ms" : 60000,
          "rating_e_decay" : 2419200,
          "revoke_pm_time_limit" : 2147483647,
          "revoke_time_limit" : 2147483647,
          "saved_gifs_limit" : 200,
          "static_maps_provider" : "",
          "stickers_faved_limit" : 5,
          "stickers_recent_limit" : 200,
          "suggested_lang_code" : "en",
          "test_mode" : false,
          "this_dc" : 2,
          "tmp_sessions" : 0,
          "venue_search_username" : "foursquare",
          "webfile_dc_id" : 4
       },
       "messageId" : "0",
       "messageLength" : 0,
       "messageSeqNo" : 0
    }
*/
