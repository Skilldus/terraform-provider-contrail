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
	discovery_service_assignment_id_perms int = iota
	discovery_service_assignment_perms2
	discovery_service_assignment_annotations
	discovery_service_assignment_display_name
	discovery_service_assignment_dsa_rules
	discovery_service_assignment_tag_refs
)

type DiscoveryServiceAssignment struct {
	contrail.ObjectBase
	id_perms     IdPermsType
	perms2       PermType2
	annotations  KeyValuePairs
	display_name string
	dsa_rules    contrail.ReferenceList
	tag_refs     contrail.ReferenceList
	valid        big.Int
	modified     big.Int
	baseMap      map[string]contrail.ReferenceList
}

func (obj *DiscoveryServiceAssignment) GetType() string {
	return "discovery-service-assignment"
}

func (obj *DiscoveryServiceAssignment) GetDefaultParent() []string {
	name := []string{}
	return name
}

func (obj *DiscoveryServiceAssignment) GetDefaultParentType() string {
	return ""
}

func (obj *DiscoveryServiceAssignment) SetName(name string) {
	obj.VSetName(obj, name)
}

func (obj *DiscoveryServiceAssignment) SetParent(parent contrail.IObject) {
	obj.VSetParent(obj, parent)
}

func (obj *DiscoveryServiceAssignment) storeReferenceBase(
	name string, refList contrail.ReferenceList) {
	if obj.baseMap == nil {
		obj.baseMap = make(map[string]contrail.ReferenceList)
	}
	refCopy := make(contrail.ReferenceList, len(refList))
	copy(refCopy, refList)
	obj.baseMap[name] = refCopy
}

func (obj *DiscoveryServiceAssignment) hasReferenceBase(name string) bool {
	if obj.baseMap == nil {
		return false
	}
	_, exists := obj.baseMap[name]
	return exists
}

func (obj *DiscoveryServiceAssignment) UpdateDone() {
	obj.modified.SetInt64(0)
	obj.baseMap = nil
}

func (obj *DiscoveryServiceAssignment) GetIdPerms() IdPermsType {
	return obj.id_perms
}

func (obj *DiscoveryServiceAssignment) SetIdPerms(value *IdPermsType) {
	obj.id_perms = *value
	obj.modified.SetBit(&obj.modified, discovery_service_assignment_id_perms, 1)
}

func (obj *DiscoveryServiceAssignment) GetPerms2() PermType2 {
	return obj.perms2
}

func (obj *DiscoveryServiceAssignment) SetPerms2(value *PermType2) {
	obj.perms2 = *value
	obj.modified.SetBit(&obj.modified, discovery_service_assignment_perms2, 1)
}

func (obj *DiscoveryServiceAssignment) GetAnnotations() KeyValuePairs {
	return obj.annotations
}

func (obj *DiscoveryServiceAssignment) SetAnnotations(value *KeyValuePairs) {
	obj.annotations = *value
	obj.modified.SetBit(&obj.modified, discovery_service_assignment_annotations, 1)
}

func (obj *DiscoveryServiceAssignment) GetDisplayName() string {
	return obj.display_name
}

func (obj *DiscoveryServiceAssignment) SetDisplayName(value string) {
	obj.display_name = value
	obj.modified.SetBit(&obj.modified, discovery_service_assignment_display_name, 1)
}

func (obj *DiscoveryServiceAssignment) readDsaRules() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(discovery_service_assignment_dsa_rules) == 0) {
		err := obj.GetField(obj, "dsa_rules")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *DiscoveryServiceAssignment) GetDsaRules() (
	contrail.ReferenceList, error) {
	err := obj.readDsaRules()
	if err != nil {
		return nil, err
	}
	return obj.dsa_rules, nil
}

func (obj *DiscoveryServiceAssignment) readTagRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(discovery_service_assignment_tag_refs) == 0) {
		err := obj.GetField(obj, "tag_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *DiscoveryServiceAssignment) GetTagRefs() (
	contrail.ReferenceList, error) {
	err := obj.readTagRefs()
	if err != nil {
		return nil, err
	}
	return obj.tag_refs, nil
}

func (obj *DiscoveryServiceAssignment) AddTag(
	rhs *Tag) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(discovery_service_assignment_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.tag_refs = append(obj.tag_refs, ref)
	obj.modified.SetBit(&obj.modified, discovery_service_assignment_tag_refs, 1)
	return nil
}

func (obj *DiscoveryServiceAssignment) DeleteTag(uuid string) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(discovery_service_assignment_tag_refs) == 0 {
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
	obj.modified.SetBit(&obj.modified, discovery_service_assignment_tag_refs, 1)
	return nil
}

func (obj *DiscoveryServiceAssignment) ClearTag() {
	if (obj.valid.Bit(discovery_service_assignment_tag_refs) != 0) &&
		(obj.modified.Bit(discovery_service_assignment_tag_refs) == 0) {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}
	obj.tag_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, discovery_service_assignment_tag_refs, 1)
	obj.modified.SetBit(&obj.modified, discovery_service_assignment_tag_refs, 1)
}

func (obj *DiscoveryServiceAssignment) SetTagList(
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

func (obj *DiscoveryServiceAssignment) MarshalJSON() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalCommon(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(discovery_service_assignment_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(discovery_service_assignment_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(discovery_service_assignment_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(discovery_service_assignment_display_name) != 0 {
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

func (obj *DiscoveryServiceAssignment) UnmarshalJSON(body []byte) error {
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
				obj.valid.SetBit(&obj.valid, discovery_service_assignment_id_perms, 1)
			}
			break
		case "perms2":
			err = json.Unmarshal(value, &obj.perms2)
			if err == nil {
				obj.valid.SetBit(&obj.valid, discovery_service_assignment_perms2, 1)
			}
			break
		case "annotations":
			err = json.Unmarshal(value, &obj.annotations)
			if err == nil {
				obj.valid.SetBit(&obj.valid, discovery_service_assignment_annotations, 1)
			}
			break
		case "display_name":
			err = json.Unmarshal(value, &obj.display_name)
			if err == nil {
				obj.valid.SetBit(&obj.valid, discovery_service_assignment_display_name, 1)
			}
			break
		case "dsa_rules":
			err = json.Unmarshal(value, &obj.dsa_rules)
			if err == nil {
				obj.valid.SetBit(&obj.valid, discovery_service_assignment_dsa_rules, 1)
			}
			break
		case "tag_refs":
			err = json.Unmarshal(value, &obj.tag_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, discovery_service_assignment_tag_refs, 1)
			}
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *DiscoveryServiceAssignment) UpdateObject() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalId(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(discovery_service_assignment_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(discovery_service_assignment_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(discovery_service_assignment_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(discovery_service_assignment_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if obj.modified.Bit(discovery_service_assignment_tag_refs) != 0 {
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

func (obj *DiscoveryServiceAssignment) UpdateReferences() error {

	if (obj.modified.Bit(discovery_service_assignment_tag_refs) != 0) &&
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

func DiscoveryServiceAssignmentByName(c contrail.ApiClient, fqn string) (*DiscoveryServiceAssignment, error) {
	obj, err := c.FindByName("discovery-service-assignment", fqn)
	if err != nil {
		return nil, err
	}
	return obj.(*DiscoveryServiceAssignment), nil
}

func DiscoveryServiceAssignmentByUuid(c contrail.ApiClient, uuid string) (*DiscoveryServiceAssignment, error) {
	obj, err := c.FindByUuid("discovery-service-assignment", uuid)
	if err != nil {
		return nil, err
	}
	return obj.(*DiscoveryServiceAssignment), nil
}
