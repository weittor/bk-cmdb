/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except 
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and 
 * limitations under the License.
 */
 
package api

import (
	"configcenter/src/framework/core/output/module/inst"
	"configcenter/src/framework/core/output/module/model"
)

// HostIteratorWrapper the host iterator wrapper
type HostIteratorWrapper struct {
	host inst.Iterator
}

// Next next the business
func (cli *HostIteratorWrapper) Next() (*HostWrapper, error) {

	host, err := cli.host.Next()

	return &HostWrapper{host: host}, err

}

// ForEach the foreach function
func (cli *HostIteratorWrapper) ForEach(callback func(host *HostWrapper) error) error {

	return cli.host.ForEach(func(item inst.Inst) error {
		return callback(&HostWrapper{host: item})
	})
}

// HostWrapper the host wrapper
type HostWrapper struct {
	host inst.Inst
}

// GetModel get the model for the host
func (cli *HostWrapper) GetModel() model.Model {
	return cli.host.GetModel()
}

// SetValue set the key value
func (cli *HostWrapper) SetValue(key string, val interface{}) error {
	return cli.host.SetValue(key, val)
}

// Save save the data
func (cli *HostWrapper) Save() error {

	if err := cli.host.SetValue(fieldImportFrom, HostImportFromAPI); nil != err {
		return err
	}
	return cli.host.Save()
}

// SetBakOperator set the bak operator
func (cli *HostWrapper) SetBakOperator(bakOperator string) error {
	return cli.host.SetValue(fieldBakOperator, bakOperator)
}

// GetBakOperator get the bak operator
func (cli *HostWrapper) GetBakOperator() (string, error) {
	vals, err := cli.host.GetValues()
	if nil != err {
		return "", err
	}
	return vals.String(fieldBakOperator), nil
}

// SetOsBit set the os bit
func (cli *HostWrapper) SetOsBit(osbit string) error {
	return cli.host.SetValue(fieldOsBit, osbit)
}

// GetOsBit get the os bit
func (cli *HostWrapper) GetOsBit() (string, error) {
	vals, err := cli.host.GetValues()
	if nil != err {
		return "", err
	}
	return vals.String(fieldOsBit), nil
}

// SetSLA set the sla
func (cli *HostWrapper) SetSLA(sla string) error {
	return cli.host.SetValue(fieldSLA, sla)
}

// GetSLA get sla
func (cli *HostWrapper) GetSLA() (string, error) {
	vals, err := cli.host.GetValues()
	if nil != err {
		return "", err
	}
	return vals.String(fieldSLA), nil
}

// SetCloudID set the cloudid for the host
func (cli *HostWrapper) SetCloudID(cloudID int64) error {
	return cli.host.SetValue(fieldCloudID, cloudID)
}

// GetCloudID get the cloudid
func (cli *HostWrapper) GetCloudID() (int, error) {
	vals, err := cli.host.GetValues()
	if nil != err {
		return 0, err
	}
	return vals.Int(fieldCloudID)
}

// SetInnerIP set the inner ip
func (cli *HostWrapper) SetInnerIP(innerIP string) error {
	return cli.host.SetValue(fieldHostInnerIP, innerIP)
}

// GetInnerIP get the inner ip
func (cli *HostWrapper) GetInnerIP() (string, error) {
	vals, err := cli.host.GetValues()
	if nil != err {
		return "", err
	}
	return vals.String(fieldHostInnerIP), nil
}

// SetOperator set the operator for the host
func (cli *HostWrapper) SetOperator(operator string) error {
	return cli.host.SetValue(fieldHostOperator, operator)
}

// GetOperator get the operator
func (cli *HostWrapper) GetOperator() (string, error) {
	vals, err := cli.host.GetValues()
	if nil != err {
		return "", err
	}
	return vals.String(fieldHostOperator), nil
}

// SetCPU set the cpu core num  for the host
func (cli *HostWrapper) SetCPU(cpu int64) error {
	return cli.host.SetValue(fieldCPU, cpu)
}

// GetCPU get the cpu core num
func (cli *HostWrapper) GetCPU() (int, error) {
	vals, err := cli.host.GetValues()
	if nil != err {
		return 0, err
	}
	return vals.Int(fieldCPU)
}

// SetCPUMhz set the cpu mhz
func (cli *HostWrapper) SetCPUMhz(cpuMhz int64) error {
	return cli.host.SetValue(fieldCPUMhz, cpuMhz)
}

// GetCPUMhz get the cpu mhz
func (cli *HostWrapper) GetCPUMhz() (int, error) {
	vals, err := cli.host.GetValues()
	if nil != err {
		return 0, err
	}
	return vals.Int(fieldCPUMhz)
}

// SetOsType set the os type for the host
func (cli *HostWrapper) SetOsType(osType string) error {
	return cli.host.SetValue(fieldOsType, osType)
}

// GetOsType get the os type
func (cli *HostWrapper) GetOsType() (string, error) {
	vals, err := cli.host.GetValues()
	if nil != err {
		return "", err
	}
	return vals.String(fieldOsType), nil
}

// SetOuterIP set the outer ip for the host
func (cli *HostWrapper) SetOuterIP(outerIP string) error {
	return cli.host.SetValue(fieldHostOuterIP, outerIP)
}

