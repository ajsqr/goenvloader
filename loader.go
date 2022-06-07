package goenvloader

import "os"

type loader struct {
	file string
}

func NewLoader(file string) *loader {
	return &loader{
		file: file,
	}
}

func (l *loader) Load() error {
	scanner, err := loadEnvFile(l.file)
	if err != nil {
		return err
	}

	tokens, err := generateTokens(scanner)
	if err != nil {
		return err
	}

	for key, value := range tokens {
		os.Setenv(key, value)
	}
	return nil
}
