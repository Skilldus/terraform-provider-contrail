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
	virtual_machine_id_perms int = iota
	virtual_machine_perms2
	virtual_machine_annotations
	virtual_machine_display_name
	virtual_machine_virtual_machine_interfaces
	virtual_machine_service_instance_refs
	virtual_machine_tag_refs
	virtual_machine_virtual_machine_interface_back_refs
	virtual_machine_virtual_router_back_refs
)

type VirtualMachine struct {
	contrail.ObjectBase
	id_perms                            IdPermsType
	perms2                              PermType2
	annotations                         KeyValuePairs
	display_name                        string
	virtual_machine_interfaces          contrail.ReferenceList
	service_instance_refs               contrail.ReferenceList
	tag_refs                            contrail.ReferenceList
	virtual_machine_interface_back_refs contrail.ReferenceList
	virtual_router_back_refs            contrail.ReferenceList
	valid                               big.Int
	modified                            big.Int
	baseMap                             map[string]contrail.ReferenceList
}

func (obj *VirtualMachine) GetType() string {
	return "virtual-machine"
}

func (obj *VirtualMachine) GetDefaultParent() []string {
	name := []string{}
	return name
}

func (obj *VirtualMachine) GetDefaultParentType() string {
	return ""
}

func (obj *VirtualMachine) SetName(name string) {
	obj.VSetName(obj, name)
}

func (obj *VirtualMachine) SetParent(parent contrail.IObject) {
	obj.VSetParent(obj, parent)
}

func (obj *VirtualMachine) storeReferenceBase(
	name string, refList contrail.ReferenceList) {
	if obj.baseMap == nil {
		obj.baseMap = make(map[string]contrail.ReferenceList)
	}
	refCopy := make(contrail.ReferenceList, len(refList))
	copy(refCopy, refList)
	obj.baseMap[name] = refCopy
}

func (obj *VirtualMachine) hasReferenceBase(name string) bool {
	if obj.baseMap == nil {
		return false
	}
	_, exists := obj.baseMap[name]
	return exists
}

func (obj *VirtualMachine) UpdateDone() {
	obj.modified.SetInt64(0)
	obj.baseMap = nil
}

func (obj *VirtualMachine) GetIdPerms() IdPermsType {
	return obj.id_perms
}

func (obj *VirtualMachine) SetIdPerms(value *IdPermsType) {
	obj.id_perms = *value
	obj.modified.SetBit(&obj.modified, virtual_machine_id_perms, 1)
}

func (obj *VirtualMachine) GetPerms2() PermType2 {
	return obj.perms2
}

func (obj *VirtualMachine) SetPerms2(value *PermType2) {
	obj.perms2 = *value
	obj.modified.SetBit(&obj.modified, virtual_machine_perms2, 1)
}

func (obj *VirtualMachine) GetAnnotations() KeyValuePairs {
	return obj.annotations
}

func (obj *VirtualMachine) SetAnnotations(value *KeyValuePairs) {
	obj.annotations = *value
	obj.modified.SetBit(&obj.modified, virtual_machine_annotations, 1)
}

func (obj *VirtualMachine) GetDisplayName() string {
	return obj.display_name
}

func (obj *VirtualMachine) SetDisplayName(value string) {
	obj.display_name = value
	obj.modified.SetBit(&obj.modified, virtual_machine_display_name, 1)
}

func (obj *VirtualMachine) readVirtualMachineInterfaces() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(virtual_machine_virtual_machine_interfaces) == 0) {
		err := obj.GetField(obj, "virtual_machine_interfaces")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *VirtualMachine) GetVirtualMachineInterfaces() (
	contrail.ReferenceList, error) {
	err := obj.readVirtualMachineInterfaces()
	if err != nil {
		return nil, err
	}
	return obj.virtual_machine_interfaces, nil
}

