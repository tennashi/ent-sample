// Code generated by entc, DO NOT EDIT.

package namespace

const (
	// Label holds the string label denoting the namespace type in the database.
	Label = "namespace"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgeMembers holds the string denoting the members edge name in mutations.
	EdgeMembers = "members"
	// EdgeAdmins holds the string denoting the admins edge name in mutations.
	EdgeAdmins = "admins"

	// Table holds the table name of the namespace in the database.
	Table = "namespaces"
	// MembersTable is the table the holds the members relation/edge. The primary key declared below.
	MembersTable = "namespace_members"
	// MembersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	MembersInverseTable = "users"
	// AdminsTable is the table the holds the admins relation/edge. The primary key declared below.
	AdminsTable = "namespace_admins"
	// AdminsInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AdminsInverseTable = "users"
)

// Columns holds all SQL columns for namespace fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// MembersPrimaryKey and MembersColumn2 are the table columns denoting the
	// primary key for the members relation (M2M).
	MembersPrimaryKey = []string{"namespace_id", "user_id"}
	// AdminsPrimaryKey and AdminsColumn2 are the table columns denoting the
	// primary key for the admins relation (M2M).
	AdminsPrimaryKey = []string{"namespace_id", "user_id"}
)