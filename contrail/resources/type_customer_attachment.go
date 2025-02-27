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
	customer_attachment_id_perms int = iota
	customer_attachment_perms2
	customer_attachment_annotations
	customer_attachment_display_name
	customer_attachment_virtual_machine_interface_refs
	customer_attachment_floating_ip_refs
	customer_attachment_tag_refs
)

type CustomerAttachment struct {
	contrail.ObjectBase
	id_perms                       IdPermsType
	perms2                         PermType2
	annotations                    KeyValuePairs
	display_name                   string
	virtual_machine_interface_refs contrail.ReferenceList
	floating_ip_refs               contrail.ReferenceList
	tag_refs                       contrail.ReferenceList
	valid                          big.Int
	modified                       big.Int
	baseMap                        map[string]contrail.ReferenceList
}

func (obj *CustomerAttachment) GetType() string {
	return "customer-attachment"
}

func (obj *CustomerAttachment) GetDefaultParent() []string {
	name := []string{}
	return name
}

func (obj *CustomerAttachment) GetDefaultParentType() string {
	return ""
}

func (obj *CustomerAttachment) SetName(name string) {
	obj.VSetName(obj, name)
}

func (obj *CustomerAttachment) SetParent(parent contrail.IObject) {
	obj.VSetParent(obj, parent)
}

func (obj *CustomerAttachment) storeReferenceBase(
	name string, refList contrail.ReferenceList) {
	if obj.baseMap == nil {
		obj.baseMap = make(map[string]contrail.ReferenceList)
	}
	refCopy := make(contrail.ReferenceList, len(refList))
	copy(refCopy, refList)
	obj.baseMap[name] = refCopy
}

func (obj *CustomerAttachment) hasReferenceBase(name string) bool {
	if obj.baseMap == nil {
		return false
	}
	_, exists := obj.baseMap[name]
	return exists
}

func (obj *CustomerAttachment) UpdateDone() {
	obj.modified.SetInt64(0)
	obj.baseMap = nil
}

func (obj *CustomerAttachment) GetIdPerms() IdPermsType {
	return obj.id_perms
}

func (obj *CustomerAttachment) SetIdPerms(value *IdPermsType) {
	obj.id_perms = *value
	obj.modified.SetBit(&obj.modified, customer_attachment_id_perms, 1)
}

func (obj *CustomerAttachment) GetPerms2() PermType2 {
	return obj.perms2
}

func (obj *CustomerAttachment) SetPerms2(value *PermType2) {
	obj.perms2 = *value
	obj.modified.SetBit(&obj.modified, customer_attachment_perms2, 1)
}

func (obj *CustomerAttachment) GetAnnotations() KeyValuePairs {
	return obj.annotations
}

func (obj *CustomerAttachment) SetAnnotations(value *KeyValuePairs) {
	obj.annotations = *value
	obj.modified.SetBit(&obj.modified, customer_attachment_annotations, 1)
}

func (obj *CustomerAttachment) GetDisplayName() string {
	return obj.display_name
}

func (obj *CustomerAttachment) SetDisplayName(value string) {
	obj.display_name = value
	obj.modified.SetBit(&obj.modified, customer_attachment_display_name, 1)
}

