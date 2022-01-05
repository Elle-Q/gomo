package config

type Qiniu struct {
	AK     string `yaml:"AK"`
	SK     string `yaml:"SK"`
	BUCKET string `yaml:"BUCKET"`
}

var QiniuConfig = new(Qiniu)
