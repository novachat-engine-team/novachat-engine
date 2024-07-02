
### Novachat_engine
Novachat-engine is a open source mtproto server for Telegram compatible Android, IOS,tdesktop(Mac/Windows/Linux/Web).
## Introduce [Telegram](https://telegram.org/)

#### Novachat_engine server
    Novachat_engine server base layer 122.

##### Features
- contacts
- sms
- channel/megagroup(unsupport broadcast)
- voice call
- sticketsets

#### ENV

    golang version: GO 1.17.11

    GO111MODULE=auto
    GOPROXY=https://goproxy.io,direct

    go mod tidy -go=1.16 && go mod tidy -go=1.17


### Support Components
    mongodb 4.4.19
    mysql 5.7
    kafka (you can find in $ROOT_DIR/lib)
    redis (must redisbloom found in $ROOT_DIR/lib)
    etcd

#### BUILDING
    1. goto $ROOT_DIR/gen
    2. mkdir out
    3. ./config.sh
    4. ./build.sh all
    5. cp start.sh restart.sh status.sh stop.sh out
    6. config env in hosts 
        ip mysql-hosts 
        ip kafka-hosts 
        ip redis-hosts 
        ip etcd-hosts 
        ip mongodb-hosts
    7. cd out
    8. ./start.sh all
       start serve
    9. ./status.sh
       check serve is running, `ls log` all serve logs


tips: building error fix in [INSTALL_fix](https://github.com/novachat-engine-team/novachat-engine/blob/main/INSTALL_fix)

#### Notes
Will coming soon

- web
- videocall
- audiocall group call
- video group call(need LAYER 131)
- bots


#### Author 
PM [Author](https://t.me/bigM1223)

Groups [Telegram](https://t.me/novachat_telegram)
