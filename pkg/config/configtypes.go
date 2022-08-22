package config

type Config struct {
	Server   ConfigServer   `json:"server"`
	Database ConfigDatabase `json:"database"`
	Discord  ConfigDiscord  `json:"discord"`
	Email    ConfigEmail    `json:"email"`
	RabbitMQ ConfigRabbitMQ `json:"rabbitmq"`
	Redis    ConfigRedis    `json:"redis"`
	Session  ConfigSession  `json:"session"`
	OAuth    ConfigOAuth    `json:"oauth"`
	VATUSA   ConfigVATUSA   `json:"vatusa"`
	Storage  ConfigStorage  `json:"storage"`
}

type ConfigServer struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type ConfigDiscord struct {
	Webhooks map[string]string `json:"webhooks"`
}

type ConfigDatabase struct {
	Host        string `json:"host"`
	Port        string `json:"port"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Database    string `json:"database"`
	Automigrate bool   `json:"automigrate"`
}

type ConfigEmail struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	From     string `json:"from"`
}

type ConfigRabbitMQ struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type ConfigRedis struct {
	Password      string   `json:"password"`
	Database      int      `json:"database"`
	Address       string   `json:"address"`
	Sentinel      bool     `json:"sentinel"`
	MasterName    string   `json:"master_name"`
	SentinelAddrs []string `json:"sentinel_addrs"`
}

type ConfigSession struct {
	Cookie ConfigSessionCookie `json:"cookie"`
}

type ConfigSessionCookie struct {
	Name   string `json:"name"`
	Secret string `json:"secret"`
	Domain string `json:"domain"`
	Path   string `json:"path"`
	MaxAge int    `json:"max_age"`
}

type ConfigOAuth struct {
	BaseURL      string `json:"base_URL"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	MyBaseURL    string `json:"my_base_URL"`

	Endpoints ConfigOAuthEndpoints `json:"endpoints"`
}

type ConfigOAuthEndpoints struct {
	Authorize string `json:"authorize"`
	Token     string `json:"token"`
	UserInfo  string `json:"user"`
}

type ConfigVATUSA struct {
	APIKey string `json:"api_key"`
}

type ConfigStorage struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Bucket    string `json:"bucket"`
	Region    string `json:"region"`
	Endpoint  string `json:"endpoint"`
}
