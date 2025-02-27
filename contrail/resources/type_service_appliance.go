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
	service_appliance_service_appliance_user_credentials int = iota
	service_appliance_service_appliance_ip_address
	service_appliance_service_appliance_properties
	service_appliance_id_perms
	service_appliance_perms2
	service_appliance_annotations
	service_appliance_display_name
	service_appliance_physical_interface_refs
	service_appliance_tag_refs
)

type ServiceAppliance struct {
	contrail.ObjectBase
	service_appliance_user_credentials UserCredentials
	service_appliance_ip_address       string
	service_appliance_properties       KeyValuePairs
	id_perms                           IdPermsType
	perms2                             PermType2
	annotations                        KeyValuePairs
	display_name                       string
	physical_interface_refs            contrail.ReferenceList
	tag_refs                           contrail.ReferenceList
	valid                              big.Int
	modified                           big.Int
	baseMap                            map[string]contrail.ReferenceList
}

func (obj *ServiceAppliance) GetType() string {
	return "service-appliance"
}

func (obj *ServiceAppliance) GetDefaultParent() []string {
	name := []string{"default-global-system-config", "default-service-appliance-set"}
	return name
}

func (obj *ServiceAppliance) GetDefaultParentType() string {
	return "service-appliance-set"
}

func (obj *ServiceAppliance) SetName(name string) {
	obj.VSetName(obj, name)
}

func (obj *ServiceAppliance) SetParent(parent contrail.IObject) {
	obj.VSetParent(obj, parent)
}

func (obj *ServiceAppliance) storeReferenceBase(
	name string, refList contrail.ReferenceList) {
	if obj.baseMap == nil {
		obj.baseMap = make(map[string]contrail.ReferenceList)
	}
	refCopy := make(contrail.ReferenceList, len(refList))
	copy(refCopy, refList)
	obj.baseMap[name] = refCopy
}

func (obj *ServiceAppliance) hasReferenceBase(name string) bool {
	if obj.baseMap == nil {
		return false
	}
	_, exists := obj.baseMap[name]
	return exists
}

func (obj *ServiceAppliance) UpdateDone() {
	obj.modified.SetInt64(0)
	obj.baseMap = nil
}

func (obj *ServiceAppliance) GetServiceApplianceUserCredentials() UserCredentials {
	return obj.service_appliance_user_credentials
}

func (obj *ServiceAppliance) SetServiceApplianceUserCredentials(value *UserCredentials) {
	obj.service_appliance_user_credentials = *value
	obj.modified.SetBit(&obj.modified, service_appliance_service_appliance_user_credentials, 1)
}

func (obj *ServiceAppliance) GetServiceApplianceIpAddress() string {
	return obj.service_appliance_ip_address
}

func (obj *ServiceAppliance) SetServiceApplianceIpAddress(value string) {
	obj.service_appliance_ip_address = value
	obj.modified.SetBit(&obj.modified, service_appliance_service_appliance_ip_address, 1)
}

func (obj *ServiceAppliance) GetServiceApplianceProperties() KeyValuePairs {
	return obj.service_appliance_properties
}

func (obj *ServiceAppliance) SetServiceApplianceProperties(value *KeyValuePairs) {
	obj.service_appliance_properties = *value
	obj.modified.SetBit(&obj.modified, service_appliance_service_appliance_properties, 1)
}

func (obj *ServiceAppliance) GetIdPerms() IdPermsType {
	return obj.id_perms
}

func (obj *ServiceAppliance) SetIdPerms(value *IdPermsType) {
	obj.id_perms = *value
	obj.modified.SetBit(&obj.modified, service_appliance_id_perms, 1)
}

func (obj *ServiceAppliance) GetPerms2() PermType2 {
	return obj.perms2
}

func (obj *ServiceAppliance) SetPerms2(value *PermType2) {
	obj.perms2 = *value
	obj.modified.SetBit(&obj.modified, service_appliance_perms2, 1)
}

func (obj *ServiceAppliance) GetAnnotations() KeyValuePairs {
	return obj.annotations
}

func (obj *ServiceAppliance) SetAnnotations(value *KeyValuePairs) {
	obj.annotations = *value
	obj.modified.SetBit(&obj.modified, service_appliance_annotations, 1)
}

func (obj *ServiceAppliance) GetDisplayName() string {
	return obj.display_name
}

func (obj *ServiceAppliance) SetDisplayName(value string) {
	obj.display_name = value
	obj.modified.SetBit(&obj.modified, service_appliance_display_name, 1)
}

