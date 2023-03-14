SELECT
  COALESCE(id, 0) AS id,
  company_id,
  candidate_id,
  COALESCE(rating, 0) AS rating,
  COALESCE(review, '') AS review
FROM
  company_reviews
WHERE
  company_id = $1
ORDER BY
  created_at DESC
LIMIT $2 OFFSET $3
