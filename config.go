package jaeger

import "fmt"

const (
	samplerDefault = 0.1
	samplerMin     = 0.01
	samplerMax     = 1.0
)

// 定义 jaeger 设置项
type Config struct {
	// 服务名
	Name string `yaml:"name" json:"name"`
	// 是否启用
	Enable bool `yaml:"enable" json:"enable"`
	// 采样率
	Sampler float64 `yaml:"sampler" json:"sampler"`
	// 端点
	Endpoint string `yaml:"endpoint" json:"endpoint"`
}

// 默认设置
func DefaultConfig() *Config {
	return &Config{
		Name:     "jaeger-service",
		Enable:   true,
		Sampler:  samplerDefault,
		Endpoint: ":6831",
	}
}

// 设置项检查
func (c *Config) Check() {
	if c.Sampler < samplerMin {
		c.Sampler = samplerMin
	}

	if c.Sampler > samplerMax {
		c.Sampler = samplerMax
	}
}

// 重写 String 方法
func (c *Config) String() string {
	return fmt.Sprintf(
		"\nJaeger Client Config:"+
			"\n\tName: %s"+
			"\n\tEnable: %t"+
			"\n\tSampler: %f"+
			"\n\tEndpoint: %s\n",
		c.Name,
		c.Enable,
		c.Sampler,
		c.Endpoint)
}
