

### mongodb 
dbpath = /opt/mongodb/data 
logpath = /opt/mongodb/logs/mongodb.log
port = 27017
fork = true
# nohttpinterface = true
# Error parsing INI config file: unrecognised option 'nohttpinterface'
#auth=true
bind_ip=0.0.0.0
replSet=rs0
wiredTigerCacheSizeGB=0.5

### mongodb 4.2
bind_ip=0.0.0.0
port=27017
dbpath=/opt/mongodb-4.2.15/data/
logpath=/opt/mongodb-4.2.15/logs/mongodb.log
pidfilepath =/opt/mongodb-4.2.15/run/mongodb.pid
logappend=true
fork=true
#maxConns=500
#noauth = true
wiredTigerCacheSizeGB=0.5


######

/opt/mongodb-4.2.15/bin/mongoexport --host=192.168.95.132 --port=27017 -d contacts -c contacts_4 -o  all/contacts_4.json 
/opt/mongodb-4.2.15/bin/mongoimport --host=192.168.95.132 --port=27017 -d contacts -c contacts_4  all/contacts_4.json

/opt/mongodb-4.2.15/bin/mongodump --host=192.168.95.132 --port=27017 -d contacts -o .

/opt/mongodb-4.2.15/bin/mongorestore --host=192.168.95.132 --port=27017 -d contacts contacts/


mongo

use access_hash
db.access_hash.insert({"_id": "access_hash"})

use chats
db.chats.insert({"_id": "chats"})

use contacts
db.chats.insert({"_id": "contacts"})

use conversations
db.chats.insert({"_id": "conversations"})

use file_info
db.chats.insert({"_id": "file_info"})

use message_boxes
db.chats.insert({"_id": "message_boxes"})

use user_status
db.chats.insert({"_id": "user_status"})

use user_updates
db.chats.insert({"_id": "user_updates"})



