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
	service_template_service_template_properties int = iota
	service_template_id_perms
	service_template_perms2
	service_template_annotations
	service_template_display_name
	service_template_service_appliance_set_refs
	service_template_tag_refs
	service_template_service_instance_back_refs
)

type ServiceTemplate struct {
	contrail.ObjectBase
	service_template_properties ServiceTemplateType
	id_perms                    IdPermsType
	perms2                      PermType2
	annotations                 KeyValuePairs
	display_name                string
	service_appliance_set_refs  contrail.ReferenceList
	tag_refs                    contrail.ReferenceList
	service_instance_back_refs  contrail.ReferenceList
	valid                       big.Int
	modified                    big.Int
	baseMap                     map[string]contrail.ReferenceList
}

func (obj *ServiceTemplate) GetType() string {
	return "service-template"
}

func (obj *ServiceTemplate) GetDefaultParent() []string {
	name := []string{"default-domain"}
	return name
}

func (obj *ServiceTemplate) GetDefaultParentType() string {
	return "domain"
}

func (obj *ServiceTemplate) SetName(name string) {
	obj.VSetName(obj, name)
}

func (obj *ServiceTemplate) SetParent(parent contrail.IObject) {
	obj.VSetParent(obj, parent)
}

func (obj *ServiceTemplate) storeReferenceBase(
	name string, refList contrail.ReferenceList) {
	if obj.baseMap == nil {
		obj.baseMap = make(map[string]contrail.ReferenceList)
	}
	refCopy := make(contrail.ReferenceList, len(refList))
	copy(refCopy, refList)
	obj.baseMap[name] = refCopy
}

func (obj *ServiceTemplate) hasReferenceBase(name string) bool {
	if obj.baseMap == nil {
		return false
	}
	_, exists := obj.baseMap[name]
	return exists
}

func (obj *ServiceTemplate) UpdateDone() {
	obj.modified.SetInt64(0)
	obj.baseMap = nil
}

func (obj *ServiceTemplate) GetServiceTemplateProperties() ServiceTemplateType {
	return obj.service_template_properties
}

func (obj *ServiceTemplate) SetServiceTemplateProperties(value *ServiceTemplateType) {
	obj.service_template_properties = *value
	obj.modified.SetBit(&obj.modified, service_template_service_template_properties, 1)
}

func (obj *ServiceTemplate) GetIdPerms() IdPermsType {
	return obj.id_perms
}

func (obj *ServiceTemplate) SetIdPerms(value *IdPermsType) {
	obj.id_perms = *value
	obj.modified.SetBit(&obj.modified, service_template_id_perms, 1)
}

func (obj *ServiceTemplate) GetPerms2() PermType2 {
	return obj.perms2
}

func (obj *ServiceTemplate) SetPerms2(value *PermType2) {
	obj.perms2 = *value
	obj.modified.SetBit(&obj.modified, service_template_perms2, 1)
}

func (obj *ServiceTemplate) GetAnnotations() KeyValuePairs {
	return obj.annotations
}

func (obj *ServiceTemplate) SetAnnotations(value *KeyValuePairs) {
	obj.annotations = *value
	obj.modified.SetBit(&obj.modified, service_template_annotations, 1)
}

func (obj *ServiceTemplate) GetDisplayName() string {
	return obj.display_name
}

func (obj *ServiceTemplate) SetDisplayName(value string) {
	obj.display_name = value
	obj.modified.SetBit(&obj.modified, service_template_display_name, 1)
}

