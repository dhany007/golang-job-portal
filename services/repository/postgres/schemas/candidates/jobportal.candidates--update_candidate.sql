UPDATE
  candidates
SET
  first_name = $1,
  last_name = $2,
  phone_number = $3,
  telp_number = $4,
  address = $5,
  profil_picture_url = $6
WHERE
  id = $7
