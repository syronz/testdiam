CER affirm -> cgrates

    Origin-Host (264): "ocs1p"
    Origin-Realm (296): ocs.mnc82.mcc418.3gppnetwork.org
    Host-IP-Address Address (257): 172.22.21.6
    VendorId (266): 3GPP (10415)
    Product-Name (269): AN-3000
    Origin-State-Id (278): 25
    Supported-Vendor-Id (265): 3GPP 10415
    Supported-Vendor-Id (265): Vodafone Information Technology and Technology Management (12645)
    Auth-Application-Id (258): Diameter Credit Control Application (4)
    Firmware-Revision (267): 1

// ----------------------------------------------

CEA cgrates -> affirm

    Result-Code (268): DIAMETER_SUCCESS (2001)
    Origin-Host (264): ocs.iqonline.com
    Origin-Realm (296): ocs.mnc82.mcc418.3gppnetwork.org
    Host-IP-Address (257): 3137322e32322e32352e31
    VendorId (266): Reserved (0)
    Product-Name (269): CGRateS
    Origin-State-Id (278): 25
    AVP: Acct-Application-Id(259) l=12 f=-M- val=Diameter Base Accounting (3)
    AVP: Auth-Application-Id(258) l=12 f=-M- val=Diameter Credit Control Application (4)
    AVP: Supported-Vendor-Id(265) l=12 f=-M- val=10415
    AVP: Vendor-Specific-Application-Id(260) l=32 f=-M-
      AVP: Vendor-Id(266) l=12 f=-M- val=10415
      AVP: Auth-Application-Id(258) l=12 f=-M- val=3GPP Gx (16777238)
    AVP: Auth-Application-Id(258) l=12 f=-M- val=NASREQ Application (1)
    AVP: Supported-Vendor-Id(265) l=12 f=-M- val=10415
    AVP: Firmware-Revision(267) l=12 f=--- val=918

// ----------------------------------------------

