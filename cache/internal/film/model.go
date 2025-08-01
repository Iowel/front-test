package film

type FilmResponse struct {
	KinopoiskID                int      `json:"kinopoiskId"`
	KinopoiskHDID              string   `json:"kinopoiskHDId"`
	ImdbID                     string   `json:"imdbId"`
	NameRu                     string   `json:"nameRu"`
	NameEn                     *string  `json:"nameEn"`
	NameOriginal               *string  `json:"nameOriginal"`
	PosterURL                  string   `json:"posterUrl"`
	PosterURLPreview           string   `json:"posterUrlPreview"`
	CoverURL                   string   `json:"coverUrl"`
	LogoURL                    *string  `json:"logoUrl"`
	ReviewsCount               int      `json:"reviewsCount"`
	RatingGoodReview           float64  `json:"ratingGoodReview"`
	RatingGoodReviewVoteCount  int      `json:"ratingGoodReviewVoteCount"`
	RatingKinopoisk            float64  `json:"ratingKinopoisk"`
	RatingKinopoiskVoteCount   int      `json:"ratingKinopoiskVoteCount"`
	RatingImdb                 float64  `json:"ratingImdb"`
	RatingImdbVoteCount        int      `json:"ratingImdbVoteCount"`
	RatingFilmCritics          *float64 `json:"ratingFilmCritics"`
	RatingFilmCriticsVoteCount int      `json:"ratingFilmCriticsVoteCount"`
	RatingAwait                *float64 `json:"ratingAwait"`
	RatingAwaitCount           int      `json:"ratingAwaitCount"`
	RatingRfCritics            float64  `json:"ratingRfCritics"`
	RatingRfCriticsVoteCount   int      `json:"ratingRfCriticsVoteCount"`
	WebURL                     string   `json:"webUrl"`
	Year                       int      `json:"year"`
	FilmLength                 *int     `json:"filmLength"`
	Slogan                     *string  `json:"slogan"`
	Description                string   `json:"description"`
	ShortDescription           string   `json:"shortDescription"`
	EditorAnnotation           *string  `json:"editorAnnotation"`
	IsTicketsAvailable         bool     `json:"isTicketsAvailable"`
	ProductionStatus           *string  `json:"productionStatus"`
	Type                       string   `json:"type"`
	RatingMpaa                 *string  `json:"ratingMpaa"`
	RatingAgeLimits            string   `json:"ratingAgeLimits"`
	Countries                  []struct {
		Country string `json:"country"`
	} `json:"countries"`
	Genres []struct {
		Genre string `json:"genre"`
	} `json:"genres"`
	StartYear int    `json:"startYear"`
	EndYear   *int   `json:"endYear"`
	Serial    bool   `json:"serial"`
	ShortFilm bool   `json:"shortFilm"`
	Completed bool   `json:"completed"`
	HasImax   bool   `json:"hasImax"`
	Has3D     bool   `json:"has3D"`
	LastSync  string `json:"lastSync"`
}

type Actor struct {
	StaffID        int64  `json:"staffId"`
	NameRu         string `json:"nameRu"`
	NameEn         string `json:"nameEn"`
	Description    string `json:"description"`
	PosterURL      string `json:"posterUrl"`
	ProfessionKey  string `json:"professionKey"`
	ProfessionText string `json:"professionText"`
}

type FilmListResponse struct {
	PagesCount int    `json:"pagesCount"`
	Films      []Film `json:"films"`
}

type Film struct {
	FilmID           int       `json:"filmId"`
	NameRu           string    `json:"nameRu"`
	NameEn           *string   `json:"nameEn"`
	Year             string    `json:"year"`
	FilmLength       string    `json:"filmLength"`
	Countries        []Country `json:"countries"`
	Genres           []Genre   `json:"genres"`
	Rating           *string   `json:"rating"`
	RatingVoteCount  int       `json:"ratingVoteCount"`
	PosterURL        string    `json:"posterUrl"`
	PosterURLPreview string    `json:"posterUrlPreview"`
	RatingChange     *string   `json:"ratingChange"`
	IsRatingUp       *bool     `json:"isRatingUp"`
	IsAfisha         int       `json:"isAfisha"`
}

type Country struct {
	Country string `json:"country"`
}

type Genre struct {
	Genre string `json:"genre"`
}

type PremieresResponse struct {
	Total int        `json:"total"`
	Docs  []*Popular `json:"docs"`
	Pages int        `json:"pages"` // добавляем поле с количеством страниц
}

type CountryPopular struct {
	Name string `json:"name"`
}

type GenrePopular struct {
	Name string `json:"name"`
}

type Popular struct {
	ID              int              `json:"id"`
	Name            string           `json:"name"`
	AlternativeName string           `json:"alternativeName"`
	Year            *int             `json:"year"`
	Description     string           `json:"description"`
	MovieLength     *int             `json:"movieLength"`
	AgeRating       *int             `json:"ageRating"`
	Rating          Rating           `json:"rating"`
	Votes           Votes            `json:"votes"`
	Poster          *Poster          `json:"poster"`
	Genres          []GenrePopular   `json:"genres"`
	Countries       []CountryPopular `json:"countries"`
	IsSeries        bool             `json:"isSeries"`
	TicketsOnSale   bool             `json:"ticketsOnSale"`
}

type Rating struct {
	Kp                 float64 `json:"kp"`
	Imdb               float64 `json:"imdb"`
	FilmCritics        float64 `json:"filmCritics"`
	RussianFilmCritics float64 `json:"russianFilmCritics"`
	Await              float64 `json:"await"`
}

type Votes struct {
	Kp                 int `json:"kp"`
	Imdb               int `json:"imdb"`
	FilmCritics        int `json:"filmCritics"`
	RussianFilmCritics int `json:"russianFilmCritics"`
	Await              int `json:"await"`
}

type Poster struct {
	Url        string `json:"url"`
	PreviewUrl string `json:"previewUrl"`
}
