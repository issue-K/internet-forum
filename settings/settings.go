package settings

import(
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() ( err error ){
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil{
		fmt.Printf("viper.ReadInConfig failed,err: %v\n",err )
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("配置文件被修改了\n")
	})
	return
}