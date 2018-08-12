package data

type IMap interface {
	Get(key string) interface{}
	Put(key string, value interface{})
	Remove(key string)
	ContainsKey(key string) bool
}
