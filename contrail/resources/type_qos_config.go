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
	qos_config_qos_config_type int = iota
	qos_config_dscp_entries
	qos_config_vlan_priority_entries
	qos_config_mpls_exp_entries
	qos_config_default_forwarding_class_id
	qos_config_id_perms
	qos_config_perms2
	qos_config_annotations
	qos_config_display_name
	qos_config_global_system_config_refs
	qos_config_tag_refs
	qos_config_virtual_network_back_refs
	qos_config_virtual_machine_interface_back_refs
)

type QosConfig struct {
	contrail.ObjectBase
	qos_config_type                     string
	dscp_entries                        QosIdForwardingClassPairs
	vlan_priority_entries               QosIdForwardingClassPairs
	mpls_exp_entries                    QosIdForwardingClassPairs
	default_forwarding_class_id         int
	id_perms                            IdPermsType
	perms2                              PermType2
	annotations                         KeyValuePairs
	display_name                        string
	global_system_config_refs           contrail.ReferenceList
	tag_refs                            contrail.ReferenceList
	virtual_network_back_refs           contrail.ReferenceList
	virtual_machine_interface_back_refs contrail.ReferenceList
	valid                               big.Int
	modified                            big.Int
	baseMap                             map[string]contrail.ReferenceList
}

func (obj *QosConfig) GetType() string {
	return "qos-config"
}

func (obj *QosConfig) GetDefaultParent() []string {
	name := []string{"default-global-system-config", "default-global-qos-config"}
	return name
}

func (obj *QosConfig) GetDefaultParentType() string {
	return "global-qos-config"
}

func (obj *QosConfig) SetName(name string) {
	obj.VSetName(obj, name)
}

func (obj *QosConfig) SetParent(parent contrail.IObject) {
	obj.VSetParent(obj, parent)
}

func (obj *QosConfig) storeReferenceBase(
	name string, refList contrail.ReferenceList) {
	if obj.baseMap == nil {
		obj.baseMap = make(map[string]contrail.ReferenceList)
	}
	refCopy := make(contrail.ReferenceList, len(refList))
	copy(refCopy, refList)
	obj.baseMap[name] = refCopy
}

func (obj *QosConfig) hasReferenceBase(name string) bool {
	if obj.baseMap == nil {
		return false
	}
	_, exists := obj.baseMap[name]
	return exists
}

func (obj *QosConfig) UpdateDone() {
	obj.modified.SetInt64(0)
	obj.baseMap = nil
}

func (obj *QosConfig) GetQosConfigType() string {
	return obj.qos_config_type
}

func (obj *QosConfig) SetQosConfigType(value string) {
	obj.qos_config_type = value
	obj.modified.SetBit(&obj.modified, qos_config_qos_config_type, 1)
}

func (obj *QosConfig) GetDscpEntries() QosIdForwardingClassPairs {
	return obj.dscp_entries
}

func (obj *QosConfig) SetDscpEntries(value *QosIdForwardingClassPairs) {
	obj.dscp_entries = *value
	obj.modified.SetBit(&obj.modified, qos_config_dscp_entries, 1)
}

func (obj *QosConfig) GetVlanPriorityEntries() QosIdForwardingClassPairs {
	return obj.vlan_priority_entries
}

func (obj *QosConfig) SetVlanPriorityEntries(value *QosIdForwardingClassPairs) {
	obj.vlan_priority_entries = *value
	obj.modified.SetBit(&obj.modified, qos_config_vlan_priority_entries, 1)
}

func (obj *QosConfig) GetMplsExpEntries() QosIdForwardingClassPairs {
	return obj.mpls_exp_entries
}

func (obj *QosConfig) SetMplsExpEntries(value *QosIdForwardingClassPairs) {
	obj.mpls_exp_entries = *value
	obj.modified.SetBit(&obj.modified, qos_config_mpls_exp_entries, 1)
}

func (obj *QosConfig) GetDefaultForwardingClassId() int {
	return obj.default_forwarding_class_id
}

