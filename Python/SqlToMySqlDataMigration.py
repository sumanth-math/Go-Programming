import pymysql
import pyodbc

connection = pymysql.connect(host='localhost',
                           user='mysqlappuser',
                           password='test@123',
                           db='business',
                           charset='utf8mb4',
                           cursorclass=pymysql.cursors.DictCursor)

# Create a Cursor object to execute queries.
mysqlDbCursor = connection.cursor()

# Select data from table using SQL query.
mysqlDbCursor.execute("SELECT * FROM BSEE_CustomerType")

print("Aurora MySQL Connection Data") 
# print the first and second columns      
for row in mysqlDbCursor.fetchall() :
  print(row['CustomerTypeName'])


msSqlconnection = pyodbc.connect("Driver={SQL Server Native Client 11.0};"
                    "Server=localhost;"
                    "Database=Business"
                    "Trusted_Connection=yes;")

# Create a Cursor object to execute queries.
mssqlDbCursor = msSqlconnection.cursor()

# Select data from table using SQL query.
mssqlDbCursor.execute("SELECT * FROM CustomerStatus")

print("\nMSSQL Connection Data")
# print the first and second columns      
for row in mssqlDbCursor.fetchall() :
  print(row[1])