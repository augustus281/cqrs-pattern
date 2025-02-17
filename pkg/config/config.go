package config

type Config struct {
	Server         ServerConfig         `mapstructure:"config"`
	PostgreSQL     PostgreSQLConfig     `mapstructure:"postgresql"`
	MongoDB        MongoDBConfig        `mapstructure:"mongo"`
	Logger         LoggerConfig         `mapstructure:"logger"`
	Jaeger         JaegerConfig         `mapstructure:"jaeger"`
	EventStore     EventStoreConfig     `mapstructure:"event_store"`
	ElasticSearch  ElasticSearchConfig  `mapstructure:"elastic_search"`
	ElasticIndexes ElasticIndexesConfig `mapstructure:"elastic_indexes"`
	ServiceName    ServiceNameConfig    `mapstructure:"service_name"`
}

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type PostgreSQLConfig struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	DBName          string `mapstructure:"dbname"`
	SslMode         string `mapstructure:"sslmode"`
	Timezone        string `mapstructure:"timezone"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

type MongoDBConfig struct {
	URI      string `mapstructure:"uri"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Db       string `mapstructure:"db"`
}

type LoggerConfig struct {
	Log_level     string `mapstructure:"log_level"`
	File_log_name string `mapstructure:"file_log_name"`
	Max_backups   int    `mapstructure:"max_backups"`
	Max_age       int    `mapstructure:"max_age"`
	Max_size      int    `mapstructure:"max_size"`
	Compress      bool   `mapstructure:"compress"`
}

type JaegerConfig struct {
	Enable      bool   `mapstructure:"enable"`
	ServiceName string `mapstructure:"service_name"`
	HostPort    string `mapstructure:"host_port"`
	LogSpans    bool   `mapstructure:"log_spans"`
}

type ElasticSearchConfig struct {
	Url         string `mapstructure:"url"`
	Sniff       bool   `mapstructure:"sniff"`
	Gzip        bool   `mapstructure:"gzip"`
	Explain     bool   `mapstructure:"explain"`
	FetchSource bool   `mapstructure:"fetch_source"`
	Version     bool   `mapstructure:"version"`
	Pretty      bool   `mapstructure:"pretty"`
}

type ElasticIndexesConfig struct {
	Orders string `mapstructure:"orders" validate:"required"`
}

type EventStoreConfig struct {
	ConnectionString string `mapstructure:"connection_string"`
}

type ServiceNameConfig struct {
	ServiceName string `mapstruct:"service_name"`
}
