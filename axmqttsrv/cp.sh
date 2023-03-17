#!/system/bin/sh

init() {
    mkdir -p /data/ax/bin
    mkdir -p /data/ax/etc/axmqttsrv
    mkdir -p /data/ax/src
    mkdir -p /data/ax/www  
    echo "init"
}

delfn(){
    rm -rf /etc/init/axmqttsrv00.rc  
    rm -rf /data/ax/etc/axmqttsrv/conf
    echo "delfn"
}


cpfn(){
    cp axmqttsrv00.rc  /etc/init/
    cp axmqttsrv      /data/ax/bin/        
    cp -a conf       /data/ax/etc/axmqttsrv/
    echo "cpfn"
}

authfn() {
    chmod +x  /data/ax/bin/axmqttsrv
    echo "authfn"
}


init 
delfn
cpfn 
authfn 

echo "done"