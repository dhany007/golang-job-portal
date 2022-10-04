SELECT
  cb.id,
  cbc.value
FROM
  company_benefits cb
  JOIN company_benefits_codes cbc ON cb.benefit_id = cbc.id
WHERE
  cb.company_id = $1
