-- SQLBook: Code
-- SQL Server script
USE master;

GO
-- Drop the database if it exists
IF EXISTS (
    SELECT name
    FROM master.dbo.sysdatabases
    WHERE
        name = N'test'
) BEGIN

DROP DATABASE test;

DROP DATABASE test2;

DROP DATABASE test3;

END;

ELSE BEGIN PRINT 'Database does not exist... created it';

CREATE DATABASE test;

CREATE DATABASE test2;

CREATE DATABASE test3;

END;

GO