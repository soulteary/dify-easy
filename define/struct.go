package Define

type HealthCheck struct {
	Test        string `yaml:"test"`
	Interval    string `yaml:"interval,omitempty"`
	StartPeriod string `yaml:"start_period,omitempty"`
	Timeout     string `yaml:"timeout,omitempty"`
	Retries     int    `yaml:"retries,omitempty"`
}

type DifyEnvirontment map[string]any

type XSharedEnv struct {
	LOG_LEVEL   string `yaml:"LOG_LEVEL"`
	LOG_FILE    string `yaml:"LOG_FILE"`
	DEBUG       string `yaml:"DEBUG"`
	FLASK_DEBUG string `yaml:"FLASK_DEBUG"`

	SECRET_KEY    string `yaml:"SECRET_KEY"`
	INIT_PASSWORD string `yaml:"INIT_PASSWORD"`

	CONSOLE_WEB_URL  string `yaml:"CONSOLE_WEB_URL"`
	CONSOLE_API_URL  string `yaml:"CONSOLE_API_URL"`
	SERVICE_API_URL  string `yaml:"SERVICE_API_URL"`
	APP_WEB_URL      string `yaml:"APP_WEB_URL"`
	CHECK_UPDATE_URL string `yaml:"CHECK_UPDATE_URL"`
	OPENAI_API_BASE  string `yaml:"OPENAI_API_BASE"`

	FILES_URL            string `yaml:"FILES_URL"`
	FILES_ACCESS_TIMEOUT string `yaml:"FILES_ACCESS_TIMEOUT"`

	APP_MAX_ACTIVE_REQUESTS string `yaml:"APP_MAX_ACTIVE_REQUESTS"`
	MIGRATION_ENABLED       string `yaml:"MIGRATION_ENABLED"`

	DEPLOY_ENV        string `yaml:"DEPLOY_ENV"`
	DIFY_BIND_ADDRESS string `yaml:"DIFY_BIND_ADDRESS"`
	DIFY_PORT         string `yaml:"DIFY_PORT"`

	SERVER_WORKER_AMOUNT             string `yaml:"SERVER_WORKER_AMOUNT"`
	SERVER_WORKER_CLASS              string `yaml:"SERVER_WORKER_CLASS"`
	CELERY_WORKER_CLASS              string `yaml:"CELERY_WORKER_CLASS"`
	GUNICORN_TIMEOUT                 string `yaml:"GUNICORN_TIMEOUT"`
	CELERY_WORKER_AMOUNT             string `yaml:"CELERY_WORKER_AMOUNT"`
	CELERY_AUTO_SCALE                string `yaml:"CELERY_AUTO_SCALE"`
	CELERY_MAX_WORKERS               string `yaml:"CELERY_MAX_WORKERS"`
	CELERY_MIN_WORKERS               string `yaml:"CELERY_MIN_WORKERS"`
	API_TOOL_DEFAULT_CONNECT_TIMEOUT string `yaml:"API_TOOL_DEFAULT_CONNECT_TIMEOUT"`
	API_TOOL_DEFAULT_READ_TIMEOUT    string `yaml:"API_TOOL_DEFAULT_READ_TIMEOUT"`

	DB_USERNAME string `yaml:"DB_USERNAME"`
	DB_PASSWORD string `yaml:"DB_PASSWORD"`
	DB_HOST     string `yaml:"DB_HOST"`
	DB_PORT     string `yaml:"DB_PORT"`
	DB_DATABASE string `yaml:"DB_DATABASE"`

	SQLALCHEMY_POOL_SIZE    string `yaml:"SQLALCHEMY_POOL_SIZE"`
	SQLALCHEMY_POOL_RECYCLE string `yaml:"SQLALCHEMY_POOL_RECYCLE"`
	SQLALCHEMY_ECHO         string `yaml:"SQLALCHEMY_ECHO"`

	REDIS_HOST                    string `yaml:"REDIS_HOST"`
	REDIS_PORT                    string `yaml:"REDIS_PORT"`
	REDIS_USERNAME                string `yaml:"REDIS_USERNAME"`
	REDIS_PASSWORD                string `yaml:"REDIS_PASSWORD"`
	REDIS_USE_SSL                 string `yaml:"REDIS_USE_SSL"`
	REDIS_DB                      int    `yaml:"REDIS_DB"`
	REDIS_USE_SENTINEL            string `yaml:"REDIS_USE_SENTINEL"`
	REDIS_SENTINELS               string `yaml:"REDIS_SENTINELS"`
	REDIS_SENTINEL_SERVICE_NAME   string `yaml:"REDIS_SENTINEL_SERVICE_NAME"`
	REDIS_SENTINEL_USERNAME       string `yaml:"REDIS_SENTINEL_USERNAME"`
	REDIS_SENTINEL_PASSWORD       string `yaml:"REDIS_SENTINEL_PASSWORD"`
	REDIS_SENTINEL_SOCKET_TIMEOUT string `yaml:"REDIS_SENTINEL_SOCKET_TIMEOUT"`

	CELERY_BROKER_URL string `yaml:"CELERY_BROKER_URL"`

	BROKER_USE_SSL string `yaml:"BROKER_USE_SSL"`

	CELERY_USE_SENTINEL            string `yaml:"CELERY_USE_SENTINEL"`
	CELERY_SENTINEL_MASTER_NAME    string `yaml:"CELERY_SENTINEL_MASTER_NAME"`
	CELERY_SENTINEL_SOCKET_TIMEOUT string `yaml:"CELERY_SENTINEL_SOCKET_TIMEOUT"`

	WEB_API_CORS_ALLOW_ORIGINS string `yaml:"WEB_API_CORS_ALLOW_ORIGINS"`
	CONSOLE_CORS_ALLOW_ORIGINS string `yaml:"CONSOLE_CORS_ALLOW_ORIGINS"`

	STORAGE_TYPE string `yaml:"STORAGE_TYPE"`

	STORAGE_LOCAL_PATH string `yaml:"STORAGE_LOCAL_PATH"`

	S3_USE_AWS_MANAGED_IAM string `yaml:"S3_USE_AWS_MANAGED_IAM"`
	S3_ENDPOINT            string `yaml:"S3_ENDPOINT"`
	S3_BUCKET_NAME         string `yaml:"S3_BUCKET_NAME"`
	S3_ACCESS_KEY          string `yaml:"S3_ACCESS_KEY"`
	S3_SECRET_KEY          string `yaml:"S3_SECRET_KEY"`
	S3_REGION              string `yaml:"S3_REGION"`

	AZURE_BLOB_ACCOUNT_NAME   string `yaml:"AZURE_BLOB_ACCOUNT_NAME"`
	AZURE_BLOB_ACCOUNT_KEY    string `yaml:"AZURE_BLOB_ACCOUNT_KEY"`
	AZURE_BLOB_CONTAINER_NAME string `yaml:"AZURE_BLOB_CONTAINER_NAME"`
	AZURE_BLOB_ACCOUNT_URL    string `yaml:"AZURE_BLOB_ACCOUNT_URL"`

	GOOGLE_STORAGE_BUCKET_NAME                 string `yaml:"GOOGLE_STORAGE_BUCKET_NAME"`
	GOOGLE_STORAGE_SERVICE_ACCOUNT_JSON_BASE64 string `yaml:"GOOGLE_STORAGE_SERVICE_ACCOUNT_JSON_BASE64"`

	ALIYUN_OSS_BUCKET_NAME  string `yaml:"ALIYUN_OSS_BUCKET_NAME"`
	ALIYUN_OSS_ACCESS_KEY   string `yaml:"ALIYUN_OSS_ACCESS_KEY"`
	ALIYUN_OSS_SECRET_KEY   string `yaml:"ALIYUN_OSS_SECRET_KEY"`
	ALIYUN_OSS_ENDPOINT     string `yaml:"ALIYUN_OSS_ENDPOINT"`
	ALIYUN_OSS_REGION       string `yaml:"ALIYUN_OSS_REGION"`
	ALIYUN_OSS_AUTH_VERSION string `yaml:"ALIYUN_OSS_AUTH_VERSION"`
	ALIYUN_OSS_PATHS        string `yaml:"ALIYUN_OSS_PATHS"`

	TENCENT_COS_BUCKET_NAME string `yaml:"TENCENT_COS_BUCKET_NAME"`
	TENCENT_COS_SECRET_KEY  string `yaml:"TENCENT_COS_SECRET_KEY"`
	TENCENT_COS_SECRET_ID   string `yaml:"TENCENT_COS_SECRET_ID"`
	TENCENT_COS_REGION      string `yaml:"TENCENT_COS_REGION"`
	TENCENT_COS_SCHEME      string `yaml:"TENCENT_COS_SCHEME"`

	HUAWEI_OBS_BUCKET_NAME string `yaml:"HUAWEI_OBS_BUCKET_NAME"`
	HUAWEI_OBS_SECRET_KEY  string `yaml:"HUAWEI_OBS_SECRET_KEY"`
	HUAWEI_OBS_ACCESS_KEY  string `yaml:"HUAWEI_OBS_ACCESS_KEY"`
	HUAWEI_OBS_SERVER      string `yaml:"HUAWEI_OBS_SERVER"`

	OCI_ENDPOINT    string `yaml:"OCI_ENDPOINT"`
	OCI_BUCKET_NAME string `yaml:"OCI_BUCKET_NAME"`
	OCI_ACCESS_KEY  string `yaml:"OCI_ACCESS_KEY"`
	OCI_SECRET_KEY  string `yaml:"OCI_SECRET_KEY"`
	OCI_REGION      string `yaml:"OCI_REGION"`

	VOLCENGINE_TOS_BUCKET_NAME string `yaml:"VOLCENGINE_TOS_BUCKET_NAME"`
	VOLCENGINE_TOS_SECRET_KEY  string `yaml:"VOLCENGINE_TOS_SECRET_KEY"`
	VOLCENGINE_TOS_ACCESS_KEY  string `yaml:"VOLCENGINE_TOS_ACCESS_KEY"`
	VOLCENGINE_TOS_ENDPOINT    string `yaml:"VOLCENGINE_TOS_ENDPOINT"`
	VOLCENGINE_TOS_REGION      string `yaml:"VOLCENGINE_TOS_REGION"`

	VECTOR_STORE string `yaml:"VECTOR_STORE"`

	WEAVIATE_ENDPOINT string `yaml:"WEAVIATE_ENDPOINT"`
	WEAVIATE_API_KEY  string `yaml:"WEAVIATE_API_KEY"`

	QDRANT_URL            string `yaml:"QDRANT_URL"`
	QDRANT_API_KEY        string `yaml:"QDRANT_API_KEY"`
	QDRANT_CLIENT_TIMEOUT string `yaml:"QDRANT_CLIENT_TIMEOUT"`
	QDRANT_GRPC_ENABLED   string `yaml:"QDRANT_GRPC_ENABLED"`
	QDRANT_GRPC_PORT      string `yaml:"QDRANT_GRPC_PORT"`

	MILVUS_URI      string `yaml:"MILVUS_URI"`
	MILVUS_TOKEN    string `yaml:"MILVUS_TOKEN"`
	MILVUS_USER     string `yaml:"MILVUS_USER"`
	MILVUS_PASSWORD string `yaml:"MILVUS_PASSWORD"`

	MYSCALE_HOST       string `yaml:"MYSCALE_HOST"`
	MYSCALE_PORT       string `yaml:"MYSCALE_PORT"`
	MYSCALE_USER       string `yaml:"MYSCALE_USER"`
	MYSCALE_PASSWORD   string `yaml:"MYSCALE_PASSWORD"`
	MYSCALE_DATABASE   string `yaml:"MYSCALE_DATABASE"`
	MYSCALE_FTS_PARAMS string `yaml:"MYSCALE_FTS_PARAMS"`

	RELYT_HOST     string `yaml:"RELYT_HOST"`
	RELYT_PORT     string `yaml:"RELYT_PORT"`
	RELYT_USER     string `yaml:"RELYT_USER"`
	RELYT_PASSWORD string `yaml:"RELYT_PASSWORD"`
	RELYT_DATABASE string `yaml:"RELYT_DATABASE"`

	PGVECTOR_HOST     string `yaml:"PGVECTOR_HOST"`
	PGVECTOR_PORT     string `yaml:"PGVECTOR_PORT"`
	PGVECTOR_USER     string `yaml:"PGVECTOR_USER"`
	PGVECTOR_PASSWORD string `yaml:"PGVECTOR_PASSWORD"`
	PGVECTOR_DATABASE string `yaml:"PGVECTOR_DATABASE"`

	TIDB_VECTOR_HOST     string `yaml:"TIDB_VECTOR_HOST"`
	TIDB_VECTOR_PORT     string `yaml:"TIDB_VECTOR_PORT"`
	TIDB_VECTOR_USER     string `yaml:"TIDB_VECTOR_USER"`
	TIDB_VECTOR_PASSWORD string `yaml:"TIDB_VECTOR_PASSWORD"`
	TIDB_VECTOR_DATABASE string `yaml:"TIDB_VECTOR_DATABASE"`

	ORACLE_HOST     string `yaml:"ORACLE_HOST"`
	ORACLE_PORT     string `yaml:"ORACLE_PORT"`
	ORACLE_USER     string `yaml:"ORACLE_USER"`
	ORACLE_PASSWORD string `yaml:"ORACLE_PASSWORD"`
	ORACLE_DATABASE string `yaml:"ORACLE_DATABASE"`

	CHROMA_HOST             string `yaml:"CHROMA_HOST"`
	CHROMA_PORT             string `yaml:"CHROMA_PORT"`
	CHROMA_TENANT           string `yaml:"CHROMA_TENANT"`
	CHROMA_DATABASE         string `yaml:"CHROMA_DATABASE"`
	CHROMA_AUTH_PROVIDER    string `yaml:"CHROMA_AUTH_PROVIDER"`
	CHROMA_AUTH_CREDENTIALS string `yaml:"CHROMA_AUTH_CREDENTIALS"`

	ELASTICSEARCH_HOST     string `yaml:"ELASTICSEARCH_HOST"`
	ELASTICSEARCH_PORT     string `yaml:"ELASTICSEARCH_PORT"`
	ELASTICSEARCH_USERNAME string `yaml:"ELASTICSEARCH_USERNAME"`
	ELASTICSEARCH_PASSWORD string `yaml:"ELASTICSEARCH_PASSWORD"`
	KIBANA_PORT            string `yaml:"KIBANA_PORT"`

	ANALYTICDB_KEY_ID             string `yaml:"ANALYTICDB_KEY_ID"`
	ANALYTICDB_KEY_SECRET         string `yaml:"ANALYTICDB_KEY_SECRET"`
	ANALYTICDB_REGION_ID          string `yaml:"ANALYTICDB_REGION_ID"`
	ANALYTICDB_INSTANCE_ID        string `yaml:"ANALYTICDB_INSTANCE_ID"`
	ANALYTICDB_ACCOUNT            string `yaml:"ANALYTICDB_ACCOUNT"`
	ANALYTICDB_PASSWORD           string `yaml:"ANALYTICDB_PASSWORD"`
	ANALYTICDB_NAMESPACE          string `yaml:"ANALYTICDB_NAMESPACE"`
	ANALYTICDB_NAMESPACE_PASSWORD string `yaml:"ANALYTICDB_NAMESPACE_PASSWORD"`

	OPENSEARCH_HOST     string `yaml:"OPENSEARCH_HOST"`
	OPENSEARCH_PORT     string `yaml:"OPENSEARCH_PORT"`
	OPENSEARCH_USER     string `yaml:"OPENSEARCH_USER"`
	OPENSEARCH_PASSWORD string `yaml:"OPENSEARCH_PASSWORD"`
	OPENSEARCH_SECURE   string `yaml:"OPENSEARCH_SECURE"`

	TENCENT_VECTOR_DB_URL      string `yaml:"TENCENT_VECTOR_DB_URL"`
	TENCENT_VECTOR_DB_API_KEY  string `yaml:"TENCENT_VECTOR_DB_API_KEY"`
	TENCENT_VECTOR_DB_TIMEOUT  string `yaml:"TENCENT_VECTOR_DB_TIMEOUT"`
	TENCENT_VECTOR_DB_USERNAME string `yaml:"TENCENT_VECTOR_DB_USERNAME"`
	TENCENT_VECTOR_DB_DATABASE string `yaml:"TENCENT_VECTOR_DB_DATABASE"`
	TENCENT_VECTOR_DB_SHARD    string `yaml:"TENCENT_VECTOR_DB_SHARD"`
	TENCENT_VECTOR_DB_REPLICAS string `yaml:"TENCENT_VECTOR_DB_REPLICAS"`

	UPLOAD_FILE_SIZE_LIMIT  string `yaml:"UPLOAD_FILE_SIZE_LIMIT"`
	UPLOAD_FILE_BATCH_LIMIT string `yaml:"UPLOAD_FILE_BATCH_LIMIT"`

	ETL_TYPE                     string `yaml:"ETL_TYPE"`
	UNSTRUCTURED_API_URL         string `yaml:"UNSTRUCTURED_API_URL"`
	MULTIMODAL_SEND_IMAGE_FORMAT string `yaml:"MULTIMODAL_SEND_IMAGE_FORMAT"`
	UPLOAD_IMAGE_FILE_SIZE_LIMIT string `yaml:"UPLOAD_IMAGE_FILE_SIZE_LIMIT"`

	SENTRY_DSN                  string `yaml:"SENTRY_DSN"`
	SENTRY_TRACES_SAMPLE_RATE   string `yaml:"SENTRY_TRACES_SAMPLE_RATE"`
	SENTRY_PROFILES_SAMPLE_RATE string `yaml:"SENTRY_PROFILES_SAMPLE_RATE"`

	NOTION_INTEGRATION_TYPE string `yaml:"NOTION_INTEGRATION_TYPE"`
	NOTION_CLIENT_SECRET    string `yaml:"NOTION_CLIENT_SECRET"`
	NOTION_CLIENT_ID        string `yaml:"NOTION_CLIENT_ID"`
	NOTION_INTERNAL_SECRET  string `yaml:"NOTION_INTERNAL_SECRET"`

	MAIL_TYPE              string `yaml:"MAIL_TYPE"`
	MAIL_DEFAULT_SEND_FROM string `yaml:"MAIL_DEFAULT_SEND_FROM"`

	SMTP_SERVER            string `yaml:"SMTP_SERVER"`
	SMTP_PORT              string `yaml:"SMTP_PORT"`
	SMTP_USERNAME          string `yaml:"SMTP_USERNAME"`
	SMTP_PASSWORD          string `yaml:"SMTP_PASSWORD"`
	SMTP_USE_TLS           string `yaml:"SMTP_USE_TLS"`
	SMTP_OPPORTUNISTIC_TLS string `yaml:"SMTP_OPPORTUNISTIC_TLS"`

	RESEND_API_KEY                          string `yaml:"RESEND_API_KEY"`
	RESEND_API_URL                          string `yaml:"RESEND_API_URL"`
	INDEXING_MAX_SEGMENTATION_TOKENS_LENGTH string `yaml:"INDEXING_MAX_SEGMENTATION_TOKENS_LENGTH"`
	INVITE_EXPIRY_HOURS                     string `yaml:"INVITE_EXPIRY_HOURS"`
	RESET_PASSWORD_TOKEN_EXPIRY_HOURS       string `yaml:"RESET_PASSWORD_TOKEN_EXPIRY_HOURS"`

	CODE_EXECUTION_ENDPOINT       string `yaml:"CODE_EXECUTION_ENDPOINT"`
	CODE_EXECUTION_API_KEY        string `yaml:"CODE_EXECUTION_API_KEY"`
	CODE_MAX_NUMBER               string `yaml:"CODE_MAX_NUMBER"`
	CODE_MIN_NUMBER               string `yaml:"CODE_MIN_NUMBER"`
	CODE_MAX_DEPTH                string `yaml:"CODE_MAX_DEPTH"`
	CODE_MAX_PRECISION            string `yaml:"CODE_MAX_PRECISION"`
	CODE_MAX_STRING_LENGTH        string `yaml:"CODE_MAX_STRING_LENGTH"`
	TEMPLATE_TRANSFORM_MAX_LENGTH string `yaml:"TEMPLATE_TRANSFORM_MAX_LENGTH"`
	CODE_MAX_STRING_ARRAY_LENGTH  string `yaml:"CODE_MAX_STRING_ARRAY_LENGTH"`
	CODE_MAX_OBJECT_ARRAY_LENGTH  string `yaml:"CODE_MAX_OBJECT_ARRAY_LENGTH"`
	CODE_MAX_NUMBER_ARRAY_LENGTH  string `yaml:"CODE_MAX_NUMBER_ARRAY_LENGTH"`

	SSRF_PROXY_HTTP_URL  string `yaml:"SSRF_PROXY_HTTP_URL"`
	SSRF_PROXY_HTTPS_URL string `yaml:"SSRF_PROXY_HTTPS_URL"`
}

type DifyAPI struct {
	Image   string `yaml:"image"`
	Restart string `yaml:"restart"`
	// Environment XSharedEnv `yaml:"environment"`
	Environment DifyEnvirontment `yaml:"environment"`
	DependsOn   []string         `yaml:"depends_on"`
	Volumes     []string         `yaml:"volumes"`
	Networks    []string         `yaml:"networks"`
}

type DifyWorker DifyAPI

type DifyWeb struct {
	Image       string              `yaml:"image"`
	Restart     string              `yaml:"restart"`
	Environment DifyWebEnvirontment `yaml:"environment"`
}

type DifyWebEnvirontment struct {
	CONSOLE_API_URL         string `yaml:"CONSOLE_API_URL"`
	APP_API_URL             string `yaml:"APP_API_URL"`
	SENTRY_DSN              string `yaml:"SENTRY_DSN"`
	NEXT_TELEMETRY_DISABLED string `yaml:"NEXT_TELEMETRY_DISABLED"`
}
