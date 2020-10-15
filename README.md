# golang_mysql_client

First setup a MySQL server and create the table customers as per this tutorial: https://www.linode.com/docs/databases/mysql/how-to-install-mysql-on-centos-7/

The code is all based on this tutorial: https://tutorialedge.net/golang/golang-mysql-tutorial/

The repo contents are explained as follows and should read in the following order:
1) README: Please read this first for instructions on how to setup your DB first.
2) connect.go: A simple example of connecting to the DB.
3) insert_into_table.go: Inserting a row into a table in the DB so that we have some data.
4) query_single_row.go: Querying a row example.
5) write_results_to_struct.go: When you interact with data you need to put that data into a data structure that makes it easy to interact with. This is showing how to do that with a row of data. 
6) sql_api.go: Finally this file shows how to serve that data back through the web server.
