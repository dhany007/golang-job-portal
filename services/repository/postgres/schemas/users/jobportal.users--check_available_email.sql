SELECT
  id,
  email
FROM
  users
WHERE
  email = $1
