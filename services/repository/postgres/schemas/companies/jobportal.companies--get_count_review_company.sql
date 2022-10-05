SELECT
  COUNT(*)
FROM
  company_reviews
WHERE
  company_id = $1
