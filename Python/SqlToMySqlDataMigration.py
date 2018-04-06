# This program migrates data from Microsoft SQL Server db to MySQL Server db

# Python MySQL Library to 
import pymysql
# Python ODBC Library to access ODBC databases (in this case SQL)
import pyodbc

# Connect to Sql DB
msSqlconnection = pyodbc.connect("Driver={SQL Server Native Client 11.0};"
                    "Server=MS SQL SERVER NAME;"
                    "Database=MS SQL DATABASENAME;"
                    "UID=USERNAME;"
                    "PWD=PASSWORD")

# Create a Cursor object to execute queries in SQL.
mssqlDbCursor = msSqlconnection.cursor()

# Connect to MySql DB
mySqlconnection = pymysql.connect(host='MYSQL SERVER NAME',
                           user='MYSQL USERNAME',
                           password='MYSQL USER PASSWORD',
                           db='MYSQL SCHEMA NAME',
                           charset='utf8mb4',
                           cursorclass=pymysql.cursors.DictCursor)

# Create a Cursor object to execute queries in MySQL.
mysqlDbCursor = mySqlconnection.cursor()

# Query to select all state from table in SQL
mssqlDbCursor.execute("SELECT * FROM dbo.State")

print("\nMSSQL Connection Data")
# print the first and second columns      
for row in mssqlDbCursor.fetchall() :
  # print(row[0], row[1])
  selectStateDataQuery = "SELECT StateCode FROM State WHERE  StateCode = '%s'" % (row[0])
  mysqlDbCursor.execute(selectStateDataQuery)

  # Insert State only if not already
  if mysqlDbCursor.rowcount == 0:
    insertStateDataQuery = "INSERT INTO State (StateCode, StateName) VALUES ('%s', '%s')" % (row[0], row[1])
    # Query to select data from table in MySQL
    mysqlDbCursor.execute(insertStateDataQuery)

# Commit the data inserted
mySqlconnection.commit()

# Query to select all state from table in SQL
mysqlDbCursor.execute("SELECT * FROM State")

print("MySQL Data (Migrated)") 
for row in mysqlDbCursor.fetchall() :
  print(row['StateCode'], row['StateName'])