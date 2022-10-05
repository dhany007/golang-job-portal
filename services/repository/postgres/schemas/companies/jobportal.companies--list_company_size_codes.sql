SELECT
  COALESCE(id, 0) AS id,
  COALESCE(value, '') AS value
FROM
  company_size_codes
