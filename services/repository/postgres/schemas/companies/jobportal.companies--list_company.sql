SELECT
  c.id,
  COALESCE(c.name, '') AS name,
  COALESCE(AVG(cr.rating), 0) AS rating,
  COALESCE(COUNT(cr.rating), 0) AS count_review
FROM
  companies c
  LEFT JOIN company_reviews cr ON c.id = cr.company_id
GROUP BY
  c.id
ORDER BY
  rating DESC
LIMIT $1
OFFSET $2
