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
	access_control_list_access_control_list_entries int = iota
	access_control_list_access_control_list_hash
	access_control_list_id_perms
	access_control_list_perms2
	access_control_list_annotations
	access_control_list_display_name
	access_control_list_tag_refs
)

type AccessControlList struct {
	contrail.ObjectBase
	access_control_list_entries AclEntriesType
	access_control_list_hash    uint64
	id_perms                    IdPermsType
	perms2                      PermType2
	annotations                 KeyValuePairs
	display_name                string
	tag_refs                    contrail.ReferenceList
	valid                       big.Int
	modified                    big.Int
	baseMap                     map[string]contrail.ReferenceList
}

func (obj *AccessControlList) GetType() string {
	return "access-control-list"
}

func (obj *AccessControlList) GetDefaultParent() []string {
	name := []string{"default-domain", "default-project", "default-virtual-network"}
	return name
}

func (obj *AccessControlList) GetDefaultParentType() string {
	return "virtual-network"
}

func (obj *AccessControlList) SetName(name string) {
	obj.VSetName(obj, name)
}

func (obj *AccessControlList) SetParent(parent contrail.IObject) {
	obj.VSetParent(obj, parent)
}

func (obj *AccessControlList) storeReferenceBase(
	name string, refList contrail.ReferenceList) {
	if obj.baseMap == nil {
		obj.baseMap = make(map[string]contrail.ReferenceList)
	}
	refCopy := make(contrail.ReferenceList, len(refList))
	copy(refCopy, refList)
	obj.baseMap[name] = refCopy
}

func (obj *AccessControlList) hasReferenceBase(name string) bool {
	if obj.baseMap == nil {
		return false
	}
	_, exists := obj.baseMap[name]
	return exists
}

func (obj *AccessControlList) UpdateDone() {
	obj.modified.SetInt64(0)
	obj.baseMap = nil
}

func (obj *AccessControlList) GetAccessControlListEntries() AclEntriesType {
	return obj.access_control_list_entries
}

func (obj *AccessControlList) SetAccessControlListEntries(value *AclEntriesType) {
	obj.access_control_list_entries = *value
	obj.modified.SetBit(&obj.modified, access_control_list_access_control_list_entries, 1)
}

func (obj *AccessControlList) GetAccessControlListHash() uint64 {
	return obj.access_control_list_hash
}

func (obj *AccessControlList) SetAccessControlListHash(value uint64) {
	obj.access_control_list_hash = value
	obj.modified.SetBit(&obj.modified, access_control_list_access_control_list_hash, 1)
}

func (obj *AccessControlList) GetIdPerms() IdPermsType {
	return obj.id_perms
}

func (obj *AccessControlList) SetIdPerms(value *IdPermsType) {
	obj.id_perms = *value
	obj.modified.SetBit(&obj.modified, access_control_list_id_perms, 1)
}

func (obj *AccessControlList) GetPerms2() PermType2 {
	return obj.perms2
}

func (obj *AccessControlList) SetPerms2(value *PermType2) {
	obj.perms2 = *value
	obj.modified.SetBit(&obj.modified, access_control_list_perms2, 1)
}

func (obj *AccessControlList) GetAnnotations() KeyValuePairs {
	return obj.annotations
}

func (obj *AccessControlList) SetAnnotations(value *KeyValuePairs) {
	obj.annotations = *value
	obj.modified.SetBit(&obj.modified, access_control_list_annotations, 1)
}

func (obj *AccessControlList) GetDisplayName() string {
	return obj.display_name
}

func (obj *AccessControlList) SetDisplayName(value string) {
	obj.display_name = value
	obj.modified.SetBit(&obj.modified, access_control_list_display_name, 1)
}

