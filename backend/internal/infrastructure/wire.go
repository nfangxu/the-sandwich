package infrastructure

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/the-sandwich/backend/internal/application/auth"
	"github.com/the-sandwich/backend/internal/application/game"
	"github.com/the-sandwich/backend/internal/application/matchmaking"
	"github.com/the-sandwich/backend/internal/domain/user"
	"github.com/the-sandwich/backend/internal/infrastructure/cache"
	"github.com/the-sandwich/backend/internal/infrastructure/persistence/mysql"
	"github.com/the-sandwich/backend/internal/infrastructure/persistence/sqlite"
)

// AppConfig holds all configuration for the application
type AppConfig struct {
	Database DatabaseConfig
	Redis    RedisConfig
	Auth     AuthConfig
	App      ServerConfig
}

type DatabaseConfig struct {
	Driver string `mapstructure:"driver"`
	DSN    string `mapstructure:"dsn"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type AuthConfig struct {
	JWTsecret string `mapstructure:"jwt_secret"`
}

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

// LoadConfig loads configuration from the specified file path
func LoadConfig(path string) (AppConfig, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return AppConfig{}, err
	}

	var cfg AppConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		return AppConfig{}, err
	}

	return cfg, nil
}

// InitializeApp initializes the application with all dependencies
func InitializeApp(cfg AppConfig) (*App, error) {
	// Initialize database based on driver
	var dbRepo user.UserRepository
	switch cfg.Database.Driver {
	case "mysql":
		db, err := mysql.NewDB(mysql.Config{DSN: cfg.Database.DSN})
		if err != nil {
			return nil, err
		}
		dbRepo = mysql.NewUserRepository(db)
	case "sqlite":
		db, err := sqlite.NewDB(sqlite.Config{DSN: cfg.Database.DSN})
		if err != nil {
			return nil, err
		}
		dbRepo = sqlite.NewUserRepository(db)
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", cfg.Database.Driver)
	}

	// Initialize Redis
	redisClient, err := cache.NewRedis(cache.Config{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	if err != nil {
		return nil, err
	}

	// Initialize repositories
	gameRepo := cache.NewGameRepository(redisClient)
	matchmakingRepo := cache.NewMatchmakingRepository(redisClient)

	// Initialize application services
	authSvc := auth.NewAuthService(dbRepo, cfg.Auth.JWTsecret)
	gameSvc := game.NewGameService(gameRepo)
	matchmakingSvc := matchmaking.NewMatchmakingService(matchmakingRepo)

	return &App{
		AuthSvc:        authSvc,
		GameSvc:        gameSvc,
		MatchmakingSvc: matchmakingSvc,
	}, nil
}

type App struct {
	AuthSvc        *auth.Service
	GameSvc        *game.GameService
	MatchmakingSvc *matchmaking.MatchmakingService
}
