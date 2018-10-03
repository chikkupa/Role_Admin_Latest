DROP VIEW IF EXISTS view_user_files;

CREATE VIEW view_user_files AS
SELECT a.id, a.user_id, a.file_category as category_id, b.name as file_category, a.filename, 
a.status, a.expiry_date, a.message 
FROM user_files a 
LEFT JOIN user_file_category b on a.file_category=b.id