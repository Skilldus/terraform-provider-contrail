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
	virtual_DNS_virtual_DNS_data int = iota
	virtual_DNS_id_perms
	virtual_DNS_perms2
	virtual_DNS_annotations
	virtual_DNS_display_name
	virtual_DNS_virtual_DNS_records
	virtual_DNS_tag_refs
	virtual_DNS_network_ipam_back_refs
)

type VirtualDns struct {
	contrail.ObjectBase
	virtual_DNS_data       VirtualDnsType
	id_perms               IdPermsType
	perms2                 PermType2
	annotations            KeyValuePairs
	display_name           string
	virtual_DNS_records    contrail.ReferenceList
	tag_refs               contrail.ReferenceList
	network_ipam_back_refs contrail.ReferenceList
	valid                  big.Int
	modified               big.Int
	baseMap                map[string]contrail.ReferenceList
}

func (obj *VirtualDns) GetType() string {
	return "virtual-DNS"
}

func (obj *VirtualDns) GetDefaultParent() []string {
	name := []string{"default-domain"}
	return name
}

func (obj *VirtualDns) GetDefaultParentType() string {
	return "domain"
}

func (obj *VirtualDns) SetName(name string) {
	obj.VSetName(obj, name)
}

func (obj *VirtualDns) SetParent(parent contrail.IObject) {
	obj.VSetParent(obj, parent)
}

func (obj *VirtualDns) storeReferenceBase(
	name string, refList contrail.ReferenceList) {
	if obj.baseMap == nil {
		obj.baseMap = make(map[string]contrail.ReferenceList)
	}
	refCopy := make(contrail.ReferenceList, len(refList))
	copy(refCopy, refList)
	obj.baseMap[name] = refCopy
}

func (obj *VirtualDns) hasReferenceBase(name string) bool {
	if obj.baseMap == nil {
		return false
	}
	_, exists := obj.baseMap[name]
	return exists
}

func (obj *VirtualDns) UpdateDone() {
	obj.modified.SetInt64(0)
	obj.baseMap = nil
}

func (obj *VirtualDns) GetVirtualDnsData() VirtualDnsType {
	return obj.virtual_DNS_data
}

func (obj *VirtualDns) SetVirtualDnsData(value *VirtualDnsType) {
	obj.virtual_DNS_data = *value
	obj.modified.SetBit(&obj.modified, virtual_DNS_virtual_DNS_data, 1)
}

func (obj *VirtualDns) GetIdPerms() IdPermsType {
	return obj.id_perms
}

func (obj *VirtualDns) SetIdPerms(value *IdPermsType) {
	obj.id_perms = *value
	obj.modified.SetBit(&obj.modified, virtual_DNS_id_perms, 1)
}

func (obj *VirtualDns) GetPerms2() PermType2 {
	return obj.perms2
}

func (obj *VirtualDns) SetPerms2(value *PermType2) {
	obj.perms2 = *value
	obj.modified.SetBit(&obj.modified, virtual_DNS_perms2, 1)
}

func (obj *VirtualDns) GetAnnotations() KeyValuePairs {
	return obj.annotations
}

func (obj *VirtualDns) SetAnnotations(value *KeyValuePairs) {
	obj.annotations = *value
	obj.modified.SetBit(&obj.modified, virtual_DNS_annotations, 1)
}

func (obj *VirtualDns) GetDisplayName() string {
	return obj.display_name
}

func (obj *VirtualDns) SetDisplayName(value string) {
	obj.display_name = value
	obj.modified.SetBit(&obj.modified, virtual_DNS_display_name, 1)
}

func (obj *VirtualDns) readVirtualDnsRecords() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(virtual_DNS_virtual_DNS_records) == 0) {
		err := obj.GetField(obj, "virtual_DNS_records")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *VirtualDns) GetVirtualDnsRecords() (
	contrail.ReferenceList, error) {
	err := obj.readVirtualDnsRecords()
	if err != nil {
		return nil, err
	}
	return obj.virtual_DNS_records, nil
}

