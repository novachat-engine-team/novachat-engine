### ETCD url
    https://etcd.io/docs/v3.5/install/

### Install
    centos:
        yum install etcd -y

    ubuntu:
        apt install etcd -y

### Config
    
    ETCD_LISTEN_CLIENT_URLS="http://0.0.0.0:2379"
    ETCD_ADVERTISE_CLIENT_URLS="http://0.0.0.0:2379"