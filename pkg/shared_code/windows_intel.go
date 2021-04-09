package shared_code

// import the libraries we need
import ( // necessary for gathering process information
	// necessary for getting netowork information
	// Software information Enumeration
	"golang.org/x/sys/windows/registry"
)

// GetCPUmodel inserts information about the target system's CPU model
// into a hostintel struct
func GetCPUmodel(intel_struct HostIntel) {
	key := `HARDWARE\DESCRIPTION\System\CentralProcessor\0`
	regValue := "Identifier"
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, key, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer k.Close()

	v, _, err := k.GetStringValue(regValue)
	if err != nil {
		return "", err
	}
	return v, err
}

// GetCPUname returns the target system's CPU name
func GetCPUname() (string, error) {
	key := `HARDWARE\DESCRIPTION\System\CentralProcessor\0`
	regValue := "ProcessorNameString"

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, key, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer k.Close()

	v, _, err := k.GetStringValue(regValue)
	if err != nil {
		return "", err
	}
	return v, err
}

// code from:
// https://github.com/bluesentinelsec/OffensiveGoLang/blob/master/pkg/windows/discovery/os.go
// GetOSinfo returns information about the target system OS
// in a struct that must be converted to json and encrypted
func GetOSinfo() (OSInfo, error) {
	var OSInformation OSInfo
	key := `SOFTWARE\Microsoft\Windows NT\CurrentVersion`

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, key, registry.QUERY_VALUE)
	if err != nil {
		return OSInformation, err
	}
	defer k.Close()
	OSInformation.InstallationType, _, _ = k.GetStringValue("InstallationType")
	OSInformation.ProductID, _, _ = k.GetStringValue("ProductId")
	OSInformation.ProductName, _, _ = k.GetStringValue("ProductName")
	OSInformation.RegisteredOwner, _, _ = k.GetStringValue("RegisteredOwner")
	OSInformation.ReleaseID, _, _ = k.GetStringValue("ReleaseId")
	OSInformation.CurrentBuild, _, _ = k.GetStringValue("CurrentBuild")
	// ToDo: convert epoc times to readable format
	OSInformation.InstallDate, _, _ = k.GetIntegerValue("InstallDate")
	OSInformation.InstallTime, _, _ = k.GetIntegerValue("InstallTime")

	return OSInformation, nil
}
