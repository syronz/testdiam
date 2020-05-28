                                           +--------------------+
                                            |    AVP Flag rules  |
                                            |----+-----+----+----|----+
                     AVP  Section           |    |     |SHLD|MUST|    |
   Attribute Name    Code Defined Data Type |MUST| MAY | NOT|NOT |Encr|
   -----------------------------------------|----+-----+----+----|----|
   CC-Correlation-Id 411  8.1    OctetString|    | P,M |    |  V | Y  |
   CC-Input-Octets   412  8.24   Unsigned64 | M  |  P  |    |  V | Y  |
   CC-Money          413  8.22   Grouped    | M  |  P  |    |  V | Y  |
   CC-Output-Octets  414  8.25   Unsigned64 | M  |  P  |    |  V | Y  |
   CC-Request-Number 415  8.2    Unsigned32 | M  |  P  |    |  V | Y  |
   CC-Request-Type   416  8.3    Enumerated | M  |  P  |    |  V | Y  |
   CC-Service-       417  8.26   Unsigned64 | M  |  P  |    |  V | Y  |
     Specific-Units                         |    |     |    |    |    |
   CC-Session-       418  8.4    Enumerated | M  |  P  |    |  V | Y  |
     Failover                               |    |     |    |    |    |
   CC-Sub-Session-Id 419  8.5    Unsigned64 | M  |  P  |    |  V | Y  |
   CC-Time           420  8.21   Unsigned32 | M  |  P  |    |  V | Y  |
   CC-Total-Octets   421  8.23   Unsigned64 | M  |  P  |    |  V | Y  |
   CC-Unit-Type      454  8.32   Enumerated | M  |  P  |    |  V | Y  |
   Check-Balance-    422  8.6    Enumerated | M  |  P  |    |  V | Y  |
     Result                                 |    |     |    |    |    |
   Cost-Information  423  8.7    Grouped    | M  |  P  |    |  V | Y  |
   Cost-Unit         424  8.12   UTF8String | M  |  P  |    |  V | Y  |
   Credit-Control    426  8.13   Enumerated | M  |  P  |    |  V | Y  |
   Credit-Control-   427  8.14   Enumerated | M  |  P  |    |  V | Y  |
     Failure-Handling                       |    |     |    |    |    |
   Currency-Code     425  8.11   Unsigned32 | M  |  P  |    |  V | Y  |
   Direct-Debiting-  428  8.15   Enumerated | M  |  P  |    |  V | Y  |
     Failure-Handling                       |    |     |    |    |    |
   Exponent          429  8.9    Integer32  | M  |  P  |    |  V | Y  |
   Final-Unit-Action 449  8.35   Enumerated | M  |  P  |    |  V | Y  |
   Final-Unit-       430  8.34   Grouped    | M  |  P  |    |  V | Y  |
     Indication                             |    |     |    |    |    |
   Granted-Service-  431  8.17   Grouped    | M  |  P  |    |  V | Y  |
     Unit                                   |    |     |    |    |    |
   G-S-U-Pool-       453  8.31   Unsigned32 | M  |  P  |    |  V | Y  |
     Identifier                             |    |     |    |    |    |




Hakala, et al.              Standards Track                    [Page 56]
 
