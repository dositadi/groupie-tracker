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

func (a *ArtistInfo) SetIsFavorited(id int, status bool) {
	if val, ok := byId[id]; ok {
		if status {
			val.IsFavorited = true
		}
		if !status {
			val.IsFavorited = false
		}
	}
	if val, ok := byCreationDate[id]; ok {
		if status {
			val.IsFavorited = true
		}
		if !status {
			val.IsFavorited = false
		}
	}
	for _, artist := range byFirstAlbum {
		if artist.Id == id {
			if status {
				artist.IsFavorited = true
			}
			if !status {
				artist.IsFavorited = false
			}
		}
	}
	for _, artist := range byName {
		if artist.Id == id {
			if status {
				artist.IsFavorited = true
			}
			if !status {
				artist.IsFavorited = false
			}
		}
	}
}
