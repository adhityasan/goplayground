package index

import "errors"

// SetValue set value of an index of integer array through its pointer.
//
// `1st` param is the array.
// `2nd` param is the index.
// `3rd` param is the new value.
func SetValue(array []int, indexOf int, newVal int) error {

	if indexOf >= len(array) {
		return errors.New("Index not found")
	}
	varpoint := &array[indexOf]
	*varpoint = newVal

	return nil
}

// SetAllValue set all index value of an array through it's pointer
//
// `1st` param is the target array.
// `2nd` param is the array to update the target array value.
//
// It return `error` if the both array length is not the same
func SetAllValue(oldarr []int, newarr []int) error {
	if len(oldarr) != len(newarr) {
		return errors.New("Error, param length is not the same")
	}

	for index := 0; index < len(oldarr); index++ {
		varpoint := &oldarr[index]
		*varpoint = newarr[index]
	}

	return nil
}
