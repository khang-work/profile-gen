package info

import "fmt"

const SoftwareName = "profile-gen"
const SoftwareDescriptiveName = "Profile Generator"

const SoftwareVersionMajor = 0
const SoftwareVersionMinor = 1
const SoftwareVersionBugfix = 0

var SoftwareVersion string = fmt.Sprintf("%d.%d.%d", SoftwareVersionMajor, SoftwareVersionMinor, SoftwareVersionBugfix)
var SoftwareTitle string = fmt.Sprintf("%s v%s", SoftwareDescriptiveName, SoftwareVersion)
