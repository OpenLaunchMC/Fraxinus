package server

import (
	"bufio"
	"os"

	jsoniter "github.com/json-iterator/go"
	
	ferrors "github.com/ProjectOpenLaunch/Fraxinus/errors"
)

func LoadExtensionsFile() ( map[string]interface{} , error ) { 
	f , err := os.OpenFile("extensions.json", os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	buf := make([]byte, 0, 1024 * 1024)
	conf := make(map[string]interface{})
	_ , err = r.Read(buf)
	if err != nil {
		return nil, err
	}
	err = jsoniter.ConfigFastest.Unmarshal(buf,conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func (s *Server)RegisterExtensions() []error {
	file , err := LoadExtensionsFile()
	list := file["extensions"].([]string)
	if err != nil {
		return err
	}
	for _, name := range list {
		ext = file[name].(interface{})
		if ext == nil {
			errs = append(errs, errors.New("Extension " + name + " not found in extensions.json"))
		s.a.Use()
	}
	return nil
}