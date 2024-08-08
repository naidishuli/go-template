package task

const (
    TypeSendEmail             = "send_email"
    TypeSendConfirmationEmail = "send_confirmation_email"
    TypeIngestCompanyFile     = "ingest_company_file"
    TypeIngestJob             = "ingest_job"
    TypeIngestJobFeed         = "ingest_job_feed"
)

const (
    QueueDefault = "default:queue"

    QueueAuth          = "auth:queue"
    QueueIngestCompany = "ingest_company:queue"
    QueueIngestJob     = "ingest_job:queue"
    QueueIngestJobFeed = "ingest_job_feed:queue"
)
