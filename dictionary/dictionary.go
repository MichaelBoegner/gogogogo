package dictionary

type Dictionary map[string]string

func Search(dictionary map[string]string, key string) (value string) {

	return dictionary[key]
}
