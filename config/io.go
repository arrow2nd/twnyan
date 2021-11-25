package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// getConfigDir 設定ファイルのディレクトリを取得
func getConfigDir() string {
	dirPath, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return filepath.Join(dirPath, ".twnyan")
}

// configFileExists 設定ファイルが存在するか
func (cfg *Config) configFileExists() bool {
	dir := cfg.Option.ConfigDir
	fileList := []string{
		filepath.Join(dir, crdFile),
		filepath.Join(dir, optFile),
		filepath.Join(dir, colFile),
	}

	// ディレクトリの存在チェック
	if _, err := os.Stat(dir); err != nil {
		if err := os.Mkdir(dir, 0777); err != nil {
			panic(err)
		}
		return false
	}

	// ファイルの存在チェック
	for _, path := range fileList {
		if _, err := os.Stat(path); err != nil {
			return false
		}
	}

	return true
}

// saveYAML ファイルに保存
func (cfg *Config) saveYaml(filename string, in interface{}) {
	buf, err := yaml.Marshal(in)
	if err != nil {
		panic(err)
	}

	path := filepath.Join(cfg.Option.ConfigDir, filename)
	err = ioutil.WriteFile(path, buf, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

// loadYAML ファイルから読込
func (cfg *Config) loadYaml(filename string, out interface{}) {
	path := filepath.Join(cfg.Option.ConfigDir, filename)
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(buf, out)
	if err != nil {
		panic(err)
	}
}
