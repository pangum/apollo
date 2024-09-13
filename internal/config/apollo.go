package config

type Apollo struct {
	// 配置端点
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty" xml:"endpoint,omitempty" toml:"endpoint,omitempty"`
	// 应用
	Appid string `json:"appid,omitempty" yaml:"appid,omitempty" xml:"appid,omitempty" toml:"appid,omitempty" validate:"required"`
	// 集群
	Cluster string `json:"cluster,omitempty" yaml:"cluster,omitempty" toml:"cluster,omitempty"`
	// 命名空间
	Namespaces []string `json:"namespaces,omitempty" yaml:"namespaces,omitempty" toml:"namespaces,omitempty"`
}
