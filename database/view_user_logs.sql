DROP VIEW IF EXISTS view_user_log;
CREATE VIEW view_user_log AS
 SELECT 
  a.id, 
  a.user_id, 
  b.username, 
  c.name,
  b.role,
  b.status,
  a.user_type, 
  DATE_FORMAT(a.date_time, '%d/%m/%Y %h:%i:%s') AS date_time, 
  a.actions 
   FROM user_log a 
    LEFT JOIN user b ON a.user_id=b.id
    LEFT JOIN  (
     SELECT user_id,name FROM user_details 
      UNION 
     SELECT user_id, name FROM admin_user_details
     ) c ON a.user_id=c.user_id;