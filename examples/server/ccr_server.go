// Copyright 2013-2015 go-diameter authors.  All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Diameter server example. This is by no means a complete server.
//
// If you'd like to test diameter over SSL, generate SSL certificates:
//   go run $GOROOT/src/crypto/tls/generate_cert.go --host localhost
//
// And start the server with `-cert_file cert.pem -key_file key.pem`.
//
// By default this server runs in a single OS thread. If you want to
// make it run on more, set the GOMAXPROCS=n environment variable.
// See Go's FAQ for details: http://golang.org/doc/faq#Why_no_multi_CPU

package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	_ "net/http/pprof"

	"github.com/fiorix/go-diameter/v4/diam"
	"github.com/fiorix/go-diameter/v4/diam/avp"
	"github.com/fiorix/go-diameter/v4/diam/datatype"
	"github.com/fiorix/go-diameter/v4/diam/dict"
	"github.com/fiorix/go-diameter/v4/diam/sm"
)

func main() {
	addr := flag.String("addr", ":3868", "address in the form of ip:port to listen on")
	ppaddr := flag.String("pprof_addr", ":9000", "address in form of ip:port for the pprof server")
	host := flag.String("diam_host", "server", "diameter identity host")
	realm := flag.String("diam_realm", "go-diameter", "diameter identity realm")
	certFile := flag.String("cert_file", "", "tls certificate file (optional)")
	keyFile := flag.String("key_file", "", "tls key file (optional)")
	silent := flag.Bool("s", false, "silent mode, useful for benchmarks")
	flag.Parse()

	// Load our custom dictionary on top of the default one, which
	// always have the Base Protocol (RFC6733) and Credit Control
	// Application (RFC4006).
	err := dict.Default.Load(bytes.NewReader([]byte(helloDictionary)))
	if err != nil {
		log.Fatal(err)
	}

	settings := &sm.Settings{
		OriginHost:       datatype.DiameterIdentity(*host),
		OriginRealm:      datatype.DiameterIdentity(*realm),
		VendorID:         13,
		ProductName:      "go-diameter",
		FirmwareRevision: 1,
	}

	// Create the state machine (mux) and set its message handlers.
	mux := sm.New(settings)
	mux.Handle("HMR", handleHMR(*silent))
	mux.Handle("CCR", handleCCR(*silent))
	mux.Handle("ACR", handleACR(*silent))
	mux.HandleFunc("ALL", handleALL) // Catch all.

	// Print error reports.
	go printErrors(mux.ErrorReports())

	if len(*ppaddr) > 0 {
		go func() { log.Fatal(http.ListenAndServe(*ppaddr, nil)) }()
	}

	err = listen(*addr, *certFile, *keyFile, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func printErrors(ec <-chan *diam.ErrorReport) {
	for err := range ec {
		log.Println(err)
	}
}

func listen(addr, cert, key string, handler diam.Handler) error {
	// Start listening for connections.
	if len(cert) > 0 && len(key) > 0 {
		log.Println("Starting secure diameter server on", addr)
		return diam.ListenAndServeTLS(addr, cert, key, handler, nil)
	}
	log.Println("Starting diameter server on", addr)
	return diam.ListenAndServe(addr, handler, nil)
}

func handleHMR(silent bool) diam.HandlerFunc {
	type HelloRequest struct {
		SessionID        datatype.UTF8String       `avp:"Session-Id"`
		OriginHost       datatype.DiameterIdentity `avp:"Origin-Host"`
		OriginRealm      datatype.DiameterIdentity `avp:"Origin-Realm"`
		DestinationRealm datatype.DiameterIdentity `avp:"Destination-Realm"`
		DestinationHost  datatype.DiameterIdentity `avp:"Destination-Host"`
		UserName         string                    `avp:"User-Name"`
	}
	return func(c diam.Conn, m *diam.Message) {
		if !silent {
			log.Printf("Received HMR from %s:\n%s", c.RemoteAddr(), m)
		}
		var hmr HelloRequest
		if err := m.Unmarshal(&hmr); err != nil {
			log.Printf("Failed to parse message from %s: %s\n%s",
				c.RemoteAddr(), err, m)
			return
		}
		a := m.Answer(diam.Success)
		a.NewAVP(avp.SessionID, avp.Mbit, 0, hmr.SessionID)
		a.NewAVP(avp.OriginHost, avp.Mbit, 0, hmr.DestinationHost)
		a.NewAVP(avp.OriginRealm, avp.Mbit, 0, hmr.DestinationRealm)
		a.NewAVP(avp.DestinationRealm, avp.Mbit, 0, hmr.OriginRealm)
		a.NewAVP(avp.DestinationHost, avp.Mbit, 0, hmr.OriginHost)
		_, err := a.WriteTo(c)
		if err != nil {
			log.Printf("Failed to write message to %s: %s\n%s\n",
				c.RemoteAddr(), err, a)
			return
		}
		if !silent {
			log.Printf("Sent HMA to %s:\n%s", c.RemoteAddr(), a)
		}
	}
}

func handleCCR(silent bool) diam.HandlerFunc {

	type SubscriptionIDSTRUCT struct {
		SubscriptionIDType datatype.Enumerated `avp:"Subscription-Id-Type"`
		SubscriptionIDData datatype.UTF8String `avp:"Subscription-Id-Data"`
	}

	type RSUnitSTRUCT struct {
		CCTotalOctets datatype.Unsigned64 `avp:"CC-Total-Octets"`
	}

	type MSCControlSTRUCT struct {
		RequestedServiceUnit RSUnitSTRUCT        `avp:"Requested-Service-Unit"`
		ServiceIdentifier    datatype.Unsigned32 `avp:"Service-Identifier"`
		RatingGroup          datatype.Unsigned32 `avp:"Rating-Group"`
	}

	type UEInfoSTRUCT struct {
		UserEquipmentInfoType  datatype.Enumerated  `avp:"User-Equipment-Info-Type"`
		UserEquipmentInfoValue datatype.OctetString `avp:"User-Equipment-Info-Value"`
	}

	type PSInformationSTRUCT struct {
		PDPAddress datatype.Address `avp:"PDP-Address"`
	}

	type SInformationSTRUCT struct {
		PSInformation PSInformationSTRUCT `avp:"PS-Information"`
	}

	type CCRPacket struct {
		SessionID                     datatype.UTF8String       `avp:"Session-Id"`
		OriginHost                    datatype.DiameterIdentity `avp:"Origin-Host"`
		OriginRealm                   datatype.DiameterIdentity `avp:"Origin-Realm"`
		DestinationRealm              datatype.DiameterIdentity `avp:"Destination-Realm"`
		AuthApplicationID             datatype.Unsigned32       `avp:"Auth-Application-Id"`
		ServiceContextID              datatype.UTF8String       `avp:"Service-Context-Id"`
		CCRequestType                 datatype.Enumerated       `avp:"CC-Request-Type"`
		CCRequestNumber               datatype.Unsigned32       `avp:"CC-Request-Number"`
		DestinationHost               datatype.DiameterIdentity `avp:"Destination-Host"`
		UserName                      string                    `avp:"User-Name"`
		OriginStateID                 datatype.Unsigned32       `avp:"Origin-State-Id"`
		EventTimestamp                datatype.Time             `avp:"Event-Timestamp"`
		SubscriptionID1               SubscriptionIDSTRUCT      `avp:"Subscription-Id"`
		SubscriptionID2               SubscriptionIDSTRUCT      `avp:"Subscription-Id"`
		MultipleServicesIndicator     datatype.Enumerated       `avp:"Multiple-Services-Indicator"`
		MultipleServicesCreditControl MSCControlSTRUCT          `avp:"Multiple-Services-Credit-Control"`
		UserEquipmentInfo             UEInfoSTRUCT              `avp:"User-Equipment-Info"`
		ServiceInformation            SInformationSTRUCT        `avp:"Service-Information"`
	}
	return func(c diam.Conn, m *diam.Message) {
		fmt.Println("\n>\n>\n>\n>\n>\n>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>------------")
		mStr := fmt.Sprint(m)
		re := regexp.MustCompile(`4175000\d*`)
		imsi := re.FindString(mStr)
		if !silent {
			log.Printf("Received CCR from %s:\n%s", c.RemoteAddr(), mStr)
		}

		f, errFile := os.OpenFile(fmt.Sprintf("cdrs/%v.log", imsi),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if errFile != nil {
			log.Println(errFile)
		}
		defer f.Close()
		nowTime := time.Now()
		if _, errFile := f.WriteString(fmt.Sprintf("\n%v\nREQUEST:\n%v\n", nowTime, mStr)); errFile != nil {
			log.Println(errFile)
		}

		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>------------2")
		var ccr CCRPacket
		if err := m.Unmarshal(&ccr); err != nil {
			log.Printf("Failed to parse message from %s: %s\n%s",
				c.RemoteAddr(), err, m)
			return
		}
		ccr.SubscriptionID2.SubscriptionIDType = datatype.Enumerated(1)
		ccr.SubscriptionID2.SubscriptionIDData = datatype.UTF8String(imsi)

		fmt.Printf("\n+++++++++++++++++++++++++++++++++++ %+v\n", ccr)
		fmt.Println("********************", string(ccr.ServiceInformation.PSInformation.PDPAddress.Serialize()))
		a := m.Answer(diam.Success)
		// a.NewAVP(avp.ResultCode, avp.Mbit, 0, datatype.Unsigned32(2001))
		a.NewAVP(avp.SessionID, avp.Mbit, 0, ccr.SessionID)
		a.NewAVP(avp.OriginHost, avp.Mbit, 0, ccr.DestinationHost)
		a.NewAVP(avp.OriginRealm, avp.Mbit, 0, ccr.DestinationRealm)
		a.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, ccr.AuthApplicationID)
		a.NewAVP(avp.CCRequestType, avp.Mbit, 0, ccr.CCRequestType)
		a.NewAVP(avp.CCRequestNumber, avp.Mbit, 0, ccr.CCRequestNumber)
		a.NewAVP(avp.MultipleServicesCreditControl, avp.Mbit, 0, &diam.GroupedAVP{
			AVP: []*diam.AVP{
				diam.NewAVP(avp.GrantedServiceUnit, avp.Mbit, 0, &diam.GroupedAVP{
					AVP: []*diam.AVP{
						diam.NewAVP(avp.CCTotalOctets, avp.Mbit, 0, datatype.Unsigned64(5242880)),
					},
				}),
				diam.NewAVP(avp.RatingGroup, avp.Mbit, 0, datatype.Unsigned32(999)),
				diam.NewAVP(avp.ServiceIdentifier, avp.Mbit, 0, datatype.Unsigned32(999)),
				diam.NewAVP(avp.ResultCode, avp.Mbit, 0, datatype.Unsigned32(2001)),
			},
		})

		if _, errFile := f.WriteString(fmt.Sprintf("\n%v --- Duration: %v\nANSWER:\n%v\n", time.Now(), time.Since(nowTime), a)); errFile != nil {
			log.Println(errFile)
		}

		fmt.Printf("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$\n$$$$$$$$$$$$\n %+v\n", a)
		_, err := a.WriteTo(c)
		if err != nil {
			log.Printf("Failed to write message to %s: %s\n%s\n",
				c.RemoteAddr(), err, a)
			return
		}
		if !silent {
			log.Printf("Sent HMA to %s:\n%s", c.RemoteAddr(), a)
		}
	}
}

func handleACR(silent bool) diam.HandlerFunc {
	type AccountingRequest struct {
		SessionID              *diam.AVP                 `avp:"Session-Id"`
		OriginHost             *diam.AVP                 `avp:"Origin-Host"`
		OriginRealm            *diam.AVP                 `avp:"Origin-Realm"`
		DestinationRealm       datatype.DiameterIdentity `avp:"Destination-Realm"`
		AccountingRecordType   *diam.AVP                 `avp:"Accounting-Record-Type"`
		AccountingRecordNumber *diam.AVP                 `avp:"Accounting-Record-Number"`
		DestinationHost        datatype.DiameterIdentity `avp:"Destination-Host"`
	}
	return func(c diam.Conn, m *diam.Message) {
		if !silent {
			log.Printf("Received ACR from %s\n%s", c.RemoteAddr(), m)
		}
		var acr AccountingRequest
		if err := m.Unmarshal(&acr); err != nil {
			log.Printf("Failed to parse message from %s: %s\n%s",
				c.RemoteAddr(), err, m)
			return
		}
		a := m.Answer(diam.Success)
		a.InsertAVP(acr.SessionID)
		a.NewAVP(avp.OriginHost, avp.Mbit, 0, acr.DestinationHost)
		a.NewAVP(avp.OriginRealm, avp.Mbit, 0, acr.DestinationRealm)
		a.AddAVP(acr.AccountingRecordType)
		a.AddAVP(acr.AccountingRecordNumber)
		_, err := a.WriteTo(c)
		if err != nil {
			log.Printf("Failed to write message to %s: %s\n%s\n",
				c.RemoteAddr(), err, a)
			return
		}
		if !silent {
			log.Printf("Sent ACA to %s:\n%s", c.RemoteAddr(), a)
		}
	}
}

func handleALL(c diam.Conn, m *diam.Message) {
	log.Printf("Received unexpected message from %s:\n%s", c.RemoteAddr(), m)
}

// helloDictionary is our custom, example dictionary.
var helloDictionary = xml.Header + `
<diameter>
	<application id="999" type="acct">
		<command code="111" short="HM" name="Hello-Message">
			<request>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="User-Name" required="false" max="1"/>
			</request>
			<answer>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Result-Code" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Error-Message" required="false" max="1"/>
			</answer>
		</command>
	</application>
</diameter>
`