func (obj *VirtualDns) readTagRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(virtual_DNS_tag_refs) == 0) {
		err := obj.GetField(obj, "tag_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *VirtualDns) GetTagRefs() (
	contrail.ReferenceList, error) {
	err := obj.readTagRefs()
	if err != nil {
		return nil, err
	}
	return obj.tag_refs, nil
}

func (obj *VirtualDns) AddTag(
	rhs *Tag) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(virtual_DNS_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.tag_refs = append(obj.tag_refs, ref)
	obj.modified.SetBit(&obj.modified, virtual_DNS_tag_refs, 1)
	return nil
}

func (obj *VirtualDns) DeleteTag(uuid string) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(virtual_DNS_tag_refs) == 0 {
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
	obj.modified.SetBit(&obj.modified, virtual_DNS_tag_refs, 1)
	return nil
}

func (obj *VirtualDns) ClearTag() {
	if (obj.valid.Bit(virtual_DNS_tag_refs) != 0) &&
		(obj.modified.Bit(virtual_DNS_tag_refs) == 0) {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}
	obj.tag_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, virtual_DNS_tag_refs, 1)
	obj.modified.SetBit(&obj.modified, virtual_DNS_tag_refs, 1)
}

func (obj *VirtualDns) SetTagList(
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

func (obj *VirtualDns) readNetworkIpamBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(virtual_DNS_network_ipam_back_refs) == 0) {
		err := obj.GetField(obj, "network_ipam_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *VirtualDns) GetNetworkIpamBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readNetworkIpamBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.network_ipam_back_refs, nil
}

func (obj *VirtualDns) MarshalJSON() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalCommon(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(virtual_DNS_virtual_DNS_data) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.virtual_DNS_data)
		if err != nil {
			return nil, err
		}
		msg["virtual_DNS_data"] = &value
	}

	if obj.modified.Bit(virtual_DNS_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(virtual_DNS_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(virtual_DNS_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(virtual_DNS_display_name) != 0 {
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

func (obj *VirtualDns) UnmarshalJSON(body []byte) error {
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
		case "virtual_DNS_data":
			err = json.Unmarshal(value, &obj.virtual_DNS_data)
			if err == nil {
				obj.valid.SetBit(&obj.valid, virtual_DNS_virtual_DNS_data, 1)
			}
			break
		case "id_perms":
			err = json.Unmarshal(value, &obj.id_perms)
			if err == nil {
				obj.valid.SetBit(&obj.valid, virtual_DNS_id_perms, 1)
			}
			break
		case "perms2":
			err = json.Unmarshal(value, &obj.perms2)
			if err == nil {
				obj.valid.SetBit(&obj.valid, virtual_DNS_perms2, 1)
			}
			break
		case "annotations":
			err = json.Unmarshal(value, &obj.annotations)
			if err == nil {
				obj.valid.SetBit(&obj.valid, virtual_DNS_annotations, 1)
			}
			break
		case "display_name":
			err = json.Unmarshal(value, &obj.display_name)
			if err == nil {
				obj.valid.SetBit(&obj.valid, virtual_DNS_display_name, 1)
			}
			break
		case "virtual_DNS_records":
			err = json.Unmarshal(value, &obj.virtual_DNS_records)
			if err == nil {
				obj.valid.SetBit(&obj.valid, virtual_DNS_virtual_DNS_records, 1)
			}
			break
		case "tag_refs":
			err = json.Unmarshal(value, &obj.tag_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, virtual_DNS_tag_refs, 1)
			}
			break
		case "network_ipam_back_refs":
			err = json.Unmarshal(value, &obj.network_ipam_back_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, virtual_DNS_network_ipam_back_refs, 1)
			}
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *VirtualDns) UpdateObject() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalId(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(virtual_DNS_virtual_DNS_data) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.virtual_DNS_data)
		if err != nil {
			return nil, err
		}
		msg["virtual_DNS_data"] = &value
	}

	if obj.modified.Bit(virtual_DNS_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(virtual_DNS_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(virtual_DNS_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(virtual_DNS_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if obj.modified.Bit(virtual_DNS_tag_refs) != 0 {
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

func (obj *VirtualDns) UpdateReferences() error {

	if (obj.modified.Bit(virtual_DNS_tag_refs) != 0) &&
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

func VirtualDnsByName(c contrail.ApiClient, fqn string) (*VirtualDns, error) {
	obj, err := c.FindByName("virtual-DNS", fqn)
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualDns), nil
}

func VirtualDnsByUuid(c contrail.ApiClient, uuid string) (*VirtualDns, error) {
	obj, err := c.FindByUuid("virtual-DNS", uuid)
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualDns), nil
}
