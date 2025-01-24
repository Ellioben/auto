/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var conf Conf

type Conf struct {
	Name string `mapstructure:"name"`
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "auto",
	Short: "1. root",
	Long:  `root is the root command of the application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		// 不输入 --name 从配置文件中读取 name
		if len(name) == 0 {
			name = viper.GetString("name")
			// 配置文件中未读取到 name，打印帮助提示
			if len(name) == 0 {
				cmd.Help()
				return
			}
		}
		fmt.Printf("Create rule %s success.\n", name)
		fmt.Print("----" + conf.Name + "\n")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.auto.yaml)")
	//viper.AddConfigPath("./cmd/config.yaml") // 可设置多个搜索路径
	//viper.ReadInConfig()
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// 在指定标志时可以用 --name，也可以使用短写 -n
	// 加载优先级比较高
	rootCmd.Flags().StringVarP(&name, "name", "n", "./cmd/config.yaml", "createCmd下的name参数")
	viper.SetConfigFile(name)   // 指定配置文件路径
	err := viper.ReadInConfig() // 读取配置信息
	if err != nil {             // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))

	}
	// 将读取的配置信息保存至全局变量Conf
	if err := viper.Unmarshal(&conf); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}
}
