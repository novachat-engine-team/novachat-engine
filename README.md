
### Novachat_engine
Novachat-engine is a open source mtproto server for Telegram compatible Android, IOS,tdesktop(Mac/Windows/Linux/Web).

## Introduce
[Telegram](https://telegram.org/)

#### Novachat_engine server
    Novachat_engine server base layer 122.


#### ENV
    
    golang version: GO 1.17.11
    
    GO111MODULE=auto
    GOPROXY=https://goproxy.io,direct
    
    go mod tidy -go=1.16 && go mod tidy -go=1.17
    

### DB
    mongodb 4.4.19
    mysql 5.7
    kafka (you can find in ROOT_DIR/lib)
    redis (must redisbloom found in ROOT_DIR/lib)
    etcd
    
#### BUILD
    1. goto ROOT_DIR/gen
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


#### Author 
PM [Author](https://t.me/bigM123)
PM [Author](https://t.me/bigM1223)

Groups [Telegram](https://t.me/novachat_telegram)
