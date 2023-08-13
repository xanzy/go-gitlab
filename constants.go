package gitlab

type PipelineScopeValue string

const (
	PipelineScopeValueRunning  PipelineScopeValue = "running"
	PipelineScopeValuePending  PipelineScopeValue = "pending"
	PipelineScopeValueFinished PipelineScopeValue = "finished"
	PipelineScopeValueBranches PipelineScopeValue = "branches"
	PipelineScopeValueTags     PipelineScopeValue = "tags"
)

type PipelineStatusValue string

const (
	PipelineStatusValueCreated            PipelineStatusValue = "created"
	PipelineStatusValueWaitingForResource PipelineStatusValue = "waiting_for_resource"
	PipelineStatusValuePreparing          PipelineStatusValue = "preparing"
	PipelineStatusValuePending            PipelineStatusValue = "pending"
	PipelineStatusValueRunning            PipelineStatusValue = "running"
	PipelineStatusValueSuccess            PipelineStatusValue = "success"
	PipelineStatusValueFailed             PipelineStatusValue = "failed"
	PipelineStatusValueCanceled           PipelineStatusValue = "canceled"
	PipelineStatusValueSkipped            PipelineStatusValue = "skipped"
	PipelineStatusValueManual             PipelineStatusValue = "manual"
	PipelineStatusValueScheduled          PipelineStatusValue = "scheduled"
)

// PipelineStatus is a helper routine that allocates a new PipelineStatusValue
// to store v and returns a pointer to it.
func PipelineStatus(v PipelineStatusValue) *PipelineStatusValue {
	p := new(PipelineStatusValue)
	*p = v
	return p
}

type PipelineSourceValue string

const (
	PipelineSourceValuePush                     PipelineSourceValue = "push"
	PipelineSourceValueWeb                      PipelineSourceValue = "web"
	PipelineSourceValueTrigger                  PipelineSourceValue = "trigger"
	PipelineSourceValueSchedule                 PipelineSourceValue = "schedule"
	PipelineSourceValueApi                      PipelineSourceValue = "api"
	PipelineSourceValueExternal                 PipelineSourceValue = "external"
	PipelineSourceValuePipeline                 PipelineSourceValue = "pipeline"
	PipelineSourceValueChat                     PipelineSourceValue = "chat"
	PipelineSourceValueWebide                   PipelineSourceValue = "webide"
	PipelineSourceValueMergeRequestEvent        PipelineSourceValue = "merge_request_event"
	PipelineSourceValueExternalPullRequestEvent PipelineSourceValue = "external_pull_request_event"
	PipelineSourceValueParentPipeline           PipelineSourceValue = "parent_pipeline"
	PipelineSourceValueOnDemandDastScane        PipelineSourceValue = "ondemand_dast_scan"
	PipelineSourceValueOnDemandDastValidation   PipelineSourceValue = "ondemand_dast_validation"
)

type PipelineSortValue string

const (
	PipelineSortValueAsc  PipelineSortValue = "asc"
	PipelineSortValueDesc PipelineSortValue = "desc"
)

type PipelineOrderByValue string

const (
	PipelineOrderByValueID        PipelineOrderByValue = "id"
	PipelineOrderByValueStatus    PipelineOrderByValue = "status"
	PipelineOrderByValueRef       PipelineOrderByValue = "ref"
	PipelineOrderByValueUpdatedAt PipelineOrderByValue = "updated_at"
	PipelineOrderByValueUserID    PipelineOrderByValue = "user_id"
)

type MergeRequestOrderByValue string

const (
	MergeRequestOrderByValueCreatedAt MergeRequestOrderByValue = "created_at"
	MergeRequestOrderByValueTitle     MergeRequestOrderByValue = "title"
	MergeRequestOrderByValueUpdatedAt MergeRequestOrderByValue = "updated_at"
)

type MergeRequestScopeValue string

const (
	MergeRequestScopeValueCreatedByMe  MergeRequestScopeValue = "created_by_me"
	MergeRequestScopeValueAssignedToMe MergeRequestScopeValue = "assigned_to_me"
	MergeRequestScopeValueAll          MergeRequestScopeValue = "all"
)

type MergeRequestSortValue string

const (
	MergeRequestSortValueAsc  MergeRequestSortValue = "asc"
	MergeRequestSortValueDesc MergeRequestSortValue = "desc"
)

type MergeRequestStateValue string

const (
	MergeRequestStateValueOpened MergeRequestStateValue = "opened"
	MergeRequestStateValueClosed MergeRequestStateValue = "closed"
	MergeRequestStateValueLocked MergeRequestStateValue = "locked"
	MergeRequestStateValueMerged MergeRequestStateValue = "merged"
)

type MergeRequestDetailedMergeStatusValue string

const (
	MergeRequestDetailedMergeStatusValueBlockedStatus          MergeRequestDetailedMergeStatusValue = "blocked_status"           //Blocked by another merge request
	MergeRequestDetailedMergeStatusValueBrokenStatus           MergeRequestDetailedMergeStatusValue = "broken_status"            //Can’t merge into the target branch due to a potential conflict
	MergeRequestDetailedMergeStatusValueChecking               MergeRequestDetailedMergeStatusValue = "checking"                 //Git is testing if a valid merge is possible
	MergeRequestDetailedMergeStatusValueUnchecked              MergeRequestDetailedMergeStatusValue = "unchecked"                //Git has not yet tested if a valid merge is possible
	MergeRequestDetailedMergeStatusValueCIMustPass             MergeRequestDetailedMergeStatusValue = "ci_must_pass"             //A CI/CD pipeline must succeed before merge
	MergeRequestDetailedMergeStatusValueCIStillRunning         MergeRequestDetailedMergeStatusValue = "ci_still_running"         //A CI/CD pipeline is still running
	MergeRequestDetailedMergeStatusValueDiscussionsNotResolved MergeRequestDetailedMergeStatusValue = "discussions_not_resolved" //All discussions must be resolved before merge
	MergeRequestDetailedMergeStatusValueDraftStatus            MergeRequestDetailedMergeStatusValue = "draft_status"             //Can’t merge because the merge request is a draft
	MergeRequestDetailedMergeStatusValueExternalStatusCheck    MergeRequestDetailedMergeStatusValue = "external_status_check"    //All status checks must pass before merge
	MergeRequestDetailedMergeStatusValueMergeable              MergeRequestDetailedMergeStatusValue = "mergeable"                //The branch can merge cleanly into the target branch
	MergeRequestDetailedMergeStatusValueNotApproved            MergeRequestDetailedMergeStatusValue = "not_approved"             //Approval is required before merge
	MergeRequestDetailedMergeStatusValueNotOpened              MergeRequestDetailedMergeStatusValue = "not_opened"               //The merge request must be open before merge.
	MergeRequestDetailedMergeStatusValuePoliciesDenied         MergeRequestDetailedMergeStatusValue = "policies_denied"          //The merge request contains denied policies
)
