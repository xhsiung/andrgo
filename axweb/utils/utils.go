package utils

import (
	"fmt"

	ini "gopkg.in/ini.v1"
)

const fileconf = "conf/conf.ini"

var (
	Cfg *ini.File
	err error
)

func init() {
	Cfg, err = ini.Load(fileconf)
	checkErr(err)
}

func ReloadCfg() {
	Cfg, err = ini.Load(fileconf)
}

func WriteCfg(head, key, val string) {
	Cfg.Section(head).Key(key).SetValue(val)
	Cfg.SaveTo(fileconf)
	ReloadCfg()
}

func WriteCfgHead(head string, kv map[string]interface{}) {
	//Cfg.Section(head).Key(key).SetValue(val)
	for k, v := range kv {
		//rv := reflect.ValueOf(v)
		vstr := fmt.Sprintf("%v", v)
		Cfg.Section(head).Key(k).SetValue(vstr)
	}

	Cfg.SaveTo(fileconf)
	ReloadCfg()
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
