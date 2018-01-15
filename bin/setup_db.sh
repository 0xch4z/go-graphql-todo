#!/usr/bin/env sh

MYSQL=(which mysql)

$MYSQL -u root -e "CREATE DATABASE IF NOT EXISTS go_graph_complex;"
