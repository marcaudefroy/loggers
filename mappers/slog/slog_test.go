package slog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"testing"
)

func TestSlogAdapter(t *testing.T) {
	// Préparer un buffer pour capturer la sortie du logger
	var buf bytes.Buffer

	// Créer un handler JSON avec le buffer comme sortie
	handler := slog.NewJSONHandler(&buf, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	// Créer un logger Contextual en utilisant l'adaptateur slog
	logger := NewLogger(slog.New(handler))

	// Ajouter des champs contextuels
	logger = logger.WithFields("component", "test", "module", "adapter")

	// Émettre un message de niveau Info
	logger.Info("Test message", "key1", "value1", "key2", 42)

	// Analyser la sortie JSON
	var logEntry map[string]interface{}
	decoder := json.NewDecoder(&buf)
	if err := decoder.Decode(&logEntry); err != nil {
		t.Fatalf("Échec du décodage de la sortie JSON : %v", err)
	}
	fmt.Println(logEntry)

	// Vérifier le niveau de log
	if level, ok := logEntry["level"]; !ok || level != "INFO" {
		t.Errorf("Niveau de log incorrect : attendu 'INFO', obtenu '%v'", level)
	}

	// Vérifier le message
	if msg, ok := logEntry["msg"]; !ok || msg != "Test message" {
		t.Errorf("Message incorrect : attendu 'Test message', obtenu '%v'", msg)
	}

	// Vérifier les champs contextuels
	expectedFields := map[string]interface{}{
		"component": "test",
		"module":    "adapter",
		"key1":      "value1",
		"key2":      float64(42), // JSON décode les nombres en float64
	}

	for key, expectedValue := range expectedFields {
		if value, ok := logEntry[key]; !ok || value != expectedValue {
			t.Errorf("Champ '%s' incorrect : attendu '%v', obtenu '%v'", key, expectedValue, value)
		}
	}
}