func (obj *QosConfig) SetDefaultForwardingClassId(value int) {
	obj.default_forwarding_class_id = value
	obj.modified.SetBit(&obj.modified, qos_config_default_forwarding_class_id, 1)
}

func (obj *QosConfig) GetIdPerms() IdPermsType {
	return obj.id_perms
}

func (obj *QosConfig) SetIdPerms(value *IdPermsType) {
	obj.id_perms = *value
	obj.modified.SetBit(&obj.modified, qos_config_id_perms, 1)
}

func (obj *QosConfig) GetPerms2() PermType2 {
	return obj.perms2
}

func (obj *QosConfig) SetPerms2(value *PermType2) {
	obj.perms2 = *value
	obj.modified.SetBit(&obj.modified, qos_config_perms2, 1)
}

func (obj *QosConfig) GetAnnotations() KeyValuePairs {
	return obj.annotations
}

func (obj *QosConfig) SetAnnotations(value *KeyValuePairs) {
	obj.annotations = *value
	obj.modified.SetBit(&obj.modified, qos_config_annotations, 1)
}

func (obj *QosConfig) GetDisplayName() string {
	return obj.display_name
}

func (obj *QosConfig) SetDisplayName(value string) {
	obj.display_name = value
	obj.modified.SetBit(&obj.modified, qos_config_display_name, 1)
}

