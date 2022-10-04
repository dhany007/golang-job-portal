SELECT
  id,
  email,
  COALESCE(hash_password, '') AS hash_password,
  COALESCE(is_active, 1) AS is_active,
  COALESCE(role, 2) AS role,
  created_at,
  modified_at
FROM
  users
WHERE
  email=$1;