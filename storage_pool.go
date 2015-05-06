package libvirt

/*
#cgo LDFLAGS: -lvirt -ldl
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>
*/
import "C"

import (
	"errors"
	"io/ioutil"
	"unsafe"
)

type VirStoragePool struct {
	ptr C.virStoragePoolPtr
}

type VirStoragePoolInfo struct {
	ptr C.virStoragePoolInfo
}

func (p *VirStoragePool) Build(flags uint32) error {
	result := C.virStoragePoolBuild(p.ptr, C.uint(flags))
	if result == -1 {
		return errors.New(GetLastError())
	}
	return nil
}

func (p *VirStoragePool) Create(flags uint32) error {
	result := C.virStoragePoolCreate(p.ptr, C.uint(flags))
	if result == -1 {
		return errors.New(GetLastError())
	}
	return nil
}

func (p *VirStoragePool) Delete(flags uint32) error {
	result := C.virStoragePoolDelete(p.ptr, C.uint(flags))
	if result == -1 {
		return errors.New(GetLastError())
	}
	return nil
}

func (p *VirStoragePool) Destroy() error {
	result := C.virStoragePoolDestroy(p.ptr)
	if result == -1 {
		return errors.New(GetLastError())
	}
	return nil
}

func (p *VirStoragePool) Free() error {
	if result := C.virStoragePoolFree(p.ptr); result != 0 {
		return errors.New(GetLastError())
	}
	return nil
}

func (p *VirStoragePool) GetAutostart() (bool, error) {
	var out C.int
	result := C.virStoragePoolGetAutostart(p.ptr, (*C.int)(unsafe.Pointer(&out)))
	if result == -1 {
		return false, errors.New(GetLastError())
	}
	switch out {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}

func (p *VirStoragePool) GetInfo() (VirStoragePoolInfo, error) {
	pi := VirStoragePoolInfo{}
	var ptr C.virStoragePoolInfo
	result := C.virStoragePoolGetInfo(p.ptr, (*C.virStoragePoolInfo)(unsafe.Pointer(&ptr)))
	if result == -1 {
		return pi, errors.New(GetLastError())
	}
	pi.ptr = ptr
	return pi, nil
}

func (p *VirStoragePool) GetName() (string, error) {
	name := C.virStoragePoolGetName(p.ptr)
	if name == nil {
		return "", errors.New(GetLastError())
	}
	return C.GoString(name), nil
}

func (p *VirStoragePool) GetUUID() ([]byte, error) {
	var cUuid [C.VIR_UUID_BUFLEN](byte)
	cuidPtr := unsafe.Pointer(&cUuid)
	result := C.virStoragePoolGetUUID(p.ptr, (*C.uchar)(cuidPtr))
	if result != 0 {
		return []byte{}, errors.New(GetLastError())
	}
	return C.GoBytes(cuidPtr, C.VIR_UUID_BUFLEN), nil
}

func (p *VirStoragePool) GetUUIDString() (string, error) {
	var cUuid [C.VIR_UUID_STRING_BUFLEN](C.char)
	cuidPtr := unsafe.Pointer(&cUuid)
	result := C.virStoragePoolGetUUIDString(p.ptr, (*C.char)(cuidPtr))
	if result != 0 {
		return "", errors.New(GetLastError())
	}
	return C.GoString((*C.char)(cuidPtr)), nil
}

func (p *VirStoragePool) GetXMLDesc(flags uint32) (string, error) {
	result := C.virStoragePoolGetXMLDesc(p.ptr, C.uint(flags))
	if result == nil {
		return "", errors.New(GetLastError())
	}
	xml := C.GoString(result)
	C.free(unsafe.Pointer(result))
	return xml, nil
}

func (p *VirStoragePool) IsActive() (bool, error) {
	result := C.virStoragePoolIsActive(p.ptr)
	if result == -1 {
		return false, errors.New(GetLastError())
	}
	if result == 1 {
		return true, nil
	}
	return false, nil
}

func (p *VirStoragePool) SetAutostart(autostart bool) error {
	var cAutostart C.int
	switch autostart {
	case true:
		cAutostart = 1
	default:
		cAutostart = 0
	}
	result := C.virStoragePoolSetAutostart(p.ptr, cAutostart)
	if result == -1 {
		return errors.New(GetLastError())
	}
	return nil
}

func (p *VirStoragePool) Refresh(flags uint32) error {
	result := C.virStoragePoolRefresh(p.ptr, C.uint(flags))
	if result == -1 {
		return errors.New(GetLastError())
	}
	return nil
}

func (p *VirStoragePool) Undefine() error {
	result := C.virStoragePoolUndefine(p.ptr)
	if result == -1 {
		return errors.New(GetLastError())
	}
	return nil
}

func (i *VirStoragePoolInfo) GetState() uint8 {
	return uint8(i.ptr.state)
}

func (i *VirStoragePoolInfo) GetCapacityInBytes() uint64 {
	return uint64(i.ptr.capacity)
}

func (i *VirStoragePoolInfo) GetAllocationInBytes() uint64 {
	return uint64(i.ptr.allocation)
}

func (i *VirStoragePoolInfo) GetAvailableInBytes() uint64 {
	return uint64(i.ptr.available)
}

func (p *VirStoragePool) StorageVolCreateXMLFromFile(xmlFile string, flags uint32) (VirStorageVol, error) {
	xmlConfig, err := ioutil.ReadFile(xmlFile)
	if err != nil {
		return VirStorageVol{}, err
	}
	return p.StorageVolCreateXML(string(xmlConfig), flags)
}

func (p *VirStoragePool) StorageVolCreateXML(xmlConfig string, flags uint32) (VirStorageVol, error) {
	cXml := C.CString(string(xmlConfig))
	defer C.free(unsafe.Pointer(cXml))
	ptr := C.virStorageVolCreateXML(p.ptr, cXml, C.uint(flags))
	if ptr == nil {
		return VirStorageVol{}, errors.New(GetLastError())
	}
	return VirStorageVol{ptr: ptr}, nil
}

func (p *VirStoragePool) LookupStorageVolByName(name string) (VirStorageVol, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	ptr := C.virStorageVolLookupByName(p.ptr, cName)
	if ptr == nil {
		return VirStorageVol{}, errors.New(GetLastError())
	}
	return VirStorageVol{ptr: ptr}, nil
}

func (p *VirStoragePool) ListVolumes() ([]string, error) {
	var volumes [256](*C.char)
	volsPtr := unsafe.Pointer(&volumes)

	numVolumes := C.virStoragePoolListVolumes(
		p.ptr,
		(**C.char)(volsPtr),
		256)

	if numVolumes == -1 {
		return nil, errors.New(GetLastError())
	}

	goVolumes := make([]string, numVolumes)
	for k := 0; k < int(numVolumes); k++ {
		goVolumes[k] = C.GoString(volumes[k])
		C.free(unsafe.Pointer(volumes[k]))
	}
	return goVolumes, nil
}