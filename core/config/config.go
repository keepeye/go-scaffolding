package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cast"
	"github.com/tomwright/dasel"
)

var cfgtree *dasel.Node

func init() {
	var v interface{}
	if _, err := toml.DecodeFile(filepath.Join(os.Getenv("WORKDIR"), "config.toml"), &v); err != nil {
		panic(err)
	}
	cfgtree = dasel.New(v)
}

func GetString(path string, a ...any) string {
	if len(a) > 0 {
		path = fmt.Sprintf(path, a...)
	}
	v, err := cfgtree.Query(path)
	if err != nil {
		return ""
	}
	return v.String()
}

func GetInt(path string, a ...any) int {
	if len(a) > 0 {
		path = fmt.Sprintf(path, a...)
	}
	v, err := cfgtree.Query(path)
	if err != nil {
		return 0
	}
	return cast.ToInt(v.InterfaceValue())
}

func GetBool(path string, a ...any) bool {
	if len(a) > 0 {
		path = fmt.Sprintf(path, a...)
	}
	v, err := cfgtree.Query(path)
	if err != nil {
		return false
	}
	return cast.ToBool(v.InterfaceValue())
}