RFC 4006          Diameter Credit-Control Application        August 2005


   G-S-U-Pool-       457  8.30   Grouped    | M  |  P  |    |  V | Y  |
     Reference                              |    |     |    |    |    |
   Multiple-Services 456  8.16   Grouped    | M  |  P  |    |  V | Y  |
    -Credit-Control                         |    |     |    |    |    |
   Multiple-Services 455  8.40   Enumerated | M  |  P  |    |  V | Y  |
    -Indicator                              |    |     |    |    |    |
   Rating-Group      432  8.29   Unsigned32 | M  |  P  |    |  V | Y  |
   Redirect-Address  433  8.38   Enumerated | M  |  P  |    |  V | Y  |
     -Type                                  |    |     |    |    |    |
   Redirect-Server   434  8.37   Grouped    | M  |  P  |    |  V | Y  |
   Redirect-Server   435  8.39   UTF8String | M  |  P  |    |  V | Y  |
     -Address                               |    |     |    |    |    |
   Requested-Action  436  8.41   Enumerated | M  |  P  |    |  V | Y  |
   Requested-Service 437  8.18   Grouped    | M  |  P  |    |  V | Y  |
     -Unit                                  |    |     |    |    |    |
   Restriction       438  8.36   IPFiltrRule| M  |  P  |    |  V | Y  |
     -Filter-Rule                           |    |     |    |    |    |
   Service-Context   461  8.42   UTF8String | M  |  P  |    |  V | Y  |
     -Id                                    |    |     |    |    |    |
   Service-          439  8.28   Unsigned32 | M  |  P  |    |  V | Y  |
     Identifier                             |    |     |    |    |    |
   Service-Parameter 440  8.43   Grouped    |    | P,M |    |  V | Y  |
     -Info                                  |    |     |    |    |    |
   Service-          441  8.44   Unsigned32 |    | P,M |    |  V | Y  |
     Parameter-Type                         |    |     |    |    |    |
   Service-          442  8.45   OctetString|    | P,M |    |  V | Y  |
     Parameter-Value                        |    |     |    |    |    |
   Subscription-Id   443  8.46   Grouped    | M  |  P  |    |  V | Y  |
   Subscription-Id   444  8.48   UTF8String | M  |  P  |    |  V | Y  |
     -Data                                  |    |     |    |    |    |
   Subscription-Id   450  8.47   Enumerated | M  |  P  |    |  V | Y  |
     -Type                                  |    |     |    |    |    |
   Tariff-Change     452  8.27   Enumerated | M  |  P  |    |  V | Y  |
     -Usage                                 |    |     |    |    |    |
   Tariff-Time       451  8.20   Time       | M  |  P  |    |  V | Y  |
     -Change                                |    |     |    |    |    |
   Unit-Value        445  8.8    Grouped    | M  |  P  |    |  V | Y  |
   Used-Service-Unit 446  8.19   Grouped    | M  |  P  |    |  V | Y  |
   User-Equipment    458  8.49   Grouped    |    | P,M |    |  V | Y  |
     -Info                                  |    |     |    |    |    |
   User-Equipment    459  8.50   Enumerated |    | P,M |    |  V | Y  |
     -Info-Type                             |    |     |    |    |    |
   User-Equipment    460  8.51   OctetString|    | P,M |    |  V | Y  |
     -Info-Value                            |    |     |    |    |    |
   Value-Digits      447  8.10   Integer64  | M  |  P  |    |  V | Y  |
   Validity-Time     448  8.33   Unsigned32 | M  |  P  |    |  V | Y  |



                                            +----------+
                                            | AVP Flag |
                                            |  rules   |
                                            |----+-----|
                   AVP  Section             |    |MUST |
   Attribute Name  Code Defined  Data Type  |MUST| NOT |
   -----------------------------------------|----+-----|
   Acct-             85  9.8.2   Unsigned32 | M  |  V  |
     Interim-Interval                       |    |     |
   Accounting-      483  9.8.7   Enumerated | M  |  V  |
     Realtime-Required                      |    |     |
   Acct-            50   9.8.5   UTF8String | M  |  V  |
     Multi-Session-Id                       |    |     |
   Accounting-      485  9.8.3   Unsigned32 | M  |  V  |
     Record-Number                          |    |     |
   Accounting-      480  9.8.1   Enumerated | M  |  V  |
     Record-Type                            |    |     |
   Acct-             44  9.8.4   OctetString| M  |  V  |
    Session-Id                              |    |     |
   Accounting-      287  9.8.6   Unsigned64 | M  |  V  |
     Sub-Session-Id                         |    |     |
   Acct-            259  6.9     Unsigned32 | M  |  V  |
     Application-Id                         |    |     |
   Auth-            258  6.8     Unsigned32 | M  |  V  |
     Application-Id                         |    |     |
   Auth-Request-    274  8.7     Enumerated | M  |  V  |
      Type                                  |    |     |
   Authorization-   291  8.9     Unsigned32 | M  |  V  |
     Lifetime                               |    |     |
   Auth-Grace-      276  8.10    Unsigned32 | M  |  V  |
     Period                                 |    |     |
   Auth-Session-    277  8.11    Enumerated | M  |  V  |
     State                                  |    |     |
   Re-Auth-Request- 285  8.12    Enumerated | M  |  V  |
     Type                                   |    |     |
   Class             25  8.20    OctetString| M  |  V  |
   Destination-Host 293  6.5     DiamIdent  | M  |  V  |
   Destination-     283  6.6     DiamIdent  | M  |  V  |
     Realm                                  |    |     |
   Disconnect-Cause 273  5.4.3   Enumerated | M  |  V  |
   Error-Message    281  7.3     UTF8String |    | V,M |
   Error-Reporting- 294  7.4     DiamIdent  |    | V,M |
     Host                                   |    |     |
   Event-Timestamp   55  8.21    Time       | M  |  V  |
   Experimental-    297  7.6     Grouped    | M  |  V  |
      Result                                |    |     |
   -----------------------------------------|----+-----|




