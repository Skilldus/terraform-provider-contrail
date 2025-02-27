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
	logical_interface_logical_interface_vlan_tag int = iota
	logical_interface_logical_interface_type
	logical_interface_id_perms
	logical_interface_perms2
	logical_interface_annotations
	logical_interface_display_name
	logical_interface_virtual_machine_interface_refs
	logical_interface_tag_refs
)

type LogicalInterface struct {
	contrail.ObjectBase
	logical_interface_vlan_tag     int
	logical_interface_type         string
	id_perms                       IdPermsType
	perms2                         PermType2
	annotations                    KeyValuePairs
	display_name                   string
	virtual_machine_interface_refs contrail.ReferenceList
	tag_refs                       contrail.ReferenceList
	valid                          big.Int
	modified                       big.Int
	baseMap                        map[string]contrail.ReferenceList
}

func (obj *LogicalInterface) GetType() string {
	return "logical-interface"
}

func (obj *LogicalInterface) GetDefaultParent() []string {
	name := []string{"default-global-system-config", "default-physical-router"}
	return name
}

func (obj *LogicalInterface) GetDefaultParentType() string {
	return "physical-router"
}

func (obj *LogicalInterface) SetName(name string) {
	obj.VSetName(obj, name)
}

func (obj *LogicalInterface) SetParent(parent contrail.IObject) {
	obj.VSetParent(obj, parent)
}

func (obj *LogicalInterface) storeReferenceBase(
	name string, refList contrail.ReferenceList) {
	if obj.baseMap == nil {
		obj.baseMap = make(map[string]contrail.ReferenceList)
	}
	refCopy := make(contrail.ReferenceList, len(refList))
	copy(refCopy, refList)
	obj.baseMap[name] = refCopy
}

func (obj *LogicalInterface) hasReferenceBase(name string) bool {
	if obj.baseMap == nil {
		return false
	}
	_, exists := obj.baseMap[name]
	return exists
}

func (obj *LogicalInterface) UpdateDone() {
	obj.modified.SetInt64(0)
	obj.baseMap = nil
}

func (obj *LogicalInterface) GetLogicalInterfaceVlanTag() int {
	return obj.logical_interface_vlan_tag
}

func (obj *LogicalInterface) SetLogicalInterfaceVlanTag(value int) {
	obj.logical_interface_vlan_tag = value
	obj.modified.SetBit(&obj.modified, logical_interface_logical_interface_vlan_tag, 1)
}

func (obj *LogicalInterface) GetLogicalInterfaceType() string {
	return obj.logical_interface_type
}

func (obj *LogicalInterface) SetLogicalInterfaceType(value string) {
	obj.logical_interface_type = value
	obj.modified.SetBit(&obj.modified, logical_interface_logical_interface_type, 1)
}

func (obj *LogicalInterface) GetIdPerms() IdPermsType {
	return obj.id_perms
}

func (obj *LogicalInterface) SetIdPerms(value *IdPermsType) {
	obj.id_perms = *value
	obj.modified.SetBit(&obj.modified, logical_interface_id_perms, 1)
}

func (obj *LogicalInterface) GetPerms2() PermType2 {
	return obj.perms2
}

func (obj *LogicalInterface) SetPerms2(value *PermType2) {
	obj.perms2 = *value
	obj.modified.SetBit(&obj.modified, logical_interface_perms2, 1)
}

func (obj *LogicalInterface) GetAnnotations() KeyValuePairs {
	return obj.annotations
}

func (obj *LogicalInterface) SetAnnotations(value *KeyValuePairs) {
	obj.annotations = *value
	obj.modified.SetBit(&obj.modified, logical_interface_annotations, 1)
}

func (obj *LogicalInterface) GetDisplayName() string {
	return obj.display_name
}

func (obj *LogicalInterface) SetDisplayName(value string) {
	obj.display_name = value
	obj.modified.SetBit(&obj.modified, logical_interface_display_name, 1)
}

