package internal

import (
	"fmt"
	"time"
)

var (
	BuildTime = ""      // 编译时间
	GitCommit = ""      // Git Commit ID
	Version   = "1.0.0" // 版本号
)

// PrintBuildInfo 打印构建信息
func PrintBuildInfo() {
	fmt.Printf("KnowTime Backend Starting...\n")

	if GitCommit != "" {
		fmt.Printf("Git Commit: %s\n", GitCommit)
	} else {
		fmt.Printf("Git Commit: unknown\n")
	}

	if BuildTime != "" {
		fmt.Printf("Build Time: %s\n", BuildTime)
	} else {
		fmt.Printf("Build Time: unknown\n")
	}

	fmt.Printf("Version: %s\n", Version)
	fmt.Println("-----------------------------------")
}

// ParseBuildTime 尝试将字符串形式的构建时间转换为time.Time
func ParseBuildTime() (*time.Time, error) {
	if BuildTime == "" {
		return nil, fmt.Errorf("build time is not set")
	}

	parsedTime, err := time.Parse(time.RFC3339, BuildTime)
	if err != nil {
		return nil, err
	}

	return &parsedTime, nil
}
