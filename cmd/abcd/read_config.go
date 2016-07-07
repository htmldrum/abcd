package main

import (
	"encoding/json"
	//"io/ioutil"
	"os"
	"strings"
	"regexp"
	"path/filepath"
	// "time"
	// log "github.com/Sirupsen/logrus"
	// "github.com/olebedev/config"
	// "strconv"
	"github.com/htmldrum/abcd/fs"
	"github.com/spf13/afero"
)

const DIRNAME = ".abcd"
const CONFNAME = ".abcd.json"
// const DBNAME = ".abcd.db"

const HOME_FLAG = `HOME=`
func ReadConfig ()(Config){
	env := os.Environ()
	conf_dir := GetConfDir(&env)
	conf_file := filepath.Join(conf_dir, CONFNAME)
	// db_file := filepath.Join(conf_dir, DBNAME)
	EnsureConfDir(conf_dir, fs.AppFs)
	return EnsureConfFile(conf_file, fs.AppFs)
}

func GetConfDir(envs *[]string)(p string){
	var home string
	for _, env := range *envs {
		if strings.Index(env,"HOME=") != -1{
			re := regexp.MustCompile("HOME=(?P<dir>.*)")
			md := re.FindStringSubmatch(env)
			if len(md) == 2 {
				home = md[1]
			}
		}
	}

	if len(home) == 0 {
		panic("You have no home :( - I will not work for you")
	}

	p = filepath.Join(home, DIRNAME)
	return
}
func GetConfFile(p string) string {
	return filepath.Join(p, CONFNAME)
}

// Just test stat'ing for now
func EnsureConfDir(p string, fss afero.Fs){
	_, err := fss.Stat(p)

	if err != nil {
		if os.IsNotExist(err) {
			err = fss.MkdirAll(p, 0700)
			if err != nil {
				panic(err)
			}
		}
	}
}

func EnsureConfFile(p string, fss afero.Fs)(Config){
	_, err := fss.Stat(p)

	if err != nil {
		if os.IsNotExist(err) {
			WriteDefaultConf(p, fss)
		} else {
			panic(err)
		}
	}
	return EnsureConfFileValid(p, fss)
}

func WriteDefaultConf(p string, fss afero.Fs)(c Config){
	defaultConfig := []byte(`{"feeds":[],"faves":[]}`)
	eFC := afero.WriteFile(fss, p, defaultConfig, 0777)
	fs.Assert(eFC)
	json.Unmarshal(defaultConfig, &c)
	return c
}
func EnsureConfFileValid(p string, ffs afero.Fs)(c Config){
	conf, eRF := afero.ReadFile(ffs, p)
	fs.Assert(eRF)

	eSer := json.Unmarshal(conf, &c)

	if fs.IsNil(eSer) {
		c = WriteDefaultConf(p, ffs)
	}
	return c
}
