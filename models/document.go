package models

import (
	"time"
)

// DocumentDto represents the equivalent Go struct for DocumentDto class.
type DocumentDto struct {
	ID                           string                   `json:"id"`
	CreatedTime                  time.Time                `json:"created_time"`
	DetailsLastUpdatedTime       *time.Time               `json:"details_last_updated_time,omitempty"`
	DetailsLastUpdatedBy         *string                  `json:"details_last_updated_by,omitempty"`
	DetailsLastUpdatedByName     *string                  `json:"details_last_updated_by_name,omitempty"`
	DetailsUpdated               bool                     `json:"details_updated"`
	Title                        string                   `json:"title"`
	DocumentCode                 *string                  `json:"document_code,omitempty"`
	Description                  *string                  `json:"description,omitempty"`
	IssuingAuthority             *string                  `json:"issuing_authority,omitempty"`
	DocumentTemplateID           *string                  `json:"document_template_id,omitempty"`
	Priority                     int                      `json:"priority"`
	InputFileURLs                []string                 `json:"input_file_urls"`
	Configs                      []DocumentTemplateDetail `json:"configs"`
	Snippet                      *string                  `json:"snippet,omitempty"`
	OriginalLangCode             string                   `json:"original_lang_code"`
	TranslateLangCode            string                   `json:"translate_lang_code"`
	Metadata                     any                      `json:"metadata,omitempty"`
	EffectiveStartTime           *time.Time               `json:"effective_start_time,omitempty"`
	EffectiveEndTime             *time.Time               `json:"effective_end_time,omitempty"`
	Keywords                     []string                 `json:"keywords,omitempty"`
	KeywordTypes                 *KeywordTypes            `json:"keyword_types,omitempty"`
	Signer                       *string                  `json:"signer,omitempty"`
	IsDetectFace                 bool                     `json:"is_detect_face"`
	CanFindDocumentByImage       bool                     `json:"can_find_document_by_image"`
	UploadConfigID               *string                  `json:"upload_config_id,omitempty"`
	Publisher                    *string                  `json:"publisher,omitempty"`
	FileType                     FileTypes                `json:"file_type"`
	Status                       DocumentStatus           `json:"status"`
	OCRProcessStatus             ProcessStatuses          `json:"ocr_process_status"`
	FaceDetectProcessStatus      ProcessStatuses          `json:"face_detect_process_status"`
	ExtractPureInfoProcessStatus ProcessStatuses          `json:"extract_pure_info_process_status"`
	ExtractContentProcessStatus  ProcessStatuses          `json:"extract_content_process_status"`
	LegalDocumentProcessStatus   ProcessStatuses          `json:"legal_document_process_status"`
	ApproveStatus                ApproveStatus            `json:"approve_status"`
	CreatorID                    string                   `json:"creator_id"`
	CreatorName                  *string                  `json:"creator_name,omitempty"`
	InsertedTime                 time.Time                `json:"inserted_time"`
	SummaryAll                   *string                  `json:"summary_all,omitempty"`
	Author                       *string                  `json:"author,omitempty"`
	Autograph                    *string                  `json:"autograph,omitempty"`
	ReliabilityLevel             *ReliabilityLevel        `json:"reliability_level,omitempty"`
	PhysicalState                *PhysicalState           `json:"physical_state,omitempty"`
	Classification               *Classification          `json:"classification,omitempty"`
	ColorType                    *ColorType               `json:"color_type,omitempty"`
	HasAttachment                *HasAttachment           `json:"has_attachment,omitempty"`
	ResumeID                     *string                  `json:"resume_id,omitempty"`
	Privacy                      Privacy                  `json:"privacy"`
	HasBackup                    bool                     `json:"has_backup"`
	BackupStatus                 BackupStatus             `json:"backup_status"`
	Genre                        *DocumentGenre           `json:"genre,omitempty"`
	Form                         *DocumentForm            `json:"form,omitempty"`
	Note                         *string                  `json:"note,omitempty"`
	InputSourceType              *string                  `json:"input_source_type,omitempty"`
	IntegrationID                *string                  `json:"integration_id,omitempty"`
}

// Enums
type FileTypes int

const (
// Define your values here based on the original C# enum.
)

type DocumentStatus int

const (
	NotStart DocumentStatus = iota
	// Add more statuses as needed
)

type ProcessStatuses int

const (
	Pending ProcessStatuses = iota
	// Add more statuses as needed
)

type ApproveStatus int

const (
	Draft ApproveStatus = iota
	// Add more statuses as needed
)

// Additional types
type DocumentTemplateDetail struct {
	// Define fields here
}

type KeywordTypes struct {
	ListNamePerson    []string `json:"list_name_person"`
	ListOrganizations []string `json:"list_organizations"`
	ListLocation      []string `json:"list_location"`
	ListJobPosition   []string `json:"list_job_position"`
}

type TagDto struct {
	// Define fields here
}

type ExtractContentDto struct {
	// Define fields here
}

type DocumentFaceDto struct {
	// Define fields here
}

type DocumentDetailDto struct {
	// Define fields here
}

type DocumentTemplateDto struct {
	// Define fields here
}

type UploadConfigDto struct {
	// Define fields here
}

type TopicDto struct {
	// Define fields here
}

type RecommendedDocumentTopicDto struct {
	// Define fields here
}

type EntityActivityDto struct {
	// Define fields here
}

type DocumentTopicDto struct {
	// Define fields here
}

type TopicTrainingDocumentPrototypeDto struct {
	// Define fields here
}

type DocumentTypeDto struct {
	// Define fields here
}

type DocumentDocumentTypeDto struct {
	// Define fields here
}

type ResumeDto struct {
	// Define fields here
}

type DocumentTagDto struct {
	// Define fields here
}

type DocumentViewLogDto struct {
	// Define fields here
}

type ReliabilityLevel int

const (
	ElectronicOriginal ReliabilityLevel = iota + 1
	Digitalization
	Mixed
)

type PhysicalState int

const (
	Good PhysicalState = iota + 1
	Normal
	Damaged
)

type Classification int

const (
	Movie Classification = iota + 1
	Image
)

type ColorType int

const (
	Color ColorType = iota + 1
	BlackWhite
)

type HasAttachment int

const (
	Yes HasAttachment = iota + 1
	No
)

type Privacy int

const (
	Public Privacy = iota
	Conditional
	Private
)

type BackupStatus int

const (
	NotBackedUp BackupStatus = iota
	BackedUp
)

type DocumentGenre int

const (
	Resolution DocumentGenre = iota
	Decree
	Directive
	Regulation
	Rule
	Announcement
	Notice
	Instruction
)

type DocumentForm int

const (
	Manuscript DocumentForm = iota
	TechnicalDocument
	AudioVisualDocument
	AudioRecordingDocument
)
