DROP VIEW IF EXISTS view_rating_valuation;

CREATE VIEW view_rating_valuation AS
SELECT  a.id, a.name, a.type_id,  COALESCE(b.name, '') AS type_name, 
 a.minimum_score, a.maximum_score, a.importance, a.added_by AS added_by_id, 
 c.username AS added_by_username, DATE(a.create_time) AS create_date, a.create_time  
FROM rating_valuation a  
LEFT JOIN rating_valuation_type b ON a.type_id=b.id 
LEFT JOIN user c ON a.added_by=c.id;