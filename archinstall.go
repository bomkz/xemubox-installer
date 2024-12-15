package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type CredConfig struct {
	RootPassword string           `json:"!root-password"`
	Users        []UserCredConfig `json:"!users"`
}

type UserCredConfig struct {
	Password string `json:"!password"`
}

type SysConfig struct {
	DiskConfig    DiskSysConfig    `json:"disk_config"`
	ProfileConfig ProfileSysConfig `json:"profile_config"`
}

type DiskSysConfig struct {
	DeviceModifications []DeviceDiskSysConfig `json:"device_modifications"`
}

type DeviceDiskSysConfig struct {
	Device string `json:"device"`
}

type ProfileSysConfig struct {
	GfxDriver string `json:"gfx_driver"`
}

func populateConfigJson() error {
	configfile, err := os.ReadFile("configuration_template.json")
	if err != nil {
		return err
	}

	var ConfData SysConfig

	if err = json.Unmarshal(configfile, &ConfData); err != nil {
		return err
	}

	if len(ConfData.DiskConfig.DeviceModifications) == 0 {
		return fmt.Errorf("Error Reading XemuBOX Configuration Template: Unexpected or no value in 'device_modifications")
	}

	ConfData.DiskConfig.DeviceModifications[0].Device = dsk

	ConfData.ProfileConfig.GfxDriver = gfx

	updatedConf, err := json.MarshalIndent(ConfData, "", "	")
	if err != nil {
		return err
	}
	err = os.WriteFile("user_configuration.json", updatedConf, os.ModePerm)
	if err != nil {
		return err
	}

	credfile, err := os.ReadFile("credential_template.json")
	if err != nil {
		return err
	}

	var CredData CredConfig

	if err = json.Unmarshal(credfile, &CredData); err != nil {
		return err
	}

	if len(CredData.Users) == 0 {
		return fmt.Errorf("Error Reading XemuBOX Configuration Template: Unexpected or no value in !users'")
	}

	CredData.RootPassword = rtp
	CredData.Users[0].Password = usr

	updatedCred, err := json.MarshalIndent(CredData, "", "	")
	if err != nil {
		return err
	}

	err = os.WriteFile("user_credentials.json", updatedCred, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
