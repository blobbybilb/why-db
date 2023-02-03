package types

type Store struct {
	Set func(cat string, key string, data string) error
	Get func(cat string, key string) (string, error)
	Add func(cat string, key string, data string) error
	Del func(cat string, key string) error
}
