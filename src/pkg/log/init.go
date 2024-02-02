package log

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	cfg, err := getZapConfig()
	if err != nil {
		return nil, err
	}

	logger, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build logger's config: %s", err)
	}

	return logger, nil
}

func getZapConfig() (*zap.Config, error) {
	v := viper.New()
	v.SetConfigFile("./configs/zap_config.yml")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &zap.Config{}
	logLevel, err := zap.ParseAtomicLevel(v.GetString("logger.level"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse log level: %s", err)
	}

	cfg.Level = logLevel
	cfg.Encoding = v.GetString("logger.encoding")
	cfg.OutputPaths = []string{v.GetString("logger.output_paths")}
	cfg.ErrorOutputPaths = []string{v.GetString("logger.error_output_paths")}

	cfg.EncoderConfig.MessageKey = v.GetString("encoder.message_key")
	cfg.EncoderConfig.LevelKey = v.GetString("encoder.level_key")
	cfg.EncoderConfig.TimeKey = v.GetString("encoder.time_key")
	cfg.EncoderConfig.CallerKey = v.GetString("encoder.caller_key")

	_ = cfg.EncoderConfig.EncodeLevel.UnmarshalText([]byte(v.GetString("encoder.level_encoder")))
	_ = cfg.EncoderConfig.EncodeTime.UnmarshalText([]byte(v.GetString("encoder.time_encoder")))
	_ = cfg.EncoderConfig.EncodeCaller.UnmarshalText([]byte(v.GetString("encoder.caller_encoder")))

	return cfg, nil
}
