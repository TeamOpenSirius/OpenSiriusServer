package config

import (
	"strings"
	"time"
)

// Config holds application configuration.
type Config struct {
	Env        string        `yaml:"env"`
	LogLevel   string        `yaml:"log_level"`
	LogPath    string        `yaml:"log_path"`
	Addr       string        `yaml:"addr"`
	JWTSecret  string        `yaml:"jwt_secret"`
	JWTIssuer  string        `yaml:"jwt_issuer"`
	AccessTTL  time.Duration `yaml:"access_ttl"`
	RefreshTTL time.Duration `yaml:"refresh_ttl"`
	Cache      CacheConfig   `yaml:"cache"`
	Game       GameConfig    `yaml:"game"`
	Debug      DebugConfig   `yaml:"debug"`
}

type CacheConfig struct {
	Enabled    bool          `yaml:"enabled"`
	MaxEntries int           `yaml:"max_entries"`
	TTL        time.Duration `yaml:"ttl"`
}

type GameConfig struct {
	BaseURL                    string              `yaml:"base_url"`
	RealtimeURL                string              `yaml:"realtime_url"`
	ApplicationVersion         string              `yaml:"application_version"`
	AssetVersion               string              `yaml:"asset_version"`
	ApiAssetsVersion           string              `yaml:"api_assets_version"`
	AssetRoot                  string              `yaml:"asset_root"`
	AssetURL                   string              `yaml:"asset_url"`
	StaticContentURL           string              `yaml:"static_content_url"`
	MasterDataURL              string              `yaml:"master_data_url"`
	MasterDataBackendURL       string              `yaml:"master_data_backend_url"`
	MasterDataURI              string              `yaml:"master_data_uri"`
	MasterDataVersion          string              `yaml:"master_data_version"`
	MasterDataPublishTimestamp int64               `yaml:"master_data_publish_timestamp"`
	MasterDataRoot             string              `yaml:"master_data_root"`
	MasterDecodedJSONRoot      string              `yaml:"master_decoded_json_root"`
	UserDataFixture            string              `yaml:"user_data_fixture"`
	UserDataSeedJSON           string              `yaml:"user_data_seed_json"`
	UserDataSeedMode           string              `yaml:"user_data_seed_mode"`
	GeneratedUser              GeneratedUserConfig `yaml:"generated_user"`
	GeneratedUserProfile       string              `yaml:"generated_user_profile"`
	GeneratedIncludeUnions     []int64             `yaml:"generated_include_unions"`
	GeneratedExcludeUnions     []int64             `yaml:"generated_exclude_unions"`
	EnableLocalNotification    bool                `yaml:"enable_local_notification"`
	UseRawFixtureUserData      bool                `yaml:"use_raw_fixture_user_data"`
	SeedCharacterBaseMasterIDs []int64             `yaml:"seed_character_base_master_ids"`
	SeedCharacterMasterIDs     []int64             `yaml:"seed_character_master_ids"`
}

type GeneratedUserConfig struct {
	PlayerRank     int64 `yaml:"player_rank"`
	CurrentStamina int64 `yaml:"current_stamina"`
	Coin           int64 `yaml:"coin"`
	FreeJewel      int64 `yaml:"free_jewel"`
	PaidJewel      int64 `yaml:"paid_jewel"`
	TutorialStatus int64 `yaml:"tutorial_status"`
}

type DebugConfig struct {
	DumpAPIResponse bool   `yaml:"dump_api_response"`
	DumpDir         string `yaml:"dump_dir"`
}

// Load loads configuration from data/config.yml, creating it with defaults if missing.
func Load() (Config, error) {
	return LoadFrom("data/config.yml")
}

// LoadFrom loads configuration from the given path, creating it with defaults if missing.
func LoadFrom(path string) (Config, error) {
	return LoadYAML(path, Default(), (*Config).applyLogLevelDefault)
}

// Save writes the configuration to a YAML file.
func (c *Config) Save(path string) error {
	return SaveYAML(path, c)
}

// Default returns the default configuration.
func Default() Config {
	cfg := Config{
		Env:        "development",
		LogLevel:   "",
		LogPath:    "data/logs/app.log",
		Addr:       ":8080",
		JWTSecret:  "change-me",
		JWTIssuer:  "ymst_server",
		AccessTTL:  mustParseDuration("30m"),
		RefreshTTL: mustParseDuration("720h"),
		Cache: CacheConfig{
			Enabled:    true,
			MaxEntries: 512,
			TTL:        mustParseDuration("30s"),
		},
		Game: GameConfig{
			BaseURL:                    "http://127.0.0.1:8080",
			RealtimeURL:                "http://127.0.0.1:8081",
			ApplicationVersion:         "2.30.1",
			AssetVersion:               "1.91.0",
			ApiAssetsVersion:           "1.12.0",
			AssetRoot:                  "data/assets/production",
			AssetURL:                   "",
			StaticContentURL:           "",
			MasterDataURL:              "",
			MasterDataURI:              "2026-05-09/mastermemory_1778239615_1778239615.db",
			MasterDataVersion:          "1778239615_1778239615",
			MasterDataPublishTimestamp: 1778239615,
			MasterDataRoot:             "data/master-data/production",
			MasterDecodedJSONRoot:      "../ymst_masterdata/data/master/decoded",
			UserDataFixture:            "references/captured/2026-05-16T06-13-39-729Z_res_lb-api.wds-stellarium.com_api_data_user.bin",
			UserDataSeedJSON:           "references/captured/2026-05-16T06-13-39-729Z_res_lb-api.wds-stellarium.com_api_data_user.bin.api.json",
			UserDataSeedMode:           "captured_json",
			GeneratedUser:              GeneratedUserConfig{PlayerRank: 1, CurrentStamina: 100, Coin: 100000, FreeJewel: 3000, PaidJewel: 0, TutorialStatus: 99},
			EnableLocalNotification:    true,
			UseRawFixtureUserData:      false,
			SeedCharacterBaseMasterIDs: []int64{1, 2, 3, 4, 5},
			SeedCharacterMasterIDs:     []int64{1, 2, 3, 4, 5},
		},
		Debug: DebugConfig{
			DumpAPIResponse: false,
			DumpDir:         "data/dumps",
		},
	}
	cfg.applyLogLevelDefault()
	return cfg
}

