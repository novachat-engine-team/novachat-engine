```
pinned: [ SKIPPED BY BIT 24 IN FIELD flags ],                                                                                                  
id: 176 [INT],
from_id: [ SKIPPED BY BIT 8 IN FIELD flags ],
peer_id: { peerUser
user_id: 6220850133 [LONG],
},     
fwd_from: [ SKIPPED BY BIT 2 IN FIELD flags ], 
via_bot_id: [ SKIPPED BY BIT 11 IN FIELD flags ],
reply_to: [ SKIPPED BY BIT 3 IN FIELD flags ],
date: 1689910316 [INT], 
message: "🗣Input you answer，CHATGPT ！\n        🗣 Please enter your question and CHATGPT will get back to you!!!\n\n            ✅>
media: [ SKIPPED BY BIT 9 IN FIELD flags ],
reply_markup: { replyInlineMarkup
            rows: [ vector<0x0> (2)
              { keyboardButtonRow
                buttons: [ vector<0x0> (2)
                  { keyboardButtonUrl                                                                                                                    
                    text: "VIP（Purchase）" [STRING],
                    url: "https://t.me/chatgptzs_MB_bot?start" [STRING],
                  },    
                  { keyboardButtonUrl
                    text: "JoinUs（Join US）" [STRING],
                    url: "https://t.me/chatgpts101" [STRING],
                  },    
                ],    
              },
              { keyboardButtonRow
                buttons: [ vector<0x0> (1)
                  { keyboardButtonCallback
                    flags: 0 [INT],
                    requires_password: [ SKIPPED BY BIT 0 IN FIELD flags ],
                    text: "name clib （Enter Mantra Club）" [STRING],
                    data: "show" [STRING],
                  },    
                ],    
              },
            ],  
          },
entities: [ vector<0x0> (6)
            { messageEntityMention                                                                                                                       
              offset: 467 [INT],
              length: 14 [INT],
            },
            { messageEntityMention
              offset: 526 [INT],
              length: 14 [INT],
            },
            { messageEntityBold
              offset: 614 [INT],
              length: 2 [INT],
            },
            { messageEntityBold
              offset: 693 [INT],
              length: 2 [INT],
            },
            { messageEntityBold
              offset: 725 [INT],
              length: 37 [INT],
            },
            { messageEntityBold
              offset: 777 [INT],
              length: 118 [INT],
            },
          ],
 
 

Send: { core_message
  msg_id: 7258116667905723088 [LONG],    
  seq_no: 80 [INT],                      
  bytes: 104 [INT],                      
  body: { msg_container                  
    messages: [ vector<0xffffffffffffffff> (2)
      { core_message                     
        msg_id: 7258116667905491260 [LONG],
        seq_no: 79 [INT],                
        bytes: 44 [INT],                 
        body: { messages_getBotCallbackAnswer
          flags: 5 [INT],                
          game: [ SKIPPED BY BIT 1 IN FIELD flags ],
          peer: { inputPeerUser          
            user_id: 6220850133 [LONG],   
            access_hash: 17801763351487381826 [LONG],
          },                             
          msg_id: 176 [INT],             
          data: "show" [STRING],         
          password: { inputCheckPasswordEmpty },
        },                               
      },                                 
      { core_message                     
        msg_id: 7258116667905708272 [LONG],
        seq_no: 80 [INT],                
        bytes: 20 [INT],                 
        body: { msgs_ack                 
          msg_ids: [ vector<0x22076cba> (1)
            7258116663832970241 [LONG],  
          ],                             
        },                               
      },                                 
    ],                                   
  },                                     
}
Recv: { core_message
  msg_id: 7258116680688923649 [LONG],
  seq_no: 79 [INT],
  bytes: 24 [INT],
  body: { rpc_result
    req_msg_id: 7258116667905491260 [LONG],
    result: { messages_botCallbackAnswer
      flags: 16 [INT],
      alert: [ SKIPPED BY BIT 1 IN FIELD flags ],
      has_url: [ SKIPPED BY BIT 3 IN FIELD flags ],
      native_ui: YES [ BY BIT 4 IN FIELD flags ],
      message: [ SKIPPED BY BIT 0 IN FIELD flags ],
      url: [ SKIPPED BY BIT 2 IN FIELD flags ],
      cache_time: 0 [INT],
    },           
  },             
}

Recv: { core_message
  msg_id: 7258116682230072321 [LONG],
  seq_no: 81 [INT],
  bytes: 2108 [INT],
  body: [GZIPPED] { updates
    updates: [ vector<0x0> (1)
      { updateNewMessage
        message: { message
          flags: 192 [INT], 
          out: [ SKIPPED BY BIT 1 IN FIELD flags ],
          mentioned: [ SKIPPED BY BIT 4 IN FIELD flags ], 
          media_unread: [ SKIPPED BY BIT 5 IN FIELD flags ],
          silent: [ SKIPPED BY BIT 13 IN FIELD flags ],
          post: [ SKIPPED BY BIT 14 IN FIELD flags ],
          from_scheduled: [ SKIPPED BY BIT 18 IN FIELD flags ],
          legacy: [ SKIPPED BY BIT 19 IN FIELD flags ], 
          edit_hide: [ SKIPPED BY BIT 21 IN FIELD flags ],
          pinned: [ SKIPPED BY BIT 24 IN FIELD flags ],
          id: 179 [INT],
          from_id: [ SKIPPED BY BIT 8 IN FIELD flags ],
          peer_id: { peerUser
            user_id: 6220850133 [LONG],
          },
          fwd_from: [ SKIPPED BY BIT 2 IN FIELD flags ], 
          via_bot_id: [ SKIPPED BY BIT 11 IN FIELD flags ],
          reply_to: [ SKIPPED BY BIT 3 IN FIELD flags ],
          date: 1689911978 [INT], 
          message: "hello xxxxx" [STRING],
          media: [ SKIPPED BY BIT 9 IN FIELD flags ],
          reply_markup: { replyInlineMarkup
            rows: [ vector<0x0> (1)
              { keyboardButtonRow
                buttons: [ vector<0x0> (1)
                  { keyboardButtonUrl
                    text: "（Join US）" [STRING],
                    url: "https://t.me/chatgpts101" [STRING],
                  },
                ],
              },
            ],  
          },    
          entities: [ vector<0x0> (8)
            { messageEntityCode
              offset: 989 [INT],
              length: 28 [INT],
            },  
            { messageEntityBotCommand
              offset: 1042 [INT],
              length: 9 [INT],
            },  
            { messageEntityBotCommand
              offset: 1059 [INT],
              length: 9 [INT],
            },  
            { messageEntityBotCommand
              offset: 1094 [INT],
              length: 9 [INT],
            },  
            { messageEntityCode
              offset: 1214 [INT],
              length: 45 [INT],
            },  
            { messageEntityCode
              offset: 1421 [INT],
              length: 190 [INT],
            },  
            { messageEntityBotCommand
              offset: 1663 [INT],
              length: 9 [INT],
            },  
            { messageEntityBotCommand
              offset: 1693 [INT],
              length: 4 [INT],
            },
          },
```