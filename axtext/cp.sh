#!/system/bin/sh

init() {
    mkdir -p /data/ax/bin
    mkdir -p /data/ax/etc    
    mkdir -p /data/ax/www  
    mkdir -p /data/ax/src
    echo "init"
}

delfn(){
    rm -rf /etc/init/axtest00.rc
    rm -rf /data/ax/bin/gotest
    echo "delfn"
}

cpfn(){
    cp axtest00.rc  /etc/init/
    cp gotest      /data/ax/bin/
    echo "cpfn"
}

authfn() {
    chmod +x  /data/ax/bin/gotest
    echo "authfn"
}

init 
delfn
cpfn 
authfn 

echo "done"
