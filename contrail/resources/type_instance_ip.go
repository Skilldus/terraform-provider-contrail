//
// Automatically generated. DO NOT EDIT.
//

package resources

import (
	"encoding/json"
	"math/big"

	"github.com/Skilldus/contrail-go-api"
)

const (
	instance_ip_instance_ip_address int = iota
	instance_ip_instance_ip_family
	instance_ip_instance_ip_mode
	instance_ip_secondary_ip_tracking_ip
	instance_ip_subnet_uuid
	instance_ip_instance_ip_secondary
	instance_ip_instance_ip_local_ip
	instance_ip_service_instance_ip
	instance_ip_service_health_check_ip
	instance_ip_id_perms
	instance_ip_perms2
	instance_ip_annotations
	instance_ip_display_name
	instance_ip_virtual_network_refs
	instance_ip_network_ipam_refs
	instance_ip_virtual_machine_interface_refs
	instance_ip_physical_router_refs
	instance_ip_virtual_router_refs
	instance_ip_floating_ips
	instance_ip_tag_refs
	instance_ip_service_instance_back_refs
)

type InstanceIp struct {
	contrail.ObjectBase
	instance_ip_address            string
	instance_ip_family             string
	instance_ip_mode               string
	secondary_ip_tracking_ip       SubnetType
	subnet_uuid                    string
	instance_ip_secondary          bool
	instance_ip_local_ip           bool
	service_instance_ip            bool
	service_health_check_ip        bool
	id_perms                       IdPermsType
	perms2                         PermType2
	annotations                    KeyValuePairs
	display_name                   string
	virtual_network_refs           contrail.ReferenceList
	network_ipam_refs              contrail.ReferenceList
	virtual_machine_interface_refs contrail.ReferenceList
	physical_router_refs           contrail.ReferenceList
	virtual_router_refs            contrail.ReferenceList
	floating_ips                   contrail.ReferenceList
	tag_refs                       contrail.ReferenceList
	service_instance_back_refs     contrail.ReferenceList
	valid                          big.Int
	modified                       big.Int
	baseMap                        map[string]contrail.ReferenceList
}

func (obj *InstanceIp) GetType() string {
	return "instance-ip"
}

func (obj *InstanceIp) GetDefaultParent() []string {
	name := []string{}
	return name
}

func (obj *InstanceIp) GetDefaultParentType() string {
	return ""
}

func (obj *InstanceIp) SetName(name string) {
	obj.VSetName(obj, name)
}

func (obj *InstanceIp) SetParent(parent contrail.IObject) {
	obj.VSetParent(obj, parent)
}

func (obj *InstanceIp) storeReferenceBase(
	name string, refList contrail.ReferenceList) {
	if obj.baseMap == nil {
		obj.baseMap = make(map[string]contrail.ReferenceList)
	}
	refCopy := make(contrail.ReferenceList, len(refList))
	copy(refCopy, refList)
	obj.baseMap[name] = refCopy
}

func (obj *InstanceIp) hasReferenceBase(name string) bool {
	if obj.baseMap == nil {
		return false
	}
	_, exists := obj.baseMap[name]
	return exists
}

func (obj *InstanceIp) UpdateDone() {
	obj.modified.SetInt64(0)
	obj.baseMap = nil
}

func (obj *InstanceIp) GetInstanceIpAddress() string {
	return obj.instance_ip_address
}

func (obj *InstanceIp) SetInstanceIpAddress(value string) {
	obj.instance_ip_address = value
	obj.modified.SetBit(&obj.modified, instance_ip_instance_ip_address, 1)
}

func (obj *InstanceIp) GetInstanceIpFamily() string {
	return obj.instance_ip_family
}

func (obj *InstanceIp) SetInstanceIpFamily(value string) {
	obj.instance_ip_family = value
	obj.modified.SetBit(&obj.modified, instance_ip_instance_ip_family, 1)
}

func (obj *InstanceIp) GetInstanceIpMode() string {
	return obj.instance_ip_mode
}

func (obj *InstanceIp) SetInstanceIpMode(value string) {
	obj.instance_ip_mode = value
	obj.modified.SetBit(&obj.modified, instance_ip_instance_ip_mode, 1)
}

func (obj *InstanceIp) GetSecondaryIpTrackingIp() SubnetType {
	return obj.secondary_ip_tracking_ip
}

func (obj *InstanceIp) SetSecondaryIpTrackingIp(value *SubnetType) {
	obj.secondary_ip_tracking_ip = *value
	obj.modified.SetBit(&obj.modified, instance_ip_secondary_ip_tracking_ip, 1)
}

func (obj *InstanceIp) GetSubnetUuid() string {
	return obj.subnet_uuid
}

func (obj *InstanceIp) SetSubnetUuid(value string) {
	obj.subnet_uuid = value
	obj.modified.SetBit(&obj.modified, instance_ip_subnet_uuid, 1)
}