func (obj *CustomerAttachment) readVirtualMachineInterfaceRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(customer_attachment_virtual_machine_interface_refs) == 0) {
		err := obj.GetField(obj, "virtual_machine_interface_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *CustomerAttachment) GetVirtualMachineInterfaceRefs() (
	contrail.ReferenceList, error) {
	err := obj.readVirtualMachineInterfaceRefs()
	if err != nil {
		return nil, err
	}
	return obj.virtual_machine_interface_refs, nil
}

func (obj *CustomerAttachment) AddVirtualMachineInterface(
	rhs *VirtualMachineInterface) error {
	err := obj.readVirtualMachineInterfaceRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(customer_attachment_virtual_machine_interface_refs) == 0 {
		obj.storeReferenceBase("virtual-machine-interface", obj.virtual_machine_interface_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.virtual_machine_interface_refs = append(obj.virtual_machine_interface_refs, ref)
	obj.modified.SetBit(&obj.modified, customer_attachment_virtual_machine_interface_refs, 1)
	return nil
}

func (obj *CustomerAttachment) DeleteVirtualMachineInterface(uuid string) error {
	err := obj.readVirtualMachineInterfaceRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(customer_attachment_virtual_machine_interface_refs) == 0 {
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
	obj.modified.SetBit(&obj.modified, customer_attachment_virtual_machine_interface_refs, 1)
	return nil
}

func (obj *CustomerAttachment) ClearVirtualMachineInterface() {
	if (obj.valid.Bit(customer_attachment_virtual_machine_interface_refs) != 0) &&
		(obj.modified.Bit(customer_attachment_virtual_machine_interface_refs) == 0) {
		obj.storeReferenceBase("virtual-machine-interface", obj.virtual_machine_interface_refs)
	}
	obj.virtual_machine_interface_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, customer_attachment_virtual_machine_interface_refs, 1)
	obj.modified.SetBit(&obj.modified, customer_attachment_virtual_machine_interface_refs, 1)
}

func (obj *CustomerAttachment) SetVirtualMachineInterfaceList(
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

func (obj *CustomerAttachment) readFloatingIpRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(customer_attachment_floating_ip_refs) == 0) {
		err := obj.GetField(obj, "floating_ip_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *CustomerAttachment) GetFloatingIpRefs() (
	contrail.ReferenceList, error) {
	err := obj.readFloatingIpRefs()
	if err != nil {
		return nil, err
	}
	return obj.floating_ip_refs, nil
}

func (obj *CustomerAttachment) AddFloatingIp(
	rhs *FloatingIp) error {
	err := obj.readFloatingIpRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(customer_attachment_floating_ip_refs) == 0 {
		obj.storeReferenceBase("floating-ip", obj.floating_ip_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.floating_ip_refs = append(obj.floating_ip_refs, ref)
	obj.modified.SetBit(&obj.modified, customer_attachment_floating_ip_refs, 1)
	return nil
}

func (obj *CustomerAttachment) DeleteFloatingIp(uuid string) error {
	err := obj.readFloatingIpRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(customer_attachment_floating_ip_refs) == 0 {
		obj.storeReferenceBase("floating-ip", obj.floating_ip_refs)
	}

	for i, ref := range obj.floating_ip_refs {
		if ref.Uuid == uuid {
			obj.floating_ip_refs = append(
				obj.floating_ip_refs[:i],
				obj.floating_ip_refs[i+1:]...)
			break
		}
	}
	obj.modified.SetBit(&obj.modified, customer_attachment_floating_ip_refs, 1)
	return nil
}

func (obj *CustomerAttachment) ClearFloatingIp() {
	if (obj.valid.Bit(customer_attachment_floating_ip_refs) != 0) &&
		(obj.modified.Bit(customer_attachment_floating_ip_refs) == 0) {
		obj.storeReferenceBase("floating-ip", obj.floating_ip_refs)
	}
	obj.floating_ip_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, customer_attachment_floating_ip_refs, 1)
	obj.modified.SetBit(&obj.modified, customer_attachment_floating_ip_refs, 1)
}

func (obj *CustomerAttachment) SetFloatingIpList(
	refList []contrail.ReferencePair) {
	obj.ClearFloatingIp()
	obj.floating_ip_refs = make([]contrail.Reference, len(refList))
	for i, pair := range refList {
		obj.floating_ip_refs[i] = contrail.Reference{
			pair.Object.GetFQName(),
			pair.Object.GetUuid(),
			pair.Object.GetHref(),
			pair.Attribute,
		}
	}
}

func (obj *CustomerAttachment) readTagRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(customer_attachment_tag_refs) == 0) {
		err := obj.GetField(obj, "tag_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *CustomerAttachment) GetTagRefs() (
	contrail.ReferenceList, error) {
	err := obj.readTagRefs()
	if err != nil {
		return nil, err
	}
	return obj.tag_refs, nil
}

func (obj *CustomerAttachment) AddTag(
	rhs *Tag) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(customer_attachment_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.tag_refs = append(obj.tag_refs, ref)
	obj.modified.SetBit(&obj.modified, customer_attachment_tag_refs, 1)
	return nil
}

func (obj *CustomerAttachment) DeleteTag(uuid string) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(customer_attachment_tag_refs) == 0 {
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
	obj.modified.SetBit(&obj.modified, customer_attachment_tag_refs, 1)
	return nil
}

func (obj *CustomerAttachment) ClearTag() {
	if (obj.valid.Bit(customer_attachment_tag_refs) != 0) &&
		(obj.modified.Bit(customer_attachment_tag_refs) == 0) {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}
	obj.tag_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, customer_attachment_tag_refs, 1)
	obj.modified.SetBit(&obj.modified, customer_attachment_tag_refs, 1)
}

func (obj *CustomerAttachment) SetTagList(
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

func (obj *CustomerAttachment) MarshalJSON() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalCommon(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(customer_attachment_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(customer_attachment_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(customer_attachment_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(customer_attachment_display_name) != 0 {
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

	if len(obj.floating_ip_refs) > 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.floating_ip_refs)
		if err != nil {
			return nil, err
		}
		msg["floating_ip_refs"] = &value
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

func (obj *CustomerAttachment) UnmarshalJSON(body []byte) error {
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
				obj.valid.SetBit(&obj.valid, customer_attachment_id_perms, 1)
			}
			break
		case "perms2":
			err = json.Unmarshal(value, &obj.perms2)
			if err == nil {
				obj.valid.SetBit(&obj.valid, customer_attachment_perms2, 1)
			}
			break
		case "annotations":
			err = json.Unmarshal(value, &obj.annotations)
			if err == nil {
				obj.valid.SetBit(&obj.valid, customer_attachment_annotations, 1)
			}
			break
		case "display_name":
			err = json.Unmarshal(value, &obj.display_name)
			if err == nil {
				obj.valid.SetBit(&obj.valid, customer_attachment_display_name, 1)
			}
			break
		case "virtual_machine_interface_refs":
			err = json.Unmarshal(value, &obj.virtual_machine_interface_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, customer_attachment_virtual_machine_interface_refs, 1)
			}
			break
		case "floating_ip_refs":
			err = json.Unmarshal(value, &obj.floating_ip_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, customer_attachment_floating_ip_refs, 1)
			}
			break
		case "tag_refs":
			err = json.Unmarshal(value, &obj.tag_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, customer_attachment_tag_refs, 1)
			}
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *CustomerAttachment) UpdateObject() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalId(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(customer_attachment_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(customer_attachment_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(customer_attachment_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(customer_attachment_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if obj.modified.Bit(customer_attachment_virtual_machine_interface_refs) != 0 {
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

	if obj.modified.Bit(customer_attachment_floating_ip_refs) != 0 {
		if len(obj.floating_ip_refs) == 0 {
			var value json.RawMessage
			value, err := json.Marshal(
				make([]contrail.Reference, 0))
			if err != nil {
				return nil, err
			}
			msg["floating_ip_refs"] = &value
		} else if !obj.hasReferenceBase("floating-ip") {
			var value json.RawMessage
			value, err := json.Marshal(&obj.floating_ip_refs)
			if err != nil {
				return nil, err
			}
			msg["floating_ip_refs"] = &value
		}
	}

	if obj.modified.Bit(customer_attachment_tag_refs) != 0 {
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

func (obj *CustomerAttachment) UpdateReferences() error {

	if (obj.modified.Bit(customer_attachment_virtual_machine_interface_refs) != 0) &&
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

	if (obj.modified.Bit(customer_attachment_floating_ip_refs) != 0) &&
		len(obj.floating_ip_refs) > 0 &&
		obj.hasReferenceBase("floating-ip") {
		err := obj.UpdateReference(
			obj, "floating-ip",
			obj.floating_ip_refs,
			obj.baseMap["floating-ip"])
		if err != nil {
			return err
		}
	}

	if (obj.modified.Bit(customer_attachment_tag_refs) != 0) &&
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

func CustomerAttachmentByName(c contrail.ApiClient, fqn string) (*CustomerAttachment, error) {
	obj, err := c.FindByName("customer-attachment", fqn)
	if err != nil {
		return nil, err
	}
	return obj.(*CustomerAttachment), nil
}

func CustomerAttachmentByUuid(c contrail.ApiClient, uuid string) (*CustomerAttachment, error) {
	obj, err := c.FindByUuid("customer-attachment", uuid)
	if err != nil {
		return nil, err
	}
	return obj.(*CustomerAttachment), nil
}
