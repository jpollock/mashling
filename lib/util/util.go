/*
* Copyright © 2017. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/TIBCOSoftware/flogo-cli/env"
	ftrigger "github.com/TIBCOSoftware/flogo-lib/core/trigger"
)

func GetGithubResource(gitHubPath string, resourceFile string) ([]byte, error) {
	gbProject := env.NewGbProjectEnv()

	gbProject.Init(os.Getenv("GOPATH"))

	resourceDir := gbProject.GetVendorSrcDir()
	resourcePath := resourceDir + "/" + gitHubPath + "/" + resourceFile

	gbProject.InstallDependency(gitHubPath, "")

	data, err := ioutil.ReadFile(resourcePath)
	if err != nil {
		return nil, err
	}

	err = gbProject.UninstallDependency(gitHubPath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetTriggerMetadata(gitHubPath string) (*ftrigger.Metadata, error) {
	gbProject := env.NewGbProjectEnv()

	gbProject.Init(os.Getenv("GOPATH"))

	resourceDir := gbProject.GetVendorSrcDir()
	triggerPath := resourceDir + "/" + gitHubPath + "/" + Gateway_Trigger_Metadata_JSON_Name
	gbProject.InstallDependency(gitHubPath, "")
	data, err := ioutil.ReadFile(triggerPath)
	if err != nil {
		// This hack only works for gb right now
		gb, ok := gbProject.(*env.GbProject)
		if !ok {
			return nil, errors.New("this is not a GB project")
		}
		_, triggerErr := os.Stat(filepath.Join(gb.VendorSrcDir, triggerPath))
		_, pathErr := os.Stat(filepath.Join(gb.VendorSrcDir, gitHubPath))
		if !os.IsNotExist(pathErr) && os.IsNotExist(triggerErr) {
			fmt.Fprintf(os.Stdout, "Nested trigger dependency was already loaded, manually loading higher level dependency...\n")
			err = tweakedFlogoCliInstallDependency(gb, gitHubPath, "")
			if err != nil {
				return nil, err
			}
			data, err = ioutil.ReadFile(triggerPath)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	triggerMetadata := &ftrigger.Metadata{}
	json.Unmarshal(data, triggerMetadata)
	return triggerMetadata, nil
}

func IsValidTriggerSetting(metadata *ftrigger.Metadata, property string) bool {
	settings := metadata.Settings
	for key := range settings {
		if key == property {
			return true
		}
	}

	return false
}

func IsValidTriggerHandlerSetting(metadata *ftrigger.Metadata, property string) bool {
	settings := metadata.Handler.Settings

	for _, element := range settings {
		if element.Name == property {
			return true
		}
	}

	return false
}

func ValidateTriggerConfigExpr(expression *string) (bool, *string) {
	if expression == nil {
		return false, nil
	}

	exprValue := *expression
	if strings.HasPrefix(exprValue, Gateway_Trigger_Config_Prefix) && strings.HasSuffix(exprValue, Gateway_Trigger_Config_Suffix) {
		//get name of the config
		str := exprValue[len(Gateway_Trigger_Config_Prefix) : len(exprValue)-1]
		return true, &str
	} else {
		return false, &exprValue
	}
}

func CheckTriggerOptimization(triggerSettings map[string]interface{}) bool {
	if val, ok := triggerSettings[Gateway_Trigger_Optimize_Property]; ok {
		optimize, err := strconv.ParseBool(val.(string))
		if err != nil {
			//check if its a boolean
			optimize, found := val.(bool)
			if !found {
				return found
			}
			return optimize
		}
		return optimize
	} else {
		return Gateway_Trigger_Optimize_Property_Default
	}
}

func ValidateEnvPropertySettingExpr(expression *string) (bool, *string) {
	if expression == nil {
		return false, nil
	}

	exprValue := *expression
	if strings.HasPrefix(exprValue, Gateway_Trigger_Setting_Env_Prefix) && strings.HasSuffix(exprValue, Gateway_Trigger_Setting_Env_Suffix) {
		//get name of the property
		str := exprValue[len(Gateway_Trigger_Setting_Env_Prefix) : len(exprValue)-1]
		return true, &str
	} else {
		return false, &exprValue
	}
}

// This is only called when a nested dependency has been loaded before the higher level one.
func tweakedFlogoCliInstallDependency(e *env.GbProject, depPath string, version string) error {
	var cmd *exec.Cmd

	cwd, _ := os.Getwd()
	defer os.Chdir(cwd)

	if version == "" {
		cmd = exec.Command("gb", "vendor", "fetch", depPath)
	} else {
		var tag string

		if version[0] != 'v' {
			tag = "v" + version
		} else {
			tag = version
		}

		cmd = exec.Command("gb", "vendor", "fetch", "-tag", tag, depPath)
	}

	os.Chdir(e.RootDir)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