func (obj *LogicalInterface) readVirtualMachineInterfaceRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(logical_interface_virtual_machine_interface_refs) == 0) {
		err := obj.GetField(obj, "virtual_machine_interface_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *LogicalInterface) GetVirtualMachineInterfaceRefs() (
	contrail.ReferenceList, error) {
	err := obj.readVirtualMachineInterfaceRefs()
	if err != nil {
		return nil, err
	}
	return obj.virtual_machine_interface_refs, nil
}

func (obj *LogicalInterface) AddVirtualMachineInterface(
	rhs *VirtualMachineInterface) error {
	err := obj.readVirtualMachineInterfaceRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(logical_interface_virtual_machine_interface_refs) == 0 {
		obj.storeReferenceBase("virtual-machine-interface", obj.virtual_machine_interface_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.virtual_machine_interface_refs = append(obj.virtual_machine_interface_refs, ref)
	obj.modified.SetBit(&obj.modified, logical_interface_virtual_machine_interface_refs, 1)
	return nil
}

func (obj *LogicalInterface) DeleteVirtualMachineInterface(uuid string) error {
	err := obj.readVirtualMachineInterfaceRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(logical_interface_virtual_machine_interface_refs) == 0 {
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
	obj.modified.SetBit(&obj.modified, logical_interface_virtual_machine_interface_refs, 1)
	return nil
}

func (obj *LogicalInterface) ClearVirtualMachineInterface() {
	if (obj.valid.Bit(logical_interface_virtual_machine_interface_refs) != 0) &&
		(obj.modified.Bit(logical_interface_virtual_machine_interface_refs) == 0) {
		obj.storeReferenceBase("virtual-machine-interface", obj.virtual_machine_interface_refs)
	}
	obj.virtual_machine_interface_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, logical_interface_virtual_machine_interface_refs, 1)
	obj.modified.SetBit(&obj.modified, logical_interface_virtual_machine_interface_refs, 1)
}

func (obj *LogicalInterface) SetVirtualMachineInterfaceList(
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

func (obj *LogicalInterface) readTagRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(logical_interface_tag_refs) == 0) {
		err := obj.GetField(obj, "tag_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *LogicalInterface) GetTagRefs() (
	contrail.ReferenceList, error) {
	err := obj.readTagRefs()
	if err != nil {
		return nil, err
	}
	return obj.tag_refs, nil
}

func (obj *LogicalInterface) AddTag(
	rhs *Tag) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(logical_interface_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.tag_refs = append(obj.tag_refs, ref)
	obj.modified.SetBit(&obj.modified, logical_interface_tag_refs, 1)
	return nil
}

func (obj *LogicalInterface) DeleteTag(uuid string) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(logical_interface_tag_refs) == 0 {
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
	obj.modified.SetBit(&obj.modified, logical_interface_tag_refs, 1)
	return nil
}

func (obj *LogicalInterface) ClearTag() {
	if (obj.valid.Bit(logical_interface_tag_refs) != 0) &&
		(obj.modified.Bit(logical_interface_tag_refs) == 0) {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}
	obj.tag_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, logical_interface_tag_refs, 1)
	obj.modified.SetBit(&obj.modified, logical_interface_tag_refs, 1)
}

func (obj *LogicalInterface) SetTagList(
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

func (obj *LogicalInterface) MarshalJSON() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalCommon(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(logical_interface_logical_interface_vlan_tag) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.logical_interface_vlan_tag)
		if err != nil {
			return nil, err
		}
		msg["logical_interface_vlan_tag"] = &value
	}

	if obj.modified.Bit(logical_interface_logical_interface_type) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.logical_interface_type)
		if err != nil {
			return nil, err
		}
		msg["logical_interface_type"] = &value
	}

	if obj.modified.Bit(logical_interface_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(logical_interface_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(logical_interface_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(logical_interface_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if len(obj.virtual_machine_interface_refs) > 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.virtual_machine_interface_refs)
		if err != nil {
			return nil, err
		}
		msg["virtual_machine_interface_refs"] = &value
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

func (obj *LogicalInterface) UnmarshalJSON(body []byte) error {
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
		case "logical_interface_vlan_tag":
			err = json.Unmarshal(value, &obj.logical_interface_vlan_tag)
			if err == nil {
				obj.valid.SetBit(&obj.valid, logical_interface_logical_interface_vlan_tag, 1)
			}
			break
		case "logical_interface_type":
			err = json.Unmarshal(value, &obj.logical_interface_type)
			if err == nil {
				obj.valid.SetBit(&obj.valid, logical_interface_logical_interface_type, 1)
			}
			break
		case "id_perms":
			err = json.Unmarshal(value, &obj.id_perms)
			if err == nil {
				obj.valid.SetBit(&obj.valid, logical_interface_id_perms, 1)
			}
			break
		case "perms2":
			err = json.Unmarshal(value, &obj.perms2)
			if err == nil {
				obj.valid.SetBit(&obj.valid, logical_interface_perms2, 1)
			}
			break
		case "annotations":
			err = json.Unmarshal(value, &obj.annotations)
			if err == nil {
				obj.valid.SetBit(&obj.valid, logical_interface_annotations, 1)
			}
			break
		case "display_name":
			err = json.Unmarshal(value, &obj.display_name)
			if err == nil {
				obj.valid.SetBit(&obj.valid, logical_interface_display_name, 1)
			}
			break
		case "virtual_machine_interface_refs":
			err = json.Unmarshal(value, &obj.virtual_machine_interface_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, logical_interface_virtual_machine_interface_refs, 1)
			}
			break
		case "tag_refs":
			err = json.Unmarshal(value, &obj.tag_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, logical_interface_tag_refs, 1)
			}
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *LogicalInterface) UpdateObject() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalId(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(logical_interface_logical_interface_vlan_tag) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.logical_interface_vlan_tag)
		if err != nil {
			return nil, err
		}
		msg["logical_interface_vlan_tag"] = &value
	}

	if obj.modified.Bit(logical_interface_logical_interface_type) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.logical_interface_type)
		if err != nil {
			return nil, err
		}
		msg["logical_interface_type"] = &value
	}

	if obj.modified.Bit(logical_interface_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(logical_interface_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(logical_interface_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(logical_interface_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if obj.modified.Bit(logical_interface_virtual_machine_interface_refs) != 0 {
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

	if obj.modified.Bit(logical_interface_tag_refs) != 0 {
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

func (obj *LogicalInterface) UpdateReferences() error {

	if (obj.modified.Bit(logical_interface_virtual_machine_interface_refs) != 0) &&
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

	if (obj.modified.Bit(logical_interface_tag_refs) != 0) &&
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

func LogicalInterfaceByName(c contrail.ApiClient, fqn string) (*LogicalInterface, error) {
	obj, err := c.FindByName("logical-interface", fqn)
	if err != nil {
		return nil, err
	}
	return obj.(*LogicalInterface), nil
}

func LogicalInterfaceByUuid(c contrail.ApiClient, uuid string) (*LogicalInterface, error) {
	obj, err := c.FindByUuid("logical-interface", uuid)
	if err != nil {
		return nil, err
	}
	return obj.(*LogicalInterface), nil
}
