
### ubuntu version >= 2004 

#### 1. Install MySQL
    wget https://dev.mysql.com/get/mysql-apt-config_0.8.12-1_all.deb
    sudo dpkg -i mysql-apt-config_0.8.12-1_all.deb
    or
    sudo dpkg-reconfigure mysql-apt-config

    select MySQL Server & Cluster (Currently select: mysql-5.7)
    select mysql-5.7
    select ok

    sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 467B942D3A79BD29
    sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys B7B3B788A8D3785C
    sudo apt update

    sudo apt install -f mysql-client=5.7* mysql-community-server=5.7* mysql-server=5.7*

    input root password: 123456

    sudo systemctl enable mysql
    sudo systemctl start mysql

#### 2. Install Etcd

    ETCD_VER=v3.2.26
    ETCD_DIR=/tmp/etcd
    ETCD_FILENAME=etcd-${ETCD_VER}-linux-amd64.tar.gz

    GITHUB_URL=https://github.com/etcd-io/etcd/releases/download
    DOWNLOAD_URL=${GITHUB_URL}

    rm -rf ${ETCD_DIR} && mkdir -p ${ETCD_DIR}

    curl -L ${DOWNLOAD_URL}/${ETCD_VER}/${ETCD_FILENAME} -o ${ETCD_DIR}/${ETCD_FILENAME}

    cd ${ETCD_DIR}
    tar xzvf ${ETCD_DIR}/${ETCD_FILENAME} -C ${ETCD_DIR} --strip-components=1
    rm -f ${ETCD_DIR}/${ETCD_FILENAME}

    ${ETCD_DIR}/etcd --version
    ${ETCD_DIR}/etcdctl version

    # start a local etcd server
    ${ETCD_DIR}/etcd &

#### 3. Install KAFKA
    
    sudo apt install openjdk-11-jdk

    KAFKA_DIR=/tmp/kafka
    KAFKA_FILENAME=kafka_2.12-2.5.0.tgz
    KAFKA=kafka_2.12-2.5.0

    rm -rf ${KAFKA_DIR} && mkdir -p ${KAFKA_DIR}

    # online kafka
        curl -L https://archive.apache.org/dist/kafka/2.5.0/${KAFKA_FILENAME} -o ${KAFKA_DIR}/${KAFKA_FILENAME}
    # local kafka
        cp $ROOT_DIR/lib/kafka_2.12-2.5.0.tar.gz ${KAFKA_DIR}/${KAFKA_FILENAME}

    cd ${KAFKA_DIR}    
    tar -vxzf ${KAFKA_FILENAME}

    cd ${KAFKA}
    cd kafka_2.12-2.5.0/

    bin/zookeeper-server-start.sh -daemon config/zookeeper.properties
    
    export KAFKA_HEAP_OPTS="-Xmx1G -Xms1G"
    bin/kafka-server-start.sh -daemon config/server.properties

#### 4. Install MongoDB

    MONGODB_DIR=/tmp/mongodb
    MONGODB_FILENAME=mongodb-linux-x86_64-4.4.19.tgz

    rm -rf ${MONGODB_DIR} && mkdir -p ${MONGODB_DIR} && mkdir -p ${MONGODB_DIR}/data 

    curl -L https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-ubuntu2004-4.4.19.tgz -o ${MONGODB_DIR}/${MONGODB_FILENAME}

    cd ${MONGODB_DIR}
    tar -vxzf ${MONGODB_FILENAME}
    mv mongodb-linux-x86_64-ubuntu2004-4.4.19 mongodb-linux-x86_64-4.4.19
    cd mongodb-linux-x86_64-4.4.19

    cat > mongodb.conf << EOF
    systemLog:
        destination: file
        logAppend: true
        path: /tmp/mongodb/data/mongod.log
    storage:
        dbPath: /tmp/mongodb/data
        journal:
            enabled: true
        engine: "wiredTiger"
        wiredTiger:
            engineConfig:
                cacheSizeGB: 0.5
    processManagement:
        fork: true  # fork and run in background
        pidFilePath: /tmp/mongodb/data/mongod.pid  # location of pidfile
        timeZoneInfo: /usr/share/zoneinfo

    net:
        port: 27017
        bindIp: 0.0.0.0 # Enter 0.0.0.0,:: to bind to all IPv4 and IPv6 addresses or, alternatively, use the net.bindIpAll setting.
    replication:
        replSetName: rs0
    EOF

    
    sudo bin/mongod -f mongodb.conf
    export PATH=$PATH:/tmp/mongodb/mongodb-linux-x86_64-4.4.19/bin

#### 5. Install redis
    sudo apt install redis