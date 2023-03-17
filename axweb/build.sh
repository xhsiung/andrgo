#!/bin/bash

CWD=`pwd`
rm -rf myweb/dist

cd myweb/src 
#rm -rf my.min.js
#uglifyjs my.js -o my.min.js -m -c
npm run build
cd $CWD

#cd www/static/js/init 
#rm -rf login.min.js
#uglifyjs login.js -o login.min.js -m -c
#cd $CWD

rm -rf www/*
cp -a  myweb/dist/*  www/

#go run main.go
