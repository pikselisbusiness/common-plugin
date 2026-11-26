package v2

import "time"

// RepairStatus is one of the organization's repair statuses. Choose by id in
// plugin settings, never by name.
type RepairStatus struct {
	StatusId uint   `json:"statusId"`
	Name     string `json:"name"`
	Color    string `json:"color"`
	OrderBy  int    `json:"orderBy"`
}

// Repair is the part of a repair a plugin keeps in step with an outside
// system. Field names follow the external API repair.
type Repair struct {
	RepairId             uint      `json:"repairId"`
	RepairStatusId       uint      `json:"repairStatusId"`
	CompanyId            uint      `json:"companyId"`
	RegistrationDatetime time.Time `json:"registrationDatetime"`
	DeviceName           string    `json:"deviceName"`
	SerialNumber         string    `json:"serialNumber"`
	CustomerToldDefect   string    `json:"customerToldDefect"`
	IsWarrantable        bool      `json:"isWarrantable"`
}

// RepairCreateRequest creates a repair for a company (the customer the
// devices come from). Field names follow the external API POST /repairs.
type RepairCreateRequest struct {
	CompanyId            uint      `json:"companyId"`
	RepairStatusId       uint      `json:"repairStatusId"`
	RegistrationDatetime time.Time `json:"registrationDatetime"` // zero: now
	DeviceName           string    `json:"deviceName"`
	SerialNumber         string    `json:"serialNumber"`
	CustomerToldDefect   string    `json:"customerToldDefect"`
	IsWarrantable        bool      `json:"isWarrantable"`
	// DevicePassword is the device unlock code the technician needs; it is
	// stored on the repair as the person-entered one is.
	DevicePassword string `json:"devicePassword"`
	Comments       string `json:"comments"`
	// ResponsibleUserId names the organization user the repair is assigned
	// to; 0 leaves it unassigned.
	ResponsibleUserId uint `json:"responsibleUserId"`
}

// RepairPatchRequest changes only the fields that are set.
type RepairPatchRequest struct {
	DeviceName         *string `json:"deviceName,omitempty"`
	SerialNumber       *string `json:"serialNumber,omitempty"`
	CustomerToldDefect *string `json:"customerToldDefect,omitempty"`
	IsWarrantable      *bool   `json:"isWarrantable,omitempty"`
}

type RepairMutateResponse struct {
	RepairId uint `json:"repairId"`
}

// RepairComment is one comment of a repair, as plain text.
type RepairComment struct {
	CommentId uint      `json:"commentId"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
	// CreatedUserId is the organization user who wrote it; 0 for comments the
	// system or a plugin wrote.
	CreatedUserId uint `json:"createdUserId"`
}

// RepairCommentsRequest pages a repair's comments oldest first.
type RepairCommentsRequest struct {
	// AfterCommentId returns only comments with a larger id (0: from the first).
	AfterCommentId uint `json:"afterCommentId"`
	Limit          int  `json:"limit"` // 0 or more than 100: 100
}

// RepairCommentCreateRequest is a plain-text internal comment.
type RepairCommentCreateRequest struct {
	Text string `json:"text"`
}

type RepairCommentMutateResponse struct {
	CommentId uint `json:"commentId"`
}

// RepairImage is one image added to a repair's photos (JPEG, PNG or WEBP),
// shaped like ProductImage.
type RepairImage struct {
	FileName string `json:"fileName"`
	FileBlob []byte `json:"fileBlob"`
}

// RepairsAPI is the host's Repairs module, for plugins that bring repairs in
// from another system and keep them in step.
type RepairsAPI interface {
	// GetRepairStatuses lists the organization's repair statuses, for plugin
	// settings where the organization picks the status of each plugin stage.
	GetRepairStatuses(context RequestContext) ([]RepairStatus, error)
	GetRepair(context RequestContext, repairId uint) (Repair, error)
	// CreateRepair creates a repair as a person would (numbering, activity,
	// repair.created event, afterCreate workflows) and records the plugin.
	CreateRepair(context RequestContext, request RepairCreateRequest) (RepairMutateResponse, error)
	// PatchRepair writes the set fields; the activity log records the change
	// and afterUpdate workflows run.
	PatchRepair(context RequestContext, repairId uint, request RepairPatchRequest) (RepairMutateResponse, error)
	// SetRepairStatus moves a repair to another of the organization's statuses
	// with the status log, event and workflows of a person's change. Status
	// SMS and emails go out as the status is configured.
	SetRepairStatus(context RequestContext, repairId uint, repairStatusId uint) error
	GetRepairComments(context RequestContext, repairId uint, request RepairCommentsRequest) ([]RepairComment, error)
	// CreateRepairComment adds a comment written by the plugin (CreatedUserId
	// 0) and notifies the repair's responsible user.
	CreateRepairComment(context RequestContext, repairId uint, request RepairCommentCreateRequest) (RepairCommentMutateResponse, error)
	// AddRepairImages appends images to the repair's photos.
	AddRepairImages(context RequestContext, repairId uint, images []RepairImage) error
}
