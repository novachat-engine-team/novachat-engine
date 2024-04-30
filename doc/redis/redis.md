



### RedisBloom
centos7:
    yum -y install cmake

cd /usr/src

https://github.com/RedisBloom/RedisBloom.git

cd RedisBloom
git checkout v2.2.8
make

mkdir -p /usr/lib64/redis/modules/

cp redisbloom.so /usr/lib64/redis/modules/

ll /usr/lib64/redis/modules/redisbloom.so

-rwxr-xr-x. 1 root root 331992 Aug 19 15:21 /usr/lib64/redis/modules/redisbloom.so

chmod 755 /usr/lib64/redis/modules/redisbloom.so

redis version > 4.0.0

centos7:
    yum install -y http://rpms.famillecollet.com/enterprise/remi-release-7.rpm
    yum --enablerepo=remi install redis

    add: /etc/redis.conf
    loadmodule /usr/lib64/redis/modules/redisbloom.so

ubuntu:
    apt install redis

    add: /etc/redis/redis.conf
    loadmodule /usr/lib64/redis/modules/redisbloom.so
    
