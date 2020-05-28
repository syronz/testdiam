affirm CER request

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

CEA from cgrates

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








