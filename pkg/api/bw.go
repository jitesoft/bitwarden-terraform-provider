package api

import (
	"github.com/jitesoft/bitwarden-terraform-provider/pkg/types"
)

// GetEntry fetches an entry from the BitWarden vault.
func GetEntry() (*types.Entry, error) {
	return nil, nil
}

// EntryExist does a lookup in the vault to see if the entry exist.
func EntryExist() (bool, error) {
	return false, nil
}

// UpdateEntry updates a given entry with new value.
func UpdateEntry() error {
	return nil
}

// DeleteEntry will remove an entry from the vault.
func DeleteEntry() error {
	return nil
}

// CreateEntry creates a entry with given name, path and value in the vault.
func CreateEntry() error {
	return nil
}
