#!/system/bin/sh

init() {
    mkdir -p /data/ax/bin    
    mkdir -p /data/ax/etc/axweb
    mkdir -p /data/ax/www  
    mkdir -p /data/ax/src
    echo "init"
}

delfn(){
    rm -rf /etc/init/axweb00.rc
    rm -rf /data/ax/www  
    rm -rf /data/ax/etc/axweb/conf
    echo "delfn"
}


cpfn(){
    cp axweb00.rc  /etc/init/
    cp gowebd      /data/ax/bin/    
    cp -a www      /data/ax/
    cp -a conf     /data/ax/etc/axweb/
    echo "cpfn"
}

authfn() {
    chmod +x  /data/ax/bin/gowebd
    echo "authfn"
}


init 
delfn
cpfn 
authfn 

echo "done"