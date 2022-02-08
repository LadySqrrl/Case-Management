-- To run docker container mysql server:
-- docker exec -NAMEOFMYSQLSERVER bash
-- mysql -uroot -p

-- Reevals due during the entire school year including name, due date, roster teacher
SELECT CONCAT(s.student_first_name, ' ', s.student_last_name), p.provider_name, s.reeval_due,
CASE WHEN s.annual_due < s.reeval_due THEN s.annual_due ELSE s.reeval_due END AS "Meeting" 
FROM students s
INNER JOIN providers p 
ON s.providerID = p.providerID
WHERE s.reeval_due < '2022-10-01'
ORDER BY "Meeting";

-- Reevals due in a given month
SELECT CONCAT(s.student_first_name, ' ', s.student_last_name), p.provider_name, s.reeval_due,
CASE WHEN s.annual_due < s.reeval_due THEN s.annual_due ELSE s.reeval_due END AS "Meeting" 
FROM students s
INNER JOIN providers p 
ON s.providerID = p.providerID
WHERE s.reeval_due BETWEEN '2021-11-01' AND '2021-12-01'
ORDER BY s.reeval_due;

