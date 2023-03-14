SELECT
  COALESCE(AVG(cr.rating), 0) AS rating,
  COALESCE(COUNT(cr.rating), 0) AS count_review
FROM
  companies c
  LEFT JOIN company_reviews cr ON c.id = cr.company_id
WHERE
  c.id = $1
GROUP BY
  c.id
