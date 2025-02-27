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
	qos_queue_min_bandwidth int = iota
	qos_queue_max_bandwidth
	qos_queue_qos_queue_identifier
	qos_queue_id_perms
	qos_queue_perms2
	qos_queue_annotations
	qos_queue_display_name
	qos_queue_tag_refs
	qos_queue_forwarding_class_back_refs
)

type QosQueue struct {
	contrail.ObjectBase
	min_bandwidth              int
	max_bandwidth              int
	qos_queue_identifier       int
	id_perms                   IdPermsType
	perms2                     PermType2
	annotations                KeyValuePairs
	display_name               string
	tag_refs                   contrail.ReferenceList
	forwarding_class_back_refs contrail.ReferenceList
	valid                      big.Int
	modified                   big.Int
	baseMap                    map[string]contrail.ReferenceList
}

func (obj *QosQueue) GetType() string {
	return "qos-queue"
}

func (obj *QosQueue) GetDefaultParent() []string {
	name := []string{"default-global-system-config", "default-global-qos-config"}
	return name
}

func (obj *QosQueue) GetDefaultParentType() string {
	return "global-qos-config"
}

func (obj *QosQueue) SetName(name string) {
	obj.VSetName(obj, name)
}

func (obj *QosQueue) SetParent(parent contrail.IObject) {
	obj.VSetParent(obj, parent)
}

func (obj *QosQueue) storeReferenceBase(
	name string, refList contrail.ReferenceList) {
	if obj.baseMap == nil {
		obj.baseMap = make(map[string]contrail.ReferenceList)
	}
	refCopy := make(contrail.ReferenceList, len(refList))
	copy(refCopy, refList)
	obj.baseMap[name] = refCopy
}

func (obj *QosQueue) hasReferenceBase(name string) bool {
	if obj.baseMap == nil {
		return false
	}
	_, exists := obj.baseMap[name]
	return exists
}

func (obj *QosQueue) UpdateDone() {
	obj.modified.SetInt64(0)
	obj.baseMap = nil
}

func (obj *QosQueue) GetMinBandwidth() int {
	return obj.min_bandwidth
}

func (obj *QosQueue) SetMinBandwidth(value int) {
	obj.min_bandwidth = value
	obj.modified.SetBit(&obj.modified, qos_queue_min_bandwidth, 1)
}

func (obj *QosQueue) GetMaxBandwidth() int {
	return obj.max_bandwidth
}

func (obj *QosQueue) SetMaxBandwidth(value int) {
	obj.max_bandwidth = value
	obj.modified.SetBit(&obj.modified, qos_queue_max_bandwidth, 1)
}

func (obj *QosQueue) GetQosQueueIdentifier() int {
	return obj.qos_queue_identifier
}

func (obj *QosQueue) SetQosQueueIdentifier(value int) {
	obj.qos_queue_identifier = value
	obj.modified.SetBit(&obj.modified, qos_queue_qos_queue_identifier, 1)
}

func (obj *QosQueue) GetIdPerms() IdPermsType {
	return obj.id_perms
}

func (obj *QosQueue) SetIdPerms(value *IdPermsType) {
	obj.id_perms = *value
	obj.modified.SetBit(&obj.modified, qos_queue_id_perms, 1)
}

func (obj *QosQueue) GetPerms2() PermType2 {
	return obj.perms2
}

func (obj *QosQueue) SetPerms2(value *PermType2) {
	obj.perms2 = *value
	obj.modified.SetBit(&obj.modified, qos_queue_perms2, 1)
}

func (obj *QosQueue) GetAnnotations() KeyValuePairs {
	return obj.annotations
}

func (obj *QosQueue) SetAnnotations(value *KeyValuePairs) {
	obj.annotations = *value
	obj.modified.SetBit(&obj.modified, qos_queue_annotations, 1)
}

func (obj *QosQueue) GetDisplayName() string {
	return obj.display_name
}

func (obj *QosQueue) SetDisplayName(value string) {
	obj.display_name = value
	obj.modified.SetBit(&obj.modified, qos_queue_display_name, 1)
}

func (obj *QosQueue) readTagRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(qos_queue_tag_refs) == 0) {
		err := obj.GetField(obj, "tag_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *QosQueue) GetTagRefs() (
	contrail.ReferenceList, error) {
	err := obj.readTagRefs()
	if err != nil {
		return nil, err
	}
	return obj.tag_refs, nil
}

func (obj *QosQueue) AddTag(
	rhs *Tag) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(qos_queue_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.tag_refs = append(obj.tag_refs, ref)
	obj.modified.SetBit(&obj.modified, qos_queue_tag_refs, 1)
	return nil
}

func (obj *QosQueue) DeleteTag(uuid string) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(qos_queue_tag_refs) == 0 {
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
	obj.modified.SetBit(&obj.modified, qos_queue_tag_refs, 1)
	return nil
}

func (obj *QosQueue) ClearTag() {
	if (obj.valid.Bit(qos_queue_tag_refs) != 0) &&
		(obj.modified.Bit(qos_queue_tag_refs) == 0) {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}
	obj.tag_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, qos_queue_tag_refs, 1)
	obj.modified.SetBit(&obj.modified, qos_queue_tag_refs, 1)
}

