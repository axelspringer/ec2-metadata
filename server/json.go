// Copyright © 2017 Axel Springer SE
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package server

// EC2MetaData wraps all the ec2 metadata
type EC2MetaData struct {
	DynamicData DynamicData `json:"dynamicData"`
	UserData    UserData    `json:"userData"`
	MetaData    MetaData    `json:"metaData"`
	Versions    []string    `json:"versions"`
}

// InstanceIdentity contains the instance-identity endpoints
type InstanceIdentity struct {
	Document InstanceIdentityDocument `json:"document"`
}

// InstanceIdentityDocument contains the ec2 instance identity document
type InstanceIdentityDocument struct {
	DevpayProductCodes string `json:"devpayProductCodes"`
	AvailabilityZone   string `json:"availabilityZone"`
	PrivateIP          string `json:"privateIp"`
	Version            string `json:"version"`
	InstanceID         string `json:"instanceId"`
	InstanceType       string `json:"instanceType"`
	AccountID          string `json:"accountId"`
	ImageID            string `json:"imageId"`
	PendingTime        string `json:"pendingTime"`
	Architecture       string `json:"architecture"`
	KernelID           string `json:"kernelId"`
	RAMDiskID          string `json:"ramdiskId"`
	Region             string `json:"region"`
}

// DynamicData contains all the dynamic data
type DynamicData struct {
	InstanceIdentity InstanceIdentity `json:"instanceIdentity"`
}

// UserData contains the user data associated with the ec2 instance
type UserData struct {
	EnvironmentCode           string `json:"environment_code"`
	EnvironmentName           string `json:"environment_name"`
	AppCode                   string `json:"app_code"`
	AppName                   string `json:"app_name"`
	AppEmail                  string `json:"app_email"`
	AppDescription            string `json:"app_description"`
	ClusterCode               string `json:"cluster_code"`
	ClusterName               string `json:"cluster_name"`
	ClusterStack              string `json:"cluster_stack"`
	InstanceConfigurationName string `json:"instance_configuration_name"`
	HostnameSuffix            string `json:"hostname_suffix"`
	VolumeSizeGb              int    `json:"volume_sizes_gb"`
	VolumeDevices             string `json:"volume_devices"`
	VolumeMountPoints         string `json:"volume_mount_points"`
	VolumeSnapShotFrequencies int    `json:"volume_snapshot_frequencies"`
	SnapshotDaysRetained      int    `json:"snapshot_days_retained"`
	RegionCode                string `json:"region_code"`
	Hostname                  string `json:"hostname"`
	VpcName                   string `json:"vpc_name"`
	SubnetName                string `json:"subnet_name"`
	NimbulKey                 string `json:"nimbul_key"`
	BuildTag                  string `json:"build_tag"`
}

// MetaData contains the ec2 metadata
type MetaData struct {
	AmiID               string `json:"ami-id"`
	AmiLaunchIndex      int    `json:"ami-launch-index"`
	AmiManifestPath     string `json:"ami-manifest-path"`
	LocalIpv4           string `json:"local-ipv4"`
	AvailabilityZone    string `json:"availability-zone"`
	Hostname            string `json:"hostname"`
	InstanceAction      string `json:"instance-action"`
	InstanceID          string `json:"instance-id"`
	InstanceType        string `json:"instance-type"`
	Mac                 string `json:"mac"`
	Profile             string `json:"profile"`
	ReservationID       string `json:"reservation-id"`
	SecurityCredentials SecurityCredentials
	SecurityGroups      []string `json:"security-groups"`
}

// SecurityCredentials contains the security credentials associate with the instance
type SecurityCredentials struct {
	User            string `json:"user"`
	AccessKeyID     string `json:"AccessKeyId"`
	SecretAccessKey string `json:"SecretAccessKey"`
	Token           string `json:"Token"`
	Expiration      string `json:"Expiration"`
}
