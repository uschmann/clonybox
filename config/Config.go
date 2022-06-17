package config

import "os"

type Config struct {
	DatabaseFile string

	RfidType string

	RfidEvdevFile   string
	Mfrc522ResetPin string
	Mfrc522IrqPin   string
}

func FromEnv() *Config {
	return &Config{
		DatabaseFile: getEnv("DB_FILE", "database.db"),

		RfidType: getEnv("RFID_TYPE", "stdio"), // stdio, evdev, mfrc522

		RfidEvdevFile: getEnv("RFID_EVDEV_FILE", "/dev/input/event0"),

		Mfrc522ResetPin: getEnv("MFRC522_RESET_PIN", "P1_22"),
		Mfrc522IrqPin:   getEnv("MFRC522_IRQ_PIN", "P1_18"),
	}
}

func getEnv(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)

	if ok {
		return value
	}

	return defaultValue
}
