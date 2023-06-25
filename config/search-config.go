package config

import "gin-demo/search"

var SEARCH *search.MeiliSearch

func getMeiliSearch() *search.MeiliSearch {
	return search.NewMeiliSearch(CONFIG.MeiliSearchConfig.ApiHost, CONFIG.MeiliSearchConfig.ApiKey)
}

func LoadMeiliSearchConfig() {
	SEARCH = getMeiliSearch()
}
