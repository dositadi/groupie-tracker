package artistapi

var byIdKey, byCreationDateKey = make(map[int]ArtistInfo), make(map[int]ArtistInfo)
var byName, byFirstAlbum = make(map[string]ArtistInfo), make(map[string]ArtistInfo)

func New() *ArtistInfo {
	return &ArtistInfo{}
}

func (a *ArtistInfo) Init() {
	byIdKey, byCreationDateKey, byName, byFirstAlbum = a.mapArtistsInfo()
	logger.PrintInfo("Artists Info initialized successfully", map[string]string{
		"Source": "Init f(n) in artistapi pkg",
	})
}

func (a *ArtistInfo) GetByIdKey() map[int]ArtistInfo {
	return byIdKey
}

func (a *ArtistInfo) GetByCreationDate() map[int]ArtistInfo {
	return byCreationDateKey
}

func (a *ArtistInfo) GetByName() map[string]ArtistInfo {
	return byName
}

func (a *ArtistInfo) GetByFirstAlbum() map[string]ArtistInfo {
	return byFirstAlbum
}

func (a *ArtistInfo) SetIsFavorited(id int, status bool) {
	if val, ok := byIdKey[id]; ok {
		if val.IsFavorited && status {
			val.IsFavorited = !val.IsFavorited
		}
		if !val.IsFavorited && !status {
			val.IsFavorited = !val.IsFavorited
		}
	}
	if val, ok := byCreationDateKey[id]; ok {
		if val.IsFavorited && status {
			val.IsFavorited = !val.IsFavorited
		}
		if !val.IsFavorited && !status {
			val.IsFavorited = !val.IsFavorited
		}
	}
	for _, artist := range byFirstAlbum {
		if artist.Id == id {
			if artist.IsFavorited && status {
				artist.IsFavorited = !artist.IsFavorited
			}
			if !artist.IsFavorited && !status {
				artist.IsFavorited = !artist.IsFavorited
			}
		}
	}
	for _, artist := range byName {
		if artist.Id == id {
			if artist.IsFavorited && status {
				artist.IsFavorited = !artist.IsFavorited
			}
			if !artist.IsFavorited && !status {
				artist.IsFavorited = !artist.IsFavorited
			}
		}
	}
}
