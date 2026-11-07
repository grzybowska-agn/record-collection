package main

type ReleaseType string

const (
	ReleaseTypeLP     ReleaseType = "LP"
	ReleaseTypeEP     ReleaseType = "EP"
	ReleaseTypeSingle ReleaseType = "Single"
)

type VinylSize int

const (
	Size7Inch  VinylSize = 7
	Size10INch VinylSize = 10
	Size12Inch VinylSize = 12
)

type Record struct {
	ID              int
	Artist          string
	Title           string
	OriginalYear    int
	CoverURL        string
	ReleaseType     ReleaseType
	SizeInches      VinylSize
	DiscCount       int
	PressingYear    int
	PressingLabel   string
	PressingCountry string
	IsReissue       bool
}