func (obj *InstanceIp) GetInstanceIpSecondary() bool {
	return obj.instance_ip_secondary
}

func (obj *InstanceIp) SetInstanceIpSecondary(value bool) {
	obj.instance_ip_secondary = value
	obj.modified.SetBit(&obj.modified, instance_ip_instance_ip_secondary, 1)
}

func (obj *InstanceIp) GetInstanceIpLocalIp() bool {
	return obj.instance_ip_local_ip
}

func (obj *InstanceIp) SetInstanceIpLocalIp(value bool) {
	obj.instance_ip_local_ip = value
	obj.modified.SetBit(&obj.modified, instance_ip_instance_ip_local_ip, 1)
}

func (obj *InstanceIp) GetServiceInstanceIp() bool {
	return obj.service_instance_ip
}

func (obj *InstanceIp) SetServiceInstanceIp(value bool) {
	obj.service_instance_ip = value
	obj.modified.SetBit(&obj.modified, instance_ip_service_instance_ip, 1)
}

func (obj *InstanceIp) GetServiceHealthCheckIp() bool {
	return obj.service_health_check_ip
}

func (obj *InstanceIp) SetServiceHealthCheckIp(value bool) {
	obj.service_health_check_ip = value
	obj.modified.SetBit(&obj.modified, instance_ip_service_health_check_ip, 1)
}

func (obj *InstanceIp) GetIdPerms() IdPermsType {
	return obj.id_perms
}

func (obj *InstanceIp) SetIdPerms(value *IdPermsType) {
	obj.id_perms = *value
	obj.modified.SetBit(&obj.modified, instance_ip_id_perms, 1)
}

func (obj *InstanceIp) GetPerms2() PermType2 {
	return obj.perms2
}

func (obj *InstanceIp) SetPerms2(value *PermType2) {
	obj.perms2 = *value
	obj.modified.SetBit(&obj.modified, instance_ip_perms2, 1)
}

func (obj *InstanceIp) GetAnnotations() KeyValuePairs {
	return obj.annotations
}

func (obj *InstanceIp) SetAnnotations(value *KeyValuePairs) {
	obj.annotations = *value
	obj.modified.SetBit(&obj.modified, instance_ip_annotations, 1)
}

func (obj *InstanceIp) GetDisplayName() string {
	return obj.display_name
}

func (obj *InstanceIp) SetDisplayName(value string) {
	obj.display_name = value
	obj.modified.SetBit(&obj.modified, instance_ip_display_name, 1)
}

func (obj *InstanceIp) readFloatingIps() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(instance_ip_floating_ips) == 0) {
		err := obj.GetField(obj, "floating_ips")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *InstanceIp) GetFloatingIps() (
	contrail.ReferenceList, error) {
	err := obj.readFloatingIps()
	if err != nil {
		return nil, err
	}
	return obj.floating_ips, nil
}

