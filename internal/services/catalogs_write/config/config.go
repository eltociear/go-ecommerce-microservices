package config

import (
    "flag"
    "fmt"
    "os"
    "path/filepath"
    "runtime"
    "strconv"
    "strings"

    "emperror.dev/errors"
    "github.com/caarlos0/env/v6"
    "github.com/spf13/viper"

    "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/constants"
    "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/eventstroredb"
    gormPostgres "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/gorm_postgres"
    "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/grpc"
    customEcho "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/http/custom_echo"
    "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/logger"
    "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/otel"
    "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/otel/metrics"
    postgres "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/postgres_pgx"
    "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/rabbitmq/config"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "catalogs_write write microservice config path")
}

type Config struct {
	DeliveryType      string                         `mapstructure:"deliveryType" env:"DeliveryType"`
	ServiceName       string                         `mapstructure:"serviceName" env:"ServiceName"`
	Logger            *logger.LogConfig              `mapstructure:"logger" envPrefix:"Logger_"`
	GRPC              *grpc.GrpcConfig               `mapstructure:"grpc" envPrefix:"GRPC_"`
	Http              *customEcho.EchoHttpConfig     `mapstructure:"http" envPrefix:"Http_"`
	Context           Context                        `mapstructure:"context" envPrefix:"Context_"`
	Postgresql        *postgres.Config               `mapstructure:"postgres" envPrefix:"Postgresql_"`
	GormPostgres      *gormPostgres.GormConfig       `mapstructure:"gormPostgres" envPrefix:"GormPostgres_"`
	RabbitMQ          *config.RabbitMQConfig         `mapstructure:"rabbitmq" envPrefix:"RabbitMQ_"`
	OTel              *otel.OpenTelemetryConfig      `mapstructure:"otel" envPrefix:"OTel_"`
	OTelMetricsConfig *metrics.OTelMetricsConfig     `mapstructure:"otelMetrics" envPrefix:"OTelMetrics_"`
	EventStoreConfig  eventstroredb.EventStoreConfig `mapstructure:"eventStoreConfig" envPrefix:"EventStoreConfig_"`
}

type Context struct {
	Timeout int `mapstructure:"timeout" env:"Timeout"`
}

func InitConfig(environment string) (*Config, error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv(constants.ConfigPath)
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			//https://stackoverflow.com/questions/31873396/is-it-possible-to-get-the-current-root-of-package-structure-as-a-string-in-golan
			//https://stackoverflow.com/questions/18537257/how-to-get-the-directory-of-the-currently-running-file
			d, err := dirname()
			if err != nil {
				return nil, err
			}

			configPath = d
		}
	}

	cfg := &Config{}

	viper.SetConfigName(fmt.Sprintf("config.%s", environment))
	viper.AddConfigPath(configPath)
	viper.SetConfigType(constants.Json)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.WrapIf(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.WrapIf(err, "viper.Unmarshal")
	}

	if err := env.Parse(cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	grpcPort := os.Getenv(constants.GrpcPort)
	if grpcPort != "" {
		cfg.GRPC.Port = grpcPort
	}

	postgresHost := os.Getenv(constants.PostgresqlHost)
	if postgresHost != "" {
		cfg.Postgresql.Host = postgresHost
	}
	postgresPort := os.Getenv(constants.PostgresqlPort)
	if postgresPort != "" {
		port, err := strconv.Atoi(postgresPort)
		if err != nil {
			return nil, err
		}
		cfg.Postgresql.Port = port
	}

	jaegerPort := os.Getenv(constants.JaegerPort)
	if jaegerPort != "" {
		cfg.OTel.JaegerExporterConfig.AgentPort = jaegerPort
	}

	jaegerHost := os.Getenv(constants.JaegerHost)
	if jaegerHost != "" {
		cfg.OTel.JaegerExporterConfig.AgentHost = jaegerHost
	}

	return cfg, nil
}

func (cfg *Config) GetMicroserviceNameUpper() string {
	return strings.ToUpper(cfg.ServiceName)
}

func (cfg *Config) GetMicroserviceName() string {
	return cfg.ServiceName
}

func filename() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("unable to get the current filename")
	}
	return filename, nil
}

func dirname() (string, error) {
	filename, err := filename()
	if err != nil {
		return "", err
	}
	return filepath.Dir(filename), nil
}