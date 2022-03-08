package setting

import (
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	//create viper and set up viper
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")

	//read the configs.yaml
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Setting{vp: vp}, nil
}

//ReadSection read each section of the configs.yaml file by key and value
func (s *Setting) ReadSection(key string, v interface{}) error {
	if err := s.vp.UnmarshalKey(key, v); err != nil {
		return err
	}

	return nil
}
