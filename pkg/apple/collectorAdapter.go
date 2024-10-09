package apple

import (
	"fmt"
)

// RegionConfig holds region-specific settings
type RegionConfig struct {
	Prefix        string
	DefaultParams map[string]string
}

// GetRegionConfig returns the configuration for a given region
func GetRegionConfig(region string) (RegionConfig, error) {
	switch region {
	case "jp":
		return RegionConfig{
			Prefix:        "https://www.apple.com/jp",
			DefaultParams: map[string]string{},
		}, nil
	case "hk":
		return RegionConfig{
			Prefix:        "https://www.apple.com/hk-zh",
			DefaultParams: map[string]string{},
		}, nil
	case "us":
		return RegionConfig{
			Prefix:        "https://www.apple.com",
			DefaultParams: map[string]string{},
		}, nil
	case "cn":
		return RegionConfig{
			Prefix:        "https://www.apple.com.cn",
			DefaultParams: map[string]string{},
		}, nil
	// Add more regions as needed
	default:
		return RegionConfig{}, fmt.Errorf("unsupported region: %s", region)
	}
}
