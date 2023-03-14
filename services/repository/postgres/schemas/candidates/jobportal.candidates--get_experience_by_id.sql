SELECT
  id,
  candidate_id,
  company_name,
  title,
  description,
  date_start,
  date_end,
  created_at,
  modified_at
FROM
  candidate_experiences
WHERE
  id = $1
