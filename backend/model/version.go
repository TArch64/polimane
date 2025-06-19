package model

const (
	PKVersion = "#VERSION"
	SKVersion = "#METADATA"
)

type Version struct {
	PK      string `dynamo:"PK,hash"`
	SK      string `dynamo:"SK,range"`
	Version int    `dynamo:"Version"`
}

func IntVersion(version int) *Version {
	return &Version{
		PK:      PKVersion,
		SK:      SKVersion,
		Version: version,
	}
}

func NewVersion() *Version {
	return IntVersion(-1)
}

func (v *Version) NextVersion() int {
	return v.Version + 1
}

func (v *Version) IsLatest(versionCount int) bool {
	return v.NextVersion() >= versionCount
}
