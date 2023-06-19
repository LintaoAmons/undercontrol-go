ALTER TABLE account 
ALTER COLUMN amount TYPE varchar(50)
USING (amount/100.0)::varchar(50); -- May cause error if there's null value in the db