func (obj *InstanceIp) readVirtualNetworkRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(instance_ip_virtual_network_refs) == 0) {
		err := obj.GetField(obj, "virtual_network_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *InstanceIp) GetVirtualNetworkRefs() (
	contrail.ReferenceList, error) {
	err := obj.readVirtualNetworkRefs()
	if err != nil {
		return nil, err
	}
	return obj.virtual_network_refs, nil
}

func (obj *InstanceIp) AddVirtualNetwork(
	rhs *VirtualNetwork) error {
	err := obj.readVirtualNetworkRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(instance_ip_virtual_network_refs) == 0 {
		obj.storeReferenceBase("virtual-network", obj.virtual_network_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.virtual_network_refs = append(obj.virtual_network_refs, ref)
	obj.modified.SetBit(&obj.modified, instance_ip_virtual_network_refs, 1)
	return nil
}

func (obj *InstanceIp) DeleteVirtualNetwork(uuid string) error {
	err := obj.readVirtualNetworkRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(instance_ip_virtual_network_refs) == 0 {
		obj.storeReferenceBase("virtual-network", obj.virtual_network_refs)
	}

	for i, ref := range obj.virtual_network_refs {
		if ref.Uuid == uuid {
			obj.virtual_network_refs = append(
				obj.virtual_network_refs[:i],
				obj.virtual_network_refs[i+1:]...)
			break
		}
	}
	obj.modified.SetBit(&obj.modified, instance_ip_virtual_network_refs, 1)
	return nil
}

func (obj *InstanceIp) ClearVirtualNetwork() {
	if (obj.valid.Bit(instance_ip_virtual_network_refs) != 0) &&
		(obj.modified.Bit(instance_ip_virtual_network_refs) == 0) {
		obj.storeReferenceBase("virtual-network", obj.virtual_network_refs)
	}
	obj.virtual_network_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, instance_ip_virtual_network_refs, 1)
	obj.modified.SetBit(&obj.modified, instance_ip_virtual_network_refs, 1)
}

func (obj *InstanceIp) SetVirtualNetworkList(
	refList []contrail.ReferencePair) {
	obj.ClearVirtualNetwork()
	obj.virtual_network_refs = make([]contrail.Reference, len(refList))
	for i, pair := range refList {
		obj.virtual_network_refs[i] = contrail.Reference{
			pair.Object.GetFQName(),
			pair.Object.GetUuid(),
			pair.Object.GetHref(),
			pair.Attribute,
		}
	}
}

func (obj *InstanceIp) readNetworkIpamRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(instance_ip_network_ipam_refs) == 0) {
		err := obj.GetField(obj, "network_ipam_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *InstanceIp) GetNetworkIpamRefs() (
	contrail.ReferenceList, error) {
	err := obj.readNetworkIpamRefs()
	if err != nil {
		return nil, err
	}
	return obj.network_ipam_refs, nil
}

func (obj *InstanceIp) AddNetworkIpam(
	rhs *NetworkIpam) error {
	err := obj.readNetworkIpamRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(instance_ip_network_ipam_refs) == 0 {
		obj.storeReferenceBase("network-ipam", obj.network_ipam_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.network_ipam_refs = append(obj.network_ipam_refs, ref)
	obj.modified.SetBit(&obj.modified, instance_ip_network_ipam_refs, 1)
	return nil
}

func (obj *InstanceIp) DeleteNetworkIpam(uuid string) error {
	err := obj.readNetworkIpamRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(instance_ip_network_ipam_refs) == 0 {
		obj.storeReferenceBase("network-ipam", obj.network_ipam_refs)
	}

	for i, ref := range obj.network_ipam_refs {
		if ref.Uuid == uuid {
			obj.network_ipam_refs = append(
				obj.network_ipam_refs[:i],
				obj.network_ipam_refs[i+1:]...)
			break
		}
	}
	obj.modified.SetBit(&obj.modified, instance_ip_network_ipam_refs, 1)
	return nil
}

func (obj *InstanceIp) ClearNetworkIpam() {
	if (obj.valid.Bit(instance_ip_network_ipam_refs) != 0) &&
		(obj.modified.Bit(instance_ip_network_ipam_refs) == 0) {
		obj.storeReferenceBase("network-ipam", obj.network_ipam_refs)
	}
	obj.network_ipam_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, instance_ip_network_ipam_refs, 1)
	obj.modified.SetBit(&obj.modified, instance_ip_network_ipam_refs, 1)
}

func (obj *InstanceIp) SetNetworkIpamList(
	refList []contrail.ReferencePair) {
	obj.ClearNetworkIpam()
	obj.network_ipam_refs = make([]contrail.Reference, len(refList))
	for i, pair := range refList {
		obj.network_ipam_refs[i] = contrail.Reference{
			pair.Object.GetFQName(),
			pair.Object.GetUuid(),
			pair.Object.GetHref(),
			pair.Attribute,
		}
	}
}

func (obj *InstanceIp) readVirtualMachineInterfaceRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(instance_ip_virtual_machine_interface_refs) == 0) {
		err := obj.GetField(obj, "virtual_machine_interface_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *InstanceIp) GetVirtualMachineInterfaceRefs() (
	contrail.ReferenceList, error) {
	err := obj.readVirtualMachineInterfaceRefs()
	if err != nil {
		return nil, err
	}
	return obj.virtual_machine_interface_refs, nil
}

func (obj *InstanceIp) AddVirtualMachineInterface(
	rhs *VirtualMachineInterface) error {
	err := obj.readVirtualMachineInterfaceRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(instance_ip_virtual_machine_interface_refs) == 0 {
		obj.storeReferenceBase("virtual-machine-interface", obj.virtual_machine_interface_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.virtual_machine_interface_refs = append(obj.virtual_machine_interface_refs, ref)
	obj.modified.SetBit(&obj.modified, instance_ip_virtual_machine_interface_refs, 1)
	return nil
}

func (obj *InstanceIp) DeleteVirtualMachineInterface(uuid string) error {
	err := obj.readVirtualMachineInterfaceRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(instance_ip_virtual_machine_interface_refs) == 0 {
		obj.storeReferenceBase("virtual-machine-interface", obj.virtual_machine_interface_refs)
	}

	for i, ref := range obj.virtual_machine_interface_refs {
		if ref.Uuid == uuid {
			obj.virtual_machine_interface_refs = append(
				obj.virtual_machine_interface_refs[:i],
				obj.virtual_machine_interface_refs[i+1:]...)
			break
		}
	}
	obj.modified.SetBit(&obj.modified, instance_ip_virtual_machine_interface_refs, 1)
	return nil
}

func (obj *InstanceIp) ClearVirtualMachineInterface() {
	if (obj.valid.Bit(instance_ip_virtual_machine_interface_refs) != 0) &&
		(obj.modified.Bit(instance_ip_virtual_machine_interface_refs) == 0) {
		obj.storeReferenceBase("virtual-machine-interface", obj.virtual_machine_interface_refs)
	}
	obj.virtual_machine_interface_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, instance_ip_virtual_machine_interface_refs, 1)
	obj.modified.SetBit(&obj.modified, instance_ip_virtual_machine_interface_refs, 1)
}

func (obj *InstanceIp) SetVirtualMachineInterfaceList(
	refList []contrail.ReferencePair) {
	obj.ClearVirtualMachineInterface()
	obj.virtual_machine_interface_refs = make([]contrail.Reference, len(refList))
	for i, pair := range refList {
		obj.virtual_machine_interface_refs[i] = contrail.Reference{
			pair.Object.GetFQName(),
			pair.Object.GetUuid(),
			pair.Object.GetHref(),
			pair.Attribute,
		}
	}
}

func (obj *InstanceIp) readPhysicalRouterRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(instance_ip_physical_router_refs) == 0) {
		err := obj.GetField(obj, "physical_router_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *InstanceIp) GetPhysicalRouterRefs() (
	contrail.ReferenceList, error) {
	err := obj.readPhysicalRouterRefs()
	if err != nil {
		return nil, err
	}
	return obj.physical_router_refs, nil
}

func (obj *InstanceIp) AddPhysicalRouter(
	rhs *PhysicalRouter) error {
	err := obj.readPhysicalRouterRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(instance_ip_physical_router_refs) == 0 {
		obj.storeReferenceBase("physical-router", obj.physical_router_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.physical_router_refs = append(obj.physical_router_refs, ref)
	obj.modified.SetBit(&obj.modified, instance_ip_physical_router_refs, 1)
	return nil
}

func (obj *InstanceIp) DeletePhysicalRouter(uuid string) error {
	err := obj.readPhysicalRouterRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(instance_ip_physical_router_refs) == 0 {
		obj.storeReferenceBase("physical-router", obj.physical_router_refs)
	}

	for i, ref := range obj.physical_router_refs {
		if ref.Uuid == uuid {
			obj.physical_router_refs = append(
				obj.physical_router_refs[:i],
				obj.physical_router_refs[i+1:]...)
			break
		}
	}
	obj.modified.SetBit(&obj.modified, instance_ip_physical_router_refs, 1)
	return nil
}

func (obj *InstanceIp) ClearPhysicalRouter() {
	if (obj.valid.Bit(instance_ip_physical_router_refs) != 0) &&
		(obj.modified.Bit(instance_ip_physical_router_refs) == 0) {
		obj.storeReferenceBase("physical-router", obj.physical_router_refs)
	}
	obj.physical_router_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, instance_ip_physical_router_refs, 1)
	obj.modified.SetBit(&obj.modified, instance_ip_physical_router_refs, 1)
}

func (obj *InstanceIp) SetPhysicalRouterList(
	refList []contrail.ReferencePair) {
	obj.ClearPhysicalRouter()
	obj.physical_router_refs = make([]contrail.Reference, len(refList))
	for i, pair := range refList {
		obj.physical_router_refs[i] = contrail.Reference{
			pair.Object.GetFQName(),
			pair.Object.GetUuid(),
			pair.Object.GetHref(),
			pair.Attribute,
		}
	}
}

func (obj *InstanceIp) readVirtualRouterRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(instance_ip_virtual_router_refs) == 0) {
		err := obj.GetField(obj, "virtual_router_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *InstanceIp) GetVirtualRouterRefs() (
	contrail.ReferenceList, error) {
	err := obj.readVirtualRouterRefs()
	if err != nil {
		return nil, err
	}
	return obj.virtual_router_refs, nil
}

func (obj *InstanceIp) AddVirtualRouter(
	rhs *VirtualRouter) error {
	err := obj.readVirtualRouterRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(instance_ip_virtual_router_refs) == 0 {
		obj.storeReferenceBase("virtual-router", obj.virtual_router_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.virtual_router_refs = append(obj.virtual_router_refs, ref)
	obj.modified.SetBit(&obj.modified, instance_ip_virtual_router_refs, 1)
	return nil
}

func (obj *InstanceIp) DeleteVirtualRouter(uuid string) error {
	err := obj.readVirtualRouterRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(instance_ip_virtual_router_refs) == 0 {
		obj.storeReferenceBase("virtual-router", obj.virtual_router_refs)
	}

	for i, ref := range obj.virtual_router_refs {
		if ref.Uuid == uuid {
			obj.virtual_router_refs = append(
				obj.virtual_router_refs[:i],
				obj.virtual_router_refs[i+1:]...)
			break
		}
	}
	obj.modified.SetBit(&obj.modified, instance_ip_virtual_router_refs, 1)
	return nil
}

func (obj *InstanceIp) ClearVirtualRouter() {
	if (obj.valid.Bit(instance_ip_virtual_router_refs) != 0) &&
		(obj.modified.Bit(instance_ip_virtual_router_refs) == 0) {
		obj.storeReferenceBase("virtual-router", obj.virtual_router_refs)
	}
	obj.virtual_router_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, instance_ip_virtual_router_refs, 1)
	obj.modified.SetBit(&obj.modified, instance_ip_virtual_router_refs, 1)
}

func (obj *InstanceIp) SetVirtualRouterList(
	refList []contrail.ReferencePair) {
	obj.ClearVirtualRouter()
	obj.virtual_router_refs = make([]contrail.Reference, len(refList))
	for i, pair := range refList {
		obj.virtual_router_refs[i] = contrail.Reference{
			pair.Object.GetFQName(),
			pair.Object.GetUuid(),
			pair.Object.GetHref(),
			pair.Attribute,
		}
	}
}

func (obj *InstanceIp) readTagRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(instance_ip_tag_refs) == 0) {
		err := obj.GetField(obj, "tag_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *InstanceIp) GetTagRefs() (
	contrail.ReferenceList, error) {
	err := obj.readTagRefs()
	if err != nil {
		return nil, err
	}
	return obj.tag_refs, nil
}

func (obj *InstanceIp) AddTag(
	rhs *Tag) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(instance_ip_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.tag_refs = append(obj.tag_refs, ref)
	obj.modified.SetBit(&obj.modified, instance_ip_tag_refs, 1)
	return nil
}

func (obj *InstanceIp) DeleteTag(uuid string) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(instance_ip_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	for i, ref := range obj.tag_refs {
		if ref.Uuid == uuid {
			obj.tag_refs = append(
				obj.tag_refs[:i],
				obj.tag_refs[i+1:]...)
			break
		}
	}
	obj.modified.SetBit(&obj.modified, instance_ip_tag_refs, 1)
	return nil
}

func (obj *InstanceIp) ClearTag() {
	if (obj.valid.Bit(instance_ip_tag_refs) != 0) &&
		(obj.modified.Bit(instance_ip_tag_refs) == 0) {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}
	obj.tag_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, instance_ip_tag_refs, 1)
	obj.modified.SetBit(&obj.modified, instance_ip_tag_refs, 1)
}

func (obj *InstanceIp) SetTagList(
	refList []contrail.ReferencePair) {
	obj.ClearTag()
	obj.tag_refs = make([]contrail.Reference, len(refList))
	for i, pair := range refList {
		obj.tag_refs[i] = contrail.Reference{
			pair.Object.GetFQName(),
			pair.Object.GetUuid(),
			pair.Object.GetHref(),
			pair.Attribute,
		}
	}
}

func (obj *InstanceIp) readServiceInstanceBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(instance_ip_service_instance_back_refs) == 0) {
		err := obj.GetField(obj, "service_instance_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *InstanceIp) GetServiceInstanceBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readServiceInstanceBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.service_instance_back_refs, nil
}

func (obj *InstanceIp) MarshalJSON() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalCommon(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(instance_ip_instance_ip_address) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.instance_ip_address)
		if err != nil {
			return nil, err
		}
		msg["instance_ip_address"] = &value
	}

	if obj.modified.Bit(instance_ip_instance_ip_family) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.instance_ip_family)
		if err != nil {
			return nil, err
		}
		msg["instance_ip_family"] = &value
	}

	if obj.modified.Bit(instance_ip_instance_ip_mode) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.instance_ip_mode)
		if err != nil {
			return nil, err
		}
		msg["instance_ip_mode"] = &value
	}

	if obj.modified.Bit(instance_ip_secondary_ip_tracking_ip) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.secondary_ip_tracking_ip)
		if err != nil {
			return nil, err
		}
		msg["secondary_ip_tracking_ip"] = &value
	}

	if obj.modified.Bit(instance_ip_subnet_uuid) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.subnet_uuid)
		if err != nil {
			return nil, err
		}
		msg["subnet_uuid"] = &value
	}

	if obj.modified.Bit(instance_ip_instance_ip_secondary) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.instance_ip_secondary)
		if err != nil {
			return nil, err
		}
		msg["instance_ip_secondary"] = &value
	}

	if obj.modified.Bit(instance_ip_instance_ip_local_ip) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.instance_ip_local_ip)
		if err != nil {
			return nil, err
		}
		msg["instance_ip_local_ip"] = &value
	}

	if obj.modified.Bit(instance_ip_service_instance_ip) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.service_instance_ip)
		if err != nil {
			return nil, err
		}
		msg["service_instance_ip"] = &value
	}

	if obj.modified.Bit(instance_ip_service_health_check_ip) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.service_health_check_ip)
		if err != nil {
			return nil, err
		}
		msg["service_health_check_ip"] = &value
	}

	if obj.modified.Bit(instance_ip_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(instance_ip_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(instance_ip_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(instance_ip_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if len(obj.virtual_network_refs) > 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.virtual_network_refs)
		if err != nil {
			return nil, err
		}
		msg["virtual_network_refs"] = &value
	}

	if len(obj.network_ipam_refs) > 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.network_ipam_refs)
		if err != nil {
			return nil, err
		}
		msg["network_ipam_refs"] = &value
	}

	if len(obj.virtual_machine_interface_refs) > 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.virtual_machine_interface_refs)
		if err != nil {
			return nil, err
		}
		msg["virtual_machine_interface_refs"] = &value
	}

	if len(obj.physical_router_refs) > 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.physical_router_refs)
		if err != nil {
			return nil, err
		}
		msg["physical_router_refs"] = &value
	}

	if len(obj.virtual_router_refs) > 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.virtual_router_refs)
		if err != nil {
			return nil, err
		}
		msg["virtual_router_refs"] = &value
	}

	if len(obj.tag_refs) > 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.tag_refs)
		if err != nil {
			return nil, err
		}
		msg["tag_refs"] = &value
	}

	return json.Marshal(msg)
}

