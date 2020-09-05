// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"

	"github.com/tennashi/ent-sample/ent/namespace"
	"github.com/tennashi/ent-sample/ent/user"

	"github.com/facebook/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeNamespace = "Namespace"
	TypeUser      = "User"
)

// NamespaceMutation represents an operation that mutate the Namespaces
// nodes in the graph.
type NamespaceMutation struct {
	config
	op             Op
	typ            string
	id             *string
	name           *string
	clearedFields  map[string]struct{}
	members        map[string]struct{}
	removedmembers map[string]struct{}
	admins         map[string]struct{}
	removedadmins  map[string]struct{}
	done           bool
	oldValue       func(context.Context) (*Namespace, error)
}

var _ ent.Mutation = (*NamespaceMutation)(nil)

// namespaceOption allows to manage the mutation configuration using functional options.
type namespaceOption func(*NamespaceMutation)

// newNamespaceMutation creates new mutation for $n.Name.
func newNamespaceMutation(c config, op Op, opts ...namespaceOption) *NamespaceMutation {
	m := &NamespaceMutation{
		config:        c,
		op:            op,
		typ:           TypeNamespace,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withNamespaceID sets the id field of the mutation.
func withNamespaceID(id string) namespaceOption {
	return func(m *NamespaceMutation) {
		var (
			err   error
			once  sync.Once
			value *Namespace
		)
		m.oldValue = func(ctx context.Context) (*Namespace, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Namespace.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withNamespace sets the old Namespace of the mutation.
func withNamespace(node *Namespace) namespaceOption {
	return func(m *NamespaceMutation) {
		m.oldValue = func(context.Context) (*Namespace, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m NamespaceMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m NamespaceMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on Namespace creation.
func (m *NamespaceMutation) SetID(id string) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *NamespaceMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the name field.
func (m *NamespaceMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *NamespaceMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old name value of the Namespace.
// If the Namespace object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *NamespaceMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName reset all changes of the "name" field.
func (m *NamespaceMutation) ResetName() {
	m.name = nil
}

// AddMemberIDs adds the members edge to User by ids.
func (m *NamespaceMutation) AddMemberIDs(ids ...string) {
	if m.members == nil {
		m.members = make(map[string]struct{})
	}
	for i := range ids {
		m.members[ids[i]] = struct{}{}
	}
}

// RemoveMemberIDs removes the members edge to User by ids.
func (m *NamespaceMutation) RemoveMemberIDs(ids ...string) {
	if m.removedmembers == nil {
		m.removedmembers = make(map[string]struct{})
	}
	for i := range ids {
		m.removedmembers[ids[i]] = struct{}{}
	}
}

// RemovedMembers returns the removed ids of members.
func (m *NamespaceMutation) RemovedMembersIDs() (ids []string) {
	for id := range m.removedmembers {
		ids = append(ids, id)
	}
	return
}

// MembersIDs returns the members ids in the mutation.
func (m *NamespaceMutation) MembersIDs() (ids []string) {
	for id := range m.members {
		ids = append(ids, id)
	}
	return
}

// ResetMembers reset all changes of the "members" edge.
func (m *NamespaceMutation) ResetMembers() {
	m.members = nil
	m.removedmembers = nil
}

// AddAdminIDs adds the admins edge to User by ids.
func (m *NamespaceMutation) AddAdminIDs(ids ...string) {
	if m.admins == nil {
		m.admins = make(map[string]struct{})
	}
	for i := range ids {
		m.admins[ids[i]] = struct{}{}
	}
}

// RemoveAdminIDs removes the admins edge to User by ids.
func (m *NamespaceMutation) RemoveAdminIDs(ids ...string) {
	if m.removedadmins == nil {
		m.removedadmins = make(map[string]struct{})
	}
	for i := range ids {
		m.removedadmins[ids[i]] = struct{}{}
	}
}

// RemovedAdmins returns the removed ids of admins.
func (m *NamespaceMutation) RemovedAdminsIDs() (ids []string) {
	for id := range m.removedadmins {
		ids = append(ids, id)
	}
	return
}

// AdminsIDs returns the admins ids in the mutation.
func (m *NamespaceMutation) AdminsIDs() (ids []string) {
	for id := range m.admins {
		ids = append(ids, id)
	}
	return
}

// ResetAdmins reset all changes of the "admins" edge.
func (m *NamespaceMutation) ResetAdmins() {
	m.admins = nil
	m.removedadmins = nil
}

// Op returns the operation name.
func (m *NamespaceMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Namespace).
func (m *NamespaceMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *NamespaceMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, namespace.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *NamespaceMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case namespace.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *NamespaceMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case namespace.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Namespace field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *NamespaceMutation) SetField(name string, value ent.Value) error {
	switch name {
	case namespace.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Namespace field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *NamespaceMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *NamespaceMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *NamespaceMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Namespace numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *NamespaceMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *NamespaceMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *NamespaceMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Namespace nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *NamespaceMutation) ResetField(name string) error {
	switch name {
	case namespace.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Namespace field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *NamespaceMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.members != nil {
		edges = append(edges, namespace.EdgeMembers)
	}
	if m.admins != nil {
		edges = append(edges, namespace.EdgeAdmins)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *NamespaceMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case namespace.EdgeMembers:
		ids := make([]ent.Value, 0, len(m.members))
		for id := range m.members {
			ids = append(ids, id)
		}
		return ids
	case namespace.EdgeAdmins:
		ids := make([]ent.Value, 0, len(m.admins))
		for id := range m.admins {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *NamespaceMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	if m.removedmembers != nil {
		edges = append(edges, namespace.EdgeMembers)
	}
	if m.removedadmins != nil {
		edges = append(edges, namespace.EdgeAdmins)
	}
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *NamespaceMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case namespace.EdgeMembers:
		ids := make([]ent.Value, 0, len(m.removedmembers))
		for id := range m.removedmembers {
			ids = append(ids, id)
		}
		return ids
	case namespace.EdgeAdmins:
		ids := make([]ent.Value, 0, len(m.removedadmins))
		for id := range m.removedadmins {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *NamespaceMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *NamespaceMutation) EdgeCleared(name string) bool {
	switch name {
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *NamespaceMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Namespace unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *NamespaceMutation) ResetEdge(name string) error {
	switch name {
	case namespace.EdgeMembers:
		m.ResetMembers()
		return nil
	case namespace.EdgeAdmins:
		m.ResetAdmins()
		return nil
	}
	return fmt.Errorf("unknown Namespace edge %s", name)
}

// UserMutation represents an operation that mutate the Users
// nodes in the graph.
type UserMutation struct {
	config
	op                      Op
	typ                     string
	id                      *string
	name                    *string
	avatar_url              *string
	clearedFields           map[string]struct{}
	namespaces              map[string]struct{}
	removednamespaces       map[string]struct{}
	owned_namespaces        map[string]struct{}
	removedowned_namespaces map[string]struct{}
	done                    bool
	oldValue                func(context.Context) (*User, error)
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows to manage the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for $n.Name.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the id field of the mutation.
func withUserID(id string) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on User creation.
func (m *UserMutation) SetID(id string) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *UserMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the name field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old name value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName reset all changes of the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// SetAvatarURL sets the avatar_url field.
func (m *UserMutation) SetAvatarURL(s string) {
	m.avatar_url = &s
}

// AvatarURL returns the avatar_url value in the mutation.
func (m *UserMutation) AvatarURL() (r string, exists bool) {
	v := m.avatar_url
	if v == nil {
		return
	}
	return *v, true
}

// OldAvatarURL returns the old avatar_url value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldAvatarURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldAvatarURL is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldAvatarURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAvatarURL: %w", err)
	}
	return oldValue.AvatarURL, nil
}

// ClearAvatarURL clears the value of avatar_url.
func (m *UserMutation) ClearAvatarURL() {
	m.avatar_url = nil
	m.clearedFields[user.FieldAvatarURL] = struct{}{}
}

// AvatarURLCleared returns if the field avatar_url was cleared in this mutation.
func (m *UserMutation) AvatarURLCleared() bool {
	_, ok := m.clearedFields[user.FieldAvatarURL]
	return ok
}

// ResetAvatarURL reset all changes of the "avatar_url" field.
func (m *UserMutation) ResetAvatarURL() {
	m.avatar_url = nil
	delete(m.clearedFields, user.FieldAvatarURL)
}

// AddNamespaceIDs adds the namespaces edge to Namespace by ids.
func (m *UserMutation) AddNamespaceIDs(ids ...string) {
	if m.namespaces == nil {
		m.namespaces = make(map[string]struct{})
	}
	for i := range ids {
		m.namespaces[ids[i]] = struct{}{}
	}
}

// RemoveNamespaceIDs removes the namespaces edge to Namespace by ids.
func (m *UserMutation) RemoveNamespaceIDs(ids ...string) {
	if m.removednamespaces == nil {
		m.removednamespaces = make(map[string]struct{})
	}
	for i := range ids {
		m.removednamespaces[ids[i]] = struct{}{}
	}
}

// RemovedNamespaces returns the removed ids of namespaces.
func (m *UserMutation) RemovedNamespacesIDs() (ids []string) {
	for id := range m.removednamespaces {
		ids = append(ids, id)
	}
	return
}

// NamespacesIDs returns the namespaces ids in the mutation.
func (m *UserMutation) NamespacesIDs() (ids []string) {
	for id := range m.namespaces {
		ids = append(ids, id)
	}
	return
}

// ResetNamespaces reset all changes of the "namespaces" edge.
func (m *UserMutation) ResetNamespaces() {
	m.namespaces = nil
	m.removednamespaces = nil
}

// AddOwnedNamespaceIDs adds the owned_namespaces edge to Namespace by ids.
func (m *UserMutation) AddOwnedNamespaceIDs(ids ...string) {
	if m.owned_namespaces == nil {
		m.owned_namespaces = make(map[string]struct{})
	}
	for i := range ids {
		m.owned_namespaces[ids[i]] = struct{}{}
	}
}

// RemoveOwnedNamespaceIDs removes the owned_namespaces edge to Namespace by ids.
func (m *UserMutation) RemoveOwnedNamespaceIDs(ids ...string) {
	if m.removedowned_namespaces == nil {
		m.removedowned_namespaces = make(map[string]struct{})
	}
	for i := range ids {
		m.removedowned_namespaces[ids[i]] = struct{}{}
	}
}

// RemovedOwnedNamespaces returns the removed ids of owned_namespaces.
func (m *UserMutation) RemovedOwnedNamespacesIDs() (ids []string) {
	for id := range m.removedowned_namespaces {
		ids = append(ids, id)
	}
	return
}

// OwnedNamespacesIDs returns the owned_namespaces ids in the mutation.
func (m *UserMutation) OwnedNamespacesIDs() (ids []string) {
	for id := range m.owned_namespaces {
		ids = append(ids, id)
	}
	return
}

// ResetOwnedNamespaces reset all changes of the "owned_namespaces" edge.
func (m *UserMutation) ResetOwnedNamespaces() {
	m.owned_namespaces = nil
	m.removedowned_namespaces = nil
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	if m.avatar_url != nil {
		fields = append(fields, user.FieldAvatarURL)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldName:
		return m.Name()
	case user.FieldAvatarURL:
		return m.AvatarURL()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldName:
		return m.OldName(ctx)
	case user.FieldAvatarURL:
		return m.OldAvatarURL(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case user.FieldAvatarURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAvatarURL(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldAvatarURL) {
		fields = append(fields, user.FieldAvatarURL)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldAvatarURL:
		m.ClearAvatarURL()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldName:
		m.ResetName()
		return nil
	case user.FieldAvatarURL:
		m.ResetAvatarURL()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.namespaces != nil {
		edges = append(edges, user.EdgeNamespaces)
	}
	if m.owned_namespaces != nil {
		edges = append(edges, user.EdgeOwnedNamespaces)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeNamespaces:
		ids := make([]ent.Value, 0, len(m.namespaces))
		for id := range m.namespaces {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeOwnedNamespaces:
		ids := make([]ent.Value, 0, len(m.owned_namespaces))
		for id := range m.owned_namespaces {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	if m.removednamespaces != nil {
		edges = append(edges, user.EdgeNamespaces)
	}
	if m.removedowned_namespaces != nil {
		edges = append(edges, user.EdgeOwnedNamespaces)
	}
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeNamespaces:
		ids := make([]ent.Value, 0, len(m.removednamespaces))
		for id := range m.removednamespaces {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeOwnedNamespaces:
		ids := make([]ent.Value, 0, len(m.removedowned_namespaces))
		for id := range m.removedowned_namespaces {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeNamespaces:
		m.ResetNamespaces()
		return nil
	case user.EdgeOwnedNamespaces:
		m.ResetOwnedNamespaces()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
