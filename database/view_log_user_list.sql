DROP VIEW IF EXISTS view_log_user_list;

CREATE VIEW view_log_user_list AS
SELECT 
max(a.id) AS id, a.user_id, b.username, c.name, b.role, b.status, max(a.date_time) AS date_time 
FROM user_log a 
LEFT JOIN user b ON a.user_id=b.id 
LEFT JOIN 
(SELECT user_id, name FROM user_details UNION SELECT user_id, name FROM admin_user_details) c ON a.user_id=c.user_id
GROUP BY user_id, name ORDER BY id DESC;