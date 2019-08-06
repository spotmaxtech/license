package license

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"io"
	"os"
	"sort"
	"strings"
)

const LicKey = "License"

type Manager struct {
	magic string
}

func NewLicenseManger(magic string) *Manager {
	license := Manager{
		magic: magic,
	}

	return &license
}

func (mgr *Manager) CreateLicense(data []byte) ([]byte, error) {
	var buffer bytes.Buffer
	buffer.Write(data)
	buffer.WriteString(mgr.magic)

	shaVal := sha512.Sum512(buffer.Bytes())
	md5Val := md5.Sum(shaVal[:])
	return md5Val[:], nil
}

func (mgr *Manager) LoadInfoFromLicense(licensePath string) (*map[string]string, error) {
	licenseMap, err := parseInputFileJSON(licensePath)
	if err != nil {
		return nil, err
	}

	generateByte := parseMapSort(licenseMap)
	verifyData, _ := mgr.CreateLicense(generateByte)

	licenseValue := (*licenseMap)[LicKey]
	license, err := hex.DecodeString(licenseValue)

	if err != nil {
		return nil, err
	}

	if ok := bytes.Equal(verifyData, license); ok == true {
		return licenseMap, nil
	}

	return nil, errors.New("license validation failed")
}

func (mgr *Manager) OutputLicenseFile(data *map[string]string, outputPath string) error {

	outputFile, err := fileIsExist(outputPath)
	if err != nil {
		panic("failed to Open file" + err.Error())
	}

	defer outputFile.Close()

	var buffer bytes.Buffer
	for k, v := range *data {
		buffer.WriteString(k + ":" + v + "\r\n")
	}

	key := parseMapSort(data)

	license, _ := mgr.CreateLicense(key)
	buffer.WriteString(LicKey + ":" + hex.EncodeToString(license) + "\r\n")

	_, err = outputFile.Write(buffer.Bytes())

	return err
}

func fileIsExist(outputPath string) (*os.File, error) {
	if _, err := os.Stat(outputPath); !os.IsNotExist(err) {
		if err := os.Remove(outputPath); err != nil {
			return nil, err
		}

	}

	return os.Create(outputPath)
}

func parseInputFileJSON(inputPath string) (*map[string]string, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		// panic("failed to load file" + err.Error())
		return nil, err
	}
	defer file.Close()

	br := bufio.NewReader(file)
	result := make(map[string]string)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF || len(a) <= 0 {
			break
		}
		str := string(a)
		semiIndex := strings.Index(str, ":")
		if semiIndex > -1 {
			key := strings.TrimSpace(string(str[:semiIndex]))
			value := strings.TrimSpace(string(str[(semiIndex + 1):]))
			result[key] = value
		}
	}

	return &result, nil
}

func parseMapSort(data *map[string]string) []byte {
	var result []string
	for k, v := range *data {
		if k != LicKey {
			result = append(result, v)
		}
	}
	sort.Strings(result)
	str := strings.Join(result, "")
	return []byte(str)
}