func (obj *AccessControlList) readTagRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(access_control_list_tag_refs) == 0) {
		err := obj.GetField(obj, "tag_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *AccessControlList) GetTagRefs() (
	contrail.ReferenceList, error) {
	err := obj.readTagRefs()
	if err != nil {
		return nil, err
	}
	return obj.tag_refs, nil
}

func (obj *AccessControlList) AddTag(
	rhs *Tag) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(access_control_list_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.tag_refs = append(obj.tag_refs, ref)
	obj.modified.SetBit(&obj.modified, access_control_list_tag_refs, 1)
	return nil
}

func (obj *AccessControlList) DeleteTag(uuid string) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(access_control_list_tag_refs) == 0 {
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
	obj.modified.SetBit(&obj.modified, access_control_list_tag_refs, 1)
	return nil
}

func (obj *AccessControlList) ClearTag() {
	if (obj.valid.Bit(access_control_list_tag_refs) != 0) &&
		(obj.modified.Bit(access_control_list_tag_refs) == 0) {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}
	obj.tag_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, access_control_list_tag_refs, 1)
	obj.modified.SetBit(&obj.modified, access_control_list_tag_refs, 1)
}

func (obj *AccessControlList) SetTagList(
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

func (obj *AccessControlList) MarshalJSON() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalCommon(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(access_control_list_access_control_list_entries) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.access_control_list_entries)
		if err != nil {
			return nil, err
		}
		msg["access_control_list_entries"] = &value
	}

	if obj.modified.Bit(access_control_list_access_control_list_hash) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.access_control_list_hash)
		if err != nil {
			return nil, err
		}
		msg["access_control_list_hash"] = &value
	}

	if obj.modified.Bit(access_control_list_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(access_control_list_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(access_control_list_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(access_control_list_display_name) != 0 {
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

func (obj *AccessControlList) UnmarshalJSON(body []byte) error {
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
		case "access_control_list_entries":
			err = json.Unmarshal(value, &obj.access_control_list_entries)
			if err == nil {
				obj.valid.SetBit(&obj.valid, access_control_list_access_control_list_entries, 1)
			}
			break
		case "access_control_list_hash":
			err = json.Unmarshal(value, &obj.access_control_list_hash)
			if err == nil {
				obj.valid.SetBit(&obj.valid, access_control_list_access_control_list_hash, 1)
			}
			break
		case "id_perms":
			err = json.Unmarshal(value, &obj.id_perms)
			if err == nil {
				obj.valid.SetBit(&obj.valid, access_control_list_id_perms, 1)
			}
			break
		case "perms2":
			err = json.Unmarshal(value, &obj.perms2)
			if err == nil {
				obj.valid.SetBit(&obj.valid, access_control_list_perms2, 1)
			}
			break
		case "annotations":
			err = json.Unmarshal(value, &obj.annotations)
			if err == nil {
				obj.valid.SetBit(&obj.valid, access_control_list_annotations, 1)
			}
			break
		case "display_name":
			err = json.Unmarshal(value, &obj.display_name)
			if err == nil {
				obj.valid.SetBit(&obj.valid, access_control_list_display_name, 1)
			}
			break
		case "tag_refs":
			err = json.Unmarshal(value, &obj.tag_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, access_control_list_tag_refs, 1)
			}
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *AccessControlList) UpdateObject() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalId(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(access_control_list_access_control_list_entries) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.access_control_list_entries)
		if err != nil {
			return nil, err
		}
		msg["access_control_list_entries"] = &value
	}

	if obj.modified.Bit(access_control_list_access_control_list_hash) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.access_control_list_hash)
		if err != nil {
			return nil, err
		}
		msg["access_control_list_hash"] = &value
	}

	if obj.modified.Bit(access_control_list_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(access_control_list_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(access_control_list_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(access_control_list_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if obj.modified.Bit(access_control_list_tag_refs) != 0 {
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

func (obj *AccessControlList) UpdateReferences() error {

	if (obj.modified.Bit(access_control_list_tag_refs) != 0) &&
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

func AccessControlListByName(c contrail.ApiClient, fqn string) (*AccessControlList, error) {
	obj, err := c.FindByName("access-control-list", fqn)
	if err != nil {
		return nil, err
	}
	return obj.(*AccessControlList), nil
}

func AccessControlListByUuid(c contrail.ApiClient, uuid string) (*AccessControlList, error) {
	obj, err := c.FindByUuid("access-control-list", uuid)
	if err != nil {
		return nil, err
	}
	return obj.(*AccessControlList), nil
}
