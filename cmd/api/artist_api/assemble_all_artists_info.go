package artistapi

func (a *ArtistInfo) assembleArtistInfoAsMap(chArtistsInfo chan *ArtistInfo) (map[int]ArtistInfo, map[int]ArtistInfo, map[string]ArtistInfo, map[string]ArtistInfo) {
	byId := make(map[int]ArtistInfo)
	byCreationDate := make(map[int]ArtistInfo)
	byName := make(map[string]ArtistInfo)
	byFirstAlbum := make(map[string]ArtistInfo)

	for artistInfo := range chArtistsInfo {
		byId[artistInfo.Id] = *artistInfo
		byCreationDate[artistInfo.CreationDate] = *artistInfo
		byName[artistInfo.Name] = *artistInfo
		byFirstAlbum[artistInfo.FirstAlbum] = *artistInfo
	}
	return byId, byCreationDate, byName, byFirstAlbum
}
