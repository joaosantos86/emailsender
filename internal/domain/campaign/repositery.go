package campaign

type Repositery interface {
	Save(campaign *Campaign) error
}