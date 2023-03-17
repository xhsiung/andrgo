#!/system/bin/sh

init() {
    mkdir -p /data/ax/bin
    mkdir -p /data/ax/etc    
    mkdir -p /data/ax/www  
    mkdir -p /data/ax/src
    echo "init"
}

delfn(){
    rm -rf /etc/init/axinit01.rc
    rm -rf /data/ax/bin/init.sh    
    echo "delfn"
}

cpfn(){
    cp axinit01.rc  /etc/init/
    cp init.sh      /data/ax/bin/
    echo "cpfn"
}

authfn() {
    chmod +x  /data/ax/bin/init.sh
    echo "authfn"
}

init 
delfn
cpfn 
authfn 

echo "done"