func (obj *VirtualMachine) readServiceInstanceRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(virtual_machine_service_instance_refs) == 0) {
		err := obj.GetField(obj, "service_instance_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *VirtualMachine) GetServiceInstanceRefs() (
	contrail.ReferenceList, error) {
	err := obj.readServiceInstanceRefs()
	if err != nil {
		return nil, err
	}
	return obj.service_instance_refs, nil
}

func (obj *VirtualMachine) AddServiceInstance(
	rhs *ServiceInstance) error {
	err := obj.readServiceInstanceRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(virtual_machine_service_instance_refs) == 0 {
		obj.storeReferenceBase("service-instance", obj.service_instance_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.service_instance_refs = append(obj.service_instance_refs, ref)
	obj.modified.SetBit(&obj.modified, virtual_machine_service_instance_refs, 1)
	return nil
}

func (obj *VirtualMachine) DeleteServiceInstance(uuid string) error {
	err := obj.readServiceInstanceRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(virtual_machine_service_instance_refs) == 0 {
		obj.storeReferenceBase("service-instance", obj.service_instance_refs)
	}

	for i, ref := range obj.service_instance_refs {
		if ref.Uuid == uuid {
			obj.service_instance_refs = append(
				obj.service_instance_refs[:i],
				obj.service_instance_refs[i+1:]...)
			break
		}
	}
	obj.modified.SetBit(&obj.modified, virtual_machine_service_instance_refs, 1)
	return nil
}

func (obj *VirtualMachine) ClearServiceInstance() {
	if (obj.valid.Bit(virtual_machine_service_instance_refs) != 0) &&
		(obj.modified.Bit(virtual_machine_service_instance_refs) == 0) {
		obj.storeReferenceBase("service-instance", obj.service_instance_refs)
	}
	obj.service_instance_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, virtual_machine_service_instance_refs, 1)
	obj.modified.SetBit(&obj.modified, virtual_machine_service_instance_refs, 1)
}

func (obj *VirtualMachine) SetServiceInstanceList(
	refList []contrail.ReferencePair) {
	obj.ClearServiceInstance()
	obj.service_instance_refs = make([]contrail.Reference, len(refList))
	for i, pair := range refList {
		obj.service_instance_refs[i] = contrail.Reference{
			pair.Object.GetFQName(),
			pair.Object.GetUuid(),
			pair.Object.GetHref(),
			pair.Attribute,
		}
	}
}

func (obj *VirtualMachine) readTagRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(virtual_machine_tag_refs) == 0) {
		err := obj.GetField(obj, "tag_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *VirtualMachine) GetTagRefs() (
	contrail.ReferenceList, error) {
	err := obj.readTagRefs()
	if err != nil {
		return nil, err
	}
	return obj.tag_refs, nil
}

func (obj *VirtualMachine) AddTag(
	rhs *Tag) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(virtual_machine_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.tag_refs = append(obj.tag_refs, ref)
	obj.modified.SetBit(&obj.modified, virtual_machine_tag_refs, 1)
	return nil
}

func (obj *VirtualMachine) DeleteTag(uuid string) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(virtual_machine_tag_refs) == 0 {
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
	obj.modified.SetBit(&obj.modified, virtual_machine_tag_refs, 1)
	return nil
}

func (obj *VirtualMachine) ClearTag() {
	if (obj.valid.Bit(virtual_machine_tag_refs) != 0) &&
		(obj.modified.Bit(virtual_machine_tag_refs) == 0) {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}
	obj.tag_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, virtual_machine_tag_refs, 1)
	obj.modified.SetBit(&obj.modified, virtual_machine_tag_refs, 1)
}

func (obj *VirtualMachine) SetTagList(
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

func (obj *VirtualMachine) readVirtualMachineInterfaceBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(virtual_machine_virtual_machine_interface_back_refs) == 0) {
		err := obj.GetField(obj, "virtual_machine_interface_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *VirtualMachine) GetVirtualMachineInterfaceBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readVirtualMachineInterfaceBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.virtual_machine_interface_back_refs, nil
}

func (obj *VirtualMachine) readVirtualRouterBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(virtual_machine_virtual_router_back_refs) == 0) {
		err := obj.GetField(obj, "virtual_router_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *VirtualMachine) GetVirtualRouterBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readVirtualRouterBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.virtual_router_back_refs, nil
}

func (obj *VirtualMachine) MarshalJSON() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalCommon(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(virtual_machine_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(virtual_machine_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(virtual_machine_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(virtual_machine_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if len(obj.service_instance_refs) > 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.service_instance_refs)
		if err != nil {
			return nil, err
		}
		msg["service_instance_refs"] = &value
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

func (obj *VirtualMachine) UnmarshalJSON(body []byte) error {
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
		case "id_perms":
			err = json.Unmarshal(value, &obj.id_perms)
			if err == nil {
				obj.valid.SetBit(&obj.valid, virtual_machine_id_perms, 1)
			}
			break
		case "perms2":
			err = json.Unmarshal(value, &obj.perms2)
			if err == nil {
				obj.valid.SetBit(&obj.valid, virtual_machine_perms2, 1)
			}
			break
		case "annotations":
			err = json.Unmarshal(value, &obj.annotations)
			if err == nil {
				obj.valid.SetBit(&obj.valid, virtual_machine_annotations, 1)
			}
			break
		case "display_name":
			err = json.Unmarshal(value, &obj.display_name)
			if err == nil {
				obj.valid.SetBit(&obj.valid, virtual_machine_display_name, 1)
			}
			break
		case "virtual_machine_interfaces":
			err = json.Unmarshal(value, &obj.virtual_machine_interfaces)
			if err == nil {
				obj.valid.SetBit(&obj.valid, virtual_machine_virtual_machine_interfaces, 1)
			}
			break
		case "service_instance_refs":
			err = json.Unmarshal(value, &obj.service_instance_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, virtual_machine_service_instance_refs, 1)
			}
			break
		case "tag_refs":
			err = json.Unmarshal(value, &obj.tag_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, virtual_machine_tag_refs, 1)
			}
			break
		case "virtual_machine_interface_back_refs":
			err = json.Unmarshal(value, &obj.virtual_machine_interface_back_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, virtual_machine_virtual_machine_interface_back_refs, 1)
			}
			break
		case "virtual_router_back_refs":
			err = json.Unmarshal(value, &obj.virtual_router_back_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, virtual_machine_virtual_router_back_refs, 1)
			}
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *VirtualMachine) UpdateObject() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalId(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(virtual_machine_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(virtual_machine_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(virtual_machine_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(virtual_machine_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if obj.modified.Bit(virtual_machine_service_instance_refs) != 0 {
		if len(obj.service_instance_refs) == 0 {
			var value json.RawMessage
			value, err := json.Marshal(
				make([]contrail.Reference, 0))
			if err != nil {
				return nil, err
			}
			msg["service_instance_refs"] = &value
		} else if !obj.hasReferenceBase("service-instance") {
			var value json.RawMessage
			value, err := json.Marshal(&obj.service_instance_refs)
			if err != nil {
				return nil, err
			}
			msg["service_instance_refs"] = &value
		}
	}

	if obj.modified.Bit(virtual_machine_tag_refs) != 0 {
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

func (obj *VirtualMachine) UpdateReferences() error {

	if (obj.modified.Bit(virtual_machine_service_instance_refs) != 0) &&
		len(obj.service_instance_refs) > 0 &&
		obj.hasReferenceBase("service-instance") {
		err := obj.UpdateReference(
			obj, "service-instance",
			obj.service_instance_refs,
			obj.baseMap["service-instance"])
		if err != nil {
			return err
		}
	}

	if (obj.modified.Bit(virtual_machine_tag_refs) != 0) &&
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

func VirtualMachineByName(c contrail.ApiClient, fqn string) (*VirtualMachine, error) {
	obj, err := c.FindByName("virtual-machine", fqn)
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachine), nil
}

func VirtualMachineByUuid(c contrail.ApiClient, uuid string) (*VirtualMachine, error) {
	obj, err := c.FindByUuid("virtual-machine", uuid)
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachine), nil
}
