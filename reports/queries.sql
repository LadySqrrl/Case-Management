SELECT CONCAT(s.student_first_name, ' ', s.student_last_name), p.provider_name, s.reeval_due,
CASE WHEN s.annual_due < s.reeval_due THEN s.annual_due ELSE s.reeval_due END AS "Meeting" 
FROM students s
INNER JOIN providers p 
ON s.providerID = p.providerID
WHERE s.reeval_due < '2022-10-01'
ORDER BY s.reeval_due;

IF annual_due < reeval_due 
SELECT annual_due 
FROM students
ELSE SELECT reeval_due 
FROM students
-- WHERE reeval_due < '2022-10-01';


SELECT CASE WHEN annual_due < reeval_due THEN annual_due ELSE reeval_due END AS "Meeting" FROM students;

