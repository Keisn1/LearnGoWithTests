package main

type Dictionary map[string]string

const (
	ErrNotFound               = DictionaryErr("Could not find the word")
	ErrWordPresent            = DictionaryErr("Word is already present, rather use Dictionary.Update")
	ErrWordDoesNotExist       = DictionaryErr("Word does not exist")
	ErrWordToDeleteNotPresent = DictionaryErr("Word to be deleted doesn't exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	def, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return def, nil
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = def
	case nil:
		return ErrWordPresent
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = def
		return nil
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		delete(d, word)
		return nil
	case ErrNotFound:
		return ErrWordToDeleteNotPresent
	default:
		return err
	}
}
