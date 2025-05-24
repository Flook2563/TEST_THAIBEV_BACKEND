package config

import (
	"log"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

func LoadFileConfig[T any](configPath string) *T {
	if len(configPath) == 0 {
		log.Fatal("Config path is required")
	}

	dir := filepath.Dir(configPath)
	filebase := filepath.Base(configPath)
	// file name without extension
	filename := strings.TrimSuffix(filebase, filepath.Ext(filebase))

	viper.SetConfigName(filename)
	viper.AddConfigPath(dir)
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	// convert _ to dot in env variable
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("fatal error config file: %+v", err)
	}

	var cfg T
	bindEnvs("env", cfg)
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("unable to decode config into struct, %+v", err)
	}

	return &cfg
}

func bindEnvs(tag string, iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup(tag)
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			bindEnvs(tag, v.Interface(), append(parts, tv)...)
		default:
			viper.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}
