package license

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEncrypt(t *testing.T) {
	Convey("test use encrypt", t, func() {
		license := NewLicenseManger("123456789")
		key, _ := license.CreateLicense([]byte("jinjin"))
		t.Logf("%x", key)
		So(len(fmt.Sprintf("%x", key)), ShouldBeGreaterThan, 30)
	})
}

func TestParseInputFile(t *testing.T) {
	Convey("test parse file", t, func() {
		inputPath, _ := filepath.Abs("./user.txt")
		mapValue, _ := parseInputFileJSON(inputPath)
		uid := (*mapValue)["UID"]
		So(uid, ShouldEqual, "123")
	})
}

func TestOutputLicenseFile(t *testing.T) {
	Convey("test use encrypt", t, func() {
		inputPath, _ := filepath.Abs("./user.txt")
		mapValue, _ := parseInputFileJSON(inputPath)

		license := NewLicenseManger("123456789")
		outputPath, _ := filepath.Abs("./license.txt")
		err := license.OutputLicenseFile(mapValue, outputPath)
		So(err, ShouldBeNil)
	})
}

func TestCompare(t *testing.T) {
	Convey("test use Compare", t, func() {
		licenseFile, _ := filepath.Abs("./license.txt")
		licenseMap, _ := parseInputFileJSON(licenseFile)
		key, err := hex.DecodeString((*licenseMap)[LicKey])
		So(err, ShouldBeNil)

		inputPath, _ := filepath.Abs("./user.txt")
		fileMap, _ := parseInputFileJSON(inputPath)
		generateByte := parseMapSort(fileMap)

		license := NewLicenseManger("123456789")
		generateKey, _ := license.CreateLicense(generateByte)
		So(bytes.Equal(key, generateKey), ShouldBeTrue)
	})
}

func TestLoadInfoFromLicenseFile(t *testing.T) {
	Convey("test use LoadInfoFromLicenseFile", t, func() {
		license := NewLicenseManger("123456789")
		licenseFile, _ := filepath.Abs("./license.txt")
		fileMap, err := license.LoadInfoFromLicense(licenseFile)
		t.Logf("%v", fileMap)
		// log.Fatalln(err)
		So(err, ShouldBeNil)
	})
}
