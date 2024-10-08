package Define

var DEFAULT_DOCKER_SHARED_ENV = XSharedEnv{
	LOG_LEVEL:   "${LOG_LEVEL:-INFO}",
	LOG_FILE:    "${LOG_FILE:-}",
	DEBUG:       "${DEBUG:-false}",
	FLASK_DEBUG: "${FLASK_DEBUG:-false}",

	SECRET_KEY:    "${SECRET_KEY:-sk-9f73s3ljTXVcMT3Blb3ljTqtsKiGHXVcMT3BlbkFJLK7U}",
	INIT_PASSWORD: "${INIT_PASSWORD:-}",

	CONSOLE_WEB_URL:      "${CONSOLE_WEB_URL:-}",
	CONSOLE_API_URL:      "${CONSOLE_API_URL:-}",
	SERVICE_API_URL:      "${SERVICE_API_URL:-}",
	APP_WEB_URL:          "${APP_WEB_URL:-}",
	CHECK_UPDATE_URL:     "${CHECK_UPDATE_URL:-https://updates.dify.ai}",
	OPENAI_API_BASE:      "${OPENAI_API_BASE:-https://api.openai.com/v1}",
	FILES_URL:            "${FILES_URL:-}",
	FILES_ACCESS_TIMEOUT: "${FILES_ACCESS_TIMEOUT:-300}",

	APP_MAX_ACTIVE_REQUESTS: "${APP_MAX_ACTIVE_REQUESTS:-0}",
	MIGRATION_ENABLED:       "${MIGRATION_ENABLED:-true}",
	DEPLOY_ENV:              "${DEPLOY_ENV:-PRODUCTION}",
	DIFY_BIND_ADDRESS:       "${DIFY_BIND_ADDRESS:-0.0.0.0}",
	DIFY_PORT:               "${DIFY_PORT:-5001}",

	SERVER_WORKER_AMOUNT: "${SERVER_WORKER_AMOUNT:-}",
	SERVER_WORKER_CLASS:  "${SERVER_WORKER_CLASS:-}",
	CELERY_WORKER_CLASS:  "${CELERY_WORKER_CLASS:-}",
	GUNICORN_TIMEOUT:     "${GUNICORN_TIMEOUT:-360}",
	CELERY_WORKER_AMOUNT: "${CELERY_WORKER_AMOUNT:-}",
	CELERY_AUTO_SCALE:    "${CELERY_AUTO_SCALE:-false}",
	CELERY_MAX_WORKERS:   "${CELERY_MAX_WORKERS:-}",
	CELERY_MIN_WORKERS:   "${CELERY_MIN_WORKERS:-}",

	API_TOOL_DEFAULT_CONNECT_TIMEOUT: "${API_TOOL_DEFAULT_CONNECT_TIMEOUT:-10}",
	API_TOOL_DEFAULT_READ_TIMEOUT:    "${API_TOOL_DEFAULT_READ_TIMEOUT:-60}",

	DB_USERNAME: "${DB_USERNAME:-postgres}",
	DB_PASSWORD: "${DB_PASSWORD:-difyai123456}",
	DB_HOST:     "${DB_HOST:-db}",
	DB_PORT:     "${DB_PORT:-5432}",
	DB_DATABASE: "${DB_DATABASE:-dify}",

	SQLALCHEMY_POOL_SIZE:    "${SQLALCHEMY_POOL_SIZE:-30}",
	SQLALCHEMY_POOL_RECYCLE: "${SQLALCHEMY_POOL_RECYCLE:-3600}",
	SQLALCHEMY_ECHO:         "${SQLALCHEMY_ECHO:-false}",

	REDIS_HOST:                    "${REDIS_HOST:-redis}",
	REDIS_PORT:                    "${REDIS_PORT:-6379}",
	REDIS_USERNAME:                "${REDIS_USERNAME:-}",
	REDIS_PASSWORD:                "${REDIS_PASSWORD:-difyai123456}",
	REDIS_USE_SSL:                 "${REDIS_USE_SSL:-false}",
	REDIS_DB:                      0,
	REDIS_USE_SENTINEL:            "${REDIS_USE_SENTINEL:-false}",
	REDIS_SENTINELS:               "${REDIS_SENTINELS:-}",
	REDIS_SENTINEL_SERVICE_NAME:   "${REDIS_SENTINEL_SERVICE_NAME:-}",
	REDIS_SENTINEL_USERNAME:       "${REDIS_SENTINEL_USERNAME:-}",
	REDIS_SENTINEL_PASSWORD:       "${REDIS_SENTINEL_PASSWORD:-}",
	REDIS_SENTINEL_SOCKET_TIMEOUT: "${REDIS_SENTINEL_SOCKET_TIMEOUT:-0.1}",

	CELERY_BROKER_URL: "${CELERY_BROKER_URL:-redis://:difyai123456@redis:6379/1}",
	BROKER_USE_SSL:    "${BROKER_USE_SSL:-false}",

	CELERY_USE_SENTINEL:            "${CELERY_USE_SENTINEL:-false}",
	CELERY_SENTINEL_MASTER_NAME:    "${CELERY_SENTINEL_MASTER_NAME:-}",
	CELERY_SENTINEL_SOCKET_TIMEOUT: "${CELERY_SENTINEL_SOCKET_TIMEOUT:-0.1}",

	WEB_API_CORS_ALLOW_ORIGINS: "${WEB_API_CORS_ALLOW_ORIGINS:-*}",
	CONSOLE_CORS_ALLOW_ORIGINS: "${CONSOLE_CORS_ALLOW_ORIGINS:-*}",

	STORAGE_TYPE: "${STORAGE_TYPE:-local}",

	STORAGE_LOCAL_PATH: "storage",

	S3_USE_AWS_MANAGED_IAM: "${S3_USE_AWS_MANAGED_IAM:-false}",
	S3_ENDPOINT:            "${S3_ENDPOINT:-}",
	S3_BUCKET_NAME:         "${S3_BUCKET_NAME:-}",
	S3_ACCESS_KEY:          "${S3_ACCESS_KEY:-}",
	S3_SECRET_KEY:          "${S3_SECRET_KEY:-}",
	S3_REGION:              "${S3_REGION:-us-east-1}",

	AZURE_BLOB_ACCOUNT_NAME:   "${AZURE_BLOB_ACCOUNT_NAME:-}",
	AZURE_BLOB_ACCOUNT_KEY:    "${AZURE_BLOB_ACCOUNT_KEY:-}",
	AZURE_BLOB_CONTAINER_NAME: "${AZURE_BLOB_CONTAINER_NAME:-}",
	AZURE_BLOB_ACCOUNT_URL:    "${AZURE_BLOB_ACCOUNT_URL:-}",

	GOOGLE_STORAGE_BUCKET_NAME:                 "${GOOGLE_STORAGE_BUCKET_NAME:-}",
	GOOGLE_STORAGE_SERVICE_ACCOUNT_JSON_BASE64: "${GOOGLE_STORAGE_SERVICE_ACCOUNT_JSON_BASE64:-}",

	ALIYUN_OSS_BUCKET_NAME:  "${ALIYUN_OSS_BUCKET_NAME:-}",
	ALIYUN_OSS_ACCESS_KEY:   "${ALIYUN_OSS_ACCESS_KEY:-}",
	ALIYUN_OSS_SECRET_KEY:   "${ALIYUN_OSS_SECRET_KEY:-}",
	ALIYUN_OSS_ENDPOINT:     "${ALIYUN_OSS_ENDPOINT:-}",
	ALIYUN_OSS_REGION:       "${ALIYUN_OSS_REGION:-}",
	ALIYUN_OSS_AUTH_VERSION: "${ALIYUN_OSS_AUTH_VERSION:-v4}",
	ALIYUN_OSS_PATHS:        "${ALIYUN_OSS_PATH:-}",

	TENCENT_COS_BUCKET_NAME: "${TENCENT_COS_BUCKET_NAME:-}",
	TENCENT_COS_SECRET_KEY:  "${TENCENT_COS_SECRET_KEY:-}",
	TENCENT_COS_SECRET_ID:   "${TENCENT_COS_SECRET_ID:-}",
	TENCENT_COS_REGION:      "${TENCENT_COS_REGION:-}",
	TENCENT_COS_SCHEME:      "${TENCENT_COS_SCHEME:-}",

	HUAWEI_OBS_BUCKET_NAME: "${HUAWEI_OBS_BUCKET_NAME:-}",
	HUAWEI_OBS_SECRET_KEY:  "${HUAWEI_OBS_SECRET_KEY:-}",
	HUAWEI_OBS_ACCESS_KEY:  "${HUAWEI_OBS_ACCESS_KEY:-}",
	HUAWEI_OBS_SERVER:      "${HUAWEI_OBS_SERVER:-}",

	OCI_ENDPOINT:    "${OCI_ENDPOINT:-}",
	OCI_BUCKET_NAME: "${OCI_BUCKET_NAME:-}",
	OCI_ACCESS_KEY:  "${OCI_ACCESS_KEY:-}",
	OCI_SECRET_KEY:  "${OCI_SECRET_KEY:-}",
	OCI_REGION:      "${OCI_REGION:-}",

	VOLCENGINE_TOS_BUCKET_NAME: "${VOLCENGINE_TOS_BUCKET_NAME:-}",
	VOLCENGINE_TOS_SECRET_KEY:  "${VOLCENGINE_TOS_SECRET_KEY:-}",
	VOLCENGINE_TOS_ACCESS_KEY:  "${VOLCENGINE_TOS_ACCESS_KEY:-}",
	VOLCENGINE_TOS_ENDPOINT:    "${VOLCENGINE_TOS_ENDPOINT:-}",
	VOLCENGINE_TOS_REGION:      "${VOLCENGINE_TOS_REGION:-}",

	VECTOR_STORE: "${VECTOR_STORE:-weaviate}",

	WEAVIATE_ENDPOINT: "${WEAVIATE_ENDPOINT:-http://weaviate:8080}",
	WEAVIATE_API_KEY:  "${WEAVIATE_API_KEY:-WVF5YThaHlkYwhGUSmCRgsX3tD5ngdN8pkih}",

	QDRANT_URL:            "${QDRANT_URL:-http://qdrant:6333}",
	QDRANT_API_KEY:        "${QDRANT_API_KEY:-difyai123456}",
	QDRANT_CLIENT_TIMEOUT: "${QDRANT_CLIENT_TIMEOUT:-20}",
	QDRANT_GRPC_ENABLED:   "${QDRANT_GRPC_ENABLED:-false}",
	QDRANT_GRPC_PORT:      "${QDRANT_GRPC_PORT:-6334}",

	MILVUS_URI:      "${MILVUS_URI:-http://127.0.0.1:19530}",
	MILVUS_TOKEN:    "${MILVUS_TOKEN:-}",
	MILVUS_USER:     "${MILVUS_USER:-root}",
	MILVUS_PASSWORD: "${MILVUS_PASSWORD:-Milvus}",

	MYSCALE_HOST:       "${MYSCALE_HOST:-myscale}",
	MYSCALE_PORT:       "${MYSCALE_PORT:-8123}",
	MYSCALE_USER:       "${MYSCALE_USER:-default}",
	MYSCALE_PASSWORD:   "${MYSCALE_PASSWORD:-}",
	MYSCALE_DATABASE:   "${MYSCALE_DATABASE:-dify}",
	MYSCALE_FTS_PARAMS: "${MYSCALE_FTS_PARAMS:-}",

	RELYT_HOST:     "${RELYT_HOST:-db}",
	RELYT_PORT:     "${RELYT_PORT:-5432}",
	RELYT_USER:     "${RELYT_USER:-postgres}",
	RELYT_PASSWORD: "${RELYT_PASSWORD:-difyai123456}",
	RELYT_DATABASE: "${RELYT_DATABASE:-postgres}",

	PGVECTOR_HOST:     "${PGVECTOR_HOST:-pgvector}",
	PGVECTOR_PORT:     "${PGVECTOR_PORT:-5432}",
	PGVECTOR_USER:     "${PGVECTOR_USER:-postgres}",
	PGVECTOR_PASSWORD: "${PGVECTOR_PASSWORD:-difyai123456}",
	PGVECTOR_DATABASE: "${PGVECTOR_DATABASE:-dify}",

	TIDB_VECTOR_HOST:     "${TIDB_VECTOR_HOST:-tidb}",
	TIDB_VECTOR_PORT:     "${TIDB_VECTOR_PORT:-4000}",
	TIDB_VECTOR_USER:     "${TIDB_VECTOR_USER:-}",
	TIDB_VECTOR_PASSWORD: "${TIDB_VECTOR_PASSWORD:-}",
	TIDB_VECTOR_DATABASE: "${TIDB_VECTOR_DATABASE:-dify}",

	ORACLE_HOST:     "${ORACLE_HOST:-oracle}",
	ORACLE_PORT:     "${ORACLE_PORT:-1521}",
	ORACLE_USER:     "${ORACLE_USER:-dify}",
	ORACLE_PASSWORD: "${ORACLE_PASSWORD:-dify}",
	ORACLE_DATABASE: "${ORACLE_DATABASE:-FREEPDB1}",

	CHROMA_HOST:             "${CHROMA_HOST:-127.0.0.1}",
	CHROMA_PORT:             "${CHROMA_PORT:-8000}",
	CHROMA_TENANT:           "${CHROMA_TENANT:-default_tenant}",
	CHROMA_DATABASE:         "${CHROMA_DATABASE:-default_database}",
	CHROMA_AUTH_PROVIDER:    "${CHROMA_AUTH_PROVIDER:-chromadb.auth.token_authn.TokenAuthClientProvider}",
	CHROMA_AUTH_CREDENTIALS: "${CHROMA_AUTH_CREDENTIALS:-}",

	ELASTICSEARCH_HOST:     "${ELASTICSEARCH_HOST:-0.0.0.0}",
	ELASTICSEARCH_PORT:     "${ELASTICSEARCH_PORT:-9200}",
	ELASTICSEARCH_USERNAME: "${ELASTICSEARCH_USERNAME:-elastic}",
	ELASTICSEARCH_PASSWORD: "${ELASTICSEARCH_PASSWORD:-elastic}",
	KIBANA_PORT:            "${KIBANA_PORT:-5601}",

	ANALYTICDB_KEY_ID:             "${ANALYTICDB_KEY_ID:-}",
	ANALYTICDB_KEY_SECRET:         "${ANALYTICDB_KEY_SECRET:-}",
	ANALYTICDB_REGION_ID:          "${ANALYTICDB_REGION_ID:-}",
	ANALYTICDB_INSTANCE_ID:        "${ANALYTICDB_INSTANCE_ID:-}",
	ANALYTICDB_ACCOUNT:            "${ANALYTICDB_ACCOUNT:-}",
	ANALYTICDB_PASSWORD:           "${ANALYTICDB_PASSWORD:-}",
	ANALYTICDB_NAMESPACE:          "${ANALYTICDB_NAMESPACE:-dify}",
	ANALYTICDB_NAMESPACE_PASSWORD: "${ANALYTICDB_NAMESPACE_PASSWORD:-}",

	OPENSEARCH_HOST:     "${OPENSEARCH_HOST:-opensearch}",
	OPENSEARCH_PORT:     "${OPENSEARCH_PORT:-9200}",
	OPENSEARCH_USER:     "${OPENSEARCH_USER:-admin}",
	OPENSEARCH_PASSWORD: "${OPENSEARCH_PASSWORD:-admin}",
	OPENSEARCH_SECURE:   "${OPENSEARCH_SECURE:-true}",

	TENCENT_VECTOR_DB_URL:      "${TENCENT_VECTOR_DB_URL:-http://127.0.0.1}",
	TENCENT_VECTOR_DB_API_KEY:  "${TENCENT_VECTOR_DB_API_KEY:-dify}",
	TENCENT_VECTOR_DB_TIMEOUT:  "${TENCENT_VECTOR_DB_TIMEOUT:-30}",
	TENCENT_VECTOR_DB_USERNAME: "${TENCENT_VECTOR_DB_USERNAME:-dify}",
	TENCENT_VECTOR_DB_DATABASE: "${TENCENT_VECTOR_DB_DATABASE:-dify}",
	TENCENT_VECTOR_DB_SHARD:    "${TENCENT_VECTOR_DB_SHARD:-1}",
	TENCENT_VECTOR_DB_REPLICAS: "${TENCENT_VECTOR_DB_REPLICAS:-2}",

	UPLOAD_FILE_SIZE_LIMIT:  "${UPLOAD_FILE_SIZE_LIMIT:-15}",
	UPLOAD_FILE_BATCH_LIMIT: "${UPLOAD_FILE_BATCH_LIMIT:-5}",

	ETL_TYPE:                     "${ETL_TYPE:-dify}",
	UNSTRUCTURED_API_URL:         "${UNSTRUCTURED_API_URL:-}",
	MULTIMODAL_SEND_IMAGE_FORMAT: "${MULTIMODAL_SEND_IMAGE_FORMAT:-base64}",
	UPLOAD_IMAGE_FILE_SIZE_LIMIT: "${UPLOAD_IMAGE_FILE_SIZE_LIMIT:-10}",

	SENTRY_DSN:                  "${API_SENTRY_DSN:-}",
	SENTRY_TRACES_SAMPLE_RATE:   "${API_SENTRY_TRACES_SAMPLE_RATE:-1.0}",
	SENTRY_PROFILES_SAMPLE_RATE: "${API_SENTRY_PROFILES_SAMPLE_RATE:-1.0}",

	NOTION_INTEGRATION_TYPE: "${NOTION_INTEGRATION_TYPE:-public}",
	NOTION_CLIENT_SECRET:    "${NOTION_CLIENT_SECRET:-}",
	NOTION_CLIENT_ID:        "${NOTION_CLIENT_ID:-}",
	NOTION_INTERNAL_SECRET:  "${NOTION_INTERNAL_SECRET:-}",

	MAIL_TYPE:              "${MAIL_TYPE:-resend}",
	MAIL_DEFAULT_SEND_FROM: "${MAIL_DEFAULT_SEND_FROM:-}",
	SMTP_SERVER:            "${SMTP_SERVER:-}",
	SMTP_PORT:              "${SMTP_PORT:-465}",
	SMTP_USERNAME:          "${SMTP_USERNAME:-}",
	SMTP_PASSWORD:          "${SMTP_PASSWORD:-}",
	SMTP_USE_TLS:           "${SMTP_USE_TLS:-true}",
	SMTP_OPPORTUNISTIC_TLS: "${SMTP_OPPORTUNISTIC_TLS:-false}",

	RESEND_API_KEY:                          "${RESEND_API_KEY:-your-resend-api-key}",
	RESEND_API_URL:                          "https://api.resend.com",
	INDEXING_MAX_SEGMENTATION_TOKENS_LENGTH: "${INDEXING_MAX_SEGMENTATION_TOKENS_LENGTH:-1000}",
	INVITE_EXPIRY_HOURS:                     "${INVITE_EXPIRY_HOURS:-72}",
	RESET_PASSWORD_TOKEN_EXPIRY_HOURS:       "${RESET_PASSWORD_TOKEN_EXPIRY_HOURS:-24}",

	CODE_EXECUTION_ENDPOINT:       "${CODE_EXECUTION_ENDPOINT:-http://sandbox:8194}",
	CODE_EXECUTION_API_KEY:        "${SANDBOX_API_KEY:-dify-sandbox}",
	CODE_MAX_NUMBER:               "${CODE_MAX_NUMBER:-9223372036854775807}",
	CODE_MIN_NUMBER:               "${CODE_MIN_NUMBER:--9223372036854775808}",
	CODE_MAX_DEPTH:                "${CODE_MAX_DEPTH:-5}",
	CODE_MAX_PRECISION:            "${CODE_MAX_PRECISION:-20}",
	CODE_MAX_STRING_LENGTH:        "${CODE_MAX_STRING_LENGTH:-80000}",
	TEMPLATE_TRANSFORM_MAX_LENGTH: "${TEMPLATE_TRANSFORM_MAX_LENGTH:-80000}",
	CODE_MAX_STRING_ARRAY_LENGTH:  "${CODE_MAX_STRING_ARRAY_LENGTH:-30}",
	CODE_MAX_OBJECT_ARRAY_LENGTH:  "${CODE_MAX_OBJECT_ARRAY_LENGTH:-30}",
	CODE_MAX_NUMBER_ARRAY_LENGTH:  "${CODE_MAX_NUMBER_ARRAY_LENGTH:-1000}",

	SSRF_PROXY_HTTP_URL:  "${SSRF_PROXY_HTTP_URL:-http://ssrf_proxy:3128}",
	SSRF_PROXY_HTTPS_URL: "${SSRF_PROXY_HTTPS_URL:-http://ssrf_proxy:3128}",
}
