// Code generated by ent, DO NOT EDIT.

package ent

import (
	"test/ent/comment"
	"test/ent/schema"
	"test/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	commentFields := schema.Comment{}.Fields()
	_ = commentFields
	// commentDescComment is the schema descriptor for comment field.
	commentDescComment := commentFields[1].Descriptor()
	// comment.DefaultComment holds the default value on creation for the comment field.
	comment.DefaultComment = commentDescComment.Default.(string)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[1].Descriptor()
	// user.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	user.AgeValidator = userDescAge.Validators[0].(func(int) error)
	// userDescNickname is the schema descriptor for nickname field.
	userDescNickname := userFields[2].Descriptor()
	// user.DefaultNickname holds the default value on creation for the nickname field.
	user.DefaultNickname = userDescNickname.Default.(string)
}
