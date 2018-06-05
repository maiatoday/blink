// +build example
//
// Do not build by default.

/*
 How to run
 Pass the Bluetooth address or name as the first param:

	go run examples/sprkplus.go SK-1234

 NOTE: sudo is required to use BLE in Linux
*/

package main
import (
	//"os"
	"time"
  "fmt"
//  "math"
  //"github.com/aws/aws-sdk-go/service/iot"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/sphero/sprkplus"
)

// ThingConfig contains credentials for a thing.
// type ThingConfig struct {
// 	// The ARN of the certificate.
// 	CertificateArn string
//
// 	// The ID of the certificate. AWS IoT issues a default subject name for the
// 	// certificate (e.g., AWS IoT Certificate).
// 	CertificateID string
//
// 	// The certificate data, in PEM format.
// 	CertificatePem string
//
// 	// The generated key pair.
// 	KeyPair *iot.KeyPair
// }
//
// // NewThingConfig create a new thing configuration using the response
// func NewThingConfig(resp *iot.CreateKeysAndCertificateOutput) *ThingConfig {
// 	return &ThingConfig{
// 		CertificateArn: *resp.CertificateArn,
// 		CertificateID:  *resp.CertificateId,
// 		CertificatePem: *resp.CertificatePem,
// 		KeyPair:        resp.KeyPair,
// 	}
// }

func main() {
  //var username = "maiatoday"

	//bleAdaptor := ble.NewClientAdaptor(os.Args[1])
  bleAdaptor := ble.NewClientAdaptor("SK-C345")
	sprk := sprkplus.NewDriver(bleAdaptor)

	work := func() {
		gobot.Every(1*time.Second, func() {
      temperature := gobot.Rand(50)
      fmt.Printf("temp %d \n", temperature)
      r,g,b := colorFromTemp(temperature)
      fmt.Printf("r %d g %d b %d \n", r, g, b)
			// r := uint8(gobot.Rand(255))
			// g := uint8(gobot.Rand(255))
			// b := uint8(gobot.Rand(255))

			sprk.SetRGB(uint8(r), uint8(g), uint8(b))
		})
	}

	robot := gobot.NewRobot("sprkBot",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{sprk},
		work,
	)

	robot.Start()
}
const MAX_TEMP = 40
const MIN_TEMP = 0
func colorFromTemp(temperature int) (r, g, b int) {
  var tt float32
  if (temperature < MIN_TEMP) {
    tt = float32(MIN_TEMP)
  } else if (temperature > MAX_TEMP) {
    tt = float32(MAX_TEMP)
  } else {
    tt = float32(temperature)
  }

  r = int(255 * (tt/MAX_TEMP))
  b = int(255 * (MAX_TEMP-tt)/MAX_TEMP)
  g = int(125)
  return
}
