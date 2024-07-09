
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

#### Installing
- [Ubuntu](https://github.com/novachat-engine-team/novachat-engine/blob/main/install/install_ubuntu.md)
- [CentOS7](https://github.com/novachat-engine-team/novachat-engine/blob/main/install/install_centos-7.md)
  
    MySQL default: root/123456

#### BUILDING
    git clone https://github.com/novachat-engine-team/novachat-engine
    1. cd novachat-engine
    2. cd gen
    3. mkdir out
    4. ./config.sh
    5. ./build.sh all
    6. cp start.sh restart.sh status.sh stop.sh out
    7. cat >> /etc/hosts << EOF 

    127.0.0.1 mysql-hosts 
    127.0.0.1 kafka-hosts 
    127.0.0.1 redis-hosts 
    127.0.0.1 etcd-hosts 
    127.0.0.1 mongodb-hosts
    EOF
    8. cd out
    9. ./start.sh all
       start serve
    10. ./status.sh
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