CCR affirm -> cgrates [INITIAL_REQUEST]

    AVP: Session-Id(263) l=57 f=-M- val=ocs1p;1587022628;275387;439986;1590579224;1097744
    AVP: Origin-Host(264) l=13 f=-M- val=ocs1p
    AVP: Origin-Realm(296) l=40 f=-M- val=ocs.mnc82.mcc418.3gppnetwork.org
    AVP: Destination-Realm(283) l=40 f=-M- val=ocs.mnc82.mcc418.3gppnetwork.org
    AVP: Auth-Application-Id(258) l=12 f=-M- val=Diameter Credit Control Application (4)
    AVP: Service-Context-Id(461) l=24 f=-M- val=ocs@iqonline.com
    AVP: CC-Request-Type(416) l=12 f=-M- val=INITIAL_REQUEST (1)
    AVP: CC-Request-Number(415) l=12 f=-M- val=0
    AVP: Destination-Host(293) l=15 f=-M- val=vepcocs
    AVP: User-Name(1) l=16 f=-M- val=internet
    AVP: Origin-State-Id(278) l=12 f=-M- val=25
    AVP: Event-Timestamp(55) l=12 f=-M- val=May 27, 2020 11:33:44.000000000 UTC
    AVP: Subscription-Id(443) l=40 f=-M-
      AVP: Subscription-Id-Type(450) l=12 f=-M- val=END_USER_E164 (0)
      AVP: Subscription-Id-Data(444) l=20 f=-M- val=963500155039
    AVP: Subscription-Id(443) l=44 f=-M-
      AVP: Subscription-Id-Type(450) l=12 f=-M- val=END_USER_IMSI (1)
      AVP: Subscription-Id-Data(444) l=23 f=-M- val=417500011046039
    AVP: Multiple-Services-Indicator(455) l=12 f=-M- val=MULTIPLE_SERVICES_SUPPORTED (1)
    AVP: Multiple-Services-Credit-Control(456) l=56 f=-M-
      AVP: Requested-Service-Unit(437) l=24 f=-M-
        AVP: CC-Total-Octets(421) l=16 f=-M- val=5242880
      AVP: Service-Identifier(439) l=12 f=-M- val=999
      AVP: Rating-Group(432) l=12 f=-M- val=999
    AVP: User-Equipment-Info(458) l=44 f=-M-
      AVP: User-Equipment-Info-Type(459) l=12 f=-M- val=IMEISV (0)
      AVP: User-Equipment-Info-Value(460) l=24 f=-M- val=33353739363531303137313934383031

    AVP: Service-Information(873) l=652 f=VM- vnd=TGPP
      AVP: PS-Information(874) l=640 f=VM- vnd=TGPP
        AVP: 3GPP-Charging-Id(2) l=16 f=VM- vnd=TGPP val=00012800
        AVP: PDN-Connection-Charging-ID(2050) l=16 f=VM- vnd=TGPP val=75776
        AVP: 3GPP-PDP-Type(3) l=16 f=VM- vnd=TGPP val=IPv4 (0)
        AVP: PDP-Address(1227) l=18 f=VM- vnd=TGPP val=10.168.10.1
        AVP: Dynamic-Address-Flag(2051) l=16 f=VM- vnd=TGPP val=Dynamic (1)
        AVP: QoS-Information(1016) l=204 f=VM- vnd=TGPP
          AVP: QoS-Class-Identifier(1028) l=16 f=VM- vnd=TGPP val=QCI_9 (9)
          AVP: Max-Requested-Bandwidth-UL(516) l=16 f=VM- vnd=TGPP val=0
          AVP: Max-Requested-Bandwidth-DL(515) l=16 f=VM- vnd=TGPP val=0
          AVP: Guaranteed-Bitrate-UL(1026) l=16 f=VM- vnd=TGPP val=0
          AVP: Guaranteed-Bitrate-DL(1025) l=16 f=VM- vnd=TGPP val=0
          AVP: Bearer-Identifier(1020) l=17 f=VM- vnd=TGPP val=3735373736
          AVP: Allocation-Retention-Priority(1034) l=60 f=V-- vnd=TGPP
          AVP: APN-Aggregate-Max-Bitrate-UL(1041) l=16 f=V-- vnd=TGPP val=1000000000
          AVP: APN-Aggregate-Max-Bitrate-DL(1040) l=16 f=V-- vnd=TGPP val=1000000000
        AVP: 3GPP-GPRS-Negotiated-QoS-Profile(5) l=35 f=VM- vnd=TGPP val=08-0409000F4240000F4240
        AVP: 3GPP-IMSI-MCC-MNC(8) l=17 f=VM- vnd=TGPP val=41750
        AVP: IMSI-Unauthenticated-Flag(2308) l=16 f=VM- vnd=TGPP val=Authenticated (0)
        AVP: 3GPP-GGSN-MCC-MNC(9) l=17 f=VM- vnd=TGPP val=41750
        AVP: 3GPP-NSAPI(10) l=13 f=VM- vnd=TGPP val=5
        AVP: Called-Station-Id(30) l=13 f=-M- val=test2
        AVP: 3GPP-Selection-Mode(12) l=13 f=VM- vnd=TGPP val=0
        AVP: 3GPP-Charging-Characteristics(13) l=16 f=VM- vnd=TGPP val=0000
        AVP: 3GPP-SGSN-MCC-MNC(18) l=17 f=VM- vnd=TGPP val=41750
        AVP: 3GPP-MS-TimeZone(23) l=14 f=VM- vnd=TGPP val=Timezone: GMT + 4 hours 0 minutes +1 hour adjustment for Daylight Saving Time
        AVP: Charging-Rule-Base-Name(1004) l=18 f=VM- vnd=TGPP val=CI_One
        AVP: 3GPP-User-Location-Info(22) l=25 f=VM- vnd=TGPP val=MCC 417 Syrian Arab Republic, MNC 50 , ECGI 0x4e2014
        AVP: 3GPP-RAT-Type(21) l=13 f=VM- vnd=TGPP val=06
        AVP: PDP-Context-Type(1247) l=16 f=VM- vnd=TGPP val=PRIMARY (0)
        AVP: 3GPP-IMEISV(20) l=28 f=VM- vnd=TGPP val=8670700211943855
        AVP: SGSN-Address(1228) l=18 f=VM- vnd=TGPP val=172.22.21.8
        AVP: GGSN-Address(847) l=18 f=VM- vnd=TGPP val=172.22.43.10


// ----------------------------------------------

