#!/bin/bash
CWD=`pwd`
mkdir -p www

cd myweb
#npm install 

rm -rf dist
npm run build
cd $CWD

rm -rf www/*
cp -a  myweb/dist/*  www/

#go run main.go
