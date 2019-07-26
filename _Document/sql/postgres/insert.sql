



INSERT INTO image_manage (
  id,
  exif_lat,
  exif_lon,
  exif_alt,
  created_by,
  updated_by
) VALUES (
  1,
  35.000001,
  139.000002,
  10.001,
  99999999,
  99999999
);




INSERT INTO role (
  id,
  roleName,
  createdBy,
  updatedBy
) VALUES (
  1,
  'ROLE_ADMIN',
  99999999,
  99999999
), (
  2,
  'ROLE_OPERATOR',
  99999999,
  99999999
), (
  3,
  'ROLE_SALESFORCE',
  99999999,
  99999999
);

-- 自動採番項目の値を指定すると、項目の最大値とシーケンス番号がずれて
-- 次回Insert時にエラーとなるため、ここでシーケンス番号を設定し直す
SELECT setval('role_id_seq', (SELECT MAX(id) FROM role));