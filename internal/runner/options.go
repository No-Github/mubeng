package runner

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/projectdiscovery/gologger"
	"ktbs.dev/mubeng/common"
	"ktbs.dev/mubeng/internal/updater"
)

// Options defines the values needed to execute the Runner.
func Options() *common.Options {
	opt := &common.Options{}

	flag.StringVar(&opt.File, "f", "", "")
	flag.StringVar(&opt.File, "file", "", "")

	flag.StringVar(&opt.Address, "a", "", "")
	flag.StringVar(&opt.Address, "address", "", "")

	flag.StringVar(&opt.Auth, "A", "", "")
	flag.StringVar(&opt.Auth, "auth", "", "")

	flag.BoolVar(&opt.Check, "c", false, "")
	flag.BoolVar(&opt.Check, "check", false, "")
	flag.BoolVar(&opt.Loop, "loop", false, "")

	flag.StringVar(&opt.CC, "only-cc", "", "")

	flag.DurationVar(&opt.Timeout, "t", 2*time.Second, "")
	flag.DurationVar(&opt.Timeout, "timeout", 2*time.Second, "")

	flag.IntVar(&opt.Rotate, "r", 1, "")
	flag.IntVar(&opt.Rotate, "rotate", 1, "")

	flag.StringVar(&opt.Method, "m", "sequent", "")
	flag.StringVar(&opt.Method, "method", "sequent", "")

	flag.BoolVar(&opt.Sync, "s", false, "")
	flag.BoolVar(&opt.Sync, "sync", false, "")

	flag.BoolVar(&opt.Verbose, "v", false, "")
	flag.BoolVar(&opt.Verbose, "verbose", false, "")

	flag.BoolVar(&opt.Daemon, "d", false, "")
	flag.BoolVar(&opt.Daemon, "daemon", false, "")

	flag.StringVar(&opt.Output, "o", "", "")
	flag.StringVar(&opt.Output, "output", "", "")

	flag.StringVar(&opt.Url, "u", "", "")
	//flag.StringVar(&opt.Out, "o", ".", "配置文件保存路径")

	flag.BoolVar(&Verify, "verify", false, "验证配置文件可用性")

	flag.BoolVar(&doUpdate, "update", false, "")

	flag.BoolVar(&version, "V", false, "")
	flag.BoolVar(&version, "version", false, "")

	flag.BoolVar(&opt.Watch, "w", false, "")
	flag.BoolVar(&opt.Watch, "watch", false, "")

	flag.IntVar(&opt.Goroutine, "g", 10, "")
	flag.IntVar(&opt.Goroutine, "goroutine", 10, "")

	flag.Usage = func() {
		//showBanner()
		showUsage()
	}
	flag.Parse()

	if version {
		showVersion()
	}
	//showBanner()

	if opt.Url != "" {
		// 发送 GET 请求获取资源
		resp, err := http.Get(opt.Url)
		if err != nil {
			fmt.Println("Error fetching resource:", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		// 创建本地文件
		file, err := os.Create("config.yaml")
		if err != nil {
			fmt.Println("Error creating file:", err)
			os.Exit(1)
		}
		defer file.Close()

		// 将获取到的资源写入本地文件
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			os.Exit(1)
		}

		// 移动文件到指定位置
		err = os.Rename("config.yaml", os.ExpandEnv("."+"/config.yaml"))
		if err != nil {
			fmt.Println("Error moving file:", err)
			os.Exit(1)
		}
		fmt.Println("配置文件下载地址:", opt.Url)
		fmt.Println("配置文件保存路径:", ".")
		os.Exit(0)
	}

	if doUpdate {
		if err := updater.New(); err != nil {
			gologger.Fatal().Msgf("Error! %s.", err)
		}
	}

	if Verify {
		fmt.Println("验证配置文件(./config.yaml)可用性中")
		Verify_Yaml("./config.yaml")
		os.Exit(0)
	}

	if err := validate(opt); err != nil {
		gologger.Fatal().Msgf("Error! %s.", err)
	}

	return opt
}
