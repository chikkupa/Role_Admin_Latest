DROP VIEW IF EXISTS view_staffs;

CREATE VIEW  view_staffs AS
SELECT a. id, a.username, a.password, a.status, a.role, a.permissions, a.verify_string, 
 b.name, b.lastname, b.email, b.age, b.gender, b.task  
FROM user a LEFT JOIN admin_user_details b ON a.id=b.user_id where a.user_type=1;