// GetOuterIP get the outerip
func (cli *HostWrapper) GetOuterIP() (string, error) {
	vals, err := cli.host.GetValues()
	if nil != err {
		return "", err
	}
	return vals.String(fieldHostOuterIP), nil
}

// SetAssetID set the assetid for the host
func (cli *HostWrapper) SetAssetID(assetID string) error {
	return cli.host.SetValue(fieldAssetID, assetID)
}

// GetAssetID get the asset id for the host
func (cli *HostWrapper) GetAssetID() (string, error) {
	vals, err := cli.host.GetValues()
	if nil != err {
		return "", err
	}
	return vals.String(fieldAssetID), nil
}

// SetInnerMac set the mac for the host
func (cli *HostWrapper) SetInnerMac(mac string) error {
	return cli.host.SetValue(fieldMac, mac)
}

// GetInnerMac get the mac for the host
func (cli *HostWrapper) GetInnerMac() (string, error) {
	vals, err := cli.host.GetValues()
	if nil != err {
		return "", err
	}

	return vals.String(fieldMac), nil
}

// SetOuterMac set the mac for the host
func (cli *HostWrapper) SetOuterMac(mac string) error {
	return cli.host.SetValue(fieldMac, mac)
}

// GetOuterMac get the mac for the host
func (cli *HostWrapper) GetOuterMac() (string, error) {
	vals, err := cli.host.GetValues()
	if nil != err {
		return "", err
	}

	return vals.String(fieldMac), nil
}

// SetSN set the sn for the host
func (cli *HostWrapper) SetSN(sn string) error {
	return cli.host.SetValue(fieldSN, sn)
}

// GetSN get the sn for the host
func (cli *HostWrapper) GetSN() (string, error) {
	vals, err := cli.host.GetValues()
	if nil != err {
		return "", err
	}
	return vals.String(fieldSN), nil
}

// SetCPUModule set the cpu module for the host
func (cli *HostWrapper) SetCPUModule(cpuModule string) error {
	return cli.host.SetValue(fieldCPUModule, cpuModule)
}

// GetCPUModule get the cpu module
func (cli *HostWrapper) GetCPUModule() (string, error) {
	vals, err := cli.host.GetValues()
	if nil != err {
		return "", err
	}
	return vals.String(fieldCPUModule), nil
}

// SetName set the host name for the host
func (cli *HostWrapper) SetName(hostName string) error {
	return cli.host.SetValue(fieldHostName, hostName)
}

// GetName get the host name
func (cli *HostWrapper) GetName() (string, error) {
	vals, err := cli.host.GetValues()
	if nil != err {
		return "", err
	}
	return vals.String(fieldHostName), nil
}

// SetServiceTerm set the service term for the host
func (cli *HostWrapper) SetServiceTerm(serviceTerm int64) error {
	return cli.host.SetValue(fieldServiceTerm, serviceTerm)
}

// GetServiceTerm get the service term
func (cli *HostWrapper) GetServiceTerm() (int, error) {
	vals, err := cli.host.GetValues()
	if nil != err {
		return 0, err
	}
	return vals.Int(fieldServiceTerm)
}

// SetComment set the comment for the host
func (cli *HostWrapper) SetComment(comment string) error {
	return cli.host.SetValue(fieldComment, comment)
}

// GetComment get the comment for the host
func (cli *HostWrapper) GetComment() (string, error) {
	vals, err := cli.host.GetValues()
	if nil != err {
		return "", err
	}
	return vals.String(fieldComment), nil
}

// SetMem set the mem for the host
func (cli *HostWrapper) SetMem(mem int64) error {
	return cli.host.SetValue(fieldMem, mem)
}

// GetMem get the mem for the host
func (cli *HostWrapper) GetMem() (int, error) {
	vals, err := cli.host.GetValues()
	if nil != err {
		return 0, err
	}
	return vals.Int(fieldMem)
}

// SetDisk set the capacity of the disk for the host
func (cli *HostWrapper) SetDisk(disk int64) error {
	return cli.host.SetValue(fieldDisk, disk)
}

// GetDisk get the disk
func (cli *HostWrapper) GetDisk() (int, error) {
	vals, err := cli.host.GetValues()
	if nil != err {
		return 0, err
	}
	return vals.Int(fieldDisk)
}

// SetOsName set the os name for the host
func (cli *HostWrapper) SetOsName(osName string) error {
	return cli.host.SetValue(fieldOsName, osName)
}

// GetOsName get the os name
func (cli *HostWrapper) GetOsName() (string, error) {
	vals, err := cli.host.GetValues()
	if nil != err {
		return "", err
	}
	return vals.String(fieldOsName), nil
}

// SetOsVersion set the os version for the host
func (cli *HostWrapper) SetOsVersion(osVersion string) error {
	return cli.host.SetValue(fieldOsVersion, osVersion)
}

// GetOsVersion get the os version
func (cli *HostWrapper) GetOsVersion() (string, error) {
	vals, err := cli.host.GetValues()
	if nil != err {
		return "", err
	}
	return vals.String(fieldOsVersion), nil
}
