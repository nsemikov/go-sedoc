package argument

import (
	sedoc "github.com/nsemikov/go-sedoc"
)

const (
	// RegexpDuration var
	RegexpDuration = "^((\\d+(\\.\\d*)?h)(\\d(\\.\\d*)?m)?(\\d(\\.\\d*)?s)?(\\d(\\.\\d*)?[nuµm]s)?)|((\\d+(\\.\\d*)?h)?(\\d(\\.\\d*)?m)(\\d(\\.\\d*)?s)?(\\d(\\.\\d*)?[nuµm]s)?)|((\\d+(\\.\\d*)?h)?(\\d(\\.\\d*)?m)?(\\d(\\.\\d*)?s)(\\d(\\.\\d*)?[nuµm]s)?)|((\\d+(\\.\\d*)?h)?(\\d(\\.\\d*)?m)?(\\d(\\.\\d*)?s)?(\\d(\\.\\d*)?[nuµm]s))$"
	// RegexpEmail var
	RegexpEmail = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	// RegexpUUID var
	RegexpUUID = "^[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}$"
	// RegexpInteger var
	RegexpInteger = "^[0-9]+$"
)

// Required func return copy of arg with Required field setted to true
func Required(arg sedoc.Argument) sedoc.Argument {
	arg.Required = true
	return arg
}

// Nullable func return copy of arg with Nullable field setted to true
func Nullable(arg sedoc.Argument) sedoc.Argument {
	arg.Nullable = true
	return arg
}

var (
	// ID is integer identifier argument
	ID = sedoc.Argument{Name: "id", Type: sedoc.ArgumentTypeInteger, Description: "Identifier", RegExp: RegexpInteger}
	// UUID is uuid identifier argument
	UUID = sedoc.Argument{Name: "id", Type: sedoc.ArgumentTypeUUID, Description: "Identifier", RegExp: RegexpUUID}
	// Login argument
	Login = sedoc.Argument{Name: "login", Type: sedoc.ArgumentTypeString, Description: "Login string"}
	// Password argument
	Password = sedoc.Argument{Name: "password", Type: sedoc.ArgumentTypeString, Description: "Password string"}
	// Surname argument
	Surname = sedoc.Argument{Name: "surname", Type: sedoc.ArgumentTypeString, Description: "Surname string"}
	// Name argument
	Name = sedoc.Argument{Name: "name", Type: sedoc.ArgumentTypeString, Description: "Name string"}
	// Patronymic argument
	Patronymic = sedoc.Argument{Name: "patronymic", Type: sedoc.ArgumentTypeString, Description: "Patronymic string"}
	// Email argument
	Email = sedoc.Argument{Name: "email", Type: sedoc.ArgumentTypeString, Description: "Email string", RegExp: RegexpEmail}
	// EmailNull argument
	EmailNull = sedoc.Argument{Name: "email_null", Type: sedoc.ArgumentTypeBoolean, Description: ""}
	// EmailNotNull argument
	EmailNotNull = sedoc.Argument{Name: "email_not_null", Type: sedoc.ArgumentTypeBoolean, Description: ""}
	// CreatedLater argument
	CreatedLater = sedoc.Argument{Name: "created_later", Type: sedoc.ArgumentTypeTime, Description: "Created later then time"}
	// CreatedEarlier argument
	CreatedEarlier = sedoc.Argument{Name: "created_earlier", Type: sedoc.ArgumentTypeTime, Description: "Created earlier then time"}
	// UpdatedLater argument
	UpdatedLater = sedoc.Argument{Name: "updated_later", Type: sedoc.ArgumentTypeTime, Description: "Updated later then time"}
	// UpdatedEarlier argument
	UpdatedEarlier = sedoc.Argument{Name: "updated_earlier", Type: sedoc.ArgumentTypeTime, Description: "Updated earlier then time"}
	// DeletedLater argument
	DeletedLater = sedoc.Argument{Name: "deleted_later", Type: sedoc.ArgumentTypeTime, Description: "Deleted later then time"}
	// DeletedEarlier argument
	DeletedEarlier = sedoc.Argument{Name: "deleted_earlier", Type: sedoc.ArgumentTypeTime, Description: "Deleted earlier then time"}
	// Count argument
	Count = sedoc.Argument{Name: "count", Type: sedoc.ArgumentTypeInteger, Description: "Count of items"}
	// Offset argument
	Offset = sedoc.Argument{Name: "offset", Type: sedoc.ArgumentTypeInteger, Description: "Items offset"}
	// Deleted argument
	Deleted = sedoc.Argument{Name: "deleted", Type: sedoc.ArgumentTypeBoolean, Description: "Get deleted and undeleted items"}
	// DeletedOnly argument
	DeletedOnly = sedoc.Argument{Name: "deleted_only", Type: sedoc.ArgumentTypeBoolean, Description: "Get only deleted items"}
	// ActivationCode argument
	ActivationCode = sedoc.Argument{Name: "activation_code", Type: sedoc.ArgumentTypeString, Description: "Signup confirmation code"}
	// Active argument
	Active = sedoc.Argument{Name: "active", Type: sedoc.ArgumentTypeBoolean, Description: "Signup confirmed"}
	// RoleID is integer role identifier argument
	RoleID = sedoc.Argument{Name: "role_id", Type: sedoc.ArgumentTypeInteger, Description: "User role identifier", RegExp: RegexpInteger}
	// RoleUUID is uuid role identifier argument
	RoleUUID = sedoc.Argument{Name: "role_id", Type: sedoc.ArgumentTypeUUID, Description: "User role identifier", RegExp: RegexpUUID}
	// FinallyDelete argument
	FinallyDelete = sedoc.Argument{Name: "finally", Type: sedoc.ArgumentTypeBoolean, Description: "Finally delete"}
)
