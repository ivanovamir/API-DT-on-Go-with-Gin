import psycopg2
import psycopg2.extras as extras


def execute_values(conn, df, table):

	tuples = [tuple(x) for x in df.to_numpy()]

	cols = ','.join(list(df.columns))
	# SQL query to execute
	query = "INSERT INTO %s(%s) VALUES %%s" % (table, cols)
	cursor = conn.cursor()
	try:
		extras.execute_values(cursor, query, tuples)
		conn.commit()
	except (Exception, psycopg2.DatabaseError) as error:
		print("Error: %s" % error)
		conn.rollback()
		cursor.close()
		return 1
	print("the dataframe is inserted")
	cursor.close()


conn = psycopg2.connect(
	database="shop", user='postgres', password='qwerty', host='178.21.8.81', port='5432'
)

execute_values(conn, data, 'products')