package singlefile

import (
	"testing"

	"github.com/LintaoAmons/undercontrol/src/domain/account"
	singlefile "github.com/LintaoAmons/undercontrol/src/persistence/common/single-file"
	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {
	// Create a mock SingleFileDB
	mockDB := &MockSingleFileDB{
		tableData: map[singlefile.Table][]interface{}{
			ACCOUNT: {
				&account.Account{Name: "Account 1"},
				&account.Account{Name: "Account 2"},
				&account.Account{Name: "Account 3"},
			},
		},
	}

	// Create a new instance of the SingleFileAccountRepo with the mockDB
	repo := NewSingleFileAccountRepo(mockDB)

	// Retrieve all accounts from the repository
	accounts := repo.FindAll()

	// Assert that the number of retrieved accounts is equal to the number of accounts in the mockDB
	assert.Equal(t, 3, len(accounts))

	// Assert that the retrieved accounts match the accounts in the mockDB
	assert.Equal(t, "Account 1", accounts[0].Name)
	assert.Equal(t, "Account 2", accounts[1].Name)
	assert.Equal(t, "Account 3", accounts[2].Name)
}

// MockSingleFileDB is a mock implementation of the SingleFileDB interface
type MockSingleFileDB struct {
	tableData map[singlefile.Table][]interface{}
}

func (m *MockSingleFileDB) GetTable(table singlefile.Table) []interface{} {
	return m.tableData[table]
}

func (m *MockSingleFileDB) SaveTable(table singlefile.Table, data []interface{}) {
	m.tableData[table] = data
}
