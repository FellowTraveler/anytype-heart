package files

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
)

// GetFileConfig - returns data from config file, if file doesn't exist returns same cfg struct
func GetFileConfig(configPath string, cfg interface{}) error {
	if reflect.ValueOf(cfg).Kind() != reflect.Ptr {
		return fmt.Errorf("cfg param must be a pointer type")
	}

	cfgFile, err := os.OpenFile(configPath, os.O_RDONLY, 0655)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	defer cfgFile.Close()

	if !errors.Is(err, os.ErrNotExist) {
		info, err := cfgFile.Stat()
		if err != nil {
			return fmt.Errorf("fail to get info about config file: %s", err.Error())
		}

		if info.Size() > 0 {
			err = json.NewDecoder(cfgFile).Decode(cfg)
			if err != nil {
				return fmt.Errorf("invalid file config format: %s", err.Error())
			}
		}
	}

	return nil
}

//WriteFileConfig - overwrites params in file only specified params which passed in cfg
func WriteFileConfig(configPath string, cfg interface{}) error {
	oldCfg := make(map[string]interface{})
	if err := GetFileConfig(configPath, &oldCfg); err != nil {
		return err
	}

	newConfig, err := toMapInterface(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal new config: %s", err.Error())
	}

	for oldKey, oldData := range oldCfg {
		if _, ok := newConfig[oldKey]; !ok {
			newConfig[oldKey] = oldData
		}
	}

	cfgFile, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0640)
	if err != nil {
		return fmt.Errorf("failed to open cfg file for updating: %s", err.Error())
	}
	defer cfgFile.Close()

	err = json.NewEncoder(cfgFile).Encode(newConfig)
	if err != nil {
		return fmt.Errorf("failed to save data to the config file: %s", err.Error())
	}

	return nil
}

func toMapInterface(cfg interface{}) (map[string]interface{}, error) {
	var m map[string]interface{}
	byteData, err := json.Marshal(cfg)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteData, &m)
	return m, err
}
