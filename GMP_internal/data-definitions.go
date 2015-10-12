package GMP_internal

import (
	"sync"
)

var ClientIdentifier string = "ANONYMOUS AARDVARK"

// String Constants
const (
	ENGA_VERSION          string = "0.0.1"
	ENGA_USERAGENT_STRING        = "Evilnode GMP GoLang Lib"
	ENGA_CID_NAME                = "ENGAV0.0.1USERTOKEN"
	ENGA_ENDPOINT_SSL            = "https://ssl.google-analytics.com/collect"
)

// Hit Types
const (
	HitTypePageView = iota
	HitTypeAppView
	HitTypeEvent
	HitTypeTransaction
	HitTypeItem
	HitTypeSocial
	HitTypeException
	HitTypeTiming
)

// GMP Param Keys
const (
	KeyVersion = iota
	KeyHitType
	KeyTrackingID
	KeyClientID
	KeyCampaignName
	KeyCampaignSource
	KeyCampaignMedium
	KeyCampaignKeyword
	KeyCampaignContent
	KeyCampaignID
	KeyDocumentHostName
	KeyDocumentPage
	KeyDocumentTitle
	KeyAdwordsID
	KeyDisplayAdsID
	KeyShouldAnonymizeIP
	KeyShouldUseSessionControl
	KeyScreenResolution
	KeyViewportSize
	KeyDocumentEncoding
	KeyScreenColors
	KeyUserLanguage
	KeyAppName
	KeyAppVersion
	KeyTransactionID
	KeyTransactionAffiliation
	KeyTransactionRevenue
	KeyTransactionShipping
	KeyTransactionTax
	KeyNonInteractiveHit
	KeyContentDescription
	KeyLinkID
	KeyEventCategory
	KeyEventAction
	KeyEventLabel
	KeyEventValue
	KeyItemName
	KeyItemPrice
	KeyItemQuantity
	KeyItemCode
	KeyItemCategory
	KeyCurrencyCode
	KeySocialNetwork
	KeySocialAction
	KeySocialActionTarget
	KeyUserTimingCategory
	KeyUserTimingVariableName
	KeyUserTimingTime
	KeyUserTimingLabel
	KeyPageLoadTime
	KeyDNSTime
	KeyPageDownloadTime
	KeyRedirectResponseTime
	KeyTCPConnectTime
	KeyServerResponseTime
	KeyExceptionDescription
	KeyExceptionFatal
	KeyExperimentID
	KeyExperimentVariant
)

var paramKeys map[int]string
var paramKeysOnce sync.Once

// Get the GMP Param key for the selected field
func ParamKey(param int) (string, bool) {
	paramKeysOnce.Do(func() {
		paramKeys = map[int]string{
			KeyVersion:                 "v",
			KeyHitType:                 "t",
			KeyTrackingID:              "tid",
			KeyClientID:                "cid",
			KeyCampaignName:            "cn",
			KeyCampaignSource:          "cs",
			KeyCampaignMedium:          "cm",
			KeyCampaignKeyword:         "ck",
			KeyCampaignContent:         "cc",
			KeyCampaignID:              "ci",
			KeyDocumentHostName:        "dh",
			KeyDocumentPage:            "dp",
			KeyDocumentTitle:           "dt",
			KeyAdwordsID:               "gclid",
			KeyDisplayAdsID:            "dclid",
			KeyShouldAnonymizeIP:       "aip",
			KeyShouldUseSessionControl: "sc",
			KeyScreenResolution:        "sr",
			KeyViewportSize:            "vp",
			KeyDocumentEncoding:        "de",
			KeyScreenColors:            "sd",
			KeyUserLanguage:            "ul",
			KeyAppName:                 "an",
			KeyAppVersion:              "av",
			KeyTransactionID:           "ti",
			KeyTransactionAffiliation:  "ta",
			KeyTransactionRevenue:      "tr",
			KeyTransactionShipping:     "ts",
			KeyTransactionTax:          "tt",
			KeyNonInteractiveHit:       "ni",
			KeyContentDescription:      "cd",
			KeyLinkID:                  "linkid",
			KeyEventCategory:           "ec",
			KeyEventAction:             "ea",
			KeyEventLabel:              "el",
			KeyEventValue:              "ev",
			KeyItemName:                "in",
			KeyItemPrice:               "ip",
			KeyItemQuantity:            "iq",
			KeyItemCode:                "ic",
			KeyItemCategory:            "iv",
			KeyCurrencyCode:            "cu",
			KeySocialNetwork:           "sn",
			KeySocialAction:            "sa",
			KeySocialActionTarget:      "st",
			KeyUserTimingCategory:      "utc",
			KeyUserTimingVariableName:  "utv",
			KeyUserTimingTime:          "utt",
			KeyUserTimingLabel:         "utl",
			KeyPageLoadTime:            "plt",
			KeyDNSTime:                 "dns",
			KeyPageDownloadTime:        "pdt",
			KeyRedirectResponseTime:    "rrt",
			KeyTCPConnectTime:          "tcp",
			KeyServerResponseTime:      "srt",
			KeyExceptionDescription:    "exd",
			KeyExceptionFatal:          "exf",
			KeyExperimentID:            "xid",
			KeyExperimentVariant:       "xvar",
		}
	})
	p, e := paramKeys[param]
	return p, e
}
