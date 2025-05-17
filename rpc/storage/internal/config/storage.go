package config

type StorageConf struct {
	AccessKey string
	SecretKey string
	Bucket    string
	Endpoint  string
	UseSSL    bool
}
