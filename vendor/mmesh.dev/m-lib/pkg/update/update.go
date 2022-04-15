package update

import (
	"bufio"
	"bytes"
	"crypto"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"

	"github.com/inconshreveable/go-update"
	"github.com/spf13/viper"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/logging"
	"mmesh.dev/m-lib/pkg/utils"
	"mmesh.dev/m-lib/pkg/utils/colors"
	"mmesh.dev/m-lib/pkg/version"
	"mmesh.dev/m-lib/pkg/xlog"
)

// const binURL string = "https://storage.googleapis.com/mmesh-io/binaries"
const binURL string = "https://mmesh.s3.nl-ams.scw.cloud/binaries"
const binVersion string = "latest"
const branchStable string = "stable"

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAyl1lTZHz7sK/TnU2YFeZ
d6aP72Ey9XzLfPeVgdgL5/pWG4fH3rUDdxZO48J3YxJaPNkq92TJbfU3VNVZReXs
dRLXY1NoiqFB6fi5hYCLSxSJXP82KSCKIPjaBrcIYA1/HBnMeIWDYwBFoznnJr1d
NOBVFncVGixKFO2K7fTwt+Kca3vIrR4H8cKkkPtm5urYo5WZ3u2GeA6KhQSXinDj
eSFwleSFju7jxCEPw5/7ScDsnF4OztAMrbaVu4DewyIaC4gNFwlmEFfVT7fca8mK
bnY0jVYNx3YRJeWWGSskv9C/TORINTU2G5K1yZNLHYzzAbV3QcKEfsStZr3P9jO9
Wpta9h7iGFcWFpiUfuW0G0+S97C77QjUiMl8CVRxQu8Rnz0Ski99yTvSXhQpNmdX
jO2q4olhZSW9zOpxXKl53w3ZGinqTw0USOsM5kaq4qNR6qkUL+f0KccivH6wilzb
Fw/GpL6Q64+iQ/T87xMIcPQ7ckYC5Ls0vyrojyCPOx4Bep3T4PUhwyL3fYKt2zaV
3GOplCDunWXZqxbN4P6LU4lMZoTaNjrzGFkjFM+S1qEJLbtixDfyPq7i6DolQfHu
yeOUZE0aR11fyjX1rzpjy9occeDdPLFcmlz09kvmrtuDh9+wHwC2QpLVRO4w+xKp
DwIt2q7+1TeUXK8X5zrCstUCAwEAAQ==
-----END PUBLIC KEY-----
`)

func IsBinaryOutdated(app string) (bool, error) {
	exe, err := os.Executable()
	if err != nil {
		return false, errors.Wrapf(err, "[%v] function os.Executable()", errors.Trace())
	}

	checksum, err := getChecksum(app)
	if err != nil {
		return false, errors.Wrapf(err, "[%v] function getChecksum()", errors.Trace())
	}

	return binaryIsOutdated(exe, checksum), nil
}

func Update(app string) error {
	exe, err := os.Executable()
	if err != nil {
		return errors.Wrapf(err, "[%v] function os.Executable()", errors.Trace())
	}

	checksum, err := getChecksum(app)
	if err != nil {
		return errors.Wrapf(err, "[%v] function getChecksum()", errors.Trace())
	}

	if !binaryIsOutdated(exe, checksum) {
		return nil
	}

	if app == version.CLI_NAME {
		fmt.Println("New version available, updating...")
	} else {
		xlog.Info("New version available, updating...")
	}

	signature, err := getSignature(app)
	if err != nil {
		return errors.Wrapf(err, "[%v] function getSignature()", errors.Trace())
	}

	opts := update.Options{
		Checksum:  checksum,
		Signature: signature,
		Hash:      crypto.SHA256,
		Verifier:  update.NewRSAVerifier(),
	}

	if err := opts.SetPublicKeyPEM(publicKey); err != nil {
		return errors.Wrapf(err, "[%v] function opts.SetPublicKeyPEM()", errors.Trace())
	}

	if err := opts.CheckPermissions(); err != nil {
		if app == version.CLI_NAME {
			fmt.Printf("Unable to update binary: %s\n", colors.DarkRed("permission denied on filesystem"))
		} else {
			xlog.Warnf("Unable to update binary: %v", err)
		}
		return errors.Cause(err)
	}

	// request the new file
	url := getURL(getBinaryName(app))
	resp, err := http.Get(url)
	if err != nil {
		return errors.Wrapf(err, "[%v] function http.Get()", errors.Trace())
	}
	defer resp.Body.Close()

	if err := update.Apply(resp.Body, opts); err != nil {
		if app != version.CLI_NAME {
			xlog.Errorf("Unable to apply update: %v", err)
		}
		if rerr := update.RollbackError(err); rerr != nil {
			if app != version.CLI_NAME {
				xlog.Errorf("Unable to rollback from bad update: %v", rerr)
			}
			return errors.Wrapf(rerr, "[%v] function update.RollbackError()", errors.Trace())
		}
		return errors.Wrapf(err, "[%v] function update.Apply()", errors.Trace())
	}

	restartProcess(app, exe)

	return nil
}

func getURL(file string) string {
	goos := runtime.GOOS
	goarch := runtime.GOARCH

	if goarch == "arm" {
		goarch = "arm_6"
	}

	binBranch := viper.GetString("version.branch")
	if len(binBranch) == 0 {
		binBranch = branchStable
	}

	return fmt.Sprintf("%s/%s/%s/%s/%s/%s", binURL, binBranch, binVersion, goos, goarch, file)
}

func binaryIsOutdated(exe string, checksum []byte) bool {
	if len(checksum) == 0 {
		return false
	}

	currentChecksum, err := utils.ChecksumSHA256(exe)
	if err != nil {
		logging.Errorf("Unable to get current binary's checksum: %v", err)
		return false
	}

	if len(currentChecksum) == 0 {
		return false
	}

	// msg.Infof("Current binary checksum:\n%x", string(currentChecksum))
	// msg.Infof("Latest binary checksum:\n%x\n", string(checksum))

	if !bytes.Equal(checksum, currentChecksum) {
		return true
	}

	return false
}

func getChecksum(app string) ([]byte, error) {
	filename := getBinaryName(app) + "_checksum.sha256"

	if err := downloadFile(filename); err != nil {
		return nil, errors.Wrapf(err, "[%v] function downloadFile()", errors.Trace())
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function os.Open()", errors.Trace())
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	words := make([]string, 0)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := f.Close(); err != nil {
		return nil, errors.Wrapf(err, "[%v] function f.Close()", errors.Trace())
	}

	if err := os.RemoveAll(filename); err != nil {
		return nil, errors.Wrapf(err, "[%v] function os.RemoveAll()", errors.Trace())
	}

	checksum, err := hex.DecodeString(words[0])
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function hex.DecodeString()", errors.Trace())
	}

	return checksum, nil
}

func getSignature(app string) ([]byte, error) {
	filename := getBinaryName(app) + "_signature.sha256"

	if err := downloadFile(filename); err != nil {
		return nil, errors.Wrapf(err, "[%v] function downloadFile()", errors.Trace())
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function ioutil.ReadFile()", errors.Trace())
	}

	if err := os.RemoveAll(filename); err != nil {
		return nil, errors.Wrapf(err, "[%v] function os.RemoveAll()", errors.Trace())
	}

	return data, nil
}

func getBinaryName(app string) string {
	if runtime.GOOS == "windows" {
		return fmt.Sprintf("%s.exe", app)
	}

	return app
}