func mustParseDuration(val string) time.Duration {
	dur, err := time.ParseDuration(val)
	if err != nil {
		return 0
	}
	return dur
}

func (c *Config) applyLogLevelDefault() {
	if c == nil {
		return
	}
	if strings.TrimSpace(c.Env) == "" {
		c.Env = "development"
	}
	if strings.TrimSpace(c.LogLevel) == "" {
		c.LogLevel = "info"
	}
	if c.Cache.MaxEntries <= 0 {
		c.Cache.MaxEntries = 512
	}
	if c.Cache.TTL <= 0 {
		c.Cache.TTL = mustParseDuration("30s")
	}
	if strings.TrimSpace(c.Game.BaseURL) == "" {
		c.Game.BaseURL = "http://127.0.0.1:8080"
	}
	if strings.TrimSpace(c.Game.RealtimeURL) == "" {
		c.Game.RealtimeURL = "http://127.0.0.1:8081"
	}
	if strings.TrimSpace(c.Game.ApplicationVersion) == "" {
		c.Game.ApplicationVersion = "2.30.1"
	}
	if strings.TrimSpace(c.Game.AssetVersion) == "" {
		c.Game.AssetVersion = "1.91.0"
	}
	if strings.TrimSpace(c.Game.ApiAssetsVersion) == "" {
		c.Game.ApiAssetsVersion = "1.12.0"
	}
	if strings.TrimSpace(c.Game.AssetRoot) == "" {
		c.Game.AssetRoot = "data/assets/production"
	}
	if strings.TrimSpace(c.Game.MasterDataURI) == "" {
		c.Game.MasterDataURI = "2026-05-09/mastermemory_1778239615_1778239615.db"
	}
	if strings.TrimSpace(c.Game.MasterDataVersion) == "" {
		c.Game.MasterDataVersion = "1778239615_1778239615"
	}
	if c.Game.MasterDataPublishTimestamp <= 0 {
		c.Game.MasterDataPublishTimestamp = 1778239615
	}
	if strings.TrimSpace(c.Game.MasterDataRoot) == "" {
		c.Game.MasterDataRoot = "data/master-data/production"
	}
	if strings.TrimSpace(c.Game.MasterDecodedJSONRoot) == "" {
		c.Game.MasterDecodedJSONRoot = "../ymst_masterdata/data/master/decoded"
	}
	if strings.TrimSpace(c.Game.UserDataFixture) == "" {
		c.Game.UserDataFixture = "references/captured/2026-05-16T06-13-39-729Z_res_lb-api.wds-stellarium.com_api_data_user.bin"
	}
	if strings.TrimSpace(c.Game.UserDataSeedJSON) == "" {
		c.Game.UserDataSeedJSON = "references/captured/2026-05-16T06-13-39-729Z_res_lb-api.wds-stellarium.com_api_data_user.bin.api.json"
	}
	if strings.TrimSpace(c.Game.UserDataSeedMode) == "" {
		c.Game.UserDataSeedMode = "captured_json"
	}
	if strings.TrimSpace(c.Game.GeneratedUserProfile) == "" {
		c.Game.GeneratedUserProfile = "template_minimal"
	}
	if c.Game.GeneratedUser.PlayerRank <= 0 {
		c.Game.GeneratedUser.PlayerRank = 1
	}
	if c.Game.GeneratedUser.CurrentStamina <= 0 {
		c.Game.GeneratedUser.CurrentStamina = 100
	}
	if c.Game.GeneratedUser.Coin <= 0 {
		c.Game.GeneratedUser.Coin = 100000
	}
	if c.Game.GeneratedUser.FreeJewel <= 0 {
		c.Game.GeneratedUser.FreeJewel = 3000
	}
	if c.Game.GeneratedUser.TutorialStatus <= 0 {
		c.Game.GeneratedUser.TutorialStatus = 99
	}
	if len(c.Game.SeedCharacterBaseMasterIDs) == 0 {
		c.Game.SeedCharacterBaseMasterIDs = []int64{1, 2, 3, 4, 5}
	}
	if len(c.Game.SeedCharacterMasterIDs) == 0 {
		c.Game.SeedCharacterMasterIDs = []int64{1, 2, 3, 4, 5}
	}
	if strings.TrimSpace(c.Debug.DumpDir) == "" {
		c.Debug.DumpDir = "data/dumps"
	}
}

func (g GameConfig) ClientVersion() string {
	if strings.TrimSpace(g.ApplicationVersion) != "" {
		return g.ApplicationVersion
	}
	return "2.30.1"
}

func (g GameConfig) APIAssetsVersion() string {
	if strings.TrimSpace(g.ApiAssetsVersion) != "" {
		return g.ApiAssetsVersion
	}
	return "1.12.0"
}
