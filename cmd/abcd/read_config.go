package main

import (
	"os"
	"strings"
	"regexp"
	"path/filepath"
	// "time"
	// log "github.com/Sirupsen/logrus"
	// "github.com/olebedev/config"
	// "strconv"
	"github.com/spf13/afero"
)

var FSp afero.Fs = afero.NewOsFs()

const DIRNAME = ".abcd"
const CONFNAME = ".abcd.json"
const DBNAME = ".abcd.bolt"

const HOME_FLAG = `HOME=`
func ReadConfig () {
	env := os.Environ()
	conf_dir := GetConfDir(&env)
	EnsureConfDir(&conf_dir)
	// fi, err := os.Lstat(conf_loc)

	// 1) Check the conf dir exists OR create
	// 2) Check the conf *file* exists OR create
	// 3) Check the bolt db exists OR create
	// EnsureConfDir()
	// EnsureConfFile()
	// EnsureConfDB()

	// try_it = true
	// while try_it {
	//	fi, not_exists := os.Lstat(conf_loc)
	//	if not_exists {
	//		mkdir, perm_error = Os.Mkdir(conf_loc, os.ModeDir)
	//		if perm_error {
	//			try_it = false
	//			panic(perm_error)
	//		} else {

	//		}
	//	}
	// }

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

func MakeHome(home_path *string){
	os.Mkdir(*home_path, os.ModeDir)
}

func EnsureConfDir(p *string){
}
