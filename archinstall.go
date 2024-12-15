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
	Sudo     bool   `json:"sudo"`
	Username string `json:"username"`
}

type SysConfig struct {
	AdditionalRepositories []string         `json:"additional-repositories"`
	ArchInstallLanguage    string           `json:"archinstall-language"`
	AudioConfig            AudioSysConfig   `json:"audio_config"`
	Bootloader             string           `json:"bootloader"`
	ConfigVersion          string           `json:"config_version"`
	DiskConfig             DiskSysConfig    `json:"disk_config"`
	HostName               string           `json:"hostname"`
	Kernels                []string         `json:"kernels"`
	LocaleConfig           LocaleSysConfig  `json:"locale_config"`
	MirrorConfig           MirrorSysConfig  `json:"mirror_config"`
	NTP                    bool             `json:"ntp"`
	Packages               []string         `json:"packages"`
	ParallelDownloads      int              `json:"parallel_downloads"`
	ProfileConfig          ProfileSysConfig `json:"profile_config"`
	Swap                   bool             `json:"swap"`
	TimeZone               string           `json:"timezone"`
	UKI                    bool             `json:"uki"`
	Version                string           `json:"version"`
}

type MirrorSysConfig struct {
	CustomMirrors []CustomMirrorSysConfig `json:"custom_mirrors"`
}

type CustomMirrorSysConfig struct {
	Name       string `json:"name"`
	SignCheck  string `json:"sign_check"`
	SignOption string `json:"sign_option"`
	Url        string `json:"url"`
}

type LocaleSysConfig struct {
	KbLayout string `json:"kb_layout"`
	SysEnc   string `json:"sys_enc"`
	SysLang  string `json:"sys_lang"`
}
type AudioSysConfig struct {
	Audio string `json:"audio"`
}

type DiskSysConfig struct {
	ConfigType          string                `json:"config_type"`
	DeviceModifications []DeviceDiskSysConfig `json:"device_modifications"`
}

type DeviceDiskSysConfig struct {
	Device     string                         `json:"device"`
	Partitions []PartitionDeviceDiskSysConfig `json:"partitions"`
	Wipe       bool                           `json:"wipe"`
}

type PartitionDeviceDiskSysConfig struct {
	BTRFS        []string                            `json:"btrfs"`
	DevPath      *string                             `json:"dev_path"`
	Flags        []string                            `json:"flags"`
	FsType       string                              `json:"fs_type"`
	MountOptions []string                            `json:"mount_options"`
	MountPoint   string                              `json:"mountpoint"`
	ObjId        string                              `json:"obj_id"`
	Size         GenericPartitionDeviceDiskSysConfig `json:"size"`
	Start        GenericPartitionDeviceDiskSysConfig `json:"start"`
	Status       string                              `json:"status"`
	Type         string                              `json:"type"`
}

type GenericPartitionDeviceDiskSysConfig struct {
	SectorSize SectorGenericPartitionDeviceDiskSysConfig `json:"sector_size"`
	Unit       string                                    `json:"unit"`
	Value      int64                                     `json:"value"`
}

type SectorGenericPartitionDeviceDiskSysConfig struct {
	Unit  string `json:"unit"`
	Value int64  `json:"value"`
}

type ProfileSysConfig struct {
	GfxDriver string                  `json:"gfx_driver"`
	Profile   ProfileProfileSysConfig `json:"profile"`
}

type ProfileProfileSysConfig struct {
	Main string `json:"main"`
}

func populateConfigJson() error {
	configfile, err := os.ReadFile("/root/xemubox-archinstall-template/configuration_template.json")
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

	ConfData.DiskConfig.DeviceModifications[0].Device = "/dev/" + dsk

	ConfData.ProfileConfig.GfxDriver = gfx
	ConfData.DiskConfig.DeviceModifications[0].Partitions[0].DevPath = nil
	ConfData.DiskConfig.DeviceModifications[0].Partitions[1].DevPath = nil

	updatedConf, err := json.MarshalIndent(ConfData, "", "	")
	if err != nil {
		return err
	}
	err = os.WriteFile("/root/user_configuration.json", updatedConf, os.ModePerm)
	if err != nil {
		return err
	}

	credfile, err := os.ReadFile("/root/xemubox-archinstall-template/credential_template.json")
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

	err = os.WriteFile("/root/user_credentials.json", updatedCred, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
