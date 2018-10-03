DROP VIEW IF EXISTS view_rating;

CREATE VIEW view_rating AS
SELECT a.id, a.valuation_id, c.name AS valuation_name, a.user_id, b.username, 
 a.description, a.date_valuation, a.given_score, a.staff_id, COALESCE(d.username, '') AS staff_username,
 a.comments, a.score_percentage, a.status, DATE(a.create_time) AS create_date, a.create_time 
FROM rating a left join user b on a.user_id=b.id 
LEFT JOIN rating_valuation c on a.valuation_id=c.id 
LEFT JOIN user d on a.staff_id=d.id;