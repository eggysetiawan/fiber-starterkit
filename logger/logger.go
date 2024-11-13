package logger

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/eggysetiawan/fiber-starterkit/config"
	"github.com/elastic/go-elasticsearch"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func NewLogger() {

	// Initialize Elasticsearch client
	var err error

	// d := time.Now().Format(time.DateOnly)

	cfg := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "@timestamps"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	cfg.EncoderConfig = encoderConfig
	// config.OutputPaths = []string{"stdout", fmt.Sprintf("./storage/logs/%s-%s.log", "london", d)}
	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stderr"}

	log, err := cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	if !config.AppConfig.GetBool("LOG_TO_ELASTIC") {
		logger = log
		return
	}

	esCfg := elasticsearch.Config{
		Addresses: strings.Split(config.AppConfig.GetString("ELASTIC_HOST"), ","),
	}

	esClient, err := elasticsearch.NewClient(esCfg)
	if err != nil {
		panic("failed to init elastic client: " + err.Error())
	}

	// Create Elasticsearch logger core
	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)
	esWriter := &ElasticsearchWriter{client: esClient}
	esCore := zapcore.NewCore(jsonEncoder, zapcore.AddSync(esWriter), zapcore.InfoLevel)

	// Combine standard logger core and Elasticsearch logger core
	core := zapcore.NewTee(
		log.Core(),
		esCore,
	)

	// Create combined logger
	logger = zap.New(core)

	// Use the logger
	zap.ReplaceGlobals(logger)
}

type ElasticsearchWriter struct {
	client *elasticsearch.Client
}

func (w *ElasticsearchWriter) Write(p []byte) (n int, err error) {
	idx := fmt.Sprintf("%s-%s", config.AppConfig.GetString("ELASTIC_INDEX"), time.Now().Format(time.DateOnly))
	// Assuming logs are formatted as JSON, so no need to encode here
	// Send log entry to Elasticsearch
	_, err = w.client.Index(
		idx,
		bytes.NewReader(p),
	)
	if err != nil {
		fmt.Println(fmt.Errorf("error indexing document %s: %v", idx, err))
		return 0, err
	}
	return len(p), nil
}

func Info(message string, fields ...zap.Field) {
	logger.Info(message, fields...)
}

func Infof(template string, fields ...interface{}) {
	logger.Sugar().Infof(template, fields...)
}

func Debug(message string, fields ...zap.Field) {
	logger.Debug(message, fields...)
}
func Error(message string, fields ...zap.Field) {
	logger.Error(message, fields...)
}

func Errorf(template string, fields ...interface{}) {
	logger.Sugar().Errorf(template, fields...)
}
