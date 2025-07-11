package fakeUserAgent

import "bytes"

// Browsers
const (
	Google           = "Google"
	Chrome           = "Chrome"
	Firefox          = "Firefox"
	Edge             = "Edge"
	Opera            = "Opera"
	Safari           = "Safari"
	AndroidBrowser   = "Android"
	YandexBrowser    = "Yandex Browser"
	SamsungInternet  = "Samsung Internet"
	OperaMobile      = "Opera Mobile"
	MobileSafari     = "Mobile Safari"
	FirefoxMobile    = "Firefox Mobile"
	FirefoxiOS       = "Firefox iOS"
	ChromeMobile     = "Chrome Mobile"
	ChromeMobileiOS  = "Chrome Mobile iOS"
	MobileSafariUIWK = "Mobile Safari UI/WKWebView"
	EdgeMobile       = "Edge Mobile"
	DuckDuckGoMobile = "DuckDuckGo Mobile"
	MiuiBrowser      = "MiuiBrowser"
	Whale            = "Whale"
	Twitter          = "Twitter"
	Facebook         = "Facebook"
	AmazonSilk       = "Amazon Silk"
)

// OS
const (
	Windows  = "Windows"
	Linux    = "Linux"
	Ubuntu   = "Ubuntu"
	ChromeOS = "Chrome OS"
	MacOSX   = "Mac OS X"
	Android  = "Android"
	IOS      = "iOS"
)

// Platforms
const (
	Mobile  = "mobile"
	Tablet  = "tablet"
	Desktop = "desktop"
)

type UserAgent struct {
	List     []UserAgents
	Filtered []UserAgents
	fallback *string
}

type FilterBy struct {
	UserAgent *UserAgent
	fallback  *string
}

type UserAgents struct {
	Useragent                string  `json:"useragent"`
	Percent                  float64 `json:"percent"`
	Type                     string  `json:"type"`
	DeviceBrand              string  `json:"device_brand"`
	Browser                  string  `json:"browser"`
	BrowserVersion           string  `json:"browser_version"`
	BrowserVersionMajorMinor float64 `json:"browser_version_major_minor"`
	Os                       string  `json:"os"`
	OsVersion                string  `json:"os_version"`
	Platform                 string  `json:"platform"`
}

// Init UserAgent client and load user-Agents from JSON file to memory cache
func New() (*UserAgent, error) {
	var jsonreader = bytes.NewReader(userAgentsFile)

	userAgents, err := getUserAgents(jsonreader)
	if err != nil {
		return nil, err
	}
	fallbackDefault := ""
	return &UserAgent{List: userAgents, fallback: &fallbackDefault}, nil
}

// Gets random user-Agent without filters
// This method don't init filter method
func (c *UserAgent) GetRandom() string {
	filteredList := Filter(c.List, func(d UserAgents) bool {
		return true
	})
	randomIndex := randFromLen(len(filteredList))
	return filteredList[randomIndex].Useragent
}

// Init filters for user-Agent
func (c *UserAgent) Filter() *FilterBy {
	return &FilterBy{UserAgent: c, fallback: c.fallback}
}

// Set fallback user_agent
func (c *UserAgent) SetFallback(fallback string) *UserAgent {
	c.fallback = &fallback
	return c
}

// Gets from filters
// Method return userAgent string only or empty string if nothing found
func (c *FilterBy) Get() string {
	filteredList := c.getList()
	if len(filteredList) == 0 {
		return *c.fallback
	}
	randomIndex := randFromLen(len(filteredList))
	return filteredList[randomIndex].Useragent
}

// Gets from filters
// Method return all struct UserAgents or empty UserAgents struct if nothing found
func (c *FilterBy) GetRaw() UserAgents {
	filteredList := c.getList()
	if len(filteredList) == 0 {
		return UserAgents{}
	}
	randomIndex := randFromLen(len(filteredList))
	return filteredList[randomIndex]
}

// Get filtered user-Agent list
// Get from memory new filtered list
func (c *FilterBy) getList() []UserAgents {
	if c.UserAgent.Filtered != nil {
		return c.UserAgent.Filtered
	}
	return c.UserAgent.List
}

// Filter by: Os
func (c *FilterBy) Os(f ...string) *FilterBy {
	filtered := Filter(c.getList(), func(d UserAgents) bool {
		return stringInSlice(d.Os, f)
	})
	c.UserAgent.Filtered = filtered
	return c
}

// Filter by: Browser
func (c *FilterBy) Browser(f ...string) *FilterBy {
	filtered := Filter(c.getList(), func(d UserAgents) bool {
		return stringInSlice(d.Browser, f)
	})
	c.UserAgent.Filtered = filtered
	return c
}

// Filter by: Platform
func (c *FilterBy) Platform(f ...string) *FilterBy {
	filtered := Filter(c.getList(), func(d UserAgents) bool {
		return stringInSlice(d.Type, f)
	})
	c.UserAgent.Filtered = filtered
	return c
}

// Filter by: Os Version
func (c *FilterBy) OsVer(f int) *FilterBy {
	filtered := Filter(c.getList(), func(d UserAgents) bool {
		return ExtractMajorVersion(d.OsVersion) == f
	})
	c.UserAgent.Filtered = filtered
	return c
}

// Filter by: Chrome browserss
func (c *FilterBy) Chrome() *FilterBy {
	return c.Browser(Chrome, ChromeMobile, ChromeMobileiOS)
}

// Filter by: Firefox browser
func (c *FilterBy) Firefox() *FilterBy {
	return c.Browser(Firefox, FirefoxMobile, FirefoxiOS)
}

// Filter by: Safari browser
func (c *FilterBy) Safari() *FilterBy {
	return c.Browser(Safari, MobileSafari, MobileSafariUIWK)
}

// Filter by: Opera browser
func (c *FilterBy) Opera() *FilterBy {
	return c.Browser(Opera, OperaMobile)
}

// Filter by: Edge browser
func (c *FilterBy) Edge() *FilterBy {
	return c.Browser(Edge, EdgeMobile)
}