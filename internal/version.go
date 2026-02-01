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
	fmt.Println("888    d8P                              88888888888 d8b                        ")
	fmt.Println("888   d8P                                   888     Y8P                        ")
	fmt.Println("888  d8P                                    888                                ")
	fmt.Println("888d88K     88888b.   .d88b.  888  888  888 888     888 88888b.d88b.   .d88b.  ")
	fmt.Println("8888888b    888  88b d88  88b 888  888  888 888     888 888  888  88b d8P  Y8b ")
	fmt.Println("888  Y88b   888  888 888  888 888  888  888 888     888 888  888  888 88888888 ")
	fmt.Println("888   Y88b  888  888 Y88..88P Y88b 888 d88P 888     888 888  888  888 Y8b.     ")
	fmt.Println("888    Y88b 888  888   Y88P     Y8888888P   888     888 888  888  888   Y8888  ")
	fmt.Println("                                                                               ")
	fmt.Println("                                                                               ")
	fmt.Println("Backend By pykelysia and lvshujun                                              ")
	fmt.Println("-----------------------------------")
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
