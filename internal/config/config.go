package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App        AppConfig        `yaml:"app"`
	Database   DatabaseConfig   `yaml:"database"`
	Redis      RedisConfig      `yaml:"redis"`
	API        APIConfig        `yaml:"api"`
	RateLimit  RateLimitConfig  `yaml:"rate_limit"`
	Pagination PaginationConfig `yaml:"pagination"`
	Logging    LoggingConfig    `yaml:"logging"`
	Security   SecurityConfig   `yaml:"security"`
	Monitoring MonitoringConfig `yaml:"monitoring"`
}

type AppConfig struct {
	Name     string `yaml:"name"`
	Version  string `yaml:"version"`
	Env      string `yaml:"env"`
	Port     int    `yaml:"port"`
	Timezone string `yaml:"timezone"`
}

type DatabaseConfig struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	Database        string `yaml:"database"`
	Charset         string `yaml:"charset"`
	ParseTime       bool   `yaml:"parse_time"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
}

type RedisConfig struct {
	Enabled  bool             `yaml:"enabled"`
	Host     string           `yaml:"host"`
	Port     int              `yaml:"port"`
	Password string           `yaml:"password"`
	DB       int              `yaml:"db"`
	CacheTTL RedisCacheTTL    `yaml:"cache_ttl"`
}

type RedisCacheTTL struct {
	SatpenList   int `yaml:"satpen_list"`
	SatpenDetail int `yaml:"satpen_detail"`
	Statistics   int `yaml:"statistics"`
	MasterData   int `yaml:"master_data"`
}

type APIConfig struct {
	BasePath       string   `yaml:"base_path"`
	RequestTimeout int      `yaml:"request_timeout"`
	MaxRequestSize int      `yaml:"max_request_size"`
	AllowedOrigins []string `yaml:"allowed_origins"`
}

type RateLimitConfig struct {
	Enabled    bool                 `yaml:"enabled"`
	Satpen     RateLimitRule        `yaml:"satpen"`
	Statistics RateLimitRule        `yaml:"statistics"`
}

type RateLimitRule struct {
	Requests int `yaml:"requests"`
	Window   int `yaml:"window"`
}

type PaginationConfig struct {
	DefaultPage  int `yaml:"default_page"`
	DefaultLimit int `yaml:"default_limit"`
	MaxLimit     int `yaml:"max_limit"`
}

type LoggingConfig struct {
	Level      string `yaml:"level"`
	Format     string `yaml:"format"`
	Output     string `yaml:"output"`
	FilePath   string `yaml:"file_path"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
	Compress   bool   `yaml:"compress"`
}

type SecurityConfig struct {
	CorsEnabled    bool     `yaml:"cors_enabled"`
	TrustedProxies []string `yaml:"trusted_proxies"`
}

type MonitoringConfig struct {
	Enabled         bool   `yaml:"enabled"`
	MetricsPath     string `yaml:"metrics_path"`
	HealthCheckPath string `yaml:"health_check_path"`
}

var GlobalConfig *Config

// LoadConfig loads configuration from config.yaml
func LoadConfig(configPath string) (*Config, error) {
	// Read config file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Parse YAML
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	// Override with environment variables if present
	overrideWithEnv(&config)

	GlobalConfig = &config
	return &config, nil
}

// overrideWithEnv overrides config with environment variables
func overrideWithEnv(config *Config) {
	if env := os.Getenv("APP_ENV"); env != "" {
		config.App.Env = env
	}
	if port := os.Getenv("APP_PORT"); port != "" {
		fmt.Sscanf(port, "%d", &config.App.Port)
	}
	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		config.Database.Host = dbHost
	}
	if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
		fmt.Sscanf(dbPort, "%d", &config.Database.Port)
	}
	if dbUser := os.Getenv("DB_USERNAME"); dbUser != "" {
		config.Database.Username = dbUser
	}
	if dbPass := os.Getenv("DB_PASSWORD"); dbPass != "" {
		config.Database.Password = dbPass
	}
	if dbName := os.Getenv("DB_DATABASE"); dbName != "" {
		config.Database.Database = dbName
	}
	if redisHost := os.Getenv("REDIS_HOST"); redisHost != "" {
		config.Redis.Host = redisHost
	}
	if redisPass := os.Getenv("REDIS_PASSWORD"); redisPass != "" {
		config.Redis.Password = redisPass
	}
}

// GetDSN returns database connection string
func (c *Config) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=Local",
		c.Database.Username,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Database,
		c.Database.Charset,
		c.Database.ParseTime,
	)
}

// GetRedisAddr returns Redis address
func (c *Config) GetRedisAddr() string {
	return fmt.Sprintf("%s:%d", c.Redis.Host, c.Redis.Port)
}

// IsDevelopment checks if app is in development mode
func (c *Config) IsDevelopment() bool {
	return c.App.Env == "development"
}

// IsProduction checks if app is in production mode
func (c *Config) IsProduction() bool {
	return c.App.Env == "production"
}

// Init initializes config from default path
func Init() {
	cfg, err := LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	GlobalConfig = cfg
	log.Printf("Config loaded successfully: %s v%s [%s]",
		cfg.App.Name, cfg.App.Version, cfg.App.Env)
}
