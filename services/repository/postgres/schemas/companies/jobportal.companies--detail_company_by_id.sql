SELECT
  COALESCE(c.email, '') AS email,
  COALESCE(c.name, '') AS name,
  COALESCE(c.description, '') AS description,
  COALESCE(c.address, '') AS address,
  COALESCE(c.website, '') AS website,
  COALESCE(c.phone_number, '') AS phone_number,
  COALESCE(c.telp_number, '') AS telp_number,
  COALESCE(c.profil_picture_url, '') AS profil_picture_url,
  COALESCE(cdc.value, '') AS dress,
  COALESCE(czc.value, '') AS size,
  c.created_at,
  c.modified_at
FROM
  companies c
  LEFT JOIN company_dresscode_codes cdc ON c.dresscode_code = cdc.id
  LEFT JOIN company_size_codes czc ON c.size_code = czc.id
WHERE
  c.id = $1
