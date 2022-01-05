package config

type Qiniu struct {
	AK     string `yaml:"AK"`
	SK     string `yaml:"SK"`
	BUCKET string `yaml:"BUCKET"`
	PubDomain string `yaml:"PubDomain"`
}

var QiniuConfig = new(Qiniu)