func (obj *ServiceAppliance) readPhysicalInterfaceRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(service_appliance_physical_interface_refs) == 0) {
		err := obj.GetField(obj, "physical_interface_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *ServiceAppliance) GetPhysicalInterfaceRefs() (
	contrail.ReferenceList, error) {
	err := obj.readPhysicalInterfaceRefs()
	if err != nil {
		return nil, err
	}
	return obj.physical_interface_refs, nil
}

func (obj *ServiceAppliance) AddPhysicalInterface(
	rhs *PhysicalInterface, data ServiceApplianceInterfaceType) error {
	err := obj.readPhysicalInterfaceRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(service_appliance_physical_interface_refs) == 0 {
		obj.storeReferenceBase("physical-interface", obj.physical_interface_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), data}
	obj.physical_interface_refs = append(obj.physical_interface_refs, ref)
	obj.modified.SetBit(&obj.modified, service_appliance_physical_interface_refs, 1)
	return nil
}

func (obj *ServiceAppliance) DeletePhysicalInterface(uuid string) error {
	err := obj.readPhysicalInterfaceRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(service_appliance_physical_interface_refs) == 0 {
		obj.storeReferenceBase("physical-interface", obj.physical_interface_refs)
	}

	for i, ref := range obj.physical_interface_refs {
		if ref.Uuid == uuid {
			obj.physical_interface_refs = append(
				obj.physical_interface_refs[:i],
				obj.physical_interface_refs[i+1:]...)
			break
		}
	}
	obj.modified.SetBit(&obj.modified, service_appliance_physical_interface_refs, 1)
	return nil
}

func (obj *ServiceAppliance) ClearPhysicalInterface() {
	if (obj.valid.Bit(service_appliance_physical_interface_refs) != 0) &&
		(obj.modified.Bit(service_appliance_physical_interface_refs) == 0) {
		obj.storeReferenceBase("physical-interface", obj.physical_interface_refs)
	}
	obj.physical_interface_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, service_appliance_physical_interface_refs, 1)
	obj.modified.SetBit(&obj.modified, service_appliance_physical_interface_refs, 1)
}

func (obj *ServiceAppliance) SetPhysicalInterfaceList(
	refList []contrail.ReferencePair) {
	obj.ClearPhysicalInterface()
	obj.physical_interface_refs = make([]contrail.Reference, len(refList))
	for i, pair := range refList {
		obj.physical_interface_refs[i] = contrail.Reference{
			pair.Object.GetFQName(),
			pair.Object.GetUuid(),
			pair.Object.GetHref(),
			pair.Attribute,
		}
	}
}

func (obj *ServiceAppliance) readTagRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(service_appliance_tag_refs) == 0) {
		err := obj.GetField(obj, "tag_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *ServiceAppliance) GetTagRefs() (
	contrail.ReferenceList, error) {
	err := obj.readTagRefs()
	if err != nil {
		return nil, err
	}
	return obj.tag_refs, nil
}

func (obj *ServiceAppliance) AddTag(
	rhs *Tag) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(service_appliance_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.tag_refs = append(obj.tag_refs, ref)
	obj.modified.SetBit(&obj.modified, service_appliance_tag_refs, 1)
	return nil
}

func (obj *ServiceAppliance) DeleteTag(uuid string) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(service_appliance_tag_refs) == 0 {
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
	obj.modified.SetBit(&obj.modified, service_appliance_tag_refs, 1)
	return nil
}

func (obj *ServiceAppliance) ClearTag() {
	if (obj.valid.Bit(service_appliance_tag_refs) != 0) &&
		(obj.modified.Bit(service_appliance_tag_refs) == 0) {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}
	obj.tag_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, service_appliance_tag_refs, 1)
	obj.modified.SetBit(&obj.modified, service_appliance_tag_refs, 1)
}

