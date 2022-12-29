package server

import (
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	a *fiber.App
	s ServerConfig
}

type ServerConfig struct {
	Host string	`json:"server.host"`
	Port int	`json:"server.port"`

	Timeout 	int		`json:"server.timeout"`
	logFormat	string	`json:"server.log_format"`
	timeFormat 	string	`json:"server.time_format"`
	timeZone 	string	`json:"server.time_zone"`

	API 			APIConfig		`json:"api"`
	TokensConfig 	TokensConfig	`json:"tokens"`
	TexturesConfig 	TexturesConfig	`json:"textures"`

	UserDB Database

	EnableCORS		bool	`json:"server.enable_cors"`
	EnableGzip		bool	`json:"server.enable_gzip"`
	EnablePrefork 	bool	`json:"server.enable_prefork"`
	ConcurrentLimit int		`json:"server.concurrent_limit"`
}

type Database struct {
	Host	 string
	Port	 int
	Username string
	Password string
	Database string
}

type APIConfig struct {
	BaseURL 				string	`json:"api.base_url"`
	ServerName				string	`json:"api.server_name"`
	ImplemantationName 		string	`json:"api.implementation_name"`
	ImplemantationVersion 	string	`json:"api.implementation_version"`
}

type TokensConfig struct {
	Expiring       bool		`json:"token.expiring"`
	ExpirationTime int		`json:"token.expiration_time"`
	TokenDB        Database	`json:"token.database"`
}

type TexturesConfig struct {
	Uploadable  []string	`json:"textures.uploadable"`
	TexturesDB  Database	`json:"textures.database"`
	ServePath   string		`json:"textures.serve_path"`
	StoragePath string		`json:"textures.storage_path"` 
	MaximalSize int			`json:"textures.maximal_size"`
}

type Extension struct {
	ExtensionName 		string 	`json:"extensions.name`
	ExtensionID 		string	`json:"extensions.id"`
	ExtensionVersion 	string	`json:"extensions.version"`
	ExtensionProvider	string	`json:"extensions.provider"`
}