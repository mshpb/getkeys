package getkeys

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

// load environment variables file that provides.
func LoadEnv(fileName string) error {
	if fileName == "" {
		return errors.New("Need at least one environment variable filename")
	}

	if err := readFile(fileName); err != nil {
		return err
	}

	return nil
}

// read environment file from same directory.
func readFile(name string) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()

	return parseFile(file)
}

// parse environment variable file data and make map from that and set key value pairs.
func parseFile(file io.Reader) error {
	// Create a map to store key-value pairs
	envMap := make(map[string]string)

	// Read and parse the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			envMap[key] = value
		}
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		return err
	}

	if err := setEnvKeys(envMap); err != nil {
		return err
	}

	return nil
}

// set key and value to the environment
func setEnvKeys(envMap map[string]string) error {
	for key, value := range envMap {
		if err := os.Setenv(key, value); err != nil {
			return err
		}
	}

	return nil
}
