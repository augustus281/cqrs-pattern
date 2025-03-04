package constants

const (
	GrpcPort                   = "GRPC_PORT"
	HttpPort                   = "HTTP_PORT"
	ConfigPath                 = "CONFIG_PATH"
	KafkaBrokers               = "KAFKA_BROKERS"
	JaegerHostPort             = "JAEGER_HOST"
	RedisAddr                  = "REDIS_ADDR"
	MongoDbURI                 = "MONGO_URI"
	EventStoreConnectionString = "EVENT_STORE_CONNECTION_STRING"
	ElasticUrl                 = "ELASTIC_URL"

	ReaderServicePort = "READER_SERVICE"

	Yaml          = "yaml"
	Tcp           = "tcp"
	Redis         = "redis"
	Kafka         = "kafka"
	Postgres      = "postgres"
	MongoDB       = "mongo"
	ElasticSearch = "elasticSearch"

	GRPC     = "GRPC"
	SIZE     = "SIZE"
	URI      = "URI"
	STATUS   = "STATUS"
	HTTP     = "HTTP"
	ERROR    = "ERROR"
	METHOD   = "METHOD"
	METADATA = "METADATA"
	REQUEST  = "REQUEST"
	REPLY    = "REPLY"
	TIME     = "TIME"

	Topic        = "topic"
	Partition    = "partition"
	Message      = "message"
	WorkerID     = "workerID"
	Offset       = "offset"
	Time         = "time"
	GroupName    = "GroupName"
	StreamID     = "StreamID"
	EventID      = "EventID"
	EventType    = "EventType"
	EventNumber  = "EventNumber"
	CreatedDate  = "CreatedDate"
	UserMetadata = "UserMetadata"

	Page   = "page"
	Size   = "size"
	Search = "search"
	ID     = "id"

	EsAll = "$all"

	Validate        = "validate"
	FieldValidation = "field validation"
	RequiredHeaders = "required header"
	Base64          = "base64"
	Unmarshal       = "unmarshal"
	Uuid            = "uuid"
	Cookie          = "cookie"
	Token           = "token"
	Bcrypt          = "bcrypt"
	SQLState        = "sqlstate"

	MessageSize = "MessageSize"

	MongoProjection   = "(MongoDB Projection)"
	ElasticProjection = "(Elastic Projection)"

	OrderIdIndex    = "order_id"
	OrderId         = "order_id"
	DeliveryAddress = "delivery_address"
	Submitted       = "submitted"
	Completed       = "completed"
	DeliveredTime   = "delivered_time"
	Payment         = "payment"
	Paid            = "paid"
	Canceled        = "canceled"
	CancelReason    = "cancel_reason"

	KafkaHeaders = "kafkaHeaders"
)
