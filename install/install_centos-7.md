
### CentOS 7
    2024-6-30 EOL, End Of Life

    sudo curl -o /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo
    sudo yum clean all
    sudo yum makecache

    or
    sudo yum install -y centos-release-vault
    sudo yum-config-manager --enable centosplus-vault
    sudo yum-config-manager --enable powertools-vault

#### 1. Install dnf
    sudo yum install epel-release -y
    sudo yum install dnf -y

#### 2. Install MySQL
    sudo dnf install wget -y
	wget https://dev.mysql.com/get/mysql57-community-release-el7-9.noarch.rpm
	sudo rpm -ivh mysql57-community-release-el7-9.noarch.rpm
    sudo rpm --import https://repo.mysql.com/RPM-GPG-KEY-mysql-2022

	sudo dnf install mysql-server
    sudo systemctl start mysqld
	sudo cat /var/log/mysqld.log | grep 'password'

    mysql -uroot -p 

    ALTER USER 'root'@'localhost' IDENTIFIED BY '23ZNsdfrhY7DzB#XPdfdsf';
    SHOW VARIABLES LIKE 'validate_password%';
    set global validate_password_policy=0;
    set global validate_password_length=6;
    ALTER USER 'root'@'localhost' IDENTIFIED BY '123456';

#### 3. Install Etcd

    ETCD_VER=v3.2.26
    ETCD_DIR=/tmp/etcd
    ETCD_FILENAME=etcd-${ETCD_VER}-linux-amd64.tar.gz

    GITHUB_URL=https://github.com/etcd-io/etcd/releases/download
    DOWNLOAD_URL=${GITHUB_URL}

    rm -rf ${ETCD_DIR} && mkdir -p ${ETCD_DIR}

    curl -L ${DOWNLOAD_URL}/${ETCD_VER}/${ETCD_FILENAME} -o ${ETCD_DIR}/${ETCD_FILENAME}
    tar xzvf ${ETCD_DIR}/${ETCD_FILENAME} -C ${ETCD_DIR} --strip-components=1
    rm -f ${ETCD_DIR}/${ETCD_FILENAME}

    ${ETCD_DIR}/etcd --version

    # start a local etcd server
    ${ETCD_DIR}/etcd &

#### 4. Install KAFKA
    
    dnf install java

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
    bin/zookeeper-server-start.sh -daemon config/zookeeper.properties

    export KAFKA_HEAP_OPTS="-Xmx1G -Xms1G"
    bin/kafka-server-start.sh -daemon config/server.properties

#### 5. Install MongoDB

    MONGODB_DIR=/tmp/mongodb
    MONGODB_FILENAME=mongodb-linux-x86_64-4.4.19.tgz

    rm -rf ${MONGODB_DIR} && mkdir -p ${MONGODB_DIR} && mkdir -p ${MONGODB_DIR}/data

    curl -L https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-rhel70-4.4.19.tgz -o ${MONGODB_DIR}/${MONGODB_FILENAME}

    cd ${MONGODB_DIR}
    tar -vxzf ${MONGODB_FILENAME}
    mv mongodb-linux-x86_64-rhel70-4.4.19 mongodb-linux-x86_64-4.4.19
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

#### 6. Install redis
    sudo dnf install -y http://rpms.famillecollet.com/enterprise/remi-release-7.rpm
    sudo dnf --enablerepo=remi list redis --showduplicates | sort -r
    sudo dnf  --enablerepo=remi install redis
    sudo systemctl restart redis