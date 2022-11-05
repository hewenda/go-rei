package config

type ModelData struct {
	TitlePattern            string      `json:"titlePattern"`
	Icon                    string      `json:"icon"`
	Description             string      `json:"description"`
	IncludeJSONModel        bool        `json:"includeJsonModel"`
	HideUniversalNavigation bool        `json:"hideUniversalNavigation"`
	HideFooter              bool        `json:"hideFooter"`
	Title                   string      `json:"title"`
	Analytics               interface{} `json:"analytics"`
	AllowZoom               bool        `json:"allowZoom"`
	StaticRoot              string      `json:"staticRoot"`
	HideHeaderFooter        bool        `json:"hideHeaderFooter"`
	HideHeader              bool        `json:"hideHeader"`
	UseAdobeVisitorService  bool        `json:"useAdobeVisitorService"`
	HideGlobalNavigation    bool        `json:"hideGlobalNavigation"`
	Favicon                 string      `json:"favicon"`
	MaskIcon                string      `json:"maskIcon"`
	Canonical               string      `json:"canonical"`
	PageData                PageData    `json:"pageData"`
	CustomStyles            interface{} `json:"customStyles"`
	CustomScripts           interface{} `json:"customScripts"`
	UseAdobeTarget          bool        `json:"useAdobeTarget"`
	CanonicalRoot           string      `json:"canonicalRoot"`
	HidePromoMessageBanner  bool        `json:"hidePromoMessageBanner"`
	Meta                    interface{} `json:"meta"`
	SchemaType              interface{} `json:"schemaType"`
	SiteID                  string      `json:"siteId"`
	IncludeAnalyticsJSON    bool        `json:"includeAnalyticsJson"`
	ContactPhoneNumber      string      `json:"contactPhoneNumber"`
}

type PageData struct {
	Product     Product     `json:"product"`
	Messages    interface{} `json:"messages"`
	Analytics   interface{} `json:"analytics"`
	AppProps    interface{} `json:"appProps"`
	PreviewMode bool        `json:"previewMode"`
	Preview     bool        `json:"preview"`
	FeDev       bool        `json:"feDev"`
}

type Colors struct {
	Name         string   `json:"name"`
	Code         string   `json:"code"`
	ColorFamily  string   `json:"colorFamily"`
	DisplayLabel string   `json:"displayLabel"`
	HexValue     []string `json:"hexValue"`
}
type Color struct {
	Name         string   `json:"name"`
	Code         string   `json:"code"`
	ColorFamily  string   `json:"colorFamily"`
	DisplayLabel string   `json:"displayLabel"`
	HexValue     []string `json:"hexValue"`
}

type Size struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	SortOrder int    `json:"sortOrder"`
}

type Price struct {
	Value     float64 `json:"value"`
	CompValue float64 `json:"compValue"`
	OfferType string  `json:"offerType"`
	Sale      bool    `json:"sale"`
	SaleOnly  bool    `json:"saleOnly"`
	Clearance bool    `json:"clearance"`
}

type Skus struct {
	SkuID                   string          `json:"skuId"`
	Status                  string          `json:"status"`
	Color                   Color           `json:"color"`
	Size                    Size            `json:"size"`
	Price                   AvailablePrices `json:"price"`
	ShippingTier            int             `json:"shippingTier"`
	Oversize                bool            `json:"oversize"`
	OversizeCharge          interface{}     `json:"oversizeCharge"`
	Hazardous               bool            `json:"hazardous"`
	AssemblyRequired        bool            `json:"assemblyRequired"`
	Campaigns               []interface{}   `json:"campaigns"`
	SourceType              string          `json:"sourceType"`
	Upc                     string          `json:"upc"`
	MembersOnly             bool            `json:"membersOnly"`
	ShippingRestrictions    interface{}     `json:"shippingRestrictions"`
	NewArrival              bool            `json:"newArrival"`
	AnyShippingRestrictions bool            `json:"anyShippingRestrictions"`
	Sellable                bool            `json:"sellable"`
}

type Images struct {
	ID                 string      `json:"id"`
	Title              interface{} `json:"title"`
	Type               string      `json:"type"`
	Description        interface{} `json:"description"`
	Tag                interface{} `json:"tag"`
	AltText            interface{} `json:"altText"`
	Caption            interface{} `json:"caption"`
	Size               string      `json:"size"`
	MaxHeight          int         `json:"maxHeight"`
	MaxWidth           int         `json:"maxWidth"`
	URI                string      `json:"uri"`
	HeroImageURL       string      `json:"heroImageUrl"`
	ThumbnailURL       string      `json:"thumbnailUrl"`
	Color              Color       `json:"color"`
	Featured           bool        `json:"featured"`
	FeaturedFromSource bool        `json:"featuredFromSource"`
}
type Videos struct {
	ID           interface{} `json:"id"`
	Title        string      `json:"title"`
	Type         interface{} `json:"type"`
	Description  string      `json:"description"`
	Tag          interface{} `json:"tag"`
	AltText      interface{} `json:"altText"`
	Caption      interface{} `json:"caption"`
	EmbedCode    string      `json:"embedCode"`
	ThumbnailURL string      `json:"thumbnailUrl"`
	HeroImageURL string      `json:"heroImageUrl"`
	ArticleURL   string      `json:"articleUrl"`
	SourceType   string      `json:"sourceType"`
	Status       string      `json:"status"`
}

