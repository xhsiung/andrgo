#!/bin/bash

adb shell rm -rf  /sdcard/upload/$1
adb push $1 /sdcard/upload/

