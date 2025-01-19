You are given a board of N rows and M columns. Each field of the board contains a single digit (0−9).


You want to find a path consisting of four neighboring fields. Two fields are neighboring if they share a common side. Also, the fields in your path should be distinct (you can't visit the same field twice).


The four digits of your path, in the order in which you visit them, create an integer. What is the biggest integer that you can achieve in this way?


Write a function


class Solution { public int solution(int[][] Board); }


that, given the board represented as a matrix of integers consisting of N rows and M columns, returns the biggest integer that you can achieve when concatenating the values in a path of length four.


Examples:


Given the following board (N=3, M=5):
image
the function should return 9121. You can choose either of the following paths (the first field is denoted by red):
image
or
image
