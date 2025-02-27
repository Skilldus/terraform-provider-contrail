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
	loadbalancer_member_loadbalancer_member_properties int = iota
	loadbalancer_member_id_perms
	loadbalancer_member_perms2
	loadbalancer_member_annotations
	loadbalancer_member_display_name
	loadbalancer_member_tag_refs
)

type LoadbalancerMember struct {
	contrail.ObjectBase
	loadbalancer_member_properties LoadbalancerMemberType
	id_perms                       IdPermsType
	perms2                         PermType2
	annotations                    KeyValuePairs
	display_name                   string
	tag_refs                       contrail.ReferenceList
	valid                          big.Int
	modified                       big.Int
	baseMap                        map[string]contrail.ReferenceList
}

func (obj *LoadbalancerMember) GetType() string {
	return "loadbalancer-member"
}

func (obj *LoadbalancerMember) GetDefaultParent() []string {
	name := []string{"default-domain", "default-project", "default-loadbalancer-pool"}
	return name
}

func (obj *LoadbalancerMember) GetDefaultParentType() string {
	return "loadbalancer-pool"
}

func (obj *LoadbalancerMember) SetName(name string) {
	obj.VSetName(obj, name)
}

func (obj *LoadbalancerMember) SetParent(parent contrail.IObject) {
	obj.VSetParent(obj, parent)
}

func (obj *LoadbalancerMember) storeReferenceBase(
	name string, refList contrail.ReferenceList) {
	if obj.baseMap == nil {
		obj.baseMap = make(map[string]contrail.ReferenceList)
	}
	refCopy := make(contrail.ReferenceList, len(refList))
	copy(refCopy, refList)
	obj.baseMap[name] = refCopy
}

func (obj *LoadbalancerMember) hasReferenceBase(name string) bool {
	if obj.baseMap == nil {
		return false
	}
	_, exists := obj.baseMap[name]
	return exists
}

func (obj *LoadbalancerMember) UpdateDone() {
	obj.modified.SetInt64(0)
	obj.baseMap = nil
}

func (obj *LoadbalancerMember) GetLoadbalancerMemberProperties() LoadbalancerMemberType {
	return obj.loadbalancer_member_properties
}

func (obj *LoadbalancerMember) SetLoadbalancerMemberProperties(value *LoadbalancerMemberType) {
	obj.loadbalancer_member_properties = *value
	obj.modified.SetBit(&obj.modified, loadbalancer_member_loadbalancer_member_properties, 1)
}

func (obj *LoadbalancerMember) GetIdPerms() IdPermsType {
	return obj.id_perms
}

func (obj *LoadbalancerMember) SetIdPerms(value *IdPermsType) {
	obj.id_perms = *value
	obj.modified.SetBit(&obj.modified, loadbalancer_member_id_perms, 1)
}

func (obj *LoadbalancerMember) GetPerms2() PermType2 {
	return obj.perms2
}

func (obj *LoadbalancerMember) SetPerms2(value *PermType2) {
	obj.perms2 = *value
	obj.modified.SetBit(&obj.modified, loadbalancer_member_perms2, 1)
}

func (obj *LoadbalancerMember) GetAnnotations() KeyValuePairs {
	return obj.annotations
}

func (obj *LoadbalancerMember) SetAnnotations(value *KeyValuePairs) {
	obj.annotations = *value
	obj.modified.SetBit(&obj.modified, loadbalancer_member_annotations, 1)
}

func (obj *LoadbalancerMember) GetDisplayName() string {
	return obj.display_name
}

func (obj *LoadbalancerMember) SetDisplayName(value string) {
	obj.display_name = value
	obj.modified.SetBit(&obj.modified, loadbalancer_member_display_name, 1)
}

