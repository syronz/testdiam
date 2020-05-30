// Copyright 2013-2015 go-diameter authors.  All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Diameter client example. This is by no means a complete client.
//
// If you'd like to test diameter over SSL, make sure the server supports
// it and add -ssl to the command line. To use client certificates,
// run the client with -ssl -cert_file cert.pem -key_file key.pem.
//
// When the client connects, the underlying state machine (diam/sm package)
// performs the handshake (CER/CEA) and returns a connection. If the
// client is configured with watchdog, it automatically sends DWR and
// handles DWA in background.
//
// The -hello command line flag makes the client connect, handshake,
// send a hello message, and disconnect. This is to demonstrate how to
// use custom dictionaries.
//
// The -bench option turns the client into a benchmark tool to test
// the server. It uses ACR/ACA messages for this.
//
// By default this client runs in a single OS thread. If you want to
// make it run on more, set the GOMAXPROCS=n environment variable.
// See Go's FAQ for details: http://golang.org/doc/faq#Why_no_multi_CPU

package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/fiorix/go-diameter/v4/diam"
	"github.com/fiorix/go-diameter/v4/diam/avp"
	"github.com/fiorix/go-diameter/v4/diam/datatype"
	"github.com/fiorix/go-diameter/v4/diam/dict"
	"github.com/fiorix/go-diameter/v4/diam/sm"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	addr := flag.String("addr", "localhost:3868", "address in form of ip:port to connect to")
	ssl := flag.Bool("ssl", false, "connect to server using tls")
	host := flag.String("diam_host", "ocs1p", "diameter identity host")
	realm := flag.String("diam_realm", "ocs.mnc82.mcc418.3gppnetwork.org", "diameter identity realm")
	certFile := flag.String("cert_file", "", "tls client certificate file (optional)")
	keyFile := flag.String("key_file", "", "tls client key file (optional)")
	initiate := flag.Bool("initiate", false, "send a initiate message, wait for the response and disconnect")
	networkType := flag.String("network_type", "tcp", "protocol type tcp/sctp")

	flag.Parse()
	if len(*addr) == 0 {
		flag.Usage()
	}

	// Load our custom dictionary on top of the default one.
	err := dict.Default.Load(bytes.NewReader([]byte(helloDictionary)))
	if err != nil {
		log.Fatal(err)
	}

	cfg := &sm.Settings{
		OriginHost:       datatype.DiameterIdentity(*host),
		OriginRealm:      datatype.DiameterIdentity(*realm),
		VendorID:         13,
		ProductName:      "go-diameter",
		OriginStateID:    datatype.Unsigned32(time.Now().Unix()),
		FirmwareRevision: 1,
		HostIPAddresses: []datatype.Address{
			datatype.Address(net.ParseIP("127.0.0.1")),
		},
	}

	// Create the state machine (it's a diam.ServeMux) and client.
	mux := sm.New(cfg)

	cli := &sm.Client{
		Dict:               dict.Default,
		Handler:            mux,
		MaxRetransmits:     3,
		RetransmitInterval: time.Second,
		EnableWatchdog:     true,
		WatchdogInterval:   5 * time.Second,
		AcctApplicationID: []*diam.AVP{
			// Advertise that we want support accounting application with id 999
			diam.NewAVP(avp.AcctApplicationID, avp.Mbit, 0, datatype.Unsigned32(helloApplication)),
		},
		AuthApplicationID: []*diam.AVP{
			// Advertise support for credit control application
			diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(4)), // RFC 4006
		},
	}

	// Set message handlers.
	done := make(chan struct{}, 1000)
	mux.Handle("CCA", handleCCA(done))

	// Print error reports.
	go printErrors(mux.ErrorReports())

	connect := func() (diam.Conn, error) {
		return dial(cli, *addr, *certFile, *keyFile, *ssl, *networkType)
	}

	if *initiate {
		c, err := connect()
		if err != nil {
			log.Fatal(err)
		}
		err = sendInitiate(c, cfg)
		if err != nil {
			log.Fatal(err)
		}
		select {
		case <-done:
		case <-time.After(5 * time.Second):
			log.Fatal("timeout: no hello answer received")
		}
		return
	}

	// Makes a persisent connection with back-off.
	log.Println("Use wireshark to see the messages, or try -hello")
	backoff := 1
	for {
		c, err := connect()
		if err != nil {
			log.Println(err)
			backoff *= 2
			if backoff > 20 {
				backoff = 20
			}
			time.Sleep(time.Duration(backoff) * time.Second)
			continue
		}
		log.Println("Client connected, handshake ok")
		backoff = 1
		<-c.(diam.CloseNotifier).CloseNotify()
		log.Println("Client disconnected")
	}
}

func printErrors(ec <-chan *diam.ErrorReport) {
	for err := range ec {
		log.Println(err)
	}
}

func dial(cli *sm.Client, addr, cert, key string, ssl bool, networkType string) (diam.Conn, error) {
	if ssl {
		return cli.DialNetworkTLS(networkType, addr, cert, key, nil)
	}
	return cli.DialNetwork(networkType, addr)
}

