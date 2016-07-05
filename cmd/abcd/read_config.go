package main

import (
	"encoding/json"
	//"fmt"
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
const DBNAME = ".abcd.bolt"

const HOME_FLAG = `HOME=`
func ReadConfig () {
	env := os.Environ()
	conf_dir := GetConfDir(&env)
	conf_file := GetConfFile(conf_dir)
	EnsureConfDir(conf_dir, fs.AppFs)
	EnsureConfFile(conf_file, fs.AppFs)

	// 3) Check the bolt db exists OR create
	// EnsureConfDB()
	// 4) Return configuration object - or pointer? erghhh
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

func MakeHome(home_path string){
	os.Mkdir(home_path, os.ModeDir)
}

// Just test stat'ing for now
func EnsureConfDir(p string, fs afero.Fs){
	_, err := fs.Stat(p)

	if err != nil {
		if os.IsNotExist(err) {
			err = fs.Mkdir(p, os.ModeDir)
			if err != nil {
				panic(err)
			}
		}
	}
}

func EnsureConfFile(p string, fss afero.Fs){
	_, err := fss.Stat(p)

	if err != nil {
		if os.IsNotExist(err) {
			WriteDefaultConf(p, fss)
		} else {
			panic(err)
		}
	}
	EnsureConfFileValid(p, fss)
}

func WriteDefaultConf(p string, fss afero.Fs){
	defaultConfig := []byte(`{"feeds":[],"faves":[]}`)
	eFC := afero.WriteFile(fss, p, defaultConfig, 0644)
	fs.Assert(eFC)
}
func EnsureConfFileValid(p string, ffs afero.Fs){
	var m interface{}
	conf, eRF := afero.ReadFile(ffs, p)
	fs.Assert(eRF)

	eSer := json.Unmarshal(conf, &m)

	if fs.IsNil(eSer) {
		WriteDefaultConf(p, ffs)
	}
}
