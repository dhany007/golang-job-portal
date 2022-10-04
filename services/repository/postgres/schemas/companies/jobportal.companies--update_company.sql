UPDATE
  companies
SET
  email = $1,
  name = $2,
  description = $3,
  address = $4,
  website = $5,
  phone_number = $6,
  telp_number = $7,
  profil_picture_url = $8,
  dresscode_code = $9,
  size_code = $10,
  modified_at = $11
WHERE
  id = $12
