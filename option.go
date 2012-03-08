package dhcp

const (
	OptionCodePad                                        = 0
	OptionCodeEnd                                        = 255
	OptionCodeSubnetMask                                 = 1
	OptionCodeTimeOffset                                 = 2
	OptionCodeRouter                                     = 3
	OptionCodeTimeServer                                 = 4
	OptionCodeNameServer                                 = 5
	OptionCodeDomainNameServer                           = 6
	OptionCodeLogServer                                  = 7
	OptionCodeCookieServer                               = 8
	OptionCodeLPRServer                                  = 9
	OptionCodeImpressServer                              = 10
	OptionCodeResourceLocationServer                     = 11
	OptionCodeHostName                                   = 12
	OptionCodeBootFileSize                               = 13
	OptionCodeMeritDumpFile                              = 14
	OptionCodeDomainName                                 = 15
	OptionCodeSwapServer                                 = 16
	OptionCodeRootPath                                   = 17
	OptionCodeExtensionsPath                             = 18
	OptionCodeIPForwardingEnableDisable                  = 19
	OptionCodeNonLocalSourceRoutingEnableDisable         = 20
	OptionCodePolicyFilter                               = 21
	OptionCodeMaximumDatagramReassemblySize              = 22
	OptionCodeDefaultIPTimetolive                        = 23
	OptionCodePathMTUAgingTimeout                        = 24
	OptionCodePathMTUPlateauTable                        = 25
	OptionCodeInterfaceMTU                               = 26
	OptionCodeAllSubnetsareLocal                         = 27
	OptionCodeBroadcastAddress                           = 28
	OptionCodePerformMaskDiscovery                       = 29
	OptionCodeMaskSupplier                               = 30
	OptionCodePerformRouterDiscovery                     = 31
	OptionCodeRouterSolicitationAddress                  = 32
	OptionCodeStaticRoute                                = 33
	OptionCodeTrailerEncapsulation                       = 34
	OptionCodeARPCacheTimeout                            = 35
	OptionCodeEthernetEncapsulation                      = 36
	OptionCodeTCPDefaultTTL                              = 37
	OptionCodeTCPKeepaliveInterval                       = 38
	OptionCodeTCPKeepaliveGarbage                        = 39
	OptionCodeNetworkInformationServiceDomain            = 40
	OptionCodeNetworkInformationServers                  = 41
	OptionCodeNetworkTimeProtocolServers                 = 42
	OptionCodeVendorSpecificInformation                  = 43
	OptionCodeNetBIOSoverTCPIPNameServer                 = 44
	OptionCodeNetBIOSoverTCPIPDatagramDistributionServer = 45
	OptionCodeNetBIOSoverTCPIPNodeType                   = 46
	OptionCodeNetBIOSoverTCPIPScope                      = 47
	OptionCodeXWindowSystemFontServer                    = 48
	OptionCodeXWindowSystemDisplayManager                = 49
	OptionCodeNetworkInformationServicePlusDomain        = 64
	OptionCodeNetworkInformationServicePlusServers       = 65
	OptionCodeMobileIPHomeAgent                          = 68
	OptionCodeSMTPServer                                 = 69
	OptionCodePOP3Server                                 = 70
	OptionCodeNNTPServer                                 = 71
	OptionCodeWWWServer                                  = 72
	OptionCodeDefaultFingerServer                        = 73
	OptionCodeDefaultIRCServer                           = 74
	OptionCodeStreetTalkServer                           = 75
	OptionCodeStreetSTDAServer                           = 76
	OptionCodeDHCPRequestedIPAddress                     = 50
	OptionCodeDHCPIPAddressLeaseTime                     = 51
	OptionCodeDHCPOptionOverload                         = 52
	OptionCodeDHCPTFTPservername                         = 66
	OptionCodeDHCPBootfilename                           = 67
	OptionCodeDHCPMessageType                            = 53
	OptionCodeDHCPServerIdentifier                       = 54
	OptionCodeDHCPParameterRequestList                   = 55
	OptionCodeDHCPMessage                                = 56
	OptionCodeDHCPMaximumMessageSize                     = 57
	OptionCodeDHCPRenewalTimeValue                       = 58
	OptionCodeDHCPRebindingTimeValue                     = 59
	OptionCodeDHCPVendorclassidentifier                  = 60
	OptionCodeDHCPClientidentifier                       = 61
)

const (
	NetBIOSoverTCPIPNodeTypeB = 0x1
	NetBIOSoverTCPIPNodeTypeP = 0x2
	NetBIOSoverTCPIPNodeTypeM = 0x4
	NetBIOSoverTCPIPNodeTypeH = 0x8
)

const (
	DHCPMessageTypeDiscover = 1
	DHCPMessageTypeOffer    = 2
	DHCPMessageTypeRequest  = 3
	DHCPMessageTypeDecline  = 4
	DHCPMessageTypeAck      = 5
	DHCPMessageTypeNak      = 6
	DHCPMessageTypeRelease  = 7
	DHCPMessageTypeInform   = 8
)

type Option struct {
	Code   uint8
	Length uint8
	Value  []byte
}

func ReadOption(b []byte) (opt *Option, err error) {
	opt = new(Option)

	if len(b) < 1 {
		return nil, ErrInvalidFormat
	}

	opt.Code = uint8(b[0])
	if opt.Code == OptionCodeEnd || opt.Code == OptionCodePad {
		return
	}

  if len(b) < 2 {
    return nil, ErrInvalidFormat
  }

	opt.Length = uint8(b[1])
	if len(b) < int(opt.Length)+2 {
		return nil, ErrInvalidFormat
	}
	opt.Value = b[2 : opt.Length+2]

	return opt, nil
}