func (obj *QosQueue) SetTagList(
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

func (obj *QosQueue) readForwardingClassBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(qos_queue_forwarding_class_back_refs) == 0) {
		err := obj.GetField(obj, "forwarding_class_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *QosQueue) GetForwardingClassBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readForwardingClassBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.forwarding_class_back_refs, nil
}

func (obj *QosQueue) MarshalJSON() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalCommon(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(qos_queue_min_bandwidth) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.min_bandwidth)
		if err != nil {
			return nil, err
		}
		msg["min_bandwidth"] = &value
	}

	if obj.modified.Bit(qos_queue_max_bandwidth) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.max_bandwidth)
		if err != nil {
			return nil, err
		}
		msg["max_bandwidth"] = &value
	}

	if obj.modified.Bit(qos_queue_qos_queue_identifier) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.qos_queue_identifier)
		if err != nil {
			return nil, err
		}
		msg["qos_queue_identifier"] = &value
	}

	if obj.modified.Bit(qos_queue_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(qos_queue_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(qos_queue_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(qos_queue_display_name) != 0 {
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

func (obj *QosQueue) UnmarshalJSON(body []byte) error {
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
		case "min_bandwidth":
			err = json.Unmarshal(value, &obj.min_bandwidth)
			if err == nil {
				obj.valid.SetBit(&obj.valid, qos_queue_min_bandwidth, 1)
			}
			break
		case "max_bandwidth":
			err = json.Unmarshal(value, &obj.max_bandwidth)
			if err == nil {
				obj.valid.SetBit(&obj.valid, qos_queue_max_bandwidth, 1)
			}
			break
		case "qos_queue_identifier":
			err = json.Unmarshal(value, &obj.qos_queue_identifier)
			if err == nil {
				obj.valid.SetBit(&obj.valid, qos_queue_qos_queue_identifier, 1)
			}
			break
		case "id_perms":
			err = json.Unmarshal(value, &obj.id_perms)
			if err == nil {
				obj.valid.SetBit(&obj.valid, qos_queue_id_perms, 1)
			}
			break
		case "perms2":
			err = json.Unmarshal(value, &obj.perms2)
			if err == nil {
				obj.valid.SetBit(&obj.valid, qos_queue_perms2, 1)
			}
			break
		case "annotations":
			err = json.Unmarshal(value, &obj.annotations)
			if err == nil {
				obj.valid.SetBit(&obj.valid, qos_queue_annotations, 1)
			}
			break
		case "display_name":
			err = json.Unmarshal(value, &obj.display_name)
			if err == nil {
				obj.valid.SetBit(&obj.valid, qos_queue_display_name, 1)
			}
			break
		case "tag_refs":
			err = json.Unmarshal(value, &obj.tag_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, qos_queue_tag_refs, 1)
			}
			break
		case "forwarding_class_back_refs":
			err = json.Unmarshal(value, &obj.forwarding_class_back_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, qos_queue_forwarding_class_back_refs, 1)
			}
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *QosQueue) UpdateObject() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalId(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(qos_queue_min_bandwidth) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.min_bandwidth)
		if err != nil {
			return nil, err
		}
		msg["min_bandwidth"] = &value
	}

	if obj.modified.Bit(qos_queue_max_bandwidth) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.max_bandwidth)
		if err != nil {
			return nil, err
		}
		msg["max_bandwidth"] = &value
	}

	if obj.modified.Bit(qos_queue_qos_queue_identifier) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.qos_queue_identifier)
		if err != nil {
			return nil, err
		}
		msg["qos_queue_identifier"] = &value
	}

	if obj.modified.Bit(qos_queue_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(qos_queue_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(qos_queue_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(qos_queue_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if obj.modified.Bit(qos_queue_tag_refs) != 0 {
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

func (obj *QosQueue) UpdateReferences() error {

	if (obj.modified.Bit(qos_queue_tag_refs) != 0) &&
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

func QosQueueByName(c contrail.ApiClient, fqn string) (*QosQueue, error) {
	obj, err := c.FindByName("qos-queue", fqn)
	if err != nil {
		return nil, err
	}
	return obj.(*QosQueue), nil
}

func QosQueueByUuid(c contrail.ApiClient, uuid string) (*QosQueue, error) {
	obj, err := c.FindByUuid("qos-queue", uuid)
	if err != nil {
		return nil, err
	}
	return obj.(*QosQueue), nil
}
