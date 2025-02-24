package config

type Config struct {
	Server             ServerConfig           `mapstructure:"server"`
	GRPC               GRPCConfig             `mapstructure:"grpc"`
	Probes             ProbesConfig           `mapstructure:"probes"`
	PostgreSQL         PostgreSQLConfig       `mapstructure:"postgres"`
	Redis              RedisConfig            `mapstructure:"redis"`
	MongoDB            MongoDBConfig          `mapstructure:"mongo"`
	MongoDBCollections MongoCollectionsConfig `mapstructure:"mongo_collections"`
	Logger             LoggerConfig           `mapstructure:"logger"`
	Jaeger             JaegerConfig           `mapstructure:"jaeger"`
	EventStore         EventStoreConfig       `mapstructure:"event_store"`
	ElasticSearch      ElasticSearchConfig    `mapstructure:"elastic_search"`
	ElasticIndexes     ElasticIndexesConfig   `mapstructure:"elastic_indexes"`
	ServiceName        ServiceNameConfig      `mapstructure:"service_name"`
	Kafka              KafkaConfig            `mapstructure:"kafka"`
	KafkaTopics        KafkaTopicsConfig      `mapstructure:"kafka_topics"`
}

type ServerConfig struct {
	Port  int    `mapstructure:"port"`
	Mode  string `mapstructure:"mode"`
	Debug bool   `mapstructure:"debug"`
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

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

type MongoDBConfig struct {
	URI      string `mapstructure:"uri"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Db       string `mapstructure:"db"`
}

type MongoCollectionsConfig struct {
	Shop string `mapstructure:"orders" validate:"required"`
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
	ServiceName string `mapstructure:"service_name"`
}

type KafkaConfig struct {
	Brokers    []string `mapstructure:"brokers"`
	GroupID    string   `mapstructure:"group_id"`
	InitTopics bool     `mapstructure:"init_topics"`
}

type KafkaTopicsConfig struct {
	EventCreated EventCreated `mapstructure:"event_created"`
}

type EventCreated struct {
	TopicName         string `mapstructure:"topic_name"`
	Partitions        int    `mapstructure:"partitions"`
	ReplicationFactor int    `mapstructure:"replication_factor"`
}

type GRPCConfig struct {
	Port        int  `mapstructure:"port"`
	Development bool `mapstructure:"development"`
}

type ProbesConfig struct {
	LivenessPath         string `mapstructure:"livenessPath"`
	ReadinessPath        string `mapstructure:"readinessPath"`
	Port                 string `mapstructure:"port"`
	Pprof                string `mapstructure:"pprof"`
	PrometheusPath       string `mapstructure:"prometheusPath"`
	PrometheusPort       int    `mapstructure:"prometheusPort"`
	CheckIntervalSeconds int    `mapstructure:"checkIntervalSeconds"`
}
