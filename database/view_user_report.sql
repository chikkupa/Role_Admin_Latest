DROP VIEW IF EXISTS view_user_report;

CREATE VIEW view_user_report AS
SELECT a.id AS user_id, a.username, b.name, b.lastname, b.create_time, 
date(b.create_time) AS create_date, a.status, 
COALESCE(GROUP_CONCAT(d.name SEPARATOR ', '), '' ) AS especialidades  
FROM user a 
LEFT JOIN user_details b ON a.id=b.user_id 
LEFT JOIN especialidade_user_files c ON a.id=c.user_id 
LEFT JOIN especialidade d ON c.especialidade_id=d.id 
WHERE a.user_type=2 GROUP BY a.id, a.username, b.name, b.lastname, b.create_time, a.status;