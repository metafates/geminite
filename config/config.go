package config

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/knadh/koanf/v2"
	"github.com/metafates/geminite/afs"
	"github.com/metafates/geminite/where"
)

var Config struct {
	WPM   int  `koanf:"wpm" validate:"min=1"`
	Cache bool `koanf:"cache"`
}

const keyDelim = "."

func readConfigFile() ([]byte, error) {
	configFile, err := afs.AFS.OpenFile(
		where.ConfigFile(),
		os.O_CREATE|os.O_RDONLY,
		os.ModePerm,
	)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	stat, err := configFile.Stat()
	if err != nil {
		return nil, err
	}

	contents := make([]byte, stat.Size())

	_, err = configFile.Read(contents)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

func loadDefaults(k *koanf.Koanf) error {
	defaults := map[string]any{
		"wpm":   250,
		"cache": true,
	}

	return k.Load(confmap.Provider(defaults, keyDelim), nil)
}

func loadConfigFile(k *koanf.Koanf) error {
	configFile, err := readConfigFile()
	if err != nil {
		return err
	}

	return k.Load(
		rawbytes.Provider(configFile),
		toml.Parser(),
	)
}

func unmarshal(k *koanf.Koanf) error {
	return k.UnmarshalWithConf("", &Config, koanf.UnmarshalConf{
		Tag: "koanf",
	})
}

func Init() error {
	k := koanf.New(keyDelim)

	for _, loader := range []func(k *koanf.Koanf) error{
		loadDefaults,
		loadConfigFile,
	} {
		if err := loader(k); err != nil {
			return err
		}
	}

	if err := unmarshal(k); err != nil {
		return err
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	return validate.Struct(Config)
}
