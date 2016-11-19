#coding=utf-8
import MySQLdb

conn = MySQLdb.connect(host = 'localhost', user = 'root', passwd = 'Terminal@207', charset = 'utf8')
conn.select_db("yanyu")
curs = conn.cursor()

with open('major') as f2:
    for major in f2:
        major= major.rstrip('\n')
        curs.execute('insert into yanyu_major(major) values(%s)', ([major]))
        print major
conn.commit()
