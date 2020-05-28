 CER client -> server

    AVP: Origin-Host(264) l=14 f=-M- val=client [OK]
    AVP: Origin-Realm(296) l=19 f=-M- val=go-diameter [OK]
    AVP: Host-IP-Address(257) l=14 f=-M- val=127.0.0.1 [OK]
    AVP: Vendor-Id(266) l=12 f=-M- val=13 [OK]
    AVP: Product-Name(269) l=19 f=--- val=go-diameter [OK]
    AVP: Origin-State-Id(278) l=12 f=-M- val=1590650433 [OK]
    AVP: Auth-Application-Id(258) l=12 f=-M- val=Diameter Credit Control Application (4) [OK]
    AVP: Inband-Security-Id(299) l=12 f=-M- val=NO_INBAND_SECURITY (0) [NO]
    AVP: Acct-Application-Id(259) l=12 f=-M- val=Unknown (999) [NO]
    AVP: Firmware-Revision(267) l=12 f=--- val=1 [OK]

// ----------------------------------------------

CEA server -> client

    AVP: Result-Code(268) l=12 f=-M- val=DIAMETER_SUCCESS (2001) [OK]
    AVP: Origin-Host(264) l=14 f=-M- val=server [OK]
    AVP: Origin-Realm(296) l=19 f=-M- val=go-diameter [OK]
    AVP: Host-IP-Address(257) l=26 f=-M- val=::1 [?]
    AVP: Vendor-Id(266) l=12 f=-M- val=13 [?]
    AVP: Product-Name(269) l=19 f=--- val=go-diameter [OK]
    AVP: Origin-State-Id(278) l=12 f=-M- val=1590650433 [OK]
    AVP: Acct-Application-Id(259) l=12 f=-M- val=Diameter Base Accounting (3) [OK]
    AVP: Auth-Application-Id(258) l=12 f=-M- val=Diameter Credit Control Application (4) [OK]
    AVP: Supported-Vendor-Id(265) l=12 f=-M- val=10415 [OK]
    AVP: Vendor-Specific-Application-Id(260) l=32 f=-M- 
      AVP: Vendor-Id(266) l=12 f=-M- val=10415 [OK]
      AVP: Auth-Application-Id(258) l=12 f=-M- val=3GPP Gx (16777238) [OK]
    AVP: Auth-Application-Id(258) l=12 f=-M- val=NASREQ Application (1) [OK]
    AVP: Supported-Vendor-Id(265) l=12 f=-M- val=10415 [DUPLICATE]
    AVP: Vendor-Specific-Application-Id(260) l=32 f=-M- [DUPLICATE]
    AVP: Supported-Vendor-Id(265) l=12 f=-M- val=10415 [DUPLICATE]
    AVP: Vendor-Specific-Application-Id(260) l=32 f=-M- [DUPLICATE]
    AVP: Supported-Vendor-Id(265) l=12 f=-M- val=10415 [DUPLICATE]
    AVP: Vendor-Specific-Application-Id(260) l=32 f=-M- [DUPLICATE]
    AVP: Acct-Application-Id(259) l=12 f=-M- val=Unknown (999) [EXTRA]
    AVP: Firmware-Revision(267) l=12 f=--- val=1 [OK]
