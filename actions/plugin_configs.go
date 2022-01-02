// Copyright (c) 2021 Damien Stuart. All rights reserved.
//
// Use of this source code is governed by the MIT License that can be found
// in the LICENSE file.
//
package plugin_data

type PluginsConfig struct {
	Noop struct {
		Test1 string `default:"hello world" yaml:"test_1"`
		Test2 string `default:"hello back" yaml:"test_2"`
	} `yaml:"noop"`
	Clickhouse struct {
		LogMaxSize    int  `default:"1024" yaml:"log_size_max"`
		LogMaxBackups int  `default:"7" yaml:"log_backups_max"`
		LogMaxAge     int  `yaml:"log_age_max"`
		LogCompress   bool `default:"false" yaml:"compress_rotated_logs"`
	} `yaml:"clickhouse"`
	Logger struct {
		LogMaxSize    int  `default:"1024" yaml:"log_size_max"`
		LogMaxBackups int  `default:"7" yaml:"log_backups_max"`
		LogMaxAge     int  `yaml:"log_age_max"`
		LogCompress   bool `default:"false" yaml:"compress_rotated_logs"`
	} `yaml:"logger"`
}
