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
	security_group_security_group_id int = iota
	security_group_configured_security_group_id
	security_group_security_group_entries
	security_group_id_perms
	security_group_perms2
	security_group_annotations
	security_group_display_name
	security_group_access_control_lists
	security_group_tag_refs
	security_group_security_logging_object_back_refs
	security_group_virtual_machine_interface_back_refs
)

type SecurityGroup struct {
	contrail.ObjectBase
	security_group_id                   string
	configured_security_group_id        int
	security_group_entries              PolicyEntriesType
	id_perms                            IdPermsType
	perms2                              PermType2
	annotations                         KeyValuePairs
	display_name                        string
	access_control_lists                contrail.ReferenceList
	tag_refs                            contrail.ReferenceList
	security_logging_object_back_refs   contrail.ReferenceList
	virtual_machine_interface_back_refs contrail.ReferenceList
	valid                               big.Int
	modified                            big.Int
	baseMap                             map[string]contrail.ReferenceList
}

func (obj *SecurityGroup) GetType() string {
	return "security-group"
}

func (obj *SecurityGroup) GetDefaultParent() []string {
	name := []string{"default-domain", "default-project"}
	return name
}

func (obj *SecurityGroup) GetDefaultParentType() string {
	return "project"
}

func (obj *SecurityGroup) SetName(name string) {
	obj.VSetName(obj, name)
}

func (obj *SecurityGroup) SetParent(parent contrail.IObject) {
	obj.VSetParent(obj, parent)
}

func (obj *SecurityGroup) storeReferenceBase(
	name string, refList contrail.ReferenceList) {
	if obj.baseMap == nil {
		obj.baseMap = make(map[string]contrail.ReferenceList)
	}
	refCopy := make(contrail.ReferenceList, len(refList))
	copy(refCopy, refList)
	obj.baseMap[name] = refCopy
}

func (obj *SecurityGroup) hasReferenceBase(name string) bool {
	if obj.baseMap == nil {
		return false
	}
	_, exists := obj.baseMap[name]
	return exists
}

func (obj *SecurityGroup) UpdateDone() {
	obj.modified.SetInt64(0)
	obj.baseMap = nil
}

func (obj *SecurityGroup) GetSecurityGroupId() string {
	return obj.security_group_id
}

func (obj *SecurityGroup) SetSecurityGroupId(value string) {
	obj.security_group_id = value
	obj.modified.SetBit(&obj.modified, security_group_security_group_id, 1)
}

func (obj *SecurityGroup) GetConfiguredSecurityGroupId() int {
	return obj.configured_security_group_id
}

func (obj *SecurityGroup) SetConfiguredSecurityGroupId(value int) {
	obj.configured_security_group_id = value
	obj.modified.SetBit(&obj.modified, security_group_configured_security_group_id, 1)
}

func (obj *SecurityGroup) GetSecurityGroupEntries() PolicyEntriesType {
	return obj.security_group_entries
}

func (obj *SecurityGroup) SetSecurityGroupEntries(value *PolicyEntriesType) {
	obj.security_group_entries = *value
	obj.modified.SetBit(&obj.modified, security_group_security_group_entries, 1)
}

func (obj *SecurityGroup) GetIdPerms() IdPermsType {
	return obj.id_perms
}

func (obj *SecurityGroup) SetIdPerms(value *IdPermsType) {
	obj.id_perms = *value
	obj.modified.SetBit(&obj.modified, security_group_id_perms, 1)
}

func (obj *SecurityGroup) GetPerms2() PermType2 {
	return obj.perms2
}

func (obj *SecurityGroup) SetPerms2(value *PermType2) {
	obj.perms2 = *value
	obj.modified.SetBit(&obj.modified, security_group_perms2, 1)
}

func (obj *SecurityGroup) GetAnnotations() KeyValuePairs {
	return obj.annotations
}

func (obj *SecurityGroup) SetAnnotations(value *KeyValuePairs) {
	obj.annotations = *value
	obj.modified.SetBit(&obj.modified, security_group_annotations, 1)
}

func (obj *SecurityGroup) GetDisplayName() string {
	return obj.display_name
}

func (obj *SecurityGroup) SetDisplayName(value string) {
	obj.display_name = value
	obj.modified.SetBit(&obj.modified, security_group_display_name, 1)
}

func (obj *SecurityGroup) readAccessControlLists() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(security_group_access_control_lists) == 0) {
		err := obj.GetField(obj, "access_control_lists")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *SecurityGroup) GetAccessControlLists() (
	contrail.ReferenceList, error) {
	err := obj.readAccessControlLists()
	if err != nil {
		return nil, err
	}
	return obj.access_control_lists, nil
}

func (obj *SecurityGroup) readTagRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(security_group_tag_refs) == 0) {
		err := obj.GetField(obj, "tag_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *SecurityGroup) GetTagRefs() (
	contrail.ReferenceList, error) {
	err := obj.readTagRefs()
	if err != nil {
		return nil, err
	}
	return obj.tag_refs, nil
}

func (obj *SecurityGroup) AddTag(
	rhs *Tag) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(security_group_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.tag_refs = append(obj.tag_refs, ref)
	obj.modified.SetBit(&obj.modified, security_group_tag_refs, 1)
	return nil
}

func (obj *SecurityGroup) DeleteTag(uuid string) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(security_group_tag_refs) == 0 {
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
	obj.modified.SetBit(&obj.modified, security_group_tag_refs, 1)
	return nil
}

func (obj *SecurityGroup) ClearTag() {
	if (obj.valid.Bit(security_group_tag_refs) != 0) &&
		(obj.modified.Bit(security_group_tag_refs) == 0) {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}
	obj.tag_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, security_group_tag_refs, 1)
	obj.modified.SetBit(&obj.modified, security_group_tag_refs, 1)
}