func (obj *QosConfig) readGlobalSystemConfigRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(qos_config_global_system_config_refs) == 0) {
		err := obj.GetField(obj, "global_system_config_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *QosConfig) GetGlobalSystemConfigRefs() (
	contrail.ReferenceList, error) {
	err := obj.readGlobalSystemConfigRefs()
	if err != nil {
		return nil, err
	}
	return obj.global_system_config_refs, nil
}

func (obj *QosConfig) AddGlobalSystemConfig(
	rhs *GlobalSystemConfig) error {
	err := obj.readGlobalSystemConfigRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(qos_config_global_system_config_refs) == 0 {
		obj.storeReferenceBase("global-system-config", obj.global_system_config_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.global_system_config_refs = append(obj.global_system_config_refs, ref)
	obj.modified.SetBit(&obj.modified, qos_config_global_system_config_refs, 1)
	return nil
}

func (obj *QosConfig) DeleteGlobalSystemConfig(uuid string) error {
	err := obj.readGlobalSystemConfigRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(qos_config_global_system_config_refs) == 0 {
		obj.storeReferenceBase("global-system-config", obj.global_system_config_refs)
	}

	for i, ref := range obj.global_system_config_refs {
		if ref.Uuid == uuid {
			obj.global_system_config_refs = append(
				obj.global_system_config_refs[:i],
				obj.global_system_config_refs[i+1:]...)
			break
		}
	}
	obj.modified.SetBit(&obj.modified, qos_config_global_system_config_refs, 1)
	return nil
}

func (obj *QosConfig) ClearGlobalSystemConfig() {
	if (obj.valid.Bit(qos_config_global_system_config_refs) != 0) &&
		(obj.modified.Bit(qos_config_global_system_config_refs) == 0) {
		obj.storeReferenceBase("global-system-config", obj.global_system_config_refs)
	}
	obj.global_system_config_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, qos_config_global_system_config_refs, 1)
	obj.modified.SetBit(&obj.modified, qos_config_global_system_config_refs, 1)
}

func (obj *QosConfig) SetGlobalSystemConfigList(
	refList []contrail.ReferencePair) {
	obj.ClearGlobalSystemConfig()
	obj.global_system_config_refs = make([]contrail.Reference, len(refList))
	for i, pair := range refList {
		obj.global_system_config_refs[i] = contrail.Reference{
			pair.Object.GetFQName(),
			pair.Object.GetUuid(),
			pair.Object.GetHref(),
			pair.Attribute,
		}
	}
}

func (obj *QosConfig) readTagRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(qos_config_tag_refs) == 0) {
		err := obj.GetField(obj, "tag_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *QosConfig) GetTagRefs() (
	contrail.ReferenceList, error) {
	err := obj.readTagRefs()
	if err != nil {
		return nil, err
	}
	return obj.tag_refs, nil
}

func (obj *QosConfig) AddTag(
	rhs *Tag) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(qos_config_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.tag_refs = append(obj.tag_refs, ref)
	obj.modified.SetBit(&obj.modified, qos_config_tag_refs, 1)
	return nil
}

func (obj *QosConfig) DeleteTag(uuid string) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(qos_config_tag_refs) == 0 {
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
	obj.modified.SetBit(&obj.modified, qos_config_tag_refs, 1)
	return nil
}

func (obj *QosConfig) ClearTag() {
	if (obj.valid.Bit(qos_config_tag_refs) != 0) &&
		(obj.modified.Bit(qos_config_tag_refs) == 0) {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}
	obj.tag_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, qos_config_tag_refs, 1)
	obj.modified.SetBit(&obj.modified, qos_config_tag_refs, 1)
}

func (obj *QosConfig) SetTagList(
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

func (obj *QosConfig) readVirtualNetworkBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(qos_config_virtual_network_back_refs) == 0) {
		err := obj.GetField(obj, "virtual_network_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *QosConfig) GetVirtualNetworkBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readVirtualNetworkBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.virtual_network_back_refs, nil
}

func (obj *QosConfig) readVirtualMachineInterfaceBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(qos_config_virtual_machine_interface_back_refs) == 0) {
		err := obj.GetField(obj, "virtual_machine_interface_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *QosConfig) GetVirtualMachineInterfaceBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readVirtualMachineInterfaceBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.virtual_machine_interface_back_refs, nil
}

func (obj *QosConfig) MarshalJSON() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalCommon(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(qos_config_qos_config_type) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.qos_config_type)
		if err != nil {
			return nil, err
		}
		msg["qos_config_type"] = &value
	}

	if obj.modified.Bit(qos_config_dscp_entries) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.dscp_entries)
		if err != nil {
			return nil, err
		}
		msg["dscp_entries"] = &value
	}

	if obj.modified.Bit(qos_config_vlan_priority_entries) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.vlan_priority_entries)
		if err != nil {
			return nil, err
		}
		msg["vlan_priority_entries"] = &value
	}

	if obj.modified.Bit(qos_config_mpls_exp_entries) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.mpls_exp_entries)
		if err != nil {
			return nil, err
		}
		msg["mpls_exp_entries"] = &value
	}

	if obj.modified.Bit(qos_config_default_forwarding_class_id) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.default_forwarding_class_id)
		if err != nil {
			return nil, err
		}
		msg["default_forwarding_class_id"] = &value
	}

	if obj.modified.Bit(qos_config_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(qos_config_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(qos_config_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(qos_config_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if len(obj.global_system_config_refs) > 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.global_system_config_refs)
		if err != nil {
			return nil, err
		}
		msg["global_system_config_refs"] = &value
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

func (obj *QosConfig) UnmarshalJSON(body []byte) error {
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
		case "qos_config_type":
			err = json.Unmarshal(value, &obj.qos_config_type)
			if err == nil {
				obj.valid.SetBit(&obj.valid, qos_config_qos_config_type, 1)
			}
			break
		case "dscp_entries":
			err = json.Unmarshal(value, &obj.dscp_entries)
			if err == nil {
				obj.valid.SetBit(&obj.valid, qos_config_dscp_entries, 1)
			}
			break
		case "vlan_priority_entries":
			err = json.Unmarshal(value, &obj.vlan_priority_entries)
			if err == nil {
				obj.valid.SetBit(&obj.valid, qos_config_vlan_priority_entries, 1)
			}
			break
		case "mpls_exp_entries":
			err = json.Unmarshal(value, &obj.mpls_exp_entries)
			if err == nil {
				obj.valid.SetBit(&obj.valid, qos_config_mpls_exp_entries, 1)
			}
			break
		case "default_forwarding_class_id":
			err = json.Unmarshal(value, &obj.default_forwarding_class_id)
			if err == nil {
				obj.valid.SetBit(&obj.valid, qos_config_default_forwarding_class_id, 1)
			}
			break
		case "id_perms":
			err = json.Unmarshal(value, &obj.id_perms)
			if err == nil {
				obj.valid.SetBit(&obj.valid, qos_config_id_perms, 1)
			}
			break
		case "perms2":
			err = json.Unmarshal(value, &obj.perms2)
			if err == nil {
				obj.valid.SetBit(&obj.valid, qos_config_perms2, 1)
			}
			break
		case "annotations":
			err = json.Unmarshal(value, &obj.annotations)
			if err == nil {
				obj.valid.SetBit(&obj.valid, qos_config_annotations, 1)
			}
			break
		case "display_name":
			err = json.Unmarshal(value, &obj.display_name)
			if err == nil {
				obj.valid.SetBit(&obj.valid, qos_config_display_name, 1)
			}
			break
		case "global_system_config_refs":
			err = json.Unmarshal(value, &obj.global_system_config_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, qos_config_global_system_config_refs, 1)
			}
			break
		case "tag_refs":
			err = json.Unmarshal(value, &obj.tag_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, qos_config_tag_refs, 1)
			}
			break
		case "virtual_network_back_refs":
			err = json.Unmarshal(value, &obj.virtual_network_back_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, qos_config_virtual_network_back_refs, 1)
			}
			break
		case "virtual_machine_interface_back_refs":
			err = json.Unmarshal(value, &obj.virtual_machine_interface_back_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, qos_config_virtual_machine_interface_back_refs, 1)
			}
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *QosConfig) UpdateObject() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalId(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(qos_config_qos_config_type) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.qos_config_type)
		if err != nil {
			return nil, err
		}
		msg["qos_config_type"] = &value
	}

	if obj.modified.Bit(qos_config_dscp_entries) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.dscp_entries)
		if err != nil {
			return nil, err
		}
		msg["dscp_entries"] = &value
	}

	if obj.modified.Bit(qos_config_vlan_priority_entries) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.vlan_priority_entries)
		if err != nil {
			return nil, err
		}
		msg["vlan_priority_entries"] = &value
	}

	if obj.modified.Bit(qos_config_mpls_exp_entries) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.mpls_exp_entries)
		if err != nil {
			return nil, err
		}
		msg["mpls_exp_entries"] = &value
	}

	if obj.modified.Bit(qos_config_default_forwarding_class_id) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.default_forwarding_class_id)
		if err != nil {
			return nil, err
		}
		msg["default_forwarding_class_id"] = &value
	}

	if obj.modified.Bit(qos_config_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(qos_config_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(qos_config_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(qos_config_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if obj.modified.Bit(qos_config_global_system_config_refs) != 0 {
		if len(obj.global_system_config_refs) == 0 {
			var value json.RawMessage
			value, err := json.Marshal(
				make([]contrail.Reference, 0))
			if err != nil {
				return nil, err
			}
			msg["global_system_config_refs"] = &value
		} else if !obj.hasReferenceBase("global-system-config") {
			var value json.RawMessage
			value, err := json.Marshal(&obj.global_system_config_refs)
			if err != nil {
				return nil, err
			}
			msg["global_system_config_refs"] = &value
		}
	}

	if obj.modified.Bit(qos_config_tag_refs) != 0 {
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

func (obj *QosConfig) UpdateReferences() error {

	if (obj.modified.Bit(qos_config_global_system_config_refs) != 0) &&
		len(obj.global_system_config_refs) > 0 &&
		obj.hasReferenceBase("global-system-config") {
		err := obj.UpdateReference(
			obj, "global-system-config",
			obj.global_system_config_refs,
			obj.baseMap["global-system-config"])
		if err != nil {
			return err
		}
	}

	if (obj.modified.Bit(qos_config_tag_refs) != 0) &&
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

func QosConfigByName(c contrail.ApiClient, fqn string) (*QosConfig, error) {
	obj, err := c.FindByName("qos-config", fqn)
	if err != nil {
		return nil, err
	}
	return obj.(*QosConfig), nil
}

func QosConfigByUuid(c contrail.ApiClient, uuid string) (*QosConfig, error) {
	obj, err := c.FindByUuid("qos-config", uuid)
	if err != nil {
		return nil, err
	}
	return obj.(*QosConfig), nil
}