Fajardo, et al.              Standards Track                   [Page 56]
 
RFC 6733                 Diameter Base Protocol             October 2012


                                            +----------+
                                            | AVP Flag |
                                            |  rules   |
                                            |----+-----|
                   AVP  Section             |    |MUST |
   Attribute Name  Code Defined  Data Type  |MUST| NOT |
   -----------------------------------------|----+-----|
   Experimental-    298  7.7     Unsigned32 | M  |  V  |
      Result-Code                           |    |     |
   Failed-AVP       279  7.5     Grouped    | M  |  V  |
   Firmware-        267  5.3.4   Unsigned32 |    | V,M |
     Revision                               |    |     |
   Host-IP-Address  257  5.3.5   Address    | M  |  V  |
   Inband-Security                          | M  |  V  |
      -Id           299  6.10    Unsigned32 |    |     |
   Multi-Round-     272  8.19    Unsigned32 | M  |  V  |
     Time-Out                               |    |     |
   Origin-Host      264  6.3     DiamIdent  | M  |  V  |
   Origin-Realm     296  6.4     DiamIdent  | M  |  V  |
   Origin-State-Id  278  8.16    Unsigned32 | M  |  V  |
   Product-Name     269  5.3.7   UTF8String |    | V,M |
   Proxy-Host       280  6.7.3   DiamIdent  | M  |  V  |
   Proxy-Info       284  6.7.2   Grouped    | M  |  V  |
   Proxy-State       33  6.7.4   OctetString| M  |  V  |
   Redirect-Host    292  6.12    DiamURI    | M  |  V  |
   Redirect-Host-   261  6.13    Enumerated | M  |  V  |
      Usage                                 |    |     |
   Redirect-Max-    262  6.14    Unsigned32 | M  |  V  |
      Cache-Time                            |    |     |
   Result-Code      268  7.1     Unsigned32 | M  |  V  |
   Route-Record     282  6.7.1   DiamIdent  | M  |  V  |
   Session-Id       263  8.8     UTF8String | M  |  V  |
   Session-Timeout   27  8.13    Unsigned32 | M  |  V  |
   Session-Binding  270  8.17    Unsigned32 | M  |  V  |
   Session-Server-  271  8.18    Enumerated | M  |  V  |
     Failover                               |    |     |
   Supported-       265  5.3.6   Unsigned32 | M  |  V  |
     Vendor-Id                              |    |     |
   Termination-     295  8.15    Enumerated | M  |  V  |
      Cause                                 |    |     |
   User-Name          1  8.14    UTF8String | M  |  V  |
   Vendor-Id        266  5.3.3   Unsigned32 | M  |  V  |
   Vendor-Specific- 260  6.11    Grouped    | M  |  V  |
      Application-Id                        |    |     |
   -----------------------------------------|----+-----|



