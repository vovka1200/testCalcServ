#!/bin/bash

echo "1+2=$(curl "http://127.0.0.1:8080/add?a=1&b=2")"
echo "3-2=$(curl "http://127.0.0.1:8080/sub?a=3&b=2")"
echo "3*4=$(curl "http://127.0.0.1:8080/mul?a=3&b=4")"
echo "5/2=$(curl "http://127.0.0.1:8080/div?a=5&b=2")"
echo "1/0=$(curl "http://127.0.0.1:8080/div?a=1&b=0")"