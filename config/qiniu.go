package config

type Qiniu struct {
	AK     string `yaml:"AK"`
	SK        string `yaml:"SK"`
	PubBucket string `yaml:"PubBucket"`
	PubDomain string `yaml:"PubDomain"`
	VideoDomain string `yaml:"VideoDomain"`
	VideoBucket string `yaml:"VideoBucket"`
}

var QiniuConfig = new(Qiniu)
