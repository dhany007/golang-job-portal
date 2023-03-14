SELECT
  id,
  COALESCE(email, '') AS email,
  COALESCE(first_name, '') AS first_name,
  COALESCE(last_name, '') AS last_name,
  COALESCE(phone_number, '') AS phone_number,
  COALESCE(telp_number, '') AS telp_number,
  COALESCE(address, '') AS address,
  COALESCE(profil_picture_url, '') AS profil_picture_url,
  created_at,
  modified_at
FROM
  candidates
WHERE
  id = $1