CCR affirm -> cgrates [UPDATE_REQUEST]

    AVP: Session-Id(263) l=56 f=-M- val=ocs1p;1587023009;188535;453670;1590612545;753680
    AVP: Origin-Host(264) l=13 f=-M- val=ocs1p
    AVP: Origin-Realm(296) l=40 f=-M- val=ocs.mnc82.mcc418.3gppnetwork.org
    AVP: Destination-Realm(283) l=40 f=-M- val=ocs.mnc82.mcc418.3gppnetwork.org
    AVP: Auth-Application-Id(258) l=12 f=-M- val=Diameter Credit Control Application (4)
    AVP: Service-Context-Id(461) l=24 f=-M- val=ocs@iqonline.com
    AVP: CC-Request-Type(416) l=12 f=-M- val=UPDATE_REQUEST (2)
    AVP: CC-Request-Number(415) l=12 f=-M- val=760
    AVP: Destination-Host(293) l=24 f=-M- val=ocs.iqonline.com
    AVP: User-Name(1) l=16 f=-M- val=internet
    AVP: Origin-State-Id(278) l=12 f=-M- val=25
    AVP: Event-Timestamp(55) l=12 f=-M- val=May 28, 2020 09:38:06.000000000 UTC
    AVP: Subscription-Id(443) l=40 f=-M-
      AVP: Subscription-Id-Type(450) l=12 f=-M- val=END_USER_E164 (0)
      AVP: Subscription-Id-Data(444) l=20 f=-M- val=963500156958
    AVP: Subscription-Id(443) l=44 f=-M-
      AVP: Subscription-Id-Type(450) l=12 f=-M- val=END_USER_IMSI (1)
      AVP: Subscription-Id-Data(444) l=23 f=-M- val=417500011047958
    AVP: Multiple-Services-Indicator(455) l=12 f=-M- val=MULTIPLE_SERVICES_SUPPORTED (1)
    AVP: Multiple-Services-Credit-Control(456) l=104 f=-M-
      AVP: Used-Service-Unit(446) l=56 f=-M-
        AVP: CC-Total-Octets(421) l=16 f=-M- val=0
        AVP: CC-Input-Octets(412) l=16 f=-M- val=0
        AVP: CC-Output-Octets(414) l=16 f=-M- val=0
      AVP: Service-Identifier(439) l=12 f=-M- val=999
      AVP: Rating-Group(432) l=12 f=-M- val=999
      AVP: 3GPP-Reporting-Reason(872) l=16 f=VM- vnd=TGPP val=QHT (1)
    AVP: User-Equipment-Info(458) l=44 f=-M-
      AVP: User-Equipment-Info-Type(459) l=12 f=-M- val=IMEISV (0)
      AVP: User-Equipment-Info-Value(460) l=24 f=-M- val=33353334313631303639333233343031
    AVP: Service-Information(873) l=628 f=VM- vnd=TGPP
      AVP: PS-Information(874) l=616 f=VM- vnd=TGPP

// ----------------------------------------------

CCR affirm -> cgrates [TERMINATION_REQUEST]

    AVP: Session-Id(263) l=57 f=-M- val=ocs1p;1587022627;280780;455076;1590651308;1122320
    AVP: Origin-Host(264) l=13 f=-M- val=ocs1p
    AVP: Origin-Realm(296) l=40 f=-M- val=ocs.mnc82.mcc418.3gppnetwork.org
    AVP: Destination-Realm(283) l=40 f=-M- val=ocs.mnc82.mcc418.3gppnetwork.org
    AVP: Auth-Application-Id(258) l=12 f=-M- val=Diameter Credit Control Application (4)
    AVP: Service-Context-Id(461) l=24 f=-M- val=ocs@iqonline.com
    AVP: CC-Request-Type(416) l=12 f=-M- val=TERMINATION_REQUEST (3)
    AVP: CC-Request-Number(415) l=12 f=-M- val=121
    AVP: Destination-Host(293) l=24 f=-M- val=ocs.iqonline.com
    AVP: User-Name(1) l=16 f=-M- val=internet
    AVP: Origin-State-Id(278) l=12 f=-M- val=25
    AVP: Event-Timestamp(55) l=12 f=-M- val=May 28, 2020 09:39:43.000000000 UTC
    AVP: Subscription-Id(443) l=40 f=-M-
      AVP: Subscription-Id-Type(450) l=12 f=-M- val=END_USER_E164 (0)
      AVP: Subscription-Id-Data(444) l=20 f=-M- val=963500156958
    AVP: Subscription-Id(443) l=44 f=-M-
      AVP: Subscription-Id-Type(450) l=12 f=-M- val=END_USER_IMSI (1)
      AVP: Subscription-Id-Data(444) l=23 f=-M- val=417500011047958
    AVP: Termination-Cause(295) l=12 f=-M- val=DIAMETER_LOGOUT (1)
    AVP: Multiple-Services-Indicator(455) l=12 f=-M- val=MULTIPLE_SERVICES_SUPPORTED (1)
    AVP: Multiple-Services-Credit-Control(456) l=104 f=-M-
      AVP: Used-Service-Unit(446) l=56 f=-M-
        AVP: CC-Total-Octets(421) l=16 f=-M- val=0
        AVP: CC-Input-Octets(412) l=16 f=-M- val=0
        AVP: CC-Output-Octets(414) l=16 f=-M- val=0
      AVP: Service-Identifier(439) l=12 f=-M- val=999
      AVP: Rating-Group(432) l=12 f=-M- val=999
      AVP: 3GPP-Reporting-Reason(872) l=16 f=VM- vnd=TGPP val=FINAL (2)
    AVP: User-Equipment-Info(458) l=44 f=-M-
      AVP: User-Equipment-Info-Type(459) l=12 f=-M- val=IMEISV (0)
      AVP: User-Equipment-Info-Value(460) l=24 f=-M- val=33353334313631303639333233343031
    AVP: Service-Information(873) l=644 f=VM- vnd=TGPP