func (obj *SecurityGroup) SetTagList(
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

func (obj *SecurityGroup) readSecurityLoggingObjectBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(security_group_security_logging_object_back_refs) == 0) {
		err := obj.GetField(obj, "security_logging_object_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *SecurityGroup) GetSecurityLoggingObjectBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readSecurityLoggingObjectBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.security_logging_object_back_refs, nil
}

func (obj *SecurityGroup) readVirtualMachineInterfaceBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(security_group_virtual_machine_interface_back_refs) == 0) {
		err := obj.GetField(obj, "virtual_machine_interface_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *SecurityGroup) GetVirtualMachineInterfaceBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readVirtualMachineInterfaceBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.virtual_machine_interface_back_refs, nil
}

func (obj *SecurityGroup) MarshalJSON() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalCommon(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(security_group_security_group_id) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.security_group_id)
		if err != nil {
			return nil, err
		}
		msg["security_group_id"] = &value
	}

	if obj.modified.Bit(security_group_configured_security_group_id) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.configured_security_group_id)
		if err != nil {
			return nil, err
		}
		msg["configured_security_group_id"] = &value
	}

	if obj.modified.Bit(security_group_security_group_entries) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.security_group_entries)
		if err != nil {
			return nil, err
		}
		msg["security_group_entries"] = &value
	}

	if obj.modified.Bit(security_group_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(security_group_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(security_group_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(security_group_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
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

func (obj *SecurityGroup) UnmarshalJSON(body []byte) error {
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
		case "security_group_id":
			err = json.Unmarshal(value, &obj.security_group_id)
			if err == nil {
				obj.valid.SetBit(&obj.valid, security_group_security_group_id, 1)
			}
			break
		case "configured_security_group_id":
			err = json.Unmarshal(value, &obj.configured_security_group_id)
			if err == nil {
				obj.valid.SetBit(&obj.valid, security_group_configured_security_group_id, 1)
			}
			break
		case "security_group_entries":
			err = json.Unmarshal(value, &obj.security_group_entries)
			if err == nil {
				obj.valid.SetBit(&obj.valid, security_group_security_group_entries, 1)
			}
			break
		case "id_perms":
			err = json.Unmarshal(value, &obj.id_perms)
			if err == nil {
				obj.valid.SetBit(&obj.valid, security_group_id_perms, 1)
			}
			break
		case "perms2":
			err = json.Unmarshal(value, &obj.perms2)
			if err == nil {
				obj.valid.SetBit(&obj.valid, security_group_perms2, 1)
			}
			break
		case "annotations":
			err = json.Unmarshal(value, &obj.annotations)
			if err == nil {
				obj.valid.SetBit(&obj.valid, security_group_annotations, 1)
			}
			break
		case "display_name":
			err = json.Unmarshal(value, &obj.display_name)
			if err == nil {
				obj.valid.SetBit(&obj.valid, security_group_display_name, 1)
			}
			break
		case "access_control_lists":
			err = json.Unmarshal(value, &obj.access_control_lists)
			if err == nil {
				obj.valid.SetBit(&obj.valid, security_group_access_control_lists, 1)
			}
			break
		case "tag_refs":
			err = json.Unmarshal(value, &obj.tag_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, security_group_tag_refs, 1)
			}
			break
		case "virtual_machine_interface_back_refs":
			err = json.Unmarshal(value, &obj.virtual_machine_interface_back_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, security_group_virtual_machine_interface_back_refs, 1)
			}
			break
		case "security_logging_object_back_refs":
			{
				type ReferenceElement struct {
					To   []string
					Uuid string
					Href string
					Attr SecurityLoggingObjectRuleListType
				}
				var array []ReferenceElement
				err = json.Unmarshal(value, &array)
				if err != nil {
					break
				}
				obj.valid.SetBit(&obj.valid, security_group_security_logging_object_back_refs, 1)
				obj.security_logging_object_back_refs = make(contrail.ReferenceList, 0)
				for _, element := range array {
					ref := contrail.Reference{
						element.To,
						element.Uuid,
						element.Href,
						element.Attr,
					}
					obj.security_logging_object_back_refs = append(obj.security_logging_object_back_refs, ref)
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

func (obj *SecurityGroup) UpdateObject() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalId(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(security_group_security_group_id) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.security_group_id)
		if err != nil {
			return nil, err
		}
		msg["security_group_id"] = &value
	}

	if obj.modified.Bit(security_group_configured_security_group_id) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.configured_security_group_id)
		if err != nil {
			return nil, err
		}
		msg["configured_security_group_id"] = &value
	}

	if obj.modified.Bit(security_group_security_group_entries) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.security_group_entries)
		if err != nil {
			return nil, err
		}
		msg["security_group_entries"] = &value
	}

	if obj.modified.Bit(security_group_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(security_group_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(security_group_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(security_group_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if obj.modified.Bit(security_group_tag_refs) != 0 {
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

func (obj *SecurityGroup) UpdateReferences() error {

	if (obj.modified.Bit(security_group_tag_refs) != 0) &&
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

func SecurityGroupByName(c contrail.ApiClient, fqn string) (*SecurityGroup, error) {
	obj, err := c.FindByName("security-group", fqn)
	if err != nil {
		return nil, err
	}
	return obj.(*SecurityGroup), nil
}

func SecurityGroupByUuid(c contrail.ApiClient, uuid string) (*SecurityGroup, error) {
	obj, err := c.FindByUuid("security-group", uuid)
	if err != nil {
		return nil, err
	}
	return obj.(*SecurityGroup), nil
}
