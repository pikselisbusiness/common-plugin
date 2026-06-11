// Package common contains stable plugin DTOs shared by versioned API packages.
//
// Plugin authors should usually import a versioned package such as
// github.com/pikselisbusiness/common-plugin/v2. This package exists for
// runtime primitives that should not be duplicated between API versions, such
// as RequestContext and small helper DTOs.
package common