func (obj *LoadbalancerMember) readTagRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(loadbalancer_member_tag_refs) == 0) {
		err := obj.GetField(obj, "tag_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *LoadbalancerMember) GetTagRefs() (
	contrail.ReferenceList, error) {
	err := obj.readTagRefs()
	if err != nil {
		return nil, err
	}
	return obj.tag_refs, nil
}

func (obj *LoadbalancerMember) AddTag(
	rhs *Tag) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(loadbalancer_member_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.tag_refs = append(obj.tag_refs, ref)
	obj.modified.SetBit(&obj.modified, loadbalancer_member_tag_refs, 1)
	return nil
}

func (obj *LoadbalancerMember) DeleteTag(uuid string) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(loadbalancer_member_tag_refs) == 0 {
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
	obj.modified.SetBit(&obj.modified, loadbalancer_member_tag_refs, 1)
	return nil
}

func (obj *LoadbalancerMember) ClearTag() {
	if (obj.valid.Bit(loadbalancer_member_tag_refs) != 0) &&
		(obj.modified.Bit(loadbalancer_member_tag_refs) == 0) {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}
	obj.tag_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, loadbalancer_member_tag_refs, 1)
	obj.modified.SetBit(&obj.modified, loadbalancer_member_tag_refs, 1)
}

func (obj *LoadbalancerMember) SetTagList(
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

func (obj *LoadbalancerMember) MarshalJSON() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalCommon(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(loadbalancer_member_loadbalancer_member_properties) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.loadbalancer_member_properties)
		if err != nil {
			return nil, err
		}
		msg["loadbalancer_member_properties"] = &value
	}

	if obj.modified.Bit(loadbalancer_member_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(loadbalancer_member_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(loadbalancer_member_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(loadbalancer_member_display_name) != 0 {
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

func (obj *LoadbalancerMember) UnmarshalJSON(body []byte) error {
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
		case "loadbalancer_member_properties":
			err = json.Unmarshal(value, &obj.loadbalancer_member_properties)
			if err == nil {
				obj.valid.SetBit(&obj.valid, loadbalancer_member_loadbalancer_member_properties, 1)
			}
			break
		case "id_perms":
			err = json.Unmarshal(value, &obj.id_perms)
			if err == nil {
				obj.valid.SetBit(&obj.valid, loadbalancer_member_id_perms, 1)
			}
			break
		case "perms2":
			err = json.Unmarshal(value, &obj.perms2)
			if err == nil {
				obj.valid.SetBit(&obj.valid, loadbalancer_member_perms2, 1)
			}
			break
		case "annotations":
			err = json.Unmarshal(value, &obj.annotations)
			if err == nil {
				obj.valid.SetBit(&obj.valid, loadbalancer_member_annotations, 1)
			}
			break
		case "display_name":
			err = json.Unmarshal(value, &obj.display_name)
			if err == nil {
				obj.valid.SetBit(&obj.valid, loadbalancer_member_display_name, 1)
			}
			break
		case "tag_refs":
			err = json.Unmarshal(value, &obj.tag_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, loadbalancer_member_tag_refs, 1)
			}
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *LoadbalancerMember) UpdateObject() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalId(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(loadbalancer_member_loadbalancer_member_properties) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.loadbalancer_member_properties)
		if err != nil {
			return nil, err
		}
		msg["loadbalancer_member_properties"] = &value
	}

	if obj.modified.Bit(loadbalancer_member_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(loadbalancer_member_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(loadbalancer_member_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(loadbalancer_member_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if obj.modified.Bit(loadbalancer_member_tag_refs) != 0 {
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

func (obj *LoadbalancerMember) UpdateReferences() error {

	if (obj.modified.Bit(loadbalancer_member_tag_refs) != 0) &&
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

func LoadbalancerMemberByName(c contrail.ApiClient, fqn string) (*LoadbalancerMember, error) {
	obj, err := c.FindByName("loadbalancer-member", fqn)
	if err != nil {
		return nil, err
	}
	return obj.(*LoadbalancerMember), nil
}

func LoadbalancerMemberByUuid(c contrail.ApiClient, uuid string) (*LoadbalancerMember, error) {
	obj, err := c.FindByUuid("loadbalancer-member", uuid)
	if err != nil {
		return nil, err
	}
	return obj.(*LoadbalancerMember), nil
}
