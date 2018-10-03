DROP VIEW IF EXISTS view_activity_records;

CREATE VIEW view_activity_records AS
select
a.id, a.user_id, b.name, b.lastname, b.crm, a.activity, a.date, a.duration, 
a.assistant, c.name as assistant_name, c.crm as assistant_crm, a.description, a.create_time
from activity_record a 
left join user_details b on a.user_id=b.user_id 
left join user_details c on a.assistant=c.user_id;