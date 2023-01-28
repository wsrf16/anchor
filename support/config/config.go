package config

import (
	"github.com/wsrf16/swiss/sugar/base/lambda"
	"github.com/wsrf16/swiss/sugar/io/viperkit"
)

const Version = "v1.2.0"

var global *RootConfig

func Global() (global *RootConfig, err error) {
	if global == nil {
		global, err = viperkit.UnmarshalCurrentFileToT[RootConfig]("/config.yaml")
		if err != nil {
			return nil, err
		}
	}

	return lambda.IfFuncPair(global == nil, func() (*RootConfig, error) {
		return viperkit.UnmarshalCurrentFileToT[RootConfig]("/config.yaml")
	}, func() (*RootConfig, error) {
		return global, nil
	})
}