type CompareAt struct {
	Value     float64 `json:"value"`
	CompValue float64 `json:"compValue"`
	OfferType string  `json:"offerType"`
	Sale      bool    `json:"sale"`
	SaleOnly  bool    `json:"saleOnly"`
	Clearance bool    `json:"clearance"`
}

type StylePrice struct {
	MinPrice               Price       `json:"minPrice"`
	MaxPrice               Price       `json:"maxPrice"`
	MinCompareAt           CompareAt   `json:"minCompareAt"`
	MaxCompareAt           CompareAt   `json:"maxCompareAt"`
	MinSavingsPercentage   interface{} `json:"minSavingsPercentage"`
	MaxSavingsPercentage   interface{} `json:"maxSavingsPercentage"`
	MinSavingsDollarAmount int         `json:"minSavingsDollarAmount"`
	MaxSavingsDollarAmount int         `json:"maxSavingsDollarAmount"`
	OfferType              string      `json:"offerType"`
	Map                    bool        `json:"map"`
}

type AvailablePrices struct {
	CompareAt           CompareAt   `json:"compareAt"`
	Price               Price       `json:"price"`
	SavingsPercentage   interface{} `json:"savingsPercentage"`
	SavingsDollarAmount int         `json:"savingsDollarAmount"`
	HideDisplayPrice    bool        `json:"hideDisplayPrice"`
	IsSale              bool        `json:"isSale"`
	Map                 bool        `json:"map"`
}

type BySku struct {
	String BySkuContent
}
type BySkuContent struct {
	SkuID                   string        `json:"skuId"`
	Status                  string        `json:"status"`
	Color                   Color         `json:"color"`
	Size                    Size          `json:"size"`
	Price                   Price         `json:"price"`
	ShippingTier            int           `json:"shippingTier"`
	Oversize                bool          `json:"oversize"`
	OversizeCharge          interface{}   `json:"oversizeCharge"`
	Hazardous               bool          `json:"hazardous"`
	AssemblyRequired        bool          `json:"assemblyRequired"`
	Campaigns               []interface{} `json:"campaigns"`
	SourceType              string        `json:"sourceType"`
	Upc                     string        `json:"upc"`
	MembersOnly             bool          `json:"membersOnly"`
	ShippingRestrictions    interface{}   `json:"shippingRestrictions"`
	NewArrival              bool          `json:"newArrival"`
	AnyShippingRestrictions bool          `json:"anyShippingRestrictions"`
	Sellable                bool          `json:"sellable"`
}

type Product struct {
	Title                          string            `json:"title"`
	StyleID                        string            `json:"styleId"`
	CanonicalURL                   string            `json:"canonicalUrl"`
	DisplayOptions                 interface{}       `json:"displayOptions"`
	Colors                         []Colors          `json:"colors"`
	Sizes                          []string          `json:"sizes"`
	Skus                           []Skus            `json:"skus"`
	Images                         []Images          `json:"images"`
	UserManuals                    []interface{}     `json:"userManuals"`
	Videos                         []Videos          `json:"videos"`
	SizeChart                      interface{}       `json:"sizeChart"`
	StylePrice                     StylePrice        `json:"stylePrice"`
	AvailablePrices                []AvailablePrices `json:"availablePrices"`
	OptionsAvailableWithoutPrice   interface{}       `json:"optionsAvailableWithoutPrice"`
	OptionsAvailableByPrice        interface{}       `json:"optionsAvailableByPrice"`
	MixedPriceOptions              interface{}       `json:"mixedPriceOptions"`
	BySize                         interface{}       `json:"bySize"`
	SkusForOption                  interface{}       `json:"skusForOption"`
	ByColor                        interface{}       `json:"byColor"`
	ShowSwatches                   bool              `json:"showSwatches"`
	TechSpecs                      interface{}       `json:"techSpecs"`
	FeaturedSpecs                  interface{}       `json:"featuredSpecs"`
	Features                       []string          `json:"features"`
	Description                    string            `json:"description"`
	OriginCountry                  string            `json:"originCountry"`
	PrimaryCategory                interface{}       `json:"primaryCategory"`
	BySku                          BySku             `json:"bySku"`
	ReviewSummary                  interface{}       `json:"reviewSummary"`
	SapGender                      string            `json:"sapGender"`
	AllDisplayableSkusArePreorder  bool              `json:"allDisplayableSkusArePreorder"`
	AllDisplayableSkusAreBackorder bool              `json:"allDisplayableSkusAreBackorder"`
	Outlet                         bool              `json:"outlet"`
	EligibleForShipping            bool              `json:"eligibleForShipping"`
	AnyOversizeCharges             bool              `json:"anyOversizeCharges"`
	OversizeCharges                []interface{}     `json:"oversizeCharges"`
	HideOversizeMessaging          bool              `json:"hideOversizeMessaging"`
	RelatedProducts                interface{}       `json:"relatedProducts"`
	RelatedLinks                   interface{}       `json:"relatedLinks"`
	ProductLineURL                 string            `json:"productLineUrl"`
	ProductLineURLLabel            string            `json:"productLineUrlLabel"`
	AnySkuShippingRestrictions     bool              `json:"anySkuShippingRestrictions"`
	SellingCopy                    interface{}       `json:"sellingCopy"`
	Badges                         interface{}       `json:"badges"`
	MembershipV2                   interface{}       `json:"membershipV2"`
	AllSkusAreMembersOnly          bool              `json:"allSkusAreMembersOnly"`
	AnySkusAreMembersOnly          bool              `json:"anySkusAreMembersOnly"`
	Brand                          interface{}       `json:"brand"`
	PublicationStatus              string            `json:"publicationStatus"`
}
