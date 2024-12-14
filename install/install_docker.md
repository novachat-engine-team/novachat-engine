
### Deploy by Docker

Download file [docker-compose.yaml](https://github.com/novachat-engine-team/novachat-engine/blob/main/install/docker-compose.yaml)
    
    curl https://github.com/novachat-engine-team/novachat-engine/blob/main/install/docker-compose.yaml -o docker-compose.yaml

#### CentOS 7
    2024-6-30 EOL, End Of Life

    sudo curl -o /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo
    sudo yum clean all
    sudo yum makecache

    sudo yum install epel-release -y
    sudo yum install dnf -y

    sudo curl -o /etc/yum.repos.d/docker-ce.repo https://download.docker.com/linux/centos/docker-ce.repo

    sudo dnf install docker-ce
    sudo dnf install docker-compose

    sudo systemctl start docker

    sudo docker-compose up -d

    sudo docker ps

    eg:
    CONTAINER ID   IMAGE                         COMMAND                  CREATED         STATUS         PORTS                                                                   NAMES
    bd90fb76ee2f   wurstmeister/kafka            "start-kafka.sh"         6 seconds ago   Up 4 seconds   0.0.0.0:9092->9092/tcp, :::9092->9092/tcp                               kafka
    451f6783c3a1   quay.io/coreos/etcd:v3.5.14   "/usr/local/bin/etcd"    6 seconds ago   Up 5 seconds   0.0.0.0:2379-2380->2379-2380/tcp, :::2379-2380->2379-2380/tcp           etcd
    65b41523e72e   wurstmeister/zookeeper        "/bin/sh -c '/usr/sb…"   6 seconds ago   Up 5 seconds   22/tcp, 2888/tcp, 3888/tcp, 0.0.0.0:2181->2181/tcp, :::2181->2181/tcp   zookeeper
    32be50ea2c5b   ninglei1996/redisbloom        "docker-entrypoint.s…"   6 seconds ago   Up 5 seconds   0.0.0.0:6379->6379/tcp, :::6379->6379/tcp                               redis
    6042a41c7137   mysql:5.7                     "docker-entrypoint.s…"   6 seconds ago   Up 5 seconds   0.0.0.0:3306->3306/tcp, :::3306->3306/tcp, 33060/tcp                    mysql
    ef43de1bcb80   mongo:4.4.19                  "docker-entrypoint.s…"   6 seconds ago   Up 5 seconds   0.0.0.0:27017->27017/tcp, :::27017->27017/tcp                           mongodb

#### Ubuntu

    sudo apt install apt-transport-https ca-certificates curl software-properties-common -y

    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
    echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

    sudo apt install -y docker-ce

    sudo apt install docker-compose -y

    sudo systemctl start docker

    sudo docker-compose up -d

    sudo docker ps

    eg:
        CONTAINER ID   IMAGE                         COMMAND                  CREATED          STATUS          PORTS                                                                   NAMES
        f3646545ff1c   wurstmeister/kafka            "start-kafka.sh"         46 seconds ago   Up 45 seconds   0.0.0.0:9092->9092/tcp, :::9092->9092/tcp                               kafka
        adff41659c90   wurstmeister/zookeeper        "/bin/sh -c '/usr/sb…"   49 seconds ago   Up 45 seconds   22/tcp, 2888/tcp, 3888/tcp, 0.0.0.0:2181->2181/tcp, :::2181->2181/tcp   zookeeper
        05b53b0def3e   mysql:5.7                     "docker-entrypoint.s…"   49 seconds ago   Up 45 seconds   0.0.0.0:3306->3306/tcp, :::3306->3306/tcp, 33060/tcp                    mysql
        f1d1384a694a   mongo:4.4.19                  "docker-entrypoint.s…"   49 seconds ago   Up 45 seconds   0.0.0.0:27017->27017/tcp, :::27017->27017/tcp                           mongodb
        ccd47e2b526a   quay.io/coreos/etcd:v3.5.14   "/usr/local/bin/etcd"    49 seconds ago   Up 45 seconds   0.0.0.0:2379-2380->2379-2380/tcp, :::2379-2380->2379-2380/tcp           etcd
        3a472b0bbeae   ninglei1996/redisbloom        "docker-entrypoint.s…"   49 seconds ago   Up 45 seconds   0.0.0.0:6379->6379/tcp, :::6379->6379/tcp                               redis
