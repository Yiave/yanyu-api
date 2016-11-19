#coding=utf-8
import MySQLdb
import json

f = file('item.json')
json = json.load(f)
conn = MySQLdb.connect(host = 'localhost', user = 'root', passwd = 'Terminal@207', charset = 'utf8')
conn.select_db("yanyu")
curs = conn.cursor()

for university in json:
    school = university['school']
    subjection = university['subjection']
    attribute = university['attribute']
    self_decision = university['self_decision']
    is_graduate_school = university['is_graduate_school']
    location = university['location']
    l = len(attribute)
    if l == 1:
        attribute = attribute[0]
    elif l == 2:
        attribute = attribute[0] + ',' + attribute[1]
    else:
        attribute = ""



    if is_graduate_school == "yes":
        is_graduate_school = "1"
    else:
        is_graduate_school = "0"

    if self_decision == "yes":
        self_decision = "1"
    else:
        self_decision = "0"

    curs.execute('insert into yanyu_university(name, location, subjection, attribution, graduate_school, self_decision) values(%s, %s, %s, %s, %s, %s)', (school, location, subjection, attribute, is_graduate_school, self_decision))

    print school, subjection, attribute, self_decision, is_graduate_school, location
conn.commit()
