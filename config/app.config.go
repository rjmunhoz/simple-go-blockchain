package config

type AppConfig struct {
  Port string
  Message string
}


func NewAppConfig() AppConfig {
  return AppConfig{
    Port: GetEnv("SERVER_BINDING_PORT", ":8080"),
    Message: GetEnv("HELLO_MESSAGE", "Hello, world!"),
  }
}
