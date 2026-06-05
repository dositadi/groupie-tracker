package artistapi

func New() *ArtistInfo {
	return &ArtistInfo{}
}

func (a *ArtistInfo) Init() {
	a.mapArtistsInfo()
}

func (a *ArtistInfo) GetByIdKey() map[int]ArtistInfo {
	return byId
}

func (a *ArtistInfo) GetByCreationDate() map[int]ArtistInfo {
	return byCreationDate
}

func (a *ArtistInfo) GetByName() map[string]ArtistInfo {
	return byName
}

func (a *ArtistInfo) GetByFirstAlbum() map[string]ArtistInfo {
	return byFirstAlbum
}