func (obj *InstanceIp) UnmarshalJSON(body []byte) error {
	var m map[string]json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	err = obj.UnmarshalCommon(m)
	if err != nil {
		return err
	}
	for key, value := range m {
		switch key {
		case "instance_ip_address":
			err = json.Unmarshal(value, &obj.instance_ip_address)
			if err == nil {
				obj.valid.SetBit(&obj.valid, instance_ip_instance_ip_address, 1)
			}
			break
		case "instance_ip_family":
			err = json.Unmarshal(value, &obj.instance_ip_family)
			if err == nil {
				obj.valid.SetBit(&obj.valid, instance_ip_instance_ip_family, 1)
			}
			break
		case "instance_ip_mode":
			err = json.Unmarshal(value, &obj.instance_ip_mode)
			if err == nil {
				obj.valid.SetBit(&obj.valid, instance_ip_instance_ip_mode, 1)
			}
			break
		case "secondary_ip_tracking_ip":
			err = json.Unmarshal(value, &obj.secondary_ip_tracking_ip)
			if err == nil {
				obj.valid.SetBit(&obj.valid, instance_ip_secondary_ip_tracking_ip, 1)
			}
			break
		case "subnet_uuid":
			err = json.Unmarshal(value, &obj.subnet_uuid)
			if err == nil {
				obj.valid.SetBit(&obj.valid, instance_ip_subnet_uuid, 1)
			}
			break
		case "instance_ip_secondary":
			err = json.Unmarshal(value, &obj.instance_ip_secondary)
			if err == nil {
				obj.valid.SetBit(&obj.valid, instance_ip_instance_ip_secondary, 1)
			}
			break
		case "instance_ip_local_ip":
			err = json.Unmarshal(value, &obj.instance_ip_local_ip)
			if err == nil {
				obj.valid.SetBit(&obj.valid, instance_ip_instance_ip_local_ip, 1)
			}
			break
		case "service_instance_ip":
			err = json.Unmarshal(value, &obj.service_instance_ip)
			if err == nil {
				obj.valid.SetBit(&obj.valid, instance_ip_service_instance_ip, 1)
			}
			break
		case "service_health_check_ip":
			err = json.Unmarshal(value, &obj.service_health_check_ip)
			if err == nil {
				obj.valid.SetBit(&obj.valid, instance_ip_service_health_check_ip, 1)
			}
			break
		case "id_perms":
			err = json.Unmarshal(value, &obj.id_perms)
			if err == nil {
				obj.valid.SetBit(&obj.valid, instance_ip_id_perms, 1)
			}
			break
		case "perms2":
			err = json.Unmarshal(value, &obj.perms2)
			if err == nil {
				obj.valid.SetBit(&obj.valid, instance_ip_perms2, 1)
			}
			break
		case "annotations":
			err = json.Unmarshal(value, &obj.annotations)
			if err == nil {
				obj.valid.SetBit(&obj.valid, instance_ip_annotations, 1)
			}
			break
		case "display_name":
			err = json.Unmarshal(value, &obj.display_name)
			if err == nil {
				obj.valid.SetBit(&obj.valid, instance_ip_display_name, 1)
			}
			break
		case "virtual_network_refs":
			err = json.Unmarshal(value, &obj.virtual_network_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, instance_ip_virtual_network_refs, 1)
			}
			break
		case "network_ipam_refs":
			err = json.Unmarshal(value, &obj.network_ipam_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, instance_ip_network_ipam_refs, 1)
			}
			break
		case "virtual_machine_interface_refs":
			err = json.Unmarshal(value, &obj.virtual_machine_interface_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, instance_ip_virtual_machine_interface_refs, 1)
			}
			break
		case "physical_router_refs":
			err = json.Unmarshal(value, &obj.physical_router_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, instance_ip_physical_router_refs, 1)
			}
			break
		case "virtual_router_refs":
			err = json.Unmarshal(value, &obj.virtual_router_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, instance_ip_virtual_router_refs, 1)
			}
			break
		case "floating_ips":
			err = json.Unmarshal(value, &obj.floating_ips)
			if err == nil {
				obj.valid.SetBit(&obj.valid, instance_ip_floating_ips, 1)
			}
			break
		case "tag_refs":
			err = json.Unmarshal(value, &obj.tag_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, instance_ip_tag_refs, 1)
			}
			break
		case "service_instance_back_refs":
			{
				type ReferenceElement struct {
					To   []string
					Uuid string
					Href string
					Attr ServiceInterfaceTag
				}
				var array []ReferenceElement
				err = json.Unmarshal(value, &array)
				if err != nil {
					break
				}
				obj.valid.SetBit(&obj.valid, instance_ip_service_instance_back_refs, 1)
				obj.service_instance_back_refs = make(contrail.ReferenceList, 0)
				for _, element := range array {
					ref := contrail.Reference{
						element.To,
						element.Uuid,
						element.Href,
						element.Attr,
					}
					obj.service_instance_back_refs = append(obj.service_instance_back_refs, ref)
				}
				break
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *InstanceIp) UpdateObject() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalId(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(instance_ip_instance_ip_address) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.instance_ip_address)
		if err != nil {
			return nil, err
		}
		msg["instance_ip_address"] = &value
	}

	if obj.modified.Bit(instance_ip_instance_ip_family) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.instance_ip_family)
		if err != nil {
			return nil, err
		}
		msg["instance_ip_family"] = &value
	}

	if obj.modified.Bit(instance_ip_instance_ip_mode) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.instance_ip_mode)
		if err != nil {
			return nil, err
		}
		msg["instance_ip_mode"] = &value
	}

	if obj.modified.Bit(instance_ip_secondary_ip_tracking_ip) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.secondary_ip_tracking_ip)
		if err != nil {
			return nil, err
		}
		msg["secondary_ip_tracking_ip"] = &value
	}

	if obj.modified.Bit(instance_ip_subnet_uuid) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.subnet_uuid)
		if err != nil {
			return nil, err
		}
		msg["subnet_uuid"] = &value
	}

	if obj.modified.Bit(instance_ip_instance_ip_secondary) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.instance_ip_secondary)
		if err != nil {
			return nil, err
		}
		msg["instance_ip_secondary"] = &value
	}

	if obj.modified.Bit(instance_ip_instance_ip_local_ip) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.instance_ip_local_ip)
		if err != nil {
			return nil, err
		}
		msg["instance_ip_local_ip"] = &value
	}

	if obj.modified.Bit(instance_ip_service_instance_ip) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.service_instance_ip)
		if err != nil {
			return nil, err
		}
		msg["service_instance_ip"] = &value
	}

	if obj.modified.Bit(instance_ip_service_health_check_ip) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.service_health_check_ip)
		if err != nil {
			return nil, err
		}
		msg["service_health_check_ip"] = &value
	}

	if obj.modified.Bit(instance_ip_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(instance_ip_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(instance_ip_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(instance_ip_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if obj.modified.Bit(instance_ip_virtual_network_refs) != 0 {
		if len(obj.virtual_network_refs) == 0 {
			var value json.RawMessage
			value, err := json.Marshal(
				make([]contrail.Reference, 0))
			if err != nil {
				return nil, err
			}
			msg["virtual_network_refs"] = &value
		} else if !obj.hasReferenceBase("virtual-network") {
			var value json.RawMessage
			value, err := json.Marshal(&obj.virtual_network_refs)
			if err != nil {
				return nil, err
			}
			msg["virtual_network_refs"] = &value
		}
	}

	if obj.modified.Bit(instance_ip_network_ipam_refs) != 0 {
		if len(obj.network_ipam_refs) == 0 {
			var value json.RawMessage
			value, err := json.Marshal(
				make([]contrail.Reference, 0))
			if err != nil {
				return nil, err
			}
			msg["network_ipam_refs"] = &value
		} else if !obj.hasReferenceBase("network-ipam") {
			var value json.RawMessage
			value, err := json.Marshal(&obj.network_ipam_refs)
			if err != nil {
				return nil, err
			}
			msg["network_ipam_refs"] = &value
		}
	}

	if obj.modified.Bit(instance_ip_virtual_machine_interface_refs) != 0 {
		if len(obj.virtual_machine_interface_refs) == 0 {
			var value json.RawMessage
			value, err := json.Marshal(
				make([]contrail.Reference, 0))
			if err != nil {
				return nil, err
			}
			msg["virtual_machine_interface_refs"] = &value
		} else if !obj.hasReferenceBase("virtual-machine-interface") {
			var value json.RawMessage
			value, err := json.Marshal(&obj.virtual_machine_interface_refs)
			if err != nil {
				return nil, err
			}
			msg["virtual_machine_interface_refs"] = &value
		}
	}

	if obj.modified.Bit(instance_ip_physical_router_refs) != 0 {
		if len(obj.physical_router_refs) == 0 {
			var value json.RawMessage
			value, err := json.Marshal(
				make([]contrail.Reference, 0))
			if err != nil {
				return nil, err
			}
			msg["physical_router_refs"] = &value
		} else if !obj.hasReferenceBase("physical-router") {
			var value json.RawMessage
			value, err := json.Marshal(&obj.physical_router_refs)
			if err != nil {
				return nil, err
			}
			msg["physical_router_refs"] = &value
		}
	}

	if obj.modified.Bit(instance_ip_virtual_router_refs) != 0 {
		if len(obj.virtual_router_refs) == 0 {
			var value json.RawMessage
			value, err := json.Marshal(
				make([]contrail.Reference, 0))
			if err != nil {
				return nil, err
			}
			msg["virtual_router_refs"] = &value
		} else if !obj.hasReferenceBase("virtual-router") {
			var value json.RawMessage
			value, err := json.Marshal(&obj.virtual_router_refs)
			if err != nil {
				return nil, err
			}
			msg["virtual_router_refs"] = &value
		}
	}

	if obj.modified.Bit(instance_ip_tag_refs) != 0 {
		if len(obj.tag_refs) == 0 {
			var value json.RawMessage
			value, err := json.Marshal(
				make([]contrail.Reference, 0))
			if err != nil {
				return nil, err
			}
			msg["tag_refs"] = &value
		} else if !obj.hasReferenceBase("tag") {
			var value json.RawMessage
			value, err := json.Marshal(&obj.tag_refs)
			if err != nil {
				return nil, err
			}
			msg["tag_refs"] = &value
		}
	}

	return json.Marshal(msg)
}

func (obj *InstanceIp) UpdateReferences() error {

	if (obj.modified.Bit(instance_ip_virtual_network_refs) != 0) &&
		len(obj.virtual_network_refs) > 0 &&
		obj.hasReferenceBase("virtual-network") {
		err := obj.UpdateReference(
			obj, "virtual-network",
			obj.virtual_network_refs,
			obj.baseMap["virtual-network"])
		if err != nil {
			return err
		}
	}

	if (obj.modified.Bit(instance_ip_network_ipam_refs) != 0) &&
		len(obj.network_ipam_refs) > 0 &&
		obj.hasReferenceBase("network-ipam") {
		err := obj.UpdateReference(
			obj, "network-ipam",
			obj.network_ipam_refs,
			obj.baseMap["network-ipam"])
		if err != nil {
			return err
		}
	}

	if (obj.modified.Bit(instance_ip_virtual_machine_interface_refs) != 0) &&
		len(obj.virtual_machine_interface_refs) > 0 &&
		obj.hasReferenceBase("virtual-machine-interface") {
		err := obj.UpdateReference(
			obj, "virtual-machine-interface",
			obj.virtual_machine_interface_refs,
			obj.baseMap["virtual-machine-interface"])
		if err != nil {
			return err
		}
	}

	if (obj.modified.Bit(instance_ip_physical_router_refs) != 0) &&
		len(obj.physical_router_refs) > 0 &&
		obj.hasReferenceBase("physical-router") {
		err := obj.UpdateReference(
			obj, "physical-router",
			obj.physical_router_refs,
			obj.baseMap["physical-router"])
		if err != nil {
			return err
		}
	}

	if (obj.modified.Bit(instance_ip_virtual_router_refs) != 0) &&
		len(obj.virtual_router_refs) > 0 &&
		obj.hasReferenceBase("virtual-router") {
		err := obj.UpdateReference(
			obj, "virtual-router",
			obj.virtual_router_refs,
			obj.baseMap["virtual-router"])
		if err != nil {
			return err
		}
	}

	if (obj.modified.Bit(instance_ip_tag_refs) != 0) &&
		len(obj.tag_refs) > 0 &&
		obj.hasReferenceBase("tag") {
		err := obj.UpdateReference(
			obj, "tag",
			obj.tag_refs,
			obj.baseMap["tag"])
		if err != nil {
			return err
		}
	}

	return nil
}

func InstanceIpByName(c contrail.ApiClient, fqn string) (*InstanceIp, error) {
	obj, err := c.FindByName("instance-ip", fqn)
	if err != nil {
		return nil, err
	}
	return obj.(*InstanceIp), nil
}

func InstanceIpByUuid(c contrail.ApiClient, uuid string) (*InstanceIp, error) {
	obj, err := c.FindByUuid("instance-ip", uuid)
	if err != nil {
		return nil, err
	}
	return obj.(*InstanceIp), nil
}