func sendInitiate(c diam.Conn, cfg *sm.Settings) error {
	// meta, ok := smpeer.FromContext(c.Context())
	// if !ok {
	// return errors.New("peer metadata unavailable")
	// }

	/*
		m.NewAVP("Session-Id", avp.Mbit, 0, datatype.UTF8String("890f81bee22a0dfddc8b9037eb367781cea1f328"))
		m.NewAVP("Service-Information", avp.Mbit, 10415, &GroupedAVP{
			AVP: []*AVP{
				NewAVP(20300, avp.Mbit, 20300, &GroupedAVP{ // IN-Information
					AVP: []*AVP{
						NewAVP(20339, avp.Mbit, 20300, datatype.Unsigned32(0)),  // Charge-Flow-Type
						NewAVP(20302, avp.Mbit, 20300, datatype.UTF8String("")), // Calling-Vlr-Number
					},
				}),
			}})
	*/
	m := diam.NewRequest(diam.CreditControl, 4, c.Dictionary())
	sid := "ocs1p;1587022628;275387;439986;1590579224;1097744"
	m.NewAVP(avp.SessionID, avp.Mbit, 0, datatype.UTF8String(sid))
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, cfg.OriginHost)
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, cfg.OriginRealm)
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("ocs.mnc82.mcc418.3gppnetwork.org"))
	m.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(4))
	m.NewAVP(avp.ServiceContextID, avp.Mbit, 0, datatype.UTF8String("ocs@iqonline.com"))
	m.NewAVP(avp.CCRequestType, avp.Mbit, 0, datatype.Enumerated(1))
	m.NewAVP(avp.CCRequestNumber, avp.Mbit, 0, datatype.Unsigned32(0))
	m.NewAVP(avp.DestinationHost, avp.Mbit, 0, datatype.DiameterIdentity("vepcocs"))
	m.NewAVP(avp.UserName, avp.Mbit, 0, datatype.UTF8String("internet"))
	m.NewAVP(avp.OriginStateID, avp.Mbit, 0, datatype.Unsigned32(25))
	m.NewAVP(avp.EventTimestamp, avp.Mbit, 0, datatype.Time(time.Unix(1377093974, 0)))
	m.NewAVP(avp.SubscriptionID, avp.Mbit, 0, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.SubscriptionIDType, avp.Mbit, 0, datatype.Enumerated(0)),
			diam.NewAVP(avp.SubscriptionIDData, avp.Mbit, 0, datatype.UTF8String("963500155039")),
		},
	})
	m.NewAVP(avp.SubscriptionID, avp.Mbit, 0, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.SubscriptionIDType, avp.Mbit, 0, datatype.Enumerated(1)),
			diam.NewAVP(avp.SubscriptionIDData, avp.Mbit, 0, datatype.UTF8String("417500011046039")),
		},
	})
	m.NewAVP(avp.MultipleServicesIndicator, avp.Mbit, 0, datatype.Enumerated(1))
	m.NewAVP(avp.MultipleServicesCreditControl, avp.Mbit, 0, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.RequestedServiceUnit, avp.Mbit, 0, &diam.GroupedAVP{
				AVP: []*diam.AVP{
					diam.NewAVP(avp.CCTotalOctets, avp.Mbit, 0, datatype.Unsigned64(5242880)),
				},
			}),
			diam.NewAVP(avp.ServiceIdentifier, avp.Mbit, 0, datatype.Unsigned32(999)),
			diam.NewAVP(avp.RatingGroup, avp.Mbit, 0, datatype.Unsigned32(999)),
		},
	})
	m.NewAVP(avp.UserEquipmentInfo, avp.Mbit, 0, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.UserEquipmentInfoType, avp.Mbit, 0, datatype.Enumerated(0)),
			diam.NewAVP(avp.UserEquipmentInfoValue, avp.Mbit, 0, datatype.OctetString("33353739363531303137313934383031")),
		},
	})
	m.NewAVP(avp.ServiceInformation, avp.Mbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.PSInformation, avp.Mbit, 10415, &diam.GroupedAVP{
				AVP: []*diam.AVP{
					diam.NewAVP(avp.PDPAddress, avp.Mbit, 10415, datatype.Address("10.168.10.1")),
					// diam.NewAVP(avp.PDPAddress, avp.Mbit, 10415, datatype.UTF8String("10.168.10.1")),
				},
			}),
		},
	})
	log.Printf("Sending HMR to %s\n%s", c.RemoteAddr(), m)
	_, err := m.WriteTo(c)
	return err
}

func handleCCA(done chan struct{}) diam.HandlerFunc {
	return func(c diam.Conn, m *diam.Message) {
		log.Printf("Received HMA from %s\n%s", c.RemoteAddr(), m)
		close(done)
	}
}

var eventRecord = datatype.Unsigned32(1) // RFC 6733: EVENT_RECORD 1

const (
	helloApplication = 999 // Our custom app from the dictionary below.
	helloMessage     = 111
)

// helloDictionary is our custom, example dictionary.
var helloDictionary = xml.Header + `
<diameter>
	<application id="999" type="acct">
		<command code="111" short="HM" name="Hello-Message">
			<request>
				<rule avp="Session-Id" required="true" max="1"/>
				<rule avp="Origin-Host" required="true" max="1"/>
				<rule avp="Origin-Realm" required="true" max="1"/>
				<rule avp="Destination-Realm" required="true" max="1"/>
				<rule avp="Destination-Host" required="true" max="1"/>
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
