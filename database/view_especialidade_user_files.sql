DROP VIEW IF EXISTS view_especialidade_user_files;

CREATE VIEW view_especialidade_user_files AS
SELECT a.id, a.especialidade_id, b.name AS especialidade_name,  a.user_id, c.username, d.name, a.filename, a.expiry_date, a.status, a.message, a.create_time 
 FROM especialidade_user_files a 
 LEFT JOIN especialidade b on a.especialidade_id = b.id
 LEFT JOIN user c on a.user_id = c.id 
 LEFT JOIN user_details d on a.user_id=d.user_id;