func (obj *ServiceTemplate) readServiceApplianceSetRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(service_template_service_appliance_set_refs) == 0) {
		err := obj.GetField(obj, "service_appliance_set_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *ServiceTemplate) GetServiceApplianceSetRefs() (
	contrail.ReferenceList, error) {
	err := obj.readServiceApplianceSetRefs()
	if err != nil {
		return nil, err
	}
	return obj.service_appliance_set_refs, nil
}

func (obj *ServiceTemplate) AddServiceApplianceSet(
	rhs *ServiceApplianceSet) error {
	err := obj.readServiceApplianceSetRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(service_template_service_appliance_set_refs) == 0 {
		obj.storeReferenceBase("service-appliance-set", obj.service_appliance_set_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.service_appliance_set_refs = append(obj.service_appliance_set_refs, ref)
	obj.modified.SetBit(&obj.modified, service_template_service_appliance_set_refs, 1)
	return nil
}

func (obj *ServiceTemplate) DeleteServiceApplianceSet(uuid string) error {
	err := obj.readServiceApplianceSetRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(service_template_service_appliance_set_refs) == 0 {
		obj.storeReferenceBase("service-appliance-set", obj.service_appliance_set_refs)
	}

	for i, ref := range obj.service_appliance_set_refs {
		if ref.Uuid == uuid {
			obj.service_appliance_set_refs = append(
				obj.service_appliance_set_refs[:i],
				obj.service_appliance_set_refs[i+1:]...)
			break
		}
	}
	obj.modified.SetBit(&obj.modified, service_template_service_appliance_set_refs, 1)
	return nil
}

func (obj *ServiceTemplate) ClearServiceApplianceSet() {
	if (obj.valid.Bit(service_template_service_appliance_set_refs) != 0) &&
		(obj.modified.Bit(service_template_service_appliance_set_refs) == 0) {
		obj.storeReferenceBase("service-appliance-set", obj.service_appliance_set_refs)
	}
	obj.service_appliance_set_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, service_template_service_appliance_set_refs, 1)
	obj.modified.SetBit(&obj.modified, service_template_service_appliance_set_refs, 1)
}

func (obj *ServiceTemplate) SetServiceApplianceSetList(
	refList []contrail.ReferencePair) {
	obj.ClearServiceApplianceSet()
	obj.service_appliance_set_refs = make([]contrail.Reference, len(refList))
	for i, pair := range refList {
		obj.service_appliance_set_refs[i] = contrail.Reference{
			pair.Object.GetFQName(),
			pair.Object.GetUuid(),
			pair.Object.GetHref(),
			pair.Attribute,
		}
	}
}

func (obj *ServiceTemplate) readTagRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(service_template_tag_refs) == 0) {
		err := obj.GetField(obj, "tag_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *ServiceTemplate) GetTagRefs() (
	contrail.ReferenceList, error) {
	err := obj.readTagRefs()
	if err != nil {
		return nil, err
	}
	return obj.tag_refs, nil
}

func (obj *ServiceTemplate) AddTag(
	rhs *Tag) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(service_template_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.tag_refs = append(obj.tag_refs, ref)
	obj.modified.SetBit(&obj.modified, service_template_tag_refs, 1)
	return nil
}

func (obj *ServiceTemplate) DeleteTag(uuid string) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(service_template_tag_refs) == 0 {
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
	obj.modified.SetBit(&obj.modified, service_template_tag_refs, 1)
	return nil
}

func (obj *ServiceTemplate) ClearTag() {
	if (obj.valid.Bit(service_template_tag_refs) != 0) &&
		(obj.modified.Bit(service_template_tag_refs) == 0) {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}
	obj.tag_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, service_template_tag_refs, 1)
	obj.modified.SetBit(&obj.modified, service_template_tag_refs, 1)
}

func (obj *ServiceTemplate) SetTagList(
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

func (obj *ServiceTemplate) readServiceInstanceBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(service_template_service_instance_back_refs) == 0) {
		err := obj.GetField(obj, "service_instance_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *ServiceTemplate) GetServiceInstanceBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readServiceInstanceBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.service_instance_back_refs, nil
}

func (obj *ServiceTemplate) MarshalJSON() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalCommon(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(service_template_service_template_properties) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.service_template_properties)
		if err != nil {
			return nil, err
		}
		msg["service_template_properties"] = &value
	}

	if obj.modified.Bit(service_template_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(service_template_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(service_template_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(service_template_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if len(obj.service_appliance_set_refs) > 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.service_appliance_set_refs)
		if err != nil {
			return nil, err
		}
		msg["service_appliance_set_refs"] = &value
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

func (obj *ServiceTemplate) UnmarshalJSON(body []byte) error {
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
		case "service_template_properties":
			err = json.Unmarshal(value, &obj.service_template_properties)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_template_service_template_properties, 1)
			}
			break
		case "id_perms":
			err = json.Unmarshal(value, &obj.id_perms)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_template_id_perms, 1)
			}
			break
		case "perms2":
			err = json.Unmarshal(value, &obj.perms2)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_template_perms2, 1)
			}
			break
		case "annotations":
			err = json.Unmarshal(value, &obj.annotations)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_template_annotations, 1)
			}
			break
		case "display_name":
			err = json.Unmarshal(value, &obj.display_name)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_template_display_name, 1)
			}
			break
		case "service_appliance_set_refs":
			err = json.Unmarshal(value, &obj.service_appliance_set_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_template_service_appliance_set_refs, 1)
			}
			break
		case "tag_refs":
			err = json.Unmarshal(value, &obj.tag_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_template_tag_refs, 1)
			}
			break
		case "service_instance_back_refs":
			err = json.Unmarshal(value, &obj.service_instance_back_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_template_service_instance_back_refs, 1)
			}
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *ServiceTemplate) UpdateObject() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalId(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(service_template_service_template_properties) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.service_template_properties)
		if err != nil {
			return nil, err
		}
		msg["service_template_properties"] = &value
	}

	if obj.modified.Bit(service_template_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(service_template_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(service_template_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(service_template_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if obj.modified.Bit(service_template_service_appliance_set_refs) != 0 {
		if len(obj.service_appliance_set_refs) == 0 {
			var value json.RawMessage
			value, err := json.Marshal(
				make([]contrail.Reference, 0))
			if err != nil {
				return nil, err
			}
			msg["service_appliance_set_refs"] = &value
		} else if !obj.hasReferenceBase("service-appliance-set") {
			var value json.RawMessage
			value, err := json.Marshal(&obj.service_appliance_set_refs)
			if err != nil {
				return nil, err
			}
			msg["service_appliance_set_refs"] = &value
		}
	}

	if obj.modified.Bit(service_template_tag_refs) != 0 {
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

func (obj *ServiceTemplate) UpdateReferences() error {

	if (obj.modified.Bit(service_template_service_appliance_set_refs) != 0) &&
		len(obj.service_appliance_set_refs) > 0 &&
		obj.hasReferenceBase("service-appliance-set") {
		err := obj.UpdateReference(
			obj, "service-appliance-set",
			obj.service_appliance_set_refs,
			obj.baseMap["service-appliance-set"])
		if err != nil {
			return err
		}
	}

	if (obj.modified.Bit(service_template_tag_refs) != 0) &&
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

func ServiceTemplateByName(c contrail.ApiClient, fqn string) (*ServiceTemplate, error) {
	obj, err := c.FindByName("service-template", fqn)
	if err != nil {
		return nil, err
	}
	return obj.(*ServiceTemplate), nil
}

func ServiceTemplateByUuid(c contrail.ApiClient, uuid string) (*ServiceTemplate, error) {
	obj, err := c.FindByUuid("service-template", uuid)
	if err != nil {
		return nil, err
	}
	return obj.(*ServiceTemplate), nil
}
