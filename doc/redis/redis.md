

### RedisBloom

found redisbloom.so.zip goto `copy redisbloom.so.zip` or copy redis start error

otherwise goto `building RedisBloom`

#### Copy redisbloom.so.zip
    ubuntu:
        sudo apt install unzip -y
    centos7:
        sudo yum install unzip -y

    cp $ROOT_DIR/lib/redisbloom.so.zip
    unzip redisbloom.so.zip

    sudo mkdir -p /usr/lib64/redis/modules/
    sudo cp redisbloom.so /usr/lib64/redis/modules/

    centos7:
        chown -R redis:redis /usr/lib64/redis
        chmod 755 /usr/lib64/redis/modules/redisbloom.so

    // add: /etc/redis.conf
    loadmodule /usr/lib64/redis/modules/redisbloom.so

    sudo systemctl restart redis

#### Building RedisBloom

    centos7:
        yum -y install cmake
        yum -y install gcc
    ubuntu: 
        apt install cmake -y
    
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
    
        chown -R redis:redis /usr/lib64/redis
        chmod 755 /usr/lib64/redis/modules/redisbloom.so
        add: /etc/redis.conf
        loadmodule /usr/lib64/redis/modules/redisbloom.so
        sudo systemctl restart redis
    
    ubuntu:
        apt install redis
    
        add: /etc/redis/redis.conf
        loadmodule /usr/lib64/redis/modules/redisbloom.so
    
        sudo systemctl restart redis
        
