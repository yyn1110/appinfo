package controllers

import (

	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Result struct {
	Kind                               string   `json:kind`
	Features                           []string `json:features`
	SupportedDevices                   []string `json:supportedDevices`
	IsGameCenterEnabled                bool     `json:isGameCenterEnabled`
	ScreenshotUrls                     []string `json:screenshotUrls`
	IpadScreenshotUrls                 []string `json:ipadScreenshotUrls`
	ArtworkUrl60                       string   `json:artworkUrl60`
	ArtworkUrl512                      string   `json:artworkUrl512`
	ArtistViewUrl                      string   `json:artistViewUrl`
	ArtistId                           float32  `json:artistId`
	ArtistName                         string   `json:artistName`
	Price                              float32  `json:price`
	Version                            string   `json:version`
	Description                        string   `json:description`
	Currency                           string   `json:currency`
	Genres                             []string `json:genres`
	GenreIds                           []string `json:genreIds`
	ReleaseDate                        string   `json:releaseDate`
	SellerName                         string   `json:sellerName`
	BundleId                           string   `json:bundleId`
	TrackId                            float32  `json:trackId`
	TrackName                          string   `json:trackName`
	PrimaryGenreName                   string   `json:primaryGenreName`
	PrimaryGenreId                     float32  `json:primaryGenreId`
	ReleaseNotes                       string   `json:releaseNotes`
	FormattedPrice                     string   `json:formattedPrice`
	WrapperType                        string   `json:wrapperType`
	TrackCensoredName                  string   `json:trackCensoredName`
	LanguageCodesISO2A                 []string `json:languageCodesISO2A`
	FileSizeBytes                      string   `json:fileSizeBytes`
	SellerUrl                          string   `json:sellerUrl`
	ContentAdvisoryRating              string   `json:contentAdvisoryRating`
	AverageUserRatingForCurrentVersion float32  `json:averageUserRatingForCurrentVersion`
	UserRatingCountForCurrentVersion   float32  `json:userRatingCountForCurrentVersion`
	ArtworkUrl100                      string   `json:artworkUrl100`
	TrackViewUrl                       string   `json:trackViewUrl`
	TrackContentRating                 string   `json:trackContentRating`
	AverageUserRating                  float32  `json:averageUserRating`
	UserRatingCount                    float32  `json:userRatingCount`
}
type ResultsInfo struct {
	ResultCount int       `json:resultCount`
	Results     []*Result `json:results`
}
type SearchController struct {
	beego.Controller
}

func (this *SearchController) Get() {
	id := this.Ctx.Params[":appid"]
	str := this.Ctx.Params[":lg"]
	if str == "" {
		this.Ctx.WriteString(RequestLgErr)
		return
	}
	_, err := strconv.Atoi(id)
	if err != nil {
		this.Ctx.WriteString(AppIDErr)
		return
	}

	var requestURL string
	if str == "en" {
		requestURL = AppleENLookURL + string(id)
	} else if str == "ch" {
		requestURL = AppleCNLookURL + string(id)
	}

	resp, err := http.Get(requestURL)
	defer resp.Body.Close()
	if err != nil {

		this.Ctx.WriteString(RequestErr)
		return
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			this.Ctx.WriteString(DecodeBodyErr)
			return
		}
		r := new(ResultsInfo)
		err = json.Unmarshal(body, r)
		if err != nil {
			this.Ctx.WriteString(DecodeErr)
			return
		}
		rest, err := json.MarshalIndent(r, "  ", "")
		if err != nil {
			this.Ctx.WriteString(DecodeErr)
			return
		}
		this.Ctx.WriteString(string(rest))
	}

}
