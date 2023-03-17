
#!/system/bin/sh

/system/bin/ifconfig dummy0:1 172.16.16.16 netmask 255.255.255.0
setenforce 0

touch /data/ax/change.log