func (obj *ServiceAppliance) SetTagList(
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

func (obj *ServiceAppliance) MarshalJSON() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalCommon(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(service_appliance_service_appliance_user_credentials) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.service_appliance_user_credentials)
		if err != nil {
			return nil, err
		}
		msg["service_appliance_user_credentials"] = &value
	}

	if obj.modified.Bit(service_appliance_service_appliance_ip_address) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.service_appliance_ip_address)
		if err != nil {
			return nil, err
		}
		msg["service_appliance_ip_address"] = &value
	}

	if obj.modified.Bit(service_appliance_service_appliance_properties) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.service_appliance_properties)
		if err != nil {
			return nil, err
		}
		msg["service_appliance_properties"] = &value
	}

	if obj.modified.Bit(service_appliance_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(service_appliance_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(service_appliance_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(service_appliance_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if len(obj.physical_interface_refs) > 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.physical_interface_refs)
		if err != nil {
			return nil, err
		}
		msg["physical_interface_refs"] = &value
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

func (obj *ServiceAppliance) UnmarshalJSON(body []byte) error {
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
		case "service_appliance_user_credentials":
			err = json.Unmarshal(value, &obj.service_appliance_user_credentials)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_appliance_service_appliance_user_credentials, 1)
			}
			break
		case "service_appliance_ip_address":
			err = json.Unmarshal(value, &obj.service_appliance_ip_address)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_appliance_service_appliance_ip_address, 1)
			}
			break
		case "service_appliance_properties":
			err = json.Unmarshal(value, &obj.service_appliance_properties)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_appliance_service_appliance_properties, 1)
			}
			break
		case "id_perms":
			err = json.Unmarshal(value, &obj.id_perms)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_appliance_id_perms, 1)
			}
			break
		case "perms2":
			err = json.Unmarshal(value, &obj.perms2)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_appliance_perms2, 1)
			}
			break
		case "annotations":
			err = json.Unmarshal(value, &obj.annotations)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_appliance_annotations, 1)
			}
			break
		case "display_name":
			err = json.Unmarshal(value, &obj.display_name)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_appliance_display_name, 1)
			}
			break
		case "tag_refs":
			err = json.Unmarshal(value, &obj.tag_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_appliance_tag_refs, 1)
			}
			break
		case "physical_interface_refs":
			{
				type ReferenceElement struct {
					To   []string
					Uuid string
					Href string
					Attr ServiceApplianceInterfaceType
				}
				var array []ReferenceElement
				err = json.Unmarshal(value, &array)
				if err != nil {
					break
				}
				obj.valid.SetBit(&obj.valid, service_appliance_physical_interface_refs, 1)
				obj.physical_interface_refs = make(contrail.ReferenceList, 0)
				for _, element := range array {
					ref := contrail.Reference{
						element.To,
						element.Uuid,
						element.Href,
						element.Attr,
					}
					obj.physical_interface_refs = append(obj.physical_interface_refs, ref)
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

func (obj *ServiceAppliance) UpdateObject() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalId(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(service_appliance_service_appliance_user_credentials) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.service_appliance_user_credentials)
		if err != nil {
			return nil, err
		}
		msg["service_appliance_user_credentials"] = &value
	}

	if obj.modified.Bit(service_appliance_service_appliance_ip_address) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.service_appliance_ip_address)
		if err != nil {
			return nil, err
		}
		msg["service_appliance_ip_address"] = &value
	}

	if obj.modified.Bit(service_appliance_service_appliance_properties) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.service_appliance_properties)
		if err != nil {
			return nil, err
		}
		msg["service_appliance_properties"] = &value
	}

	if obj.modified.Bit(service_appliance_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(service_appliance_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(service_appliance_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(service_appliance_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if obj.modified.Bit(service_appliance_physical_interface_refs) != 0 {
		if len(obj.physical_interface_refs) == 0 {
			var value json.RawMessage
			value, err := json.Marshal(
				make([]contrail.Reference, 0))
			if err != nil {
				return nil, err
			}
			msg["physical_interface_refs"] = &value
		} else if !obj.hasReferenceBase("physical-interface") {
			var value json.RawMessage
			value, err := json.Marshal(&obj.physical_interface_refs)
			if err != nil {
				return nil, err
			}
			msg["physical_interface_refs"] = &value
		}
	}

	if obj.modified.Bit(service_appliance_tag_refs) != 0 {
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

func (obj *ServiceAppliance) UpdateReferences() error {

	if (obj.modified.Bit(service_appliance_physical_interface_refs) != 0) &&
		len(obj.physical_interface_refs) > 0 &&
		obj.hasReferenceBase("physical-interface") {
		err := obj.UpdateReference(
			obj, "physical-interface",
			obj.physical_interface_refs,
			obj.baseMap["physical-interface"])
		if err != nil {
			return err
		}
	}

	if (obj.modified.Bit(service_appliance_tag_refs) != 0) &&
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

func ServiceApplianceByName(c contrail.ApiClient, fqn string) (*ServiceAppliance, error) {
	obj, err := c.FindByName("service-appliance", fqn)
	if err != nil {
		return nil, err
	}
	return obj.(*ServiceAppliance), nil
}

func ServiceApplianceByUuid(c contrail.ApiClient, uuid string) (*ServiceAppliance, error) {
	obj, err := c.FindByUuid("service-appliance", uuid)
	if err != nil {
		return nil, err
	}
	return obj.(*ServiceAppliance), nil
}
