package libvirt

// virDomainState
const (
	VIR_DOMAIN_NOSTATE     = 0
	VIR_DOMAIN_RUNNING     = 1
	VIR_DOMAIN_BLOCKED     = 2
	VIR_DOMAIN_PAUSED      = 3
	VIR_DOMAIN_SHUTDOWN    = 4
	VIR_DOMAIN_SHUTOFF     = 5
	VIR_DOMAIN_CRASHED     = 6
	VIR_DOMAIN_PMSUSPENDED = 7
)

// virDomainVcpuFlags
const (
	VIR_DOMAIN_VCPU_CURRENT = 0
	VIR_DOMAIN_VCPU_LIVE    = 1
	VIR_DOMAIN_VCPU_CONFIG  = 2
	VIR_DOMAIN_VCPU_MAXIMUM = 4
	VIR_DOMAIN_VCPU_GUEST   = 8
)

// virDomainMemoryModFlags
const (
	VIR_DOMAIN_MEM_CURRENT = 0
	VIR_DOMAIN_MEM_LIVE    = 1
	VIR_DOMAIN_MEM_CONFIG  = 2
	VIR_DOMAIN_MEM_MAXIMUM = 4
)

// virStoragePoolState
const (
	VIR_STORAGE_POOL_INACTIVE     = 0 // Not running
	VIR_STORAGE_POOL_BUILDING     = 1 // Initializing pool,not available
	VIR_STORAGE_POOL_RUNNING      = 2 // Running normally
	VIR_STORAGE_POOL_DEGRADED     = 3 // Running degraded
	VIR_STORAGE_POOL_INACCESSIBLE = 4 // Running,but not accessible
)

// virStoragePoolBuildFlags
const (
	VIR_STORAGE_POOL_BUILD_NEW          = 0 // Regular build from scratch
	VIR_STORAGE_POOL_BUILD_REPAIR       = 1 // Repair / reinitialize
	VIR_STORAGE_POOL_BUILD_RESIZE       = 2 // Extend existing pool
	VIR_STORAGE_POOL_BUILD_NO_OVERWRITE = 4 // Do not overwrite existing pool
	VIR_STORAGE_POOL_BUILD_OVERWRITE    = 8 // Overwrite data
)

// virDomainDestroyFlags
const (
	VIR_DOMAIN_DESTROY_DEFAULT  = 0
	VIR_DOMAIN_DESTROY_GRACEFUL = 1
)

// virDomainShutdownFlags
const (
	VIR_DOMAIN_SHUTDOWN_DEFAULT        = 0
	VIR_DOMAIN_SHUTDOWN_ACPI_POWER_BTN = 1
	VIR_DOMAIN_SHUTDOWN_GUEST_AGENT    = 2
	VIR_DOMAIN_SHUTDOWN_INITCTL        = 4
	VIR_DOMAIN_SHUTDOWN_SIGNAL         = 8
)

// virDomainAttachDeviceFlags
const (
	VIR_DOMAIN_DEVICE_MODIFY_CURRENT = 0
	VIR_DOMAIN_DEVICE_MODIFY_LIVE    = 1
	VIR_DOMAIN_DEVICE_MODIFY_CONFIG  = 2
	VIR_DOMAIN_DEVICE_MODIFY_FORCE   = 4
)

// virStorageVolDeleteFlags
const (
	VIR_STORAGE_VOL_DELETE_NORMAL = 0 // Delete metadata only (fast)
	VIR_STORAGE_VOL_DELETE_ZEROED = 1 // Clear all data to zeros (slow)
)

// virStorageVolResizeFlags
const (
	VIR_STORAGE_VOL_RESIZE_ALLOCATE = 1 // force allocation of new size
	VIR_STORAGE_VOL_RESIZE_DELTA    = 2 // size is relative to current
	VIR_STORAGE_VOL_RESIZE_SHRINK   = 4 // allow decrease in capacity
)

// virStorageVolType
const (
	VIR_STORAGE_VOL_FILE  = 0 // Regular file based volumes
	VIR_STORAGE_VOL_BLOCK = 1 // Block based volumes
	VIR_STORAGE_VOL_DIR   = 2 // Directory-passthrough based volume
)

// virStorageVolWipeAlgorithm
const (
	VIR_STORAGE_VOL_WIPE_ALG_ZERO       = 0 // 1-pass, all zeroes
	VIR_STORAGE_VOL_WIPE_ALG_NNSA       = 1 // 4-pass NNSA Policy Letter NAP-14.1-C (XVI-8)
	VIR_STORAGE_VOL_WIPE_ALG_DOD        = 2 // 4-pass DoD 5220.22-M section 8-306 procedure
	VIR_STORAGE_VOL_WIPE_ALG_BSI        = 3 // 9-pass method recommended by the German Center of Security in Information Technologies
	VIR_STORAGE_VOL_WIPE_ALG_GUTMANN    = 4 // The canonical 35-pass sequence
	VIR_STORAGE_VOL_WIPE_ALG_SCHNEIER   = 5 // 7-pass method described by Bruce Schneier in "Applied Cryptography" (1996)
	VIR_STORAGE_VOL_WIPE_ALG_PFITZNER7  = 6 // 7-pass random
	VIR_STORAGE_VOL_WIPE_ALG_PFITZNER33 = 7 // 33-pass random
	VIR_STORAGE_VOL_WIPE_ALG_RANDOM     = 8 // 1-pass random
)

// virSecretUsageType
const (
	VIR_SECRET_USAGE_TYPE_NONE   = 0
	VIR_SECRET_USAGE_TYPE_VOLUME = 1
	VIR_SECRET_USAGE_TYPE_CEPH   = 2
	VIR_SECRET_USAGE_TYPE_ISCSI  = 3
)

// virStreamFlags
const (
	VIR_STREAM_NONBLOCK = 1
)

// virKeycodeSet
const (
	VIR_KEYCODE_SET_LINUX  = 0
	VIR_KEYCODE_SET_XT     = 1
	VIR_KEYCODE_SET_ATSET1 = 2
	VIR_KEYCODE_SET_ATSET2 = 3
	VIR_KEYCODE_SET_ATSET3 = 4
	VIR_KEYCODE_SET_OSX    = 5
	VIR_KEYCODE_SET_XT_KBD = 6
	VIR_KEYCODE_SET_USB    = 7
	VIR_KEYCODE_SET_WIN32  = 8
	VIR_KEYCODE_SET_RFB    = 9
)

// virDomainCreateFlags
const (
	VIR_DOMAIN_NONE               = 0
	VIR_DOMAIN_START_PAUSED       = 1
	VIR_DOMAIN_START_AUTODESTROY  = 2
	VIR_DOMAIN_START_BYPASS_CACHE = 4
	VIR_DOMAIN_START_FORCE_BOOT   = 8
)

const VIR_DOMAIN_MEMORY_PARAM_UNLIMITED = 9007199254740991