// ----------------------------------------------

CCA cgrates -> affirm [INITIAL_REQUEST]

    AVP: Result-Code(268) l=12 f=-M- val=DIAMETER_SUCCESS (2001)
    AVP: Session-Id(263) l=57 f=-M- val=ocs1p;1587022628;275387;439986;1590579224;1097744
    AVP: Origin-Host(264) l=24 f=-M- val=ocs.iqonline.com
    AVP: Origin-Realm(296) l=40 f=-M- val=ocs.mnc82.mcc418.3gppnetwork.org
    AVP: Auth-Application-Id(258) l=12 f=-M- val=Diameter Credit Control Application (4)
    AVP: CC-Request-Type(416) l=12 f=-M- val=INITIAL_REQUEST (1)
    AVP: CC-Request-Number(415) l=12 f=-M- val=0
    AVP: Multiple-Services-Credit-Control(456) l=80 f=-M-
      AVP: Granted-Service-Unit(431) l=24 f=-M-
        AVP: CC-Total-Octets(421) l=16 f=-M- val=5242880
      AVP: Rating-Group(432) l=12 f=-M- val=999
      AVP: Service-Identifier(439) l=12 f=-M- val=999
      AVP: Validity-Time(448) l=12 f=-M- val=600
      AVP: Result-Code(268) l=12 f=-M- val=DIAMETER_SUCCESS (2001)
    ???????
    AVP: Multiple-Services-Credit-Control(456) l=80 f=-M-
      AVP: Final-Unit-Indication(430) l=68 f=-M-
        AVP: Final-Unit-Action(449) l=12 f=-M- val=REDIRECT (1)
        AVP: Redirect-Server(434) l=48 f=-M-
          AVP: Redirect-Address-Type(433) l=12 f=-M- val=URL (2)
          AVP: Redirect-Server-Address(435) l=27 f=-M- val=http://my.rcell.me/
      AVP: Rating-Group(432) l=12 f=-M- val=999
      AVP: Service-Identifier(439) l=12 f=-M- val=999
      AVP: Result-Code(268) l=12 f=-M- val=DIAMETER_CREDIT_LIMIT_REACHED (4012)



  [UPDATE_REQUEST]

    AVP: Multiple-Services-Credit-Control(456) l=148 f=-M-
      AVP: Final-Unit-Indication(430) l=68 f=-M-
        AVP: Final-Unit-Action(449) l=12 f=-M- val=REDIRECT (1)
        AVP: Redirect-Server(434) l=48 f=-M-
          AVP: Redirect-Address-Type(433) l=12 f=-M- val=URL (2)
          AVP: Redirect-Server-Address(435) l=27 f=-M- val=http://my.rcell.me/
      AVP: Granted-Service-Unit(431) l=24 f=-M-
        AVP: CC-Total-Octets(421) l=16 f=-M- val=5242880
      AVP: Rating-Group(432) l=12 f=-M- val=999
      AVP: Service-Identifier(439) l=12 f=-M- val=999
      AVP: Validity-Time(448) l=12 f=-M- val=60
      AVP: Result-Code(268) l=12 f=-M- val=DIAMETER_SUCCESS (2001)








