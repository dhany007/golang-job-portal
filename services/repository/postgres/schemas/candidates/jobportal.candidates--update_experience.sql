UPDATE
  candidate_experiences
SET
  company_name = $1,
  title = $2,
  description = $3,
  modified_at = $4
WHERE
  candidate_id = $5
  AND id = $6